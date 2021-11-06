package main

import (
	"context"
	"fmt"
	"os"

	"github.com/xenolog/negra/pkg/cli"
	"k8s.io/klog"
)

func main() {
	var exitCode int
	defer func() { os.Exit(exitCode) }() // should be first and in this form
	defer klog.Flush()

	ctx := context.TODO()
	if err := cli.RootCmd.ExecuteContext(ctx); err != nil {
		exitCode = 1
	}
	fmt.Println()
}
