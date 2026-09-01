package provider

import "context"
import (
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*UploadAcmeServerConfigDetailsAction)(nil)

// UploadAcmeServerConfigDetailsAction is the generated Terraform action implementation.
type UploadAcmeServerConfigDetailsAction struct {
}

// UploadAcmeServerConfigDetailsActionModel describes the action configuration shape.
type UploadAcmeServerConfigDetailsActionModel struct {
	AcmeServerUrl types.String `tfsdk:"acme_server_url"`
	Alias         types.String `tfsdk:"alias"`
	Certificate   types.String `tfsdk:"certificate"`
}

// NewUploadAcmeServerConfigDetailsAction returns a new instance of the generated action.
func NewUploadAcmeServerConfigDetailsAction() action.Action {
	return &UploadAcmeServerConfigDetailsAction{}
}

// Metadata returns the action type name.
func (r *UploadAcmeServerConfigDetailsAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_upload_acme_server_config_details"
}

// Schema returns the action schema.
func (r *UploadAcmeServerConfigDetailsAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "upload ACME server details", Attributes: map[string]schema.Attribute{"acme_server_url": schema.StringAttribute{MarkdownDescription: "ACME server url", Required: true}, "alias": schema.StringAttribute{MarkdownDescription: "ACME server alias", Required: true}, "certificate": schema.StringAttribute{MarkdownDescription: "User uploaded file, only .crt/.cert or .pem format is supported", Required: true}}}
}

// Invoke executes the action against the remote API.
// The generated Invoke method is intentionally stubbed; the remote API is not wired.
func (r *UploadAcmeServerConfigDetailsAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config UploadAcmeServerConfigDetailsActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Invoke is not wired to a remote API endpoint.")
}
