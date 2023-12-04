// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface ApiOidcConfigIdentityMap {
    /**
     * The username (email or service account id) of the CC user that the token should map to.
     */
    ccIdentity: string;
    /**
     * Indicates that the tokenPrincipal field is a regex value.
     */
    isRegex: boolean;
    /**
     * The token value that needs to be mapped.
     */
    tokenIdentity: string;
}

export interface ClusterDedicated {
    diskIops: number;
    machineType: string;
    memoryGib: number;
    numVirtualCpus: number;
    privateNetworkVisibility: boolean;
    storageGib: number;
}

export interface ClusterRegion {
    internalDns: string;
    name: string;
    nodeCount: number;
    primary: boolean;
    sqlDns: string;
    uiDns: string;
}

export interface ClusterServerless {
    routingId: string;
    spendLimit?: number;
    usageLimits?: outputs.ClusterServerlessUsageLimits;
}

export interface ClusterServerlessUsageLimits {
    requestUnitLimit: number;
    storageMibLimit: number;
}

export interface CmekAdditionalRegion {
    internalDns: string;
    name: string;
    nodeCount: number;
    primary: boolean;
    sqlDns: string;
    uiDns: string;
}

export interface CmekRegion {
    key: outputs.CmekRegionKey;
    region: string;
    status: string;
}

export interface CmekRegionKey {
    authPrincipal: string;
    createdAt: string;
    status: string;
    type: string;
    updatedAt: string;
    uri: string;
    userMessage: string;
}

export interface GetCockroachClusterDedicated {
    diskIops: number;
    machineType: string;
    memoryGib: number;
    numVirtualCpus: number;
    privateNetworkVisibility: boolean;
    storageGib: number;
}

export interface GetCockroachClusterRegion {
    internalDns: string;
    name: string;
    nodeCount: number;
    primary: boolean;
    sqlDns: string;
    uiDns: string;
}

export interface GetCockroachClusterServerless {
    routingId: string;
    spendLimit: number;
    usageLimits: outputs.GetCockroachClusterServerlessUsageLimits;
}

export interface GetCockroachClusterServerlessUsageLimits {
    requestUnitLimit: number;
    storageMibLimit: number;
}

export interface GetConnectionStringConnectionParams {
    database: string;
    host: string;
    password: string;
    port: string;
    username: string;
}

export interface LogExportConfigGroup {
    channels: string[];
    logName: string;
    minLevel?: string;
    redact?: boolean;
}

export interface PrivateEndpointServicesService {
    aws: outputs.PrivateEndpointServicesServiceAws;
    /**
     * Cloud provider associated with this service.
     */
    cloudProvider: string;
    /**
     * Cloud provider region code associated with this service.
     */
    regionName: string;
    /**
     * Operation status of the service.
     */
    status: string;
}

export interface PrivateEndpointServicesServiceAws {
    availabilityZoneIds: string[];
    serviceId: string;
    serviceName: string;
}

export interface UserRoleGrantsRole {
    resourceId?: string;
    resourceType: string;
    roleName: string;
}

