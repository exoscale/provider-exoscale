// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	antiaffinitygroup "github.com/exoscale/provider-exoscale/internal/controller/namespaced/compute/antiaffinitygroup"
	blockstoragevolume "github.com/exoscale/provider-exoscale/internal/controller/namespaced/compute/blockstoragevolume"
	elasticip "github.com/exoscale/provider-exoscale/internal/controller/namespaced/compute/elasticip"
	instance "github.com/exoscale/provider-exoscale/internal/controller/namespaced/compute/instance"
	instancepool "github.com/exoscale/provider-exoscale/internal/controller/namespaced/compute/instancepool"
	nlb "github.com/exoscale/provider-exoscale/internal/controller/namespaced/compute/nlb"
	nlbservice "github.com/exoscale/provider-exoscale/internal/controller/namespaced/compute/nlbservice"
	privatenetwork "github.com/exoscale/provider-exoscale/internal/controller/namespaced/compute/privatenetwork"
	securitygroup "github.com/exoscale/provider-exoscale/internal/controller/namespaced/compute/securitygroup"
	securitygrouprules "github.com/exoscale/provider-exoscale/internal/controller/namespaced/compute/securitygrouprules"
	skscluster "github.com/exoscale/provider-exoscale/internal/controller/namespaced/compute/skscluster"
	sksnodepool "github.com/exoscale/provider-exoscale/internal/controller/namespaced/compute/sksnodepool"
	sshkey "github.com/exoscale/provider-exoscale/internal/controller/namespaced/compute/sshkey"
	dbaasdatabasemysql "github.com/exoscale/provider-exoscale/internal/controller/namespaced/dbaas/dbaasdatabasemysql"
	dbaasdatabasepg "github.com/exoscale/provider-exoscale/internal/controller/namespaced/dbaas/dbaasdatabasepg"
	dbaasservice "github.com/exoscale/provider-exoscale/internal/controller/namespaced/dbaas/dbaasservice"
	dbaasuserkafka "github.com/exoscale/provider-exoscale/internal/controller/namespaced/dbaas/dbaasuserkafka"
	dbaasusermysql "github.com/exoscale/provider-exoscale/internal/controller/namespaced/dbaas/dbaasusermysql"
	dbaasuserpg "github.com/exoscale/provider-exoscale/internal/controller/namespaced/dbaas/dbaasuserpg"
	providerconfig "github.com/exoscale/provider-exoscale/internal/controller/namespaced/providerconfig"
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
		dbaasdatabasemysql.Setup,
		dbaasdatabasepg.Setup,
		dbaasservice.Setup,
		dbaasuserkafka.Setup,
		dbaasusermysql.Setup,
		dbaasuserpg.Setup,
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
		dbaasdatabasemysql.SetupGated,
		dbaasdatabasepg.SetupGated,
		dbaasservice.SetupGated,
		dbaasuserkafka.SetupGated,
		dbaasusermysql.SetupGated,
		dbaasuserpg.SetupGated,
		providerconfig.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
