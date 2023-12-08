package text

import (
	"github.com/fatih/color"
)

var Example = `  Generate a ` + color.GreenString("dependency graph") + ` for the project ` + color.CyanString("my_project") + ` with output in ` + color.BlueString("PNG") + ` format.
    tgraph -n ` + color.CyanString("my_project") + ` -f ` + color.BlueString("output.png") + `

  Exclude ` + color.RedString("test") + ` and ` + color.RedString("integration") + ` folders while generating the graph.
    tgraph -n ` + color.CyanString("my_project") + ` -i ` + color.RedString("test") + " -i " + color.RedString("integration") + ` -f ` + color.BlueString("output.png") + `

  General use case:
    tgraph ([-n project]) ([-i ignore]) ([-f output | -o open])`
