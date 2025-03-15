//go:build tools
// +build tools

package tools

import (
	_ "k8s.io/cluster-bootstrap"
	_ "k8s.io/controller-manager"
	_ "k8s.io/cri-api"
	_ "k8s.io/cri-client"
	_ "k8s.io/csi-translation-lib"
	_ "k8s.io/dynamic-resource-allocation"
	_ "k8s.io/endpointslice"
	_ "k8s.io/kube-controller-manager"
	_ "k8s.io/kube-proxy"
	_ "k8s.io/kube-scheduler"
	_ "k8s.io/kubelet"
	_ "k8s.io/mount-utils"
	_ "k8s.io/pod-security-admission"
	_ "k8s.io/sample-apiserver"
)
