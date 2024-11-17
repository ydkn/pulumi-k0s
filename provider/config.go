package provider

import (
	"context"
	"strings"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/ydkn/pulumi-k0s/provider/internal/introspect"
)

type Config struct {
	SkipDowngradeCheck *bool `pulumi:"skipDowngradeCheck,optional"`
	NoDrain            *bool `pulumi:"noDrain,optional"`
	NoWait             *bool `pulumi:"noWait,optional"`
	Concurrency        *int  `pulumi:"concurrency,optional"`
	ConcurrentUploads  *int  `pulumi:"concurrentUploads,optional"`
}

func (c *Config) Annotate(a infer.Annotator) {
	skipDowngradeCheckValue := false
	a.Describe(&c.SkipDowngradeCheck, "Skip downgrade check")
	a.SetDefault(&c.SkipDowngradeCheck, &skipDowngradeCheckValue, "PULUMI_K0S_SKIP_DOWNGRADE_CHECK")

	noDrainValue := false
	a.Describe(&c.NoDrain, "Do not drain worker nodes when upgrading")
	a.SetDefault(&c.NoDrain, &noDrainValue, "PULUMI_K0S_NO_DRAIN")

	noWaitValue := false
	a.Describe(&c.NoWait, "Do not wait for worker nodes to join")
	a.SetDefault(&c.NoWait, &noWaitValue, "PULUMI_K0S_NO_WAIT")

	concurrencyValue := 30
	a.Describe(&c.Concurrency, "Maximum number of hosts to configure in parallel, set to 0 for unlimited")
	a.SetDefault(&c.Concurrency, &concurrencyValue, "PULUMI_K0S_CONCURRENCY")

	concurrentUploadsValue := 5
	a.Describe(&c.ConcurrentUploads, "Maximum number of files to upload in parallel, set to 0 for unlimited")
	a.SetDefault(&c.ConcurrentUploads, &concurrentUploadsValue, "PULUMI_K0S_CONCURRENT_UPLOADS")
}

func (c *Config) Diff(ctx context.Context, name string, olds Config, news Config) (p.DiffResponse, error) {
	diffResponse := p.DiffResponse{
		DeleteBeforeReplace: false,
		HasChanges:          false,
		DetailedDiff:        map[string]p.PropertyDiff{},
	}

	oldsProps, err := introspect.NewPropertiesMap(olds)
	if err != nil {
		return p.DiffResponse{}, err
	}

	newsProps, err := introspect.NewPropertiesMap(news)
	if err != nil {
		return p.DiffResponse{}, err
	}

	for key := range propertyMapDiff(oldsProps, newsProps, []resource.PropertyKey{}) {
		diffResponse.DetailedDiff[strings.SplitN(string(key), ".", 2)[0]] = p.PropertyDiff{
			Kind:      p.Update,
			InputDiff: true,
		}
	}

	if len(diffResponse.DetailedDiff) > 0 {
		diffResponse.HasChanges = true
	}

	return diffResponse, nil
}

func (c *Config) Check(ctx context.Context, name string, olds Config, news Config) (Config, []p.CheckFailure, error) {
	return Config{
		SkipDowngradeCheck: news.SkipDowngradeCheck,
		NoDrain:            news.NoDrain,
		NoWait:             news.NoWait,
		Concurrency:        news.Concurrency,
		ConcurrentUploads:  news.ConcurrentUploads,
	}, []p.CheckFailure{}, nil
}
