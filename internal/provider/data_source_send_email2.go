package provider

import "context"
import (
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var _ datasource.DataSource = (*SendEmail2DataSource)(nil)

// SendEmail2DataSource is the generated Terraform data source implementation.
type SendEmail2DataSource struct {
}

// SendEmail2DataSourceModel describes the data source state shape.
type SendEmail2DataSourceModel struct {
	Date types.Int64 `tfsdk:"date"`
}

// NewSendEmail2DataSource returns a new instance of the generated data source.
func NewSendEmail2DataSource() datasource.DataSource {
	return &SendEmail2DataSource{}
}

// Metadata returns the data source type name.
func (d *SendEmail2DataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_send_email2"
}

// Schema returns the data source schema.
func (d *SendEmail2DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Emails the usage report in PDF format for period containing date in format YYYYMMDD", Attributes: map[string]schema.Attribute{"date": schema.Int64Attribute{MarkdownDescription: "Date of processed volumes to retrieve in format YYYYMMDD", Required: true}}}
}

// Read fetches remote state into the data source model.
func (d *SendEmail2DataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config SendEmail2DataSourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.State.Set(ctx, &config)
}
