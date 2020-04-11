package fanout

import (
	"fmt"
	"os"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	coreinformersv1 "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

type podFanout struct {
	client      kubernetes.Interface
	namespace   string
	targetPort  int
	selector    labels.Selector
	podInformer coreinformersv1.PodInformer
}

func (f *podFanout) Receivers() ([]string, error) {
	var ep []string
	podList, err := f.podInformer.Lister().Pods(f.namespace).List(f.selector)
	if err != nil {
		return ep, err
	}
	for _, pod := range podList {
		ep = append(ep, fmt.Sprintf("%s:%d", pod.Status.PodIP, f.targetPort))
	}
	return ep, nil
}

// NewPodFanout creates a new serviceFanout from the given arguments
func NewPodFanout(namespace string, requirements string, targetPort int) (Fanout, error) {
	f := &podFanout{
		namespace:  namespace,
		targetPort: targetPort,
	}
	config, err := rest.InClusterConfig()
	if err != nil {
		kubeconfig := os.Getenv("KUBECONFIG")
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			return f, err
		}
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return f, err
	}
	f.client = client
	labelSelector, err := metav1.ParseToLabelSelector(requirements)
	if err != nil {
		return f, err
	}
	f.selector, err = metav1.LabelSelectorAsSelector(labelSelector)
	if err != nil {
		return f, err
	}
	informer := informers.NewSharedInformerFactoryWithOptions(client, 0, informers.WithNamespace(namespace))
	f.podInformer = informer.Core().V1().Pods()
	f.podInformer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
			},
		},
	)
	fmt.Printf("Starting pod fanout with config %+v\n", f)
	informer.Start(wait.NeverStop)
	err = wait.Poll(time.Second, 30, func() (bool, error) {
		return f.podInformer.Informer().HasSynced(), nil
	})
	return f, nil
}
