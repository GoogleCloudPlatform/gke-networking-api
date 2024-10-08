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

	v1 "github.com/GoogleCloudPlatform/gke-networking-api/apis/nodetopology/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNodeTopologies implements NodeTopologyInterface
type FakeNodeTopologies struct {
	Fake *FakeNetworkingV1
}

var nodetopologiesResource = v1.SchemeGroupVersion.WithResource("nodetopologies")

var nodetopologiesKind = v1.SchemeGroupVersion.WithKind("NodeTopology")

// Get takes name of the nodeTopology, and returns the corresponding nodeTopology object, and an error if there is any.
func (c *FakeNodeTopologies) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.NodeTopology, err error) {
	emptyResult := &v1.NodeTopology{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(nodetopologiesResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.NodeTopology), err
}

// List takes label and field selectors, and returns the list of NodeTopologies that match those selectors.
func (c *FakeNodeTopologies) List(ctx context.Context, opts metav1.ListOptions) (result *v1.NodeTopologyList, err error) {
	emptyResult := &v1.NodeTopologyList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(nodetopologiesResource, nodetopologiesKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.NodeTopologyList{ListMeta: obj.(*v1.NodeTopologyList).ListMeta}
	for _, item := range obj.(*v1.NodeTopologyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested nodeTopologies.
func (c *FakeNodeTopologies) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(nodetopologiesResource, opts))
}

// Create takes the representation of a nodeTopology and creates it.  Returns the server's representation of the nodeTopology, and an error, if there is any.
func (c *FakeNodeTopologies) Create(ctx context.Context, nodeTopology *v1.NodeTopology, opts metav1.CreateOptions) (result *v1.NodeTopology, err error) {
	emptyResult := &v1.NodeTopology{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(nodetopologiesResource, nodeTopology, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.NodeTopology), err
}

// Update takes the representation of a nodeTopology and updates it. Returns the server's representation of the nodeTopology, and an error, if there is any.
func (c *FakeNodeTopologies) Update(ctx context.Context, nodeTopology *v1.NodeTopology, opts metav1.UpdateOptions) (result *v1.NodeTopology, err error) {
	emptyResult := &v1.NodeTopology{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(nodetopologiesResource, nodeTopology, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.NodeTopology), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNodeTopologies) UpdateStatus(ctx context.Context, nodeTopology *v1.NodeTopology, opts metav1.UpdateOptions) (result *v1.NodeTopology, err error) {
	emptyResult := &v1.NodeTopology{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceActionWithOptions(nodetopologiesResource, "status", nodeTopology, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.NodeTopology), err
}

// Delete takes name of the nodeTopology and deletes it. Returns an error if one occurs.
func (c *FakeNodeTopologies) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(nodetopologiesResource, name, opts), &v1.NodeTopology{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNodeTopologies) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(nodetopologiesResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1.NodeTopologyList{})
	return err
}

// Patch applies the patch and returns the patched nodeTopology.
func (c *FakeNodeTopologies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.NodeTopology, err error) {
	emptyResult := &v1.NodeTopology{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(nodetopologiesResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.NodeTopology), err
}
