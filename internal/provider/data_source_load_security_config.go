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
	_ datasource.DataSource              = (*LoadSecurityConfigDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadSecurityConfigDataSource)(nil)
)

// LoadSecurityConfigDataSource is the generated Terraform data source implementation.
type LoadSecurityConfigDataSource struct {
	client *client.Client
}

// LoadSecurityConfigDataSourceModel describes the data source state shape.
type LoadSecurityConfigDataSourceModel struct {
	AllowBlankPassword   types.Bool   `tfsdk:"allow_blank_password" json:"allowBlankPassword"`
	ClusterId            types.String `tfsdk:"cluster_id" json:"clusterId"`
	FipsEnabled          types.Bool   `tfsdk:"fips_enabled" json:"fipsEnabled"`
	FipsModeState        types.Bool   `tfsdk:"fips_mode_state" json:"fipsModeState"`
	MinPasswordLen       types.Int64  `tfsdk:"min_password_len" json:"minPasswordLen"`
	SecureCrypto         types.Bool   `tfsdk:"secure_crypto" json:"secureCrypto"`
	SecureCryptoEnforced types.Bool   `tfsdk:"secure_crypto_enforced" json:"secureCryptoEnforced"`
	SecurePasswords      types.Bool   `tfsdk:"secure_passwords" json:"securePasswords"`
}

// NewLoadSecurityConfigDataSource returns a new instance of the generated data source.
func NewLoadSecurityConfigDataSource() datasource.DataSource {
	return &LoadSecurityConfigDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadSecurityConfigDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_security_config"
}

// Schema returns the data source schema.
func (d *LoadSecurityConfigDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load Security config", Attributes: map[string]schema.Attribute{"allow_blank_password": schema.BoolAttribute{Computed: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "id of the defining cluster", Required: true}, "fips_enabled": schema.BoolAttribute{MarkdownDescription: "enable/disable fips mode(pending fips mode), system reload is necessary to activate fips mode", Computed: true}, "fips_mode_state": schema.BoolAttribute{MarkdownDescription: "current fips mode", Computed: true}, "min_password_len": schema.Int64Attribute{Computed: true}, "secure_crypto": schema.BoolAttribute{MarkdownDescription: "Secure crypto mode. Value takes effect after reload", Computed: true}, "secure_crypto_enforced": schema.BoolAttribute{MarkdownDescription: "Secure crypto mode enforced", Computed: true}, "secure_passwords": schema.BoolAttribute{MarkdownDescription: "Secure passwords mode", Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadSecurityConfigDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadSecurityConfigDataSourceModel
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
func (d *LoadSecurityConfigDataSource) readRemote(ctx context.Context, config *LoadSecurityConfigDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/system/security"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_security_config", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_security_config", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_security_config", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_security_config", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_security_config", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_security_config", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_security_config", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_security_config", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_security_config", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_security_config", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_load_security_config", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_load_security_config", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_security_config", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_security_config", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadSecurityConfigDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
