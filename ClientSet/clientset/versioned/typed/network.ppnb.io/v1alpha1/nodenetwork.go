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
// Code ClientSet by client-gen. DO NOT EDIT.

package v1alpha1

import (
	context "context"

	"github.com/Frank-svg-dev/pam-cni/ClientSet/clientset/versioned/scheme"
	networkppnbiov1alpha1 "github.com/Frank-svg-dev/pam-cni/pkg/apis/network.ppnb.io/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// NodeNetworksGetter has a method to return a NodeNetworkInterface.
// A group's client should implement this interface.
type NodeNetworksGetter interface {
	NodeNetworks() NodeNetworkInterface
}

// NodeNetworkInterface has methods to work with NodeNetwork resources.
type NodeNetworkInterface interface {
	Create(ctx context.Context, nodeNetwork *networkppnbiov1alpha1.NodeNetwork, opts v1.CreateOptions) (*networkppnbiov1alpha1.NodeNetwork, error)
	Update(ctx context.Context, nodeNetwork *networkppnbiov1alpha1.NodeNetwork, opts v1.UpdateOptions) (*networkppnbiov1alpha1.NodeNetwork, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, nodeNetwork *networkppnbiov1alpha1.NodeNetwork, opts v1.UpdateOptions) (*networkppnbiov1alpha1.NodeNetwork, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*networkppnbiov1alpha1.NodeNetwork, error)
	List(ctx context.Context, opts v1.ListOptions) (*networkppnbiov1alpha1.NodeNetworkList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *networkppnbiov1alpha1.NodeNetwork, err error)
	NodeNetworkExpansion
}

// nodeNetworks implements NodeNetworkInterface
type nodeNetworks struct {
	*gentype.ClientWithList[*networkppnbiov1alpha1.NodeNetwork, *networkppnbiov1alpha1.NodeNetworkList]
}

// newNodeNetworks returns a NodeNetworks
func newNodeNetworks(c *NetworkV1alpha1Client) *nodeNetworks {
	return &nodeNetworks{
		gentype.NewClientWithList[*networkppnbiov1alpha1.NodeNetwork, *networkppnbiov1alpha1.NodeNetworkList](
			"nodenetworks",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *networkppnbiov1alpha1.NodeNetwork { return &networkppnbiov1alpha1.NodeNetwork{} },
			func() *networkppnbiov1alpha1.NodeNetworkList { return &networkppnbiov1alpha1.NodeNetworkList{} },
		),
	}
}
