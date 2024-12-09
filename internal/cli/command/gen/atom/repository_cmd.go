package atom

import (
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/component/atom"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/spf13/cobra"
)

func RunRepository(m filesystem.Manager) *cobra.Command {
	var (
		module   string
		name     string
		database string
	)

	cmd := &cobra.Command{
		Use:   "repository",
		Short: "Generates a new repository",
		Run: func(cmd *cobra.Command, args []string) {
			if err := ValidateRepositoryCommandInput(module, name, database); err != nil {
				panic(err)
			}

			repository := atom.MakeRepository(m, module, name, database)
			if err := m.GenerateFile(repository); err != nil {
				panic(err)
			}

			if database != "" {
				persistenceRepository := atom.MakePersistenceRepository(m, module, name, database)
				if err := m.GenerateFile(persistenceRepository); err != nil {
					panic(err)
				}
			}
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module name")
	cmd.Flags().StringVarP(&name, "name", "n", "", "model to be managed by the repository")
	cmd.Flags().StringVarP(&database, "database", "d", "", "database to implement repository, if it is not set, it will be not created")

	return cmd
}

func ValidateRepositoryCommandInput(module, name, database string) error {
	args := []input.Arg{
		{
			FieldName:  "module",
			Value:      module,
			IsRequired: true,
		},
		{
			FieldName:  "name",
			Value:      name,
			IsRequired: true,
		},
		{
			FieldName: "database",
			Value:     database,
		},
	}

	if err := input.ValidateArgsList(args); err != nil {
		return err
	}

	return nil
}
