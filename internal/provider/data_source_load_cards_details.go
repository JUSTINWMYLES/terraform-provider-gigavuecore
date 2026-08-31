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
	_ datasource.DataSource              = (*LoadCardsDetailsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadCardsDetailsDataSource)(nil)
)

// LoadCardsDetailsDataSource is the generated Terraform data source implementation.
type LoadCardsDetailsDataSource struct {
	client *client.Client
}

// LoadCardsDetailsDataSourceModel describes the data source state shape.
type LoadCardsDetailsDataSourceModel struct {
	ClusterId types.String `tfsdk:"cluster_id" json:"clusterId"`
	Items     types.List   `tfsdk:"items"`
	NodeId    types.String `tfsdk:"node_id" json:"nodeId"`
	Page      types.String `tfsdk:"page"`
	Sort      types.String `tfsdk:"sort"`
}

// NewLoadCardsDetailsDataSource returns a new instance of the generated data source.
func NewLoadCardsDetailsDataSource() datasource.DataSource {
	return &LoadCardsDetailsDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadCardsDetailsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_cards_details"
}

// Schema returns the data source schema.
func (d *LoadCardsDetailsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load all Device Cards details", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID. Either 'clusterId' or 'nodeId' is required", Optional: true}, "items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"admin_status": schema.StringAttribute{Computed: true}, "alarm_buffer_threshold": schema.Int64Attribute{MarkdownDescription: "card micro burst threshold", Computed: true}, "config_status": schema.StringAttribute{Computed: true}, "fabric_hash_adv": schema.BoolAttribute{MarkdownDescription: "Advanced Fabric Hash. Supported only for Q02X32/Q08 cards", Computed: true}, "health_state": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "health_state_reasons": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"message": schema.StringAttribute{MarkdownDescription: "Read-only. Describes the reason for component's health state", Computed: true}, "severity": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "traffic_health_state_computation_type": schema.StringAttribute{MarkdownDescription: "Traffic Health State Computation Type", Computed: true}}}}, "hw_revision": schema.StringAttribute{Computed: true}, "hw_type": schema.StringAttribute{Computed: true}, "mode": schema.StringAttribute{Computed: true}, "oper_status": schema.StringAttribute{Computed: true}, "pld_info": schema.SingleNestedAttribute{MarkdownDescription: "only applicable to HC3 chassis", Computed: true, Attributes: map[string]schema.Attribute{"need_upgrade": schema.BoolAttribute{Computed: true}, "pld_revision": schema.StringAttribute{MarkdownDescription: "PLD revision", Computed: true}}}, "power_priority": schema.Int64Attribute{MarkdownDescription: "Power slot Priority", Computed: true}, "power_req": schema.Int64Attribute{MarkdownDescription: "watt", Computed: true}, "product_code": schema.StringAttribute{Computed: true}, "serial_number": schema.StringAttribute{Computed: true}, "slot_id": schema.StringAttribute{MarkdownDescription: "device card slot id", Computed: true}, "temperatures": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"board": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "bottom_switch": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "bottom_switch_major": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "bottom_switch_minor": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "bottom_switch_shut": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "cav_cpu": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "cpu": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "cpu_major": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "cpu_minor": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "cpu_shut": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "e1_cpu": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "e1_port": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "e2_cpu": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "exhaust": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "intake": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "near_cpu": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "near_cpu_major": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "near_cpu_minor": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "near_cpu_shut": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "qsfp_cage": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "qsfp_cage_major": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "qsfp_cage_minor": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "qsfp_cage_shut": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "rear_panel": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "rear_panel_major": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "rear_panel_minor": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "rear_panel_shut": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "sfp_cage": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "sfp_cage_major": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "sfp_cage_minor": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "sfp_cage_shut": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "switch": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "switch_major": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "switch_minor": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "switch_shut": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "time_stamp": schema.StringAttribute{MarkdownDescription: "date-time of stats collection in RFC 3339 format", Computed: true}, "top_switch": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "top_switch_major": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "top_switch_minor": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "top_switch_shut": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}}}, "voltages": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"time_stamp": schema.StringAttribute{MarkdownDescription: "date-time of stats collection in RFC 3339 format", Computed: true}, "value": schema.Float64Attribute{MarkdownDescription: "volt", Computed: true}, "voltage": schema.StringAttribute{Computed: true}}}}}}}, "node_id": schema.StringAttribute{MarkdownDescription: "ID of the target device. Either 'clusterId' or 'nodeId' is required", Optional: true}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned", Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadCardsDetailsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadCardsDetailsDataSourceModel
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
func (d *LoadCardsDetailsDataSource) readListRemote(ctx context.Context, config *LoadCardsDetailsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/inventory/chassis/cards"
	params := url.Values{}
	if !config.ClusterId.IsNull() {
		params.Set("clusterId", config.ClusterId.ValueString())
	}
	if !config.NodeId.IsNull() {
		params.Set("nodeId", config.NodeId.ValueString())
	}
	if !config.Page.IsNull() {
		params.Set("page", config.Page.ValueString())
	}
	if !config.Sort.IsNull() {
		params.Set("sort", config.Sort.ValueString())
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
		resp.Diagnostics.AddError("Error reading gigavuecore_load_cards_details", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_cards_details", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["cards"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_cards_details", fmt.Sprintf("Could not decode list page: missing %q array", "cards"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_cards_details", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadCardsDetailsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
