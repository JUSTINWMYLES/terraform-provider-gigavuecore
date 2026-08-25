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
	_ datasource.DataSource              = (*GetSslDecryptionKeysStoreStateDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetSslDecryptionKeysStoreStateDataSource)(nil)
)

// GetSslDecryptionKeysStoreStateDataSource is the generated Terraform data source implementation.
type GetSslDecryptionKeysStoreStateDataSource struct {
	client *client.Client
}

// GetSslDecryptionKeysStoreStateDataSourceModel describes the data source state shape.
type GetSslDecryptionKeysStoreStateDataSourceModel struct {
	ClusterId   types.String `tfsdk:"cluster_id" json:"clusterId"`
	Locked      types.Bool   `tfsdk:"locked"`
	PasswordSet types.Bool   `tfsdk:"password_set" json:"passwordSet"`
}

// NewGetSslDecryptionKeysStoreStateDataSource returns a new instance of the generated data source.
func NewGetSslDecryptionKeysStoreStateDataSource() datasource.DataSource {
	return &GetSslDecryptionKeysStoreStateDataSource{}
}

// Metadata returns the data source type name.
func (d *GetSslDecryptionKeysStoreStateDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_ssl_decryption_keys_store_state"
}

// Schema returns the data source schema.
func (d *GetSslDecryptionKeysStoreStateDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get keystore keychain State", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "locked": schema.BoolAttribute{MarkdownDescription: "indicates whether SSL Decryption Key Store is unlocked", Computed: true}, "password_set": schema.BoolAttribute{MarkdownDescription: "indicates whether SSL Decryption Key Store password is set", Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetSslDecryptionKeysStoreStateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetSslDecryptionKeysStoreStateDataSourceModel
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
func (d *GetSslDecryptionKeysStoreStateDataSource) readRemote(ctx context.Context, config *GetSslDecryptionKeysStoreStateDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/apps/keystore/keychain"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_ssl_decryption_keys_store_state", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_ssl_decryption_keys_store_state", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_ssl_decryption_keys_store_state", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ssl_decryption_keys_store_state", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ssl_decryption_keys_store_state", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ssl_decryption_keys_store_state", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ssl_decryption_keys_store_state", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ssl_decryption_keys_store_state", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ssl_decryption_keys_store_state", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ssl_decryption_keys_store_state", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_ssl_decryption_keys_store_state", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ssl_decryption_keys_store_state", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_ssl_decryption_keys_store_state", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_ssl_decryption_keys_store_state", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetSslDecryptionKeysStoreStateDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
