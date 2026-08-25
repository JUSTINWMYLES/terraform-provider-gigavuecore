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
	_ datasource.DataSource              = (*GetGigaInsightNodesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetGigaInsightNodesDataSource)(nil)
)

// GetGigaInsightNodesDataSource is the generated Terraform data source implementation.
type GetGigaInsightNodesDataSource struct {
	client *client.Client
}

// GetGigaInsightNodesDataSourceModel describes the data source state shape.
type GetGigaInsightNodesDataSourceModel struct {
	Alias types.String `tfsdk:"alias"`
	Items types.List   `tfsdk:"items"`
}

// NewGetGigaInsightNodesDataSource returns a new instance of the generated data source.
func NewGetGigaInsightNodesDataSource() datasource.DataSource {
	return &GetGigaInsightNodesDataSource{}
}

// Metadata returns the data source type name.
func (d *GetGigaInsightNodesDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_giga_insight_nodes"
}

// Schema returns the data source schema.
func (d *GetGigaInsightNodesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get all GigaInsight Nodes", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Filter by node alias", Optional: true}, "items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Alias name for the GI node", Computed: true}, "created_date": schema.StringAttribute{MarkdownDescription: "Date when the node was created", Computed: true}, "ipv4_address": schema.StringAttribute{MarkdownDescription: "IPv4 address of the GI node", Computed: true}, "ipv6_address": schema.StringAttribute{MarkdownDescription: "IPv6 address of the GI node", Computed: true}, "last_config_updated_at": schema.StringAttribute{MarkdownDescription: "Date when the node was last online", Computed: true}, "last_status_updated_at": schema.StringAttribute{MarkdownDescription: "Date when the node was last updated", Computed: true}, "node_id": schema.StringAttribute{MarkdownDescription: "Unique identifier for the GI node", Computed: true}, "node_version": schema.StringAttribute{MarkdownDescription: "Version of the GigaInsight node software", Computed: true}, "prompt_bundle": schema.SingleNestedAttribute{MarkdownDescription: "Prompt bundle information for the GI node", Computed: true, Attributes: map[string]schema.Attribute{"last_updated_at": schema.StringAttribute{MarkdownDescription: "Date when the prompt bundle was last updated", Computed: true}, "prompt_bundle_version": schema.StringAttribute{MarkdownDescription: "Version of the prompt bundle", Computed: true}}}, "status": schema.StringAttribute{MarkdownDescription: "Current status of the GI node", Computed: true}}}}}}
}

// Read fetches remote state into the data source model.
func (d *GetGigaInsightNodesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetGigaInsightNodesDataSourceModel
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
func (d *GetGigaInsightNodesDataSource) readListRemote(ctx context.Context, config *GetGigaInsightNodesDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/gi/nodes"
	params := url.Values{}
	if !config.Alias.IsNull() {
		params.Set("alias", config.Alias.ValueString())
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_giga_insight_nodes", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageItems := []any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageItems); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_giga_insight_nodes", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_giga_insight_nodes", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetGigaInsightNodesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
