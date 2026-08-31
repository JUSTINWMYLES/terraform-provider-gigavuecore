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
	_ datasource.DataSource              = (*LoadAllMetadataCacheDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadAllMetadataCacheDataSource)(nil)
)

// LoadAllMetadataCacheDataSource is the generated Terraform data source implementation.
type LoadAllMetadataCacheDataSource struct {
	client *client.Client
}

// LoadAllMetadataCacheDataSourceModel describes the data source state shape.
type LoadAllMetadataCacheDataSourceModel struct {
	Items types.List   `tfsdk:"items"`
	Page  types.String `tfsdk:"page"`
	Sort  types.String `tfsdk:"sort"`
}

// NewLoadAllMetadataCacheDataSource returns a new instance of the generated data source.
func NewLoadAllMetadataCacheDataSource() datasource.DataSource {
	return &LoadAllMetadataCacheDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadAllMetadataCacheDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_all_metadata_cache"
}

// Schema returns the data source schema.
func (d *LoadAllMetadataCacheDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load All Metadata Cache", Attributes: map[string]schema.Attribute{"items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"advance_hash": schema.BoolAttribute{MarkdownDescription: "When scheduling flows for processing, use external encapsulate packet header if disable, use inner packet header if enable.", Computed: true}, "alias": schema.StringAttribute{Computed: true}, "description": schema.StringAttribute{Computed: true}, "dpi_inject_limit": schema.Int64Attribute{Computed: true}, "event": schema.StringAttribute{Computed: true}, "exporters": schema.ListAttribute{MarkdownDescription: "alias of metadata exporters to attach this cache", Computed: true, ElementType: types.StringType}, "flow_behavior": schema.StringAttribute{MarkdownDescription: "direction for flow identification", Computed: true}, "match": schema.SingleNestedAttribute{MarkdownDescription: "match criteria for record generation", Computed: true, Attributes: map[string]schema.Attribute{"datalink": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"mac_dst": schema.BoolAttribute{Computed: true}, "mac_src": schema.BoolAttribute{Computed: true}, "vlan": schema.BoolAttribute{Computed: true}}}, "interface": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"in_name_width": schema.Int64Attribute{Computed: true}, "in_physical_width": schema.Int64Attribute{Computed: true}}}, "ip": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"version": schema.BoolAttribute{Computed: true}}}, "ipv4": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"destination": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"prefix_min_mask": schema.StringAttribute{MarkdownDescription: "ipv4 destination prefix minimum-mask - netmask or mask length", Computed: true}}}, "dscp": schema.BoolAttribute{Computed: true}, "fragmentation": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"flags": schema.BoolAttribute{Computed: true}, "offset": schema.BoolAttribute{Computed: true}}}, "header_len": schema.BoolAttribute{Computed: true}, "option_map": schema.BoolAttribute{Computed: true}, "precedence": schema.BoolAttribute{Computed: true}, "protocol": schema.BoolAttribute{Computed: true}, "section": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"header_size": schema.Int64Attribute{Computed: true}, "payload_size": schema.Int64Attribute{Computed: true}}}, "source": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"prefix_min_mask": schema.StringAttribute{MarkdownDescription: "ipv4 source prefix minimum-mask - netmask or mask length", Computed: true}}}, "tos": schema.BoolAttribute{Computed: true}, "total_length": schema.BoolAttribute{Computed: true}, "ttl": schema.BoolAttribute{Computed: true}}}, "ipv6": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"destination": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"prefix_min_mask": schema.StringAttribute{Computed: true}}}, "dscp": schema.BoolAttribute{Computed: true}, "extension_map": schema.BoolAttribute{Computed: true}, "flow_label": schema.BoolAttribute{Computed: true}, "fragmentation": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"flags": schema.BoolAttribute{Computed: true}, "offset": schema.BoolAttribute{Computed: true}}}, "hop_limit": schema.BoolAttribute{Computed: true}, "length": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"header": schema.BoolAttribute{Computed: true}, "payload": schema.BoolAttribute{Computed: true}, "total": schema.BoolAttribute{Computed: true}}}, "next_header": schema.BoolAttribute{Computed: true}, "precedence": schema.BoolAttribute{Computed: true}, "section": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"header_size": schema.Int64Attribute{Computed: true}, "payload_size": schema.Int64Attribute{Computed: true}}}, "source": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"prefix_min_mask": schema.StringAttribute{MarkdownDescription: "ipv6 source prefix minimum-mask - netmask or mask length", Computed: true}}}, "traffic_class": schema.BoolAttribute{Computed: true}}}, "transport": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"dst_port": schema.BoolAttribute{Computed: true}, "icmp": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"ipv4_code": schema.BoolAttribute{Computed: true}, "ipv4_type": schema.BoolAttribute{Computed: true}, "ipv6_code": schema.BoolAttribute{Computed: true}, "ipv6_type": schema.BoolAttribute{Computed: true}}}, "src_port": schema.BoolAttribute{Computed: true}, "tcp": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"ack_number": schema.BoolAttribute{Computed: true}, "dst_port": schema.BoolAttribute{Computed: true}, "flags": schema.BoolAttribute{Computed: true}, "header_len": schema.BoolAttribute{Computed: true}, "seq_number": schema.BoolAttribute{Computed: true}, "src_port": schema.BoolAttribute{Computed: true}, "urgent_ptr": schema.BoolAttribute{Computed: true}, "window_size": schema.BoolAttribute{Computed: true}}}, "udp": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"dst_port": schema.BoolAttribute{Computed: true}, "msg_len": schema.BoolAttribute{Computed: true}, "src_port": schema.BoolAttribute{Computed: true}}}}}}}, "multi_collect": schema.BoolAttribute{MarkdownDescription: "Collect all attributes as it is discovered when enable. It will export the same record once when disable.", Computed: true}, "network_profiles": schema.ListAttribute{MarkdownDescription: "alias of metadata network profiles to attach this cache", Computed: true, ElementType: types.StringType}, "observation_domain_id": schema.Int64Attribute{Computed: true}, "sampling": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"mode": schema.StringAttribute{Computed: true}, "single_sampling_rate": schema.Int64Attribute{MarkdownDescription: "Packet interval window size. Valid values: 10-16000 (in packets)", Computed: true}}}, "size": schema.SingleNestedAttribute{MarkdownDescription: "size of the flows", Computed: true, Attributes: map[string]schema.Attribute{"flows": schema.Int64Attribute{MarkdownDescription: "size of flows in millions", Computed: true}}}, "timeout": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"idle": schema.Int64Attribute{MarkdownDescription: "idle timeout in seconds. max value 7days. default 30 min", Computed: true}}}}}}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned", Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadAllMetadataCacheDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadAllMetadataCacheDataSourceModel
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
func (d *LoadAllMetadataCacheDataSource) readListRemote(ctx context.Context, config *LoadAllMetadataCacheDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/apps/metadata/cache"
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
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_metadata_cache", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_metadata_cache", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["metadataCaches"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_metadata_cache", fmt.Sprintf("Could not decode list page: missing %q array", "metadataCaches"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_metadata_cache", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadAllMetadataCacheDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
