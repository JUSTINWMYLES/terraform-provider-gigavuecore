package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)
import (
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*LoadCardDetailsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadCardDetailsDataSource)(nil)
)

// LoadCardDetailsDataSource is the generated Terraform data source implementation.
type LoadCardDetailsDataSource struct {
	client *client.Client
}

// LoadCardDetailsDataSourceModel describes the data source state shape.
type LoadCardDetailsDataSourceModel struct {
	AdminStatus          types.String `tfsdk:"admin_status" json:"adminStatus"`
	AlarmBufferThreshold types.Int64  `tfsdk:"alarm_buffer_threshold" json:"alarmBufferThreshold"`
	ClusterId            types.String `tfsdk:"cluster_id" json:"clusterId"`
	ConfigStatus         types.String `tfsdk:"config_status" json:"configStatus"`
	FabricHashAdv        types.Bool   `tfsdk:"fabric_hash_adv" json:"fabricHashAdv"`
	HealthState          types.String `tfsdk:"health_state" json:"healthState"`
	HealthStateReasons   types.List   `tfsdk:"health_state_reasons" json:"healthStateReasons"`
	HwRevision           types.String `tfsdk:"hw_revision" json:"hwRevision"`
	HwType               types.String `tfsdk:"hw_type" json:"hwType"`
	Mode                 types.String `tfsdk:"mode"`
	NodeId               types.String `tfsdk:"node_id" json:"nodeId"`
	OperStatus           types.String `tfsdk:"oper_status" json:"operStatus"`
	PldInfo              types.Object `tfsdk:"pld_info" json:"pldInfo"`
	PowerPriority        types.Int64  `tfsdk:"power_priority" json:"powerPriority"`
	PowerReq             types.Int64  `tfsdk:"power_req" json:"powerReq"`
	ProductCode          types.String `tfsdk:"product_code" json:"productCode"`
	SerialNumber         types.String `tfsdk:"serial_number" json:"serialNumber"`
	SlotId               types.String `tfsdk:"slot_id" json:"slotId"`
	Temperatures         types.Object `tfsdk:"temperatures"`
	Voltages             types.List   `tfsdk:"voltages"`
}

// NewLoadCardDetailsDataSource returns a new instance of the generated data source.
func NewLoadCardDetailsDataSource() datasource.DataSource {
	return &LoadCardDetailsDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadCardDetailsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_card_details"
}

// Schema returns the data source schema.
func (d *LoadCardDetailsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load Device Card details", Attributes: map[string]schema.Attribute{"admin_status": schema.StringAttribute{Computed: true}, "alarm_buffer_threshold": schema.Int64Attribute{MarkdownDescription: "card micro burst threshold", Computed: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID. Either 'clusterId' or 'nodeId' is required", Optional: true}, "config_status": schema.StringAttribute{Computed: true}, "fabric_hash_adv": schema.BoolAttribute{MarkdownDescription: "Advanced Fabric Hash. Supported only for Q02X32/Q08 cards", Computed: true}, "health_state": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "health_state_reasons": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"message": schema.StringAttribute{MarkdownDescription: "Read-only. Describes the reason for component's health state", Computed: true}, "severity": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "traffic_health_state_computation_type": schema.StringAttribute{MarkdownDescription: "Traffic Health State Computation Type", Computed: true}}}}, "hw_revision": schema.StringAttribute{Computed: true}, "hw_type": schema.StringAttribute{Computed: true}, "mode": schema.StringAttribute{Computed: true}, "node_id": schema.StringAttribute{MarkdownDescription: "ID of the target device. Either 'clusterId' or 'nodeId' is required", Optional: true}, "oper_status": schema.StringAttribute{Computed: true}, "pld_info": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"need_upgrade": schema.BoolAttribute{Computed: true}, "pld_revision": schema.StringAttribute{MarkdownDescription: "PLD revision", Computed: true}}}, "power_priority": schema.Int64Attribute{MarkdownDescription: "Power slot Priority", Computed: true}, "power_req": schema.Int64Attribute{MarkdownDescription: "watt", Computed: true}, "product_code": schema.StringAttribute{Computed: true}, "serial_number": schema.StringAttribute{Computed: true}, "slot_id": schema.StringAttribute{MarkdownDescription: "device card slot id", Required: true}, "temperatures": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"board": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "bottom_switch": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "bottom_switch_major": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "bottom_switch_minor": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "bottom_switch_shut": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "cav_cpu": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "cpu": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "cpu_major": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "cpu_minor": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "cpu_shut": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "e1_cpu": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "e1_port": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "e2_cpu": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "exhaust": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "intake": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "near_cpu": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "near_cpu_major": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "near_cpu_minor": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "near_cpu_shut": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "qsfp_cage": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "qsfp_cage_major": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "qsfp_cage_minor": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "qsfp_cage_shut": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "rear_panel": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "rear_panel_major": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "rear_panel_minor": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "rear_panel_shut": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "sfp_cage": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "sfp_cage_major": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "sfp_cage_minor": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "sfp_cage_shut": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "switch": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "switch_major": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "switch_minor": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "switch_shut": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "time_stamp": schema.StringAttribute{MarkdownDescription: "date-time of stats collection in RFC 3339 format", Computed: true}, "top_switch": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "top_switch_major": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "top_switch_minor": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}, "top_switch_shut": schema.Float64Attribute{MarkdownDescription: "celsius", Computed: true}}}, "voltages": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"time_stamp": schema.StringAttribute{MarkdownDescription: "date-time of stats collection in RFC 3339 format", Computed: true}, "value": schema.Float64Attribute{MarkdownDescription: "volt", Computed: true}, "voltage": schema.StringAttribute{Computed: true}}}}}}
}

// Read fetches remote state into the data source model.
func (d *LoadCardDetailsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadCardDetailsDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	d.readRemote(ctx, &config, resp)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}

// readRemote performs the read HTTP exchange and decodes the response into config. Extracted from Read so the request/response logic is unit-testable without a tfsdk.Config.
func (d *LoadCardDetailsDataSource) readRemote(ctx context.Context, config *LoadCardDetailsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/inventory/chassis/cards/{slotId}"
	reqPath = strings.ReplaceAll(reqPath, "{slotId}", url.PathEscape(config.SlotId.ValueString()))
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_card_details", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !config.ClusterId.IsNull() {
		query.Set("clusterId", config.ClusterId.ValueString())
	}
	if !config.NodeId.IsNull() {
		query.Set("nodeId", config.NodeId.ValueString())
	}
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_card_details", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_card_details", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_card_details", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_card_details", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_card_details", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_card_details", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_card_details", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_card_details", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_card_details", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_load_card_details", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_load_card_details", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_card_details", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["card"]; ok {
		if m, ok := v.(map[string]any); ok {
			data = m
		} else if arr, ok := v.([]any); ok && len(arr) > 0 {
			if m, ok := arr[0].(map[string]any); ok {
				data = m
			}
		}
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_card_details", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadCardDetailsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
