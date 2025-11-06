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
// Code ClientSet by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	context "context"
	time "time"

	"github.com/Frank-svg-dev/pam-cni/ClientSet/clientset/versioned"
	"github.com/Frank-svg-dev/pam-cni/ClientSet/informers/externalversions/internalinterfaces"
	networkppnbiov1alpha1 "github.com/Frank-svg-dev/pam-cni/ClientSet/listers/network.ppnb.io/v1alpha1"
	apisnetworkppnbiov1alpha1 "github.com/Frank-svg-dev/pam-cni/pkg/apis/network.ppnb.io/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// NodeNetworkInformer provides access to a shared informer and lister for
// NodeNetworks.
type NodeNetworkInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() networkppnbiov1alpha1.NodeNetworkLister
}

type nodeNetworkInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewNodeNetworkInformer constructs a new informer for NodeNetwork type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewNodeNetworkInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredNodeNetworkInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredNodeNetworkInformer constructs a new informer for NodeNetwork type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredNodeNetworkInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkV1alpha1().NodeNetworks().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkV1alpha1().NodeNetworks().Watch(context.TODO(), options)
			},
		},
		&apisnetworkppnbiov1alpha1.NodeNetwork{},
		resyncPeriod,
		indexers,
	)
}

func (f *nodeNetworkInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredNodeNetworkInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *nodeNetworkInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apisnetworkppnbiov1alpha1.NodeNetwork{}, f.defaultInformer)
}

func (f *nodeNetworkInformer) Lister() networkppnbiov1alpha1.NodeNetworkLister {
	return networkppnbiov1alpha1.NewNodeNetworkLister(f.Informer().GetIndexer())
}
