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
	_ datasource.DataSource              = (*LoadAllEnhancedSlicingsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadAllEnhancedSlicingsDataSource)(nil)
)

// LoadAllEnhancedSlicingsDataSource is the generated Terraform data source implementation.
type LoadAllEnhancedSlicingsDataSource struct {
	client *client.Client
}

// LoadAllEnhancedSlicingsDataSourceModel describes the data source state shape.
type LoadAllEnhancedSlicingsDataSourceModel struct {
	Context          types.Object `tfsdk:"context"`
	EnhancedSlicings types.List   `tfsdk:"enhanced_slicings" json:"enhancedSlicings"`
	Page             types.String `tfsdk:"page"`
	Sort             types.String `tfsdk:"sort"`
}

// NewLoadAllEnhancedSlicingsDataSource returns a new instance of the generated data source.
func NewLoadAllEnhancedSlicingsDataSource() datasource.DataSource {
	return &LoadAllEnhancedSlicingsDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadAllEnhancedSlicingsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_all_enhanced_slicings"
}

// Schema returns the data source schema.
func (d *LoadAllEnhancedSlicingsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "new in version H 5.7", Attributes: map[string]schema.Attribute{"context": schema.SingleNestedAttribute{MarkdownDescription: "Gigamon query result context", Computed: true, Attributes: map[string]schema.Attribute{"page_no": schema.Int64Attribute{MarkdownDescription: "page number of the returned result set", Computed: true}, "page_size": schema.Int64Attribute{MarkdownDescription: "page size of the returned result set", Computed: true}, "sort": schema.ListAttribute{MarkdownDescription: "sorting info of the returned result set. list of fields in the array indicate sorting order", Computed: true, ElementType: types.StringType}, "total_items": schema.Int64Attribute{MarkdownDescription: "total number of items in the queried entity type", Computed: true}}}, "enhanced_slicings": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "apps enhanced slicing alias", Computed: true}, "hash_field": schema.StringAttribute{MarkdownDescription: "Hash field for session", Computed: true}, "max_sessions": schema.Int64Attribute{MarkdownDescription: "Maximum number of session entries (in millions). Used only for flow-session option", Computed: true}, "protocol_fields": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"gtp": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"flow_session": schema.SingleNestedAttribute{MarkdownDescription: "Always use 4-tuple (IP, L4 port) to identify a flow. Used only if user wants to start slicing after some packet count from the beginning of a session flow", Computed: true, Attributes: map[string]schema.Attribute{"action": schema.StringAttribute{MarkdownDescription: "Slice or drop packets after skip count", Computed: true}, "skip_pkt_count": schema.Int64Attribute{MarkdownDescription: "Start slicing after 'packet-count' is reached. Only valid with 'flowSession'", Computed: true}, "timeout": schema.Int64Attribute{MarkdownDescription: "Session inactivity timeout in seconds", Computed: true}, "value": schema.StringAttribute{MarkdownDescription: "outer: first occurrence of IP and L4 port ; inner: second occurrence of IP and L4 port", Computed: true}}}, "offset": schema.Int64Attribute{MarkdownDescription: "Number of bytes after 'protocol' header. 64-9000 if 'protocol' is none, 0-9000 for all protocols", Computed: true}, "protocol": schema.StringAttribute{Computed: true}}}, "gtp_u": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"flow_session": schema.SingleNestedAttribute{MarkdownDescription: "Always use 4-tuple (IP, L4 port) to identify a flow. Used only if user wants to start slicing after some packet count from the beginning of a session flow", Computed: true, Attributes: map[string]schema.Attribute{"action": schema.StringAttribute{MarkdownDescription: "Slice or drop packets after skip count", Computed: true}, "skip_pkt_count": schema.Int64Attribute{MarkdownDescription: "Start slicing after 'packet-count' is reached. Only valid with 'flowSession'", Computed: true}, "timeout": schema.Int64Attribute{MarkdownDescription: "Session inactivity timeout in seconds", Computed: true}, "value": schema.StringAttribute{MarkdownDescription: "outer: first occurrence of IP and L4 port ; inner: second occurrence of IP and L4 port", Computed: true}}}, "l4_port": schema.Int64Attribute{MarkdownDescription: "Layer 4 Port Number", Computed: true}, "offset": schema.Int64Attribute{MarkdownDescription: "Number of bytes after 'protocol' header. 64-9000 if 'protocol' is none, 0-9000 for all protocols", Computed: true}, "protocol": schema.StringAttribute{Computed: true}}}, "ip": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"flow_session": schema.SingleNestedAttribute{MarkdownDescription: "Always use 4-tuple (IP, L4 port) to identify a flow. Used only if user wants to start slicing after some packet count from the beginning of a session flow", Computed: true, Attributes: map[string]schema.Attribute{"action": schema.StringAttribute{MarkdownDescription: "Slice or drop packets after skip count", Computed: true}, "skip_pkt_count": schema.Int64Attribute{MarkdownDescription: "Start slicing after 'packet-count' is reached. Only valid with 'flowSession'", Computed: true}, "timeout": schema.Int64Attribute{MarkdownDescription: "Session inactivity timeout in seconds", Computed: true}, "value": schema.StringAttribute{MarkdownDescription: "outer: first occurrence of IP and L4 port ; inner: second occurrence of IP and L4 port", Computed: true}}}, "offset": schema.Int64Attribute{MarkdownDescription: "Number of bytes after 'protocol' header. 64-9000 if 'protocol' is none, 0-9000 for all protocols", Computed: true}, "protocol": schema.StringAttribute{Computed: true}, "value": schema.StringAttribute{Computed: true}}}, "none": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"flow_session": schema.SingleNestedAttribute{MarkdownDescription: "Always use 4-tuple (IP, L4 port) to identify a flow. Used only if user wants to start slicing after some packet count from the beginning of a session flow", Computed: true, Attributes: map[string]schema.Attribute{"action": schema.StringAttribute{MarkdownDescription: "Slice or drop packets after skip count", Computed: true}, "skip_pkt_count": schema.Int64Attribute{MarkdownDescription: "Start slicing after 'packet-count' is reached. Only valid with 'flowSession'", Computed: true}, "timeout": schema.Int64Attribute{MarkdownDescription: "Session inactivity timeout in seconds", Computed: true}, "value": schema.StringAttribute{MarkdownDescription: "outer: first occurrence of IP and L4 port ; inner: second occurrence of IP and L4 port", Computed: true}}}, "offset": schema.Int64Attribute{MarkdownDescription: "Number of bytes after 'protocol' header. 64-9000 if 'protocol' is none, 0-9000 for all protocols", Computed: true}}}, "transport": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"flow_session": schema.SingleNestedAttribute{MarkdownDescription: "Always use 4-tuple (IP, L4 port) to identify a flow. Used only if user wants to start slicing after some packet count from the beginning of a session flow", Computed: true, Attributes: map[string]schema.Attribute{"action": schema.StringAttribute{MarkdownDescription: "Slice or drop packets after skip count", Computed: true}, "skip_pkt_count": schema.Int64Attribute{MarkdownDescription: "Start slicing after 'packet-count' is reached. Only valid with 'flowSession'", Computed: true}, "timeout": schema.Int64Attribute{MarkdownDescription: "Session inactivity timeout in seconds", Computed: true}, "value": schema.StringAttribute{MarkdownDescription: "outer: first occurrence of IP and L4 port ; inner: second occurrence of IP and L4 port", Computed: true}}}, "l4_port": schema.Int64Attribute{MarkdownDescription: "Layer 4 Port Number", Computed: true}, "offset": schema.Int64Attribute{MarkdownDescription: "Number of bytes after 'protocol' header. 64-9000 if 'protocol' is none, 0-9000 for all protocols", Computed: true}, "protocol": schema.StringAttribute{Computed: true}, "value": schema.StringAttribute{Computed: true}}}}}}}}}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned", Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadAllEnhancedSlicingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadAllEnhancedSlicingsDataSourceModel
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
func (d *LoadAllEnhancedSlicingsDataSource) readRemote(ctx context.Context, config *LoadAllEnhancedSlicingsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/apps/enhancedSlicing"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_enhanced_slicings", fmt.Sprintf("Could not build request: %s", err))
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
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_enhanced_slicings", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_enhanced_slicings", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_enhanced_slicings", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_enhanced_slicings", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_enhanced_slicings", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_enhanced_slicings", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_enhanced_slicings", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_enhanced_slicings", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_enhanced_slicings", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_load_all_enhanced_slicings", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_enhanced_slicings", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_enhanced_slicings", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_enhanced_slicings", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadAllEnhancedSlicingsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
