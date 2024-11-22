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

// Log Export configuration for a cluster.
//
// ## Import
//
// format: <cluster id>
//
// ```sh
// $ pulumi import cockroach:index/logExportConfig:LogExportConfig example 1f69fdd2-600a-4cfc-a9ba-16995df0d77d
// ```
type LogExportConfig struct {
	pulumi.CustomResourceState

	// Either the AWS Role ARN that identifies a role that the cluster account can assume to write to CloudWatch or the GCP
	// Project ID that the cluster service account has permissions to write to for cloud logging.
	AuthPrincipal pulumi.StringOutput `pulumi:"authPrincipal"`
	// Cluster ID.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Indicates when log export was initially configured.
	CreatedAt pulumi.StringOutput             `pulumi:"createdAt"`
	Groups    LogExportConfigGroupArrayOutput `pulumi:"groups"`
	// An identifier for the logs in the customer's log sink.
	LogName pulumi.StringOutput `pulumi:"logName"`
	// Controls what CRDB channels do not get exported.
	OmittedChannels pulumi.StringArrayOutput `pulumi:"omittedChannels"`
	// Controls whether logs are redacted before forwarding to customer sinks.
	Redact pulumi.BoolPtrOutput `pulumi:"redact"`
	// Controls whether all logs are sent to a specific region in the customer sink.
	Region pulumi.StringOutput `pulumi:"region"`
	// Encodes the possible states that a log export configuration can be in as it is created, deployed, and disabled.
	Status pulumi.StringOutput `pulumi:"status"`
	// The cloud selection being exported to along with the cloud logging platform. Possible values are: * AWS_CLOUDWATCH *
	// GCP_CLOUD_LOGGING * AZURE_LOG_ANALYTICS
	Type pulumi.StringOutput `pulumi:"type"`
	// Indicates when the log export configuration was last updated.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// Elaborates on the log export status and hints at how to fix issues that may have occurred during asynchronous
	// operations.
	UserMessage pulumi.StringOutput `pulumi:"userMessage"`
}

