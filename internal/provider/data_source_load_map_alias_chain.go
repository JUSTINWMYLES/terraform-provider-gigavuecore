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
	_ datasource.DataSource              = (*LoadMapAliasChainDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadMapAliasChainDataSource)(nil)
)

// LoadMapAliasChainDataSource is the generated Terraform data source implementation.
type LoadMapAliasChainDataSource struct {
	client *client.Client
}

// LoadMapAliasChainDataSourceModel describes the data source state shape.
type LoadMapAliasChainDataSourceModel struct {
	ClusterId          types.String `tfsdk:"cluster_id" json:"clusterId"`
	Collector          types.String `tfsdk:"collector"`
	HealthState        types.String `tfsdk:"health_state" json:"healthState"`
	HealthStateReasons types.List   `tfsdk:"health_state_reasons" json:"healthStateReasons"`
	Id                 types.String `tfsdk:"id"`
	MapChainId         types.String `tfsdk:"map_chain_id" json:"mapChainId"`
	OrderedMapAliases  types.Set    `tfsdk:"ordered_map_aliases" json:"orderedMapAliases"`
	SrcPorts           types.List   `tfsdk:"src_ports" json:"srcPorts"`
	SrcPortsAsId       types.String `tfsdk:"src_ports_as_id" json:"srcPortsAsId"`
}

// NewLoadMapAliasChainDataSource returns a new instance of the generated data source.
func NewLoadMapAliasChainDataSource() datasource.DataSource {
	return &LoadMapAliasChainDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadMapAliasChainDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_map_alias_chain"
}

// Schema returns the data source schema.
func (d *LoadMapAliasChainDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load aliased map chain by source ports or mapChain ID", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "id of the defining cluster", Required: true}, "collector": schema.StringAttribute{MarkdownDescription: "alias of the chain collector map. must reference a map of matching type and  'collector' subtype", Computed: true}, "health_state": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "health_state_reasons": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"message": schema.StringAttribute{MarkdownDescription: "Read-only. Describes the reason for component's health state", Computed: true}, "severity": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "traffic_health_state_computation_type": schema.StringAttribute{MarkdownDescription: "Traffic Health State Computation Type", Computed: true}}}}, "id": schema.StringAttribute{MarkdownDescription: "srcPortsAsId or mapChainId for which the map chain ordered by map alias is to be displayed", Required: true}, "map_chain_id": schema.StringAttribute{MarkdownDescription: "mapChain ID - replaces srcPortsAsId - should be used anywhere that requires srcPortsAsId", Computed: true}, "ordered_map_aliases": schema.SetAttribute{MarkdownDescription: "aliases of the maps in the chain", Computed: true, ElementType: types.StringType}, "src_ports": schema.ListAttribute{MarkdownDescription: "network ports for the maps chain", Computed: true, ElementType: types.StringType}, "src_ports_as_id": schema.StringAttribute{MarkdownDescription: "(Deprecated - use mapChainId instead) mapChain ID", Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadMapAliasChainDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadMapAliasChainDataSourceModel
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
func (d *LoadMapAliasChainDataSource) readRemote(ctx context.Context, config *LoadMapAliasChainDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/mapChains/aliases/{id}"
	reqPath = strings.ReplaceAll(reqPath, "{id}", url.PathEscape(config.Id.ValueString()))
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_map_alias_chain", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_map_alias_chain", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_map_alias_chain", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_map_alias_chain", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_map_alias_chain", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_map_alias_chain", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_map_alias_chain", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_map_alias_chain", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_map_alias_chain", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_map_alias_chain", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_load_map_alias_chain", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_load_map_alias_chain", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_map_alias_chain", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["mapAliasChain"]; ok {
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
		resp.Diagnostics.AddError("Error reading gigavuecore_load_map_alias_chain", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadMapAliasChainDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
