package cmd

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/kameshsampath/build-status-checker/pkg/build"
	"github.com/kameshsampath/build-status-checker/pkg/helpers"
	"github.com/kameshsampath/build-status-checker/pkg/types"

	log "github.com/sirupsen/logrus"
)

//PollCommand - polls for a Knative build status
func PollCommand(gopts *types.KbscOptions) *cobra.Command {
	var opt types.PollOptions
	cmd := &cobra.Command{
		Use:   "poll",
		Short: "polls a knative build by name and waits for its status",
		RunE: func(cmd *cobra.Command, args []string) error {
			helpers.SetLogLevel(gopts.LogLevel)
			log.Debugf("Using kubeconfig : %s", gopts.KubeConfig)
			log.Infof("Polling build: '%s' in namespace '%s' ", opt.BuildName, opt.Namespace)
			config, err := clientcmd.BuildConfigFromFlags("", gopts.KubeConfig)
			if err != nil {
				return errors.Wrap(err, "could not create kubernetes client config")
			}
			if err := build.PollAndWait(config, opt.BuildName, opt.Namespace, gopts); err != nil {
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