// NewLogExportConfig registers a new resource with the given unique name, arguments, and options.
func NewLogExportConfig(ctx *pulumi.Context,
	name string, args *LogExportConfigArgs, opts ...pulumi.ResourceOption) (*LogExportConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuthPrincipal == nil {
		return nil, errors.New("invalid value for required argument 'AuthPrincipal'")
	}
	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.LogName == nil {
		return nil, errors.New("invalid value for required argument 'LogName'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LogExportConfig
	err := ctx.RegisterResource("cockroach:index/logExportConfig:LogExportConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogExportConfig gets an existing LogExportConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogExportConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogExportConfigState, opts ...pulumi.ResourceOption) (*LogExportConfig, error) {
	var resource LogExportConfig
	err := ctx.ReadResource("cockroach:index/logExportConfig:LogExportConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogExportConfig resources.
type logExportConfigState struct {
	// Either the AWS Role ARN that identifies a role that the cluster account can assume to write to CloudWatch or the GCP
	// Project ID that the cluster service account has permissions to write to for cloud logging.
	AuthPrincipal *string `pulumi:"authPrincipal"`
	// Cluster ID.
	ClusterId *string `pulumi:"clusterId"`
	// Indicates when log export was initially configured.
	CreatedAt *string                `pulumi:"createdAt"`
	Groups    []LogExportConfigGroup `pulumi:"groups"`
	// An identifier for the logs in the customer's log sink.
	LogName *string `pulumi:"logName"`
	// Controls what CRDB channels do not get exported.
	OmittedChannels []string `pulumi:"omittedChannels"`
	// Controls whether logs are redacted before forwarding to customer sinks.
	Redact *bool `pulumi:"redact"`
	// Controls whether all logs are sent to a specific region in the customer sink.
	Region *string `pulumi:"region"`
	// Encodes the possible states that a log export configuration can be in as it is created, deployed, and disabled.
	Status *string `pulumi:"status"`
	// The cloud selection being exported to along with the cloud logging platform. Possible values are: * AWS_CLOUDWATCH *
	// GCP_CLOUD_LOGGING * AZURE_LOG_ANALYTICS
	Type *string `pulumi:"type"`
	// Indicates when the log export configuration was last updated.
	UpdatedAt *string `pulumi:"updatedAt"`
	// Elaborates on the log export status and hints at how to fix issues that may have occurred during asynchronous
	// operations.
	UserMessage *string `pulumi:"userMessage"`
}

type LogExportConfigState struct {
	// Either the AWS Role ARN that identifies a role that the cluster account can assume to write to CloudWatch or the GCP
	// Project ID that the cluster service account has permissions to write to for cloud logging.
	AuthPrincipal pulumi.StringPtrInput
	// Cluster ID.
	ClusterId pulumi.StringPtrInput
	// Indicates when log export was initially configured.
	CreatedAt pulumi.StringPtrInput
	Groups    LogExportConfigGroupArrayInput
	// An identifier for the logs in the customer's log sink.
	LogName pulumi.StringPtrInput
	// Controls what CRDB channels do not get exported.
	OmittedChannels pulumi.StringArrayInput
	// Controls whether logs are redacted before forwarding to customer sinks.
	Redact pulumi.BoolPtrInput
	// Controls whether all logs are sent to a specific region in the customer sink.
	Region pulumi.StringPtrInput
	// Encodes the possible states that a log export configuration can be in as it is created, deployed, and disabled.
	Status pulumi.StringPtrInput
	// The cloud selection being exported to along with the cloud logging platform. Possible values are: * AWS_CLOUDWATCH *
	// GCP_CLOUD_LOGGING * AZURE_LOG_ANALYTICS
	Type pulumi.StringPtrInput
	// Indicates when the log export configuration was last updated.
	UpdatedAt pulumi.StringPtrInput
	// Elaborates on the log export status and hints at how to fix issues that may have occurred during asynchronous
	// operations.
	UserMessage pulumi.StringPtrInput
}

func (LogExportConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*logExportConfigState)(nil)).Elem()
}

type logExportConfigArgs struct {
	// Either the AWS Role ARN that identifies a role that the cluster account can assume to write to CloudWatch or the GCP
	// Project ID that the cluster service account has permissions to write to for cloud logging.
	AuthPrincipal string `pulumi:"authPrincipal"`
	// Cluster ID.
	ClusterId string                 `pulumi:"clusterId"`
	Groups    []LogExportConfigGroup `pulumi:"groups"`
	// An identifier for the logs in the customer's log sink.
	LogName string `pulumi:"logName"`
	// Controls what CRDB channels do not get exported.
	OmittedChannels []string `pulumi:"omittedChannels"`
	// Controls whether logs are redacted before forwarding to customer sinks.
	Redact *bool `pulumi:"redact"`
	// Controls whether all logs are sent to a specific region in the customer sink.
	Region *string `pulumi:"region"`
	// The cloud selection being exported to along with the cloud logging platform. Possible values are: * AWS_CLOUDWATCH *
	// GCP_CLOUD_LOGGING * AZURE_LOG_ANALYTICS
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a LogExportConfig resource.
type LogExportConfigArgs struct {
	// Either the AWS Role ARN that identifies a role that the cluster account can assume to write to CloudWatch or the GCP
	// Project ID that the cluster service account has permissions to write to for cloud logging.
	AuthPrincipal pulumi.StringInput
	// Cluster ID.
	ClusterId pulumi.StringInput
	Groups    LogExportConfigGroupArrayInput
	// An identifier for the logs in the customer's log sink.
	LogName pulumi.StringInput
	// Controls what CRDB channels do not get exported.
	OmittedChannels pulumi.StringArrayInput
	// Controls whether logs are redacted before forwarding to customer sinks.
	Redact pulumi.BoolPtrInput
	// Controls whether all logs are sent to a specific region in the customer sink.
	Region pulumi.StringPtrInput
	// The cloud selection being exported to along with the cloud logging platform. Possible values are: * AWS_CLOUDWATCH *
	// GCP_CLOUD_LOGGING * AZURE_LOG_ANALYTICS
	Type pulumi.StringInput
}

func (LogExportConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logExportConfigArgs)(nil)).Elem()
}

type LogExportConfigInput interface {
	pulumi.Input

	ToLogExportConfigOutput() LogExportConfigOutput
	ToLogExportConfigOutputWithContext(ctx context.Context) LogExportConfigOutput
}

func (*LogExportConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**LogExportConfig)(nil)).Elem()
}

