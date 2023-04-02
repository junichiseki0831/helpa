package vo

import (
	"encoding/base64"
	"errors"
	"net/http"
)

type Image struct {
	binary []byte
}

func NewImage(imageBase64 string) (Image, error) {
	dec, err := base64.StdEncoding.DecodeString(imageBase64)
	if err != nil {
		return Image{binary: []byte{}}, errors.New("decode error")
	}
	contentType := http.DetectContentType(dec)
	if contentType != "image/jpeg" && contentType != "image/png" {
		return Image{binary: []byte{}}, errors.New("unsupported image format")
	}
	return Image{binary: dec}, nil
}

func (i *Image) Binary() []byte {
	return i.binary
}
