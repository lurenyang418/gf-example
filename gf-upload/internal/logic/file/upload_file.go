package file

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (s *sFile) UploadFile(ctx context.Context, file *ghttp.UploadFile) (string, error) {
	filename := file.Filename
	// err
	file.Save("upload")
	return filename, nil
}
