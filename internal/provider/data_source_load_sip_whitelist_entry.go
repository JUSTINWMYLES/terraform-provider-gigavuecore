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
	_ datasource.DataSource              = (*LoadSipWhitelistEntryDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadSipWhitelistEntryDataSource)(nil)
)

// LoadSipWhitelistEntryDataSource is the generated Terraform data source implementation.
type LoadSipWhitelistEntryDataSource struct {
	client *client.Client
}

// LoadSipWhitelistEntryDataSourceModel describes the data source state shape.
type LoadSipWhitelistEntryDataSourceModel struct {
	ActiveSessions types.Int64  `tfsdk:"active_sessions" json:"activeSessions"`
	Alias          types.String `tfsdk:"alias"`
	CallerId       types.String `tfsdk:"caller_id" json:"callerId"`
	IdRange        types.Object `tfsdk:"id_range" json:"idRange"`
	IpAddress      types.Object `tfsdk:"ip_address" json:"ipAddress"`
}

// NewLoadSipWhitelistEntryDataSource returns a new instance of the generated data source.
func NewLoadSipWhitelistEntryDataSource() datasource.DataSource {
	return &LoadSipWhitelistEntryDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadSipWhitelistEntryDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_sip_whitelist_entry"
}

// Schema returns the data source schema.
func (d *LoadSipWhitelistEntryDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Check if active sessions for  Whitelist Entry exists", Attributes: map[string]schema.Attribute{"active_sessions": schema.Int64Attribute{MarkdownDescription: "Number of active sessions", Computed: true}, "alias": schema.StringAttribute{MarkdownDescription: "alias of the target SIP Whitelist", Required: true}, "caller_id": schema.StringAttribute{MarkdownDescription: "sip caller id", Computed: true, Optional: true}, "id_range": schema.SingleNestedAttribute{MarkdownDescription: "range of values from value to valueMax", Computed: true, Attributes: map[string]schema.Attribute{"value": schema.StringAttribute{Computed: true}, "value_max": schema.StringAttribute{Computed: true}}}, "ip_address": schema.SingleNestedAttribute{MarkdownDescription: "ipaddress based whitelist entry being queried. Example:ipAddress=(value:1.1.1.1)", Computed: true, Optional: true, Attributes: map[string]schema.Attribute{"value": schema.StringAttribute{MarkdownDescription: "ipv4 or ipv6", Computed: true}}}}}
}

// Read fetches remote state into the data source model.
func (d *LoadSipWhitelistEntryDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadSipWhitelistEntryDataSourceModel
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
func (d *LoadSipWhitelistEntryDataSource) readRemote(ctx context.Context, config *LoadSipWhitelistEntryDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/apps/sip/whitelists/{alias}/entries"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(config.Alias.ValueString()))
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_sip_whitelist_entry", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !config.CallerId.IsNull() {
		query.Set("callerId", config.CallerId.ValueString())
	}
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_sip_whitelist_entry", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_sip_whitelist_entry", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_sip_whitelist_entry", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_sip_whitelist_entry", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_sip_whitelist_entry", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_sip_whitelist_entry", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_sip_whitelist_entry", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_sip_whitelist_entry", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_sip_whitelist_entry", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_load_sip_whitelist_entry", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_load_sip_whitelist_entry", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_sip_whitelist_entry", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["sipWhitelistEntry"]; ok {
		if m, ok := v.(map[string]any); ok {
			data = m
		}
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_sip_whitelist_entry", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadSipWhitelistEntryDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
