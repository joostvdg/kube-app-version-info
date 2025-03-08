# kube-app-version-info

Small Go application to monitor the version information of applications installed in Kubernetes. Via Helm Charts, ArgoCD Applications and so on. Aiming to compare the currently installed version with the latest version available.

## Set up local development environment

* Kind
* ArgoCD

```bash
kind create cluster --name kube-app-version-info
```

```bash 
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
```

```bash
kubectl apply -f examples/argocd-application-1.yaml
```

```bash
argocd app create guestbook --repo https://github.com/argoproj/argocd-example-apps.git --path guestbook --dest-server https://kubernetes.default.svc --dest-namespace default
```


## References

* https://blog.logrocket.com/handling-go-configuration-viper/
* https://github.com/golang-standards/project-layout
* https://argo-cd.readthedocs.io/en/stable/cli_installation/
* https://kind.sigs.k8s.io/

== Fucked Up Kubernetes Client Dependecies

```gomod
	k8s.io/cloud-provider v0.31.2
    k8s.io/cluster-bootstrap v0.31.2
    k8s.io/controller-manager v0.31.2
    k8s.io/cri-api v0.31.2
    k8s.io/cri-client v0.31.2
    k8s.io/csi-translation-lib v0.31.2
    k8s.io/dynamic-resource-allocation v0.31.2
    k8s.io/endpointslice v0.31.2
    k8s.io/kube-controller-manager v0.31.2
    k8s.io/kube-proxy v0.31.2
    k8s.io/kube-scheduler v0.31.2
    k8s.io/kubelet v0.31.2
    k8s.io/mount-utils v0.31.2
    k8s.io/pod-security-admission v0.31.2
    k8s.io/sample-apiserver v0.31.2
```