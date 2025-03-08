package main

import (
	"github.com/joostvdg/kube-app-version-info/internal/kubernetes_watchers"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"net/http"
	"os"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		return
	}
	port := viper.Get("PORT")

	log.Printf("Call out to the Kubernetes API to list all the applications in the cluster")
	kubernetes_watchers.ListApps()

	log.Printf("Starting webserver  on: %s", port)
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Start the watcher
	log.Printf("Starting the Kubernetes Watcher")
	config, err := clientcmd.BuildConfigFromFlags("", os.Getenv("KUBECONFIG"))
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err2 := kubernetes.NewForConfig(config)
	if err2 != nil {
		panic(err.Error())
	}

	argoAppsWatcher := kubernetes_watchers.KubernetesWatcher{
		Clientset: clientset,
	}
	argoAppsWatcher.WatchArgoCDApplications()

	echoPort := ":" + port.(string)
	e.Logger.Fatal(e.Start(echoPort))
}
