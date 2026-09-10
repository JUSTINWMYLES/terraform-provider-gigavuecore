package provider

import "context"
import (
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*DeactivateByAidAction)(nil)

// DeactivateByAidAction is the generated Terraform action implementation.
type DeactivateByAidAction struct {
}

// DeactivateByAidActionModel describes the action configuration shape.
type DeactivateByAidActionModel struct {
	Aid types.String `tfsdk:"aid"`
}

// NewDeactivateByAidAction returns a new instance of the generated action.
func NewDeactivateByAidAction() action.Action {
	return &DeactivateByAidAction{}
}

// Metadata returns the action type name.
func (r *DeactivateByAidAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_deactivate_by_aid"
}

// Schema returns the action schema.
func (r *DeactivateByAidAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Deactivate a license (floating or VBL) by its Activation ID; floating licenses still assigned to at least one card or chassis will be skipped", Attributes: map[string]schema.Attribute{"aid": schema.StringAttribute{MarkdownDescription: "Activation ID (created when license is generated) of the feature activation", Required: true}}}
}

// Invoke executes the action against the remote API.
// The generated Invoke method is intentionally stubbed; the remote API is not wired.
func (r *DeactivateByAidAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config DeactivateByAidActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Invoke is not wired to a remote API endpoint.")
}
