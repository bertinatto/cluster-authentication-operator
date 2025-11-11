package deployment

import (
	"context"
	"fmt"
	"os"

	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"

	"github.com/openshift/cluster-authentication-operator/bindata"
	"github.com/openshift/library-go/pkg/controller/factory"
	"github.com/openshift/library-go/pkg/operator/events"
	"github.com/openshift/library-go/pkg/operator/resource/resourceapply"
	"github.com/openshift/library-go/pkg/operator/resource/resourceread"
	"github.com/openshift/library-go/pkg/operator/v1helpers"
)

type hypershiftOAuthServerController struct {
	kubeClient    kubernetes.Interface
	eventRecorder events.Recorder
}

func NewHyperShiftOAuthServerController(
	kubeClient kubernetes.Interface,
	kubeInformers v1helpers.KubeInformersForNamespaces,
	eventsRecorder events.Recorder,
) factory.Controller {
	targetNS := "openshift-authentication"

	c := &hypershiftOAuthServerController{kubeClient: kubeClient, eventRecorder: eventsRecorder}

	return factory.New().
		WithSync(c.sync).
		WithInformers(
			kubeInformers.InformersFor(targetNS).Apps().V1().Deployments().Informer(),
		).
		ToController("HyperShiftOAuthServerController", eventsRecorder)
}

func (c *hypershiftOAuthServerController) sync(ctx context.Context, syncContext factory.SyncContext) error {
	requiredDeployment := resourceread.ReadDeploymentV1OrDie(bindata.MustAsset("oauth-openshift/hypershift/deployment.yaml"))
	deployment := requiredDeployment.DeepCopy()

	oauthImage := os.Getenv("OPERAND_OAUTH_SERVER_IMAGE")
	if len(oauthImage) == 0 {
		klog.Warningf("OPERAND_OAUTH_SERVER_IMAGE is not set")
	}

	cliImage := os.Getenv("CLI_IMAGE")
	if len(cliImage) == 0 {
		klog.Warningf("CLI_IMAGE is not set")
	}

	for i := range deployment.Spec.Template.Spec.Containers {
		switch deployment.Spec.Template.Spec.Containers[i].Name {
		case "oauth-openshift":
			if len(oauthImage) > 0 {
				deployment.Spec.Template.Spec.Containers[i].Image = oauthImage
			}
		case "audit-logs":
			if len(cliImage) > 0 {
				deployment.Spec.Template.Spec.Containers[i].Image = cliImage
			}
		}
	}

	_, _, err := resourceapply.ApplyDeployment(
		ctx,
		c.kubeClient.AppsV1(),
		c.eventRecorder,
		deployment,
		-1,
	)
	if err != nil {
		return fmt.Errorf("failed to apply HyperShift oauth-server deployment: %w", err)
	}

	klog.V(2).Info("Successfully reconciled HyperShift oauth-server deployment")
	return nil
}
