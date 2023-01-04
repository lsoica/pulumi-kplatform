package main

import (
	"github.com/pulumi/pulumi-rke/sdk/v3/go/rke"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := rke.NewCluster(ctx, "my-cluster", &rke.ClusterArgs{
			Nodes: rke.ClusterNodeArray{
				rke.ClusterNodeArgs{
					Address:    pulumi.String("127.0.0.1"),
					Port:       pulumi.String("2222"),
					User:       pulumi.String("vagrant"),
					Roles:      pulumi.StringArray{pulumi.String("controlplane"), pulumi.String("etcd"), pulumi.String("worker")},
					SshKeyPath: pulumi.String("/home/moebius/work/pulumi/.vagrant/machines/default/virtualbox/private_key"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
