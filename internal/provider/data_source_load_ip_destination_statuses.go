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
	_ datasource.DataSource              = (*LoadIpDestinationStatusesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadIpDestinationStatusesDataSource)(nil)
)

// LoadIpDestinationStatusesDataSource is the generated Terraform data source implementation.
type LoadIpDestinationStatusesDataSource struct {
	client *client.Client
}

// LoadIpDestinationStatusesDataSourceModel describes the data source state shape.
type LoadIpDestinationStatusesDataSourceModel struct {
	ClusterId        types.String `tfsdk:"cluster_id" json:"clusterId"`
	Context          types.Object `tfsdk:"context"`
	GsGroupAlias     types.String `tfsdk:"gs_group_alias" json:"gsGroupAlias"`
	IpInterfaceAlias types.String `tfsdk:"ip_interface_alias" json:"ipInterfaceAlias"`
	IpInterfaces     types.List   `tfsdk:"ip_interfaces" json:"ipInterfaces"`
	IpType           types.String `tfsdk:"ip_type" json:"ipType"`
}

// NewLoadIpDestinationStatusesDataSource returns a new instance of the generated data source.
func NewLoadIpDestinationStatusesDataSource() datasource.DataSource {
	return &LoadIpDestinationStatusesDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadIpDestinationStatusesDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_ip_destination_statuses"
}

// Schema returns the data source schema.
func (d *LoadIpDestinationStatusesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load IP destination status of all or those matching the IP Interface alias or IP type or GsGroup", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "context": schema.SingleNestedAttribute{MarkdownDescription: "Gigamon query result context", Computed: true, Attributes: map[string]schema.Attribute{"page_no": schema.Int64Attribute{MarkdownDescription: "page number of the returned result set", Computed: true}, "page_size": schema.Int64Attribute{MarkdownDescription: "page size of the returned result set", Computed: true}, "sort": schema.ListAttribute{MarkdownDescription: "sorting info of the returned result set. list of fields in the array indicate sorting order", Computed: true, ElementType: types.StringType}, "total_items": schema.Int64Attribute{MarkdownDescription: "total number of items in the queried entity type", Computed: true}}}, "gs_group_alias": schema.StringAttribute{MarkdownDescription: "GsGroup alias", Optional: true}, "ip_interface_alias": schema.StringAttribute{MarkdownDescription: "IP Interface alias", Optional: true}, "ip_interfaces": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"destination": schema.StringAttribute{MarkdownDescription: "IP Destination", Computed: true}, "gs_group": schema.StringAttribute{MarkdownDescription: "Associated GsGroup", Computed: true}, "interface_alias": schema.StringAttribute{MarkdownDescription: "Interface Alias", Computed: true}, "status": schema.StringAttribute{MarkdownDescription: "Status", Computed: true}, "te_id": schema.StringAttribute{MarkdownDescription: "Te-ID", Computed: true}}}}, "ip_type": schema.StringAttribute{MarkdownDescription: "IP type", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadIpDestinationStatusesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadIpDestinationStatusesDataSourceModel
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
func (d *LoadIpDestinationStatusesDataSource) readRemote(ctx context.Context, config *LoadIpDestinationStatusesDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/ip/destination"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_ip_destination_statuses", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	if !config.IpInterfaceAlias.IsNull() {
		query.Set("ipInterfaceAlias", config.IpInterfaceAlias.ValueString())
	}
	if !config.GsGroupAlias.IsNull() {
		query.Set("gsGroupAlias", config.GsGroupAlias.ValueString())
	}
	if !config.IpType.IsNull() {
		query.Set("ipType", config.IpType.ValueString())
	}
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_ip_destination_statuses", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_ip_destination_statuses", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_ip_destination_statuses", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_ip_destination_statuses", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_ip_destination_statuses", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_ip_destination_statuses", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_ip_destination_statuses", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_ip_destination_statuses", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_ip_destination_statuses", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_load_ip_destination_statuses", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_load_ip_destination_statuses", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_ip_destination_statuses", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_ip_destination_statuses", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadIpDestinationStatusesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
