// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.cockroach.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FolderState extends com.pulumi.resources.ResourceArgs {

    public static final FolderState Empty = new FolderState();

    /**
     * Name of the folder.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the folder.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * ID of the parent folder. Use &#39;root&#39; for the root level (no parent folder).
     * 
     */
    @Import(name="parentId")
    private @Nullable Output<String> parentId;

    /**
     * @return ID of the parent folder. Use &#39;root&#39; for the root level (no parent folder).
     * 
     */
    public Optional<Output<String>> parentId() {
        return Optional.ofNullable(this.parentId);
    }

    private FolderState() {}

    private FolderState(FolderState $) {
        this.name = $.name;
        this.parentId = $.parentId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FolderState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FolderState $;

        public Builder() {
            $ = new FolderState();
        }

        public Builder(FolderState defaults) {
            $ = new FolderState(Objects.requireNonNull(defaults));
        }

        /**
         * @param name Name of the folder.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the folder.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param parentId ID of the parent folder. Use &#39;root&#39; for the root level (no parent folder).
         * 
         * @return builder
         * 
         */
        public Builder parentId(@Nullable Output<String> parentId) {
            $.parentId = parentId;
            return this;
        }

        /**
         * @param parentId ID of the parent folder. Use &#39;root&#39; for the root level (no parent folder).
         * 
         * @return builder
         * 
         */
        public Builder parentId(String parentId) {
            return parentId(Output.of(parentId));
        }

        public FolderState build() {
            return $;
        }
    }

}
