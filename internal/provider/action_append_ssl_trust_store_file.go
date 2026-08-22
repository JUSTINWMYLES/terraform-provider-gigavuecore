package provider

import "context"
import (
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*AppendSslTrustStoreFileAction)(nil)

// AppendSslTrustStoreFileAction is the generated Terraform action implementation.
type AppendSslTrustStoreFileAction struct {
}

// AppendSslTrustStoreFileActionModel describes the action configuration shape.
type AppendSslTrustStoreFileActionModel struct {
	File types.String `tfsdk:"file"`
}

// NewAppendSslTrustStoreFileAction returns a new instance of the generated action.
func NewAppendSslTrustStoreFileAction() action.Action {
	return &AppendSslTrustStoreFileAction{}
}

// Metadata returns the action type name.
func (r *AppendSslTrustStoreFileAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_append_ssl_trust_store_file"
}

// Schema returns the action schema.
func (r *AppendSslTrustStoreFileAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Append inline SSL trust-store from local file", Attributes: map[string]schema.Attribute{"file": schema.StringAttribute{MarkdownDescription: "file to upload to device", Required: true}}}
}

// Invoke executes the action against the remote API.
// The generated Invoke method is intentionally stubbed; the remote API is not wired.
func (r *AppendSslTrustStoreFileAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config AppendSslTrustStoreFileActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Invoke is not wired to a remote API endpoint.")
}
