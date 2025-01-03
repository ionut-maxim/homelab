import * as pulumi from "@pulumi/pulumi";
import * as talos from "@pulumiverse/talos";


const CLUSTER_NAME = "ninthfloor"

const networkStack = new pulumi.StackReference("ionut-maxim/network/main")

const ipAddresses = networkStack.getOutput("ip-addresses") as pulumi.Output<Record<string, string>>;

const node03ip = ipAddresses["node03"]
const kubeVIPIp = ipAddresses["kubernetes-vip"]

const secrets = new talos.machine.Secrets("secretsV1", {talosVersion: "v1.9.1"});

const configuration = talos.machine.getConfigurationOutput({
    clusterName: CLUSTER_NAME,
    machineType: "controlplane",
    clusterEndpoint: "https://kubernetes.internal:6443",
    machineSecrets: secrets.machineSecrets,
});


const configurationApply = new talos.machine.ConfigurationApply("configurationApply", {
    clientConfiguration: secrets.clientConfiguration,
    machineConfigurationInput: configuration.machineConfiguration,
    node: node03ip,
    onDestroy: {
        reset: true,
        reboot: true,
        graceful: false
    },
    configPatches: [JSON.stringify({
        machine: {
            install: {
                disk: "/dev/nvme0n1",
            }, network: {
                interfaces: [{
                    interface: "eno1", dhcp: true, vip: {ip: "10.20.30.103"}
                }]
            }
        }, cluster: {
            apiServer: {
                certSANs: ["kubernetes.internal"]
            },
        }
    })],
});

const bootstrap = new talos.machine.Bootstrap("bootstrap", {
    node: node03ip, clientConfiguration: secrets.clientConfiguration,
}, {
    dependsOn: [configurationApply],
});

export const talosConfig = talos.client.getConfigurationOutput({
    clusterName: CLUSTER_NAME, clientConfiguration: secrets.clientConfiguration, nodes: [node03ip], endpoints: [node03ip]
}).talosConfig

export const kubeConfig = talos.cluster.getKubeconfigOutput({
    node: node03ip, endpoint: node03ip, clientConfiguration: bootstrap.clientConfiguration
}).kubeconfigRaw