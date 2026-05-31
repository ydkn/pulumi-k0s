// Copyright 2025, Florian Schwab.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	"bufio"
	"bytes"
	"context"

	"github.com/k0sproject/k0sctl/action"
	"github.com/k0sproject/k0sctl/phase"
)

type K0sctl struct {
	spec       *ClusterArgs
	kubeconfig *string
}

func NewK0sctl(args *ClusterArgs) *K0sctl {
	return &K0sctl{spec: args}
}

func (k *K0sctl) Validate() error {
	cluster, cleanup, err := k.spec.k0sctl()

	defer cleanup()

	if err != nil {
		return err
	}

	return cluster.Validate()
}

func (k *K0sctl) Apply(ctx context.Context, config *Config) error {
	cluster, cleanup, err := k.spec.k0sctl()

	defer cleanup()

	if err != nil {
		return err
	}

	skipDowngradeCheck := configDefaultSkipDowngradeCheck
	if config.SkipDowngradeCheck != nil {
		skipDowngradeCheck = *config.SkipDowngradeCheck
	}

	noDrain := configDefaultNoDrain
	if config.NoDrain != nil {
		noDrain = *config.NoDrain
	}

	noWait := configDefaultNoWait
	if config.NoWait != nil {
		noWait = *config.NoWait
	}

	concurrency := configDefaultConcurrency
	if config.Concurrency != nil {
		concurrency = *config.Concurrency
	}

	concurrentUploads := configDefaultConcurrentUploads
	if config.ConcurrentUploads != nil {
		concurrentUploads = *config.ConcurrentUploads
	}

	manager := phase.Manager{
		Config:            cluster,
		Concurrency:       concurrency,
		ConcurrentUploads: concurrentUploads,
	}

	var kubeconfigBytes bytes.Buffer

	kubeconfigWriter := bufio.NewWriter(&kubeconfigBytes)

	applyAction := action.Apply{
		ApplyOptions: action.ApplyOptions{
			Manager:               &manager,
			KubeconfigOut:         kubeconfigWriter,
			KubeconfigAPIAddress:  k.spec.APIAddress(),
			NoWait:                noWait,
			NoDrain:               noDrain,
			DisableDowngradeCheck: skipDowngradeCheck,
			RestoreFrom:           "",
		},
	}

	if err := applyAction.Run(ctx); err != nil {
		return err
	}

	if manager.Config.Metadata != nil && manager.Config.Metadata.Kubeconfig != "" {
		kubeconfig := kubeconfigBytes.String()
		k.kubeconfig = &kubeconfig
	}

	return nil
}

func (k *K0sctl) Kubeconfig(ctx context.Context) error {
	cluster, cleanup, err := k.spec.k0sctl()

	defer cleanup()

	if err != nil {
		return err
	}

	manager := phase.Manager{Config: cluster}

	kubeconfigAction := action.Kubeconfig{
		Manager:              &manager,
		KubeconfigAPIAddress: k.spec.APIAddress(),
	}

	if err := kubeconfigAction.Run(ctx); err != nil {
		return err
	}

	if manager.Config.Metadata != nil && manager.Config.Metadata.Kubeconfig != "" {
		k.kubeconfig = &manager.Config.Metadata.Kubeconfig
	}

	return nil
}

func (k *K0sctl) Reset(ctx context.Context) error {
	cluster, cleanup, err := k.spec.k0sctl()

	defer cleanup()

	if err != nil {
		return err
	}

	manager := phase.Manager{Config: cluster}

	resetAction := action.Reset{
		Manager: &manager,
		Stdout:  nil,
	}

	if err := resetAction.Run(ctx); err != nil {
		return err
	}

	return nil
}
