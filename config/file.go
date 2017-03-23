package config

import "os"

const (
	codeFileTmpDirEvn = "CODE_FILE_TEMP_DIR_EVN"
)

var (
	CodeFileTmpDir string
)

func init() {
	CodeFileTmpDir = os.Getenv(codeFileTmpDirEvn)
	if CodeFileTmpDir == "" {
		CodeFileTmpDir = `/Users/tang/Documents/test`
	}
}
