// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cockroach;

import com.pulumi.cockroach.Utilities;
import com.pulumi.cockroach.inputs.GetClusterCertArgs;
import com.pulumi.cockroach.inputs.GetClusterCertPlainArgs;
import com.pulumi.cockroach.inputs.GetCockroachClusterArgs;
import com.pulumi.cockroach.inputs.GetCockroachClusterPlainArgs;
import com.pulumi.cockroach.inputs.GetConnectionStringArgs;
import com.pulumi.cockroach.inputs.GetConnectionStringPlainArgs;
import com.pulumi.cockroach.inputs.GetPersonUserArgs;
import com.pulumi.cockroach.inputs.GetPersonUserPlainArgs;
import com.pulumi.cockroach.outputs.GetClusterCertResult;
import com.pulumi.cockroach.outputs.GetCockroachClusterResult;
import com.pulumi.cockroach.outputs.GetConnectionStringResult;
import com.pulumi.cockroach.outputs.GetOrganizationResult;
import com.pulumi.cockroach.outputs.GetPersonUserResult;
import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import com.pulumi.resources.InvokeArgs;
import java.util.concurrent.CompletableFuture;

public final class CockroachFunctions {
    /**
     * TLS certificate for the specified CockroachDB cluster. Certificates for dedicated clusters should be written to `$HOME/Library/CockroachCloud/certs/&lt;cluster name&gt;-ca.crt` on MacOS or Linux, or `$env:appdata\CockroachCloud\certs\&lt;cluster name&gt;-ca.crt` on Windows.
     * 
     * Serverless clusters use the root PostgreSQL CA cert. If it isn&#39;t already installed, the certificate can be appended to `$HOME/.postgresql/root.crt` on MacOS or Linux, or `$env:appdata\postgresql\root.crt` on Windows.
     * 
     */
    public static Output<GetClusterCertResult> getClusterCert(GetClusterCertArgs args) {
        return getClusterCert(args, InvokeOptions.Empty);
    }
    /**
     * TLS certificate for the specified CockroachDB cluster. Certificates for dedicated clusters should be written to `$HOME/Library/CockroachCloud/certs/&lt;cluster name&gt;-ca.crt` on MacOS or Linux, or `$env:appdata\CockroachCloud\certs\&lt;cluster name&gt;-ca.crt` on Windows.
     * 
     * Serverless clusters use the root PostgreSQL CA cert. If it isn&#39;t already installed, the certificate can be appended to `$HOME/.postgresql/root.crt` on MacOS or Linux, or `$env:appdata\postgresql\root.crt` on Windows.
     * 
     */
    public static CompletableFuture<GetClusterCertResult> getClusterCertPlain(GetClusterCertPlainArgs args) {
        return getClusterCertPlain(args, InvokeOptions.Empty);
    }
    /**
     * TLS certificate for the specified CockroachDB cluster. Certificates for dedicated clusters should be written to `$HOME/Library/CockroachCloud/certs/&lt;cluster name&gt;-ca.crt` on MacOS or Linux, or `$env:appdata\CockroachCloud\certs\&lt;cluster name&gt;-ca.crt` on Windows.
     * 
     * Serverless clusters use the root PostgreSQL CA cert. If it isn&#39;t already installed, the certificate can be appended to `$HOME/.postgresql/root.crt` on MacOS or Linux, or `$env:appdata\postgresql\root.crt` on Windows.
     * 
     */
    public static Output<GetClusterCertResult> getClusterCert(GetClusterCertArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("cockroach:index/getClusterCert:getClusterCert", TypeShape.of(GetClusterCertResult.class), args, Utilities.withVersion(options));
    }
    /**
     * TLS certificate for the specified CockroachDB cluster. Certificates for dedicated clusters should be written to `$HOME/Library/CockroachCloud/certs/&lt;cluster name&gt;-ca.crt` on MacOS or Linux, or `$env:appdata\CockroachCloud\certs\&lt;cluster name&gt;-ca.crt` on Windows.
     * 
     * Serverless clusters use the root PostgreSQL CA cert. If it isn&#39;t already installed, the certificate can be appended to `$HOME/.postgresql/root.crt` on MacOS or Linux, or `$env:appdata\postgresql\root.crt` on Windows.
     * 
     */
    public static CompletableFuture<GetClusterCertResult> getClusterCertPlain(GetClusterCertPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("cockroach:index/getClusterCert:getClusterCert", TypeShape.of(GetClusterCertResult.class), args, Utilities.withVersion(options));
    }
    /**
     * CockroachDB Cloud cluster. Can be Dedicated or Serverless.
     * 
     */
    public static Output<GetCockroachClusterResult> getCockroachCluster(GetCockroachClusterArgs args) {
        return getCockroachCluster(args, InvokeOptions.Empty);
    }
    /**
     * CockroachDB Cloud cluster. Can be Dedicated or Serverless.
     * 
     */
    public static CompletableFuture<GetCockroachClusterResult> getCockroachClusterPlain(GetCockroachClusterPlainArgs args) {
        return getCockroachClusterPlain(args, InvokeOptions.Empty);
    }
    /**
     * CockroachDB Cloud cluster. Can be Dedicated or Serverless.
     * 
     */
    public static Output<GetCockroachClusterResult> getCockroachCluster(GetCockroachClusterArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("cockroach:index/getCockroachCluster:getCockroachCluster", TypeShape.of(GetCockroachClusterResult.class), args, Utilities.withVersion(options));
    }
    /**
     * CockroachDB Cloud cluster. Can be Dedicated or Serverless.
     * 
     */
    public static CompletableFuture<GetCockroachClusterResult> getCockroachClusterPlain(GetCockroachClusterPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("cockroach:index/getCockroachCluster:getCockroachCluster", TypeShape.of(GetCockroachClusterResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Generic connection string for a cluster.
     * 
     */
    public static Output<GetConnectionStringResult> getConnectionString(GetConnectionStringArgs args) {
        return getConnectionString(args, InvokeOptions.Empty);
    }
    /**
     * Generic connection string for a cluster.
     * 
     */
    public static CompletableFuture<GetConnectionStringResult> getConnectionStringPlain(GetConnectionStringPlainArgs args) {
        return getConnectionStringPlain(args, InvokeOptions.Empty);
    }
    /**
     * Generic connection string for a cluster.
     * 
     */
    public static Output<GetConnectionStringResult> getConnectionString(GetConnectionStringArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("cockroach:index/getConnectionString:getConnectionString", TypeShape.of(GetConnectionStringResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Generic connection string for a cluster.
     * 
     */
    public static CompletableFuture<GetConnectionStringResult> getConnectionStringPlain(GetConnectionStringPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("cockroach:index/getConnectionString:getConnectionString", TypeShape.of(GetConnectionStringResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Information about the organization associated with the user&#39;s API key.
     * 
     */
    public static Output<GetOrganizationResult> getOrganization() {
        return getOrganization(InvokeArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Information about the organization associated with the user&#39;s API key.
     * 
     */
    public static CompletableFuture<GetOrganizationResult> getOrganizationPlain() {
        return getOrganizationPlain(InvokeArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Information about the organization associated with the user&#39;s API key.
     * 
     */
    public static Output<GetOrganizationResult> getOrganization(InvokeArgs args) {
        return getOrganization(args, InvokeOptions.Empty);
    }
    /**
     * Information about the organization associated with the user&#39;s API key.
     * 
     */
    public static CompletableFuture<GetOrganizationResult> getOrganizationPlain(InvokeArgs args) {
        return getOrganizationPlain(args, InvokeOptions.Empty);
    }
    /**
     * Information about the organization associated with the user&#39;s API key.
     * 
     */
    public static Output<GetOrganizationResult> getOrganization(InvokeArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("cockroach:index/getOrganization:getOrganization", TypeShape.of(GetOrganizationResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Information about the organization associated with the user&#39;s API key.
     * 
     */
    public static CompletableFuture<GetOrganizationResult> getOrganizationPlain(InvokeArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("cockroach:index/getOrganization:getOrganization", TypeShape.of(GetOrganizationResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Information about an individual user.
     * 
     */
    public static Output<GetPersonUserResult> getPersonUser(GetPersonUserArgs args) {
        return getPersonUser(args, InvokeOptions.Empty);
    }
    /**
     * Information about an individual user.
     * 
     */
    public static CompletableFuture<GetPersonUserResult> getPersonUserPlain(GetPersonUserPlainArgs args) {
        return getPersonUserPlain(args, InvokeOptions.Empty);
    }
    /**
     * Information about an individual user.
     * 
     */
    public static Output<GetPersonUserResult> getPersonUser(GetPersonUserArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("cockroach:index/getPersonUser:getPersonUser", TypeShape.of(GetPersonUserResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Information about an individual user.
     * 
     */
    public static CompletableFuture<GetPersonUserResult> getPersonUserPlain(GetPersonUserPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("cockroach:index/getPersonUser:getPersonUser", TypeShape.of(GetPersonUserResult.class), args, Utilities.withVersion(options));
    }
}
