package provider

import "context"
import (
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*UploadDeviceConfigFileAction)(nil)

// UploadDeviceConfigFileAction is the generated Terraform action implementation.
type UploadDeviceConfigFileAction struct {
}

// UploadDeviceConfigFileActionModel describes the action configuration shape.
type UploadDeviceConfigFileActionModel struct {
	File types.String `tfsdk:"file"`
}

// NewUploadDeviceConfigFileAction returns a new instance of the generated action.
func NewUploadDeviceConfigFileAction() action.Action {
	return &UploadDeviceConfigFileAction{}
}

// Metadata returns the action type name.
func (r *UploadDeviceConfigFileAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_upload_device_config_file"
}

// Schema returns the action schema.
func (r *UploadDeviceConfigFileAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "upload device config file", Attributes: map[string]schema.Attribute{"file": schema.StringAttribute{MarkdownDescription: "Attached Config file. In text format", Required: true}}}
}

// Invoke executes the action against the remote API.
// The generated Invoke method is intentionally stubbed; the remote API is not wired.
func (r *UploadDeviceConfigFileAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config UploadDeviceConfigFileActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Invoke is not wired to a remote API endpoint.")
}
