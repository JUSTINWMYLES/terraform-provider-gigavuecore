package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
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
	ClusterId types.String `tfsdk:"cluster_id" json:"clusterId"`
	Items     types.List   `tfsdk:"items"`
	Page      types.String `tfsdk:"page"`
	Sort      types.String `tfsdk:"sort"`
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
	resp.Schema = schema.Schema{MarkdownDescription: "Load all Inline Tool Groups", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Inline Tool Group alias. Unique within a cluster", Computed: true}, "comment": schema.StringAttribute{Computed: true}, "current_state": schema.SingleNestedAttribute{MarkdownDescription: "Inline Tool Group Smart Load Balancing definition", Computed: true, Attributes: map[string]schema.Attribute{"inline_tools": schema.ListAttribute{MarkdownDescription: "Current list of inlineTools", Computed: true, ElementType: types.StringType}, "spare_tool": schema.StringAttribute{MarkdownDescription: "Current spare tool", Computed: true}, "spare_tool_status": schema.StringAttribute{MarkdownDescription: "Current spare tool status", Computed: true}, "switched_inline_tool": schema.StringAttribute{MarkdownDescription: "Current switched inlineTool", Computed: true}}}, "enabled": schema.BoolAttribute{MarkdownDescription: "setting to false is equivalent to forcing the inline tool group failure (useful for taking the inline tool group out of commission for maintenance or other purposes)", Computed: true}, "failover_action": schema.StringAttribute{Computed: true}, "failover_mode": schema.StringAttribute{MarkdownDescription: "the way of handling a failure of an individual member of the inline tool list when no spare inline tool is configured or if the spare inline tool is failed", Computed: true}, "flex_status": schema.StringAttribute{Computed: true}, "flex_traffic_path": schema.StringAttribute{Computed: true}, "hash": schema.StringAttribute{Computed: true}, "health_state": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "health_state_reasons": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"message": schema.StringAttribute{MarkdownDescription: "Read-only. Describes the reason for component's health state", Computed: true}, "severity": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "traffic_health_state_computation_type": schema.StringAttribute{MarkdownDescription: "Traffic Health State Computation Type", Computed: true}}}}, "inline_tools": schema.ListAttribute{MarkdownDescription: "If no spare inline tool configured, list of aliases of inline tools participating in hash-based traffic distribution. If the spare inline tool is configured, list of aliases of primary inline tools to which traffic is forwarded as long as all of them are healthy. The number of inline tools in the list must be between 1 and 64 if the spare inline tool is configured or between 2 and 64 otherwise", Computed: true, ElementType: types.StringType}, "min_group_size": schema.Int64Attribute{MarkdownDescription: " the minimum number of inline tools (the inline tools in the list plus spare if configured) that must be up so that the entire inline-tool-group is considered up", Computed: true}, "oper_status": schema.StringAttribute{Computed: true}, "release_spare_if_possible": schema.BoolAttribute{MarkdownDescription: "when set to true, if the spare inline tool became active it remains active regardless of the health state of the originally failed primary inline tool", Computed: true}, "spare_inline_tool": schema.StringAttribute{MarkdownDescription: "alias of an inline tool to which traffic is forwarded when the first failure occurs in the set of primary inline tools", Computed: true}}}}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned", Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadAllInlineToolGroupsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadAllInlineToolGroupsDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	d.readListRemote(ctx, &config, resp)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}

// readListRemote performs the paginated read HTTP exchange and decodes the response array into config. Extracted from Read so the request/response logic is unit-testable without a tfsdk.Config.
func (d *LoadAllInlineToolGroupsDataSource) readListRemote(ctx context.Context, config *LoadAllInlineToolGroupsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/inline/toolGroups"
	params := url.Values{}
	params.Set("clusterId", config.ClusterId.ValueString())
	if !config.Page.IsNull() {
		params.Set("page", config.Page.ValueString())
	}
	if !config.Sort.IsNull() {
		params.Set("sort", config.Sort.ValueString())
	}
	var nextURL string
	fetch := func(ctx context.Context, p url.Values) (*http.Response, error) {
		httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
		if err != nil {
			return nil, err
		}
		if nextURL != "" {
			parsed, perr := url.Parse(nextURL)
			if perr != nil {
				return nil, perr
			}
			httpReq.URL = parsed
		} else {
			httpReq.URL.RawQuery = p.Encode()
		}
		return d.client.Do(httpReq)
	}
	pages, err := client.ListAllPages(ctx, params, fetch, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_inline_tool_groups", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_inline_tool_groups", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["inlineToolGroups"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_inline_tool_groups", fmt.Sprintf("Could not decode list page: missing %q array", "inlineToolGroups"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
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
