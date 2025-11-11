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

package controller

import (
	"context"
	"fmt"
	"net"

	networkv1alpha1 "github.com/Frank-svg-dev/pam-cni/api/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// NodeNetworkReconciler reconciles a NodeNetwork object
type NodeNetworkReconciler struct {
	client.Client
	Scheme   *runtime.Scheme
	CidrList []string
	Err      error
}

// +kubebuilder:rbac:groups=network.ppnb.io,resources=nodenetworks,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=network.ppnb.io,resources=nodenetworks/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=network.ppnb.io,resources=nodenetworks/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the NodeNetwork object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.19.4/pkg/reconcile
func (r *NodeNetworkReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)
	nodeNetwork := &networkv1alpha1.NodeNetwork{}

	if err := r.Get(ctx, req.NamespacedName, nodeNetwork); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if nodeNetwork.DeletionTimestamp != nil {
		logger.Info("NodeNetwork is being deleted", "name", nodeNetwork.Name)

		r.CidrList = append([]string{nodeNetwork.Spec.CIDR}, r.CidrList...)

		nodeNetwork.Finalizers = []string{}

		if err := r.Update(ctx, nodeNetwork); err != nil {
			logger.Error(err, fmt.Sprintf("update subnet spec failed %+v", err.Error()))
			return ctrl.Result{}, err
		}

		return ctrl.Result{}, nil
	}

	nnList := &networkv1alpha1.NodeNetworkList{}

	if err := r.List(context.Background(), nnList); err != nil {
		panic(fmt.Sprintf("Failed to get  CidrList, err: %s\n", r.Err.Error()))
	}

	for _, nn := range nnList.Items {
		for k, cidr := range r.CidrList {
			if cidr == nn.Spec.CIDR {
				r.CidrList = append(r.CidrList[:k], r.CidrList[k+1:]...)
			}
		}
	}

	logger.Info(fmt.Sprintf("subnet is %+v", nodeNetwork.Name))

	if nodeNetwork.Spec.CIDR == "" {
		//分配Cidr
		nodeNetwork.Spec.CIDR = r.CidrList[0]
		r.CidrList = r.CidrList[1:]
		nodeNetwork.Finalizers = []string{nodeNetwork.Name}

		logger.Info(fmt.Sprintf("%s 分配 cidr 成功为: %+v, 预计预热 %v 个地址\n", nodeNetwork.Name, nodeNetwork.Spec.CIDR, nodeNetwork.Spec.NodePreAllocate))

		if err := r.Update(ctx, nodeNetwork); err != nil {
			logger.Error(err, fmt.Sprintf("update subnet spec failed %+v", err.Error()))
			return ctrl.Result{}, err
		}
	}

	if nodeNetwork.Spec.NodePreAllocate == 0 {
		nodeNetwork.Spec.NodePreAllocate = 10
	}

	nodeNetwork.Status.Ready = true
	nodeNetwork.Status.Message = "Node CIDR allocation is complete. The remaining steps will be handled by the daemon."

	if err := r.Status().Update(ctx, nodeNetwork); err != nil {
		logger.Error(err, fmt.Sprintf("update subnet status failed %+v", err.Error()))
		return ctrl.Result{}, err
	}

	logger.Info(fmt.Sprintf("subnet %+v, cidr is : %s", nodeNetwork.Name, nodeNetwork.Spec.CIDR))

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *NodeNetworkReconciler) SetupWithManager(mgr ctrl.Manager) error {
	r.CidrList, r.Err = r.generateCidrList("200.0.0.0/16")

	if r.Err != nil || len(r.CidrList) == 0 {
		panic(fmt.Sprintf("Failed to generate CidrList, err: %s\n", r.Err))
	}

	return ctrl.NewControllerManagedBy(mgr).
		For(&networkv1alpha1.NodeNetwork{}).
		Named("nodenetwork").
		Complete(r)
}

func (r *NodeNetworkReconciler) generateCidrList(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var subnets []string

	maskSize, _ := ipnet.Mask.Size()
	if maskSize > 16 {
		return nil, fmt.Errorf("CIDR too small, must be /16 or larger")
	}

	start := ip.To4()
	if start == nil {
		return nil, fmt.Errorf("invalid IP")
	}

	// 计算需要拆多少个 /24
	total := 1 << (24 - maskSize) // /16 -> 256, /8 -> 65536

	for i := 0; i < total; i++ {
		octet2 := int(start[1]) + ((i >> 8) & 0xFF)
		octet3 := int(start[2]) + (i & 0xFF)
		subnet := fmt.Sprintf("%d.%d.%d.0/24", start[0], octet2, octet3)
		subnets = append(subnets, subnet)
	}

	return subnets, nil
}
