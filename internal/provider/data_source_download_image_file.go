package provider

import "context"
import (
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var _ datasource.DataSource = (*DownloadImageFileDataSource)(nil)

// DownloadImageFileDataSource is the generated Terraform data source implementation.
type DownloadImageFileDataSource struct {
}

// DownloadImageFileDataSourceModel describes the data source state shape.
type DownloadImageFileDataSourceModel struct {
	FileName types.String `tfsdk:"file_name" json:"fileName"`
}

// NewDownloadImageFileDataSource returns a new instance of the generated data source.
func NewDownloadImageFileDataSource() datasource.DataSource {
	return &DownloadImageFileDataSource{}
}

// Metadata returns the data source type name.
func (d *DownloadImageFileDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_download_image_file"
}

// Schema returns the data source schema.
func (d *DownloadImageFileDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Download Image File", Attributes: map[string]schema.Attribute{"file_name": schema.StringAttribute{MarkdownDescription: "Filename of the target image file", Required: true}}}
}

// Read fetches remote state into the data source model.
func (d *DownloadImageFileDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config DownloadImageFileDataSourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.State.Set(ctx, &config)
}
