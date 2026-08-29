package provider

import "context"
import (
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*UploadSslTrustStoreFileAction)(nil)

// UploadSslTrustStoreFileAction is the generated Terraform action implementation.
type UploadSslTrustStoreFileAction struct {
}

// UploadSslTrustStoreFileActionModel describes the action configuration shape.
type UploadSslTrustStoreFileActionModel struct {
	ClusterId types.String `tfsdk:"cluster_id"`
	File      types.String `tfsdk:"file"`
}

// NewUploadSslTrustStoreFileAction returns a new instance of the generated action.
func NewUploadSslTrustStoreFileAction() action.Action {
	return &UploadSslTrustStoreFileAction{}
}

// Metadata returns the action type name.
func (r *UploadSslTrustStoreFileAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_upload_ssl_trust_store_file"
}

// Schema returns the action schema.
func (r *UploadSslTrustStoreFileAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Upload inline SSL trust-store from local file", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "file": schema.StringAttribute{MarkdownDescription: "file to upload to device", Required: true}}}
}

// Invoke executes the action against the remote API.
// The generated Invoke method is intentionally stubbed; the remote API is not wired.
func (r *UploadSslTrustStoreFileAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config UploadSslTrustStoreFileActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Invoke is not wired to a remote API endpoint.")
}
