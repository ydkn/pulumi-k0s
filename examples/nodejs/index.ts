import * as k0s from "../../sdk/nodejs/bin/index.js"; // "@ydkn/pulumi-k0s";

const myProvider = new k0s.Provider("myProvider", {});
const myCluster = new k0s.Cluster(
  "myCluster",
  {
    spec: {
      hosts: [
        {
          role: "controller+worker",
          hostname: "my-node",
          localhost: { enabled: true },
        },
      ],
    },
  },
  {
    provider: myProvider,
  },
);
export const output = {
  value: myCluster.kubeconfig,
};
