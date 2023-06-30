// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cockroach.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CaCertState extends com.pulumi.resources.ResourceArgs {

    public static final CaCertState Empty = new CaCertState();

    /**
     * Status of client CA certs on a cluster
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return Status of client CA certs on a cluster
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * X509 certificate in PEM format
     * 
     */
    @Import(name="x509PemCert")
    private @Nullable Output<String> x509PemCert;

    /**
     * @return X509 certificate in PEM format
     * 
     */
    public Optional<Output<String>> x509PemCert() {
        return Optional.ofNullable(this.x509PemCert);
    }

    private CaCertState() {}

    private CaCertState(CaCertState $) {
        this.status = $.status;
        this.x509PemCert = $.x509PemCert;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CaCertState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CaCertState $;

        public Builder() {
            $ = new CaCertState();
        }

        public Builder(CaCertState defaults) {
            $ = new CaCertState(Objects.requireNonNull(defaults));
        }

        /**
         * @param status Status of client CA certs on a cluster
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status Status of client CA certs on a cluster
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param x509PemCert X509 certificate in PEM format
         * 
         * @return builder
         * 
         */
        public Builder x509PemCert(@Nullable Output<String> x509PemCert) {
            $.x509PemCert = x509PemCert;
            return this;
        }

        /**
         * @param x509PemCert X509 certificate in PEM format
         * 
         * @return builder
         * 
         */
        public Builder x509PemCert(String x509PemCert) {
            return x509PemCert(Output.of(x509PemCert));
        }

        public CaCertState build() {
            return $;
        }
    }

}