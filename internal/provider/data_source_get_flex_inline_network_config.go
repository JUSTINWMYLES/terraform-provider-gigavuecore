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
	_ datasource.DataSource              = (*GetFlexInlineNetworkConfigDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetFlexInlineNetworkConfigDataSource)(nil)
)

// GetFlexInlineNetworkConfigDataSource is the generated Terraform data source implementation.
type GetFlexInlineNetworkConfigDataSource struct {
	client *client.Client
}

// GetFlexInlineNetworkConfigDataSourceModel describes the data source state shape.
type GetFlexInlineNetworkConfigDataSourceModel struct {
	Alias               types.String  `tfsdk:"alias"`
	ClusterId           types.String  `tfsdk:"cluster_id" json:"clusterId"`
	ConfigData          types.Dynamic `tfsdk:"config_data" json:"configData"`
	ConfigStatus        types.String  `tfsdk:"config_status" json:"configStatus"`
	ConfigStatusReasons types.List    `tfsdk:"config_status_reasons" json:"configStatusReasons"`
	ConfigType          types.String  `tfsdk:"config_type" json:"configType"`
	HealthState         types.String  `tfsdk:"health_state" json:"healthState"`
	HealthStateReasons  types.List    `tfsdk:"health_state_reasons" json:"healthStateReasons"`
	SolutionAlias       types.String  `tfsdk:"solution_alias" json:"solutionAlias"`
}

// NewGetFlexInlineNetworkConfigDataSource returns a new instance of the generated data source.
func NewGetFlexInlineNetworkConfigDataSource() datasource.DataSource {
	return &GetFlexInlineNetworkConfigDataSource{}
}

// Metadata returns the data source type name.
func (d *GetFlexInlineNetworkConfigDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_flex_inline_network_config"
}

// Schema returns the data source schema.
func (d *GetFlexInlineNetworkConfigDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get inline networks details used in the given solution", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Solution alias for which inline network detail is needed", Required: true}, "cluster_id": schema.StringAttribute{Computed: true}, "config_data": schema.DynamicAttribute{MarkdownDescription: "Data holding configuration details", Computed: true}, "config_status": schema.StringAttribute{MarkdownDescription: "Status of the created config object", Computed: true}, "config_status_reasons": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "config_type": schema.StringAttribute{MarkdownDescription: "Type of the configuration object", Computed: true}, "health_state": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "health_state_reasons": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"message": schema.StringAttribute{MarkdownDescription: "Read-only. Describes the reason for component's health state", Computed: true}, "severity": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "traffic_health_state_computation_type": schema.StringAttribute{MarkdownDescription: "Traffic Health State Computation Type", Computed: true}}}}, "solution_alias": schema.StringAttribute{MarkdownDescription: "Alias of the solution", Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetFlexInlineNetworkConfigDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetFlexInlineNetworkConfigDataSourceModel
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
func (d *GetFlexInlineNetworkConfigDataSource) readRemote(ctx context.Context, config *GetFlexInlineNetworkConfigDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/flexInline/{alias}/inline/networks"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(config.Alias.ValueString()))
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_flex_inline_network_config", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_flex_inline_network_config", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_flex_inline_network_config", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_flex_inline_network_config", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_flex_inline_network_config", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_flex_inline_network_config", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_flex_inline_network_config", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_flex_inline_network_config", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_flex_inline_network_config", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_flex_inline_network_config", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_flex_inline_network_config", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_flex_inline_network_config", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_flex_inline_network_config", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["flexInlineConfig"]; ok {
		if m, ok := v.(map[string]any); ok {
			data = m
		} else if arr, ok := v.([]any); ok && len(arr) > 0 {
			if m, ok := arr[0].(map[string]any); ok {
				data = m
			}
		}
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_flex_inline_network_config", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetFlexInlineNetworkConfigDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
