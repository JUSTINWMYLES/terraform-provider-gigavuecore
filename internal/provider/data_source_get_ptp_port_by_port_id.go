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
	_ datasource.DataSource              = (*GetPtpPortByPortIdDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetPtpPortByPortIdDataSource)(nil)
)

// GetPtpPortByPortIdDataSource is the generated Terraform data source implementation.
type GetPtpPortByPortIdDataSource struct {
	client *client.Client
}

// GetPtpPortByPortIdDataSourceModel describes the data source state shape.
type GetPtpPortByPortIdDataSourceModel struct {
	AnnounceInterval         types.Int64  `tfsdk:"announce_interval" json:"announceInterval"`
	AnnounceReceiptTimeout   types.Int64  `tfsdk:"announce_receipt_timeout" json:"announceReceiptTimeout"`
	ClockIdentity            types.String `tfsdk:"clock_identity" json:"clockIdentity"`
	ClockPortId              types.Int64  `tfsdk:"clock_port_id" json:"clockPortId"`
	ClockPortState           types.String `tfsdk:"clock_port_state" json:"clockPortState"`
	ClockPortStateAlias      types.String `tfsdk:"clock_port_state_alias" json:"clockPortStateAlias"`
	DelayMechanism           types.String `tfsdk:"delay_mechanism" json:"delayMechanism"`
	DelayRequestInterval     types.Int64  `tfsdk:"delay_request_interval" json:"delayRequestInterval"`
	Enabled                  types.Bool   `tfsdk:"enabled"`
	LocalPriority            types.Int64  `tfsdk:"local_priority" json:"localPriority"`
	PeerDelayRequestInterval types.Int64  `tfsdk:"peer_delay_request_interval" json:"peerDelayRequestInterval"`
	PeerMeanPathDelay        types.Int64  `tfsdk:"peer_mean_path_delay" json:"peerMeanPathDelay"`
	PortId                   types.String `tfsdk:"port_id" json:"portId"`
	PtpVersion               types.String `tfsdk:"ptp_version" json:"ptpVersion"`
	SyncInterval             types.Int64  `tfsdk:"sync_interval" json:"syncInterval"`
	VlanId                   types.Int64  `tfsdk:"vlan_id" json:"vlanId"`
}

// NewGetPtpPortByPortIdDataSource returns a new instance of the generated data source.
func NewGetPtpPortByPortIdDataSource() datasource.DataSource {
	return &GetPtpPortByPortIdDataSource{}
}

// Metadata returns the data source type name.
func (d *GetPtpPortByPortIdDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_ptp_port_by_port_id"
}

// Schema returns the data source schema.
func (d *GetPtpPortByPortIdDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Reads the get ptp port by port id data source.", Attributes: map[string]schema.Attribute{"announce_interval": schema.Int64Attribute{MarkdownDescription: "Logarithm to the base 2 of the mean announce Interval", Computed: true}, "announce_receipt_timeout": schema.Int64Attribute{MarkdownDescription: "Specify the number of announceInterval that has to pass without receipt of an Announce message before the occurrence of the expire event", Computed: true}, "clock_identity": schema.StringAttribute{MarkdownDescription: "Clock Identifier as per the IEEE 1588 standard which is represented by an integer 8-octet array", Computed: true}, "clock_port_id": schema.Int64Attribute{MarkdownDescription: "The value of the portNumber for a port on a PTP node supporting a single PTP port shall be 1. The values of the port numbers for the N ports on a PTP node supporting N PTP ports shall be 1, 2, ...N, respectively. The all-zeros and all-ones portNumber values are reserved.", Computed: true}, "clock_port_state": schema.StringAttribute{MarkdownDescription: "Value of the current state of the protocol engine associated with the given PTP port. (deprecated: use clockPortStateAlias)", Computed: true}, "clock_port_state_alias": schema.StringAttribute{MarkdownDescription: "Value of the current state of the protocol engine associated with the given PTP port", Computed: true}, "delay_mechanism": schema.StringAttribute{MarkdownDescription: "The propagation delay measuring option used by the port in computing mean path delay", Computed: true}, "delay_request_interval": schema.Int64Attribute{MarkdownDescription: "Logarithm to the base 2 of the permitted mean time interval between successive Delay_req messages", Computed: true}, "enabled": schema.BoolAttribute{Computed: true}, "local_priority": schema.Int64Attribute{Computed: true}, "peer_delay_request_interval": schema.Int64Attribute{MarkdownDescription: "Logarithm to the base 2 of the delayReqInterval", Computed: true}, "peer_mean_path_delay": schema.Int64Attribute{MarkdownDescription: "If the value of the delay Mechanism member is peer-to-peer (P2P), the value of peerMeanPathDelay shall be an estimate of the current one-way propagation delay on the link, i.e.,meanPathDelay, attached to this port computed using the peer delay mechanism", Computed: true}, "port_id": schema.StringAttribute{MarkdownDescription: "Port ID in [box/slot/portid] format", Required: true}, "ptp_version": schema.StringAttribute{MarkdownDescription: "Indicates the version of the PTP standard implemented on the port. IEEE Std 1588-2008 corresponds to PTP version 2 whereas 1588-2002 is PTP version 1. For this implementation, PTPv2 will be supported", Computed: true}, "sync_interval": schema.Int64Attribute{MarkdownDescription: "Logarithm to the base 2 of the mean SyncInterval for multicast messages", Computed: true}, "vlan_id": schema.Int64Attribute{Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetPtpPortByPortIdDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetPtpPortByPortIdDataSourceModel
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
func (d *GetPtpPortByPortIdDataSource) readRemote(ctx context.Context, config *GetPtpPortByPortIdDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/ptp/portState/{portId}"
	reqPath = strings.ReplaceAll(reqPath, "{portId}", url.PathEscape(config.PortId.ValueString()))
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_port_by_port_id", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_port_by_port_id", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_port_by_port_id", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_port_by_port_id", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_port_by_port_id", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_port_by_port_id", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_port_by_port_id", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_port_by_port_id", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_port_by_port_id", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_port_by_port_id", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_port_by_port_id", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_port_by_port_id", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_port_by_port_id", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["portState"]; ok {
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_port_by_port_id", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetPtpPortByPortIdDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
