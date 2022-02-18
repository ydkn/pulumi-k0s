import * as pulumi from "@pulumi/pulumi";
import * as k0s from "@ydkn/pulumi-k0s";

const provider = new k0s.Provider("k0s", { noDrain: true });

const cluster = new k0s.Cluster("my-cluster", {
  metadata: { name: "my-cluster" },
  spec: {
    hosts: [
      {
        role: "controller",
        ssh: { address: "10.0.0.1" },
      },
      {
        role: "worker",
        ssh: { address: "10.0.0.2" },
      },
      {
        role: "worker",
        ssh: { address: "10.0.0.3" },
      },
    ],
  },
});

export const kubeconfig = pulumi.secret(cluster.kubeconfig);
