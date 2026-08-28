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
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*GetAllicapClientDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetAllicapClientDataSource)(nil)
)

// GetAllicapClientDataSource is the generated Terraform data source implementation.
type GetAllicapClientDataSource struct {
	client *client.Client
}

// GetAllicapClientDataSourceModel describes the data source state shape.
type GetAllicapClientDataSourceModel struct {
	ClusterId types.String `tfsdk:"cluster_id" json:"clusterId"`
	Items     types.List   `tfsdk:"items"`
}

// NewGetAllicapClientDataSource returns a new instance of the generated data source.
func NewGetAllicapClientDataSource() datasource.DataSource {
	return &GetAllicapClientDataSource{}
}

// Metadata returns the data source type name.
func (d *GetAllicapClientDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_allicap_client"
}

// Schema returns the data source schema.
func (d *GetAllicapClientDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get All icap client apps across FM", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "If provided, icap client only for that cluster are returned", Optional: true}, "items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Icap Alias", Computed: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "id of the defining cluster", Computed: true}, "config_status": schema.StringAttribute{MarkdownDescription: "Configuration status of this ICAP Client app", Computed: true}, "config_status_reasons": schema.ListAttribute{MarkdownDescription: "In case of configuration failure, this message provides details about the possible cause of the failure", Computed: true, ElementType: types.StringType}, "gs_engines": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "gs_grp_alias": schema.StringAttribute{MarkdownDescription: "GSGroup Alias // User provided (or) Auto-generated", Computed: true}, "gsop_alias": schema.StringAttribute{MarkdownDescription: "GSOP Alias", Computed: true}, "health_state": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "health_state_reasons": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"message": schema.StringAttribute{MarkdownDescription: "Read-only. Describes the reason for component's health state", Computed: true}, "severity": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "traffic_health_state_computation_type": schema.StringAttribute{MarkdownDescription: "Traffic Health State Computation Type", Computed: true}}}}, "icap_map_alias": schema.StringAttribute{MarkdownDescription: "Icap MAP Alias", Computed: true}, "icap_profile_config": schema.SingleNestedAttribute{MarkdownDescription: "ICAP Profile", Computed: true, Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Icap Alias", Computed: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "id of the defining cluster", Computed: true}, "exceed_action": schema.StringAttribute{MarkdownDescription: "Icap Profile action incase of Http request buffer exceeded", Computed: true}, "http_req_buf": schema.Int64Attribute{MarkdownDescription: "Icap Profile Http request buffer in KB", Computed: true}, "inactivity_timeout": schema.Int64Attribute{MarkdownDescription: "Icap Inactivity timeout in minutes", Computed: true}, "preview": schema.Int64Attribute{MarkdownDescription: "Icap Preview bytes in KB", Computed: true}, "resp_mod": schema.StringAttribute{MarkdownDescription: "Icap Response Modification Enable|Disable", Computed: true}, "resp_timeout": schema.Int64Attribute{MarkdownDescription: "Icap Server Response Timeout value in seconds", Computed: true}, "resp_timeout_action": schema.StringAttribute{MarkdownDescription: "Response Timeout action Drop|Bypass", Computed: true}, "server_group": schema.StringAttribute{MarkdownDescription: "Icap Server Group Alias", Computed: true}, "src_max_l4_port": schema.Int64Attribute{MarkdownDescription: "Icap Service Source l4 port maximum", Computed: true}, "src_min_l4_port": schema.Int64Attribute{MarkdownDescription: "Icap Service Source l4 port minimum", Computed: true}}}, "icap_server_grp_alias": schema.StringAttribute{MarkdownDescription: "Icap Server Group Alias // Auto-generated", Computed: true}, "icap_servers": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Icap Server Alias", Computed: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "id of the defining cluster", Computed: true}, "comment": schema.StringAttribute{MarkdownDescription: "Icap Server Comment", Computed: true}, "l3_address": schema.StringAttribute{MarkdownDescription: "Icap Server IP Address", Computed: true}, "l4_port": schema.Int64Attribute{MarkdownDescription: "Icap Server l4 Port", Computed: true}, "options_service_url": schema.StringAttribute{MarkdownDescription: "Options Service URL", Computed: true}, "reqmod_service_url": schema.StringAttribute{MarkdownDescription: "Request Modification Service URL", Computed: true}, "respmod_service_url": schema.StringAttribute{MarkdownDescription: "Response Modification Service URL", Computed: true}}}}, "ing_alias": schema.StringAttribute{MarkdownDescription: "Inline network group Alias //Auto-generated", Computed: true}, "inline_networks": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "ip_interface_alias": schema.StringAttribute{MarkdownDescription: "IpInterface Alias", Computed: true}}}}}}
}

// Read fetches remote state into the data source model.
func (d *GetAllicapClientDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetAllicapClientDataSourceModel
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
func (d *GetAllicapClientDataSource) readListRemote(ctx context.Context, config *GetAllicapClientDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/icap"
	params := url.Values{}
	if !config.ClusterId.IsNull() {
		params.Set("clusterId", config.ClusterId.ValueString())
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_allicap_client", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_allicap_client", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["gigaIcapClients"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_allicap_client", fmt.Sprintf("Could not decode list page: missing %q array", "gigaIcapClients"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_allicap_client", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetAllicapClientDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
