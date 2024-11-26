package molecule

import (
	"fmt"

	"github.com/charmingruby/bob/internal/command/generate/atom"
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/component/input"
	"github.com/charmingruby/bob/internal/command/shared/component/structure"
	"github.com/charmingruby/bob/internal/command/shared/fs"
	"github.com/spf13/cobra"
)

func RunService(m component.Manager) *cobra.Command {
	var (
		module string
		repo   string
	)

	cmd := &cobra.Command{
		Use:   "service",
		Short: "Generates a new service molecule",
		Run: func(cmd *cobra.Command, args []string) {
			if err := fs.GenerateNestedDirectories(
				fmt.Sprintf("%s/%s", m.SourceDirectory, module),
				[]string{"core", "service"},
			); err != nil {
				panic(err)
			}

			if err := input.ValidateOnlyModuleCommandInput(module); err != nil {
				panic(err)
			}

			MakeService(m, repo, module)
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module name")
	cmd.Flags().StringVarP(&repo, "repo", "r", "", "repository dependency")

	return cmd
}

func MakeService(m component.Manager, repo string, module string) {
	hasRepo := repo != ""

	if !hasRepo {
		if err := fs.GenerateFile(makeServiceRegistryAtomComponent(
			m,
			module,
		)); err != nil {
			panic(err)
		}
	} else {
		if err := fs.GenerateFile(makeServiceRegistryWithRepositoryAtomComponent(
			m,
			module,
			repo,
		)); err != nil {
			panic(err)
		}
	}

	sampleActor := module

	if err := fs.GenerateFile(atom.MakeServiceComponent(m.SourceDirectory, module, sampleActor)); err != nil {
		panic(err)
	}
}

func makeServiceRegistryWithRepositoryAtomComponent(m component.Manager, module, name string) fs.File {
	return atom.MakeRegistryComponent(atom.RegistryParams{
		Module:       module,
		TemplateName: "service_registry_with_repository",
		TemplateData: structure.DependentPackage{
			SourcePath: m.DependencyPath(module),
			Module:     module,
			Name:       name,
		},
		RegistryName:         "service",
		DestinationDirectory: m.AppendToModuleDirectory(module, "/core/service"),
	})
}

type serviceRegistryAtomData struct {
	Module string
}

func makeServiceRegistryAtomComponent(m component.Manager, module string) fs.File {
	return atom.MakeRegistryComponent(atom.RegistryParams{
		Module:       module,
		TemplateName: "service_registry",
		TemplateData: serviceRegistryAtomData{
			Module: module,
		},
		RegistryName:         "service",
		DestinationDirectory: m.AppendToModuleDirectory(module, "/core/service"),
	})
}

func ServicePath() string {
	return "core/service"
}
