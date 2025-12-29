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
	securitygroup "github.com/exoscale/provider-exoscale/internal/controller/cluster/compute/securitygroup"
	securitygrouprules "github.com/exoscale/provider-exoscale/internal/controller/cluster/compute/securitygrouprules"
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
		securitygroup.Setup,
		securitygrouprules.Setup,
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
		securitygroup.SetupGated,
		securitygrouprules.SetupGated,
		sshkey.SetupGated,
		providerconfig.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
