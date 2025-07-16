package captcha

import (
	"context"

	"gf-captcha/internal/model"

	"github.com/gogf/gf/v2/os/glog"
)

func (s *sCaptcha) CaptchaGet(ctx context.Context) (*model.CaptchaGetResp, error) {
	id, b64s, _, err := s.captcha.Generate()
	glog.Info(ctx, "captcha id:", id, "b64s:", b64s)
	if err != nil {
		glog.Error(ctx, err, 12312)
		return nil, err
	}

	return &model.CaptchaGetResp{
		CaptchaId:    id,
		CaptchaImage: b64s,
	}, nil
}
