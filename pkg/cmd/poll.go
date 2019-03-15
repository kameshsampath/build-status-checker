package cmd

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/kameshsampath/build-status-checker/pkg/build"
	"github.com/kameshsampath/build-status-checker/pkg/types"
)

//PollCommand - polls for a Knative build status
func PollCommand(opts *types.KbscOptions) *cobra.Command {
	var opt types.PollOptions
	cmd := &cobra.Command{
		Use:   "poll",
		Short: "polls a knative build by name and waits for its status",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Polling build:", opt.BuildName)
			config, err := clientcmd.BuildConfigFromFlags("", opts.KubeConfig)
			if err != nil {
				return errors.Wrap(err, "could not create kubernetes client config")
			}
			if err := build.PollAndWait(config, opt.BuildName, opt.Namespace); err != nil {
				return err
			}
			return nil
		},
	}

	cmd.Flags().StringVarP(&opt.BuildName, "buildname", "b", "", "the knative build name")
	cmd.Flags().StringVarP(&opt.Namespace, "namespace", "n", "default", "the kubernetes namespace of the build")
	cmd.MarkFlagRequired("buildname")

	return cmd
}
