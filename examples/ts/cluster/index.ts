import * as pulumi from "@pulumi/pulumi";
import * as cockroach from "@lbrlabs/pulumi-cockroach";

const cluster = new cockroach.Cluster("cluster", {
  cloudProvider: "aws",
  name: "cockroach-provider-ts",
  regions: [
    {
      name: "us-west-2",
    },
  ],
});
