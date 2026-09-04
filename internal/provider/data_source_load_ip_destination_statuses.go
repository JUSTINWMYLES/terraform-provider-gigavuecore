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
	_ datasource.DataSource              = (*LoadIpDestinationStatusesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadIpDestinationStatusesDataSource)(nil)
)

// LoadIpDestinationStatusesDataSource is the generated Terraform data source implementation.
type LoadIpDestinationStatusesDataSource struct {
	client *client.Client
}

// LoadIpDestinationStatusesDataSourceModel describes the data source state shape.
type LoadIpDestinationStatusesDataSourceModel struct {
	ClusterId        types.String `tfsdk:"cluster_id" json:"clusterId"`
	GsGroupAlias     types.String `tfsdk:"gs_group_alias" json:"gsGroupAlias"`
	IpInterfaceAlias types.String `tfsdk:"ip_interface_alias" json:"ipInterfaceAlias"`
	IpType           types.String `tfsdk:"ip_type" json:"ipType"`
	Items            types.List   `tfsdk:"items"`
}

// NewLoadIpDestinationStatusesDataSource returns a new instance of the generated data source.
func NewLoadIpDestinationStatusesDataSource() datasource.DataSource {
	return &LoadIpDestinationStatusesDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadIpDestinationStatusesDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_ip_destination_statuses"
}

// Schema returns the data source schema.
func (d *LoadIpDestinationStatusesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load IP destination status of all or those matching the IP Interface alias or IP type or GsGroup", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "gs_group_alias": schema.StringAttribute{MarkdownDescription: "GsGroup alias", Optional: true}, "ip_interface_alias": schema.StringAttribute{MarkdownDescription: "IP Interface alias", Optional: true}, "ip_type": schema.StringAttribute{MarkdownDescription: "IP type", Optional: true}, "items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"destination": schema.StringAttribute{MarkdownDescription: "IP Destination", Computed: true}, "gs_group": schema.StringAttribute{MarkdownDescription: "Associated GsGroup", Computed: true}, "interface_alias": schema.StringAttribute{MarkdownDescription: "Interface Alias", Computed: true}, "status": schema.StringAttribute{MarkdownDescription: "Status", Computed: true}, "te_id": schema.StringAttribute{MarkdownDescription: "Te-ID", Computed: true}}}}}}
}

// Read fetches remote state into the data source model.
func (d *LoadIpDestinationStatusesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadIpDestinationStatusesDataSourceModel
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
func (d *LoadIpDestinationStatusesDataSource) readListRemote(ctx context.Context, config *LoadIpDestinationStatusesDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/ip/destination"
	params := url.Values{}
	params.Set("clusterId", config.ClusterId.ValueString())
	if !config.IpInterfaceAlias.IsNull() {
		params.Set("ipInterfaceAlias", config.IpInterfaceAlias.ValueString())
	}
	if !config.GsGroupAlias.IsNull() {
		params.Set("gsGroupAlias", config.GsGroupAlias.ValueString())
	}
	if !config.IpType.IsNull() {
		params.Set("ipType", config.IpType.ValueString())
	}
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
		resp.Diagnostics.AddError("Error reading gigavuecore_load_ip_destination_statuses", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_ip_destination_statuses", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["ipInterfaces"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_ip_destination_statuses", fmt.Sprintf("Could not decode list page: missing %q array", "ipInterfaces"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_ip_destination_statuses", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadIpDestinationStatusesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
