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
	_ datasource.DataSource              = (*LoadAllSessionAwareApfProfilesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadAllSessionAwareApfProfilesDataSource)(nil)
)

// LoadAllSessionAwareApfProfilesDataSource is the generated Terraform data source implementation.
type LoadAllSessionAwareApfProfilesDataSource struct {
	client *client.Client
}

// LoadAllSessionAwareApfProfilesDataSourceModel describes the data source state shape.
type LoadAllSessionAwareApfProfilesDataSourceModel struct {
	ClusterId types.String `tfsdk:"cluster_id" json:"clusterId"`
	Items     types.List   `tfsdk:"items"`
	Page      types.String `tfsdk:"page"`
	Sort      types.String `tfsdk:"sort"`
}

// NewLoadAllSessionAwareApfProfilesDataSource returns a new instance of the generated data source.
func NewLoadAllSessionAwareApfProfilesDataSource() datasource.DataSource {
	return &LoadAllSessionAwareApfProfilesDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadAllSessionAwareApfProfilesDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_all_session_aware_apf_profiles"
}

// Schema returns the data source schema.
func (d *LoadAllSessionAwareApfProfilesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load all defined SA-APF Profiles", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{Computed: true}, "bidi": schema.BoolAttribute{MarkdownDescription: "include reverse traffic for the selected sessionFields. Not applicable for 'mplsLabel', 'gtpuTeid', 'vlanId'", Computed: true}, "buffering": schema.SingleNestedAttribute{MarkdownDescription: "Session-Aware APF Buffereing settings", Computed: true, Attributes: map[string]schema.Attribute{"buffer_count_before_match": schema.Int64Attribute{MarkdownDescription: "Maximum number of packets BSAPF will buffer per session before APF match", Computed: true}, "enabled": schema.BoolAttribute{Computed: true}, "protocol": schema.StringAttribute{MarkdownDescription: "When buffering is enabled, changing protocol requires reboot before new value takes effect.", Computed: true}}}, "cluster_id": schema.StringAttribute{MarkdownDescription: "id of the defining cluster", Computed: true}, "packet_count": schema.Int64Attribute{MarkdownDescription: "For each session, forward this number of packets. Valid range is 2-100; 0 disables packetCount.", Computed: true}, "session_fields": schema.SetNestedAttribute{MarkdownDescription: "A sessionField cannot contain overlapping session attribute, position pairs (e.g. ipv4-5tuple pos 1 and ipv4-src pos 1 is not allowed). Max number of session fields is 20", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"pos": schema.Int64Attribute{MarkdownDescription: "Value of 1 also an alias for 'outer'. Value of 2 also an alias for 'inner'. Not applicable for 'gtpuTeid'. For 'fiveTuple' 'outer' is not supported", Computed: true}, "type": schema.StringAttribute{Computed: true}}}}, "timeout": schema.Int64Attribute{MarkdownDescription: "in seconds", Computed: true}}}}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned", Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadAllSessionAwareApfProfilesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadAllSessionAwareApfProfilesDataSourceModel
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
func (d *LoadAllSessionAwareApfProfilesDataSource) readListRemote(ctx context.Context, config *LoadAllSessionAwareApfProfilesDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/apps/saApfProfiles"
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
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_session_aware_apf_profiles", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_session_aware_apf_profiles", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["saApfProfiles"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_session_aware_apf_profiles", fmt.Sprintf("Could not decode list page: missing %q array", "saApfProfiles"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_session_aware_apf_profiles", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadAllSessionAwareApfProfilesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
