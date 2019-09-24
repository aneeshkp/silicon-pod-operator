package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//PhaseType ...
type PhaseType string

// Const for phasetype
const (
	CollectdPhaseNone     PhaseType = ""
	CollectdPhaseCreating           = "Creating"
	CollectdPhaseRunning            = "Running"
	CollectdPhaseFailed             = "Failed"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SiliconPodSpec defines the desired state of SiliconPod
// +k8s:openapi-gen=true
type SiliconPodSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Size int32 `json:"size"`
}

// SiliconPodStatus defines the observed state of SiliconPod
// +k8s:openapi-gen=true
type SiliconPodStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Phase     PhaseType `json:"phase,omitempty"`
	RevNumber string    `json:"revNumber,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SiliconPod is the Schema for the siliconpods API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type SiliconPod struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SiliconPodSpec   `json:"spec,omitempty"`
	Status SiliconPodStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SiliconPodList contains a list of SiliconPod
type SiliconPodList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SiliconPod `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SiliconPod{}, &SiliconPodList{})
}
