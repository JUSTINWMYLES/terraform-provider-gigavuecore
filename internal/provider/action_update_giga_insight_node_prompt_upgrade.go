package provider

import "context"
import (
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*UpdateGigaInsightNodePromptUpgradeAction)(nil)

// UpdateGigaInsightNodePromptUpgradeAction is the generated Terraform action implementation.
type UpdateGigaInsightNodePromptUpgradeAction struct {
}

// UpdateGigaInsightNodePromptUpgradeActionModel describes the action configuration shape.
type UpdateGigaInsightNodePromptUpgradeActionModel struct {
	Bundle types.String `tfsdk:"bundle"`
	NodeId types.String `tfsdk:"node_id"`
}

// NewUpdateGigaInsightNodePromptUpgradeAction returns a new instance of the generated action.
func NewUpdateGigaInsightNodePromptUpgradeAction() action.Action {
	return &UpdateGigaInsightNodePromptUpgradeAction{}
}

// Metadata returns the action type name.
func (r *UpdateGigaInsightNodePromptUpgradeAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_update_giga_insight_node_prompt_upgrade"
}

// Schema returns the action schema.
func (r *UpdateGigaInsightNodePromptUpgradeAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Upload prompt bundle upgrade for a GigaInsight Node", Attributes: map[string]schema.Attribute{"bundle": schema.StringAttribute{Optional: true}, "node_id": schema.StringAttribute{Required: true}}}
}

// Invoke executes the action against the remote API.
// The generated Invoke method is intentionally stubbed; the remote API is not wired.
func (r *UpdateGigaInsightNodePromptUpgradeAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config UpdateGigaInsightNodePromptUpgradeActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Invoke is not wired to a remote API endpoint.")
}
