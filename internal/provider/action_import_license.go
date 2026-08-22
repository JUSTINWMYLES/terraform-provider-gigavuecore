package provider

import "context"
import (
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*ImportLicenseAction)(nil)

// ImportLicenseAction is the generated Terraform action implementation.
type ImportLicenseAction struct {
}

// ImportLicenseActionModel describes the action configuration shape.
type ImportLicenseActionModel struct {
	LicFileName types.String `tfsdk:"lic_file_name" json:"licFileName"`
}

// NewImportLicenseAction returns a new instance of the generated action.
func NewImportLicenseAction() action.Action {
	return &ImportLicenseAction{}
}

// Metadata returns the action type name.
func (r *ImportLicenseAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_import_license"
}

// Schema returns the action schema.
func (r *ImportLicenseAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Import FM License", Attributes: map[string]schema.Attribute{"lic_file_name": schema.StringAttribute{MarkdownDescription: "User uploaded license file", Required: true}}}
}

// Invoke executes the action against the remote API.
// The generated Invoke method is intentionally stubbed; the remote API is not wired.
func (r *ImportLicenseAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config ImportLicenseActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Invoke is not wired to a remote API endpoint.")
}
