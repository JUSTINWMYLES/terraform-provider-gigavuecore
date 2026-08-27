package provider

import "context"
import (
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*UploadKeystoreKeyAction)(nil)

// UploadKeystoreKeyAction is the generated Terraform action implementation.
type UploadKeystoreKeyAction struct {
}

// UploadKeystoreKeyActionModel describes the action configuration shape.
type UploadKeystoreKeyActionModel struct {
	Alias      types.String `tfsdk:"alias"`
	BodyAlias  types.String `tfsdk:"body_alias" json:"alias"`
	ClusterId  types.String `tfsdk:"cluster_id"`
	Comment    types.String `tfsdk:"comment"`
	File       types.String `tfsdk:"file"`
	Passphrase types.String `tfsdk:"passphrase"`
	Type       types.String `tfsdk:"type"`
}

// NewUploadKeystoreKeyAction returns a new instance of the generated action.
func NewUploadKeystoreKeyAction() action.Action {
	return &UploadKeystoreKeyAction{}
}

// Metadata returns the action type name.
func (r *UploadKeystoreKeyAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_upload_keystore_key"
}

// Schema returns the action schema.
func (r *UploadKeystoreKeyAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "upload a key to the keystore from local file", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "alias for the key", Required: true}, "body_alias": schema.StringAttribute{MarkdownDescription: "alias for the key", Required: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "comment": schema.StringAttribute{MarkdownDescription: "key comment", Optional: true}, "file": schema.StringAttribute{MarkdownDescription: "file to upload to device", Required: true}, "passphrase": schema.StringAttribute{MarkdownDescription: "passphrase applicable for pkcs12 only", Optional: true}, "type": schema.StringAttribute{MarkdownDescription: "key type", Required: true}}}
}

// Invoke executes the action against the remote API.
// The generated Invoke method is intentionally stubbed; the remote API is not wired.
func (r *UploadKeystoreKeyAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config UploadKeystoreKeyActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Invoke is not wired to a remote API endpoint.")
}
