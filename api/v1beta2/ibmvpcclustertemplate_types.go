/*
Copyright The Kubernetes Authors.

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

package v1beta2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// IBMVPCClusterTemplateSpec defines the desired state of IBMVPCClusterTemplate
type IBMVPCClusterTemplateSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of IBMVPCClusterTemplate. Edit ibmvpcclustertemplate_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// IBMVPCClusterTemplateStatus defines the observed state of IBMVPCClusterTemplate
type IBMVPCClusterTemplateStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// IBMVPCClusterTemplate is the Schema for the ibmvpcclustertemplates API
type IBMVPCClusterTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IBMVPCClusterTemplateSpec   `json:"spec,omitempty"`
	Status IBMVPCClusterTemplateStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// IBMVPCClusterTemplateList contains a list of IBMVPCClusterTemplate
type IBMVPCClusterTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IBMVPCClusterTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&IBMVPCClusterTemplate{}, &IBMVPCClusterTemplateList{})
}
