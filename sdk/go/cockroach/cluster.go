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

// CockroachDB Cloud cluster.
//
// ## Import
//
// format: <cluster id>
//
// ```sh
// $ pulumi import cockroach:index/cluster:Cluster my_cluster 1f69fdd2-600a-4cfc-a9ba-16995df0d77d
// ```
type Cluster struct {
	pulumi.CustomResourceState

	// The cloud provider account ID that hosts the cluster. Needed for CMEK and other advanced features.
	AccountId pulumi.StringOutput `pulumi:"accountId"`
	// The backup settings for a cluster. Each cluster has backup settings that determine if backups are enabled, how
	// frequently they are taken, and how long they are retained for. Use this attribute to manage those settings.
	BackupConfig ClusterBackupConfigOutput `pulumi:"backupConfig"`
	// Cloud provider used to host the cluster. Allowed values are: * GCP * AWS * AZURE
	CloudProvider pulumi.StringOutput `pulumi:"cloudProvider"`
	// Major version of CockroachDB running on the cluster. This value can be used to orchestrate version upgrades. Supported
	// for ADVANCED and STANDARD clusters (when `serverless.upgrade_type` set to 'MANUAL').
	CockroachVersion pulumi.StringOutput `pulumi:"cockroachVersion"`
	// ID of the user who created the cluster.
	CreatorId pulumi.StringOutput       `pulumi:"creatorId"`
	Dedicated ClusterDedicatedPtrOutput `pulumi:"dedicated"`
	// Set to true to enable delete protection on the cluster. If unset, the server chooses the value on cluster creation, and
	// preserves the value on cluster update.
	DeleteProtection pulumi.BoolOutput `pulumi:"deleteProtection"`
	// Name of the cluster.
	Name pulumi.StringOutput `pulumi:"name"`
	// Describes the current long-running operation, if any.
	OperationStatus pulumi.StringOutput `pulumi:"operationStatus"`
	// The ID of the cluster's parent folder. 'root' is used for a cluster at the root level.
	ParentId pulumi.StringOutput `pulumi:"parentId"`
	// Denotes cluster plan type: 'BASIC' or 'STANDARD' or 'ADVANCED'.
	Plan       pulumi.StringOutput        `pulumi:"plan"`
	Regions    ClusterRegionArrayOutput   `pulumi:"regions"`
	Serverless ClusterServerlessPtrOutput `pulumi:"serverless"`
	// Describes whether the cluster is being created, updated, deleted, etc.
	State pulumi.StringOutput `pulumi:"state"`
	// Describes the status of any in-progress CockroachDB upgrade or rollback.
	UpgradeStatus pulumi.StringOutput `pulumi:"upgradeStatus"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CloudProvider == nil {
		return nil, errors.New("invalid value for required argument 'CloudProvider'")
	}
	if args.Regions == nil {
		return nil, errors.New("invalid value for required argument 'Regions'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Cluster
	err := ctx.RegisterResource("cockroach:index/cluster:Cluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCluster gets an existing Cluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterState, opts ...pulumi.ResourceOption) (*Cluster, error) {
	var resource Cluster
	err := ctx.ReadResource("cockroach:index/cluster:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cluster resources.
type clusterState struct {
	// The cloud provider account ID that hosts the cluster. Needed for CMEK and other advanced features.
	AccountId *string `pulumi:"accountId"`
	// The backup settings for a cluster. Each cluster has backup settings that determine if backups are enabled, how
	// frequently they are taken, and how long they are retained for. Use this attribute to manage those settings.
	BackupConfig *ClusterBackupConfig `pulumi:"backupConfig"`
	// Cloud provider used to host the cluster. Allowed values are: * GCP * AWS * AZURE
	CloudProvider *string `pulumi:"cloudProvider"`
	// Major version of CockroachDB running on the cluster. This value can be used to orchestrate version upgrades. Supported
	// for ADVANCED and STANDARD clusters (when `serverless.upgrade_type` set to 'MANUAL').
	CockroachVersion *string `pulumi:"cockroachVersion"`
	// ID of the user who created the cluster.
	CreatorId *string           `pulumi:"creatorId"`
	Dedicated *ClusterDedicated `pulumi:"dedicated"`
	// Set to true to enable delete protection on the cluster. If unset, the server chooses the value on cluster creation, and
	// preserves the value on cluster update.
	DeleteProtection *bool `pulumi:"deleteProtection"`
	// Name of the cluster.
	Name *string `pulumi:"name"`
	// Describes the current long-running operation, if any.
	OperationStatus *string `pulumi:"operationStatus"`
	// The ID of the cluster's parent folder. 'root' is used for a cluster at the root level.
	ParentId *string `pulumi:"parentId"`
	// Denotes cluster plan type: 'BASIC' or 'STANDARD' or 'ADVANCED'.
	Plan       *string            `pulumi:"plan"`
	Regions    []ClusterRegion    `pulumi:"regions"`
	Serverless *ClusterServerless `pulumi:"serverless"`
	// Describes whether the cluster is being created, updated, deleted, etc.
	State *string `pulumi:"state"`
	// Describes the status of any in-progress CockroachDB upgrade or rollback.
	UpgradeStatus *string `pulumi:"upgradeStatus"`
}

type ClusterState struct {
	// The cloud provider account ID that hosts the cluster. Needed for CMEK and other advanced features.
	AccountId pulumi.StringPtrInput
	// The backup settings for a cluster. Each cluster has backup settings that determine if backups are enabled, how
	// frequently they are taken, and how long they are retained for. Use this attribute to manage those settings.
	BackupConfig ClusterBackupConfigPtrInput
	// Cloud provider used to host the cluster. Allowed values are: * GCP * AWS * AZURE
	CloudProvider pulumi.StringPtrInput
	// Major version of CockroachDB running on the cluster. This value can be used to orchestrate version upgrades. Supported
	// for ADVANCED and STANDARD clusters (when `serverless.upgrade_type` set to 'MANUAL').
	CockroachVersion pulumi.StringPtrInput
	// ID of the user who created the cluster.
	CreatorId pulumi.StringPtrInput
	Dedicated ClusterDedicatedPtrInput
	// Set to true to enable delete protection on the cluster. If unset, the server chooses the value on cluster creation, and
	// preserves the value on cluster update.
	DeleteProtection pulumi.BoolPtrInput
	// Name of the cluster.
	Name pulumi.StringPtrInput
	// Describes the current long-running operation, if any.
	OperationStatus pulumi.StringPtrInput
	// The ID of the cluster's parent folder. 'root' is used for a cluster at the root level.
	ParentId pulumi.StringPtrInput
	// Denotes cluster plan type: 'BASIC' or 'STANDARD' or 'ADVANCED'.
	Plan       pulumi.StringPtrInput
	Regions    ClusterRegionArrayInput
	Serverless ClusterServerlessPtrInput
	// Describes whether the cluster is being created, updated, deleted, etc.
	State pulumi.StringPtrInput
	// Describes the status of any in-progress CockroachDB upgrade or rollback.
	UpgradeStatus pulumi.StringPtrInput
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	// The backup settings for a cluster. Each cluster has backup settings that determine if backups are enabled, how
	// frequently they are taken, and how long they are retained for. Use this attribute to manage those settings.
	BackupConfig *ClusterBackupConfig `pulumi:"backupConfig"`
	// Cloud provider used to host the cluster. Allowed values are: * GCP * AWS * AZURE
	CloudProvider string `pulumi:"cloudProvider"`
	// Major version of CockroachDB running on the cluster. This value can be used to orchestrate version upgrades. Supported
	// for ADVANCED and STANDARD clusters (when `serverless.upgrade_type` set to 'MANUAL').
	CockroachVersion *string           `pulumi:"cockroachVersion"`
	Dedicated        *ClusterDedicated `pulumi:"dedicated"`
	// Set to true to enable delete protection on the cluster. If unset, the server chooses the value on cluster creation, and
	// preserves the value on cluster update.
	DeleteProtection *bool `pulumi:"deleteProtection"`
	// Name of the cluster.
	Name *string `pulumi:"name"`
	// The ID of the cluster's parent folder. 'root' is used for a cluster at the root level.
	ParentId *string `pulumi:"parentId"`
	// Denotes cluster plan type: 'BASIC' or 'STANDARD' or 'ADVANCED'.
	Plan       *string            `pulumi:"plan"`
	Regions    []ClusterRegion    `pulumi:"regions"`
	Serverless *ClusterServerless `pulumi:"serverless"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// The backup settings for a cluster. Each cluster has backup settings that determine if backups are enabled, how
	// frequently they are taken, and how long they are retained for. Use this attribute to manage those settings.
	BackupConfig ClusterBackupConfigPtrInput
	// Cloud provider used to host the cluster. Allowed values are: * GCP * AWS * AZURE
	CloudProvider pulumi.StringInput
	// Major version of CockroachDB running on the cluster. This value can be used to orchestrate version upgrades. Supported
	// for ADVANCED and STANDARD clusters (when `serverless.upgrade_type` set to 'MANUAL').
	CockroachVersion pulumi.StringPtrInput
	Dedicated        ClusterDedicatedPtrInput
	// Set to true to enable delete protection on the cluster. If unset, the server chooses the value on cluster creation, and
	// preserves the value on cluster update.
	DeleteProtection pulumi.BoolPtrInput
	// Name of the cluster.
	Name pulumi.StringPtrInput
	// The ID of the cluster's parent folder. 'root' is used for a cluster at the root level.
	ParentId pulumi.StringPtrInput
	// Denotes cluster plan type: 'BASIC' or 'STANDARD' or 'ADVANCED'.
	Plan       pulumi.StringPtrInput
	Regions    ClusterRegionArrayInput
	Serverless ClusterServerlessPtrInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}

