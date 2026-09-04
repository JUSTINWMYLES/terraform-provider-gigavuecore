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
	_ datasource.DataSource              = (*GetAllFlowFilteringSummaryDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetAllFlowFilteringSummaryDataSource)(nil)
)

// GetAllFlowFilteringSummaryDataSource is the generated Terraform data source implementation.
type GetAllFlowFilteringSummaryDataSource struct {
	client *client.Client
}

// GetAllFlowFilteringSummaryDataSourceModel describes the data source state shape.
type GetAllFlowFilteringSummaryDataSourceModel struct {
	Items types.List `tfsdk:"items"`
}

// NewGetAllFlowFilteringSummaryDataSource returns a new instance of the generated data source.
func NewGetAllFlowFilteringSummaryDataSource() datasource.DataSource {
	return &GetAllFlowFilteringSummaryDataSource{}
}

// Metadata returns the data source type name.
func (d *GetAllFlowFilteringSummaryDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_all_flow_filtering_summary"
}

// Schema returns the data source schema.
func (d *GetAllFlowFilteringSummaryDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load all Flow Filtering Report Summary", Attributes: map[string]schema.Attribute{"items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"control_only_session": schema.Int64Attribute{MarkdownDescription: "number of sessions without user bearers", Computed: true}, "control_tunnels": schema.Int64Attribute{MarkdownDescription: "total number of control tunnels", Computed: true}, "control_user_tunnels": schema.Int64Attribute{MarkdownDescription: "total number of control and user tunnels", Computed: true}, "gsgroup": schema.StringAttribute{MarkdownDescription: "alias of gsgroup", Computed: true}, "gtp_corelation_statistics": schema.ListNestedAttribute{MarkdownDescription: "control messages and user data message counters", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"gtp_control_message_stats": schema.SingleNestedAttribute{MarkdownDescription: "Statistics for Control Messages GTP-C", Computed: true, Attributes: map[string]schema.Attribute{"gtp_c_stats": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"col_no_rule": schema.Int64Attribute{Computed: true}, "col_no_session": schema.Int64Attribute{Computed: true}, "col_no_tnlx": schema.Int64Attribute{Computed: true}, "col_other": schema.Int64Attribute{Computed: true}, "col_parse_er": schema.Int64Attribute{Computed: true}, "control_message": schema.StringAttribute{Computed: true}, "tool_pass": schema.Int64Attribute{Computed: true}}}}, "gtp_version": schema.StringAttribute{Computed: true}}}, "gtp_user_message_stats": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"collector": schema.Int64Attribute{Computed: true}, "drop": schema.Int64Attribute{Computed: true}, "tool_pass": schema.Int64Attribute{Computed: true}}}}}}}, "gtp_interface_statistics": schema.ListNestedAttribute{MarkdownDescription: "number of sessions and tunnels by interface", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"dropped_bytes": schema.Int64Attribute{Computed: true}, "dropped_pkts": schema.Int64Attribute{Computed: true}, "gtp_c_packets": schema.Int64Attribute{Computed: true}, "interface_type": schema.StringAttribute{MarkdownDescription: "interface type", Computed: true}, "rx_bytes": schema.Int64Attribute{Computed: true}, "rx_pkts": schema.Int64Attribute{Computed: true}, "tx_bytes": schema.Int64Attribute{Computed: true}, "tx_pkts": schema.Int64Attribute{Computed: true}}}}, "gtp_pfcp_statistics": schema.SingleNestedAttribute{MarkdownDescription: "PFCP message stats", Computed: true, Attributes: map[string]schema.Attribute{"asso_rel_req": schema.Int64Attribute{Computed: true}, "asso_rel_rsp": schema.Int64Attribute{Computed: true}, "asso_setup_req": schema.Int64Attribute{Computed: true}, "asso_setup_rsp": schema.Int64Attribute{Computed: true}, "asso_upd_req": schema.Int64Attribute{Computed: true}, "asso_upd_rsp": schema.Int64Attribute{Computed: true}, "bundle_msg": schema.Int64Attribute{Computed: true}, "heart_beat_req": schema.Int64Attribute{Computed: true}, "heart_beat_rsp": schema.Int64Attribute{Computed: true}, "node_rpt_req": schema.Int64Attribute{Computed: true}, "node_rpt_rsp": schema.Int64Attribute{Computed: true}, "pfd_mgmt_req": schema.Int64Attribute{Computed: true}, "pfd_mgmt_rsp": schema.Int64Attribute{Computed: true}, "res_msg_type16_to49": schema.Int64Attribute{Computed: true}, "res_msg_type58_to255": schema.Int64Attribute{Computed: true}, "sess_set_del_req": schema.Int64Attribute{Computed: true}, "sess_set_del_rsp": schema.Int64Attribute{Computed: true}, "ver_not_supp_rsp": schema.Int64Attribute{Computed: true}}}, "gtp_session_statistics": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"interface_type": schema.StringAttribute{MarkdownDescription: "interface type", Computed: true}, "sessions": schema.Int64Attribute{Computed: true}, "tunnels": schema.Int64Attribute{Computed: true}}}}, "pending_session": schema.Int64Attribute{MarkdownDescription: "number of sessions waiting for control message response", Computed: true}}}}}}
}

// Read fetches remote state into the data source model.
func (d *GetAllFlowFilteringSummaryDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetAllFlowFilteringSummaryDataSourceModel
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
func (d *GetAllFlowFilteringSummaryDataSource) readListRemote(ctx context.Context, config *GetAllFlowFilteringSummaryDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/gsGroups/flowOpsReport/flowFiltering/summary"
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_flow_filtering_summary", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_flow_filtering_summary", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["flowFilteringReportsSummary"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_flow_filtering_summary", fmt.Sprintf("Could not decode list page: missing %q array", "flowFilteringReportsSummary"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_flow_filtering_summary", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetAllFlowFilteringSummaryDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
