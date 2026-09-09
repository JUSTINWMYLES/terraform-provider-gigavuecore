package provider

import "context"
import (
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*ImportTagResourcesAction)(nil)

// ImportTagResourcesAction is the generated Terraform action implementation.
type ImportTagResourcesAction struct {
}

// ImportTagResourcesActionModel describes the action configuration shape.
type ImportTagResourcesActionModel struct {
	Input     types.String `tfsdk:"input"`
	Operation types.String `tfsdk:"operation"`
}

// NewImportTagResourcesAction returns a new instance of the generated action.
func NewImportTagResourcesAction() action.Action {
	return &ImportTagResourcesAction{}
}

// Metadata returns the action type name.
func (r *ImportTagResourcesAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_import_tag_resources"
}

// Schema returns the action schema.
func (r *ImportTagResourcesAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Upload an csv file from local to import tag resources", Attributes: map[string]schema.Attribute{"input": schema.StringAttribute{MarkdownDescription: "User uploaded file, only csv format is supported", Required: true}, "operation": schema.StringAttribute{MarkdownDescription: "specifies the operation type", Required: true}}}
}

// Invoke executes the action against the remote API.
// The generated Invoke method is intentionally stubbed; the remote API is not wired.
func (r *ImportTagResourcesAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config ImportTagResourcesActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Invoke is not wired to a remote API endpoint.")
}
