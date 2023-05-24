package vo

import (
	"encoding/base64"
	"helpa/src/support/smperr"
	"net/http"
)

type Image struct {
	binary []byte
	base64 string
}

func NewImage(imageBase64 string) (Image, error) {
	dec, err := base64.StdEncoding.DecodeString(imageBase64)
	if err != nil {
		return Image{binary: []byte{}, base64: ""}, smperr.BadRequest("illegal base64 data at input byte")
	}
	contentType := http.DetectContentType(dec)
	if contentType != "image/jpeg" && contentType != "image/png" {
		return Image{binary: []byte{}}, smperr.BadRequest("unsupported image format")
	}
	return Image{binary: dec, base64: imageBase64}, nil
}

func (i Image) Binary() []byte {
	return i.binary
}

func (i Image) Base64() string {
	return i.base64
}