func (i *LogExportConfig) ToLogExportConfigOutput() LogExportConfigOutput {
	return i.ToLogExportConfigOutputWithContext(context.Background())
}

func (i *LogExportConfig) ToLogExportConfigOutputWithContext(ctx context.Context) LogExportConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogExportConfigOutput)
}

// LogExportConfigArrayInput is an input type that accepts LogExportConfigArray and LogExportConfigArrayOutput values.
// You can construct a concrete instance of `LogExportConfigArrayInput` via:
//
//	LogExportConfigArray{ LogExportConfigArgs{...} }
type LogExportConfigArrayInput interface {
	pulumi.Input

	ToLogExportConfigArrayOutput() LogExportConfigArrayOutput
	ToLogExportConfigArrayOutputWithContext(context.Context) LogExportConfigArrayOutput
}

type LogExportConfigArray []LogExportConfigInput

func (LogExportConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogExportConfig)(nil)).Elem()
}

func (i LogExportConfigArray) ToLogExportConfigArrayOutput() LogExportConfigArrayOutput {
	return i.ToLogExportConfigArrayOutputWithContext(context.Background())
}

func (i LogExportConfigArray) ToLogExportConfigArrayOutputWithContext(ctx context.Context) LogExportConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogExportConfigArrayOutput)
}

// LogExportConfigMapInput is an input type that accepts LogExportConfigMap and LogExportConfigMapOutput values.
// You can construct a concrete instance of `LogExportConfigMapInput` via:
//
//	LogExportConfigMap{ "key": LogExportConfigArgs{...} }
type LogExportConfigMapInput interface {
	pulumi.Input

	ToLogExportConfigMapOutput() LogExportConfigMapOutput
	ToLogExportConfigMapOutputWithContext(context.Context) LogExportConfigMapOutput
}

type LogExportConfigMap map[string]LogExportConfigInput

func (LogExportConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogExportConfig)(nil)).Elem()
}

func (i LogExportConfigMap) ToLogExportConfigMapOutput() LogExportConfigMapOutput {
	return i.ToLogExportConfigMapOutputWithContext(context.Background())
}

func (i LogExportConfigMap) ToLogExportConfigMapOutputWithContext(ctx context.Context) LogExportConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogExportConfigMapOutput)
}

type LogExportConfigOutput struct{ *pulumi.OutputState }

func (LogExportConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogExportConfig)(nil)).Elem()
}

func (o LogExportConfigOutput) ToLogExportConfigOutput() LogExportConfigOutput {
	return o
}

func (o LogExportConfigOutput) ToLogExportConfigOutputWithContext(ctx context.Context) LogExportConfigOutput {
	return o
}

// Either the AWS Role ARN that identifies a role that the cluster account can assume to write to CloudWatch or the GCP
// Project ID that the cluster service account has permissions to write to for cloud logging.
func (o LogExportConfigOutput) AuthPrincipal() pulumi.StringOutput {
	return o.ApplyT(func(v *LogExportConfig) pulumi.StringOutput { return v.AuthPrincipal }).(pulumi.StringOutput)
}

