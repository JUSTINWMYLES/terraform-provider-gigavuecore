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
	_ datasource.DataSource              = (*LoadAllMetadataApplicationProfilesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadAllMetadataApplicationProfilesDataSource)(nil)
)

// LoadAllMetadataApplicationProfilesDataSource is the generated Terraform data source implementation.
type LoadAllMetadataApplicationProfilesDataSource struct {
	client *client.Client
}

// LoadAllMetadataApplicationProfilesDataSourceModel describes the data source state shape.
type LoadAllMetadataApplicationProfilesDataSourceModel struct {
	ApplicationProfiles types.List   `tfsdk:"application_profiles" json:"applicationProfiles"`
	Context             types.Object `tfsdk:"context"`
	Page                types.String `tfsdk:"page"`
	Sort                types.String `tfsdk:"sort"`
}

// NewLoadAllMetadataApplicationProfilesDataSource returns a new instance of the generated data source.
func NewLoadAllMetadataApplicationProfilesDataSource() datasource.DataSource {
	return &LoadAllMetadataApplicationProfilesDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadAllMetadataApplicationProfilesDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_all_metadata_application_profiles"
}

// Schema returns the data source schema.
func (d *LoadAllMetadataApplicationProfilesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load All Metadata Application Profiles", Attributes: map[string]schema.Attribute{"application_profiles": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "application profile alias", Computed: true}, "application_id": schema.BoolAttribute{MarkdownDescription: "only valid with 'export' type", Computed: true}, "applications": schema.ListNestedAttribute{MarkdownDescription: "application and attributes.", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"attributes": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"name": schema.StringAttribute{MarkdownDescription: "attribute name", Computed: true}, "value": schema.StringAttribute{MarkdownDescription: "application's attribute value", Computed: true}}}}, "is_user_defined": schema.BoolAttribute{MarkdownDescription: "Default value is false. Must be set to true only for App Intel solution while configuring user defined apps", Computed: true}, "name": schema.StringAttribute{MarkdownDescription: "application name", Computed: true}}}}, "counter": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"bytes": schema.BoolAttribute{Computed: true}, "bytes_long": schema.BoolAttribute{Computed: true}, "inner_byte": schema.BoolAttribute{Computed: true}, "inner_byte_long": schema.BoolAttribute{Computed: true}, "packets": schema.BoolAttribute{Computed: true}, "packets_long": schema.BoolAttribute{Computed: true}}}, "datalink": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"mac_dst": schema.BoolAttribute{Computed: true}, "mac_src": schema.BoolAttribute{Computed: true}, "vlan": schema.BoolAttribute{Computed: true}}}, "description": schema.StringAttribute{Computed: true}, "flow": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"end_reason": schema.BoolAttribute{Computed: true}}}, "gtpu": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"qfi": schema.BoolAttribute{Computed: true}, "teid": schema.BoolAttribute{Computed: true}}}, "interface": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"in_name_width": schema.Int64Attribute{Computed: true}, "in_physical_width": schema.Int64Attribute{Computed: true}, "out_physical_width": schema.Int64Attribute{Computed: true}}}, "ip": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"version": schema.BoolAttribute{Computed: true}}}, "ipv4": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"destination": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"prefix_min_mask": schema.StringAttribute{Computed: true}}}, "dscp": schema.BoolAttribute{Computed: true}, "fragmentation": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"flags": schema.BoolAttribute{Computed: true}, "offset": schema.BoolAttribute{Computed: true}}}, "header_len": schema.BoolAttribute{Computed: true}, "option_map": schema.BoolAttribute{Computed: true}, "precedence": schema.BoolAttribute{Computed: true}, "protocol": schema.BoolAttribute{Computed: true}, "section": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"header_size": schema.Int64Attribute{Computed: true}, "payload_size": schema.Int64Attribute{Computed: true}}}, "source": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"prefix_min_mask": schema.StringAttribute{MarkdownDescription: "ipv4 source prefix minimum-mask - netmask or mask length", Computed: true}}}, "tos": schema.BoolAttribute{Computed: true}, "total_length": schema.BoolAttribute{Computed: true}, "ttl": schema.BoolAttribute{Computed: true}}}, "ipv6": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"destination": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"prefix_min_mask": schema.StringAttribute{Computed: true}}}, "dscp": schema.BoolAttribute{Computed: true}, "extension_map": schema.BoolAttribute{Computed: true}, "flow_label": schema.BoolAttribute{Computed: true}, "fragmentation": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"flags": schema.BoolAttribute{Computed: true}, "offset": schema.BoolAttribute{Computed: true}}}, "hop_limit": schema.BoolAttribute{Computed: true}, "length": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"header": schema.BoolAttribute{Computed: true}, "payload": schema.BoolAttribute{Computed: true}, "total": schema.BoolAttribute{Computed: true}}}, "next_header": schema.BoolAttribute{Computed: true}, "precedence": schema.BoolAttribute{Computed: true}, "section": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"header_size": schema.Int64Attribute{Computed: true}, "payload_size": schema.Int64Attribute{Computed: true}}}, "source": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"prefix_min_mask": schema.StringAttribute{Computed: true}}}, "traffic_class": schema.BoolAttribute{Computed: true}}}, "outer_ipv4": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"destination": schema.BoolAttribute{Computed: true}, "source": schema.BoolAttribute{Computed: true}}}, "outer_ipv6": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"destination": schema.BoolAttribute{Computed: true}, "source": schema.BoolAttribute{Computed: true}}}, "timestamp": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"flow_end_msec": schema.BoolAttribute{Computed: true}, "flow_endsec": schema.BoolAttribute{Computed: true}, "flow_start_msec": schema.BoolAttribute{Computed: true}, "flow_startsec": schema.BoolAttribute{Computed: true}, "sys_up_time_first": schema.BoolAttribute{Computed: true}, "sys_up_time_last": schema.BoolAttribute{Computed: true}}}, "transport": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"dst_port": schema.BoolAttribute{Computed: true}, "icmp": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"ipv4_code": schema.BoolAttribute{Computed: true}, "ipv4_type": schema.BoolAttribute{Computed: true}, "ipv6_code": schema.BoolAttribute{Computed: true}, "ipv6_type": schema.BoolAttribute{Computed: true}}}, "src_port": schema.BoolAttribute{Computed: true}, "tcp": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"ack_number": schema.BoolAttribute{Computed: true}, "dst_port": schema.BoolAttribute{Computed: true}, "flags": schema.BoolAttribute{Computed: true}, "header_len": schema.BoolAttribute{Computed: true}, "seq_number": schema.BoolAttribute{Computed: true}, "src_port": schema.BoolAttribute{Computed: true}, "urgent_ptr": schema.BoolAttribute{Computed: true}, "window_size": schema.BoolAttribute{Computed: true}}}, "udp": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"dst_port": schema.BoolAttribute{Computed: true}, "msg_len": schema.BoolAttribute{Computed: true}, "src_port": schema.BoolAttribute{Computed: true}}}}}, "type": schema.StringAttribute{Computed: true}}}}, "context": schema.SingleNestedAttribute{MarkdownDescription: "Gigamon query result context", Computed: true, Attributes: map[string]schema.Attribute{"page_no": schema.Int64Attribute{MarkdownDescription: "page number of the returned result set", Computed: true}, "page_size": schema.Int64Attribute{MarkdownDescription: "page size of the returned result set", Computed: true}, "sort": schema.ListAttribute{MarkdownDescription: "sorting info of the returned result set. list of fields in the array indicate sorting order", Computed: true, ElementType: types.StringType}, "total_items": schema.Int64Attribute{MarkdownDescription: "total number of items in the queried entity type", Computed: true}}}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned", Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadAllMetadataApplicationProfilesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadAllMetadataApplicationProfilesDataSourceModel
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
func (d *LoadAllMetadataApplicationProfilesDataSource) readRemote(ctx context.Context, config *LoadAllMetadataApplicationProfilesDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/apps/metadata/applicationProfiles"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_metadata_application_profiles", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !config.Page.IsNull() {
		query.Set("page", config.Page.ValueString())
	}
	if !config.Sort.IsNull() {
		query.Set("sort", config.Sort.ValueString())
	}
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_metadata_application_profiles", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_metadata_application_profiles", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_metadata_application_profiles", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_metadata_application_profiles", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_metadata_application_profiles", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_metadata_application_profiles", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_metadata_application_profiles", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_metadata_application_profiles", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_metadata_application_profiles", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_load_all_metadata_application_profiles", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_metadata_application_profiles", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_metadata_application_profiles", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_metadata_application_profiles", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadAllMetadataApplicationProfilesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
