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
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
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
	Items types.List   `tfsdk:"items"`
	Page  types.String `tfsdk:"page"`
	Sort  types.String `tfsdk:"sort"`
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
	resp.Schema = schema.Schema{MarkdownDescription: "new in version H 5.7", Attributes: map[string]schema.Attribute{"items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "apps enhanced slicing alias", Computed: true}, "hash_field": schema.StringAttribute{MarkdownDescription: "Hash field for session", Computed: true}, "max_sessions": schema.Int64Attribute{MarkdownDescription: "Maximum number of session entries (in millions). Used only for flow-session option", Computed: true}, "protocol_fields": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"gtp": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"flow_session": schema.SingleNestedAttribute{MarkdownDescription: "Always use 4-tuple (IP, L4 port) to identify a flow. Used only if user wants to start slicing after some packet count from the beginning of a session flow", Computed: true, Attributes: map[string]schema.Attribute{"action": schema.StringAttribute{MarkdownDescription: "Slice or drop packets after skip count", Computed: true}, "skip_pkt_count": schema.Int64Attribute{MarkdownDescription: "Start slicing after 'packet-count' is reached. Only valid with 'flowSession'", Computed: true}, "timeout": schema.Int64Attribute{MarkdownDescription: "Session inactivity timeout in seconds", Computed: true}, "value": schema.StringAttribute{MarkdownDescription: "outer: first occurrence of IP and L4 port ; inner: second occurrence of IP and L4 port", Computed: true}}}, "offset": schema.Int64Attribute{MarkdownDescription: "Number of bytes after 'protocol' header. 64-9000 if 'protocol' is none, 0-9000 for all protocols", Computed: true}, "protocol": schema.StringAttribute{Computed: true}}}, "gtp_u": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"flow_session": schema.SingleNestedAttribute{MarkdownDescription: "Always use 4-tuple (IP, L4 port) to identify a flow. Used only if user wants to start slicing after some packet count from the beginning of a session flow", Computed: true, Attributes: map[string]schema.Attribute{"action": schema.StringAttribute{MarkdownDescription: "Slice or drop packets after skip count", Computed: true}, "skip_pkt_count": schema.Int64Attribute{MarkdownDescription: "Start slicing after 'packet-count' is reached. Only valid with 'flowSession'", Computed: true}, "timeout": schema.Int64Attribute{MarkdownDescription: "Session inactivity timeout in seconds", Computed: true}, "value": schema.StringAttribute{MarkdownDescription: "outer: first occurrence of IP and L4 port ; inner: second occurrence of IP and L4 port", Computed: true}}}, "l4_port": schema.Int64Attribute{MarkdownDescription: "Layer 4 Port Number", Computed: true}, "offset": schema.Int64Attribute{MarkdownDescription: "Number of bytes after 'protocol' header. 64-9000 if 'protocol' is none, 0-9000 for all protocols", Computed: true}, "protocol": schema.StringAttribute{Computed: true}}}, "ip": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"flow_session": schema.SingleNestedAttribute{MarkdownDescription: "Always use 4-tuple (IP, L4 port) to identify a flow. Used only if user wants to start slicing after some packet count from the beginning of a session flow", Computed: true, Attributes: map[string]schema.Attribute{"action": schema.StringAttribute{MarkdownDescription: "Slice or drop packets after skip count", Computed: true}, "skip_pkt_count": schema.Int64Attribute{MarkdownDescription: "Start slicing after 'packet-count' is reached. Only valid with 'flowSession'", Computed: true}, "timeout": schema.Int64Attribute{MarkdownDescription: "Session inactivity timeout in seconds", Computed: true}, "value": schema.StringAttribute{MarkdownDescription: "outer: first occurrence of IP and L4 port ; inner: second occurrence of IP and L4 port", Computed: true}}}, "offset": schema.Int64Attribute{MarkdownDescription: "Number of bytes after 'protocol' header. 64-9000 if 'protocol' is none, 0-9000 for all protocols", Computed: true}, "protocol": schema.StringAttribute{Computed: true}, "value": schema.StringAttribute{Computed: true}}}, "none": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"flow_session": schema.SingleNestedAttribute{MarkdownDescription: "Always use 4-tuple (IP, L4 port) to identify a flow. Used only if user wants to start slicing after some packet count from the beginning of a session flow", Computed: true, Attributes: map[string]schema.Attribute{"action": schema.StringAttribute{MarkdownDescription: "Slice or drop packets after skip count", Computed: true}, "skip_pkt_count": schema.Int64Attribute{MarkdownDescription: "Start slicing after 'packet-count' is reached. Only valid with 'flowSession'", Computed: true}, "timeout": schema.Int64Attribute{MarkdownDescription: "Session inactivity timeout in seconds", Computed: true}, "value": schema.StringAttribute{MarkdownDescription: "outer: first occurrence of IP and L4 port ; inner: second occurrence of IP and L4 port", Computed: true}}}, "offset": schema.Int64Attribute{MarkdownDescription: "Number of bytes after 'protocol' header. 64-9000 if 'protocol' is none, 0-9000 for all protocols", Computed: true}}}, "transport": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"flow_session": schema.SingleNestedAttribute{MarkdownDescription: "Always use 4-tuple (IP, L4 port) to identify a flow. Used only if user wants to start slicing after some packet count from the beginning of a session flow", Computed: true, Attributes: map[string]schema.Attribute{"action": schema.StringAttribute{MarkdownDescription: "Slice or drop packets after skip count", Computed: true}, "skip_pkt_count": schema.Int64Attribute{MarkdownDescription: "Start slicing after 'packet-count' is reached. Only valid with 'flowSession'", Computed: true}, "timeout": schema.Int64Attribute{MarkdownDescription: "Session inactivity timeout in seconds", Computed: true}, "value": schema.StringAttribute{MarkdownDescription: "outer: first occurrence of IP and L4 port ; inner: second occurrence of IP and L4 port", Computed: true}}}, "l4_port": schema.Int64Attribute{MarkdownDescription: "Layer 4 Port Number", Computed: true}, "offset": schema.Int64Attribute{MarkdownDescription: "Number of bytes after 'protocol' header. 64-9000 if 'protocol' is none, 0-9000 for all protocols", Computed: true}, "protocol": schema.StringAttribute{Computed: true}, "value": schema.StringAttribute{Computed: true}}}}}}}}}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned", Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadAllEnhancedSlicingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadAllEnhancedSlicingsDataSourceModel
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
func (d *LoadAllEnhancedSlicingsDataSource) readListRemote(ctx context.Context, config *LoadAllEnhancedSlicingsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/apps/enhancedSlicing"
	params := url.Values{}
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
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_enhanced_slicings", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_enhanced_slicings", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["enhancedSlicings"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_enhanced_slicings", fmt.Sprintf("Could not decode list page: missing %q array", "enhancedSlicings"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
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
