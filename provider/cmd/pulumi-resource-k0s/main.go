package main

import (
	"fmt"
	"os"

	p "github.com/pulumi/pulumi-go-provider"
	k0s "github.com/ydkn/pulumi-k0s/provider"
)

// Serve the provider against Pulumi's Provider protocol.
func main() {
	if err := p.RunProvider(k0s.Name, k0s.Version, k0s.Provider()); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())

		os.Exit(1)
	}
}
