// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cockroach.outputs;

import com.pulumi.cockroach.outputs.PrivateEndpointServicesServiceAws;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class PrivateEndpointServicesService {
    private @Nullable PrivateEndpointServicesServiceAws aws;
    private @Nullable String cloudProvider;
    private @Nullable String regionName;
    private @Nullable String status;

    private PrivateEndpointServicesService() {}
    public Optional<PrivateEndpointServicesServiceAws> aws() {
        return Optional.ofNullable(this.aws);
    }
    public Optional<String> cloudProvider() {
        return Optional.ofNullable(this.cloudProvider);
    }
    public Optional<String> regionName() {
        return Optional.ofNullable(this.regionName);
    }
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(PrivateEndpointServicesService defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable PrivateEndpointServicesServiceAws aws;
        private @Nullable String cloudProvider;
        private @Nullable String regionName;
        private @Nullable String status;
        public Builder() {}
        public Builder(PrivateEndpointServicesService defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.aws = defaults.aws;
    	      this.cloudProvider = defaults.cloudProvider;
    	      this.regionName = defaults.regionName;
    	      this.status = defaults.status;
        }

        @CustomType.Setter
        public Builder aws(@Nullable PrivateEndpointServicesServiceAws aws) {
            this.aws = aws;
            return this;
        }
        @CustomType.Setter
        public Builder cloudProvider(@Nullable String cloudProvider) {
            this.cloudProvider = cloudProvider;
            return this;
        }
        @CustomType.Setter
        public Builder regionName(@Nullable String regionName) {
            this.regionName = regionName;
            return this;
        }
        @CustomType.Setter
        public Builder status(@Nullable String status) {
            this.status = status;
            return this;
        }
        public PrivateEndpointServicesService build() {
            final var o = new PrivateEndpointServicesService();
            o.aws = aws;
            o.cloudProvider = cloudProvider;
            o.regionName = regionName;
            o.status = status;
            return o;
        }
    }
}