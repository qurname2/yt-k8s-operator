/*
Copyright 2023.

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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// RemoteDataNodesSpec defines the desired state of RemoteDataNodes
type RemoteDataNodesSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	RemoteClusterSpec *corev1.LocalObjectReference `json:"remoteClusterSpec"`
	CommonSpec        `json:",inline"`
	DataNodesSpec     `json:",inline"`
}

type RemoteDataNodeReleaseStatus string

const (
	RemoteDataNodeReleaseStatusPending RemoteDataNodeReleaseStatus = "Pending"
	RemoteDataNodeReleaseStatusRunning RemoteDataNodeReleaseStatus = "Running"
)

// RemoteDataNodesStatus defines the observed state of RemoteDataNodes
type RemoteDataNodesStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	ReleaseStatus RemoteDataNodeReleaseStatus `json:"releaseStatus,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:printcolumn:name="ReleaseStatus",type="string",JSONPath=".status.releaseStatus",description="Release status"
//+kubebuilder:resource:categories=ytsaurus-all;yt-all
//+kubebuilder:subresource:status

// RemoteDataNodes is the Schema for the RemoteDatanodes API
type RemoteDataNodes struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RemoteDataNodesSpec   `json:"spec,omitempty"`
	Status RemoteDataNodesStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// RemoteDataNodesList contains a list of RemoteDataNodes
type RemoteDataNodesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RemoteDataNodes `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RemoteDataNodes{}, &RemoteDataNodesList{})
}
