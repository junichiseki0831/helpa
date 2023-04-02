package vo_test

import (
	_ "embed"
	"encoding/base64"
	"helpa/src/core/domain/shared/vo"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed sample/sample.png
var sampleImagePng []byte

//go:embed sample/sample.jpg
var sampleImagejpg []byte

//go:embed sample/sample.gif
var sampleImageGif []byte

func TestNewImage(t *testing.T) {
	for _, tt := range []struct {
		name  string
		input []byte
		isErr bool
	}{
		{
			name:  "正常系1",
			input: sampleImagePng,
		},
		{
			name:  "正常系2",
			input: sampleImagejpg,
		},
		{
			name:  "異常系1",
			input: []byte{123},
			isErr: true,
		},
		{
			name:  "異常系2",
			input: sampleImageGif,
			isErr: true,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			base64String := base64.StdEncoding.EncodeToString(tt.input)
			got, err := vo.NewImage(base64String)
			if tt.isErr {
				assert.NotNil(t, err)
				return
			}
			assert.Nil(t, err)
			assert.Equal(t, tt.input, got.Binary())
		})
	}
}
