package fileImpl

import (
	"github.com/rodrigodmd/go-common/base/env"
	"github.com/rodrigodmd/go-common/base/file"
)

type ConfigFile struct {
	FileName string
	Model    interface{}
}

func (cf ConfigFile) FillPath() *file.File {
	return &file.File{
		AbsPath: file.FullPath(env.BaseUrl() + cf.FileName),
		Model:   cf.Model,
	}
}
