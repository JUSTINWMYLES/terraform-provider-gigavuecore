package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)
import (
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*GetSupportedSnmpV3UserProtocolsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetSupportedSnmpV3UserProtocolsDataSource)(nil)
)

// GetSupportedSnmpV3UserProtocolsDataSource is the generated Terraform data source implementation.
type GetSupportedSnmpV3UserProtocolsDataSource struct {
	client *client.Client
}

// GetSupportedSnmpV3UserProtocolsDataSourceModel describes the data source state shape.
type GetSupportedSnmpV3UserProtocolsDataSourceModel struct {
	AuthProtocol types.String `tfsdk:"auth_protocol" json:"authProtocol"`
	PrivProtocol types.String `tfsdk:"priv_protocol" json:"privProtocol"`
	SwVersion    types.String `tfsdk:"sw_version" json:"swVersion"`
}

// NewGetSupportedSnmpV3UserProtocolsDataSource returns a new instance of the generated data source.
func NewGetSupportedSnmpV3UserProtocolsDataSource() datasource.DataSource {
	return &GetSupportedSnmpV3UserProtocolsDataSource{}
}

// Metadata returns the data source type name.
func (d *GetSupportedSnmpV3UserProtocolsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_supported_snmp_v3_user_protocols"
}

// Schema returns the data source schema.
func (d *GetSupportedSnmpV3UserProtocolsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Find SNMPv3 User Protocols by deviceice swVersion", Attributes: map[string]schema.Attribute{"auth_protocol": schema.StringAttribute{Computed: true}, "priv_protocol": schema.StringAttribute{Computed: true}, "sw_version": schema.StringAttribute{MarkdownDescription: "Device SW Version", Computed: true, Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetSupportedSnmpV3UserProtocolsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetSupportedSnmpV3UserProtocolsDataSourceModel
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
func (d *GetSupportedSnmpV3UserProtocolsDataSource) readRemote(ctx context.Context, config *GetSupportedSnmpV3UserProtocolsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/system/snmp/v3Users/supportedSnmpV3UserProtocols"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_supported_snmp_v3_user_protocols", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !config.SwVersion.IsNull() {
		query.Set("swVersion", config.SwVersion.ValueString())
	}
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_supported_snmp_v3_user_protocols", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_supported_snmp_v3_user_protocols", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_supported_snmp_v3_user_protocols", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_supported_snmp_v3_user_protocols", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_supported_snmp_v3_user_protocols", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_supported_snmp_v3_user_protocols", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_supported_snmp_v3_user_protocols", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_supported_snmp_v3_user_protocols", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_supported_snmp_v3_user_protocols", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_supported_snmp_v3_user_protocols", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_supported_snmp_v3_user_protocols", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_supported_snmp_v3_user_protocols", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_supported_snmp_v3_user_protocols", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetSupportedSnmpV3UserProtocolsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
