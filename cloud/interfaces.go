/*
Copyright 2021 The Kubernetes Authors.

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

package cloud

import (
	"context"

	"github.com/GoogleCloudPlatform/k8s-cloud-provider/pkg/cloud"
	corev1 "k8s.io/api/core/v1"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1beta1"
	capierrors "sigs.k8s.io/cluster-api/errors"

	infrav1 "sigs.k8s.io/cluster-api-provider-gcp/api/v1beta1"
)

// Cloud alias for cloud.Cloud interface.
type Cloud = cloud.Cloud

// Reconciler is a generic interface used by components offering a type of service.
type Reconciler interface {
	Reconcile(ctx context.Context) error
	Delete(ctx context.Context) error
}

// Client is an interface which can get cloud client.
type Client interface {
	Cloud() Cloud
}

// ClusterGetter is an interface which can get cluster informations.
type ClusterGetter interface {
	Client
	Project() string
	Region() string
	Name() string
	Namespace() string
	NetworkName() string
	Network() *infrav1.Network
	AdditionalLabels() infrav1.Labels
	FailureDomains() clusterv1.FailureDomains
	ControlPlaneEndpoint() clusterv1.APIEndpoint
}

// ClusterSetter is an interface which can set cluster informations.
type ClusterSetter interface {
	SetControlPlaneEndpoint(endpoint clusterv1.APIEndpoint)
}

// Cluster is an interface which can get and set cluster informations.
type Cluster interface {
	ClusterGetter
	ClusterSetter
}

// MachineGetter is an interface which can get machine informations.
type MachineGetter interface {
	Client
	Name() string
	Namespace() string
	Zone() string
	Project() string
	Role() string
	IsControlPlane() bool
	ControlPlaneGroupName() string
	GetInstanceID() *string
	GetProviderID() string
	GetBootstrapData() (string, error)
	GetInstanceStatus() *infrav1.InstanceStatus
}

// MachineSetter is an interface which can set machine informations.
type MachineSetter interface {
	SetProviderID()
	SetInstanceStatus(v infrav1.InstanceStatus)
	SetFailureMessage(v error)
	SetFailureReason(v capierrors.MachineStatusError)
	SetAnnotation(key, value string)
	SetAddresses(addressList []corev1.NodeAddress)
}

// Machine is an interface which can get and set machine informations.
type Machine interface {
	MachineGetter
	MachineSetter
}

type MachinePoolSetter interface {
}

type MachinePoolGetter interface {
	Client
	Name() string
	Zone() string
	GetBootstrapData(ctx context.Context) (string, error)
}

type MachinePool interface {
	MachinePoolGetter
	MachinePoolSetter
}
