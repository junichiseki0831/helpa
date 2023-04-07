package domain_test

import (
	_ "embed"
	"encoding/base64"
	domain "helpa/src/core/domain/userdm"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

//go:embed testdata/sample.png
var sampleImagePng []byte

func TestNewUser(t *testing.T) {
	for _, tt := range []struct {
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
		isErr        bool
	}{
		{
			title:        "正常系",
			id:           "123",
			name:         "testName",
			password:     "12345671",
			email:        "test1@test.com",
			introduction: "testIntroduction",
			note:         "testNote",
			image:        sampleImagePng,
			createdAt:    time.Now(),
			updatedAt:    time.Now(),
		},
		{
			title:        "異常系: idの引数空",
			id:           "",
			name:         "testName",
			password:     "12345672",
			email:        "test2@test.com",
			introduction: "testIntroduction",
			note:         "testNote",
			image:        sampleImagePng,
			createdAt:    time.Now(),
			updatedAt:    time.Now(),
			isErr:        true,
		},
		{
			title:        "異常系: パスワードが8文字以下",
			id:           "124",
			name:         "testName",
			password:     "123",
			email:        "test3@test.com",
			introduction: "testIntroduction",
			note:         "testNote",
			image:        sampleImagePng,
			createdAt:    time.Now(),
			updatedAt:    time.Now(),
			isErr:        true,
		},
		{
			title:        "異常系: emailが@なし",
			id:           "125",
			name:         "testName",
			password:     "12345673",
			email:        "test4com",
			introduction: "testIntroduction",
			note:         "testNote",
			image:        sampleImagePng,
			createdAt:    time.Now(),
			updatedAt:    time.Now(),
			isErr:        true,
		},
		{
			title:        "異常系: createdAtの値が空",
			id:           "126",
			name:         "testName",
			password:     "12345674",
			email:        "test5@test.com",
			introduction: "testIntroduction",
			note:         "testNote",
			image:        sampleImagePng,
			createdAt:    time.Time{},
			updatedAt:    time.Now(),
			isErr:        true,
		},
		{
			title:        "異常系: updatedAtの値が空",
			id:           "127",
			name:         "testName",
			password:     "12345675",
			email:        "test6@test.com",
			introduction: "testIntroduction",
			note:         "testNote",
			image:        sampleImagePng,
			createdAt:    time.Now(),
			updatedAt:    time.Time{},
			isErr:        true,
		},
		{
			title:        "異常系: imageの値が空",
			id:           "123",
			name:         "testName",
			password:     "12345671",
			email:        "test1@test.com",
			introduction: "testIntroduction",
			note:         "testNote",
			image:        []byte{},
			createdAt:    time.Now(),
			updatedAt:    time.Now(),
			isErr:        true,
		},
	} {
		t.Run(tt.title, func(t *testing.T) {
			base64String := base64.StdEncoding.EncodeToString(tt.image)
			got, err := domain.GenForTest(tt.id, tt.name, tt.password, tt.email, tt.introduction, tt.note, base64String, tt.createdAt, tt.updatedAt)
			if tt.isErr {
				assert.NotNil(t, err)
				assert.Empty(t, got)
				return
			}
			assert.Equal(t, tt.id, got.ID().String())
			assert.Equal(t, tt.name, got.Name())
			assert.Equal(t, tt.password, got.Password().String())
			assert.Equal(t, tt.email, got.Email().String())
			assert.Equal(t, tt.introduction, got.Introduction())
			assert.Equal(t, tt.note, got.Note())
			assert.Equal(t, tt.image, got.Image().Binary())
			assert.True(t, got.CreatedAt().Value().Equal(tt.createdAt))
			assert.True(t, got.UpdatedAt().Value().Equal(tt.updatedAt))
			assert.Nil(t, err)

		})
	}
}
