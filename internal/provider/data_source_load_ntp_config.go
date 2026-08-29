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
	_ datasource.DataSource              = (*LoadNtpConfigDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadNtpConfigDataSource)(nil)
)

// LoadNtpConfigDataSource is the generated Terraform data source implementation.
type LoadNtpConfigDataSource struct {
	client *client.Client
}

// LoadNtpConfigDataSourceModel describes the data source state shape.
type LoadNtpConfigDataSourceModel struct {
	AuthEnabled types.Bool   `tfsdk:"auth_enabled" json:"authEnabled"`
	AuthKeys    types.List   `tfsdk:"auth_keys" json:"authKeys"`
	ClockSync   types.Bool   `tfsdk:"clock_sync" json:"clockSync"`
	ClusterId   types.String `tfsdk:"cluster_id" json:"clusterId"`
	Enabled     types.Bool   `tfsdk:"enabled"`
	NtpServers  types.List   `tfsdk:"ntp_servers" json:"ntpServers"`
	RefServer   types.String `tfsdk:"ref_server" json:"refServer"`
	Statuses    types.List   `tfsdk:"statuses"`
}

// NewLoadNtpConfigDataSource returns a new instance of the generated data source.
func NewLoadNtpConfigDataSource() datasource.DataSource {
	return &LoadNtpConfigDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadNtpConfigDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_ntp_config"
}

// Schema returns the data source schema.
func (d *LoadNtpConfigDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load NTP Configuration", Attributes: map[string]schema.Attribute{"auth_enabled": schema.BoolAttribute{Computed: true}, "auth_keys": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"key": schema.StringAttribute{MarkdownDescription: "MD5 key", Computed: true}, "key_number": schema.Int64Attribute{Computed: true}, "trusted": schema.BoolAttribute{Computed: true}}}}, "clock_sync": schema.BoolAttribute{MarkdownDescription: "indicates system's clock is synchronized with referenced NTP server", Computed: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "enabled": schema.BoolAttribute{MarkdownDescription: "enable/disable use of NTP for synchronization of the system's clock", Computed: true}, "ntp_servers": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"enabled": schema.BoolAttribute{Computed: true}, "key_enabled": schema.BoolAttribute{Computed: true}, "key_number": schema.Int64Attribute{Computed: true}, "preferred": schema.BoolAttribute{Computed: true}, "server": schema.StringAttribute{MarkdownDescription: "ipv4 or ipv6 or hostname", Computed: true}, "version": schema.StringAttribute{Computed: true}}}}, "ref_server": schema.StringAttribute{MarkdownDescription: "Address of NTP server used to synchronize the system's clock. ipv4 or ipv6 or hostname", Computed: true}, "statuses": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"address": schema.StringAttribute{MarkdownDescription: "ip address of the server", Computed: true}, "last_resp": schema.Int64Attribute{MarkdownDescription: "seconds", Computed: true}, "offset": schema.Float64Attribute{MarkdownDescription: "milliseconds in float", Computed: true}, "poll_interval": schema.Int64Attribute{MarkdownDescription: "seconds", Computed: true}, "ref_clock": schema.StringAttribute{Computed: true}, "status": schema.StringAttribute{MarkdownDescription: "server status", Computed: true}, "stratum": schema.Int64Attribute{Computed: true}}}}}}
}

// Read fetches remote state into the data source model.
func (d *LoadNtpConfigDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadNtpConfigDataSourceModel
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
func (d *LoadNtpConfigDataSource) readRemote(ctx context.Context, config *LoadNtpConfigDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/system/time/ntp"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_ntp_config", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_ntp_config", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_ntp_config", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_ntp_config", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_ntp_config", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_ntp_config", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_ntp_config", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_ntp_config", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_ntp_config", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_ntp_config", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_load_ntp_config", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_load_ntp_config", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_ntp_config", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_ntp_config", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadNtpConfigDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
