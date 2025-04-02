package k8s

import (
	"context"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"time"
)

type KubernetesClient struct {
	clientSet *kubernetes.Clientset
}

func NewKubernetesClient(api, token string) (*KubernetesClient, error) {
	config := &rest.Config{
		Host:        api,
		BearerToken: token,
		TLSClientConfig: rest.TLSClientConfig{
			Insecure: true,
		},
		Timeout: time.Second * 2,
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	cli := &KubernetesClient{
		clientSet: clientSet,
	}
	return cli, err
}

func (s *KubernetesClient) Deployment(ctx context.Context, namespace string) (records *v1.DeploymentList, err error) {
	return s.clientSet.AppsV1().Deployments(namespace).List(ctx, metav1.ListOptions{})
}
