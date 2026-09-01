package provider

import "context"
import (
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var _ datasource.DataSource = (*SendEmailDataSource)(nil)

// SendEmailDataSource is the generated Terraform data source implementation.
type SendEmailDataSource struct {
}

// SendEmailDataSourceModel describes the data source state shape.
type SendEmailDataSourceModel struct {
	Date types.Int64 `tfsdk:"date"`
}

// NewSendEmailDataSource returns a new instance of the generated data source.
func NewSendEmailDataSource() datasource.DataSource {
	return &SendEmailDataSource{}
}

// Metadata returns the data source type name.
func (d *SendEmailDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_send_email"
}

// Schema returns the data source schema.
func (d *SendEmailDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Emails the usage report as encoded JSON for period containing date in format YYYYMMDD", Attributes: map[string]schema.Attribute{"date": schema.Int64Attribute{MarkdownDescription: "Date of processed volumes to retrieve in format YYYYMMDD", Required: true}}}
}

// Read fetches remote state into the data source model.
func (d *SendEmailDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config SendEmailDataSourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.State.Set(ctx, &config)
}
