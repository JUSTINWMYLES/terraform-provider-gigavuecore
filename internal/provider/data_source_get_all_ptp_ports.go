package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)
import (
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*GetAllPtpPortsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetAllPtpPortsDataSource)(nil)
)

// GetAllPtpPortsDataSource is the generated Terraform data source implementation.
type GetAllPtpPortsDataSource struct {
	client *client.Client
}

// GetAllPtpPortsDataSourceModel describes the data source state shape.
type GetAllPtpPortsDataSourceModel struct {
	BoxId types.Int64 `tfsdk:"box_id" json:"boxId"`
	Items types.List  `tfsdk:"items"`
}

// NewGetAllPtpPortsDataSource returns a new instance of the generated data source.
func NewGetAllPtpPortsDataSource() datasource.DataSource {
	return &GetAllPtpPortsDataSource{}
}

// Metadata returns the data source type name.
func (d *GetAllPtpPortsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_all_ptp_ports"
}

// Schema returns the data source schema.
func (d *GetAllPtpPortsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Reads the get all ptp ports data source.", Attributes: map[string]schema.Attribute{"box_id": schema.Int64Attribute{MarkdownDescription: "specify the cluster node by boxId. By default all nodes are selected.", Optional: true}, "items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"announce_interval": schema.Int64Attribute{MarkdownDescription: "Logarithm to the base 2 of the mean announce Interval", Computed: true}, "announce_receipt_timeout": schema.Int64Attribute{MarkdownDescription: "Specify the number of announceInterval that has to pass without receipt of an Announce message before the occurrence of the expire event", Computed: true}, "clock_identity": schema.StringAttribute{MarkdownDescription: "Clock Identifier as per the IEEE 1588 standard which is represented by an integer 8-octet array", Computed: true}, "clock_port_id": schema.Int64Attribute{MarkdownDescription: "The value of the portNumber for a port on a PTP node supporting a single PTP port shall be 1. The values of the port numbers for the N ports on a PTP node supporting N PTP ports shall be 1, 2, ...N, respectively. The all-zeros and all-ones portNumber values are reserved.", Computed: true}, "clock_port_state": schema.StringAttribute{MarkdownDescription: "Value of the current state of the protocol engine associated with the given PTP port. (deprecated: use clockPortStateAlias)", Computed: true}, "clock_port_state_alias": schema.StringAttribute{MarkdownDescription: "Value of the current state of the protocol engine associated with the given PTP port", Computed: true}, "delay_mechanism": schema.StringAttribute{MarkdownDescription: "The propagation delay measuring option used by the port in computing mean path delay", Computed: true}, "delay_request_interval": schema.Int64Attribute{MarkdownDescription: "Logarithm to the base 2 of the permitted mean time interval between successive Delay_req messages", Computed: true}, "enabled": schema.BoolAttribute{Computed: true}, "local_priority": schema.Int64Attribute{Computed: true}, "peer_delay_request_interval": schema.Int64Attribute{MarkdownDescription: "Logarithm to the base 2 of the delayReqInterval", Computed: true}, "peer_mean_path_delay": schema.Int64Attribute{MarkdownDescription: "If the value of the delay Mechanism member is peer-to-peer (P2P), the value of peerMeanPathDelay shall be an estimate of the current one-way propagation delay on the link, i.e.,meanPathDelay, attached to this port computed using the peer delay mechanism", Computed: true}, "port_id": schema.StringAttribute{MarkdownDescription: "Port ID in [box/slot/portid] format", Computed: true}, "ptp_version": schema.StringAttribute{MarkdownDescription: "Indicates the version of the PTP standard implemented on the port. IEEE Std 1588-2008 corresponds to PTP version 2 whereas 1588-2002 is PTP version 1. For this implementation, PTPv2 will be supported", Computed: true}, "sync_interval": schema.Int64Attribute{MarkdownDescription: "Logarithm to the base 2 of the mean SyncInterval for multicast messages", Computed: true}, "vlan_id": schema.Int64Attribute{Computed: true}}}}}}
}

// Read fetches remote state into the data source model.
func (d *GetAllPtpPortsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetAllPtpPortsDataSourceModel
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
func (d *GetAllPtpPortsDataSource) readListRemote(ctx context.Context, config *GetAllPtpPortsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/ptp/portState"
	params := url.Values{}
	if !config.BoxId.IsNull() {
		params.Set("boxId", strconv.FormatInt(config.BoxId.ValueInt64(), 10))
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_ptp_ports", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_ptp_ports", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["portStates"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_ptp_ports", fmt.Sprintf("Could not decode list page: missing %q array", "portStates"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_ptp_ports", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetAllPtpPortsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
