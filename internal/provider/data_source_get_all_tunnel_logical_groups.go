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
	_ datasource.DataSource              = (*GetAllTunnelLogicalGroupsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetAllTunnelLogicalGroupsDataSource)(nil)
)

// GetAllTunnelLogicalGroupsDataSource is the generated Terraform data source implementation.
type GetAllTunnelLogicalGroupsDataSource struct {
	client *client.Client
}

// GetAllTunnelLogicalGroupsDataSourceModel describes the data source state shape.
type GetAllTunnelLogicalGroupsDataSourceModel struct {
	Alias               types.String  `tfsdk:"alias"`
	Context             types.Object  `tfsdk:"context"`
	DestinationId       types.String  `tfsdk:"destination_id" json:"destinationId"`
	IsActive            types.String  `tfsdk:"is_active" json:"isActive"`
	Page                types.String  `tfsdk:"page"`
	Sort                types.String  `tfsdk:"sort"`
	SourceId            types.String  `tfsdk:"source_id" json:"sourceId"`
	State               types.String  `tfsdk:"state"`
	TunnelId            types.String  `tfsdk:"tunnel_id" json:"tunnelId"`
	TunnelLogicalGroups types.Dynamic `tfsdk:"tunnel_logical_groups" json:"tunnelLogicalGroups"`
	TunnelType          types.String  `tfsdk:"tunnel_type" json:"tunnelType"`
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
	resp.Schema = schema.Schema{MarkdownDescription: "Get All Tunnel Logical Groups", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{Optional: true}, "context": schema.SingleNestedAttribute{MarkdownDescription: "Gigamon query result context", Computed: true, Attributes: map[string]schema.Attribute{"page_no": schema.Int64Attribute{MarkdownDescription: "page number of the returned result set", Computed: true}, "page_size": schema.Int64Attribute{MarkdownDescription: "page size of the returned result set", Computed: true}, "sort": schema.ListAttribute{MarkdownDescription: "sorting info of the returned result set. list of fields in the array indicate sorting order", Computed: true, ElementType: types.StringType}, "total_items": schema.Int64Attribute{MarkdownDescription: "total number of items in the queried entity type", Computed: true}}}, "destination_id": schema.StringAttribute{Optional: true}, "is_active": schema.StringAttribute{Optional: true}, "page": schema.StringAttribute{Optional: true}, "sort": schema.StringAttribute{Optional: true}, "source_id": schema.StringAttribute{Optional: true}, "state": schema.StringAttribute{Optional: true}, "tunnel_id": schema.StringAttribute{Optional: true}, "tunnel_logical_groups": schema.DynamicAttribute{Computed: true}, "tunnel_type": schema.StringAttribute{Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetAllTunnelLogicalGroupsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetAllTunnelLogicalGroupsDataSourceModel
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
func (d *GetAllTunnelLogicalGroupsDataSource) readRemote(ctx context.Context, config *GetAllTunnelLogicalGroupsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/tunnels/logicalgroups"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_tunnel_logical_groups", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !config.Page.IsNull() {
		query.Set("page", config.Page.ValueString())
	}
	if !config.Sort.IsNull() {
		query.Set("sort", config.Sort.ValueString())
	}
	if !config.Alias.IsNull() {
		query.Set("alias", config.Alias.ValueString())
	}
	if !config.TunnelType.IsNull() {
		query.Set("tunnelType", config.TunnelType.ValueString())
	}
	if !config.State.IsNull() {
		query.Set("state", config.State.ValueString())
	}
	if !config.IsActive.IsNull() {
		query.Set("isActive", config.IsActive.ValueString())
	}
	if !config.TunnelId.IsNull() {
		query.Set("tunnelId", config.TunnelId.ValueString())
	}
	if !config.SourceId.IsNull() {
		query.Set("sourceId", config.SourceId.ValueString())
	}
	if !config.DestinationId.IsNull() {
		query.Set("destinationId", config.DestinationId.ValueString())
	}
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_tunnel_logical_groups", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_tunnel_logical_groups", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_tunnel_logical_groups", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_tunnel_logical_groups", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_tunnel_logical_groups", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_tunnel_logical_groups", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_tunnel_logical_groups", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_tunnel_logical_groups", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_tunnel_logical_groups", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_all_tunnel_logical_groups", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_tunnel_logical_groups", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_tunnel_logical_groups", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
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
