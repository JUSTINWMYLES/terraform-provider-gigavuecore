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
	_ datasource.DataSource              = (*LoadPermittedAddressesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadPermittedAddressesDataSource)(nil)
)

// LoadPermittedAddressesDataSource is the generated Terraform data source implementation.
type LoadPermittedAddressesDataSource struct {
	client *client.Client
}

// LoadPermittedAddressesDataSourceModel describes the data source state shape.
type LoadPermittedAddressesDataSourceModel struct {
	Address     types.String `tfsdk:"address"`
	AddressType types.String `tfsdk:"address_type" json:"addressType"`
	Alias       types.String `tfsdk:"alias"`
	Items       types.List   `tfsdk:"items"`
	Page        types.String `tfsdk:"page"`
	Sort        types.String `tfsdk:"sort"`
}

// NewLoadPermittedAddressesDataSource returns a new instance of the generated data source.
func NewLoadPermittedAddressesDataSource() datasource.DataSource {
	return &LoadPermittedAddressesDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadPermittedAddressesDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_permitted_addresses"
}

// Schema returns the data source schema.
func (d *LoadPermittedAddressesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load Permitted Addresses", Attributes: map[string]schema.Attribute{"address": schema.StringAttribute{MarkdownDescription: "address", Optional: true}, "address_type": schema.StringAttribute{MarkdownDescription: "addressType", Optional: true}, "alias": schema.StringAttribute{MarkdownDescription: "alias", Optional: true}, "items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"address": schema.StringAttribute{MarkdownDescription: "Permitted address", Computed: true}, "address_type": schema.StringAttribute{MarkdownDescription: "Address type", Computed: true}, "alias": schema.StringAttribute{MarkdownDescription: "Alias for the address", Computed: true}, "subscribed_events": schema.ListAttribute{MarkdownDescription: "List of events subscribed by the address", Computed: true, ElementType: types.StringType}}}}, "page": schema.StringAttribute{MarkdownDescription: "page", Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadPermittedAddressesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadPermittedAddressesDataSourceModel
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
func (d *LoadPermittedAddressesDataSource) readListRemote(ctx context.Context, config *LoadPermittedAddressesDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/notification/email/allowlistAddress"
	params := url.Values{}
	if !config.Address.IsNull() {
		params.Set("address", config.Address.ValueString())
	}
	if !config.Alias.IsNull() {
		params.Set("alias", config.Alias.ValueString())
	}
	if !config.AddressType.IsNull() {
		params.Set("addressType", config.AddressType.ValueString())
	}
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
		resp.Diagnostics.AddError("Error reading gigavuecore_load_permitted_addresses", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_permitted_addresses", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["allowListAddresses"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_permitted_addresses", fmt.Sprintf("Could not decode list page: missing %q array", "allowListAddresses"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_permitted_addresses", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadPermittedAddressesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