type ClusterInput interface {
	pulumi.Input

	ToClusterOutput() ClusterOutput
	ToClusterOutputWithContext(ctx context.Context) ClusterOutput
}

func (*Cluster) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (i *Cluster) ToClusterOutput() ClusterOutput {
	return i.ToClusterOutputWithContext(context.Background())
}

func (i *Cluster) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterOutput)
}

// ClusterArrayInput is an input type that accepts ClusterArray and ClusterArrayOutput values.
// You can construct a concrete instance of `ClusterArrayInput` via:
//
//	ClusterArray{ ClusterArgs{...} }
type ClusterArrayInput interface {
	pulumi.Input

	ToClusterArrayOutput() ClusterArrayOutput
	ToClusterArrayOutputWithContext(context.Context) ClusterArrayOutput
}

type ClusterArray []ClusterInput

func (ClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Cluster)(nil)).Elem()
}

func (i ClusterArray) ToClusterArrayOutput() ClusterArrayOutput {
	return i.ToClusterArrayOutputWithContext(context.Background())
}

func (i ClusterArray) ToClusterArrayOutputWithContext(ctx context.Context) ClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterArrayOutput)
}

// ClusterMapInput is an input type that accepts ClusterMap and ClusterMapOutput values.
// You can construct a concrete instance of `ClusterMapInput` via:
//
//	ClusterMap{ "key": ClusterArgs{...} }
type ClusterMapInput interface {
	pulumi.Input

	ToClusterMapOutput() ClusterMapOutput
	ToClusterMapOutputWithContext(context.Context) ClusterMapOutput
}

