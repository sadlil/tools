package cmd

import (
	"fmt"

	"github.com/sadlil/tools/pkg/executor"
	"github.com/spf13/cobra"
)

func NewDefaultGitctlCommand() *cobra.Command {
	root := &cobra.Command{
		Use: "gitctl",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	AddFeatureCommand(root)

	return root
}

func AddFeatureCommand(r *cobra.Command) {
	f := &cobra.Command{
		Use: "feature",
		Run: func(cmd *cobra.Command, args []string) {
			exec := executor.NewCommand("git")
			exec.PipeSTD()

			fmt.Println(exec.Run("checkout", "master"))
			fmt.Println(exec.Run("pull", "origin", "master"))
			fmt.Println(exec.Run("checkout", "-b", "srhythom/"+args[0]))
		},
	}

	r.AddCommand(f)
}
