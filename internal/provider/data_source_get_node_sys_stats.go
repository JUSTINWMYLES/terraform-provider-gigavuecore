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
	_ datasource.DataSource              = (*GetNodeSysStatsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetNodeSysStatsDataSource)(nil)
)

// GetNodeSysStatsDataSource is the generated Terraform data source implementation.
type GetNodeSysStatsDataSource struct {
	client *client.Client
}

// GetNodeSysStatsDataSourceModel describes the data source state shape.
type GetNodeSysStatsDataSourceModel struct {
	EnvId   types.String `tfsdk:"env_id" json:"envId"`
	Hosts   types.Object `tfsdk:"hosts"`
	NodeId  types.String `tfsdk:"node_id" json:"nodeId"`
	UnifyId types.String `tfsdk:"unify_id" json:"unifyId"`
}

// NewGetNodeSysStatsDataSource returns a new instance of the generated data source.
func NewGetNodeSysStatsDataSource() datasource.DataSource {
	return &GetNodeSysStatsDataSource{}
}

// Metadata returns the data source type name.
func (d *GetNodeSysStatsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_node_sys_stats"
}

// Schema returns the data source schema.
func (d *GetNodeSysStatsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Obtain a node sysStats of unified deployment node via environment, connection and node id", Attributes: map[string]schema.Attribute{"env_id": schema.StringAttribute{MarkdownDescription: "unified environment identifier", Required: true}, "hosts": schema.SingleNestedAttribute{MarkdownDescription: "Node sysstat hosts struct of a node in unified deployment", Computed: true, Attributes: map[string]schema.Attribute{"date": schema.StringAttribute{Computed: true}, "machine": schema.StringAttribute{Computed: true}, "nodename": schema.StringAttribute{Computed: true}, "number_of_avail_cpus": schema.Int64Attribute{Computed: true}, "number_of_cpus": schema.Int64Attribute{Computed: true}, "release": schema.StringAttribute{Computed: true}, "sysname": schema.StringAttribute{Computed: true}}}, "node_id": schema.StringAttribute{MarkdownDescription: "unified resource node identifier", Required: true}, "unify_id": schema.StringAttribute{MarkdownDescription: "unified resource identifier", Required: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetNodeSysStatsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetNodeSysStatsDataSourceModel
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
func (d *GetNodeSysStatsDataSource) readRemote(ctx context.Context, config *GetNodeSysStatsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/unifiedResources/env/{envId}/unifyId/{unifyId}/node/{nodeId}/sysStats"
	reqPath = strings.ReplaceAll(reqPath, "{envId}", url.PathEscape(config.EnvId.ValueString()))
	reqPath = strings.ReplaceAll(reqPath, "{unifyId}", url.PathEscape(config.UnifyId.ValueString()))
	reqPath = strings.ReplaceAll(reqPath, "{nodeId}", url.PathEscape(config.NodeId.ValueString()))
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_node_sys_stats", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_node_sys_stats", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_node_sys_stats", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_node_sys_stats", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_node_sys_stats", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_node_sys_stats", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_node_sys_stats", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_node_sys_stats", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_node_sys_stats", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_node_sys_stats", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_node_sys_stats", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_node_sys_stats", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_node_sys_stats", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["sysstat"]; ok {
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_node_sys_stats", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetNodeSysStatsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
