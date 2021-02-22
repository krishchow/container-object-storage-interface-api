/*
Copyright 2020 The Kubernetes Authors.

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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type BucketRequestSpec struct {
	// +optional
	BucketInstanceName string `json:"bucketInstanceName,omitempty"`
	// +optional
	BucketPrefix string `json:"bucketPrefix,omitempty"`
	// +optional
	BucketClassName string `json:"bucketClassName,omitempty"`
}

type BucketRequestStatus struct {
	// +optional
	Message string `json:"message,omitempty"`
	// +optional
	BucketAvailable bool `json:"bucketAvailable"`
}

type AnonymousAccessMode struct {
	// +optional
	Private bool `json:"private,omitempty"`
	// +optional
	PublicReadOnly bool `json:"publicReadOnly,omitempty"`
	// +optional
	PublicReadWrite bool `json:"publicReadWrite,omitempty"`
	// +optional
	PublicWriteOnly bool `json:"publicWriteOnly,omitempty"`
}

type BucketSpec struct {
	// +optional
	Provisioner string `json:"provisioner,omitempty"`
	// +kubebuilder:default:=retain
	RetentionPolicy RetentionPolicy `json:"retentionPolicy"`
	// +optional
	AnonymousAccessMode AnonymousAccessMode `json:"anonymousAccessMode,omitempty"`
	// +optional
	BucketClassName string                  `json:"bucketClassName,omitempty"`
	BucketRequest   *corev1.ObjectReference `json:"bucketRequest,omitempty"`
	// +listType=atomic
	AllowedNamespaces []string `json:"allowedNamespaces,omitempty"`
	Protocol          Protocol `json:"protocol"`
	// +optional
	Parameters map[string]string `json:"parameters,omitempty"`
}

type BucketStatus struct {
	// +optional
	Message string `json:"message,omitempty"`
	// +optional
	BucketAvailable bool `json:"bucketAvailable,omitempty"`
}

type RetentionPolicy string

const (
	RetentionPolicyRetain RetentionPolicy = "Retain"
	RetentionPolicyDelete RetentionPolicy = "Delete"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

type BucketRequest struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BucketRequestSpec `json:"spec,omitempty"`
	// +optional
	Status BucketRequestStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type BucketRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketRequest `json:"items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:storageversion
// +kubebuilder:subresource:status

type Bucket struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BucketSpec `json:"spec,omitempty"`
	// +optional
	Status BucketStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type BucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Bucket `json:"items"`
}

// +genclient
// +genclient:nonNamespaced
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:storageversion

type BucketClass struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +optional
	// +kubebuilder:default:=false
	IsDefaultBucketClass bool `json:"isDefaultBucketClass,omitempty"`
	// +listType=atomic
	// +optional
	AllowedNamespaces []string `json:"allowedNamespaces,omitempty"`
	Protocol          Protocol `json:"protocol"`
	// +optional
	AnonymousAccessMode AnonymousAccessMode `json:"anonymousAccessMode,omitempty"`
	// +kubebuilder:default:=retain
	RetentionPolicy RetentionPolicy `json:"retentionPolicy,omitempty"`
	// +optional
	Parameters map[string]string `json:"parameters,omitempty"`
	// +optional
	Provisioner string `json:"provisioner,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type BucketClassList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketClass `json:"items"`
}

// +genclient
// +genclient:nonNamespaced
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:storageversion

type BucketAccessClass struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +optional
	Provisioner            string                  `json:"provisioner,omitempty"`
	PolicyActionsConfigMap *corev1.ObjectReference `json:"policyActionsConfigMap,omitempty"`
	// +optional
	Parameters map[string]string `json:"parameters,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type BucketAccessClassList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketAccessClass `json:"items"`
}

type BucketAccessSpec struct {
	// +optional
	BucketInstanceName string `json:"bucketInstanceName,omitempty"`
	// +optional
	BucketAccessRequest *corev1.ObjectReference `json:"bucketAccessRequest,omitempty"`
	// +optional
	ServiceAccount *corev1.ObjectReference `json:"serviceAccount,omitempty"`
	// +optional
	MintedSecretName           string `json:"mintedSecretName,omitempty"`
	PolicyActionsConfigMapData string `json:"policyActionsConfigMapData,omitempty"`
	// +optional
	Principal string `json:"principal,omitempty"`
	// +optional
	Provisioner string `json:"provisioner,omitempty"`
	// +optional
	Parameters map[string]string `json:"parameters,omitempty"`
}

type BucketAccessStatus struct {
	// +optional
	Message string `json:"message,omitempty"`
	// +optional
	AccessGranted bool `json:"accessGranted,omitempty"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:storageversion
// +kubebuilder:subresource:status

type BucketAccess struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BucketAccessSpec `json:"spec,omitempty"`
	// +optional
	Status BucketAccessStatus `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type BucketAccessList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketAccess `json:"items"`
}

type BucketAccessRequestSpec struct {
	// +optional
	ServiceAccountName    string `json:"serviceAccountName,omitempty"`
	BucketRequestName     string `json:"bucketRequestName"`
	BucketAccessClassName string `json:"bucketAccessClassName"`
	// +optional
	BucketAccessName string `json:"bucketAccessName,omitempty"`
}

type BucketAccessRequestStatus struct {
	// +optional
	Message string `json:"message,omitempty"`
	// +optional
	AccessGranted bool `json:"accessGranted"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:storageversion
// +kubebuilder:subresource:status

type BucketAccessRequest struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BucketAccessRequestSpec `json:"spec,omitempty"`
	// +optional
	Status BucketAccessRequestStatus `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type BucketAccessRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketAccessRequest `json:"items"`
}


type COSIDriverSpec struct {
	// +optional
	Provisioner string `json:"provisioner,omitempty"`
}

type COSIDriverStatus struct {
	// +optional
	Registered bool `json:"registered"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:storageversion
// +kubebuilder:subresource:status

type COSIDriver struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              COSIDriverSpec `json:"spec,omitempty"`
	// +optional
	Status COSIDriverStatus `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type COSIDriverList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []COSIDriver `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Bucket{}, &BucketList{})
	SchemeBuilder.Register(&BucketRequest{}, &BucketRequestList{})
	SchemeBuilder.Register(&BucketClass{}, &BucketClassList{})

	SchemeBuilder.Register(&BucketAccess{}, &BucketAccessList{})
	SchemeBuilder.Register(&BucketAccessRequest{}, &BucketAccessRequestList{})
	SchemeBuilder.Register(&BucketAccessClass{}, &BucketAccessClassList{})

	SchemeBuilder.Register(&COSIDriver{}, &COSIDriverList{})
}
