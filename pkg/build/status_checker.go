package build

import (
	"fmt"
	"os"
	"time"

	knativebuild "github.com/knative/build/pkg/client/clientset/versioned"
	duckv1alpha1 "github.com/knative/pkg/apis/duck/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

//PollAndWait - will check the status of the knative build `buildName` in namespace `namespace`
func PollAndWait(config *rest.Config, buildName string, namespace string) error {

	//create client set
	clientset, err := knativebuild.NewForConfig(config)

	if err != nil {
		return err
	}

	for {
		time.Sleep(10 * time.Second)
		//TODO pull events than pod
		build, err := clientset.BuildV1alpha1().Builds(namespace).Get(buildName, v1.GetOptions{})
		if err != nil {
			return err
		}
		// there cant be more then one build pod with same name
		if build != nil {
			var bc = build.Status.GetCondition(duckv1alpha1.ConditionSucceeded)
			if bc != nil {
				if bc.Status == corev1.ConditionTrue {
					fmt.Printf("Build %s in namespace %s completed \n", buildName, namespace)
					return err
				} else if bc.Status == corev1.ConditionFalse {
					fmt.Fprintf(os.Stderr, "Build %s in namespace %s failed \n  %s \n", buildName, namespace, bc.Message)
					return err
				}
			}
		} else {
			fmt.Fprintf(os.Stderr, "No build pod(s) with name '%s' running in namespace '%s' \n ", buildName, namespace)
			return nil
		}
	}

}
