package provider

import "context"
import (
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*ImportTagsAction)(nil)

// ImportTagsAction is the generated Terraform action implementation.
type ImportTagsAction struct {
}

// ImportTagsActionModel describes the action configuration shape.
type ImportTagsActionModel struct {
	Input     types.String `tfsdk:"input"`
	Operation types.String `tfsdk:"operation"`
}

// NewImportTagsAction returns a new instance of the generated action.
func NewImportTagsAction() action.Action {
	return &ImportTagsAction{}
}

// Metadata returns the action type name.
func (r *ImportTagsAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_import_tags"
}

// Schema returns the action schema.
func (r *ImportTagsAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Upload an csv file from local  to import tags", Attributes: map[string]schema.Attribute{"input": schema.StringAttribute{MarkdownDescription: "User uploaded file, only csv format is supported", Required: true}, "operation": schema.StringAttribute{MarkdownDescription: "specifies the operation type", Required: true}}}
}

// Invoke executes the action against the remote API.
// The generated Invoke method is intentionally stubbed; the remote API is not wired.
func (r *ImportTagsAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config ImportTagsActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Invoke is not wired to a remote API endpoint.")
}
