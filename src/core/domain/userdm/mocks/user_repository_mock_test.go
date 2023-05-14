package mock_domain_test

import (
	context "context"
	_ "embed"
	"encoding/base64"
	domain "helpa/src/core/domain/userdm"
	mock_domain "helpa/src/core/domain/userdm/mocks"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

//go:embed testdata/sample.png
var sampleImagePng []byte

func TestMockUserRepository_FindByName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_domain.NewMockUserRepository(ctrl)

	type args struct {
		title        string
		id           string
		name         string
		password     string
		email        string
		introduction string
		note         string
		image        []byte
		createdAt    time.Time
		updatedAt    time.Time
	}
	for _, tt := range []struct {
		args  args
		isErr bool
	}{
		{
			args: args{
				title:        "正常系",
				id:           "123",
				name:         "testUser",
				password:     "12345671",
				email:        "test1@test.com",
				introduction: "testIntroduction",
				note:         "testNote",
				image:        sampleImagePng,
				createdAt:    time.Now(),
				updatedAt:    time.Now(),
			},
			isErr: false,
		},
		{
			args: args{
				title:        "異常系",
				id:           "12900000",
				name:         "",
				password:     "12345672",
				email:        "test2@test.com",
				introduction: "testIntroduction",
				note:         "testNote",
				image:        sampleImagePng,
				createdAt:    time.Now(),
				updatedAt:    time.Now(),
			},
			isErr: true,
		},
	} {
		tt := tt
		t.Run(tt.args.title, func(t *testing.T) {
			t.Parallel()
			base64String := base64.StdEncoding.EncodeToString(tt.args.image)
			got, err := domain.GenForTest(tt.args.id, tt.args.name, tt.args.password, tt.args.email, tt.args.introduction, tt.args.note, base64String, tt.args.createdAt, tt.args.updatedAt)
			if tt.isErr {
				assert.NotNil(t, err)
				assert.Empty(t, got)
				return
			}

			mockRepo.EXPECT().FindByName(context.Background(), tt.args.name).Return([]domain.User{*got}, err)
			users, err := mockRepo.FindByName(context.Background(), tt.args.name)
			assert.Equal(t, tt.args.name, users[0].Name())
			assert.Nil(t, err)
		})
	}
}
