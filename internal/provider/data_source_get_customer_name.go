package provider

import "context"
import (
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

// Compile-time interface assertion.
var _ datasource.DataSource = (*GetCustomerNameDataSource)(nil)

// GetCustomerNameDataSource is the generated Terraform data source implementation.
type GetCustomerNameDataSource struct {
}

// GetCustomerNameDataSourceModel describes the data source state shape.
type GetCustomerNameDataSourceModel struct {
}

// NewGetCustomerNameDataSource returns a new instance of the generated data source.
func NewGetCustomerNameDataSource() datasource.DataSource {
	return &GetCustomerNameDataSource{}
}

// Metadata returns the data source type name.
func (d *GetCustomerNameDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_customer_name"
}

// Schema returns the data source schema.
func (d *GetCustomerNameDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Returns the name of the customer that owns this FM"}
}

// Read fetches remote state into the data source model.
func (d *GetCustomerNameDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetCustomerNameDataSourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.State.Set(ctx, &config)
}
