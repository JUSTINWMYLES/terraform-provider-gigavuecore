package provider

import "context"
import (
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*UploadImageFileAction)(nil)

// UploadImageFileAction is the generated Terraform action implementation.
type UploadImageFileAction struct {
}

// UploadImageFileActionModel describes the action configuration shape.
type UploadImageFileActionModel struct {
	Image types.String `tfsdk:"image"`
}

// NewUploadImageFileAction returns a new instance of the generated action.
func NewUploadImageFileAction() action.Action {
	return &UploadImageFileAction{}
}

// Metadata returns the action type name.
func (r *UploadImageFileAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_upload_image_file"
}

// Schema returns the action schema.
func (r *UploadImageFileAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Upload an image file from local file", Attributes: map[string]schema.Attribute{"image": schema.StringAttribute{MarkdownDescription: "User uploaded file", Required: true}}}
}

// Invoke executes the action against the remote API.
// The generated Invoke method is intentionally stubbed; the remote API is not wired.
func (r *UploadImageFileAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config UploadImageFileActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Invoke is not wired to a remote API endpoint.")
}
