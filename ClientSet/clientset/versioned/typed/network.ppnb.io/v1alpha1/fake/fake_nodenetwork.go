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

package fake

import (
	networkppnbiov1alpha1 "github.com/Frank-svg-dev/pam-cni/ClientSet/clientset/versioned/typed/network.ppnb.io/v1alpha1"
	v1alpha1 "github.com/Frank-svg-dev/pam-cni/pkg/apis/network.ppnb.io/v1alpha1"
	gentype "k8s.io/client-go/gentype"
)

// fakeNodeNetworks implements NodeNetworkInterface
type fakeNodeNetworks struct {
	*gentype.FakeClientWithList[*v1alpha1.NodeNetwork, *v1alpha1.NodeNetworkList]
	Fake *FakeNetworkV1alpha1
}

func newFakeNodeNetworks(fake *FakeNetworkV1alpha1) networkppnbiov1alpha1.NodeNetworkInterface {
	return &fakeNodeNetworks{
		gentype.NewFakeClientWithList[*v1alpha1.NodeNetwork, *v1alpha1.NodeNetworkList](
			fake.Fake,
			"",
			v1alpha1.SchemeGroupVersion.WithResource("nodenetworks"),
			v1alpha1.SchemeGroupVersion.WithKind("NodeNetwork"),
			func() *v1alpha1.NodeNetwork { return &v1alpha1.NodeNetwork{} },
			func() *v1alpha1.NodeNetworkList { return &v1alpha1.NodeNetworkList{} },
			func(dst, src *v1alpha1.NodeNetworkList) { dst.ListMeta = src.ListMeta },
			func(list *v1alpha1.NodeNetworkList) []*v1alpha1.NodeNetwork {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v1alpha1.NodeNetworkList, items []*v1alpha1.NodeNetwork) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
