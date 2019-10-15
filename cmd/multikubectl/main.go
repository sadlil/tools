package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/sadlil/tools/pkg/kubectl/cmd"
	"k8s.io/kubectl/pkg/util/logs"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	command := cmd.NewDefaultMultiKubectlCommand()

	// cliflag.InitFlags()
	logs.InitLogs()
	defer logs.FlushLogs()

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
