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

// Configuration to manage external JSON Web Token (JWT) Issuers for authentication to the CockroachDB Cloud API.
//
// ## Import
//
// # JWT Issuer ID can be found by running a GET against the Cockroach Cloud API to
//
// list all existing JWT issuers.
//
// https://www.cockroachlabs.com/docs/api/cloud/v1#get-/api/v1/jwt-issuers
//
// format: <jwt issuer id>
//
// ```sh
// $ pulumi import cockroach:index/jwtIssuer:JwtIssuer my_issuer 1f69fdd2-600a-4cfc-a9ba-16995df0d77d
// ```
type JwtIssuer struct {
	pulumi.CustomResourceState

	// The intended audience for consuming the JWT.
	Audience pulumi.StringOutput `pulumi:"audience"`
	// Used to identify the user from the external Identity Provider. Defaults to "sub".
	Claim pulumi.StringPtrOutput `pulumi:"claim"`
	// A list of mappings to map the external token identity into CockroachDB Cloud.
	IdentityMaps JwtIssuerIdentityMapArrayOutput `pulumi:"identityMaps"`
	// The URL of the server issuing JWTs.
	IssuerUrl pulumi.StringOutput `pulumi:"issuerUrl"`
	// A set of public keys (JWKS) used to verify the JWT.
	Jwks pulumi.StringPtrOutput `pulumi:"jwks"`
}

// NewJwtIssuer registers a new resource with the given unique name, arguments, and options.
func NewJwtIssuer(ctx *pulumi.Context,
	name string, args *JwtIssuerArgs, opts ...pulumi.ResourceOption) (*JwtIssuer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Audience == nil {
		return nil, errors.New("invalid value for required argument 'Audience'")
	}
	if args.IssuerUrl == nil {
		return nil, errors.New("invalid value for required argument 'IssuerUrl'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource JwtIssuer
	err := ctx.RegisterResource("cockroach:index/jwtIssuer:JwtIssuer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJwtIssuer gets an existing JwtIssuer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJwtIssuer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JwtIssuerState, opts ...pulumi.ResourceOption) (*JwtIssuer, error) {
	var resource JwtIssuer
	err := ctx.ReadResource("cockroach:index/jwtIssuer:JwtIssuer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering JwtIssuer resources.
type jwtIssuerState struct {
	// The intended audience for consuming the JWT.
	Audience *string `pulumi:"audience"`
	// Used to identify the user from the external Identity Provider. Defaults to "sub".
	Claim *string `pulumi:"claim"`
	// A list of mappings to map the external token identity into CockroachDB Cloud.
	IdentityMaps []JwtIssuerIdentityMap `pulumi:"identityMaps"`
	// The URL of the server issuing JWTs.
	IssuerUrl *string `pulumi:"issuerUrl"`
	// A set of public keys (JWKS) used to verify the JWT.
	Jwks *string `pulumi:"jwks"`
}

type JwtIssuerState struct {
	// The intended audience for consuming the JWT.
	Audience pulumi.StringPtrInput
	// Used to identify the user from the external Identity Provider. Defaults to "sub".
	Claim pulumi.StringPtrInput
	// A list of mappings to map the external token identity into CockroachDB Cloud.
	IdentityMaps JwtIssuerIdentityMapArrayInput
	// The URL of the server issuing JWTs.
	IssuerUrl pulumi.StringPtrInput
	// A set of public keys (JWKS) used to verify the JWT.
	Jwks pulumi.StringPtrInput
}

func (JwtIssuerState) ElementType() reflect.Type {
	return reflect.TypeOf((*jwtIssuerState)(nil)).Elem()
}

type jwtIssuerArgs struct {
	// The intended audience for consuming the JWT.
	Audience string `pulumi:"audience"`
	// Used to identify the user from the external Identity Provider. Defaults to "sub".
	Claim *string `pulumi:"claim"`
	// A list of mappings to map the external token identity into CockroachDB Cloud.
	IdentityMaps []JwtIssuerIdentityMap `pulumi:"identityMaps"`
	// The URL of the server issuing JWTs.
	IssuerUrl string `pulumi:"issuerUrl"`
	// A set of public keys (JWKS) used to verify the JWT.
	Jwks *string `pulumi:"jwks"`
}

// The set of arguments for constructing a JwtIssuer resource.
type JwtIssuerArgs struct {
	// The intended audience for consuming the JWT.
	Audience pulumi.StringInput
	// Used to identify the user from the external Identity Provider. Defaults to "sub".
	Claim pulumi.StringPtrInput
	// A list of mappings to map the external token identity into CockroachDB Cloud.
	IdentityMaps JwtIssuerIdentityMapArrayInput
	// The URL of the server issuing JWTs.
	IssuerUrl pulumi.StringInput
	// A set of public keys (JWKS) used to verify the JWT.
	Jwks pulumi.StringPtrInput
}

func (JwtIssuerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jwtIssuerArgs)(nil)).Elem()
}

type JwtIssuerInput interface {
	pulumi.Input

	ToJwtIssuerOutput() JwtIssuerOutput
	ToJwtIssuerOutputWithContext(ctx context.Context) JwtIssuerOutput
}

func (*JwtIssuer) ElementType() reflect.Type {
	return reflect.TypeOf((**JwtIssuer)(nil)).Elem()
}

func (i *JwtIssuer) ToJwtIssuerOutput() JwtIssuerOutput {
	return i.ToJwtIssuerOutputWithContext(context.Background())
}

func (i *JwtIssuer) ToJwtIssuerOutputWithContext(ctx context.Context) JwtIssuerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JwtIssuerOutput)
}

// JwtIssuerArrayInput is an input type that accepts JwtIssuerArray and JwtIssuerArrayOutput values.
// You can construct a concrete instance of `JwtIssuerArrayInput` via:
//
//	JwtIssuerArray{ JwtIssuerArgs{...} }
type JwtIssuerArrayInput interface {
	pulumi.Input

	ToJwtIssuerArrayOutput() JwtIssuerArrayOutput
	ToJwtIssuerArrayOutputWithContext(context.Context) JwtIssuerArrayOutput
}

type JwtIssuerArray []JwtIssuerInput

func (JwtIssuerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*JwtIssuer)(nil)).Elem()
}

