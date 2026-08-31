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
	_ datasource.DataSource              = (*FetchInlineSslProfileAllListInfoDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*FetchInlineSslProfileAllListInfoDataSource)(nil)
)

// FetchInlineSslProfileAllListInfoDataSource is the generated Terraform data source implementation.
type FetchInlineSslProfileAllListInfoDataSource struct {
	client *client.Client
}

// FetchInlineSslProfileAllListInfoDataSourceModel describes the data source state shape.
type FetchInlineSslProfileAllListInfoDataSourceModel struct {
	Alias         types.String `tfsdk:"alias"`
	Blacklist     types.Object `tfsdk:"blacklist"`
	Decryptlist   types.Object `tfsdk:"decryptlist"`
	Nodecryptlist types.Object `tfsdk:"nodecryptlist"`
	Whitelist     types.Object `tfsdk:"whitelist"`
}

// NewFetchInlineSslProfileAllListInfoDataSource returns a new instance of the generated data source.
func NewFetchInlineSslProfileAllListInfoDataSource() datasource.DataSource {
	return &FetchInlineSslProfileAllListInfoDataSource{}
}

// Metadata returns the data source type name.
func (d *FetchInlineSslProfileAllListInfoDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_fetch_inline_ssl_profile_all_list_info"
}

// Schema returns the data source schema.
func (d *FetchInlineSslProfileAllListInfoDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Fetch inline SSL profile nodecryptlist and decryptlist information", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Alias of the inline SSL profile", Required: true}, "blacklist": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"num_entries": schema.Int64Attribute{MarkdownDescription: "the number of entries in the whitelist or blacklist. (deprecated: use nodecryptlist and decryptlist instead of whitelist and blacklist)", Computed: true}}}, "decryptlist": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"num_entries": schema.Int64Attribute{MarkdownDescription: "the number of entries in the nodecryptlist or decryptlist", Computed: true}}}, "nodecryptlist": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"num_entries": schema.Int64Attribute{MarkdownDescription: "the number of entries in the nodecryptlist or decryptlist", Computed: true}}}, "whitelist": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"num_entries": schema.Int64Attribute{MarkdownDescription: "the number of entries in the whitelist or blacklist. (deprecated: use nodecryptlist and decryptlist instead of whitelist and blacklist)", Computed: true}}}}}
}

// Read fetches remote state into the data source model.
func (d *FetchInlineSslProfileAllListInfoDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config FetchInlineSslProfileAllListInfoDataSourceModel
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
func (d *FetchInlineSslProfileAllListInfoDataSource) readRemote(ctx context.Context, config *FetchInlineSslProfileAllListInfoDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/apps/inlineSsl/profiles/{alias}/list"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(config.Alias.ValueString()))
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_fetch_inline_ssl_profile_all_list_info", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_fetch_inline_ssl_profile_all_list_info", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_fetch_inline_ssl_profile_all_list_info", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_fetch_inline_ssl_profile_all_list_info", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_fetch_inline_ssl_profile_all_list_info", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_fetch_inline_ssl_profile_all_list_info", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_fetch_inline_ssl_profile_all_list_info", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_fetch_inline_ssl_profile_all_list_info", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_fetch_inline_ssl_profile_all_list_info", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_fetch_inline_ssl_profile_all_list_info", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_fetch_inline_ssl_profile_all_list_info", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_fetch_inline_ssl_profile_all_list_info", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_fetch_inline_ssl_profile_all_list_info", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_fetch_inline_ssl_profile_all_list_info", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *FetchInlineSslProfileAllListInfoDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
