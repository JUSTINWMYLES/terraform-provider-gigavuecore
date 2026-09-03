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
	_ datasource.DataSource              = (*GetFlexInlineVlanConfigDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetFlexInlineVlanConfigDataSource)(nil)
)

// GetFlexInlineVlanConfigDataSource is the generated Terraform data source implementation.
type GetFlexInlineVlanConfigDataSource struct {
	client *client.Client
}

// GetFlexInlineVlanConfigDataSourceModel describes the data source state shape.
type GetFlexInlineVlanConfigDataSourceModel struct {
	ClusterId          types.String `tfsdk:"cluster_id" json:"clusterId"`
	InlineNetworkAlias types.String `tfsdk:"inline_network_alias" json:"inlineNetworkAlias"`
	Items              types.List   `tfsdk:"items"`
}

// NewGetFlexInlineVlanConfigDataSource returns a new instance of the generated data source.
func NewGetFlexInlineVlanConfigDataSource() datasource.DataSource {
	return &GetFlexInlineVlanConfigDataSource{}
}

// Metadata returns the data source type name.
func (d *GetFlexInlineVlanConfigDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_flex_inline_vlan_config"
}

// Schema returns the data source schema.
func (d *GetFlexInlineVlanConfigDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Gets all vlan configs for the given cluster", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Gets vlan configs for the provided cluster", Required: true}, "inline_network_alias": schema.StringAttribute{MarkdownDescription: "if provided, returns vlan configs only for that inline network construct", Optional: true}, "items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"inline_network_alias": schema.StringAttribute{MarkdownDescription: "Alias of the inline network", Computed: true}, "map_alias": schema.StringAttribute{MarkdownDescription: "Alias of the map", Computed: true}, "vlan_configs": schema.ListNestedAttribute{MarkdownDescription: "Array holding vlan ids", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"flex_inline_tag_type": schema.StringAttribute{MarkdownDescription: "type of vlan id, auto (or) vlan", Computed: true}, "flex_inline_tag_vlan_id": schema.Int64Attribute{MarkdownDescription: "Tool side vlan id", Computed: true}, "flex_inline_vlan_id": schema.Int64Attribute{MarkdownDescription: "Network side vlan id", Computed: true}, "network_alias": schema.StringAttribute{MarkdownDescription: "Alias of the network source", Computed: true}, "tag_protocol_id": schema.StringAttribute{MarkdownDescription: "When tool VLAN tag is added , this protocol Id will be added which egress out the traffic", Computed: true}}}}}}}}}
}

// Read fetches remote state into the data source model.
func (d *GetFlexInlineVlanConfigDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetFlexInlineVlanConfigDataSourceModel
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
func (d *GetFlexInlineVlanConfigDataSource) readListRemote(ctx context.Context, config *GetFlexInlineVlanConfigDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/flexInline/mapVlanConfigs"
	params := url.Values{}
	params.Set("clusterId", config.ClusterId.ValueString())
	if !config.InlineNetworkAlias.IsNull() {
		params.Set("inlineNetworkAlias", config.InlineNetworkAlias.ValueString())
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_flex_inline_vlan_config", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_flex_inline_vlan_config", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["mapVlanConfigs"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_flex_inline_vlan_config", fmt.Sprintf("Could not decode list page: missing %q array", "mapVlanConfigs"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_flex_inline_vlan_config", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetFlexInlineVlanConfigDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
