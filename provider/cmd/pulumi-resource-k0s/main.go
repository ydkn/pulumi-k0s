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
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/k0sproject/rig"
	presource "github.com/pulumi/pulumi/sdk/v3/go/common/resource"
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

const abc = `{"metadata":{"V":{"name":{"V":"c49544"}}},"spec":{"V":{"Element":{"V":{"hosts":{"V":[{"V":{"hooks":{"V":{"apply":{"V":{"after":{"V":[]},"before":{"V":[{"V":"sudo sh -c 'if [ ! -L /var/lib/kubelet ]; then ln -s /var/lib/k0s/kubelet /var/lib/kubelet; fi'"}]}}},"backup":{"V":{"after":{"V":[]}}},"reset":{"V":{"after":{"V":[{"V":"sudo sh -c 'if [ -L /var/lib/kubelet ]; then rm -rf /var/lib/kubelet; fi'"}]}}}}},"hostname":{"V":"c49544-controller-1256c6"},"installFlags":{"V":[{"V":"--enable-k0s-cloud-provider=true"},{"V":"--enable-cloud-provider=true"},{"V":"--labels=[secret].network/gateway=true"}]},"noTaints":{"V":true},"openSSH":{"V":{"address":{"V":"172.28.128.32"},"disableMultiplexing":{"V":true},"key":{"V":{"Element":{"V":"-----BEGIN OPENSSH PRIVATE KEY-----\nb3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtz\nc2gtZWQyNTUxOQAAACDOKcWz+8oSAnhqN553+VPsVBIfAoyAL21BtCa4zLaKXAAA\nAIhgIn1XYCJ9VwAAAAtzc2gtZWQyNTUxOQAAACDOKcWz+8oSAnhqN553+VPsVBIf\nAoyAL21BtCa4zLaKXAAAAEDxeoU2yk4MdHXUKg5LEINoKw19japGZjRvOnLmF+rZ\nR84pxbP7yhICeGo3nnf5U+xUEh8CjIAvbUG0JrjMtopcAAAAAAECAwQF\n-----END OPENSSH PRIVATE KEY-----\n"}}},"port":{"V":22},"user":{"V":"pulumi"}}},"privateAddress":{"V":"172.28.128.32"},"role":{"V":"controller+worker"},"uploadBinary":{"V":false}}}]},"k0s":{"V":{"config":{"V":{"spec":{"V":{"api":{"V":{"address":{"V":"172.28.128.32"},"sans":{"V":[{"V":"api.c49544.[secret].network"}]}}},"network":{"V":{"calico":{"V":{"ipAutodetectionMethod":{"V":"interface=^(end0|enabcm6e4ei0|eth0)$"},"mode":{"V":"bird"},"wireguard":{"V":false}}},"clusterDomain":{"V":"cluster.c49544.[secret].network"},"dualStack":{"V":{"IPv6podCIDR":{"V":"fdd0:cafe:af02:10::/96"},"IPv6serviceCIDR":{"V":"fdd0:cafe:af02:11::/112"},"enabled":{"V":true}}},"kubeProxy":{"V":{"mode":{"V":"nftables"}}},"podCIDR":{"V":"172.28.144.0/20"},"provider":{"V":"calico"},"serviceCIDR":{"V":"172.28.160.0/22"}}},"telemetry":{"V":{"enabled":{"V":false}}}}}}}}}}}}}}`

func TestAbc() {
	m := presource.PropertyMap{}

	_ = json.Unmarshal([]byte(abc), &m)
	fmt.Println(m)

	b := presource.FromResourcePropertyMap(m)
	fmt.Println(b)
}
