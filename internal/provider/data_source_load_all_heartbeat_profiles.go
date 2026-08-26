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
	_ datasource.DataSource              = (*LoadAllHeartbeatProfilesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadAllHeartbeatProfilesDataSource)(nil)
)

// LoadAllHeartbeatProfilesDataSource is the generated Terraform data source implementation.
type LoadAllHeartbeatProfilesDataSource struct {
	client *client.Client
}

// LoadAllHeartbeatProfilesDataSourceModel describes the data source state shape.
type LoadAllHeartbeatProfilesDataSourceModel struct {
	ClusterId types.String `tfsdk:"cluster_id" json:"clusterId"`
	Items     types.List   `tfsdk:"items"`
	Page      types.String `tfsdk:"page"`
	Sort      types.String `tfsdk:"sort"`
}

// NewLoadAllHeartbeatProfilesDataSource returns a new instance of the generated data source.
func NewLoadAllHeartbeatProfilesDataSource() datasource.DataSource {
	return &LoadAllHeartbeatProfilesDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadAllHeartbeatProfilesDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_all_heartbeat_profiles"
}

// Schema returns the data source schema.
func (d *LoadAllHeartbeatProfilesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load all Heartbeat Profiles", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Heartbeat Profile alias. Unique within a cluster", Computed: true}, "custom_packet": schema.StringAttribute{MarkdownDescription: "Base64-encoded custom pcap packet. If omitted, a standard ICMP ARP packet will be used as a heartbeat packet. Custom heartbeat packets can be ICMP/IPv4 or ICMPv6/IPv6", Computed: true}, "custom_packet_alias": schema.StringAttribute{MarkdownDescription: "(deprecated) Alias of referenced 'custom heartbeat packet' entry. Used for 'inline heartbeat profile' creation", Computed: true}, "custom_packet_file_name": schema.StringAttribute{MarkdownDescription: "File name of referenced 'custom heartbeat packet' entry", Computed: true}, "direction": schema.StringAttribute{Computed: true}, "packet_format": schema.StringAttribute{Computed: true}, "period": schema.Int64Attribute{MarkdownDescription: "number of milliseconds between sending subsequent heartbeat packets", Computed: true}, "recovery_time": schema.Int64Attribute{MarkdownDescription: " the minimum number of seconds with successfully received packet to declare that the inline tool is up", Computed: true}, "retries": schema.Int64Attribute{MarkdownDescription: "number of consecutive timed-out heartbeat packets at which the system will trigger a failover condition", Computed: true}, "timeout": schema.Int64Attribute{MarkdownDescription: "number of milliseconds allowed for a heartbeat packet between sending and receiving", Computed: true}}}}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned", Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadAllHeartbeatProfilesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadAllHeartbeatProfilesDataSourceModel
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
func (d *LoadAllHeartbeatProfilesDataSource) readListRemote(ctx context.Context, config *LoadAllHeartbeatProfilesDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/inline/hbProfiles"
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
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_heartbeat_profiles", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_heartbeat_profiles", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["heartbeatProfiles"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_heartbeat_profiles", fmt.Sprintf("Could not decode list page: missing %q array", "heartbeatProfiles"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_heartbeat_profiles", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadAllHeartbeatProfilesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
