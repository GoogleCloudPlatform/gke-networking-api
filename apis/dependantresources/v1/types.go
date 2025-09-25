/*
Copyright 2025 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced,shortName=slbs
// +kubebuilder:storageversion
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DependantResources holds the mapping between a GKE Service and the GCE resources it configures.
// +k8s:openapi-gen=true
type DependantResources struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the desired state of DependantResources.
	// This field is currently unused.
	// +optional
	Spec DependantResourcesSpec `json:"spec,omitempty"`

	// Status defines the observed state of DependantResources.
	// This field is populated by the controller with the list of GCE resources.
	// +optional
	Status DependantResourcesStatus `json:"status,omitempty"`
}

// DependantResourcesSpec is the spec for a DependantResources resource
// +k8s:openapi-gen=true
type DependantResourcesSpec struct{}

// DependantResourcesStatus defines the observed state of DependantResources.
// +k8s:openapi-gen=true
type DependantResourcesStatus struct {
	// Resources is a list of URLs for resources provisioned
	// by the controller for the referenced object. This field is managed
	// exclusively by the controller.
	// +optional
	// +listType=set
	Resources []string `json:"resources,omitempty"`
}

// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DependantResourcesList contains a list of DependantResources.
type DependantResourcesList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	// Items is the list of DependantResources.
	Items []DependantResources `json:"items"`
}