type ClusterMap map[string]ClusterInput

func (ClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Cluster)(nil)).Elem()
}

func (i ClusterMap) ToClusterMapOutput() ClusterMapOutput {
	return i.ToClusterMapOutputWithContext(context.Background())
}

func (i ClusterMap) ToClusterMapOutputWithContext(ctx context.Context) ClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterMapOutput)
}

type ClusterOutput struct{ *pulumi.OutputState }

func (ClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (o ClusterOutput) ToClusterOutput() ClusterOutput {
	return o
}

func (o ClusterOutput) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return o
}

// The cloud provider account ID that hosts the cluster. Needed for CMEK and other advanced features.
func (o ClusterOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.AccountId }).(pulumi.StringOutput)
}

// The backup settings for a cluster. Each cluster has backup settings that determine if backups are enabled, how
// frequently they are taken, and how long they are retained for. Use this attribute to manage those settings.
func (o ClusterOutput) BackupConfig() ClusterBackupConfigOutput {
	return o.ApplyT(func(v *Cluster) ClusterBackupConfigOutput { return v.BackupConfig }).(ClusterBackupConfigOutput)
}

// Cloud provider used to host the cluster. Allowed values are: * GCP * AWS * AZURE
func (o ClusterOutput) CloudProvider() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.CloudProvider }).(pulumi.StringOutput)
}

