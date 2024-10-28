// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cockroach

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-cockroach/sdk/go/cockroach/internal"
)

// CockroachDB Cloud cluster. Can be Dedicated or Serverless.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//	"github.com/pulumiverse/pulumi-cockroach/sdk/go/cockroach"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			clusterId := cfg.Require("clusterId")
//			_, err := cockroach.GetCockroachCluster(ctx, &cockroach.GetCockroachClusterArgs{
//				Id: clusterId,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetCockroachCluster(ctx *pulumi.Context, args *GetCockroachClusterArgs, opts ...pulumi.InvokeOption) (*GetCockroachClusterResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetCockroachClusterResult
	err := ctx.Invoke("cockroach:index/getCockroachCluster:getCockroachCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCockroachCluster.
type GetCockroachClusterArgs struct {
	Id string `pulumi:"id"`
}

// A collection of values returned by getCockroachCluster.
type GetCockroachClusterResult struct {
	AccountId        string                        `pulumi:"accountId"`
	CloudProvider    string                        `pulumi:"cloudProvider"`
	CockroachVersion string                        `pulumi:"cockroachVersion"`
	CreatorId        string                        `pulumi:"creatorId"`
	Dedicated        GetCockroachClusterDedicated  `pulumi:"dedicated"`
	Id               string                        `pulumi:"id"`
	Name             string                        `pulumi:"name"`
	OperationStatus  string                        `pulumi:"operationStatus"`
	ParentId         string                        `pulumi:"parentId"`
	Plan             string                        `pulumi:"plan"`
	Regions          []GetCockroachClusterRegion   `pulumi:"regions"`
	Serverless       GetCockroachClusterServerless `pulumi:"serverless"`
	State            string                        `pulumi:"state"`
	UpgradeStatus    string                        `pulumi:"upgradeStatus"`
}

func GetCockroachClusterOutput(ctx *pulumi.Context, args GetCockroachClusterOutputArgs, opts ...pulumi.InvokeOption) GetCockroachClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCockroachClusterResultOutput, error) {
			args := v.(GetCockroachClusterArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv GetCockroachClusterResult
			secret, err := ctx.InvokePackageRaw("cockroach:index/getCockroachCluster:getCockroachCluster", args, &rv, "", opts...)
			if err != nil {
				return GetCockroachClusterResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(GetCockroachClusterResultOutput)
			if secret {
				return pulumi.ToSecret(output).(GetCockroachClusterResultOutput), nil
			}
			return output, nil
		}).(GetCockroachClusterResultOutput)
}

// A collection of arguments for invoking getCockroachCluster.
type GetCockroachClusterOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (GetCockroachClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCockroachClusterArgs)(nil)).Elem()
}

// A collection of values returned by getCockroachCluster.
type GetCockroachClusterResultOutput struct{ *pulumi.OutputState }

func (GetCockroachClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCockroachClusterResult)(nil)).Elem()
}

func (o GetCockroachClusterResultOutput) ToGetCockroachClusterResultOutput() GetCockroachClusterResultOutput {
	return o
}

func (o GetCockroachClusterResultOutput) ToGetCockroachClusterResultOutputWithContext(ctx context.Context) GetCockroachClusterResultOutput {
	return o
}

func (o GetCockroachClusterResultOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v GetCockroachClusterResult) string { return v.AccountId }).(pulumi.StringOutput)
}

func (o GetCockroachClusterResultOutput) CloudProvider() pulumi.StringOutput {
	return o.ApplyT(func(v GetCockroachClusterResult) string { return v.CloudProvider }).(pulumi.StringOutput)
}

func (o GetCockroachClusterResultOutput) CockroachVersion() pulumi.StringOutput {
	return o.ApplyT(func(v GetCockroachClusterResult) string { return v.CockroachVersion }).(pulumi.StringOutput)
}

func (o GetCockroachClusterResultOutput) CreatorId() pulumi.StringOutput {
	return o.ApplyT(func(v GetCockroachClusterResult) string { return v.CreatorId }).(pulumi.StringOutput)
}

func (o GetCockroachClusterResultOutput) Dedicated() GetCockroachClusterDedicatedOutput {
	return o.ApplyT(func(v GetCockroachClusterResult) GetCockroachClusterDedicated { return v.Dedicated }).(GetCockroachClusterDedicatedOutput)
}

func (o GetCockroachClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCockroachClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetCockroachClusterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetCockroachClusterResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetCockroachClusterResultOutput) OperationStatus() pulumi.StringOutput {
	return o.ApplyT(func(v GetCockroachClusterResult) string { return v.OperationStatus }).(pulumi.StringOutput)
}

func (o GetCockroachClusterResultOutput) ParentId() pulumi.StringOutput {
	return o.ApplyT(func(v GetCockroachClusterResult) string { return v.ParentId }).(pulumi.StringOutput)
}

func (o GetCockroachClusterResultOutput) Plan() pulumi.StringOutput {
	return o.ApplyT(func(v GetCockroachClusterResult) string { return v.Plan }).(pulumi.StringOutput)
}

func (o GetCockroachClusterResultOutput) Regions() GetCockroachClusterRegionArrayOutput {
	return o.ApplyT(func(v GetCockroachClusterResult) []GetCockroachClusterRegion { return v.Regions }).(GetCockroachClusterRegionArrayOutput)
}

func (o GetCockroachClusterResultOutput) Serverless() GetCockroachClusterServerlessOutput {
	return o.ApplyT(func(v GetCockroachClusterResult) GetCockroachClusterServerless { return v.Serverless }).(GetCockroachClusterServerlessOutput)
}

func (o GetCockroachClusterResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v GetCockroachClusterResult) string { return v.State }).(pulumi.StringOutput)
}

func (o GetCockroachClusterResultOutput) UpgradeStatus() pulumi.StringOutput {
	return o.ApplyT(func(v GetCockroachClusterResult) string { return v.UpgradeStatus }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCockroachClusterResultOutput{})
}
