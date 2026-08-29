package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)
import (
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*FmHaDetailsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*FmHaDetailsDataSource)(nil)
)

// FmHaDetailsDataSource is the generated Terraform data source implementation.
type FmHaDetailsDataSource struct {
	client *client.Client
}

// FmHaDetailsDataSourceModel describes the data source state shape.
type FmHaDetailsDataSourceModel struct {
	FipsEnabled       types.Bool   `tfsdk:"fips_enabled" json:"fipsEnabled"`
	FmHaTunnel        types.Object `tfsdk:"fm_ha_tunnel" json:"fmHaTunnel"`
	HaStatus          types.Object `tfsdk:"ha_status" json:"haStatus"`
	Hostname          types.String `tfsdk:"hostname"`
	LoadSystemDetails types.Bool   `tfsdk:"load_system_details" json:"loadSystemDetails"`
	Name              types.String `tfsdk:"name"`
	Nodes             types.Object `tfsdk:"nodes"`
}

// NewFmHaDetailsDataSource returns a new instance of the generated data source.
func NewFmHaDetailsDataSource() datasource.DataSource {
	return &FmHaDetailsDataSource{}
}

// Metadata returns the data source type name.
func (d *FmHaDetailsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_fm_ha_details"
}

// Schema returns the data source schema.
func (d *FmHaDetailsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get fmHa group details", Attributes: map[string]schema.Attribute{"fips_enabled": schema.BoolAttribute{MarkdownDescription: "fips enabled", Computed: true}, "fm_ha_tunnel": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"tunnel_auth_mode": schema.StringAttribute{MarkdownDescription: "Auth mode for setting up FMHA tunnel. PSK is Pre Shared Key. PKI is Public Key Infrastructure(Certificate based)", Computed: true}, "tunnel_status": schema.SingleNestedAttribute{MarkdownDescription: "Includes details of the configured tunnels and their health status in FMHA cluster", Computed: true, Attributes: map[string]schema.Attribute{"total_tunnels": schema.SingleNestedAttribute{MarkdownDescription: "Lists all the configured tunnels in each node in a FMHA cluster", Computed: true, Attributes: map[string]schema.Attribute{"x198_51_100_42": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "x198_51_100_43": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "x198_51_100_44": schema.ListAttribute{Computed: true, ElementType: types.StringType}}}, "tunnel_health": schema.StringAttribute{MarkdownDescription: "Overall health status of all tunnels configured in FMHA cluster", Computed: true}, "unreachable_tunnels": schema.SingleNestedAttribute{MarkdownDescription: "Lists the unreachable tunnels from a particular FMHA node", Computed: true, Attributes: map[string]schema.Attribute{"x198_51_100_42": schema.ListAttribute{Computed: true, ElementType: types.StringType}}}}}}}, "ha_status": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"active_eligible": schema.StringAttribute{MarkdownDescription: "active eligible for ha status", Computed: true}, "support": schema.StringAttribute{MarkdownDescription: "support for ha status", Computed: true}}}, "hostname": schema.StringAttribute{MarkdownDescription: "DNS Name or IP Address of HA group", Computed: true}, "load_system_details": schema.BoolAttribute{MarkdownDescription: "Load system details", Required: true}, "name": schema.StringAttribute{MarkdownDescription: "name of the HA group", Computed: true}, "nodes": schema.SingleNestedAttribute{MarkdownDescription: "Nodes of HA group", Computed: true, Attributes: map[string]schema.Attribute{"cluster_ip_address": schema.StringAttribute{MarkdownDescription: "Cluster IP address for FM HA node", Computed: true}, "entity_id": schema.StringAttribute{MarkdownDescription: "Entity Id for FM HA node", Computed: true}, "hostname": schema.StringAttribute{MarkdownDescription: "DNS Name or IP Address for FM HA node", Computed: true}, "idp_meta_data_url": schema.StringAttribute{MarkdownDescription: "IDP Meta data URL for FM HA node", Computed: true}, "management_ip_address": schema.StringAttribute{MarkdownDescription: "Management IP address for FM HA node ", Computed: true}, "password": schema.StringAttribute{MarkdownDescription: "password for FM HA node", Computed: true, Sensitive: true}, "public_ip_address": schema.StringAttribute{MarkdownDescription: "Public IP address for FM HA node", Computed: true}, "reachable": schema.BoolAttribute{MarkdownDescription: "Reachable for FM HA node", Computed: true}, "seed_node": schema.BoolAttribute{MarkdownDescription: "Seed node for FM HA node", Computed: true}, "username": schema.StringAttribute{MarkdownDescription: "username for FM HA node", Computed: true}}}}}
}

// Read fetches remote state into the data source model.
func (d *FmHaDetailsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config FmHaDetailsDataSourceModel
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
func (d *FmHaDetailsDataSource) readRemote(ctx context.Context, config *FmHaDetailsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/fmHa"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_fm_ha_details", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("loadSystemDetails", strconv.FormatBool(config.LoadSystemDetails.ValueBool()))
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_fm_ha_details", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_fm_ha_details", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_fm_ha_details", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_fm_ha_details", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_fm_ha_details", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_fm_ha_details", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_fm_ha_details", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_fm_ha_details", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_fm_ha_details", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_fm_ha_details", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_fm_ha_details", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_fm_ha_details", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["gigaFmHaGroup"]; ok {
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
		resp.Diagnostics.AddError("Error reading gigavuecore_fm_ha_details", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *FmHaDetailsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
