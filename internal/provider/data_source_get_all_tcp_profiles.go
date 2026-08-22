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
	_ datasource.DataSource              = (*GetAllTcpProfilesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetAllTcpProfilesDataSource)(nil)
)

// GetAllTcpProfilesDataSource is the generated Terraform data source implementation.
type GetAllTcpProfilesDataSource struct {
	client *client.Client
}

// GetAllTcpProfilesDataSourceModel describes the data source state shape.
type GetAllTcpProfilesDataSourceModel struct {
	AppsTcpProfiles types.List   `tfsdk:"apps_tcp_profiles" json:"appsTcpProfiles"`
	ClusterId       types.String `tfsdk:"cluster_id" json:"clusterId"`
	Context         types.Object `tfsdk:"context"`
}

// NewGetAllTcpProfilesDataSource returns a new instance of the generated data source.
func NewGetAllTcpProfilesDataSource() datasource.DataSource {
	return &GetAllTcpProfilesDataSource{}
}

// Metadata returns the data source type name.
func (d *GetAllTcpProfilesDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_all_tcp_profiles"
}

// Schema returns the data source schema.
func (d *GetAllTcpProfilesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get all Apps TCP Profile", Attributes: map[string]schema.Attribute{"apps_tcp_profiles": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{Computed: true}, "keep_alive_timer": schema.Int64Attribute{Computed: true}, "selective_ack": schema.StringAttribute{Computed: true}, "syn_retries": schema.Int64Attribute{Computed: true}}}}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "context": schema.SingleNestedAttribute{MarkdownDescription: "Gigamon query result context", Computed: true, Attributes: map[string]schema.Attribute{"page_no": schema.Int64Attribute{MarkdownDescription: "page number of the returned result set", Computed: true}, "page_size": schema.Int64Attribute{MarkdownDescription: "page size of the returned result set", Computed: true}, "sort": schema.ListAttribute{MarkdownDescription: "sorting info of the returned result set. list of fields in the array indicate sorting order", Computed: true, ElementType: types.StringType}, "total_items": schema.Int64Attribute{MarkdownDescription: "total number of items in the queried entity type", Computed: true}}}}}
}

// Read fetches remote state into the data source model.
func (d *GetAllTcpProfilesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetAllTcpProfilesDataSourceModel
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
func (d *GetAllTcpProfilesDataSource) readRemote(ctx context.Context, config *GetAllTcpProfilesDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/apps/tcp/profile"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_tcp_profiles", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_tcp_profiles", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_tcp_profiles", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_tcp_profiles", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_tcp_profiles", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_tcp_profiles", "Entity Not Found. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_all_tcp_profiles", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_tcp_profiles", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_tcp_profiles", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_tcp_profiles", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetAllTcpProfilesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
