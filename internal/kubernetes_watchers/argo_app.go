package kubernetes_watchers

import (
	"context"
	"fmt"
	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	"github.com/joostvdg/kube-app-version-info/internal/applications"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/json"
	"k8s.io/client-go/dynamic/dynamicinformer"
	"os"

	"github.com/rs/zerolog/log"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"

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

func (w *KubernetesWatcher) WatchArgoCDApplications(storage applications.Store) {
	gvr := schema.GroupVersionResource{Group: "argoproj.io", Version: "v1alpha1", Resource: "applications"}

	log.Info().Msg("Watching ArgoCD Applications")
	factory := dynamicinformer.NewFilteredDynamicSharedInformerFactory(w.DynamicClient, 0, metav1.NamespaceAll, nil)

	log.Info().Msg("Creating informer")
	informer := factory.ForResource(gvr).Informer()
	log.Info().Msg("Adding event handler")
	_, err := informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			argoApp, err := transformToArgoApp(obj)
			if err != nil {
				log.Error().Err(err).Msg("Failed to transform to ArgoCD Application")
				return
			}
			applications.ProcessArgoCDApplication(storage, argoApp)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			argoApp, err := transformToArgoApp(newObj)
			if err != nil {
				log.Error().Err(err).Msg("Failed to transform to ArgoCD Application")
				return
			}
			applications.ProcessArgoCDApplication(storage, argoApp)
		},
		DeleteFunc: func(obj interface{}) {
			argoApp, err := transformToArgoApp(obj)
			if err != nil {
				log.Error().Err(err).Msg("Failed to transform to ArgoCD Application")
				return
			}
			applications.ProcessArgoCDApplication(storage, argoApp)
		},
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to add event handler")
		return
	}

	stopCh := make(chan struct{})
	defer close(stopCh)
	informer.Run(stopCh)
}

func transformToArgoApp(obj interface{}) (*v1alpha1.Application, error) {
	unstructuredObj := obj.(*unstructured.Unstructured)
	jsonData, err := unstructuredObj.MarshalJSON()
	if err != nil {
		log.Error().Err(err).Msg("Failed to marshal unstructured object to JSON")
		return nil, err
	}
	app := &v1alpha1.Application{}
	err = json.Unmarshal(jsonData, app)
	if err != nil {
		log.Error().Err(err).Msg("Failed to unmarshal JSON to ArgoCD Application")
		return nil, err
	}
	return app, nil
}
