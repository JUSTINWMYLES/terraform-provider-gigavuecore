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
	_ datasource.DataSource              = (*GetKeystorePreferenceDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetKeystorePreferenceDataSource)(nil)
)

// GetKeystorePreferenceDataSource is the generated Terraform data source implementation.
type GetKeystorePreferenceDataSource struct {
	client *client.Client
}

// GetKeystorePreferenceDataSourceModel describes the data source state shape.
type GetKeystorePreferenceDataSourceModel struct {
	AutoDelete    types.Bool   `tfsdk:"auto_delete" json:"autoDelete"`
	AutoEnable    types.Bool   `tfsdk:"auto_enable" json:"autoEnable"`
	AutoPurge     types.Bool   `tfsdk:"auto_purge" json:"autoPurge"`
	ClusterId     types.String `tfsdk:"cluster_id" json:"clusterId"`
	MaxKeys       types.Int64  `tfsdk:"max_keys" json:"maxKeys"`
	RetentionTime types.Int64  `tfsdk:"retention_time" json:"retentionTime"`
}

// NewGetKeystorePreferenceDataSource returns a new instance of the generated data source.
func NewGetKeystorePreferenceDataSource() datasource.DataSource {
	return &GetKeystorePreferenceDataSource{}
}

// Metadata returns the data source type name.
func (d *GetKeystorePreferenceDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_keystore_preference"
}

// Schema returns the data source schema.
func (d *GetKeystorePreferenceDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load keystore preference", Attributes: map[string]schema.Attribute{"auto_delete": schema.BoolAttribute{Computed: true}, "auto_enable": schema.BoolAttribute{Computed: true}, "auto_purge": schema.BoolAttribute{Computed: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "max_keys": schema.Int64Attribute{Computed: true}, "retention_time": schema.Int64Attribute{Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetKeystorePreferenceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetKeystorePreferenceDataSourceModel
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
func (d *GetKeystorePreferenceDataSource) readRemote(ctx context.Context, config *GetKeystorePreferenceDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/apps/keystore/preference"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_keystore_preference", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_keystore_preference", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_keystore_preference", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_keystore_preference", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_keystore_preference", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_keystore_preference", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_keystore_preference", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_keystore_preference", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_keystore_preference", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_keystore_preference", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_keystore_preference", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_keystore_preference", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_keystore_preference", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["preference"]; ok {
		if m, ok := v.(map[string]any); ok {
			data = m
		}
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_keystore_preference", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetKeystorePreferenceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
