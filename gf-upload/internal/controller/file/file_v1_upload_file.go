package file

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "gf-upload/api/file/v1"
	"gf-upload/internal/service"
)

func (c *ControllerV1) UploadFile(ctx context.Context, req *v1.UploadFileReq) (res *v1.UploadFileRes, err error) {
	if req.File == nil {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "请选择要上传的文件")
	}
	r, err := service.File().UploadFile(ctx, req.File)
	if err != nil {
		return nil, err
	}

	return &v1.UploadFileRes{Url: r}, nil
}
