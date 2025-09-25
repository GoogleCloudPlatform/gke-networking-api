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

// ServiceLoadBalancerStatus holds the mapping between a GKE Service and the GCE resources it configures.
type ServiceLoadBalancerStatus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the desired state of ServiceLoadBalancerStatus.
	// This field is currently unused.
	// +optional
	Spec ServiceLoadBalancerStatusSpec `json:"spec,omitempty"`

	// Status defines the observed state of ServiceLoadBalancerStatus.
	// This field is populated by the controller with the list of GCE resources.
	// +optional
	Status ServiceLoadBalancerStatusStatus `json:"status,omitempty"`
}

// ServiceLoadBalancerStatusSpec is the spec for a ServiceLoadBalancerStatus resource
type ServiceLoadBalancerStatusSpec struct{}

// ServiceLoadBalancerStatusStatus defines the observed state of ServiceLoadBalancerStatus.
type ServiceLoadBalancerStatusStatus struct {
	// GceResources is a list of URLs for GCE resources provisioned and managed
	// by the GKE controller for the referenced Service. This field is managed
	// exclusively by the controller.
	// +optional
	// +listType=set
	GceResources []string `json:"gceResources,omitempty"`
}

// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ServiceLoadBalancerStatusList contains a list of ServiceLoadBalancerStatus.
type ServiceLoadBalancerStatusList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	// Items is the list of ServiceLoadBalancerStatus.
	Items []ServiceLoadBalancerStatus `json:"items"`
}
