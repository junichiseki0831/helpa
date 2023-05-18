package mock_domain_test

import (
	context "context"
	"errors"
	app "helpa/src/core/app/userapp"
	mock_domain "helpa/src/core/domain/userdm/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestMockUserRepository_Store(t *testing.T) {

	type args struct {
		name         string
		password     string
		email        string
		introduction string
		note         string
		image        string
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
				image:        "iVBORw0KGgoAAAANSUhEUgAAASwAAADIBAMAAACg8cFmAAAAG1BMVEUAAAD///8fHx9fX18/Pz+fn5/f399/f3+/v7/RRZlbAAAACXBIWXMAAA7EAAAOxAGVKw4bAAACrUlEQVR4nO3Xv2/aQBjGcfdsQsZcgMBo8oN2BJEho02Tdo3bqOoYojZdEynqDBmi/tm9e23wXRyYjkzfj4T9gA336vVxhigCAAAAAAAAAAAAAAAAAAAAAAAA4FPTaVrFz+fN1HQynVUpeSMFcqW1PnqQ+gqtB6mfmhJzTD9LPHsjBdIaPKXqYmBj3p+pvx0/NS3/zNRXndkC9VP6cT72UihtO0D0eG1adJSazTJz0xtu7GZ0aDbDO7P51PVSWPsH5nEoA966aZO4bzbL1GzUvZfCaptC8mtJHTdtZEpIepIex04KLDZlFanEgZs2WppZKT2NPiycFFjrQBpQDVgnq10WGWfuG4oo2ruVZK5/nQIbmstWXonoMXWSHLuT3WThvuHFNEcutb3+dQpL3adR0i/zZFwn2SX2exnF/iU1z3JZ7KK446SgkuK5/tB8XKdqb9vlN8uesqq646SAfr5ouxbF1bIzfKhTVbVp16tm2dk0ySSqvpNClvVv/jvbVpZtl98s6c+OyzKfd9HbWlZydOI3S9mnOy/LDHG9rawony+80/cO36esdnc9YfNx8mrKm8M69U6X22W+nuj5Lqa8pXr1ApHVaXV4srxzz26VPzOqZaHrpNAGm5dTu2aVa9dKIWXsfjmVsjbdfGTNyp12tcpuvsPNxzaobI7ykrBrltuuopxKVXPMfatOgcWdLT9sZM2q27Vfvf4OP2zsN35fZuxw4SYpWdasdbvUcjW6fB9lCatTKGfpaih1bz5czTM3OeXlVZWjX6u3Shp1vRTKaHB+fFLIJP7Sm50WXT8ZcblLUtkp/X1qmWfJ/Ob4qvyzsU7BXJj/Uj0ZUv1opoZEl2wNZ3Otv8mrdQomuVz/8zy9bKZt1GXWSAAAAAAAAAAAAAAAAAAAAAAAAAB8/wEsNXDTxGtkyAAAAABJRU5ErkJggg==",
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
				image:        "iVBORw0KGgoAAAANSUhEUgAAASwAAADIBAMAAACg8cFmAAAAG1BMVEUAAAD///8fHx9fX18/Pz+fn5/f399/f3+/v7/RRZlbAAAACXBIWXMAAA7EAAAOxAGVKw4bAAACrUlEQVR4nO3Xv2/aQBjGcfdsQsZcgMBo8oN2BJEho02Tdo3bqOoYojZdEynqDBmi/tm9e23wXRyYjkzfj4T9gA336vVxhigCAAAAAAAAAAAAAAAAAAAAAAAA4FPTaVrFz+fN1HQynVUpeSMFcqW1PnqQ+gqtB6mfmhJzTD9LPHsjBdIaPKXqYmBj3p+pvx0/NS3/zNRXndkC9VP6cT72UihtO0D0eG1adJSazTJz0xtu7GZ0aDbDO7P51PVSWPsH5nEoA966aZO4bzbL1GzUvZfCaptC8mtJHTdtZEpIepIex04KLDZlFanEgZs2WppZKT2NPiycFFjrQBpQDVgnq10WGWfuG4oo2ruVZK5/nQIbmstWXonoMXWSHLuT3WThvuHFNEcutb3+dQpL3adR0i/zZFwn2SX2exnF/iU1z3JZ7KK446SgkuK5/tB8XKdqb9vlN8uesqq646SAfr5ouxbF1bIzfKhTVbVp16tm2dk0ySSqvpNClvVv/jvbVpZtl98s6c+OyzKfd9HbWlZydOI3S9mnOy/LDHG9rawony+80/cO36esdnc9YfNx8mrKm8M69U6X22W+nuj5Lqa8pXr1ApHVaXV4srxzz26VPzOqZaHrpNAGm5dTu2aVa9dKIWXsfjmVsjbdfGTNyp12tcpuvsPNxzaobI7ykrBrltuuopxKVXPMfatOgcWdLT9sZM2q27Vfvf4OP2zsN35fZuxw4SYpWdasdbvUcjW6fB9lCatTKGfpaih1bz5czTM3OeXlVZWjX6u3Shp1vRTKaHB+fFLIJP7Sm50WXT8ZcblLUtkp/X1qmWfJ/Ob4qvyzsU7BXJj/Uj0ZUv1opoZEl2wNZ3Otv8mrdQomuVz/8zy9bKZt1GXWSAAAAAAAAAAAAAAAAAAAAAAAAAB8/wEsNXDTxGtkyAAAAABJRU5ErkJggg==",
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
				image:        "iVBORw0KGgoAAAANSUhEUgAAASwAAADIBAMAAACg8cFmAAAAG1BMVEUAAAD///8fHx9fX18/Pz+fn5/f399/f3+/v7/RRZlbAAAACXBIWXMAAA7EAAAOxAGVKw4bAAACrUlEQVR4nO3Xv2/aQBjGcfdsQsZcgMBo8oN2BJEho02Tdo3bqOoYojZdEynqDBmi/tm9e23wXRyYjkzfj4T9gA336vVxhigCAAAAAAAAAAAAAAAAAAAAAAAA4FPTaVrFz+fN1HQynVUpeSMFcqW1PnqQ+gqtB6mfmhJzTD9LPHsjBdIaPKXqYmBj3p+pvx0/NS3/zNRXndkC9VP6cT72UihtO0D0eG1adJSazTJz0xtu7GZ0aDbDO7P51PVSWPsH5nEoA966aZO4bzbL1GzUvZfCaptC8mtJHTdtZEpIepIex04KLDZlFanEgZs2WppZKT2NPiycFFjrQBpQDVgnq10WGWfuG4oo2ruVZK5/nQIbmstWXonoMXWSHLuT3WThvuHFNEcutb3+dQpL3adR0i/zZFwn2SX2exnF/iU1z3JZ7KK446SgkuK5/tB8XKdqb9vlN8uesqq646SAfr5ouxbF1bIzfKhTVbVp16tm2dk0ySSqvpNClvVv/jvbVpZtl98s6c+OyzKfd9HbWlZydOI3S9mnOy/LDHG9rawony+80/cO36esdnc9YfNx8mrKm8M69U6X22W+nuj5Lqa8pXr1ApHVaXV4srxzz26VPzOqZaHrpNAGm5dTu2aVa9dKIWXsfjmVsjbdfGTNyp12tcpuvsPNxzaobI7ykrBrltuuopxKVXPMfatOgcWdLT9sZM2q27Vfvf4OP2zsN35fZuxw4SYpWdasdbvUcjW6fB9lCatTKGfpaih1bz5czTM3OeXlVZWjX6u3Shp1vRTKaHB+fFLIJP7Sm50WXT8ZcblLUtkp/X1qmWfJ/Ob4qvyzsU7BXJj/Uj0ZUv1opoZEl2wNZ3Otv8mrdQomuVz/8zy9bKZt1GXWSAAAAAAAAAAAAAAAAAAAAAAAAAB8/wEsNXDTxGtkyAAAAABJRU5ErkJggg==",
			},
			title:        "異常系: entityエラー（パスワードが8文字以下）",
			mockRepoFunc: nil,
			isErr:        true,
		},
	} {
		tt := tt
		t.Run(tt.title, func(t *testing.T) {
			t.Parallel()
			request := app.CreateUserRequest{
				Name:         tt.args.name,
				Password:     tt.args.password,
				Email:        tt.args.email,
				Introduction: tt.args.introduction,
				Note:         tt.args.note,
				Image:        tt.args.image,
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
