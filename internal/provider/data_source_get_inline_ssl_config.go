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
	_ datasource.DataSource              = (*GetInlineSslConfigDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetInlineSslConfigDataSource)(nil)
)

// GetInlineSslConfigDataSource is the generated Terraform data source implementation.
type GetInlineSslConfigDataSource struct {
	client *client.Client
}

// GetInlineSslConfigDataSourceModel describes the data source state shape.
type GetInlineSslConfigDataSourceModel struct {
	Caching       types.Object `tfsdk:"caching"`
	ClusterId     types.String `tfsdk:"cluster_id" json:"clusterId"`
	DheCiphersuit types.String `tfsdk:"dhe_ciphersuit" json:"dheCiphersuit"`
	Monitor       types.Object `tfsdk:"monitor"`
	Resumption    types.Object `tfsdk:"resumption"`
	SslVersions   types.Object `tfsdk:"ssl_versions" json:"sslVersions"`
	StartTls      types.Object `tfsdk:"start_tls" json:"startTls"`
}

// NewGetInlineSslConfigDataSource returns a new instance of the generated data source.
func NewGetInlineSslConfigDataSource() datasource.DataSource {
	return &GetInlineSslConfigDataSource{}
}

// Metadata returns the data source type name.
func (d *GetInlineSslConfigDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_inline_ssl_config"
}

// Schema returns the data source schema.
func (d *GetInlineSslConfigDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load inline SSL global configuration", Attributes: map[string]schema.Attribute{"caching": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"persistence": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"enable": schema.BoolAttribute{Computed: true}}}}}, "cluster_id": schema.StringAttribute{MarkdownDescription: "id of the defining cluster", Required: true}, "dhe_ciphersuit": schema.StringAttribute{Computed: true}, "monitor": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"enable": schema.BoolAttribute{Computed: true}}}, "resumption": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"client": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"enable": schema.BoolAttribute{Computed: true}}}}}, "ssl_versions": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"connection_reset_action_for_max_version": schema.StringAttribute{MarkdownDescription: "Action to take to reset connection if TLS version is higher than configured", Computed: true}, "connection_reset_action_for_min_version": schema.StringAttribute{MarkdownDescription: "Action to take to reset connection if TLS version is lower than configured", Computed: true}, "max_version": schema.StringAttribute{MarkdownDescription: "maxVersion must be greater than minVersion", Computed: true}, "min_version": schema.StringAttribute{Computed: true}}}, "start_tls": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"enable": schema.BoolAttribute{Computed: true}}}}}
}

// Read fetches remote state into the data source model.
func (d *GetInlineSslConfigDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetInlineSslConfigDataSourceModel
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
func (d *GetInlineSslConfigDataSource) readRemote(ctx context.Context, config *GetInlineSslConfigDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/apps/inlineSsl"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_inline_ssl_config", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_inline_ssl_config", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_inline_ssl_config", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_inline_ssl_config", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_inline_ssl_config", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_inline_ssl_config", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_inline_ssl_config", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_inline_ssl_config", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_inline_ssl_config", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_inline_ssl_config", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_inline_ssl_config", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_inline_ssl_config", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_inline_ssl_config", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_inline_ssl_config", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetInlineSslConfigDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
