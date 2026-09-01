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
	_ datasource.DataSource              = (*GetAllFlexInlineSolutionsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetAllFlexInlineSolutionsDataSource)(nil)
)

// GetAllFlexInlineSolutionsDataSource is the generated Terraform data source implementation.
type GetAllFlexInlineSolutionsDataSource struct {
	client *client.Client
}

// GetAllFlexInlineSolutionsDataSourceModel describes the data source state shape.
type GetAllFlexInlineSolutionsDataSourceModel struct {
	Items types.List `tfsdk:"items"`
}

// NewGetAllFlexInlineSolutionsDataSource returns a new instance of the generated data source.
func NewGetAllFlexInlineSolutionsDataSource() datasource.DataSource {
	return &GetAllFlexInlineSolutionsDataSource{}
}

// Metadata returns the data source type name.
func (d *GetAllFlexInlineSolutionsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_all_flex_inline_solutions"
}

// Schema returns the data source schema.
func (d *GetAllFlexInlineSolutionsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get All flexInline solutions present on FM", Attributes: map[string]schema.Attribute{"items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Alias of the solution", Computed: true}, "cluster_configs": schema.ListNestedAttribute{MarkdownDescription: "FlexInline solution source details", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{Computed: true}, "export_criteria": schema.SingleNestedAttribute{MarkdownDescription: "Based on this criteria and source on which this criteria is defined, traffic is guided to the other node", Computed: true, Attributes: map[string]schema.Attribute{"lsb": schema.Int64Attribute{Computed: true}}}, "export_type": schema.StringAttribute{Computed: true}, "ib_pathway": schema.StringAttribute{MarkdownDescription: "Inter-broker pathway to guide traffic from one node to other", Computed: true}, "source": schema.SingleNestedAttribute{MarkdownDescription: "Object holding flexInline source details", Computed: true, Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Alias of the source", Computed: true}, "type": schema.StringAttribute{MarkdownDescription: "Source type", Computed: true}}}}}}, "config_status": schema.StringAttribute{Computed: true}, "config_status_reasons": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "health_state": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "health_state_reasons": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"message": schema.StringAttribute{MarkdownDescription: "Read-only. Describes the reason for component's health state", Computed: true}, "severity": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "traffic_health_state_computation_type": schema.StringAttribute{MarkdownDescription: "Traffic Health State Computation Type", Computed: true}}}}, "resilient_config": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"side_a": schema.StringAttribute{MarkdownDescription: "Side A of the resilient map", Computed: true}, "side_b": schema.StringAttribute{MarkdownDescription: "Side B of the resilient map", Computed: true}}}, "target_traffic_path": schema.StringAttribute{MarkdownDescription: "Inline network target traffic path", Computed: true}}}}}}
}

// Read fetches remote state into the data source model.
func (d *GetAllFlexInlineSolutionsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetAllFlexInlineSolutionsDataSourceModel
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
func (d *GetAllFlexInlineSolutionsDataSource) readListRemote(ctx context.Context, config *GetAllFlexInlineSolutionsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/flexInline"
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_flex_inline_solutions", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_flex_inline_solutions", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["flexInlineSolutions"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_flex_inline_solutions", fmt.Sprintf("Could not decode list page: missing %q array", "flexInlineSolutions"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_flex_inline_solutions", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetAllFlexInlineSolutionsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
