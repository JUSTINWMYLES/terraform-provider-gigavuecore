package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)
import (
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*GetAllClusterLicensesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetAllClusterLicensesDataSource)(nil)
)

// GetAllClusterLicensesDataSource is the generated Terraform data source implementation.
type GetAllClusterLicensesDataSource struct {
	client *client.Client
}

// GetAllClusterLicensesDataSourceModel describes the data source state shape.
type GetAllClusterLicensesDataSourceModel struct {
	Items types.List `tfsdk:"items"`
}

// NewGetAllClusterLicensesDataSource returns a new instance of the generated data source.
func NewGetAllClusterLicensesDataSource() datasource.DataSource {
	return &GetAllClusterLicensesDataSource{}
}

// Metadata returns the data source type name.
func (d *GetAllClusterLicensesDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_all_cluster_licenses"
}

// Schema returns the data source schema.
func (d *GetAllClusterLicensesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get all clusters with all the chassis and cards information and licenses installed in  Hierarchical View", Attributes: map[string]schema.Attribute{"items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"chassis_modules": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"box_id": schema.Int64Attribute{Computed: true}, "cards": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"box_id": schema.Int64Attribute{Computed: true}, "cluster_name": schema.StringAttribute{MarkdownDescription: "name of the cluster of chassis", Computed: true}, "hw_type": schema.StringAttribute{MarkdownDescription: "the hardware type for the card", Computed: true}, "licenses": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"end_date": schema.Int64Attribute{MarkdownDescription: "The license end date code in format YYYYMMDD", Computed: true}, "features": schema.StringAttribute{MarkdownDescription: "Comma-separated list of licensed application names", Computed: true}, "floated": schema.BoolAttribute{MarkdownDescription: "true if the license was assigned using current Fabric Manager from a floating license pool, false otherwise", Computed: true}, "grace_period": schema.Int64Attribute{MarkdownDescription: "number of days after the end date of the license for which the user can still use the license as if it was still valid", Computed: true}, "license_status": schema.StringAttribute{MarkdownDescription: "status of the license", Computed: true}, "license_type": schema.StringAttribute{MarkdownDescription: "type of the license", Computed: true}, "notif_period": schema.Int64Attribute{MarkdownDescription: "number of days before the end date of the license when notification is sent out about impending license expiry", Computed: true}, "start_date": schema.Int64Attribute{MarkdownDescription: "The license start date code in format YYYYMMDD", Computed: true}}}}, "serial_number": schema.StringAttribute{Computed: true}, "slot_id": schema.StringAttribute{MarkdownDescription: "chassis slot number where the card is inserted", Computed: true}}}}, "cluster_name": schema.StringAttribute{Computed: true}, "host_name": schema.StringAttribute{Computed: true}, "hw_type": schema.StringAttribute{Computed: true}, "licenses": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"end_date": schema.Int64Attribute{MarkdownDescription: "The license end date code in format YYYYMMDD", Computed: true}, "features": schema.StringAttribute{MarkdownDescription: "Comma-separated list of licensed application names", Computed: true}, "floated": schema.BoolAttribute{MarkdownDescription: "true if the license was assigned using current Fabric Manager from a floating license pool, false otherwise", Computed: true}, "grace_period": schema.Int64Attribute{MarkdownDescription: "number of days after the end date of the license for which the user can still use the license as if it was still valid", Computed: true}, "license_status": schema.StringAttribute{MarkdownDescription: "status of the license", Computed: true}, "license_type": schema.StringAttribute{MarkdownDescription: "type of the license", Computed: true}, "notif_period": schema.Int64Attribute{MarkdownDescription: "number of days before the end date of the license when notification is sent out about impending license expiry", Computed: true}, "start_date": schema.Int64Attribute{MarkdownDescription: "The license start date code in format YYYYMMDD", Computed: true}}}}, "model": schema.StringAttribute{MarkdownDescription: "Gigamon physical device models", Computed: true}, "node_id": schema.StringAttribute{Computed: true}, "serial_number": schema.StringAttribute{Computed: true}}}}, "cluster_name": schema.StringAttribute{Computed: true}, "hash_code": schema.StringAttribute{Computed: true}}}}}}
}

// Read fetches remote state into the data source model.
func (d *GetAllClusterLicensesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetAllClusterLicensesDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	d.readListRemote(ctx, &config, resp)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}

// readListRemote performs the paginated read HTTP exchange and decodes the response array into config. Extracted from Read so the request/response logic is unit-testable without a tfsdk.Config.
func (d *GetAllClusterLicensesDataSource) readListRemote(ctx context.Context, config *GetAllClusterLicensesDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/licensing/module/all"
	params := url.Values{}
	var nextURL string
	fetch := func(ctx context.Context, p url.Values) (*http.Response, error) {
		httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
		if err != nil {
			return nil, err
		}
		if nextURL != "" {
			parsed, perr := url.Parse(nextURL)
			if perr != nil {
				return nil, perr
			}
			httpReq.URL = parsed
		} else {
			httpReq.URL.RawQuery = p.Encode()
		}
		return d.client.Do(httpReq)
	}
	pages, err := client.ListAllPages(ctx, params, fetch, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_licenses", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_licenses", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["clusterLicensesList"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_licenses", fmt.Sprintf("Could not decode list page: missing %q array", "clusterLicensesList"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_licenses", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetAllClusterLicensesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Data Source Configure Type", fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData))
		return
	}
	d.client = c
}