func (i JwtIssuerArray) ToJwtIssuerArrayOutput() JwtIssuerArrayOutput {
	return i.ToJwtIssuerArrayOutputWithContext(context.Background())
}

func (i JwtIssuerArray) ToJwtIssuerArrayOutputWithContext(ctx context.Context) JwtIssuerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JwtIssuerArrayOutput)
}

// JwtIssuerMapInput is an input type that accepts JwtIssuerMap and JwtIssuerMapOutput values.
// You can construct a concrete instance of `JwtIssuerMapInput` via:
//
//	JwtIssuerMap{ "key": JwtIssuerArgs{...} }
type JwtIssuerMapInput interface {
	pulumi.Input

	ToJwtIssuerMapOutput() JwtIssuerMapOutput
	ToJwtIssuerMapOutputWithContext(context.Context) JwtIssuerMapOutput
}

type JwtIssuerMap map[string]JwtIssuerInput

func (JwtIssuerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*JwtIssuer)(nil)).Elem()
}

func (i JwtIssuerMap) ToJwtIssuerMapOutput() JwtIssuerMapOutput {
	return i.ToJwtIssuerMapOutputWithContext(context.Background())
}

func (i JwtIssuerMap) ToJwtIssuerMapOutputWithContext(ctx context.Context) JwtIssuerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JwtIssuerMapOutput)
}

type JwtIssuerOutput struct{ *pulumi.OutputState }

func (JwtIssuerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JwtIssuer)(nil)).Elem()
}

func (o JwtIssuerOutput) ToJwtIssuerOutput() JwtIssuerOutput {
	return o
}

func (o JwtIssuerOutput) ToJwtIssuerOutputWithContext(ctx context.Context) JwtIssuerOutput {
	return o
}

// The intended audience for consuming the JWT.
func (o JwtIssuerOutput) Audience() pulumi.StringOutput {
	return o.ApplyT(func(v *JwtIssuer) pulumi.StringOutput { return v.Audience }).(pulumi.StringOutput)
}

// Used to identify the user from the external Identity Provider. Defaults to "sub".
func (o JwtIssuerOutput) Claim() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JwtIssuer) pulumi.StringPtrOutput { return v.Claim }).(pulumi.StringPtrOutput)
}

// A list of mappings to map the external token identity into CockroachDB Cloud.
func (o JwtIssuerOutput) IdentityMaps() JwtIssuerIdentityMapArrayOutput {
	return o.ApplyT(func(v *JwtIssuer) JwtIssuerIdentityMapArrayOutput { return v.IdentityMaps }).(JwtIssuerIdentityMapArrayOutput)
}

// The URL of the server issuing JWTs.
func (o JwtIssuerOutput) IssuerUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *JwtIssuer) pulumi.StringOutput { return v.IssuerUrl }).(pulumi.StringOutput)
}

// A set of public keys (JWKS) used to verify the JWT.
func (o JwtIssuerOutput) Jwks() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JwtIssuer) pulumi.StringPtrOutput { return v.Jwks }).(pulumi.StringPtrOutput)
}

type JwtIssuerArrayOutput struct{ *pulumi.OutputState }

func (JwtIssuerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*JwtIssuer)(nil)).Elem()
}

func (o JwtIssuerArrayOutput) ToJwtIssuerArrayOutput() JwtIssuerArrayOutput {
	return o
}

func (o JwtIssuerArrayOutput) ToJwtIssuerArrayOutputWithContext(ctx context.Context) JwtIssuerArrayOutput {
	return o
}

func (o JwtIssuerArrayOutput) Index(i pulumi.IntInput) JwtIssuerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *JwtIssuer {
		return vs[0].([]*JwtIssuer)[vs[1].(int)]
	}).(JwtIssuerOutput)
}

type JwtIssuerMapOutput struct{ *pulumi.OutputState }

func (JwtIssuerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*JwtIssuer)(nil)).Elem()
}

func (o JwtIssuerMapOutput) ToJwtIssuerMapOutput() JwtIssuerMapOutput {
	return o
}

func (o JwtIssuerMapOutput) ToJwtIssuerMapOutputWithContext(ctx context.Context) JwtIssuerMapOutput {
	return o
}

func (o JwtIssuerMapOutput) MapIndex(k pulumi.StringInput) JwtIssuerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *JwtIssuer {
		return vs[0].(map[string]*JwtIssuer)[vs[1].(string)]
	}).(JwtIssuerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*JwtIssuerInput)(nil)).Elem(), &JwtIssuer{})
	pulumi.RegisterInputType(reflect.TypeOf((*JwtIssuerArrayInput)(nil)).Elem(), JwtIssuerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*JwtIssuerMapInput)(nil)).Elem(), JwtIssuerMap{})
	pulumi.RegisterOutputType(JwtIssuerOutput{})
	pulumi.RegisterOutputType(JwtIssuerArrayOutput{})
	pulumi.RegisterOutputType(JwtIssuerMapOutput{})
}
