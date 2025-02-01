package util

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/dchest/captcha"
	"os"
	"testing"
)

var imageFile *os.File
var audioFile *os.File
var _0to9 = []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

func init() {
	var err error
	imageFile, err = OSOpenFileWithCreate("testdata/captcha/img.jpeg")
	Must(err)
	audioFile, err = OSOpenFileWithCreate("testdata/captcha/aud.wav")
	Must(err)
}

func TestCaptcha(t *testing.T) {
	var err error
	_, err = captcha.NewImage("1", _0to9, 60, 25).
		WriteTo(imageFile)
	Must(err)
	_, err = captcha.NewAudio("1", _0to9, "zh").
		WriteTo(audioFile)
	Must(err)
}

func randByte(n int) []byte {
	return s2b(gofakeit.Password(false, false, true, false, false, n))
}
