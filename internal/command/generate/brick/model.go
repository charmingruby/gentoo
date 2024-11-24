package brick

import (
	"fmt"

	"github.com/charmingruby/bob/internal/command/shared/constant"
	"github.com/charmingruby/bob/internal/command/shared/fs"
	"github.com/charmingruby/bob/internal/command/shared/validator/input"
	"github.com/spf13/cobra"
)

func RunModel(destinationDirectory string) *cobra.Command {
	var (
		module string
		name   string
	)

	cmd := &cobra.Command{
		Use:   "model",
		Short: "Generates a new model",
		Run: func(cmd *cobra.Command, args []string) {
			if err := input.ValidateDefaultCommandInput(module, name); err != nil {
				panic(err)
			}

			if err := fs.GenerateFile(makeModelComponent(
				destinationDirectory,
				module,
				name,
			)); err != nil {
				panic(err)
			}
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module name")
	cmd.Flags().StringVarP(&name, "name", "n", "", "model name")

	return cmd
}

func makeModelComponent(destinationDirectory, module, name string) fs.File {
	component := *New(ComponentInput{
		Directory: fmt.Sprintf("%s/%s/core/%s",
			destinationDirectory,
			module,
			constant.MODEL_TEMPLATE,
		),
		Module:  module,
		Name:    name,
		Suffix:  "",
		HasTest: true,
	}, WithDefaultTemplate())

	file := fs.File{
		CommandType:          constant.GENERATE_COMMAND,
		TemplateName:         constant.MODEL_TEMPLATE,
		TemplateData:         component.Data,
		FileName:             component.Name,
		FileSuffix:           "",
		DestinationDirectory: component.Directory,
		HasTest:              component.HasTest,
	}

	return file
}