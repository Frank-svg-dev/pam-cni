/*
Copyright 2025.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// NodeNetworkSpec defines the desired state of NodeNetwork.
// NodeNetworkSpec 定义节点网络的期望状态
// +kubebuilder:printcolumn:name="CIDR",type="string",JSONPath=".spec.cidr",description="Assigned CIDR"
// +kubebuilder:printcolumn:name="NodePreAllocate",type="int",JSONPath=".spec.NodePreAllocate",description="Worker node NodePreAllocate"
type NodeNetworkSpec struct {
	// CIDR 是分配给该节点的网段，例如 "192.168.10.0/24", 由判断后写
	CIDR string `json:"cidr,omitempty"`
	//节点预热几个地址, 默认10个，可以通过 daemon的annotaion来判断
	NodePreAllocate int `json:"nodePreAllocate,omitempty"`
	//节点虚机的uuid
	InstanceID string `json:"instanceID,omitempty"`
	//节点虚机的PortID， 用于daemon预热地址
	//PortID  string `json:"portID,omitempty"`
	//PortMac string `json:"portMac,omitempty"`

	//// NodeName 是当前 CRD 绑定的节点名称
	//NodeName string `json:"nodeName,omitempty"`
}

// NodeNetworkStatus defines the observed state of NodeNetwork.
type NodeNetworkStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Ready bool `json:"ready"`

	// LastSyncTime 是最近一次同步到 Neutron 的时间
	LastSyncTime metav1.Time `json:"lastSyncTime,omitempty"`

	// Message 记录一些调试信息
	Message string `json:"message,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// NodeNetwork is the Schema for the nodenetworks API.
// +kubebuilder:resource:scope=Cluster
// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type NodeNetwork struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NodeNetworkSpec   `json:"spec,omitempty"`
	Status NodeNetworkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
// NodeNetworkList contains a list of NodeNetwork.
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type NodeNetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodeNetwork `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NodeNetwork{}, &NodeNetworkList{})
}
