package model

// 通过网络文件地址上传 Req https://web.asign.cn/platform/openDoc/docDetail?mid=uploadFileUrl
type UploadFileByUrlReqBody struct {
	FileUrl        string `json:"fileurl" dc:"网络文件URL，爱签平台将从该地址拉取文档fileUrl需要进行编码。"`
	Filename       string `json:"filename" dc:"指定文件名称，必须包含扩展名。扩展名需要与上传文件扩展名一致"`
	Extension      string `json:"extension" dc:"文件扩展名，支持文件类型doc、docx、pdf、xls、xlsx、png、jpg、jpeg"`
	FileType       int    `json:"fileType" dc:"文件类型, 1：合同附件（默认）2：印章图片 3：模板附件"`
	StoreCloudSign string `json:"storeCloudSign" dc:"私有云项目标识（只有合同附件才支持私有云）"`
}

// 通过网络文件地址上传 Res
type UploadFileByUrlResBody string // 本次文件上传对应的文件标识

// 下载合同文件 Req https://web.asign.cn/platform/openDoc/docDetail?mid=downloadContract
type DownloadContractFileReqBody struct {
	ContractNo       string `json:"contractNo" dc:"合同唯一编码"`
	Outfile          string `json:"outfile,omitempty" dc:"文件本地路径"`
	Force            int    `json:"force,omitempty" dc:"强制下载标识,0(默认):未签署完的无法下载,1:无论什么状态都强制下载"`
	DownloadFileType int    `json:"downloadFileType,omitempty" dc:"下载文件类型，多附件下载"`
}

// 下载合同文件 Res
type DownloadContractFileResBody struct {
	FileName string `json:"fileName" dc:"文件名"`
	Md5      string `json:"md5" dc:"文件md5值"`
	FileType int    `json:"fileType" dc:"文件md5值"`
	Size     int    `json:"size" dc:"文件大小"`
	Data     string `json:"data" dc:"Base64字符串"`
}
