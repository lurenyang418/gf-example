package captcha

import (
	"gf-captcha/internal/service"
	"gf-captcha/internal/util"
	"image/color"

	"github.com/mojocn/base64Captcha"
)

type sCaptcha struct {
	captcha *base64Captcha.Captcha
}

func init() {
	service.RegisterCaptcha(New())
}

func New() *sCaptcha {
	driverString := base64Captcha.DriverString{
		Height:          40,
		Width:           100,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          4,
		Source:          "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
		BgColor:         &color.RGBA{R: 3, G: 102, B: 214, A: 125},
		Fonts:           []string{"wqy-microhei.ttc"},
	}
	captcha := base64Captcha.NewCaptcha(
		driverString.ConvertFonts(),
		util.NewMiniRedisStore("captcha:"),
	)
	return &sCaptcha{
		captcha: captcha,
	}
}
