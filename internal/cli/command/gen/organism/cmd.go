package organism

import (
	"github.com/charmingruby/bob/internal/cli/command/gen/organism/module"
	"github.com/charmingruby/bob/internal/filesystem"
	"github.com/spf13/cobra"
)

func SetupCMD(fs filesystem.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "org",
		Short: "Generates organisms",
	}

	cmd.AddCommand(
		module.SetupCMD(fs),
	)

	return cmd
}
