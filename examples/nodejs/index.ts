import * as pulumi from "@pulumi/pulumi";
import * as k0s from "@ydkn/pulumi-k0s";

const myProvider = new k0s.Provider("myProvider", {noDrain: "true"});
const myCluster = new k0s.Cluster("myCluster", {spec: {
    hosts: [{
        role: "controller+worker",
        localhost: {
            enabled: true,
        },
    }],
}}, {
    provider: myProvider,
});
export const output = {
    value: myCluster.kubeconfig,
};
