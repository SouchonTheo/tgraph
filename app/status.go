package app

import (
	"fmt"

	"github.com/theosouchon/tgraph/app/utils"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check the status of dependencies",
	Run:   checkDependencies,
}

func init() {
	RootCmd.AddCommand(statusCmd)
}

func checkDependencies(cmd *cobra.Command, args []string) {
	bloopInstalled := utils.IsCommandAvailable("bloop")
	if bloopInstalled {
		fmt.Println("bloop is installed.")
	} else {
		fmt.Println("bloop is not installed. Please install it.")
	}

	dotInstalled := utils.IsCommandAvailable("dot")
	if dotInstalled {
		fmt.Println("dot (Graphviz) is installed.")
	} else {
		fmt.Println("dot (Graphviz) is not installed. Please install Graphviz.")
	}
}
