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
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*GetPtpCountersByAliasDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetPtpCountersByAliasDataSource)(nil)
)

// GetPtpCountersByAliasDataSource is the generated Terraform data source implementation.
type GetPtpCountersByAliasDataSource struct {
	client *client.Client
}

// GetPtpCountersByAliasDataSourceModel describes the data source state shape.
type GetPtpCountersByAliasDataSourceModel struct {
	Alias                  types.String `tfsdk:"alias"`
	BoxId                  types.String `tfsdk:"box_id" json:"boxId"`
	ClusterId              types.String `tfsdk:"cluster_id" json:"clusterId"`
	DiscardedPackets       types.Int64  `tfsdk:"discarded_packets" json:"discardedPackets"`
	Ipv4PtpRxPackets       types.Int64  `tfsdk:"ipv4_ptp_rx_packets" json:"Ipv4PtpRxPackets"`
	Ipv6PtpRxPackets       types.Int64  `tfsdk:"ipv6_ptp_rx_packets" json:"Ipv6PtpRxPackets"`
	L2PtpRxPackets         types.Int64  `tfsdk:"l2_ptp_rx_packets" json:"L2PtpRxPackets"`
	QueueOverflowRxPackets types.Int64  `tfsdk:"queue_overflow_rx_packets" json:"QueueOverflowRxPackets"`
	RcpuEncapRxPackets     types.Int64  `tfsdk:"rcpu_encap_rx_packets" json:"RcpuEncapRxPackets"`
	RxPackets              types.Int64  `tfsdk:"rx_packets" json:"rxPackets"`
	TxPackets              types.Int64  `tfsdk:"tx_packets" json:"txPackets"`
	UdpPtpRxPackets        types.Int64  `tfsdk:"udp_ptp_rx_packets" json:"UdpPtpRxPackets"`
}

// NewGetPtpCountersByAliasDataSource returns a new instance of the generated data source.
func NewGetPtpCountersByAliasDataSource() datasource.DataSource {
	return &GetPtpCountersByAliasDataSource{}
}

// Metadata returns the data source type name.
func (d *GetPtpCountersByAliasDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_ptp_counters_by_alias"
}

// Schema returns the data source schema.
func (d *GetPtpCountersByAliasDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Reads the get ptp counters by alias data source.", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "alias of the time stamping PTP configuration", Required: true}, "box_id": schema.StringAttribute{MarkdownDescription: "specify the cluster node by boxId. By default all nodes are selected.", Computed: true, Optional: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "discarded_packets": schema.Int64Attribute{Computed: true}, "ipv4_ptp_rx_packets": schema.Int64Attribute{Computed: true}, "ipv6_ptp_rx_packets": schema.Int64Attribute{Computed: true}, "l2_ptp_rx_packets": schema.Int64Attribute{Computed: true}, "queue_overflow_rx_packets": schema.Int64Attribute{Computed: true}, "rcpu_encap_rx_packets": schema.Int64Attribute{Computed: true}, "rx_packets": schema.Int64Attribute{Computed: true}, "tx_packets": schema.Int64Attribute{Computed: true}, "udp_ptp_rx_packets": schema.Int64Attribute{Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetPtpCountersByAliasDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetPtpCountersByAliasDataSourceModel
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
func (d *GetPtpCountersByAliasDataSource) readRemote(ctx context.Context, config *GetPtpCountersByAliasDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/nodeCounters/ptp/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(config.Alias.ValueString()))
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_counters_by_alias", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	if !config.BoxId.IsNull() {
		query.Set("boxId", config.BoxId.ValueString())
	}
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_counters_by_alias", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_counters_by_alias", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_counters_by_alias", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_counters_by_alias", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_counters_by_alias", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_counters_by_alias", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_counters_by_alias", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_counters_by_alias", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_counters_by_alias", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_counters_by_alias", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_counters_by_alias", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_counters_by_alias", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["counters"]; ok {
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_counters_by_alias", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetPtpCountersByAliasDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