// Major version of CockroachDB running on the cluster. This value can be used to orchestrate version upgrades. Supported
// for ADVANCED and STANDARD clusters (when `serverless.upgrade_type` set to 'MANUAL').
func (o ClusterOutput) CockroachVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.CockroachVersion }).(pulumi.StringOutput)
}

// ID of the user who created the cluster.
func (o ClusterOutput) CreatorId() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.CreatorId }).(pulumi.StringOutput)
}

func (o ClusterOutput) Dedicated() ClusterDedicatedPtrOutput {
	return o.ApplyT(func(v *Cluster) ClusterDedicatedPtrOutput { return v.Dedicated }).(ClusterDedicatedPtrOutput)
}

// Set to true to enable delete protection on the cluster. If unset, the server chooses the value on cluster creation, and
// preserves the value on cluster update.
func (o ClusterOutput) DeleteProtection() pulumi.BoolOutput {
	return o.ApplyT(func(v *Cluster) pulumi.BoolOutput { return v.DeleteProtection }).(pulumi.BoolOutput)
}

// Name of the cluster.
func (o ClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Describes the current long-running operation, if any.
func (o ClusterOutput) OperationStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.OperationStatus }).(pulumi.StringOutput)
}

// The ID of the cluster's parent folder. 'root' is used for a cluster at the root level.
func (o ClusterOutput) ParentId() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ParentId }).(pulumi.StringOutput)
}

// Denotes cluster plan type: 'BASIC' or 'STANDARD' or 'ADVANCED'.
func (o ClusterOutput) Plan() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Plan }).(pulumi.StringOutput)
}

func (o ClusterOutput) Regions() ClusterRegionArrayOutput {
	return o.ApplyT(func(v *Cluster) ClusterRegionArrayOutput { return v.Regions }).(ClusterRegionArrayOutput)
}

func (o ClusterOutput) Serverless() ClusterServerlessPtrOutput {
	return o.ApplyT(func(v *Cluster) ClusterServerlessPtrOutput { return v.Serverless }).(ClusterServerlessPtrOutput)
}

// Describes whether the cluster is being created, updated, deleted, etc.
func (o ClusterOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Describes the status of any in-progress CockroachDB upgrade or rollback.
func (o ClusterOutput) UpgradeStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.UpgradeStatus }).(pulumi.StringOutput)
}

type ClusterArrayOutput struct{ *pulumi.OutputState }

func (ClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Cluster)(nil)).Elem()
}

func (o ClusterArrayOutput) ToClusterArrayOutput() ClusterArrayOutput {
	return o
}

func (o ClusterArrayOutput) ToClusterArrayOutputWithContext(ctx context.Context) ClusterArrayOutput {
	return o
}

func (o ClusterArrayOutput) Index(i pulumi.IntInput) ClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Cluster {
		return vs[0].([]*Cluster)[vs[1].(int)]
	}).(ClusterOutput)
}

type ClusterMapOutput struct{ *pulumi.OutputState }

func (ClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Cluster)(nil)).Elem()
}

func (o ClusterMapOutput) ToClusterMapOutput() ClusterMapOutput {
	return o
}

func (o ClusterMapOutput) ToClusterMapOutputWithContext(ctx context.Context) ClusterMapOutput {
	return o
}

func (o ClusterMapOutput) MapIndex(k pulumi.StringInput) ClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Cluster {
		return vs[0].(map[string]*Cluster)[vs[1].(string)]
	}).(ClusterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterInput)(nil)).Elem(), &Cluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterArrayInput)(nil)).Elem(), ClusterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterMapInput)(nil)).Elem(), ClusterMap{})
	pulumi.RegisterOutputType(ClusterOutput{})
	pulumi.RegisterOutputType(ClusterArrayOutput{})
	pulumi.RegisterOutputType(ClusterMapOutput{})
}
