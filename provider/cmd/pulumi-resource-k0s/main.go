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

package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/k0sproject/rig"
	"github.com/sirupsen/logrus"
	k0s "github.com/ydkn/pulumi-k0s/provider"
)

// Serve the provider against Pulumi's Provider protocol.
func main() {
	// Disable output of k0sctl
	logrus.SetOutput(io.Discard)
	rig.SetLogger(logrus.StandardLogger())

	err := k0s.Provider().Run(context.Background(), k0s.Name, k0s.Version)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(1)
	}
}
