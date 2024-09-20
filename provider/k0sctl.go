package provider

import (
	"bufio"
	"bytes"
	"strconv"
	"strings"

	"github.com/k0sproject/k0sctl/action"
	"github.com/k0sproject/k0sctl/phase"
)

type K0sctl struct {
	news    *ClusterInputs
	cluster k0sctlCluster
}

func NewK0sctl(news *ClusterInputs) *K0sctl {
	return &K0sctl{news: news, cluster: k0sctlCluster(*news)}
}

func (k *K0sctl) Validate() error {
	return k.validate()
}

func (k *K0sctl) Apply(config *Config) error {
	cluster, cleanup, err := k.cluster.k0sctl()

	defer cleanup()

	if err != nil {
		return err
	}

	skipDowngradeCheck := false
	if config.SkipDowngradeCheck != nil {
		skipDowngradeCheck = strings.ToLower(*config.SkipDowngradeCheck) == "true"
	}

	noDrain := false
	if config.NoDrain != nil {
		noDrain = strings.ToLower(*config.NoDrain) == "true"
	}

	noWait := false
	if config.NoWait != nil {
		noWait = strings.ToLower(*config.NoWait) == "true"
	}

	concurrency := 30
	if config.Concurrency != nil {
		i, err := strconv.ParseUint(*config.Concurrency, 10, 64)
		if err != nil {
			return err
		}

		concurrency = int(i)
	}

	concurrentUploads := 5
	if config.ConcurrentUploads != nil {
		i, err := strconv.ParseUint(*config.ConcurrentUploads, 10, 64)
		if err != nil {
			return err
		}

		concurrentUploads = int(i)
	}

	manager := phase.Manager{Config: cluster, Concurrency: concurrency, ConcurrentUploads: concurrentUploads}

	var kubeconfigBytes bytes.Buffer

	kubeconfigWriter := bufio.NewWriter(&kubeconfigBytes)

	applyAction := action.Apply{
		Force:                 false,
		Manager:               &manager,
		KubeconfigOut:         kubeconfigWriter,
		KubeconfigAPIAddress:  k.cluster.APIAddress(),
		NoWait:                noWait,
		NoDrain:               noDrain,
		DisableDowngradeCheck: skipDowngradeCheck,
		RestoreFrom:           "",
	}

	if err := applyAction.Run(); err != nil {
		return err
	}

	if manager.Config.Metadata != nil && manager.Config.Metadata.Kubeconfig != "" {
		kubeconfig := kubeconfigBytes.String()
		k.news.Kubeconfig = &kubeconfig
	}

	return nil
}

func (k *K0sctl) Kubeconfig() error {
	cluster, cleanup, err := k.cluster.k0sctl()

	defer cleanup()

	if err != nil {
		return err
	}

	manager := phase.Manager{Config: cluster}

	kubeconfigAction := action.Kubeconfig{
		Manager:              &manager,
		KubeconfigAPIAddress: k.cluster.APIAddress(),
	}

	if err := kubeconfigAction.Run(); err != nil {
		return err
	}

	if manager.Config.Metadata != nil && manager.Config.Metadata.Kubeconfig != "" {
		k.news.Kubeconfig = &manager.Config.Metadata.Kubeconfig
	}

	return nil
}

func (k *K0sctl) Reset() error {
	cluster, cleanup, err := k.cluster.k0sctl()

	defer cleanup()

	if err != nil {
		return err
	}

	manager := phase.Manager{Config: cluster}

	resetAction := action.Reset{
		Manager: &manager,
		Force:   true,
		Stdout:  nil,
	}

	if err := resetAction.Run(); err != nil {
		return err
	}

	return nil
}

func (k *K0sctl) validate() error {
	cluster, cleanup, err := k.cluster.k0sctl()

	defer cleanup()

	if err != nil {
		return err
	}

	return cluster.Validate()
}
