package kubernetes_watchers

import (
	"context"
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"

	argo "github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

func ListApps() {
	// This function will list all the apps in the Cluster
	// Configure the Kubernetes client
	config, err := clientcmd.BuildConfigFromFlags("", os.Getenv("KUBECONFIG"))
	if err != nil {
		panic(err.Error())
	}

	// Create a dynamic client
	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// Specify the GVR: GroupVersionResource
	gvr := schema.GroupVersionResource{Group: "argoproj.io", Version: "v1alpha1", Resource: "applications"}

	// Retrieve all Appliciations in the cluster
	apps, err := dynamicClient.Resource(gvr).List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	for _, app := range apps.Items {
		fmt.Printf("Application: %s\n", app.GetName())
	}

}

// KubernetesWatcher watches for changes in ArgoCD/Helm CRs
type KubernetesWatcher struct {
	Clientset     *kubernetes.Clientset
	DynamicClient dynamic.Interface
}

func (w *KubernetesWatcher) WatchArgoCDApplications() {
	gvr := schema.GroupVersionResource{Group: "argoproj.io", Version: "v1alpha1", Resource: "applications"}
	resourceInterface := w.DynamicClient.Resource(gvr).Namespace(metav1.NamespaceAll)

	dynamicInformer := cache.NewSharedInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				return resourceInterface.List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				return resourceInterface.Watch(context.TODO(), options)
			},
		},
		&unstructured.Unstructured{},
		0, // resync period, set to 0 to disable
		cache.Indexers{},
	)

	dynamicInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			app, ok := obj.(*argo.Application)
			if ok {
				log.Info().Msgf("New ArgoCD Application detected: %s", app.Name)
				// Process and store application version
			}
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			app, ok := newObj.(*argo.Application)
			if ok {
				log.Info().Msgf("Updated ArgoCD Application detected: %s", app.Name)
				// Update version tracking
			}
		},
		DeleteFunc: func(obj interface{}) {
			app, ok := obj.(*argo.Application)
			if ok {
				log.Info().Msgf("Deleted ArgoCD Application: %s", app.Name)
			}
		},
	})

	stopCh := make(chan struct{})
	defer close(stopCh)
	dynamicInformer.Run(stopCh)
}
