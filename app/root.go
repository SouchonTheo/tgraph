package app

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
	"github.com/theosouchon/tgraph/app/utils/text"
)

var RootCmd = &cobra.Command{
	Use:     "tgraph",
	Short:   "Tool to generate dependency graph from bloop",
	Run:     runCommand,
	Example: text.Example,
}

var project string
var output string
var open string
var sedCommands []string

func initConfig() {

}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVarP(&project, "project", "n", "projects", "Name of your project")
	RootCmd.PersistentFlags().StringVarP(&output, "output", "f", "", "Output file")
	RootCmd.PersistentFlags().StringVarP(&open, "open", "o", "", "Application to open the graph")
	RootCmd.PersistentFlags().StringArrayVarP(&sedCommands, "ignore", "i", nil, "Folders to ignore")
}

func runCommand(cmd *cobra.Command, args []string) {
	command := fmt.Sprintf("bloop %s --dot-graph", project)

	if len(sedCommands) > 0 {
		command += fmt.Sprintf(" | %s", buildSedExpression(sedCommands))
	}

	command += generateOutput()

	cmdExec := exec.Command("sh", "-c", command)
	err := cmdExec.Run()

	if err != nil {
		fmt.Printf("Erreur lors de l'exécution de la commande : %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Graph généré avec succès dans %s pour le projet %s.\n", output, project)
}

func buildSedExpression(commands []string) string {
	var sedCommands []string
	for _, cmd := range commands {
		parts := strings.SplitN(cmd, " ", 2)
		sedCommands = append(sedCommands, fmt.Sprintf(`sed "/-%s/d"`, parts[0]))
	}
	return strings.Join(sedCommands, " | ")
}

func generateOutput() string {
	var command string = ""
	if output == "" {
		if open != "" {
			command += fmt.Sprintf(" | dot -Tsvg -o %s", open)
		} else {
			command += chooseApp()
		}
	} else {
		if strings.HasSuffix(output, ".svg") {
			command += fmt.Sprintf(" | dot -Tsvg -o %s", output)
		} else if strings.HasSuffix(output, ".png") {
			command += fmt.Sprintf(" | dot -Tpng -o %s", output)
		}
	}
	return command
}

func chooseApp() string {
	switch runtime.GOOS {
	case "darwin":
		return "| dot -Tpng | open -f -a Preview"
	case "windows":
		return "| dot -Tpng | start"
	default:
		return "| dot -Tpng | xdg-open"
	}
}
