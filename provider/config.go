package provider

import (
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Config struct {
	SkipDowngradeCheck *bool `pulumi:"skipDowngradeCheck,optional"`
	NoDrain            *bool `pulumi:"noDrain,optional"`
}

func (c *Config) Annotate(a infer.Annotator) {
	a.Describe(&c.SkipDowngradeCheck, "Do not check if a downgrade would be performed.")
	//a.SetDefault(&c.SkipDowngradeCheck, false, "PULUMI_K0S_SKIP_DOWNGRADE_CHECK")

	a.Describe(&c.NoDrain, "Do not drain nodes before upgrades/updates.")
	//a.SetDefault(&c.NoDrain, false, "PULUMI_K0S_NO_DRAIN")
}
