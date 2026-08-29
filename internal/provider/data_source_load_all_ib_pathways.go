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
	_ datasource.DataSource              = (*LoadAllIbPathwaysDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadAllIbPathwaysDataSource)(nil)
)

// LoadAllIbPathwaysDataSource is the generated Terraform data source implementation.
type LoadAllIbPathwaysDataSource struct {
	client *client.Client
}

// LoadAllIbPathwaysDataSourceModel describes the data source state shape.
type LoadAllIbPathwaysDataSourceModel struct {
	Items types.List   `tfsdk:"items"`
	Page  types.String `tfsdk:"page"`
	Sort  types.String `tfsdk:"sort"`
}

// NewLoadAllIbPathwaysDataSource returns a new instance of the generated data source.
func NewLoadAllIbPathwaysDataSource() datasource.DataSource {
	return &LoadAllIbPathwaysDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadAllIbPathwaysDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_all_ib_pathways"
}

// Schema returns the data source schema.
func (d *LoadAllIbPathwaysDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load all Inter-broker Pathways", Attributes: map[string]schema.Attribute{"items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Inline resilient inter-broker pathway alias", Computed: true}, "comment": schema.StringAttribute{Computed: true}, "min_ports_up": schema.Int64Attribute{MarkdownDescription: "minimum number of ports in the 'up' state needed to declare the ib-pathway to be in the 'up' state.", Computed: true}, "operational_state": schema.StringAttribute{Computed: true}, "ports": schema.ListAttribute{MarkdownDescription: "list of local network ports with same speed", Computed: true, ElementType: types.StringType}, "traffic_path": schema.StringAttribute{Computed: true}}}}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned", Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadAllIbPathwaysDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadAllIbPathwaysDataSourceModel
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
func (d *LoadAllIbPathwaysDataSource) readListRemote(ctx context.Context, config *LoadAllIbPathwaysDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/inline/ibPathways"
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
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_ib_pathways", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_ib_pathways", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["ibPathways"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_ib_pathways", fmt.Sprintf("Could not decode list page: missing %q array", "ibPathways"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_ib_pathways", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadAllIbPathwaysDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
