package provider

import "context"
import (
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*AppendInlineSslTrustStoreFileAction)(nil)

// AppendInlineSslTrustStoreFileAction is the generated Terraform action implementation.
type AppendInlineSslTrustStoreFileAction struct {
}

// AppendInlineSslTrustStoreFileActionModel describes the action configuration shape.
type AppendInlineSslTrustStoreFileActionModel struct {
	File types.String `tfsdk:"file"`
}

// NewAppendInlineSslTrustStoreFileAction returns a new instance of the generated action.
func NewAppendInlineSslTrustStoreFileAction() action.Action {
	return &AppendInlineSslTrustStoreFileAction{}
}

// Metadata returns the action type name.
func (r *AppendInlineSslTrustStoreFileAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_append_inline_ssl_trust_store_file"
}

// Schema returns the action schema.
func (r *AppendInlineSslTrustStoreFileAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Append inline SSL trust-store from local file", Attributes: map[string]schema.Attribute{"file": schema.StringAttribute{MarkdownDescription: "file to upload to device", Required: true}}}
}

// Invoke executes the action against the remote API.
// The generated Invoke method is intentionally stubbed; the remote API is not wired.
func (r *AppendInlineSslTrustStoreFileAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config AppendInlineSslTrustStoreFileActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Invoke is not wired to a remote API endpoint.")
}
