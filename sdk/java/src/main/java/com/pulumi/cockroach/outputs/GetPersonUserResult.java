// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cockroach.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetPersonUserResult {
    /**
     * @return Email address used to find the User ID.
     * 
     */
    private String email;
    /**
     * @return User ID.
     * 
     */
    private String id;

    private GetPersonUserResult() {}
    /**
     * @return Email address used to find the User ID.
     * 
     */
    public String email() {
        return this.email;
    }
    /**
     * @return User ID.
     * 
     */
    public String id() {
        return this.id;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetPersonUserResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String email;
        private String id;
        public Builder() {}
        public Builder(GetPersonUserResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.email = defaults.email;
    	      this.id = defaults.id;
        }

        @CustomType.Setter
        public Builder email(String email) {
            this.email = Objects.requireNonNull(email);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        public GetPersonUserResult build() {
            final var o = new GetPersonUserResult();
            o.email = email;
            o.id = id;
            return o;
        }
    }
}
