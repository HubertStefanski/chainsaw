package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Describe defines how to describe resources.
type Describe struct {
	// Timeout for the operation. Overrides the global timeout set in the Configuration.
	// +optional
	Timeout *metav1.Duration `json:"timeout,omitempty"`

	// Resource type.
	Resource string `json:"resource"`

	// Namespace of the referent.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	// +optional
	Namespace string `json:"namespace,omitempty"`

	// Name of the referent.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	// +optional
	Name string `json:"name,omitempty"`

	// Selector defines labels selector.
	// +optional
	Selector string `json:"selector,omitempty"`

	// Show Events indicates whether to include related events.
	// +optional
	ShowEvents *bool `json:"showEvents,omitempty"`
}
