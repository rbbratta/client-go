// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/openshift/client-go/monitoring/clientset/versioned/typed/monitoring/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeMonitoringV1alpha1 struct {
	*testing.Fake
}

func (c *FakeMonitoringV1alpha1) AlertRelabelConfigs(namespace string) v1alpha1.AlertRelabelConfigInterface {
	return &FakeAlertRelabelConfigs{c, namespace}
}

func (c *FakeMonitoringV1alpha1) AlertingRules(namespace string) v1alpha1.AlertingRuleInterface {
	return &FakeAlertingRules{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeMonitoringV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
