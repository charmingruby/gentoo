package brick

import (
	"github.com/charmingruby/bob/internal/command/shared/constant"
	"github.com/charmingruby/bob/internal/command/shared/formatter"
	"github.com/charmingruby/bob/internal/command/shared/fs"
)

type Component struct {
	Directory string
	Module    string
	Name      string
	Suffix    string
	Data      any
	HasTest   bool
}

type ComponentInput struct {
	Module    string
	Name      string
	Suffix    string
	Directory string
	HasTest   bool
}

type ComponentOption func(*Component)

func New(in ComponentInput, opts ...ComponentOption) *Component {
	component := &Component{
		Module:    in.Module,
		Name:      in.Name,
		Directory: in.Directory,
		Suffix:    in.Suffix,
		HasTest:   in.HasTest,
	}

	component.format()

	for _, opt := range opts {
		opt(component)
	}

	return component
}

type DefaultTemplateParams struct {
	Name string
}

func WithDefaultTemplate() ComponentOption {
	return func(s *Component) {
		s.Data = DefaultTemplateParams{
			Name: s.Name,
		}
	}
}

type ModuleDependenciesTemplateParams struct {
	SourcePath string
	Module     string
	Name       string
}

func WithModuleDependenciesTemplate(sourcePath string) ComponentOption {
	println(sourcePath)
	return func(s *Component) {
		s.Data = ModuleDependenciesTemplateParams{
			Module:     s.Module,
			SourcePath: sourcePath,
			Name:       s.Name,
		}
	}
}

func (r *Component) format() {
	r.Name = formatter.ToCamelCase(r.Name)

	if r.Suffix != "" {
		r.Suffix = formatter.ToCamelCase(r.Suffix)
	}
}

func NewFileFromBrick(component Component, template string) fs.File {
	return fs.File{
		CommandType:          constant.GENERATE_COMMAND,
		TemplateName:         template,
		TemplateData:         component.Data,
		FileName:             component.Name,
		FileSuffix:           component.Suffix,
		DestinationDirectory: component.Directory,
		HasTest:              component.HasTest,
	}
}
