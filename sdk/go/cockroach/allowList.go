// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cockroach

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-cockroach/sdk/go/cockroach/internal"
)

// List of IP ranges allowed to access the cluster.
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
//			_, err := cockroach.NewAllowList(ctx, "vpn", &cockroach.AllowListArgs{
//				Name:      pulumi.String("vpn"),
//				CidrIp:    pulumi.String("123.123.1.1"),
//				CidrMask:  pulumi.Int(32),
//				Ui:        pulumi.Bool(true),
//				Sql:       pulumi.Bool(true),
//				ClusterId: pulumi.Any(staging.Id),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// format: <cluster id>:<cidr ip>/<cidr mask>
//
// ```sh
// $ pulumi import cockroach:index/allowList:AllowList home_office 1f69fdd2-600a-4cfc-a9ba-16995df0d77d:123.123.1.1/32
// ```
type AllowList struct {
	pulumi.CustomResourceState

	// IP address component of the [CIDR](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing#CIDR_notation) range for this entry.
	CidrIp pulumi.StringOutput `pulumi:"cidrIp"`
	// The [CIDR](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing#CIDR_notation) notation prefix length. A number ranging from 0 to 32 indicating the size of the network. Use 32 to allow a single IP address.
	CidrMask  pulumi.IntOutput    `pulumi:"cidrMask"`
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Name of this allowlist entry. If not set explicitly, this value does not sync with the server.
	Name pulumi.StringOutput `pulumi:"name"`
	// Set to 'true' to allow SQL connections from this CIDR range.
	Sql pulumi.BoolOutput `pulumi:"sql"`
	// Set to 'true' to allow access to the management console from this CIDR range.
	Ui pulumi.BoolOutput `pulumi:"ui"`
}

// NewAllowList registers a new resource with the given unique name, arguments, and options.
func NewAllowList(ctx *pulumi.Context,
	name string, args *AllowListArgs, opts ...pulumi.ResourceOption) (*AllowList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CidrIp == nil {
		return nil, errors.New("invalid value for required argument 'CidrIp'")
	}
	if args.CidrMask == nil {
		return nil, errors.New("invalid value for required argument 'CidrMask'")
	}
	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.Sql == nil {
		return nil, errors.New("invalid value for required argument 'Sql'")
	}
	if args.Ui == nil {
		return nil, errors.New("invalid value for required argument 'Ui'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AllowList
	err := ctx.RegisterResource("cockroach:index/allowList:AllowList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAllowList gets an existing AllowList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAllowList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AllowListState, opts ...pulumi.ResourceOption) (*AllowList, error) {
	var resource AllowList
	err := ctx.ReadResource("cockroach:index/allowList:AllowList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AllowList resources.
type allowListState struct {
	// IP address component of the [CIDR](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing#CIDR_notation) range for this entry.
	CidrIp *string `pulumi:"cidrIp"`
	// The [CIDR](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing#CIDR_notation) notation prefix length. A number ranging from 0 to 32 indicating the size of the network. Use 32 to allow a single IP address.
	CidrMask  *int    `pulumi:"cidrMask"`
	ClusterId *string `pulumi:"clusterId"`
	// Name of this allowlist entry. If not set explicitly, this value does not sync with the server.
	Name *string `pulumi:"name"`
	// Set to 'true' to allow SQL connections from this CIDR range.
	Sql *bool `pulumi:"sql"`
	// Set to 'true' to allow access to the management console from this CIDR range.
	Ui *bool `pulumi:"ui"`
}

type AllowListState struct {
	// IP address component of the [CIDR](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing#CIDR_notation) range for this entry.
	CidrIp pulumi.StringPtrInput
	// The [CIDR](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing#CIDR_notation) notation prefix length. A number ranging from 0 to 32 indicating the size of the network. Use 32 to allow a single IP address.
	CidrMask  pulumi.IntPtrInput
	ClusterId pulumi.StringPtrInput
	// Name of this allowlist entry. If not set explicitly, this value does not sync with the server.
	Name pulumi.StringPtrInput
	// Set to 'true' to allow SQL connections from this CIDR range.
	Sql pulumi.BoolPtrInput
	// Set to 'true' to allow access to the management console from this CIDR range.
	Ui pulumi.BoolPtrInput
}

func (AllowListState) ElementType() reflect.Type {
	return reflect.TypeOf((*allowListState)(nil)).Elem()
}

type allowListArgs struct {
	// IP address component of the [CIDR](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing#CIDR_notation) range for this entry.
	CidrIp string `pulumi:"cidrIp"`
	// The [CIDR](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing#CIDR_notation) notation prefix length. A number ranging from 0 to 32 indicating the size of the network. Use 32 to allow a single IP address.
	CidrMask  int    `pulumi:"cidrMask"`
	ClusterId string `pulumi:"clusterId"`
	// Name of this allowlist entry. If not set explicitly, this value does not sync with the server.
	Name *string `pulumi:"name"`
	// Set to 'true' to allow SQL connections from this CIDR range.
	Sql bool `pulumi:"sql"`
	// Set to 'true' to allow access to the management console from this CIDR range.
	Ui bool `pulumi:"ui"`
}

// The set of arguments for constructing a AllowList resource.
type AllowListArgs struct {
	// IP address component of the [CIDR](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing#CIDR_notation) range for this entry.
	CidrIp pulumi.StringInput
	// The [CIDR](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing#CIDR_notation) notation prefix length. A number ranging from 0 to 32 indicating the size of the network. Use 32 to allow a single IP address.
	CidrMask  pulumi.IntInput
	ClusterId pulumi.StringInput
	// Name of this allowlist entry. If not set explicitly, this value does not sync with the server.
	Name pulumi.StringPtrInput
	// Set to 'true' to allow SQL connections from this CIDR range.
	Sql pulumi.BoolInput
	// Set to 'true' to allow access to the management console from this CIDR range.
	Ui pulumi.BoolInput
}

func (AllowListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*allowListArgs)(nil)).Elem()
}

type AllowListInput interface {
	pulumi.Input

	ToAllowListOutput() AllowListOutput
	ToAllowListOutputWithContext(ctx context.Context) AllowListOutput
}

func (*AllowList) ElementType() reflect.Type {
	return reflect.TypeOf((**AllowList)(nil)).Elem()
}

func (i *AllowList) ToAllowListOutput() AllowListOutput {
	return i.ToAllowListOutputWithContext(context.Background())
}

func (i *AllowList) ToAllowListOutputWithContext(ctx context.Context) AllowListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AllowListOutput)
}

// AllowListArrayInput is an input type that accepts AllowListArray and AllowListArrayOutput values.
// You can construct a concrete instance of `AllowListArrayInput` via:
//
//	AllowListArray{ AllowListArgs{...} }
type AllowListArrayInput interface {
	pulumi.Input

	ToAllowListArrayOutput() AllowListArrayOutput
	ToAllowListArrayOutputWithContext(context.Context) AllowListArrayOutput
}

type AllowListArray []AllowListInput

func (AllowListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AllowList)(nil)).Elem()
}

