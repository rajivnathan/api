package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// TotalStatusEmailHashLabelKey is used for the totalstatus email hash label key
	TotalStatusEmailHashLabelKey = LabelKeyPrefix + "email-hash"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TotalStatusSpec defines the desired state of TotalStatus
// +k8s:openapi-gen=true
type TotalStatusSpec struct {
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

	Status string `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TotalStatus is used to track overall toolchain status
// +k8s:openapi-gen=true
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:printcolumn:name="Status",type="string",JSONPath=`.spec.status`
// +kubebuilder:validation:XPreserveUnknownFields
// +operator-sdk:gen-csv:customresourcedefinitions.displayName="Total Status"
type TotalStatus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec TotalStatusSpec `json:"spec,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TotalStatusList contains a list of TotalStatus
type TotalStatusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TotalStatus `json:"items"`
}
