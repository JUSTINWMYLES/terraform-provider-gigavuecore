package provider

import "context"
import (
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*UploadSslClientTrustStoreFileAction)(nil)

// UploadSslClientTrustStoreFileAction is the generated Terraform action implementation.
type UploadSslClientTrustStoreFileAction struct {
}

// UploadSslClientTrustStoreFileActionModel describes the action configuration shape.
type UploadSslClientTrustStoreFileActionModel struct {
	Alias     types.String `tfsdk:"alias"`
	ClusterId types.String `tfsdk:"cluster_id"`
	File      types.String `tfsdk:"file"`
}

// NewUploadSslClientTrustStoreFileAction returns a new instance of the generated action.
func NewUploadSslClientTrustStoreFileAction() action.Action {
	return &UploadSslClientTrustStoreFileAction{}
}

// Metadata returns the action type name.
func (r *UploadSslClientTrustStoreFileAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_upload_ssl_client_trust_store_file"
}

// Schema returns the action schema.
func (r *UploadSslClientTrustStoreFileAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Create/Replace the SSL Client trust-store from local file", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Client Trust Store alias", Required: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "file": schema.StringAttribute{MarkdownDescription: "file to upload to device", Required: true}}}
}

// Invoke executes the action against the remote API.
// The generated Invoke method is intentionally stubbed; the remote API is not wired.
func (r *UploadSslClientTrustStoreFileAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config UploadSslClientTrustStoreFileActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Invoke is not wired to a remote API endpoint.")
}
