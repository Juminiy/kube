package k8s_api

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	k8scli "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
	"sync"
)

// describe in config file
var (
	_kubeConfigPath string
)

// global immutable var
var (
	_clientSet  *k8scli.Clientset
	_clientOnce sync.Once
)

func Init() {
	_clientOnce.Do(func() {
		restConfig, err := clientcmd.BuildConfigFromFlags(
			"",
			filepath.Join(homedir.HomeDir(), ".kube", "config"))
		if err != nil {
			stdlog.InfoF("init k8s client error: %s, kube config file path: %s", err.Error(), _kubeConfigPath)
			return
		}

		_clientSet, err = k8scli.NewForConfig(restConfig)
		stdlog.InfoF("init k8s client error: %s", err.Error())
	})
}

func Get() *k8scli.Clientset {
	return _clientSet
}

func WithKubeConfigPath(kubeConfigPath string) {
	_kubeConfigPath = kubeConfigPath
}
