package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)
import (
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*GetFlowFilteringDeltaReportSummaryDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetFlowFilteringDeltaReportSummaryDataSource)(nil)
)

// GetFlowFilteringDeltaReportSummaryDataSource is the generated Terraform data source implementation.
type GetFlowFilteringDeltaReportSummaryDataSource struct {
	client *client.Client
}

// GetFlowFilteringDeltaReportSummaryDataSourceModel describes the data source state shape.
type GetFlowFilteringDeltaReportSummaryDataSourceModel struct {
	ClusterId                types.String `tfsdk:"cluster_id" json:"clusterId"`
	ControlOnlySession       types.Object `tfsdk:"control_only_session" json:"controlOnlySession"`
	ControlTunnels           types.Object `tfsdk:"control_tunnels" json:"controlTunnels"`
	ControlUserTunnels       types.Object `tfsdk:"control_user_tunnels" json:"controlUserTunnels"`
	EndTime                  types.String `tfsdk:"end_time" json:"endTime"`
	ErrorMsg                 types.String `tfsdk:"error_msg" json:"errorMsg"`
	GsGroupAlias             types.String `tfsdk:"gs_group_alias" json:"gsGroupAlias"`
	Gsgroup                  types.String `tfsdk:"gsgroup"`
	GtpCorrelationStatistics types.List   `tfsdk:"gtp_correlation_statistics" json:"gtpCorrelationStatistics"`
	GtpInterfaceStatistics   types.List   `tfsdk:"gtp_interface_statistics" json:"gtpInterfaceStatistics"`
	GtpSessionStatistics     types.List   `tfsdk:"gtp_session_statistics" json:"gtpSessionStatistics"`
	PendingSession           types.Object `tfsdk:"pending_session" json:"pendingSession"`
	ResetHappened            types.Bool   `tfsdk:"reset_happened" json:"resetHappened"`
	Since                    types.String `tfsdk:"since"`
	StartTime                types.String `tfsdk:"start_time" json:"startTime"`
}

// NewGetFlowFilteringDeltaReportSummaryDataSource returns a new instance of the generated data source.
func NewGetFlowFilteringDeltaReportSummaryDataSource() datasource.DataSource {
	return &GetFlowFilteringDeltaReportSummaryDataSource{}
}

// Metadata returns the data source type name.
func (d *GetFlowFilteringDeltaReportSummaryDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_flow_filtering_delta_report_summary"
}

