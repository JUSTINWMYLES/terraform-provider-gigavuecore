package provider

import "context"
import (
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*UploadGtpWhitelistEntriesFromFileAction)(nil)

// UploadGtpWhitelistEntriesFromFileAction is the generated Terraform action implementation.
type UploadGtpWhitelistEntriesFromFileAction struct {
}

// UploadGtpWhitelistEntriesFromFileActionModel describes the action configuration shape.
type UploadGtpWhitelistEntriesFromFileActionModel struct {
	Alias     types.String `tfsdk:"alias"`
	ClusterId types.String `tfsdk:"cluster_id"`
	Entries   types.String `tfsdk:"entries"`
	Usage     types.String `tfsdk:"usage"`
}

// NewUploadGtpWhitelistEntriesFromFileAction returns a new instance of the generated action.
func NewUploadGtpWhitelistEntriesFromFileAction() action.Action {
	return &UploadGtpWhitelistEntriesFromFileAction{}
}

// Metadata returns the action type name.
func (r *UploadGtpWhitelistEntriesFromFileAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_upload_gtp_whitelist_entries_from_file"
}

// Schema returns the action schema.
func (r *UploadGtpWhitelistEntriesFromFileAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Upload GTP Whitelist Entries from file", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "alias of the target GTP Whitelist", Required: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "entries": schema.StringAttribute{MarkdownDescription: "File containing GTP Whitelist Entries (one entry per line)", Required: true}, "usage": schema.StringAttribute{MarkdownDescription: "When 'delete' is specified, the uploaded list is treated as a delete request and all the entries in the list should be removed from the existing whitelist identified by the 'alias' parameter", Required: true}}}
}

// Invoke executes the action against the remote API.
// The generated Invoke method is intentionally stubbed; the remote API is not wired.
func (r *UploadGtpWhitelistEntriesFromFileAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config UploadGtpWhitelistEntriesFromFileActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Invoke is not wired to a remote API endpoint.")
}
