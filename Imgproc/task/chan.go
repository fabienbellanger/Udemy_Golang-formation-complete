package task

import "github.com/fabienbellanger/Udemy_Golang-formation-complete/Imgproc/filter"

type ChanTask struct {
	dirCtx
	Filter   filter.Filter
	PoolSize int
}

type jobReq struct {
	src string
	dst string
}

func NewChanTask(srcDir, dstDir string, filter filter.Filter, poolSize int) Tasker {
	return &ChanTask{
		Filter: filter,
		dirCtx: dirCtx{
			SrcDir: srcDir,
			DstDir: dstDir,
			files:  buildFileList(srcDir),
		},
		PoolSize: poolSize,
	}
}

func (c *ChanTask) Process() error {
	// size := len(c.files)
	// jobs := make(chan jobReq, size)
	// results := make(chan string, size)

	return nil
}