// Schema returns the data source schema.
func (d *GetFlowFilteringDeltaReportSummaryDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load Flow Filtering Delta Report Summary", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "control_only_session": schema.SingleNestedAttribute{MarkdownDescription: "number of sessions without user bearers", Computed: true, Attributes: map[string]schema.Attribute{"max": schema.Int64Attribute{MarkdownDescription: "This is the latest record of data available", Computed: true}, "min": schema.Int64Attribute{MarkdownDescription: "This is the earliest record of data available", Computed: true}}}, "control_tunnels": schema.SingleNestedAttribute{MarkdownDescription: "total number of control tunnels", Computed: true, Attributes: map[string]schema.Attribute{"max": schema.Int64Attribute{MarkdownDescription: "This is the latest record of data available", Computed: true}, "min": schema.Int64Attribute{MarkdownDescription: "This is the earliest record of data available", Computed: true}}}, "control_user_tunnels": schema.SingleNestedAttribute{MarkdownDescription: "total number of control and user tunnels", Computed: true, Attributes: map[string]schema.Attribute{"max": schema.Int64Attribute{MarkdownDescription: "This is the latest record of data available", Computed: true}, "min": schema.Int64Attribute{MarkdownDescription: "This is the earliest record of data available", Computed: true}}}, "end_time": schema.StringAttribute{MarkdownDescription: "Timestamp of the latest record of data", Computed: true}, "error_msg": schema.StringAttribute{MarkdownDescription: "Description of error regarding reset/ Maintenance", Computed: true}, "gs_group_alias": schema.StringAttribute{MarkdownDescription: "alias of the target GS Group", Required: true}, "gsgroup": schema.StringAttribute{MarkdownDescription: "alias of gsgroup", Computed: true}, "gtp_correlation_statistics": schema.ListNestedAttribute{MarkdownDescription: "control messages and user data message counters", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"gtp_control_message_stats": schema.SingleNestedAttribute{MarkdownDescription: "Statistics for Control Messages GTP-C", Computed: true, Attributes: map[string]schema.Attribute{"gtp_c_stats": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"col_no_rule": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"delta": schema.Int64Attribute{MarkdownDescription: "This is the differential data calculated for the selected duration", Computed: true}, "max": schema.Int64Attribute{MarkdownDescription: "This is the latest record of data available", Computed: true}, "min": schema.Int64Attribute{MarkdownDescription: "This is the earliest record of data available", Computed: true}}}, "col_no_session": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"delta": schema.Int64Attribute{MarkdownDescription: "This is the differential data calculated for the selected duration", Computed: true}, "max": schema.Int64Attribute{MarkdownDescription: "This is the latest record of data available", Computed: true}, "min": schema.Int64Attribute{MarkdownDescription: "This is the earliest record of data available", Computed: true}}}, "col_no_tnlx": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"delta": schema.Int64Attribute{MarkdownDescription: "This is the differential data calculated for the selected duration", Computed: true}, "max": schema.Int64Attribute{MarkdownDescription: "This is the latest record of data available", Computed: true}, "min": schema.Int64Attribute{MarkdownDescription: "This is the earliest record of data available", Computed: true}}}, "col_other": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"delta": schema.Int64Attribute{MarkdownDescription: "This is the differential data calculated for the selected duration", Computed: true}, "max": schema.Int64Attribute{MarkdownDescription: "This is the latest record of data available", Computed: true}, "min": schema.Int64Attribute{MarkdownDescription: "This is the earliest record of data available", Computed: true}}}, "col_parse_er": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"delta": schema.Int64Attribute{MarkdownDescription: "This is the differential data calculated for the selected duration", Computed: true}, "max": schema.Int64Attribute{MarkdownDescription: "This is the latest record of data available", Computed: true}, "min": schema.Int64Attribute{MarkdownDescription: "This is the earliest record of data available", Computed: true}}}, "control_message": schema.StringAttribute{Computed: true}, "tool_pass": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"delta": schema.Int64Attribute{MarkdownDescription: "This is the differential data calculated for the selected duration", Computed: true}, "max": schema.Int64Attribute{MarkdownDescription: "This is the latest record of data available", Computed: true}, "min": schema.Int64Attribute{MarkdownDescription: "This is the earliest record of data available", Computed: true}}}}}}, "gtp_version": schema.StringAttribute{Computed: true}}}, "gtp_user_message_stats": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"collector": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"delta": schema.Int64Attribute{MarkdownDescription: "This is the differential data calculated for the selected duration", Computed: true}, "max": schema.Int64Attribute{MarkdownDescription: "This is the latest record of data available", Computed: true}, "min": schema.Int64Attribute{MarkdownDescription: "This is the earliest record of data available", Computed: true}}}, "drop": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"delta": schema.Int64Attribute{MarkdownDescription: "This is the differential data calculated for the selected duration", Computed: true}, "max": schema.Int64Attribute{MarkdownDescription: "This is the latest record of data available", Computed: true}, "min": schema.Int64Attribute{MarkdownDescription: "This is the earliest record of data available", Computed: true}}}, "tool_pass": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"delta": schema.Int64Attribute{MarkdownDescription: "This is the differential data calculated for the selected duration", Computed: true}, "max": schema.Int64Attribute{MarkdownDescription: "This is the latest record of data available", Computed: true}, "min": schema.Int64Attribute{MarkdownDescription: "This is the earliest record of data available", Computed: true}}}}}}}}}, "gtp_interface_statistics": schema.ListNestedAttribute{MarkdownDescription: "number of sessions and tunnels by interface", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"dropped_bytes": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"delta": schema.Int64Attribute{MarkdownDescription: "This is the differential data calculated for the selected duration", Computed: true}, "max": schema.Int64Attribute{MarkdownDescription: "This is the latest record of data available", Computed: true}, "min": schema.Int64Attribute{MarkdownDescription: "This is the earliest record of data available", Computed: true}}}, "dropped_pkts": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"delta": schema.Int64Attribute{MarkdownDescription: "This is the differential data calculated for the selected duration", Computed: true}, "max": schema.Int64Attribute{MarkdownDescription: "This is the latest record of data available", Computed: true}, "min": schema.Int64Attribute{MarkdownDescription: "This is the earliest record of data available", Computed: true}}}, "gtp_c_packets": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"delta": schema.Int64Attribute{MarkdownDescription: "This is the differential data calculated for the selected duration", Computed: true}, "max": schema.Int64Attribute{MarkdownDescription: "This is the latest record of data available", Computed: true}, "min": schema.Int64Attribute{MarkdownDescription: "This is the earliest record of data available", Computed: true}}}, "interface_type": schema.StringAttribute{MarkdownDescription: "interface type", Computed: true}, "rx_bytes": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"delta": schema.Int64Attribute{MarkdownDescription: "This is the differential data calculated for the selected duration", Computed: true}, "max": schema.Int64Attribute{MarkdownDescription: "This is the latest record of data available", Computed: true}, "min": schema.Int64Attribute{MarkdownDescription: "This is the earliest record of data available", Computed: true}}}, "rx_pkts": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"delta": schema.Int64Attribute{MarkdownDescription: "This is the differential data calculated for the selected duration", Computed: true}, "max": schema.Int64Attribute{MarkdownDescription: "This is the latest record of data available", Computed: true}, "min": schema.Int64Attribute{MarkdownDescription: "This is the earliest record of data available", Computed: true}}}, "tx_bytes": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"delta": schema.Int64Attribute{MarkdownDescription: "This is the differential data calculated for the selected duration", Computed: true}, "max": schema.Int64Attribute{MarkdownDescription: "This is the latest record of data available", Computed: true}, "min": schema.Int64Attribute{MarkdownDescription: "This is the earliest record of data available", Computed: true}}}, "tx_pkts": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"delta": schema.Int64Attribute{MarkdownDescription: "This is the differential data calculated for the selected duration", Computed: true}, "max": schema.Int64Attribute{MarkdownDescription: "This is the latest record of data available", Computed: true}, "min": schema.Int64Attribute{MarkdownDescription: "This is the earliest record of data available", Computed: true}}}}}}, "gtp_session_statistics": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"interface_type": schema.StringAttribute{MarkdownDescription: "interface type", Computed: true}, "sessions": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"max": schema.Int64Attribute{MarkdownDescription: "This is the latest record of data available", Computed: true}, "min": schema.Int64Attribute{MarkdownDescription: "This is the earliest record of data available", Computed: true}}}, "tunnels": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"max": schema.Int64Attribute{MarkdownDescription: "This is the latest record of data available", Computed: true}, "min": schema.Int64Attribute{MarkdownDescription: "This is the earliest record of data available", Computed: true}}}}}}, "pending_session": schema.SingleNestedAttribute{MarkdownDescription: "number of sessions waiting for control message response", Computed: true, Attributes: map[string]schema.Attribute{"max": schema.Int64Attribute{MarkdownDescription: "This is the latest record of data available", Computed: true}, "min": schema.Int64Attribute{MarkdownDescription: "This is the earliest record of data available", Computed: true}}}, "reset_happened": schema.BoolAttribute{MarkdownDescription: "flag to notify reset/ Maintenance happened if true", Computed: true}, "since": schema.StringAttribute{MarkdownDescription: "Start time reference in the past relative to 'now'. In the {number}-{timeUnit} format, where {timeUnit} is one of: ['minute', 'hour', 'day', 'week', 'month']. (Ex: '3-hour').", Required: true}, "start_time": schema.StringAttribute{MarkdownDescription: "Timestamp of the earliest record of data", Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetFlowFilteringDeltaReportSummaryDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetFlowFilteringDeltaReportSummaryDataSourceModel
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
func (d *GetFlowFilteringDeltaReportSummaryDataSource) readRemote(ctx context.Context, config *GetFlowFilteringDeltaReportSummaryDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/flowFiltering/deltaReport"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_flow_filtering_delta_report_summary", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	query.Set("gsGroupAlias", config.GsGroupAlias.ValueString())
	query.Set("since", config.Since.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_flow_filtering_delta_report_summary", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_flow_filtering_delta_report_summary", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_flow_filtering_delta_report_summary", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_flow_filtering_delta_report_summary", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_flow_filtering_delta_report_summary", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_flow_filtering_delta_report_summary", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_flow_filtering_delta_report_summary", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_flow_filtering_delta_report_summary", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_flow_filtering_delta_report_summary", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_flow_filtering_delta_report_summary", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_flow_filtering_delta_report_summary", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_flow_filtering_delta_report_summary", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["flowFilteringDeltaReportSummary"]; ok {
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_flow_filtering_delta_report_summary", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetFlowFilteringDeltaReportSummaryDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
