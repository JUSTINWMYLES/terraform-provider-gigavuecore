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
	_ datasource.DataSource              = (*GetAllTunnelLogicalGroupsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetAllTunnelLogicalGroupsDataSource)(nil)
)

// GetAllTunnelLogicalGroupsDataSource is the generated Terraform data source implementation.
type GetAllTunnelLogicalGroupsDataSource struct {
	client *client.Client
}

// GetAllTunnelLogicalGroupsDataSourceModel describes the data source state shape.
type GetAllTunnelLogicalGroupsDataSourceModel struct {
	Alias         types.String  `tfsdk:"alias"`
	DestinationId types.String  `tfsdk:"destination_id" json:"destinationId"`
	IsActive      types.String  `tfsdk:"is_active" json:"isActive"`
	Items         types.Dynamic `tfsdk:"items"`
	Page          types.String  `tfsdk:"page"`
	Sort          types.String  `tfsdk:"sort"`
	SourceId      types.String  `tfsdk:"source_id" json:"sourceId"`
	State         types.String  `tfsdk:"state"`
	TunnelId      types.String  `tfsdk:"tunnel_id" json:"tunnelId"`
	TunnelType    types.String  `tfsdk:"tunnel_type" json:"tunnelType"`
}

// NewGetAllTunnelLogicalGroupsDataSource returns a new instance of the generated data source.
func NewGetAllTunnelLogicalGroupsDataSource() datasource.DataSource {
	return &GetAllTunnelLogicalGroupsDataSource{}
}

// Metadata returns the data source type name.
func (d *GetAllTunnelLogicalGroupsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_all_tunnel_logical_groups"
}

// Schema returns the data source schema.
func (d *GetAllTunnelLogicalGroupsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get All Tunnel Logical Groups", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Tunnel Logical Group Alias", Optional: true}, "destination_id": schema.StringAttribute{MarkdownDescription: "Destination Id(ClusterName/HostName) of Decapsulation Tunnel in Tunnel Logical Group", Optional: true}, "is_active": schema.StringAttribute{MarkdownDescription: "Tunnel Logical Group Active/Inactive State. 'true' or 'false'", Optional: true}, "items": schema.DynamicAttribute{Computed: true}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned. Defaults to (1:30) to prevent reading entire DB by accident", Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is DESC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}, "source_id": schema.StringAttribute{MarkdownDescription: "Source Id(ClusterName/HostName) of Encapsulation Tunnel in Tunnel Logical Group", Optional: true}, "state": schema.StringAttribute{MarkdownDescription: "Tunnel Discovery State", Optional: true}, "tunnel_id": schema.StringAttribute{MarkdownDescription: "Tunnel logical group Id used for encapsulation and decapsulation", Optional: true}, "tunnel_type": schema.StringAttribute{MarkdownDescription: "Tunnel Type", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetAllTunnelLogicalGroupsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetAllTunnelLogicalGroupsDataSourceModel
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
func (d *GetAllTunnelLogicalGroupsDataSource) readListRemote(ctx context.Context, config *GetAllTunnelLogicalGroupsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/tunnels/logicalgroups"
	params := url.Values{}
	if !config.Page.IsNull() {
		params.Set("page", config.Page.ValueString())
	}
	if !config.Sort.IsNull() {
		params.Set("sort", config.Sort.ValueString())
	}
	if !config.Alias.IsNull() {
		params.Set("alias", config.Alias.ValueString())
	}
	if !config.TunnelType.IsNull() {
		params.Set("tunnelType", config.TunnelType.ValueString())
	}
	if !config.State.IsNull() {
		params.Set("state", config.State.ValueString())
	}
	if !config.IsActive.IsNull() {
		params.Set("isActive", config.IsActive.ValueString())
	}
	if !config.TunnelId.IsNull() {
		params.Set("tunnelId", config.TunnelId.ValueString())
	}
	if !config.SourceId.IsNull() {
		params.Set("sourceId", config.SourceId.ValueString())
	}
	if !config.DestinationId.IsNull() {
		params.Set("destinationId", config.DestinationId.ValueString())
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_tunnel_logical_groups", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_tunnel_logical_groups", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["tunnelLogicalGroups"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_tunnel_logical_groups", fmt.Sprintf("Could not decode list page: missing %q array", "tunnelLogicalGroups"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_tunnel_logical_groups", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetAllTunnelLogicalGroupsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
