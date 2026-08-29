package provider

import "context"
import (
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*UploadCertificateSslNodeCertificateUploadNameAction)(nil)

// UploadCertificateSslNodeCertificateUploadNameAction is the generated Terraform action implementation.
type UploadCertificateSslNodeCertificateUploadNameAction struct {
}

// UploadCertificateSslNodeCertificateUploadNameActionModel describes the action configuration shape.
type UploadCertificateSslNodeCertificateUploadNameActionModel struct {
	AddToCaList types.Bool   `tfsdk:"add_to_ca_list"`
	Certificate types.String `tfsdk:"certificate"`
	ClusterId   types.String `tfsdk:"cluster_id"`
	Comment     types.String `tfsdk:"comment"`
	Name        types.String `tfsdk:"name"`
}

// NewUploadCertificateSslNodeCertificateUploadNameAction returns a new instance of the generated action.
func NewUploadCertificateSslNodeCertificateUploadNameAction() action.Action {
	return &UploadCertificateSslNodeCertificateUploadNameAction{}
}

// Metadata returns the action type name.
func (r *UploadCertificateSslNodeCertificateUploadNameAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_upload_certificate_ssl_node_certificate_upload_name"
}

// Schema returns the action schema.
func (r *UploadCertificateSslNodeCertificateUploadNameAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "upload a new certificate and add it to the device' CA List the device", Attributes: map[string]schema.Attribute{"add_to_ca_list": schema.BoolAttribute{Optional: true}, "certificate": schema.StringAttribute{MarkdownDescription: "User uploaded file, only .crt/.cert or .pem format is supported", Required: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "comment": schema.StringAttribute{MarkdownDescription: "a short description of the certificate", Optional: true}, "name": schema.StringAttribute{MarkdownDescription: "the certificate name", Required: true}}}
}

// Invoke executes the action against the remote API.
// The generated Invoke method is intentionally stubbed; the remote API is not wired.
func (r *UploadCertificateSslNodeCertificateUploadNameAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config UploadCertificateSslNodeCertificateUploadNameActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Invoke is not wired to a remote API endpoint.")
}
