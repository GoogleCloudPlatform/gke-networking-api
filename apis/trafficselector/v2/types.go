/*
Copyright 2026 Google LLC

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

// All fields in this package are required unless Explicitly marked optional
// +kubebuilder:validation:Required
// +k8s:deepcopy-gen=package
package v2

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Constants for TrafficSelector Condition Types
const (
	// ConditionTypeAccepted indicates whether the TrafficSelector resource spec
	// is valid and has been acknowledged by the controller.
	TrafficSelectorConditionTypeAccepted string = "Accepted"

	// TrafficSelectorConditionTypeProgrammed indicates whether the desired state
	// specified in the TrafficSelector has been successfully programmed into the
	// dataplane.
	TrafficSelectorConditionTypeProgrammed string = "Programmed"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster,shortName=ts
// +kubebuilder:printcolumn:name="Priority",type=integer,JSONPath=`.spec.priority`,description="Priority of the TrafficSelector (lower values have higher precedence)"
// +kubebuilder:printcolumn:name="Action Kind",type=string,JSONPath=`.spec.action.kind`,description="Kind of the target action resource"
// +kubebuilder:printcolumn:name="Action Name",type=string,JSONPath=`.spec.action.name`,description="Name of the target action resource"
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
// +kubebuilder:printcolumn:name="Egress CIDRs",type=string,priority=1,JSONPath=`.spec.egress.to.ipBlock.cidrs`,description="Destination CIDR blocks for egress traffic"
// +kubebuilder:printcolumn:name="Ingress CIDRs",type=string,priority=1,JSONPath=`.spec.ingress.from.ipBlock.cidrs`,description="Source CIDR blocks for ingress traffic"

// TrafficSelector is the Schema for the trafficselectors API
// +k8s:openapi-gen=true
type TrafficSelector struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TrafficSelectorSpec   `json:"spec"`
	Status TrafficSelectorStatus `json:"status,omitempty"`
}

// TrafficSelectorSpec defines the desired state of TrafficSelector
// +kubebuilder:validation:XValidation:rule="(has(self.ingress) && !has(self.egress)) || (!has(self.ingress) && has(self.egress))",message="Exactly one of ingress or egress must be specified"
// +k8s:openapi-gen=true
type TrafficSelectorSpec struct {
	// Priority of the TrafficSelector. Lower values have higher precedence.
	// Default value is 1000 if not specified.
	// +optional
	// +kubebuilder:default=1000
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=65535
	Priority *int32 `json:"priority,omitempty"`

	// Subject defines the objects to which the TrafficSelector applies.
	// +kubebuilder:validation:Required
	Subject TrafficSelectorSubject `json:"subject"`

	// Ingress rule to be applied to the selected objects with subject as the frame of reference
	// Exactly one of Ingress or Egress must be specified.
	// +optional
	Ingress *TrafficSelectorIngressRule `json:"ingress,omitempty"`

	// Egress rule to be applied to the selected objects with subject as the frame of reference.
	// Exactly one of Ingress or Egress must be specified.
	// +optional
	Egress *TrafficSelectorEgressRule `json:"egress,omitempty"`

	// Action defines the specific behavior applied to matched traffic.
	// This field defines what action should be taken for traffic selected by either Ingress
	// or Egress rules.
	// +kubebuilder:validation:Required
	Action Action `json:"action"`
}

// Action defines the type and name of the action to be performed on the selected traffic.
// +k8s:openapi-gen=true
type Action struct {
	// This is a reference to another Kubernetes resource
	// (e.g., CloudNATGateway or the ServiceFunctionChain or EgressDSCP) that
	// provides parameters for this action.
	corev1.ObjectReference `json:",inline"`
	// Description is a human-readable description of this action step.
	// +optional
	Description string `json:"description,omitempty"`
}

// TrafficSelectorStatus defines the observed state of TrafficSelector
// +k8s:openapi-gen=true
type TrafficSelectorStatus struct {
	// Conditions describe the current conditions of the TrafficSelector.
	// Implementations should ahere to standard Kubernetes condition conventions.
	//
	// Known condition types are:
	//
	// * "Accepted": This condition indicates whether the TrafficSelector
	//   resource spec is valid and has been acknowledged by the controller.
	//   Reasons for a False status could include syntax errors, invalid
	//   references, or unsupported configurations.
	//
	// * "Programmed": This condition indicates whether the desired state
	//   specified in the TrafficSelector has been successfully programmed into
	//   the underlying dataplane. A True status means the rule and action
	//   is active. A False status indicates that the controller was unable to
	//   apply the configuration.
	//
	// +optional
	// +listType=map
	// +listMapKey=type
	// +kubebuilder:validation:MaxItems=8
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// Exactly one of {`Pods/VMs`, `Services`} pointers must be set.
// +kubebuilder:validation:XValidation:rule="has(self.pods) != has(self.services)",message="Exactly one of Pods or Services must be set"
// +k8s:openapi-gen=true
type TrafficSelectorSubject struct {
	// The TrafficSelector will be applicable to only the interfaces that belong to the specified Network.
	// All service functions in the referenced service chain must belong to the same network as the
	// TrafficSelector to maintain network isolation. The default pod network is chosen if not
	// present.
	// +optional
	Network *string `json:"network,omitempty"`

	// Pods allows the user to select a given set of pod(s) in selected namespace(s).
	// +optional
	Pods *NamespacedPodsSubject `json:"pods,omitempty"`

	// Services allows the user to select a given set of services in selected namespace(s).
	// +optional
	Services *ServiceSelectorSubject `json:"services,omitempty"`
}

// +k8s:openapi-gen=true
type NamespacedPodsSubject struct {
	// NamespaceSelector is used to explicitly select namespaces; if empty, it selects all Namespaces.
	// +optional
	NamespaceSelector *metav1.LabelSelector `json:"namespaceSelector,omitempty"`

	// PodSelector is used to explicitly select pods; if empty, it selects all Pods.
	// +optional
	PodSelector *metav1.LabelSelector `json:"podSelector,omitempty"`

	// NodeSelector restricts the selection to Pods running on Nodes whose labels
	// match this selector. If empty, Pods on all Nodes are eligible for selection.
	// +optional
	NodeSelector *metav1.LabelSelector `json:"nodeSelector,omitempty"`
}

// ServiceSelectorSubject allows selecting services based on namespace and labels.
// +k8s:openapi-gen=true
type ServiceSelectorSubject struct {
	// NamespaceSelector is used to explicitly select namespaces; if empty, it selects all Namespaces.
	NamespaceSelector *metav1.LabelSelector `json:"namespaceSelector,omitempty"`

	// ServiceSelector is used to explicitly select services; if empty, it selects all Services.
	ServiceSelector *metav1.LabelSelector `json:"serviceSelector,omitempty"`
}

// +kubebuilder:validation:MinProperties=1
// +k8s:openapi-gen=true
type TrafficSelectorIngressRule struct {
	// From defines the source peer from which traffic will be steered by this TrafficSelector.
	// If this field is not present or empty, this rule matches all source IPs.
	// +optional
	From *TrafficSelectorPeer `json:"from,omitempty"`
	// Ports allows for matching traffic based on port and protocols.
	// If Ports is not set then the rule does not filter traffic via port.
	// However, this is not currently supported and Ports (with at least 1 item) is required.
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:MaxItems=10
	Ports []TrafficSelectorPort `json:"ports,omitempty"`
}

// +kubebuilder:validation:MinProperties=1
// +k8s:openapi-gen=true
type TrafficSelectorEgressRule struct {
	// To defines the destination peer from which traffic will be steered by this TrafficSelector.
	// If this field is not present or empty, this rule matches all destination IPs.
	// +optional
	To *TrafficSelectorPeer `json:"to,omitempty"`

	// Ports allow for matching traffic based on port and protocols.
	// If Ports are not set then the rule does not filter traffic via port.
	// Support for port matching depends on the configured action; some actions require at least 1 port while others do not support port filtering.
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:MaxItems=10
	Ports []TrafficSelectorPort `json:"ports,omitempty"`
}

// +kubebuilder:validation:MaxProperties=1
// +k8s:openapi-gen=true
type TrafficSelectorPeer struct {
	// IPBlock selects IP CIDR ranges.
	IPBlock *TrafficSelectorIPBlock `json:"ipBlock,omitempty"`
}

// CIDRString represents a string that should contain an IP address/CIDR blocks.
// We enforce a generous MaxLength to bound the CEL estimator cost during CRD validation.
// +kubebuilder:validation:MaxLength=48
// +k8s:openapi-gen=true
type CIDRString string

// +k8s:openapi-gen=true
type TrafficSelectorIPBlock struct {
	// CIDRs is a list of strings representing multiple IP Blocks.
	// +listType=atomic
	// +kubebuilder:validation:MaxItems=100
	// +kubebuilder:validation:XValidation:rule="self.all(x, isCIDR(x))",message="CIDRs elements must be valid CIDRs (e.g. 192.168.1.0/24)"
	CIDRs []CIDRString `json:"cidrs"`
}

// TrafficSelectorPort describes how to select network ports on pod(s).
// +kubebuilder:validation:MinProperties=1
// +kubebuilder:validation:MaxProperties=1
// +k8s:openapi-gen=true
type TrafficSelectorPort struct {
	// Port selects a port on a pod(s).
	// +optional
	PortNumber *Port `json:"portNumber,omitempty"`

	// AllPorts selects all ports for the specified protocol.
	// +optional
	AllPorts *AllPorts `json:"allPorts,omitempty"`
}

// +k8s:openapi-gen=true
type Port struct {
	// Protocol is the network protocol (TCP, UDP) which traffic must match.
	Protocol TrafficSelectorProtocol `json:"protocol"`

	// Port defines a network port value.
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	Port int32 `json:"port"`
}

// +k8s:openapi-gen=true
type AllPorts struct {
	// Protocol is the network protocol (TCP, UDP) which traffic must match.
	Protocol TrafficSelectorProtocol `json:"protocol"`
}

// +kubebuilder:validation:Enum=UDP;TCP
// +k8s:openapi-gen=true
type TrafficSelectorProtocol string

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// TrafficSelectorList contains a list of TrafficSelector
type TrafficSelectorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TrafficSelector `json:"items"`
}
