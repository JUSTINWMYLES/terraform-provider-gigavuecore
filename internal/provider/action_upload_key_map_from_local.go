package provider

import "context"
import (
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*UploadKeyMapFromLocalAction)(nil)

// UploadKeyMapFromLocalAction is the generated Terraform action implementation.
type UploadKeyMapFromLocalAction struct {
}

// UploadKeyMapFromLocalActionModel describes the action configuration shape.
type UploadKeyMapFromLocalActionModel struct {
	Alias     types.String `tfsdk:"alias"`
	ClusterId types.String `tfsdk:"cluster_id"`
	File      types.String `tfsdk:"file"`
}

// NewUploadKeyMapFromLocalAction returns a new instance of the generated action.
func NewUploadKeyMapFromLocalAction() action.Action {
	return &UploadKeyMapFromLocalAction{}
}

// Metadata returns the action type name.
func (r *UploadKeyMapFromLocalAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_upload_key_map_from_local"
}

// Schema returns the action schema.
func (r *UploadKeyMapFromLocalAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Upload Key Map file from local", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "alias of target HSM Group", Required: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target cluster ID.", Required: true}, "file": schema.StringAttribute{MarkdownDescription: "key map file to upload", Required: true}}}
}

// Invoke executes the action against the remote API.
// The generated Invoke method is intentionally stubbed; the remote API is not wired.
func (r *UploadKeyMapFromLocalAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config UploadKeyMapFromLocalActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Invoke is not wired to a remote API endpoint.")
}
