package files

import "os"

type (
	FileInfo struct {
		Data     []byte `json:"data"`
		FileName string `json:"fileName"`
		FileSize int64  `json:"fileSize"`
	}

	SaveFileOut struct {
		Error    error
		FilePath string
		FileInfo os.FileInfo
	}
)
