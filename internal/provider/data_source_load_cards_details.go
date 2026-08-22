package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
	Cards     types.List   `tfsdk:"cards"`
	ClusterId types.String `tfsdk:"cluster_id" json:"clusterId"`
	Context   types.Object `tfsdk:"context"`
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
	resp.Schema = schema.Schema{MarkdownDescription: "Load all Device Cards details", Attributes: map[string]schema.Attribute{"cards": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"admin_status": schema.StringAttribute{Computed: true}, "alarm_buffer_threshold": schema.Int64Attribute{MarkdownDescription: "card micro burst threshold", Computed: true}, "config_status": schema.StringAttribute{Computed: true}, "fabric_hash_adv": schema.BoolAttribute{MarkdownDescription: "Advanced Fabric Hash. Supported only for Q02X32/Q08 cards", Computed: true}, "health_state": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "health_state_reasons": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"message": schema.StringAttribute{MarkdownDescription: "Read-only. Describes the reason for component's health state", Computed: true}, "severity": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "traffic_health_state_computation_type": schema.StringAttribute{MarkdownDescription: "Traffic Health State Computation Type", Computed: true}}}}, "hw_revision": schema.StringAttribute{Computed: true}, "hw_type": schema.StringAttribute{Computed: true}, "mode": schema.StringAttribute{Computed: true}, "oper_status": schema.StringAttribute{Computed: true}, "pld_info": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"need_upgrade": schema.BoolAttribute{Computed: true}, "pld_revision": schema.StringAttribute{MarkdownDescription: "PLD revision", Computed: true}}}, "power_priority": schema.Int64Attribute{MarkdownDescription: "Power slot Priority", Computed: true}, "power_req": schema.Int64Attribute{MarkdownDescription: "watt", Computed: true}, "product_code": schema.StringAttribute{Computed: true}, "serial_number": schema.StringAttribute{Computed: true}, "slot_id": schema.StringAttribute{MarkdownDescription: "device card slot id", Computed: true}, "temperatures": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"board": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "bottom_switch": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "bottom_switch_major": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "bottom_switch_minor": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "bottom_switch_shut": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "cav_cpu": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "cpu": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "cpu_major": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "cpu_minor": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "cpu_shut": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "e1_cpu": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "e1_port": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "e2_cpu": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "exhaust": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "intake": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "near_cpu": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "near_cpu_major": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "near_cpu_minor": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "near_cpu_shut": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "qsfp_cage": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "qsfp_cage_major": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "qsfp_cage_minor": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "qsfp_cage_shut": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "rear_panel": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "rear_panel_major": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "rear_panel_minor": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "rear_panel_shut": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "sfp_cage": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "sfp_cage_major": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "sfp_cage_minor": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "sfp_cage_shut": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "switch": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "switch_major": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "switch_minor": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "switch_shut": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "time_stamp": schema.StringAttribute{MarkdownDescription: "date-time of stats collection in RFC 3339 format", Computed: true}, "top_switch": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "top_switch_major": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "top_switch_minor": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "top_switch_shut": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}}}, "voltages": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"time_stamp": schema.StringAttribute{MarkdownDescription: "date-time of stats collection in RFC 3339 format", Computed: true}, "value": schema.Float64Attribute{MarkdownDescription: "volt", Computed: true}, "voltage": schema.StringAttribute{Computed: true}}}}}}}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID. Either 'clusterId' or 'nodeId' is required", Optional: true}, "context": schema.SingleNestedAttribute{MarkdownDescription: "Gigamon query result context", Computed: true, Attributes: map[string]schema.Attribute{"page_no": schema.Int64Attribute{MarkdownDescription: "page number of the returned result set", Computed: true}, "page_size": schema.Int64Attribute{MarkdownDescription: "page size of the returned result set", Computed: true}, "sort": schema.ListAttribute{MarkdownDescription: "sorting info of the returned result set. list of fields in the array indicate sorting order", Computed: true, ElementType: types.StringType}, "total_items": schema.Int64Attribute{MarkdownDescription: "total number of items in the queried entity type", Computed: true}}}, "node_id": schema.StringAttribute{MarkdownDescription: "ID of the target device. Either 'clusterId' or 'nodeId' is required", Optional: true}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned", Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadCardsDetailsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadCardsDetailsDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	d.readRemote(ctx, &config, resp)
	if resp.Diagnostics.HasError() {
		return
	}
	if config.Page.IsNull() {
		resp.Diagnostics.AddWarning("Single-page result", fmt.Sprintf("This data source reads a single page of a paginated API endpoint and does not aggregate results across pages. The \"page\" argument is unset, so the default page is returned; set it to retrieve a different page."))
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}

// readRemote performs the read HTTP exchange and decodes the response into config. Extracted from Read so the request/response logic is unit-testable without a tfsdk.Config.
func (d *LoadCardsDetailsDataSource) readRemote(ctx context.Context, config *LoadCardsDetailsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/inventory/chassis/cards"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_cards_details", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !config.ClusterId.IsNull() {
		query.Set("clusterId", config.ClusterId.ValueString())
	}
	if !config.NodeId.IsNull() {
		query.Set("nodeId", config.NodeId.ValueString())
	}
	if !config.Page.IsNull() {
		query.Set("page", config.Page.ValueString())
	}
	if !config.Sort.IsNull() {
		query.Set("sort", config.Sort.ValueString())
	}
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_cards_details", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_cards_details", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_cards_details", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_cards_details", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_cards_details", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_cards_details", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_cards_details", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_cards_details", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_cards_details", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_load_cards_details", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_load_cards_details", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_cards_details", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
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
