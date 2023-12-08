package main

import (
	"fmt"
	"os"

	"github.com/theosouchon/tgraph/app"
)

func main() {
	if err := app.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
