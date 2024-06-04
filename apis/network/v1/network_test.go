/*
Copyright 2024 Google LLC
Copyright The Kubernetes Authors.

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
package v1

import (
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestIsInUse(t *testing.T) {
	makeNetwork := func(annotations map[string]string) *Network {
		return &Network{
			ObjectMeta: metav1.ObjectMeta{
				Name:        "test",
				Annotations: annotations,
			},
		}
	}

	tests := []struct {
		name  string
		input *Network
		want  bool
	}{
		{
			name:  "nil annotation",
			input: makeNetwork(nil),
			want:  false,
		},
		{
			name:  "missing mn annotation",
			input: makeNetwork(map[string]string{"aaa": "bbb"}),
			want:  false,
		},
		{
			name:  "mn annotation with non true value",
			input: makeNetwork(map[string]string{NetworkInUseAnnotationKey: "false"}),
			want:  false,
		},
		{
			name:  "mn annotation with correct value",
			input: makeNetwork(map[string]string{NetworkInUseAnnotationKey: NetworkInUseAnnotationValTrue}),
			want:  true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.input.InUse()
			if got != tc.want {
				t.Fatalf("IsInUse(%+v) returns %v but want %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestIsDefaultNetwork(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "empty",
			input: "",
			want:  false,
		},
		{
			name:  "old name",
			input: DefaultNetworkName,
			want:  true,
		},
		{
			name:  "new name",
			input: DefaultPodNetworkName,
			want:  true,
		},
		{
			name:  "non-default name",
			input: "test",
			want:  false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := IsDefaultNetwork(tc.input)
			if got != tc.want {
				t.Fatalf("IsDefaultNetwork(%+v) returns %v but want %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestDefaultNetworkIfEmpty(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "empty",
			input: "",
			want:  DefaultPodNetworkName,
		},
		{
			name:  "non-default name",
			input: "test",
			want:  "test",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := DefaultNetworkIfEmpty(tc.input)
			if got != tc.want {
				t.Fatalf("IsDefaultNetwork(%+v) returns %v but want %v", tc.input, got, tc.want)
			}
		})
	}
}
