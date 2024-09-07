package k8s_api

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	k8scli "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
	"sync"
)

// describe in config file
var (
	_kubeConfigPath string
	_imageRegistry  string
)

// global immutable var
var (
	_clientSet  *k8scli.Clientset
	_clientOnce sync.Once
)

func Load() {
	_clientOnce.Do(func() {
		kubeConfigDefaultPath := filepath.Join(homedir.HomeDir(), ".kube", "config")
		if !util.OSFilePathExists(_kubeConfigPath) {
			stdlog.ErrorF("kubeConfigPath: %s does not exists, use kubeConfigDefaultPath: %s", _kubeConfigPath, kubeConfigDefaultPath)
			stdlog.InfoF("use default kubeConfigDefaultPath: %s", kubeConfigDefaultPath)
			if !util.OSFilePathExists(kubeConfigDefaultPath) {
				stdlog.FatalF("kubeConfigDefaultPath: %s does not exists", kubeConfigDefaultPath)
			}
			_kubeConfigPath = kubeConfigDefaultPath
		}
		stdlog.InfoF("init kubernetes client with kubeConfigPath: %s", _kubeConfigPath)

		restConfig, err := clientcmd.BuildConfigFromFlags(
			"", _kubeConfigPath)
		if err != nil {
			stdlog.FatalF("init k8s client error: %s, kube config file path: %s", err.Error(), _kubeConfigPath)
		}

		_clientSet, err = k8scli.NewForConfig(restConfig)
		if err != nil {
			stdlog.FatalF("init k8s client error: %s", err.Error())
		}
	})
}

func GetClientSet() *k8scli.Clientset {
	return _clientSet
}

func GetImageRegistry() string {
	return _imageRegistry
}

func WithKubeConfigPath(kubeConfigPath string) {
	_kubeConfigPath = kubeConfigPath
}

func WithImageRegistry(imageRegistry string) {
	_imageRegistry = util.URLWithoutHTTP(imageRegistry)
}
