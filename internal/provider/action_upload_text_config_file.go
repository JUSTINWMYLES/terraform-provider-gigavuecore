package provider

import "context"
import (
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*UploadTextConfigFileAction)(nil)

// UploadTextConfigFileAction is the generated Terraform action implementation.
type UploadTextConfigFileAction struct {
}

// UploadTextConfigFileActionModel describes the action configuration shape.
type UploadTextConfigFileActionModel struct {
	ClusterId types.String `tfsdk:"cluster_id"`
	Config    types.String `tfsdk:"config"`
}

// NewUploadTextConfigFileAction returns a new instance of the generated action.
func NewUploadTextConfigFileAction() action.Action {
	return &UploadTextConfigFileAction{}
}

// Metadata returns the action type name.
func (r *UploadTextConfigFileAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_upload_text_config_file"
}

// Schema returns the action schema.
func (r *UploadTextConfigFileAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Upload a text configuration file from local file.", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "config": schema.StringAttribute{MarkdownDescription: "User uploaded text configuration file", Required: true}}}
}

// Invoke executes the action against the remote API.
// The generated Invoke method is intentionally stubbed; the remote API is not wired.
func (r *UploadTextConfigFileAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config UploadTextConfigFileActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Invoke is not wired to a remote API endpoint.")
}
