package task

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Tasker interface {
	Process() error
}

type dirCtx struct {
	SrcDir string
	DstDir string
	files  []string
}

func buildFileList(srcDir string) []string {
	files := []string{}
	fmt.Println("generating file list...")
	filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() || (!strings.HasSuffix(path, ".jpg") && !strings.HasSuffix(path, ".jpeg")) {
			return nil
		}

		files = append(files, path)
		return nil
	})
	return files
}
