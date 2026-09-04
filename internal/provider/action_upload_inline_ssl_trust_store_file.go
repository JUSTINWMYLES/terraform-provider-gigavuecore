package provider

import "context"
import (
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*UploadInlineSslTrustStoreFileAction)(nil)

// UploadInlineSslTrustStoreFileAction is the generated Terraform action implementation.
type UploadInlineSslTrustStoreFileAction struct {
}

// UploadInlineSslTrustStoreFileActionModel describes the action configuration shape.
type UploadInlineSslTrustStoreFileActionModel struct {
	ClusterId types.String `tfsdk:"cluster_id"`
	File      types.String `tfsdk:"file"`
}

// NewUploadInlineSslTrustStoreFileAction returns a new instance of the generated action.
func NewUploadInlineSslTrustStoreFileAction() action.Action {
	return &UploadInlineSslTrustStoreFileAction{}
}

// Metadata returns the action type name.
func (r *UploadInlineSslTrustStoreFileAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_upload_inline_ssl_trust_store_file"
}

// Schema returns the action schema.
func (r *UploadInlineSslTrustStoreFileAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Upload inline SSL trust-store from local file", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "file": schema.StringAttribute{MarkdownDescription: "file to upload to device", Required: true}}}
}

// Invoke executes the action against the remote API.
// The generated Invoke method is intentionally stubbed; the remote API is not wired.
func (r *UploadInlineSslTrustStoreFileAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config UploadInlineSslTrustStoreFileActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Invoke is not wired to a remote API endpoint.")
}
