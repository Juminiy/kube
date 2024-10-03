package service_api

import (
	"context"
	"github.com/Juminiy/kube/pkg/k8s_api"
	kubek8sinternal "github.com/Juminiy/kube/pkg/k8s_api/internal_api"
	"github.com/Juminiy/kube/pkg/util"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

type ServiceConfig struct {
	Namespace string
	MetaName  string

	SpecSelector map[string]string
	SpecPorts    []corev1.ServicePort

	cli     typedcorev1.ServiceInterface
	service *corev1.Service
	ctx     context.Context
	cbk     *kubek8sinternal.CallBack
}

func NewService(serviceConfig *ServiceConfig) error {
	serviceConfig.cli = k8s_api.GetClientSet().CoreV1().Services(serviceConfig.Namespace)

	serviceConfig.service = &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: serviceConfig.MetaName,
		},
		Spec: corev1.ServiceSpec{
			Selector: serviceConfig.SpecSelector,
			Ports:    serviceConfig.SpecPorts,
			Type:     corev1.ServiceTypeNodePort,
		},
	}

	serviceConfig.ctx = util.TODOContext()

	serviceConfig.cbk = kubek8sinternal.LogLatestCallBack()

	return nil
}

func (c *ServiceConfig) Create() error {
	c.cbk.Latest, c.cbk.LatestErr = c.cli.Create(c.ctx, c.service, metav1.CreateOptions{})
	if c.cbk.LatestErr != nil {
		return c.cbk.LatestErr
	}

	return c.cbk.Create()
}

func (c *ServiceConfig) Update() error {
	c.cbk.Latest, c.cbk.LatestErr = c.cli.Update(c.ctx, c.service, metav1.UpdateOptions{})
	if c.cbk.LatestErr != nil {
		return c.cbk.LatestErr
	}

	return c.cbk.Update()
}

func (c *ServiceConfig) Delete() error {
	deleteErr := c.cli.Delete(c.ctx, c.MetaName, metav1.DeleteOptions{})
	if deleteErr != nil {
		return deleteErr
	}

	return c.cbk.Delete()
}

func (c *ServiceConfig) List() error {
	c.cbk.Latest, c.cbk.LatestErr = c.cli.List(c.ctx, metav1.ListOptions{})
	if c.cbk.LatestErr != nil {
		return c.cbk.LatestErr
	}

	return c.cbk.List()
}
