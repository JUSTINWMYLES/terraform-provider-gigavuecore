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
	_ datasource.DataSource              = (*GetTunnelLogicalGroupStatsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetTunnelLogicalGroupStatsDataSource)(nil)
)

// GetTunnelLogicalGroupStatsDataSource is the generated Terraform data source implementation.
type GetTunnelLogicalGroupStatsDataSource struct {
	client *client.Client
}

// GetTunnelLogicalGroupStatsDataSourceModel describes the data source state shape.
type GetTunnelLogicalGroupStatsDataSourceModel struct {
	Alias                      types.String  `tfsdk:"alias"`
	Deviation                  types.Dynamic `tfsdk:"deviation"`
	LastSuccessfulAttemptTime  types.String  `tfsdk:"last_successful_attempt_time" json:"lastSuccessfulAttemptTime"`
	OctetsRx                   types.Dynamic `tfsdk:"octets_rx" json:"octetsRx"`
	OctetsTx                   types.Dynamic `tfsdk:"octets_tx" json:"octetsTx"`
	PacketsRx                  types.Dynamic `tfsdk:"packets_rx" json:"packetsRx"`
	PacketsTx                  types.Dynamic `tfsdk:"packets_tx" json:"packetsTx"`
	ProcessedTrafficPercentage types.Dynamic `tfsdk:"processed_traffic_percentage" json:"processedTrafficPercentage"`
}

// NewGetTunnelLogicalGroupStatsDataSource returns a new instance of the generated data source.
func NewGetTunnelLogicalGroupStatsDataSource() datasource.DataSource {
	return &GetTunnelLogicalGroupStatsDataSource{}
}

// Metadata returns the data source type name.
func (d *GetTunnelLogicalGroupStatsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_tunnel_logical_group_stats"
}

// Schema returns the data source schema.
func (d *GetTunnelLogicalGroupStatsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get Stats for Tunnel Logical Groups", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Tunnel Logical Group Alias", Required: true}, "deviation": schema.DynamicAttribute{MarkdownDescription: "Traffic Deviation in Percentage", Computed: true}, "last_successful_attempt_time": schema.StringAttribute{MarkdownDescription: "Last Stats Polled Timestamp UTC Format", Computed: true}, "octets_rx": schema.DynamicAttribute{MarkdownDescription: "Decapsulation end Octets Tx", Computed: true}, "octets_tx": schema.DynamicAttribute{MarkdownDescription: "Encapsulation end Octets Tx", Computed: true}, "packets_rx": schema.DynamicAttribute{MarkdownDescription: "Decapsulation end Packets Tx", Computed: true}, "packets_tx": schema.DynamicAttribute{MarkdownDescription: "Encapsulation end Packets Tx", Computed: true}, "processed_traffic_percentage": schema.DynamicAttribute{MarkdownDescription: "Amount of Traffic Processed based on Encap Packets Tx and Decap Packets Rx", Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetTunnelLogicalGroupStatsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetTunnelLogicalGroupStatsDataSourceModel
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
func (d *GetTunnelLogicalGroupStatsDataSource) readRemote(ctx context.Context, config *GetTunnelLogicalGroupStatsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/tunnels/logicalgroups/stats"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_tunnel_logical_group_stats", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("alias", config.Alias.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_tunnel_logical_group_stats", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_tunnel_logical_group_stats", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_tunnel_logical_group_stats", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_tunnel_logical_group_stats", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_tunnel_logical_group_stats", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_tunnel_logical_group_stats", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_tunnel_logical_group_stats", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_tunnel_logical_group_stats", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_tunnel_logical_group_stats", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_tunnel_logical_group_stats", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_tunnel_logical_group_stats", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_tunnel_logical_group_stats", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_tunnel_logical_group_stats", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetTunnelLogicalGroupStatsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
