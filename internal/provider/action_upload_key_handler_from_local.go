package provider

import "context"
import (
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*UploadKeyHandlerFromLocalAction)(nil)

// UploadKeyHandlerFromLocalAction is the generated Terraform action implementation.
type UploadKeyHandlerFromLocalAction struct {
}

// UploadKeyHandlerFromLocalActionModel describes the action configuration shape.
type UploadKeyHandlerFromLocalActionModel struct {
	Alias types.String `tfsdk:"alias"`
	File  types.String `tfsdk:"file"`
}

// NewUploadKeyHandlerFromLocalAction returns a new instance of the generated action.
func NewUploadKeyHandlerFromLocalAction() action.Action {
	return &UploadKeyHandlerFromLocalAction{}
}

// Metadata returns the action type name.
func (r *UploadKeyHandlerFromLocalAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_upload_key_handler_from_local"
}

// Schema returns the action schema.
func (r *UploadKeyHandlerFromLocalAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Upload world and module file from local", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "alias of target HSM Group", Required: true}, "file": schema.StringAttribute{MarkdownDescription: "world or module file to upload", Required: true}}}
}

// Invoke executes the action against the remote API.
// The generated Invoke method is intentionally stubbed; the remote API is not wired.
func (r *UploadKeyHandlerFromLocalAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config UploadKeyHandlerFromLocalActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Invoke is not wired to a remote API endpoint.")
}