func (i AllowListArray) ToAllowListArrayOutput() AllowListArrayOutput {
	return i.ToAllowListArrayOutputWithContext(context.Background())
}

func (i AllowListArray) ToAllowListArrayOutputWithContext(ctx context.Context) AllowListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AllowListArrayOutput)
}

// AllowListMapInput is an input type that accepts AllowListMap and AllowListMapOutput values.
// You can construct a concrete instance of `AllowListMapInput` via:
//
//	AllowListMap{ "key": AllowListArgs{...} }
type AllowListMapInput interface {
	pulumi.Input

	ToAllowListMapOutput() AllowListMapOutput
	ToAllowListMapOutputWithContext(context.Context) AllowListMapOutput
}

type AllowListMap map[string]AllowListInput

func (AllowListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AllowList)(nil)).Elem()
}

func (i AllowListMap) ToAllowListMapOutput() AllowListMapOutput {
	return i.ToAllowListMapOutputWithContext(context.Background())
}

func (i AllowListMap) ToAllowListMapOutputWithContext(ctx context.Context) AllowListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AllowListMapOutput)
}

type AllowListOutput struct{ *pulumi.OutputState }

func (AllowListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AllowList)(nil)).Elem()
}

func (o AllowListOutput) ToAllowListOutput() AllowListOutput {
	return o
}

func (o AllowListOutput) ToAllowListOutputWithContext(ctx context.Context) AllowListOutput {
	return o
}

// IP address component of the [CIDR](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing#CIDR_notation) range for this entry.
func (o AllowListOutput) CidrIp() pulumi.StringOutput {
	return o.ApplyT(func(v *AllowList) pulumi.StringOutput { return v.CidrIp }).(pulumi.StringOutput)
}

// The [CIDR](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing#CIDR_notation) notation prefix length. A number ranging from 0 to 32 indicating the size of the network. Use 32 to allow a single IP address.
func (o AllowListOutput) CidrMask() pulumi.IntOutput {
	return o.ApplyT(func(v *AllowList) pulumi.IntOutput { return v.CidrMask }).(pulumi.IntOutput)
}

func (o AllowListOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *AllowList) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// Name of this allowlist entry. If not set explicitly, this value does not sync with the server.
func (o AllowListOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AllowList) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Set to 'true' to allow SQL connections from this CIDR range.
func (o AllowListOutput) Sql() pulumi.BoolOutput {
	return o.ApplyT(func(v *AllowList) pulumi.BoolOutput { return v.Sql }).(pulumi.BoolOutput)
}

// Set to 'true' to allow access to the management console from this CIDR range.
func (o AllowListOutput) Ui() pulumi.BoolOutput {
	return o.ApplyT(func(v *AllowList) pulumi.BoolOutput { return v.Ui }).(pulumi.BoolOutput)
}

type AllowListArrayOutput struct{ *pulumi.OutputState }

func (AllowListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AllowList)(nil)).Elem()
}

func (o AllowListArrayOutput) ToAllowListArrayOutput() AllowListArrayOutput {
	return o
}

func (o AllowListArrayOutput) ToAllowListArrayOutputWithContext(ctx context.Context) AllowListArrayOutput {
	return o
}

func (o AllowListArrayOutput) Index(i pulumi.IntInput) AllowListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AllowList {
		return vs[0].([]*AllowList)[vs[1].(int)]
	}).(AllowListOutput)
}

type AllowListMapOutput struct{ *pulumi.OutputState }

func (AllowListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AllowList)(nil)).Elem()
}

func (o AllowListMapOutput) ToAllowListMapOutput() AllowListMapOutput {
	return o
}

func (o AllowListMapOutput) ToAllowListMapOutputWithContext(ctx context.Context) AllowListMapOutput {
	return o
}

func (o AllowListMapOutput) MapIndex(k pulumi.StringInput) AllowListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AllowList {
		return vs[0].(map[string]*AllowList)[vs[1].(string)]
	}).(AllowListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AllowListInput)(nil)).Elem(), &AllowList{})
	pulumi.RegisterInputType(reflect.TypeOf((*AllowListArrayInput)(nil)).Elem(), AllowListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AllowListMapInput)(nil)).Elem(), AllowListMap{})
	pulumi.RegisterOutputType(AllowListOutput{})
	pulumi.RegisterOutputType(AllowListArrayOutput{})
	pulumi.RegisterOutputType(AllowListMapOutput{})
}
