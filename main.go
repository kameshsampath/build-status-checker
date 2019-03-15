package main

import (
	"fmt"
	"os"

	"github.com/kameshsampath/build-status-checker/pkg/cmd"
	"github.com/kameshsampath/build-status-checker/pkg/types"
	"github.com/spf13/cobra"
)

func main() {
	var opt types.KbscOptions

	rootCmd := &cobra.Command{
		Use:   "kbsc",
		Short: "Check status of knative build",
	}

	rootCmd.PersistentFlags().StringVarP(&opt.KubeConfig, "kubeconfig", "c", "", "Path to a kubeconfig. Only required if out-of-cluster")
	rootCmd.PersistentFlags().StringVarP(&opt.LogLevel, "loglevel", "l", "", "The logrus log level. Default: Info")
	rootCmd.AddCommand(cmd.PollCommand(&opt))

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
