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
	_ datasource.DataSource              = (*LoadsGlobalFmTemplateWithValuesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadsGlobalFmTemplateWithValuesDataSource)(nil)
)

// LoadsGlobalFmTemplateWithValuesDataSource is the generated Terraform data source implementation.
type LoadsGlobalFmTemplateWithValuesDataSource struct {
	client *client.Client
}

// LoadsGlobalFmTemplateWithValuesDataSourceModel describes the data source state shape.
type LoadsGlobalFmTemplateWithValuesDataSourceModel struct {
	Classic types.Object `tfsdk:"classic"`
	Crypto  types.Object `tfsdk:"crypto"`
	Fips    types.Object `tfsdk:"fips"`
}

// NewLoadsGlobalFmTemplateWithValuesDataSource returns a new instance of the generated data source.
func NewLoadsGlobalFmTemplateWithValuesDataSource() datasource.DataSource {
	return &LoadsGlobalFmTemplateWithValuesDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadsGlobalFmTemplateWithValuesDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_loads_global_fm_template_with_values"
}

// Schema returns the data source schema.
func (d *LoadsGlobalFmTemplateWithValuesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "new in FM 6.6", Attributes: map[string]schema.Attribute{"classic": schema.SingleNestedAttribute{MarkdownDescription: "System SSH Cipher", Computed: true, Attributes: map[string]schema.Attribute{"client_ciphers": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "client_hostkey": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "client_kex": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "client_macs": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "server_ciphers": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "server_hostkey": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "server_kex": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "server_macs": schema.ListAttribute{Computed: true, ElementType: types.StringType}}}, "crypto": schema.SingleNestedAttribute{MarkdownDescription: "System SSH Cipher", Computed: true, Attributes: map[string]schema.Attribute{"client_ciphers": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "client_hostkey": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "client_kex": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "client_macs": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "server_ciphers": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "server_hostkey": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "server_kex": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "server_macs": schema.ListAttribute{Computed: true, ElementType: types.StringType}}}, "fips": schema.SingleNestedAttribute{MarkdownDescription: "System SSH Cipher", Computed: true, Attributes: map[string]schema.Attribute{"client_ciphers": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "client_hostkey": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "client_kex": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "client_macs": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "server_ciphers": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "server_hostkey": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "server_kex": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "server_macs": schema.ListAttribute{Computed: true, ElementType: types.StringType}}}}}
}

// Read fetches remote state into the data source model.
func (d *LoadsGlobalFmTemplateWithValuesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadsGlobalFmTemplateWithValuesDataSourceModel
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
func (d *LoadsGlobalFmTemplateWithValuesDataSource) readRemote(ctx context.Context, config *LoadsGlobalFmTemplateWithValuesDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/fm/templates/sshCiphers"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_loads_global_fm_template_with_values", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_loads_global_fm_template_with_values", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_loads_global_fm_template_with_values", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_loads_global_fm_template_with_values", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_loads_global_fm_template_with_values", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_loads_global_fm_template_with_values", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_loads_global_fm_template_with_values", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_loads_global_fm_template_with_values", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_loads_global_fm_template_with_values", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_loads_global_fm_template_with_values", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_loads_global_fm_template_with_values", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_loads_global_fm_template_with_values", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_loads_global_fm_template_with_values", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_loads_global_fm_template_with_values", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadsGlobalFmTemplateWithValuesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