// Cluster ID.
func (o LogExportConfigOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *LogExportConfig) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// Indicates when log export was initially configured.
func (o LogExportConfigOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *LogExportConfig) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LogExportConfigOutput) Groups() LogExportConfigGroupArrayOutput {
	return o.ApplyT(func(v *LogExportConfig) LogExportConfigGroupArrayOutput { return v.Groups }).(LogExportConfigGroupArrayOutput)
}

// An identifier for the logs in the customer's log sink.
func (o LogExportConfigOutput) LogName() pulumi.StringOutput {
	return o.ApplyT(func(v *LogExportConfig) pulumi.StringOutput { return v.LogName }).(pulumi.StringOutput)
}

// Controls what CRDB channels do not get exported.
func (o LogExportConfigOutput) OmittedChannels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LogExportConfig) pulumi.StringArrayOutput { return v.OmittedChannels }).(pulumi.StringArrayOutput)
}

// Controls whether logs are redacted before forwarding to customer sinks.
func (o LogExportConfigOutput) Redact() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LogExportConfig) pulumi.BoolPtrOutput { return v.Redact }).(pulumi.BoolPtrOutput)
}

// Controls whether all logs are sent to a specific region in the customer sink.
func (o LogExportConfigOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *LogExportConfig) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Encodes the possible states that a log export configuration can be in as it is created, deployed, and disabled.
func (o LogExportConfigOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *LogExportConfig) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The cloud selection being exported to along with the cloud logging platform. Possible values are: * AWS_CLOUDWATCH *
// GCP_CLOUD_LOGGING * AZURE_LOG_ANALYTICS
func (o LogExportConfigOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *LogExportConfig) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Indicates when the log export configuration was last updated.
func (o LogExportConfigOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *LogExportConfig) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// Elaborates on the log export status and hints at how to fix issues that may have occurred during asynchronous
// operations.
func (o LogExportConfigOutput) UserMessage() pulumi.StringOutput {
	return o.ApplyT(func(v *LogExportConfig) pulumi.StringOutput { return v.UserMessage }).(pulumi.StringOutput)
}

type LogExportConfigArrayOutput struct{ *pulumi.OutputState }

func (LogExportConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogExportConfig)(nil)).Elem()
}

func (o LogExportConfigArrayOutput) ToLogExportConfigArrayOutput() LogExportConfigArrayOutput {
	return o
}

func (o LogExportConfigArrayOutput) ToLogExportConfigArrayOutputWithContext(ctx context.Context) LogExportConfigArrayOutput {
	return o
}

func (o LogExportConfigArrayOutput) Index(i pulumi.IntInput) LogExportConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LogExportConfig {
		return vs[0].([]*LogExportConfig)[vs[1].(int)]
	}).(LogExportConfigOutput)
}

type LogExportConfigMapOutput struct{ *pulumi.OutputState }

func (LogExportConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogExportConfig)(nil)).Elem()
}

func (o LogExportConfigMapOutput) ToLogExportConfigMapOutput() LogExportConfigMapOutput {
	return o
}

func (o LogExportConfigMapOutput) ToLogExportConfigMapOutputWithContext(ctx context.Context) LogExportConfigMapOutput {
	return o
}

func (o LogExportConfigMapOutput) MapIndex(k pulumi.StringInput) LogExportConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LogExportConfig {
		return vs[0].(map[string]*LogExportConfig)[vs[1].(string)]
	}).(LogExportConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogExportConfigInput)(nil)).Elem(), &LogExportConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogExportConfigArrayInput)(nil)).Elem(), LogExportConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogExportConfigMapInput)(nil)).Elem(), LogExportConfigMap{})
	pulumi.RegisterOutputType(LogExportConfigOutput{})
	pulumi.RegisterOutputType(LogExportConfigArrayOutput{})
	pulumi.RegisterOutputType(LogExportConfigMapOutput{})
}
