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
	_ datasource.DataSource              = (*GetFlowSipSummaryDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetFlowSipSummaryDataSource)(nil)
)

// GetFlowSipSummaryDataSource is the generated Terraform data source implementation.
type GetFlowSipSummaryDataSource struct {
	client *client.Client
}

// GetFlowSipSummaryDataSourceModel describes the data source state shape.
type GetFlowSipSummaryDataSourceModel struct {
	Alias              types.String `tfsdk:"alias"`
	CallerIdPattern    types.String `tfsdk:"caller_id_pattern" json:"callerIdPattern"`
	ClusterId          types.String `tfsdk:"cluster_id" json:"clusterId"`
	Gsgroup            types.String `tfsdk:"gsgroup"`
	RtpResourceSummary types.Object `tfsdk:"rtp_resource_summary" json:"rtpResourceSummary"`
	RtpSessions        types.Int64  `tfsdk:"rtp_sessions" json:"rtpSessions"`
	SipMessagesStats   types.List   `tfsdk:"sip_messages_stats" json:"sipMessagesStats"`
	SipResourceSummary types.Object `tfsdk:"sip_resource_summary" json:"sipResourceSummary"`
	SipSessions        types.Int64  `tfsdk:"sip_sessions" json:"sipSessions"`
}

// NewGetFlowSipSummaryDataSource returns a new instance of the generated data source.
func NewGetFlowSipSummaryDataSource() datasource.DataSource {
	return &GetFlowSipSummaryDataSource{}
}

// Metadata returns the data source type name.
func (d *GetFlowSipSummaryDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_flow_sip_summary"
}

// Schema returns the data source schema.
func (d *GetFlowSipSummaryDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load Flow Sip Report Summary", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "alias of the target GS Group", Required: true}, "caller_id_pattern": schema.StringAttribute{MarkdownDescription: "callerid pattern based active flows", Optional: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "gsgroup": schema.StringAttribute{MarkdownDescription: "alias of gsgroup", Computed: true}, "rtp_resource_summary": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"data_pools_avail": schema.Int64Attribute{Computed: true}, "data_pools_in_use": schema.Int64Attribute{Computed: true}}}, "rtp_sessions": schema.Int64Attribute{Computed: true}, "sip_messages_stats": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"drop": schema.Int64Attribute{Computed: true}, "no_match": schema.Int64Attribute{Computed: true}, "no_rule": schema.Int64Attribute{Computed: true}, "no_session": schema.Int64Attribute{Computed: true}, "other": schema.Int64Attribute{Computed: true}, "sip_message": schema.StringAttribute{Computed: true}, "tool_pass": schema.Int64Attribute{Computed: true}}}}, "sip_resource_summary": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"num_sessions": schema.Int64Attribute{Computed: true}, "session_avail": schema.Int64Attribute{Computed: true}, "sip_parse_errors": schema.Int64Attribute{Computed: true}}}, "sip_sessions": schema.Int64Attribute{Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetFlowSipSummaryDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetFlowSipSummaryDataSourceModel
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
func (d *GetFlowSipSummaryDataSource) readRemote(ctx context.Context, config *GetFlowSipSummaryDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/gsGroups/{alias}/flowOpsReport/flowSip/summary"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(config.Alias.ValueString()))
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_flow_sip_summary", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	if !config.CallerIdPattern.IsNull() {
		query.Set("callerIdPattern", config.CallerIdPattern.ValueString())
	}
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_flow_sip_summary", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_flow_sip_summary", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_flow_sip_summary", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_flow_sip_summary", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_flow_sip_summary", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_flow_sip_summary", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_flow_sip_summary", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_flow_sip_summary", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_flow_sip_summary", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_flow_sip_summary", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_flow_sip_summary", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_flow_sip_summary", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["flowSipReportSummary"]; ok {
		if m, ok := v.(map[string]any); ok {
			data = m
		}
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_flow_sip_summary", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetFlowSipSummaryDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
