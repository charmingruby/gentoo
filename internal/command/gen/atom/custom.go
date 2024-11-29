package atom

import (
	"github.com/charmingruby/bob/internal/command/shared/component/constant"
	"github.com/charmingruby/bob/internal/command/shared/filesystem"
)

type CustomComponentInput struct {
	BaseComponent Component
	TemplateName  string
	TemplateData  any
	FileName      string
	FileSuffix    string
}

func MakeCustomComponent(in CustomComponentInput) filesystem.File {
	file := filesystem.File{
		CommandType:          constant.GENERATE_COMMAND,
		TemplateName:         in.TemplateName,
		TemplateData:         in.TemplateData,
		FileName:             in.FileName,
		FileSuffix:           in.FileSuffix,
		DestinationDirectory: in.BaseComponent.DestinationDirectory,
		HasTest:              in.BaseComponent.HasTest,
	}

	return file
}
