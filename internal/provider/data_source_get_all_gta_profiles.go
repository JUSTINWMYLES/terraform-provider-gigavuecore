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
	_ datasource.DataSource              = (*GetAllGtaProfilesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetAllGtaProfilesDataSource)(nil)
)

// GetAllGtaProfilesDataSource is the generated Terraform data source implementation.
type GetAllGtaProfilesDataSource struct {
	client *client.Client
}

// GetAllGtaProfilesDataSourceModel describes the data source state shape.
type GetAllGtaProfilesDataSourceModel struct {
	Items types.List `tfsdk:"items"`
}

// NewGetAllGtaProfilesDataSource returns a new instance of the generated data source.
func NewGetAllGtaProfilesDataSource() datasource.DataSource {
	return &GetAllGtaProfilesDataSource{}
}

// Metadata returns the data source type name.
func (d *GetAllGtaProfilesDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_all_gta_profiles"
}

// Schema returns the data source schema.
func (d *GetAllGtaProfilesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "get all gta profiles", Attributes: map[string]schema.Attribute{"items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "alias of the GTA profile", Computed: true}, "control_node": schema.StringAttribute{Computed: true}, "core_network_nodes": schema.SingleNestedAttribute{MarkdownDescription: "core network nodes specified in one of 3 ways: list of addresses, range of addresses, or an address subnet", Computed: true, Attributes: map[string]schema.Attribute{"address_list": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"value": schema.ListAttribute{MarkdownDescription: "ipv4 or ipv6 addresses", Computed: true, ElementType: types.StringType}}}, "address_range": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"ip_ranges": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"max_value": schema.StringAttribute{MarkdownDescription: "ipv4 or ipv6 maximum value", Computed: true}, "value": schema.StringAttribute{MarkdownDescription: "ipv4 or ipv6 minimum value", Computed: true}}}}}}, "address_subnet": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"values": schema.ListAttribute{MarkdownDescription: "array of ipv4 or ipv6 subnet address space. Mutually exclusive with value", Computed: true, ElementType: types.StringType}}}}}, "dst_port": schema.Int64Attribute{Computed: true}, "src_port": schema.Int64Attribute{Computed: true}, "user_node": schema.StringAttribute{Computed: true}}}}}}
}

// Read fetches remote state into the data source model.
func (d *GetAllGtaProfilesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetAllGtaProfilesDataSourceModel
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
func (d *GetAllGtaProfilesDataSource) readListRemote(ctx context.Context, config *GetAllGtaProfilesDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/gtaProfiles"
	params := url.Values{}
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_gta_profiles", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_gta_profiles", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["gtaProfiles"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_gta_profiles", fmt.Sprintf("Could not decode list page: missing %q array", "gtaProfiles"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_gta_profiles", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetAllGtaProfilesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
