package provider

import "context"
import (
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*UploadCertificateAction)(nil)

// UploadCertificateAction is the generated Terraform action implementation.
type UploadCertificateAction struct {
}

// UploadCertificateActionModel describes the action configuration shape.
type UploadCertificateActionModel struct {
	Config           types.Object `tfsdk:"config"`
	ConfigLevel      types.String `tfsdk:"config_level" json:"configLevel"`
	ConfigLevelValue types.List   `tfsdk:"config_level_value" json:"configLevelValue"`
	ConfigType       types.String `tfsdk:"config_type" json:"configType"`
	Modifiable       types.Bool   `tfsdk:"modifiable"`
	RefCount         types.Int64  `tfsdk:"ref_count" json:"refCount"`
	TemplateName     types.String `tfsdk:"template_name" json:"templateName"`
	UpdateTime       types.String `tfsdk:"update_time" json:"updateTime"`
}

// NewUploadCertificateAction returns a new instance of the generated action.
func NewUploadCertificateAction() action.Action {
	return &UploadCertificateAction{}
}

// Metadata returns the action type name.
func (r *UploadCertificateAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_upload_certificate"
}

// Schema returns the action schema.
func (r *UploadCertificateAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Upload a certificate file to push SSL Certificate Global configuration to all the devices", Attributes: map[string]schema.Attribute{"config": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"device_ssl_certificate_configs": schema.ListNestedAttribute{MarkdownDescription: "Available when ConfigType is SSL_CERTIFICATE_TEMPLATE", Optional: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"issuer": schema.StringAttribute{MarkdownDescription: "issuer details of the certificate", Optional: true}, "not_after": schema.StringAttribute{MarkdownDescription: "date and time when certificate stops being valid ([rfc3339](https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14))", Optional: true}, "not_before": schema.StringAttribute{MarkdownDescription: "date and time when certificate starts being valid ([rfc3339](https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14))", Optional: true}, "operation_type": schema.StringAttribute{Required: true}, "signature_algorithm": schema.StringAttribute{Optional: true}, "subject": schema.StringAttribute{MarkdownDescription: "subject name of the certificate", Optional: true}, "trusted_ca": schema.SingleNestedAttribute{Required: true, Attributes: map[string]schema.Attribute{"name": schema.StringAttribute{MarkdownDescription: "name of the certificate", Required: true}}}, "upload_spec": schema.SingleNestedAttribute{Required: true, Attributes: map[string]schema.Attribute{"info": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"comment": schema.StringAttribute{MarkdownDescription: "a short description of the certificate", Optional: true}, "name": schema.StringAttribute{MarkdownDescription: "name of the certificate", Required: true}, "passphrase": schema.StringAttribute{MarkdownDescription: "used to decrypt pkcs12 and private keys", Optional: true}, "type": schema.StringAttribute{MarkdownDescription: "type of the certificate", Required: true}}}, "pem": schema.StringAttribute{MarkdownDescription: "contents of the certificate in pem format", Optional: true}}}}}}}}, "config_level": schema.StringAttribute{MarkdownDescription: "Scope of the applied FM template", Optional: true}, "config_level_value": schema.ListAttribute{Optional: true, ElementType: types.StringType}, "config_type": schema.StringAttribute{MarkdownDescription: "Configuration Type of the FM template", Optional: true}, "modifiable": schema.BoolAttribute{Optional: true}, "ref_count": schema.Int64Attribute{Optional: true}, "template_name": schema.StringAttribute{Optional: true}, "update_time": schema.StringAttribute{Optional: true}}}
}

// Invoke executes the action against the remote API.
// The generated Invoke method is intentionally stubbed; the remote API is not wired.
func (r *UploadCertificateAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config UploadCertificateActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Invoke is not wired to a remote API endpoint.")
}
