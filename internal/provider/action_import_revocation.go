package provider

import "context"
import (
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*ImportRevocationAction)(nil)

// ImportRevocationAction is the generated Terraform action implementation.
type ImportRevocationAction struct {
}

// ImportRevocationActionModel describes the action configuration shape.
type ImportRevocationActionModel struct {
	RevocationFileName types.String `tfsdk:"revocation_file_name" json:"revocationFileName"`
}

// NewImportRevocationAction returns a new instance of the generated action.
func NewImportRevocationAction() action.Action {
	return &ImportRevocationAction{}
}

// Metadata returns the action type name.
func (r *ImportRevocationAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_import_revocation"
}

// Schema returns the action schema.
func (r *ImportRevocationAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Import FM License Revocation", Attributes: map[string]schema.Attribute{"revocation_file_name": schema.StringAttribute{MarkdownDescription: "User uploaded license file for revocation", Required: true}}}
}

// Invoke executes the action against the remote API.
// The generated Invoke method is intentionally stubbed; the remote API is not wired.
func (r *ImportRevocationAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config ImportRevocationActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Invoke is not wired to a remote API endpoint.")
}
