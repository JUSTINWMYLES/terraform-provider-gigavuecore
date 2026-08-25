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
	_ datasource.DataSource              = (*GetAllNtpServerDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetAllNtpServerDataSource)(nil)
)

// GetAllNtpServerDataSource is the generated Terraform data source implementation.
type GetAllNtpServerDataSource struct {
	client *client.Client
}

// GetAllNtpServerDataSourceModel describes the data source state shape.
type GetAllNtpServerDataSourceModel struct {
	AuthStatus   types.String `tfsdk:"auth_status" json:"authStatus"`
	Enabled      types.Bool   `tfsdk:"enabled"`
	FmIp         types.String `tfsdk:"fm_ip" json:"fmIp"`
	NtpServers   types.List   `tfsdk:"ntp_servers" json:"ntpServers"`
	Page         types.String `tfsdk:"page"`
	ServerHost   types.String `tfsdk:"server_host" json:"serverHost"`
	ServerStatus types.String `tfsdk:"server_status" json:"serverStatus"`
}

// NewGetAllNtpServerDataSource returns a new instance of the generated data source.
func NewGetAllNtpServerDataSource() datasource.DataSource {
	return &GetAllNtpServerDataSource{}
}

// Metadata returns the data source type name.
func (d *GetAllNtpServerDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_all_ntp_server"
}

// Schema returns the data source schema.
func (d *GetAllNtpServerDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get all NTP Servers in FM", Attributes: map[string]schema.Attribute{"auth_status": schema.StringAttribute{MarkdownDescription: "Auth status of NTP", Optional: true}, "enabled": schema.BoolAttribute{Computed: true}, "fm_ip": schema.StringAttribute{MarkdownDescription: "FMHighAvailability node IpAddress/domainName", Optional: true}, "ntp_servers": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"auth_required": schema.BoolAttribute{MarkdownDescription: "authentication enabled status", Computed: true}, "fm_ip": schema.StringAttribute{MarkdownDescription: "FM HighAvailability Node address/host", Computed: true}, "is_user_ntp_server": schema.BoolAttribute{MarkdownDescription: "To differentiate user created ntp server and default ntp server", Computed: true}, "ntp_auth": schema.SingleNestedAttribute{MarkdownDescription: "NTP Auth Details", Computed: true, Attributes: map[string]schema.Attribute{"key": schema.StringAttribute{MarkdownDescription: "key for the ntp server", Computed: true}, "type": schema.StringAttribute{MarkdownDescription: "Type of algorithm used for ntp server authentication", Computed: true}, "value": schema.StringAttribute{MarkdownDescription: "value for the ntp server authentication", Computed: true}}}, "server_host": schema.StringAttribute{MarkdownDescription: "Ip/Host Address", Computed: true}, "server_status": schema.SingleNestedAttribute{MarkdownDescription: "Server Status", Computed: true, Attributes: map[string]schema.Attribute{"offset": schema.Int64Attribute{Computed: true}, "poll_interval": schema.StringAttribute{Computed: true}, "status": schema.StringAttribute{MarkdownDescription: "status of the server", Computed: true}, "stratum": schema.StringAttribute{MarkdownDescription: "status of the server", Computed: true}}}, "version": schema.Int64Attribute{MarkdownDescription: "version of server", Computed: true}}}}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned", Optional: true}, "server_host": schema.StringAttribute{MarkdownDescription: "NTP server/host address", Optional: true}, "server_status": schema.StringAttribute{MarkdownDescription: "NTP server status", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetAllNtpServerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetAllNtpServerDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	d.readRemote(ctx, &config, resp)
	if resp.Diagnostics.HasError() {
		return
	}
	if config.Page.IsNull() {
		resp.Diagnostics.AddWarning("Single-page result", fmt.Sprintf("This data source reads a single page of a paginated API endpoint and does not aggregate results across pages. The \"page\" argument is unset, so the default page is returned; set it to retrieve a different page."))
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}

// readRemote performs the read HTTP exchange and decodes the response into config. Extracted from Read so the request/response logic is unit-testable without a tfsdk.Config.
func (d *GetAllNtpServerDataSource) readRemote(ctx context.Context, config *GetAllNtpServerDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/fm/system/time/ntp/servers"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_ntp_server", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !config.ServerHost.IsNull() {
		query.Set("serverHost", config.ServerHost.ValueString())
	}
	if !config.ServerStatus.IsNull() {
		query.Set("serverStatus", config.ServerStatus.ValueString())
	}
	if !config.AuthStatus.IsNull() {
		query.Set("authStatus", config.AuthStatus.ValueString())
	}
	if !config.FmIp.IsNull() {
		query.Set("fmIp", config.FmIp.ValueString())
	}
	if !config.Page.IsNull() {
		query.Set("page", config.Page.ValueString())
	}
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_ntp_server", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_ntp_server", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_ntp_server", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_ntp_server", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_ntp_server", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_ntp_server", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_ntp_server", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_ntp_server", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_ntp_server", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_all_ntp_server", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_ntp_server", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_ntp_server", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_ntp_server", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetAllNtpServerDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
