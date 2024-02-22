package provider

import (
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Config struct {
	SkipDowngradeCheck *string `pulumi:"skipDowngradeCheck,optional"`
	NoDrain            *string `pulumi:"noDrain,optional"`
	NoWait             *string `pulumi:"noWait,optional"`
	Concurrency        *string `pulumi:"concurrency,optional"`
	ConcurrentUploads  *string `pulumi:"concurrentUploads,optional"`
}

func (c *Config) Annotate(a infer.Annotator) {
	a.Describe(&c.SkipDowngradeCheck, "Skip downgrade check")
	a.SetDefault(&c.SkipDowngradeCheck, "false", "PULUMI_K0S_SKIP_DOWNGRADE_CHECK")

	a.Describe(&c.NoDrain, "Do not drain worker nodes when upgrading")
	a.SetDefault(&c.NoDrain, "false", "PULUMI_K0S_NO_DRAIN")

	a.Describe(&c.NoWait, "Do not wait for worker nodes to join")
	a.SetDefault(&c.NoWait, "false", "PULUMI_K0S_NO_WAIT")

	a.Describe(&c.Concurrency, "Maximum number of hosts to configure in parallel, set to 0 for unlimited")
	a.SetDefault(&c.Concurrency, "30", "PULUMI_K0S_CONCURRENCY")

	a.Describe(&c.ConcurrentUploads, "Maximum number of files to upload in parallel, set to 0 for unlimited")
	a.SetDefault(&c.ConcurrentUploads, "5", "PULUMI_K0S_CONCURRENT_UPLOADS")
}
