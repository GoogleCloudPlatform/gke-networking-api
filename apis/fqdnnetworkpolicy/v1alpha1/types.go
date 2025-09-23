/*
Copyright 2022 Google LLC

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

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:shortName=fqdnnp

// FQDNNetworkPolicy specifies which Fully-Qualified Domain Names a
// workload is allowed to egress to. FQDNNetworkPolicy has no effect
// on the allowed ingress traffic.
//
// FQDNNetworkPolicy works additively with Kubernetes Network Policies.
// Traffic matching either a Kubernetes NetworkPolicy or a FQDNNetworkPolicy
// will be allowed to egress the workload.
type FQDNNetworkPolicy struct {
	v1.TypeMeta   `json:",inline"`
	v1.ObjectMeta `json:"metadata,omitempty"`

	// Spec is the desired configuration for FQDN Network Policy.
	Spec FQDNNetworkPolicySpec `json:"spec"`
}

type FQDNNetworkPolicySpec struct {
	// PodSelector defines which workloads are selected by the Policy.
	// An empty PodSelector selects all pods in the namespace.
	PodSelector v1.LabelSelector `json:"podSelector"`

	// Egress specifies a list of rules applied to the selected pods.
	// All egress rules are enforced using an implicit deny model
	// (like Kubernetes Network Policies) -- only packets destined for
	// IPs resolving the matching FQDNs and ports are allowed.
	//
	// Egress may not be missing or empty - it must contain at least 1
	// entry.
	//
	// +kubebuilder:validation:MinItems=1
	Egress []FQDNNetworkPolicyEgressRule `json:"egress"`
}

// FQDNNetworkPolicyEgressRule identifies a set of endpoints traffic is allowed
// to. The exact L4 endpoints are constructed as a cross-product between the
// matches and ports lists.
type FQDNNetworkPolicyEgressRule struct {
	// Matches specifies the FQDN peers to which egress traffic is allowed.
	// Matches may not be missing or empty - it must contain at least 1
	// entry.
	//
	// +kubebuilder:validation:MinItems=1
	Matches []FQDNNetworkPolicyMatch `json:"matches"`

	// Ports specifies the destination L4 port and protocol allowed to
	// egress the pod.
	// If ports is missing or empty, all ports and protocols are allowed.
	//
	// +optional
	Ports []FQDNNetworkPolicyPort `json:"ports,omitempty"`
}

// FQDNNetworkPolicyMatch specifies which FQDNs are allowed as peers of the
// selected pods.Exactly one of the sub-fields of FQDNNetworkPolicyMatch must be
// set. An empty struct is not allowed. A struct setting both Name and Pattern
// is not allowed.
//
// +kubebuilder:validation:MinProperties=1
// +kubebuilder:validation:MaxProperties=1
type FQDNNetworkPolicyMatch struct {
	// Name specifies the literal FQDN to match. If this is specified,
	// no other match types may be specified in the same struct.
	//
	// +optional
	// +kubebuilder:validation:Pattern=`^([a-zA-Z0-9]([-a-zA-Z0-9_]*[a-zA-Z0-9])*\.?)*$`
	Name string `json:"name,omitempty"`

	// Pattern allows matching FQDNs using wildcard specifiers. If this
	// is specified, no other match types may be specified in the same
	// struct.
	// "*" matches 0 or more DNS valid characters (except for "."), and may occur
	// anywhere in the pattern.
	// As a special case, the "*" character by itself acts as a wildcard,
	// matching all domain names.
	//
	// Examples:
	//
	//   - `*.google.com` matches subdomains of google.com at that level
	//     - "www.google.com" and "mail.google.com" match, however "google.com",
	//       "sub.subdomain.google.com", and "kubernetes.io" do not.
	//   - `sub*.google.com` matches subdomains of google.com where the subdomain
	//     component begins with "sub"
	//     - "sub.google.com" and "subdomain.google.com" match, however
	//       "www.google.com", "mail.google.com", and "google.com do not.
	//   - `*.*.google.com` matches subdomains of google.com at that level
	//     - "sub.subdomain.google.com" matches, however "google.com" and
	//       "mail.google.com" do not.
	//   - `*a*.google.com` matches subdomains of google.com that contain an "a" at
	//     that level
	//     - "mail.google.com", "maps.google.com", and
	//       "subdomain.maps.google.com" match, however "google.com" and
	//       "cloud.google.com" do not.
	//   - `*` by itself is a wild-card that matches all domains
	//     - "www.google.com", "kubernetes.io", and "sub.subdomain.google.com"
	//       all match
	//
	// +optional
	// +kubebuilder:validation:Pattern=`^([a-zA-Z0-9*]([-a-zA-Z0-9_*]*[a-zA-Z0-9*])*\.?)*$`
	Pattern string `json:"pattern,omitempty"`
}

// FQDNNetworkPolicyPort specifies which remote port and protocol is a valid
// peer of the selected pod.
type FQDNNetworkPolicyPort struct {
	// Protocol is the L4 protocol. Valid options are "TCP", "UDP",
	// or "ALL".
	// If Protocol is missing or empty, it defaults to allowing all
	// protocols.
	//
	// +optional
	// +kubebuilder:validation:Enum=TCP;UDP;ALL
	// +kubebuilder:validation:default:=ALL
	Protocol string `json:"protocol,omitempty"`

	// Port is L4 Port.
	// If Port is missing or empty, it matches all ports.
	//
	// +optional
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	Port *int32 `json:"port,omitempty"`
}

// +genclient
// +genclient:onlyVerbs=get
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FQDNNetworkPolicyList contains a list of FQDNNetworkPolicy resources.
type FQDNNetworkPolicyList struct {
	v1.TypeMeta `json:",inline"`
	v1.ListMeta `json:"metadata,omitempty"`

	// Items is a slice of FQDNNetworkPolicy resources.
	Items []FQDNNetworkPolicy `json:"items"`
}
