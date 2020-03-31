/*


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

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	loadbalancerv1beta1 "github.com/takutakahashi/loadbalancer-controller/api/v1beta1"
	"github.com/takutakahashi/loadbalancer-controller/pkg/terraform"
)

// AWSBackendReconciler reconciles a AWSBackend object
type AWSBackendReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=loadbalancer.takutakahashi.dev,resources=awsbackends,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=loadbalancer.takutakahashi.dev,resources=awsbackends/status,verbs=get;update;patch

func (r *AWSBackendReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	_ = r.Log.WithValues("awsbackend", req.NamespacedName)
	var backend loadbalancerv1beta1.AWSBackend
	err := r.Get(ctx, req.NamespacedName, &backend)
	if err != nil {
		return ctrl.Result{}, err
	}
	tc, err := terraform.NewClientForAWSBackend(backend)
	if err != nil {
		return ctrl.Result{}, err
	}
	err = tc.Apply()
	if err != nil {
		return ctrl.Result{}, err
	}
	return ctrl.Result{}, nil
}

func (r *AWSBackendReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&loadbalancerv1beta1.AWSBackend{}).
		Complete(r)
}
