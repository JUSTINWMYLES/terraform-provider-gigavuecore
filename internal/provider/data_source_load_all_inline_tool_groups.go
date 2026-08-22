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
	_ datasource.DataSource              = (*LoadAllInlineToolGroupsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadAllInlineToolGroupsDataSource)(nil)
)

// LoadAllInlineToolGroupsDataSource is the generated Terraform data source implementation.
type LoadAllInlineToolGroupsDataSource struct {
	client *client.Client
}

// LoadAllInlineToolGroupsDataSourceModel describes the data source state shape.
type LoadAllInlineToolGroupsDataSourceModel struct {
	ClusterId        types.String `tfsdk:"cluster_id" json:"clusterId"`
	Context          types.Object `tfsdk:"context"`
	InlineToolGroups types.List   `tfsdk:"inline_tool_groups" json:"inlineToolGroups"`
	Page             types.String `tfsdk:"page"`
	Sort             types.String `tfsdk:"sort"`
}

// NewLoadAllInlineToolGroupsDataSource returns a new instance of the generated data source.
func NewLoadAllInlineToolGroupsDataSource() datasource.DataSource {
	return &LoadAllInlineToolGroupsDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadAllInlineToolGroupsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_all_inline_tool_groups"
}

// Schema returns the data source schema.
func (d *LoadAllInlineToolGroupsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load all Inline Tool Groups", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "context": schema.SingleNestedAttribute{MarkdownDescription: "Gigamon query result context", Computed: true, Attributes: map[string]schema.Attribute{"page_no": schema.Int64Attribute{MarkdownDescription: "page number of the returned result set", Computed: true}, "page_size": schema.Int64Attribute{MarkdownDescription: "page size of the returned result set", Computed: true}, "sort": schema.ListAttribute{MarkdownDescription: "sorting info of the returned result set. list of fields in the array indicate sorting order", Computed: true, ElementType: types.StringType}, "total_items": schema.Int64Attribute{MarkdownDescription: "total number of items in the queried entity type", Computed: true}}}, "inline_tool_groups": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Inline Tool Group alias. Unique within a cluster", Computed: true}, "comment": schema.StringAttribute{Computed: true}, "current_state": schema.SingleNestedAttribute{MarkdownDescription: "Inline Tool Group Smart Load Balancing definition", Computed: true, Attributes: map[string]schema.Attribute{"inline_tools": schema.ListAttribute{MarkdownDescription: "Current list of inlineTools", Computed: true, ElementType: types.StringType}, "spare_tool": schema.StringAttribute{MarkdownDescription: "Current spare tool", Computed: true}, "spare_tool_status": schema.StringAttribute{MarkdownDescription: "Current spare tool status", Computed: true}, "switched_inline_tool": schema.StringAttribute{MarkdownDescription: "Current switched inlineTool", Computed: true}}}, "enabled": schema.BoolAttribute{MarkdownDescription: "setting to false is equivalent to forcing the inline tool group failure (useful for taking the inline tool group out of commission for maintenance or other purposes)", Computed: true}, "failover_action": schema.StringAttribute{Computed: true}, "failover_mode": schema.StringAttribute{MarkdownDescription: "the way of handling a failure of an individual member of the inline tool list when no spare inline tool is configured or if the spare inline tool is failed", Computed: true}, "flex_status": schema.StringAttribute{Computed: true}, "flex_traffic_path": schema.StringAttribute{Computed: true}, "hash": schema.StringAttribute{Computed: true}, "health_state": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "health_state_reasons": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"message": schema.StringAttribute{MarkdownDescription: "Read-only. Describes the reason for component's health state", Computed: true}, "severity": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "traffic_health_state_computation_type": schema.StringAttribute{MarkdownDescription: "Traffic Health State Computation Type", Computed: true}}}}, "inline_tools": schema.ListAttribute{MarkdownDescription: "If no spare inline tool configured, list of aliases of inline tools participating in hash-based traffic distribution. If the spare inline tool is configured, list of aliases of primary inline tools to which traffic is forwarded as long as all of them are healthy. The number of inline tools in the list must be between 1 and 64 if the spare inline tool is configured or between 2 and 64 otherwise", Computed: true, ElementType: types.StringType}, "min_group_size": schema.Int64Attribute{MarkdownDescription: " the minimum number of inline tools (the inline tools in the list plus spare if configured) that must be up so that the entire inline-tool-group is considered up", Computed: true}, "oper_status": schema.StringAttribute{Computed: true}, "release_spare_if_possible": schema.BoolAttribute{MarkdownDescription: "when set to true, if the spare inline tool became active it remains active regardless of the health state of the originally failed primary inline tool", Computed: true}, "spare_inline_tool": schema.StringAttribute{MarkdownDescription: "alias of an inline tool to which traffic is forwarded when the first failure occurs in the set of primary inline tools", Computed: true}}}}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned", Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadAllInlineToolGroupsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadAllInlineToolGroupsDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	d.readRemote(ctx, &config, resp)
	if resp.Diagnostics.HasError() {
		return
	}
	if config.Page.IsNull() {
		resp.Diagnostics.AddWarning("Single-page result", fmt.Sprintf("This data source reads a single page of a paginated API endpoint and does not aggregate results across pages. The \"page\" argument is unset, so the default page is returned; set it to retrieve a different page."))
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}

// readRemote performs the read HTTP exchange and decodes the response into config. Extracted from Read so the request/response logic is unit-testable without a tfsdk.Config.
func (d *LoadAllInlineToolGroupsDataSource) readRemote(ctx context.Context, config *LoadAllInlineToolGroupsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/inline/toolGroups"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_inline_tool_groups", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	if !config.Page.IsNull() {
		query.Set("page", config.Page.ValueString())
	}
	if !config.Sort.IsNull() {
		query.Set("sort", config.Sort.ValueString())
	}
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_inline_tool_groups", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_inline_tool_groups", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_inline_tool_groups", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_inline_tool_groups", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_inline_tool_groups", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_inline_tool_groups", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_inline_tool_groups", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_inline_tool_groups", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_inline_tool_groups", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_load_all_inline_tool_groups", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_inline_tool_groups", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_inline_tool_groups", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_inline_tool_groups", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadAllInlineToolGroupsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
