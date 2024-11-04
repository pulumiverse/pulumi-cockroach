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

// CockroachDB Cloud folder.
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
//			aTeam, err := cockroach.NewFolder(ctx, "a_team", &cockroach.FolderArgs{
//				Name:     pulumi.String("a-team"),
//				ParentId: pulumi.String("root"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = cockroach.NewFolder(ctx, "a_team_dev", &cockroach.FolderArgs{
//				Name:     pulumi.String("dev"),
//				ParentId: aTeam.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type Folder struct {
	pulumi.CustomResourceState

	// Name of the folder.
	Name pulumi.StringOutput `pulumi:"name"`
	// ID of the parent folder. Use 'root' for the root level (no parent folder).
	ParentId pulumi.StringOutput `pulumi:"parentId"`
}

// NewFolder registers a new resource with the given unique name, arguments, and options.
func NewFolder(ctx *pulumi.Context,
	name string, args *FolderArgs, opts ...pulumi.ResourceOption) (*Folder, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ParentId == nil {
		return nil, errors.New("invalid value for required argument 'ParentId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Folder
	err := ctx.RegisterResource("cockroach:index/folder:Folder", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFolder gets an existing Folder resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFolder(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FolderState, opts ...pulumi.ResourceOption) (*Folder, error) {
	var resource Folder
	err := ctx.ReadResource("cockroach:index/folder:Folder", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Folder resources.
type folderState struct {
	// Name of the folder.
	Name *string `pulumi:"name"`
	// ID of the parent folder. Use 'root' for the root level (no parent folder).
	ParentId *string `pulumi:"parentId"`
}

type FolderState struct {
	// Name of the folder.
	Name pulumi.StringPtrInput
	// ID of the parent folder. Use 'root' for the root level (no parent folder).
	ParentId pulumi.StringPtrInput
}

func (FolderState) ElementType() reflect.Type {
	return reflect.TypeOf((*folderState)(nil)).Elem()
}

type folderArgs struct {
	// Name of the folder.
	Name *string `pulumi:"name"`
	// ID of the parent folder. Use 'root' for the root level (no parent folder).
	ParentId string `pulumi:"parentId"`
}

// The set of arguments for constructing a Folder resource.
type FolderArgs struct {
	// Name of the folder.
	Name pulumi.StringPtrInput
	// ID of the parent folder. Use 'root' for the root level (no parent folder).
	ParentId pulumi.StringInput
}

func (FolderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*folderArgs)(nil)).Elem()
}

type FolderInput interface {
	pulumi.Input

	ToFolderOutput() FolderOutput
	ToFolderOutputWithContext(ctx context.Context) FolderOutput
}

func (*Folder) ElementType() reflect.Type {
	return reflect.TypeOf((**Folder)(nil)).Elem()
}

func (i *Folder) ToFolderOutput() FolderOutput {
	return i.ToFolderOutputWithContext(context.Background())
}

func (i *Folder) ToFolderOutputWithContext(ctx context.Context) FolderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FolderOutput)
}

// FolderArrayInput is an input type that accepts FolderArray and FolderArrayOutput values.
// You can construct a concrete instance of `FolderArrayInput` via:
//
//	FolderArray{ FolderArgs{...} }
type FolderArrayInput interface {
	pulumi.Input

	ToFolderArrayOutput() FolderArrayOutput
	ToFolderArrayOutputWithContext(context.Context) FolderArrayOutput
}

type FolderArray []FolderInput

func (FolderArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Folder)(nil)).Elem()
}

func (i FolderArray) ToFolderArrayOutput() FolderArrayOutput {
	return i.ToFolderArrayOutputWithContext(context.Background())
}

func (i FolderArray) ToFolderArrayOutputWithContext(ctx context.Context) FolderArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FolderArrayOutput)
}

// FolderMapInput is an input type that accepts FolderMap and FolderMapOutput values.
// You can construct a concrete instance of `FolderMapInput` via:
//
//	FolderMap{ "key": FolderArgs{...} }
type FolderMapInput interface {
	pulumi.Input

	ToFolderMapOutput() FolderMapOutput
	ToFolderMapOutputWithContext(context.Context) FolderMapOutput
}

type FolderMap map[string]FolderInput

func (FolderMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Folder)(nil)).Elem()
}

func (i FolderMap) ToFolderMapOutput() FolderMapOutput {
	return i.ToFolderMapOutputWithContext(context.Background())
}

func (i FolderMap) ToFolderMapOutputWithContext(ctx context.Context) FolderMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FolderMapOutput)
}

type FolderOutput struct{ *pulumi.OutputState }

func (FolderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Folder)(nil)).Elem()
}

func (o FolderOutput) ToFolderOutput() FolderOutput {
	return o
}

func (o FolderOutput) ToFolderOutputWithContext(ctx context.Context) FolderOutput {
	return o
}

// Name of the folder.
func (o FolderOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Folder) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ID of the parent folder. Use 'root' for the root level (no parent folder).
func (o FolderOutput) ParentId() pulumi.StringOutput {
	return o.ApplyT(func(v *Folder) pulumi.StringOutput { return v.ParentId }).(pulumi.StringOutput)
}

type FolderArrayOutput struct{ *pulumi.OutputState }

func (FolderArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Folder)(nil)).Elem()
}

func (o FolderArrayOutput) ToFolderArrayOutput() FolderArrayOutput {
	return o
}

func (o FolderArrayOutput) ToFolderArrayOutputWithContext(ctx context.Context) FolderArrayOutput {
	return o
}

func (o FolderArrayOutput) Index(i pulumi.IntInput) FolderOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Folder {
		return vs[0].([]*Folder)[vs[1].(int)]
	}).(FolderOutput)
}

type FolderMapOutput struct{ *pulumi.OutputState }

func (FolderMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Folder)(nil)).Elem()
}

func (o FolderMapOutput) ToFolderMapOutput() FolderMapOutput {
	return o
}

func (o FolderMapOutput) ToFolderMapOutputWithContext(ctx context.Context) FolderMapOutput {
	return o
}

func (o FolderMapOutput) MapIndex(k pulumi.StringInput) FolderOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Folder {
		return vs[0].(map[string]*Folder)[vs[1].(string)]
	}).(FolderOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FolderInput)(nil)).Elem(), &Folder{})
	pulumi.RegisterInputType(reflect.TypeOf((*FolderArrayInput)(nil)).Elem(), FolderArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FolderMapInput)(nil)).Elem(), FolderMap{})
	pulumi.RegisterOutputType(FolderOutput{})
	pulumi.RegisterOutputType(FolderArrayOutput{})
	pulumi.RegisterOutputType(FolderMapOutput{})
}
