// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ## Import
 *
 * format: <user id>
 *
 * ```sh
 * $ pulumi import cockroach:index/userRoleGrants:UserRoleGrants service_account 1f69fdd2-600a-4cfc-a9ba-16995df0d77d
 * ```
 */
export class UserRoleGrants extends pulumi.CustomResource {
    /**
     * Get an existing UserRoleGrants resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserRoleGrantsState, opts?: pulumi.CustomResourceOptions): UserRoleGrants {
        return new UserRoleGrants(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'cockroach:index/userRoleGrants:UserRoleGrants';

    /**
     * Returns true if the given object is an instance of UserRoleGrants.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserRoleGrants {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserRoleGrants.__pulumiType;
    }

    /**
     * The list of roles to include. ORG_MEMBER must be included.
     */
    public readonly roles!: pulumi.Output<outputs.UserRoleGrantsRole[]>;
    /**
     * ID of the user to grant these roles to.
     */
    public readonly userId!: pulumi.Output<string>;

    /**
     * Create a UserRoleGrants resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserRoleGrantsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserRoleGrantsArgs | UserRoleGrantsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserRoleGrantsState | undefined;
            resourceInputs["roles"] = state ? state.roles : undefined;
            resourceInputs["userId"] = state ? state.userId : undefined;
        } else {
            const args = argsOrState as UserRoleGrantsArgs | undefined;
            if ((!args || args.roles === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roles'");
            }
            if ((!args || args.userId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userId'");
            }
            resourceInputs["roles"] = args ? args.roles : undefined;
            resourceInputs["userId"] = args ? args.userId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(UserRoleGrants.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserRoleGrants resources.
 */
export interface UserRoleGrantsState {
    /**
     * The list of roles to include. ORG_MEMBER must be included.
     */
    roles?: pulumi.Input<pulumi.Input<inputs.UserRoleGrantsRole>[]>;
    /**
     * ID of the user to grant these roles to.
     */
    userId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UserRoleGrants resource.
 */
export interface UserRoleGrantsArgs {
    /**
     * The list of roles to include. ORG_MEMBER must be included.
     */
    roles: pulumi.Input<pulumi.Input<inputs.UserRoleGrantsRole>[]>;
    /**
     * ID of the user to grant these roles to.
     */
    userId: pulumi.Input<string>;
}
