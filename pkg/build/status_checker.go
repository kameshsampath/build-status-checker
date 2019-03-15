package build

import (
	"fmt"

	"github.com/kameshsampath/build-status-checker/pkg/types"
	"github.com/knative/build/pkg/apis/build/v1alpha1"
	knativebuild "github.com/knative/build/pkg/client/clientset/versioned"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"

	duckv1alpha1 "github.com/knative/pkg/apis/duck/v1alpha1"
	corev1 "k8s.io/api/core/v1"

	log "github.com/sirupsen/logrus"
)

//PollAndWait - will check the status of the knative build `buildName` in namespace `namespace`
func PollAndWait(config *rest.Config, buildName string, namespace string, gopts *types.KbscOptions) error {
	logL, err := log.ParseLevel(gopts.LogLevel)
	if err == nil {
		log.SetLevel(logL)
	}

	nameSelector := fmt.Sprintf("metadata.name=%s", buildName)
	log.Debugf("Applying field selector %s", nameSelector)

	//create client set
	clientset, err := knativebuild.NewForConfig(config)

	if err != nil {
		return err
	}

	for {
		w, err := clientset.BuildV1alpha1().Builds(namespace).Watch(v1.ListOptions{FieldSelector: nameSelector})
		if err != nil {
			panic(err)
		}
		//build, err := clientset.BuildV1alpha1().Builds(namespace).Get(buildName, v1.GetOptions{})
		for e := range w.ResultChan() {
			// convert the object to v1alpha1.Build
			b := e.Object.(*v1alpha1.Build)
			log.Debugf("Current Status %v", b.Status)
			var bc = b.Status.GetCondition(duckv1alpha1.ConditionSucceeded)
			if bc != nil {
				if bc.Status == corev1.ConditionTrue {
					log.Infof("Build %s in namespace %s completed \n", buildName, namespace)
					return nil
				} else if bc.Status == corev1.ConditionFalse {
					log.Errorf("Build %s in namespace %s has failed \n  %s \n", buildName, namespace, bc.Message)
					return err
				}
			}
		}
	}
}
