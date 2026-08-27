package provider

import "context"
import (
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*UploadSystemConfigFileAction)(nil)

// UploadSystemConfigFileAction is the generated Terraform action implementation.
type UploadSystemConfigFileAction struct {
}

// UploadSystemConfigFileActionModel describes the action configuration shape.
type UploadSystemConfigFileActionModel struct {
	ClusterId types.String `tfsdk:"cluster_id"`
	Config    types.String `tfsdk:"config"`
}

// NewUploadSystemConfigFileAction returns a new instance of the generated action.
func NewUploadSystemConfigFileAction() action.Action {
	return &UploadSystemConfigFileAction{}
}

// Metadata returns the action type name.
func (r *UploadSystemConfigFileAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_upload_system_config_file"
}

// Schema returns the action schema.
func (r *UploadSystemConfigFileAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Upload a configuration file from local file.", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "config": schema.StringAttribute{MarkdownDescription: "User uploaded file", Required: true}}}
}

// Invoke executes the action against the remote API.
// The generated Invoke method is intentionally stubbed; the remote API is not wired.
func (r *UploadSystemConfigFileAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config UploadSystemConfigFileActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Invoke is not wired to a remote API endpoint.")
}
