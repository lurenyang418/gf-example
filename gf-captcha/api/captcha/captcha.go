// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package captcha

import (
	"context"

	"gf-captcha/api/captcha/v1"
)

type ICaptchaV1 interface {
	CaptchaGet(ctx context.Context, req *v1.CaptchaGetReq) (res *v1.CaptchaGetRes, err error)
	CaptchaVerify(ctx context.Context, req *v1.CaptchaVerifyReq) (res *v1.CaptchaVerifyRes, err error)
}
