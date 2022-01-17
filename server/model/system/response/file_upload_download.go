package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/system"

type ExaFileResponse struct {
	File system.ExaFileUploadAndDownload `json:"file"`
}
