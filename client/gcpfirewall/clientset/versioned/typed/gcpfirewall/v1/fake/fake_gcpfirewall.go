/*
Copyright 2024 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "github.com/GoogleCloudPlatform/gke-networking-api/apis/gcpfirewall/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeGCPFirewalls implements GCPFirewallInterface
type FakeGCPFirewalls struct {
	Fake *FakeNetworkingV1
}

var gcpfirewallsResource = v1.SchemeGroupVersion.WithResource("gcpfirewalls")

var gcpfirewallsKind = v1.SchemeGroupVersion.WithKind("GCPFirewall")

// Get takes name of the gCPFirewall, and returns the corresponding gCPFirewall object, and an error if there is any.
func (c *FakeGCPFirewalls) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.GCPFirewall, err error) {
	emptyResult := &v1.GCPFirewall{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(gcpfirewallsResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.GCPFirewall), err
}

// List takes label and field selectors, and returns the list of GCPFirewalls that match those selectors.
func (c *FakeGCPFirewalls) List(ctx context.Context, opts metav1.ListOptions) (result *v1.GCPFirewallList, err error) {
	emptyResult := &v1.GCPFirewallList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(gcpfirewallsResource, gcpfirewallsKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.GCPFirewallList{ListMeta: obj.(*v1.GCPFirewallList).ListMeta}
	for _, item := range obj.(*v1.GCPFirewallList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested gCPFirewalls.
func (c *FakeGCPFirewalls) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(gcpfirewallsResource, opts))
}

// Create takes the representation of a gCPFirewall and creates it.  Returns the server's representation of the gCPFirewall, and an error, if there is any.
func (c *FakeGCPFirewalls) Create(ctx context.Context, gCPFirewall *v1.GCPFirewall, opts metav1.CreateOptions) (result *v1.GCPFirewall, err error) {
	emptyResult := &v1.GCPFirewall{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(gcpfirewallsResource, gCPFirewall, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.GCPFirewall), err
}

// Update takes the representation of a gCPFirewall and updates it. Returns the server's representation of the gCPFirewall, and an error, if there is any.
func (c *FakeGCPFirewalls) Update(ctx context.Context, gCPFirewall *v1.GCPFirewall, opts metav1.UpdateOptions) (result *v1.GCPFirewall, err error) {
	emptyResult := &v1.GCPFirewall{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(gcpfirewallsResource, gCPFirewall, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.GCPFirewall), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGCPFirewalls) UpdateStatus(ctx context.Context, gCPFirewall *v1.GCPFirewall, opts metav1.UpdateOptions) (result *v1.GCPFirewall, err error) {
	emptyResult := &v1.GCPFirewall{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceActionWithOptions(gcpfirewallsResource, "status", gCPFirewall, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.GCPFirewall), err
}

// Delete takes name of the gCPFirewall and deletes it. Returns an error if one occurs.
func (c *FakeGCPFirewalls) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(gcpfirewallsResource, name, opts), &v1.GCPFirewall{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGCPFirewalls) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(gcpfirewallsResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1.GCPFirewallList{})
	return err
}

// Patch applies the patch and returns the patched gCPFirewall.
func (c *FakeGCPFirewalls) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.GCPFirewall, err error) {
	emptyResult := &v1.GCPFirewall{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(gcpfirewallsResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.GCPFirewall), err
}
