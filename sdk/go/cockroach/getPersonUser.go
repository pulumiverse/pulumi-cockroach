// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cockroach

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-cockroach/sdk/go/cockroach/internal"
)

// Information about an individual user.
func GetPersonUser(ctx *pulumi.Context, args *GetPersonUserArgs, opts ...pulumi.InvokeOption) (*GetPersonUserResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetPersonUserResult
	err := ctx.Invoke("cockroach:index/getPersonUser:getPersonUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPersonUser.
type GetPersonUserArgs struct {
	// Email address used to find the User ID.
	Email string `pulumi:"email"`
}

// A collection of values returned by getPersonUser.
type GetPersonUserResult struct {
	// Email address used to find the User ID.
	Email string `pulumi:"email"`
	// User ID.
	Id string `pulumi:"id"`
}

func GetPersonUserOutput(ctx *pulumi.Context, args GetPersonUserOutputArgs, opts ...pulumi.InvokeOption) GetPersonUserResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPersonUserResult, error) {
			args := v.(GetPersonUserArgs)
			r, err := GetPersonUser(ctx, &args, opts...)
			var s GetPersonUserResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetPersonUserResultOutput)
}

// A collection of arguments for invoking getPersonUser.
type GetPersonUserOutputArgs struct {
	// Email address used to find the User ID.
	Email pulumi.StringInput `pulumi:"email"`
}

func (GetPersonUserOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPersonUserArgs)(nil)).Elem()
}

// A collection of values returned by getPersonUser.
type GetPersonUserResultOutput struct{ *pulumi.OutputState }

func (GetPersonUserResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPersonUserResult)(nil)).Elem()
}

func (o GetPersonUserResultOutput) ToGetPersonUserResultOutput() GetPersonUserResultOutput {
	return o
}

func (o GetPersonUserResultOutput) ToGetPersonUserResultOutputWithContext(ctx context.Context) GetPersonUserResultOutput {
	return o
}

// Email address used to find the User ID.
func (o GetPersonUserResultOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v GetPersonUserResult) string { return v.Email }).(pulumi.StringOutput)
}

// User ID.
func (o GetPersonUserResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetPersonUserResult) string { return v.Id }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPersonUserResultOutput{})
}
