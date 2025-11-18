package asign

import (
	"context"

	"github.com/cyjaysong/asign-go/model"
)

// 创建待签署文件
func (the Client) DownloadContractFile(ctx context.Context, req *model.DownloadContractFileReqBody) (res *model.BaseRes[model.DownloadContractFileResBody], err error) {
	path := "/contract/downloadContract"
	res = new(model.BaseRes[model.DownloadContractFileResBody])
	if err = the.post(ctx, path, req, res); err != nil {
		return nil, err
	}
	return
}
