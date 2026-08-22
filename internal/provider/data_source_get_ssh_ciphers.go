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
	_ datasource.DataSource              = (*GetSshCiphersDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetSshCiphersDataSource)(nil)
)

// GetSshCiphersDataSource is the generated Terraform data source implementation.
type GetSshCiphersDataSource struct {
	client *client.Client
}

// GetSshCiphersDataSourceModel describes the data source state shape.
type GetSshCiphersDataSourceModel struct {
	ClientCiphers types.List `tfsdk:"client_ciphers" json:"clientCiphers"`
	ClientHostkey types.List `tfsdk:"client_hostkey" json:"clientHostkey"`
	ClientKex     types.List `tfsdk:"client_kex" json:"clientKex"`
	ClientMacs    types.List `tfsdk:"client_macs" json:"clientMacs"`
	ServerCiphers types.List `tfsdk:"server_ciphers" json:"serverCiphers"`
	ServerHostkey types.List `tfsdk:"server_hostkey" json:"serverHostkey"`
	ServerKex     types.List `tfsdk:"server_kex" json:"serverKex"`
	ServerMacs    types.List `tfsdk:"server_macs" json:"serverMacs"`
}

// NewGetSshCiphersDataSource returns a new instance of the generated data source.
func NewGetSshCiphersDataSource() datasource.DataSource {
	return &GetSshCiphersDataSource{}
}

// Metadata returns the data source type name.
func (d *GetSshCiphersDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_ssh_ciphers"
}

// Schema returns the data source schema.
func (d *GetSshCiphersDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load all ciphers, Kex, Macs and HostKey", Attributes: map[string]schema.Attribute{"client_ciphers": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "client_hostkey": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "client_kex": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "client_macs": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "server_ciphers": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "server_hostkey": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "server_kex": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "server_macs": schema.ListAttribute{Computed: true, ElementType: types.StringType}}}
}

// Read fetches remote state into the data source model.
func (d *GetSshCiphersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetSshCiphersDataSourceModel
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
func (d *GetSshCiphersDataSource) readRemote(ctx context.Context, config *GetSshCiphersDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/system/ssh/config"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_ssh_ciphers", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_ssh_ciphers", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_ssh_ciphers", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ssh_ciphers", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ssh_ciphers", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ssh_ciphers", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ssh_ciphers", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ssh_ciphers", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ssh_ciphers", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ssh_ciphers", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_ssh_ciphers", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ssh_ciphers", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_ssh_ciphers", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_ssh_ciphers", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetSshCiphersDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
