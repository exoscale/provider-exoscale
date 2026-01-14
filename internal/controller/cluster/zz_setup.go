// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	antiaffinitygroup "github.com/exoscale/provider-exoscale/internal/controller/cluster/compute/antiaffinitygroup"
	blockstoragevolume "github.com/exoscale/provider-exoscale/internal/controller/cluster/compute/blockstoragevolume"
	elasticip "github.com/exoscale/provider-exoscale/internal/controller/cluster/compute/elasticip"
	instance "github.com/exoscale/provider-exoscale/internal/controller/cluster/compute/instance"
	instancepool "github.com/exoscale/provider-exoscale/internal/controller/cluster/compute/instancepool"
	nlb "github.com/exoscale/provider-exoscale/internal/controller/cluster/compute/nlb"
	nlbservice "github.com/exoscale/provider-exoscale/internal/controller/cluster/compute/nlbservice"
	privatenetwork "github.com/exoscale/provider-exoscale/internal/controller/cluster/compute/privatenetwork"
	securitygroup "github.com/exoscale/provider-exoscale/internal/controller/cluster/compute/securitygroup"
	securitygrouprules "github.com/exoscale/provider-exoscale/internal/controller/cluster/compute/securitygrouprules"
	skscluster "github.com/exoscale/provider-exoscale/internal/controller/cluster/compute/skscluster"
	sksnodepool "github.com/exoscale/provider-exoscale/internal/controller/cluster/compute/sksnodepool"
	sshkey "github.com/exoscale/provider-exoscale/internal/controller/cluster/compute/sshkey"
	providerconfig "github.com/exoscale/provider-exoscale/internal/controller/cluster/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		antiaffinitygroup.Setup,
		blockstoragevolume.Setup,
		elasticip.Setup,
		instance.Setup,
		instancepool.Setup,
		nlb.Setup,
		nlbservice.Setup,
		privatenetwork.Setup,
		securitygroup.Setup,
		securitygrouprules.Setup,
		skscluster.Setup,
		sksnodepool.Setup,
		sshkey.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		antiaffinitygroup.SetupGated,
		blockstoragevolume.SetupGated,
		elasticip.SetupGated,
		instance.SetupGated,
		instancepool.SetupGated,
		nlb.SetupGated,
		nlbservice.SetupGated,
		privatenetwork.SetupGated,
		securitygroup.SetupGated,
		securitygrouprules.SetupGated,
		skscluster.SetupGated,
		sksnodepool.SetupGated,
		sshkey.SetupGated,
		providerconfig.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
