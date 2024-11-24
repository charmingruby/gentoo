package brick

import (
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/constant"
	"github.com/charmingruby/bob/internal/command/shared/fs"
	"github.com/charmingruby/bob/internal/command/shared/validator/input"
	"github.com/spf13/cobra"
)

func RunRepository(m component.Manager) *cobra.Command {
	var (
		module string
		name   string
	)

	cmd := &cobra.Command{
		Use:   "repository",
		Short: "Generates a new repository",
		Run: func(cmd *cobra.Command, args []string) {
			if err := input.ValidateDefaultCommandInput(module, name); err != nil {
				panic(err)
			}

			if err := fs.GenerateFile(MakeRepositoryComponent(
				m.SourceDirectory,
				module,
				name,
				m.DependencyPath(module),
			)); err != nil {
				panic(err)
			}
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module name")
	cmd.Flags().StringVarP(&name, "name", "n", "", "model to be managed by the repository")

	return cmd
}

func MakeRepositoryComponent(directory, module, name, dependencyPath string) fs.File {
	component := *New(ComponentInput{
		Directory: component.ModulePath(directory, module, RepositoryPath()),
		Module:    module,
		Name:      name,
		Suffix:    "repository",
		HasTest:   false,
	}, WithModuleDependenciesTemplate(dependencyPath))

	return NewFileFromBrick(component, constant.REPOSITORY_TEMPLATE)
}

func RepositoryPath() string {
	return "core/repository"
}
