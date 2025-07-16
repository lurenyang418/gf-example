package model

type CaptchaGetResp struct {
	CaptchaId    string `json:"captchaId"`
	CaptchaImage string `json:"captchaImage" dc:"base64 encoded captcha image"`
}

type CaptchaVerifyReq struct {
	CaptchaId string `json:"captchaId" v:"required#Captcha ID is required"`
	Answer    string `json:"answer"    v:"required#Answer is required"`
}
