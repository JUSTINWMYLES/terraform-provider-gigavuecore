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
	Config           types.Dynamic `tfsdk:"config"`
	ConfigLevel      types.String  `tfsdk:"config_level" json:"configLevel"`
	ConfigLevelValue types.List    `tfsdk:"config_level_value" json:"configLevelValue"`
	ConfigType       types.String  `tfsdk:"config_type" json:"configType"`
	Modifiable       types.Bool    `tfsdk:"modifiable"`
	RefCount         types.Int64   `tfsdk:"ref_count" json:"refCount"`
	TemplateName     types.String  `tfsdk:"template_name" json:"templateName"`
	UpdateTime       types.String  `tfsdk:"update_time" json:"updateTime"`
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
	resp.Schema = schema.Schema{Description: "Upload a certificate file to push SSL Certificate Global configuration to all the devices", Attributes: map[string]schema.Attribute{"config": schema.DynamicAttribute{Optional: true}, "config_level": schema.StringAttribute{MarkdownDescription: "Scope of the applied FM template", Optional: true}, "config_level_value": schema.ListAttribute{Optional: true, ElementType: types.StringType}, "config_type": schema.StringAttribute{MarkdownDescription: "Configuration Type of the FM template", Optional: true}, "modifiable": schema.BoolAttribute{Optional: true}, "ref_count": schema.Int64Attribute{Optional: true}, "template_name": schema.StringAttribute{Optional: true}, "update_time": schema.StringAttribute{Optional: true}}}
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
