// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "deploy-operator/pkg/clientset/typed/myresource/v1alpha1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeMygroupV1alpha1 struct {
	*testing.Fake
}

func (c *FakeMygroupV1alpha1) MyResources(namespace string) v1alpha1.MyResourceInterface {
	return &FakeMyResources{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeMygroupV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
