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
	_ datasource.DataSource              = (*GetFlowFilteringSummaryDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetFlowFilteringSummaryDataSource)(nil)
)

// GetFlowFilteringSummaryDataSource is the generated Terraform data source implementation.
type GetFlowFilteringSummaryDataSource struct {
	client *client.Client
}

// GetFlowFilteringSummaryDataSourceModel describes the data source state shape.
type GetFlowFilteringSummaryDataSourceModel struct {
	Alias                   types.String `tfsdk:"alias"`
	ClusterId               types.String `tfsdk:"cluster_id" json:"clusterId"`
	ControlOnlySession      types.Int64  `tfsdk:"control_only_session" json:"controlOnlySession"`
	ControlTunnels          types.Int64  `tfsdk:"control_tunnels" json:"controlTunnels"`
	ControlUserTunnels      types.Int64  `tfsdk:"control_user_tunnels" json:"controlUserTunnels"`
	Gsgroup                 types.String `tfsdk:"gsgroup"`
	GtpCorelationStatistics types.List   `tfsdk:"gtp_corelation_statistics" json:"gtpCorelationStatistics"`
	GtpImsiPattern          types.String `tfsdk:"gtp_imsi_pattern" json:"gtpImsiPattern"`
	GtpInterfaceStatistics  types.List   `tfsdk:"gtp_interface_statistics" json:"gtpInterfaceStatistics"`
	GtpPfcpStatistics       types.Object `tfsdk:"gtp_pfcp_statistics" json:"gtpPfcpStatistics"`
	GtpSessionStatistics    types.List   `tfsdk:"gtp_session_statistics" json:"gtpSessionStatistics"`
	PendingSession          types.Int64  `tfsdk:"pending_session" json:"pendingSession"`
}

// NewGetFlowFilteringSummaryDataSource returns a new instance of the generated data source.
func NewGetFlowFilteringSummaryDataSource() datasource.DataSource {
	return &GetFlowFilteringSummaryDataSource{}
}

// Metadata returns the data source type name.
func (d *GetFlowFilteringSummaryDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_flow_filtering_summary"
}

// Schema returns the data source schema.
func (d *GetFlowFilteringSummaryDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load Flow Filtering Report Summary", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "alias of the target GS Group", Required: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "control_only_session": schema.Int64Attribute{MarkdownDescription: "number of sessions without user bearers", Computed: true}, "control_tunnels": schema.Int64Attribute{MarkdownDescription: "total number of control tunnels", Computed: true}, "control_user_tunnels": schema.Int64Attribute{MarkdownDescription: "total number of control and user tunnels", Computed: true}, "gsgroup": schema.StringAttribute{MarkdownDescription: "alias of gsgroup", Computed: true}, "gtp_corelation_statistics": schema.ListNestedAttribute{MarkdownDescription: "control messages and user data message counters", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"gtp_control_message_stats": schema.SingleNestedAttribute{MarkdownDescription: "Statistics for Control Messages GTP-C", Computed: true, Attributes: map[string]schema.Attribute{"gtp_c_stats": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"col_no_rule": schema.Int64Attribute{Computed: true}, "col_no_session": schema.Int64Attribute{Computed: true}, "col_no_tnlx": schema.Int64Attribute{Computed: true}, "col_other": schema.Int64Attribute{Computed: true}, "col_parse_er": schema.Int64Attribute{Computed: true}, "control_message": schema.StringAttribute{Computed: true}, "tool_pass": schema.Int64Attribute{Computed: true}}}}, "gtp_version": schema.StringAttribute{Computed: true}}}, "gtp_user_message_stats": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"collector": schema.Int64Attribute{Computed: true}, "drop": schema.Int64Attribute{Computed: true}, "tool_pass": schema.Int64Attribute{Computed: true}}}}}}}, "gtp_imsi_pattern": schema.StringAttribute{MarkdownDescription: "gtp imsi pattern based active flows", Optional: true}, "gtp_interface_statistics": schema.ListNestedAttribute{MarkdownDescription: "number of sessions and tunnels by interface", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"dropped_bytes": schema.Int64Attribute{Computed: true}, "dropped_pkts": schema.Int64Attribute{Computed: true}, "gtp_c_packets": schema.Int64Attribute{Computed: true}, "interface_type": schema.StringAttribute{MarkdownDescription: "interface type", Computed: true}, "rx_bytes": schema.Int64Attribute{Computed: true}, "rx_pkts": schema.Int64Attribute{Computed: true}, "tx_bytes": schema.Int64Attribute{Computed: true}, "tx_pkts": schema.Int64Attribute{Computed: true}}}}, "gtp_pfcp_statistics": schema.SingleNestedAttribute{MarkdownDescription: "PFCP message stats", Computed: true, Attributes: map[string]schema.Attribute{"asso_rel_req": schema.Int64Attribute{Computed: true}, "asso_rel_rsp": schema.Int64Attribute{Computed: true}, "asso_setup_req": schema.Int64Attribute{Computed: true}, "asso_setup_rsp": schema.Int64Attribute{Computed: true}, "asso_upd_req": schema.Int64Attribute{Computed: true}, "asso_upd_rsp": schema.Int64Attribute{Computed: true}, "bundle_msg": schema.Int64Attribute{Computed: true}, "heart_beat_req": schema.Int64Attribute{Computed: true}, "heart_beat_rsp": schema.Int64Attribute{Computed: true}, "node_rpt_req": schema.Int64Attribute{Computed: true}, "node_rpt_rsp": schema.Int64Attribute{Computed: true}, "pfd_mgmt_req": schema.Int64Attribute{Computed: true}, "pfd_mgmt_rsp": schema.Int64Attribute{Computed: true}, "res_msg_type16_to49": schema.Int64Attribute{Computed: true}, "res_msg_type58_to255": schema.Int64Attribute{Computed: true}, "sess_set_del_req": schema.Int64Attribute{Computed: true}, "sess_set_del_rsp": schema.Int64Attribute{Computed: true}, "ver_not_supp_rsp": schema.Int64Attribute{Computed: true}}}, "gtp_session_statistics": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"interface_type": schema.StringAttribute{MarkdownDescription: "interface type", Computed: true}, "sessions": schema.Int64Attribute{Computed: true}, "tunnels": schema.Int64Attribute{Computed: true}}}}, "pending_session": schema.Int64Attribute{MarkdownDescription: "number of sessions waiting for control message response", Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetFlowFilteringSummaryDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetFlowFilteringSummaryDataSourceModel
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
func (d *GetFlowFilteringSummaryDataSource) readRemote(ctx context.Context, config *GetFlowFilteringSummaryDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/gsGroups/{alias}/flowOpsReport/flowFiltering/summary"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(config.Alias.ValueString()))
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_flow_filtering_summary", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	if !config.GtpImsiPattern.IsNull() {
		query.Set("gtpImsiPattern", config.GtpImsiPattern.ValueString())
	}
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_flow_filtering_summary", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_flow_filtering_summary", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_flow_filtering_summary", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_flow_filtering_summary", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_flow_filtering_summary", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_flow_filtering_summary", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_flow_filtering_summary", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_flow_filtering_summary", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_flow_filtering_summary", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_flow_filtering_summary", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_flow_filtering_summary", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_flow_filtering_summary", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["flowFilteringReportSummary"]; ok {
		if m, ok := v.(map[string]any); ok {
			data = m
		}
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_flow_filtering_summary", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetFlowFilteringSummaryDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
