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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	networkv1alpha1 "github.com/GoogleCloudPlatform/gke-networking-api/apis/network/v1alpha1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// GKENetworkParamSetLister helps list GKENetworkParamSets.
// All objects returned here must be treated as read-only.
type GKENetworkParamSetLister interface {
	// List lists all GKENetworkParamSets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*networkv1alpha1.GKENetworkParamSet, err error)
	// Get retrieves the GKENetworkParamSet from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*networkv1alpha1.GKENetworkParamSet, error)
	GKENetworkParamSetListerExpansion
}

// gKENetworkParamSetLister implements the GKENetworkParamSetLister interface.
type gKENetworkParamSetLister struct {
	listers.ResourceIndexer[*networkv1alpha1.GKENetworkParamSet]
}

// NewGKENetworkParamSetLister returns a new GKENetworkParamSetLister.
func NewGKENetworkParamSetLister(indexer cache.Indexer) GKENetworkParamSetLister {
	return &gKENetworkParamSetLister{listers.New[*networkv1alpha1.GKENetworkParamSet](indexer, networkv1alpha1.Resource("gkenetworkparamset"))}
}
