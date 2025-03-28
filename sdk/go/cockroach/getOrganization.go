// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cockroach

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-cockroach/sdk/go/cockroach/internal"
)

// Information about the organization associated with the user's API key.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-cockroach/sdk/go/cockroach"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cockroach.GetOrganization(ctx, map[string]interface{}{}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetOrganization(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetOrganizationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetOrganizationResult
	err := ctx.Invoke("cockroach:index/getOrganization:getOrganization", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getOrganization.
type GetOrganizationResult struct {
	// Indicates when the organization was created.
	CreatedAt string `pulumi:"createdAt"`
	// Organization ID.
	Id string `pulumi:"id"`
	// A short ID used by CockroachDB Support.
	Label string `pulumi:"label"`
	// Name of the organization.
	Name string `pulumi:"name"`
}

func GetOrganizationOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) GetOrganizationResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (GetOrganizationResultOutput, error) {
		options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
		return ctx.InvokeOutput("cockroach:index/getOrganization:getOrganization", nil, GetOrganizationResultOutput{}, options).(GetOrganizationResultOutput), nil
	}).(GetOrganizationResultOutput)
}

// A collection of values returned by getOrganization.
type GetOrganizationResultOutput struct{ *pulumi.OutputState }

func (GetOrganizationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOrganizationResult)(nil)).Elem()
}

func (o GetOrganizationResultOutput) ToGetOrganizationResultOutput() GetOrganizationResultOutput {
	return o
}

func (o GetOrganizationResultOutput) ToGetOrganizationResultOutputWithContext(ctx context.Context) GetOrganizationResultOutput {
	return o
}

// Indicates when the organization was created.
func (o GetOrganizationResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// Organization ID.
func (o GetOrganizationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationResult) string { return v.Id }).(pulumi.StringOutput)
}

// A short ID used by CockroachDB Support.
func (o GetOrganizationResultOutput) Label() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationResult) string { return v.Label }).(pulumi.StringOutput)
}

// Name of the organization.
func (o GetOrganizationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationResult) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOrganizationResultOutput{})
}
