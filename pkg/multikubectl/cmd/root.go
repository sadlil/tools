package cmd

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	"k8s.io/klog"
)

func NewDefaultMultiKubectlCommand() *cobra.Command {
	return &cobra.Command{
		Use:                "multikubectl",
		Run:                run,
		TraverseChildren:   false,
		DisableFlagParsing: true,
	}
}

func run(cmd *cobra.Command, args []string) {
	kubeConfig, err := clientcmd.LoadFromFile(os.Getenv("HOME") + "/.kube/config")
	if err != nil {
		klog.Fatal(err)
	}

	for name, cluster := range kubeConfig.Clusters {
		if strings.HasPrefix(name, "bplatform") {
			fmt.Println("============================================================")
			fmt.Printf("Cluster: %s:\n", name)

			var user *clientcmdapi.AuthInfo
			for name, auth := range kubeConfig.AuthInfos {
				if strings.EqualFold(name, "bplatform-default") {
					user = auth
					break
				}
			}

			newCfg := &clientcmdapi.Config{
				Clusters: map[string]*clientcmdapi.Cluster{
					name: cluster,
				},
				AuthInfos: map[string]*clientcmdapi.AuthInfo{
					name: user,
				},
				CurrentContext: name,
				Contexts: map[string]*clientcmdapi.Context{
					name: {
						Cluster:  name,
						AuthInfo: name,
					},
				},
			}

			cfgFile := os.TempDir() + "/.multikube/cluster-" + name
			err = clientcmd.WriteToFile(*newCfg, cfgFile)
			if err != nil {
				klog.Fatal(err)
			}

			_ = os.Setenv("KUBECONFIG", cfgFile)
			cmd := exec.Command("kubectl", args...)
			cmd.Stderr = NewEmptyLineRemoveWriter(os.Stderr)
			cmd.Stdin = os.Stdin
			cmd.Stdout = NewEmptyLineRemoveWriter(os.Stdout)
			_ = cmd.Run()
			_ = os.Unsetenv("KUBECONFIG")
			_ = os.Remove(cfgFile)
		}
	}
}

type writer struct {
	w io.Writer
	counter int
	buf []byte
}

func NewEmptyLineRemoveWriter(w io.Writer) io.Writer {
	return &writer{
		w: w,
	}
}

func (w *writer) Write(b []byte) (int, error) {
	return w.w.Write(b)
}