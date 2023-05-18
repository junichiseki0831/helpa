package mock_domain_test

import (
	context "context"
	_ "embed"
	"encoding/base64"
	"errors"
	app "helpa/src/core/app/userapp"
	mock_domain "helpa/src/core/domain/userdm/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

//go:embed testdata/sample.png
var sampleImagePng []byte

func TestMockUserRepository_Store(t *testing.T) {

	type args struct {
		name         string
		password     string
		email        string
		introduction string
		note         string
		image        []byte
	}

	ctrl := gomock.NewController(t)
	mockRepo := mock_domain.NewMockUserRepository(ctrl)

	for _, tt := range []struct {
		args         args
		title        string
		mockRepoFunc func(repo *mock_domain.MockUserRepository)
		isErr        bool
	}{
		{
			args: args{
				name:         "testUser",
				password:     "12345671",
				email:        "test1@test.com",
				introduction: "testIntroduction",
				note:         "testNote",
				image:        sampleImagePng,
			},
			title: "正常系",
			mockRepoFunc: func(repo *mock_domain.MockUserRepository) {
				repo.EXPECT().Store(context.Background(), gomock.Any()).Return(nil)
			},
			isErr: false,
		},
		{
			args: args{
				name:         "testUser",
				password:     "12345671",
				email:        "test1@test.com",
				introduction: "testIntroduction",
				note:         "testNote",
				image:        sampleImagePng,
			},
			title: "異常系: Storeがエラーを返す",
			mockRepoFunc: func(repo *mock_domain.MockUserRepository) {
				repo.EXPECT().Store(context.Background(), gomock.Any()).Return(errors.New("some error"))
			},
			isErr: true,
		},
		{
			args: args{
				name:         "testUser",
				password:     "123",
				email:        "test1@test.com",
				introduction: "testIntroduction",
				note:         "testNote",
				image:        sampleImagePng,
			},
			title:        "異常系: entityエラー（パスワードが8文字以下）",
			mockRepoFunc: nil,
			isErr:        true,
		},
	} {
		tt := tt
		t.Run(tt.title, func(t *testing.T) {
			t.Parallel()
			base64String := base64.StdEncoding.EncodeToString(tt.args.image)
			request := app.CreateUserRequest{
				Name:         tt.args.name,
				Password:     tt.args.password,
				Email:        tt.args.email,
				Introduction: tt.args.introduction,
				Note:         tt.args.note,
				Image:        base64String,
			}

			if tt.mockRepoFunc != nil {
				tt.mockRepoFunc(mockRepo)
				err := app.NewCreateUserAppService(mockRepo).Exec(context.Background(), &request)
				if tt.isErr {
					assert.Error(t, err)
					return
				}
				assert.Nil(t, err)
			} else {
				err := app.NewCreateUserAppService(mockRepo).Exec(context.Background(), &request)
				if tt.isErr {
					assert.Error(t, err)
					return
				}
				assert.Nil(t, err)
			}

		})
	}
}
