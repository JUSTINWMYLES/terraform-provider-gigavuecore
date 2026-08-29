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
	_ datasource.DataSource              = (*LoadSpineLinkAllDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadSpineLinkAllDataSource)(nil)
)

// LoadSpineLinkAllDataSource is the generated Terraform data source implementation.
type LoadSpineLinkAllDataSource struct {
	client *client.Client
}

// LoadSpineLinkAllDataSourceModel describes the data source state shape.
type LoadSpineLinkAllDataSourceModel struct {
	Items types.List `tfsdk:"items"`
}

// NewLoadSpineLinkAllDataSource returns a new instance of the generated data source.
func NewLoadSpineLinkAllDataSource() datasource.DataSource {
	return &LoadSpineLinkAllDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadSpineLinkAllDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_spine_link_all"
}

// Schema returns the data source schema.
func (d *LoadSpineLinkAllDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load spine-link configuration", Attributes: map[string]schema.Attribute{"items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{Computed: true}, "comment": schema.StringAttribute{Computed: true}, "leaf_box_id": schema.Int64Attribute{MarkdownDescription: "box-id of the leaf node this spineLink is configured on", Computed: true}, "links": schema.ListNestedAttribute{MarkdownDescription: "List of links from this Leaf node to the Spine Nodes", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"leaf_gigastream": schema.StringAttribute{MarkdownDescription: "leaf node stack gigastream alias", Computed: true}, "spine_box_id": schema.Int64Attribute{MarkdownDescription: "box-id of the spine nodes this spineLink is connected to", Computed: true}, "spine_gigastream": schema.StringAttribute{MarkdownDescription: "spine node gigastream alias", Computed: true}, "stack_link": schema.StringAttribute{MarkdownDescription: "Alias of the Stack Link this Spine Link is running over. Will be empty if corresponding StackLink is not yet created", Computed: true}}}}}}}}}
}

// Read fetches remote state into the data source model.
func (d *LoadSpineLinkAllDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadSpineLinkAllDataSourceModel
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
func (d *LoadSpineLinkAllDataSource) readListRemote(ctx context.Context, config *LoadSpineLinkAllDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/spineLinks"
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
		resp.Diagnostics.AddError("Error reading gigavuecore_load_spine_link_all", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_spine_link_all", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["spineLinks"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_spine_link_all", fmt.Sprintf("Could not decode list page: missing %q array", "spineLinks"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_spine_link_all", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadSpineLinkAllDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
