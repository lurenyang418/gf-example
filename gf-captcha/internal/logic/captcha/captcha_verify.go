package captcha

import (
	"context"

	"gf-captcha/internal/model"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (s *sCaptcha) CaptchaVerify(ctx context.Context, req *model.CaptchaVerifyReq) error {
	if ok := s.captcha.Verify(req.CaptchaId, req.Answer, true); !ok {
		return gerror.NewCode(gcode.CodeInvalidParameter, "captcha verification failed")
	}

	return nil
}
