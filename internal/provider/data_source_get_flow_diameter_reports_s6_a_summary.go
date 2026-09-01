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
	_ datasource.DataSource              = (*GetFlowDiameterReportsS6ASummaryDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetFlowDiameterReportsS6ASummaryDataSource)(nil)
)

// GetFlowDiameterReportsS6ASummaryDataSource is the generated Terraform data source implementation.
type GetFlowDiameterReportsS6ASummaryDataSource struct {
	client *client.Client
}

// GetFlowDiameterReportsS6ASummaryDataSourceModel describes the data source state shape.
type GetFlowDiameterReportsS6ASummaryDataSourceModel struct {
	Items types.List `tfsdk:"items"`
}

// NewGetFlowDiameterReportsS6ASummaryDataSource returns a new instance of the generated data source.
func NewGetFlowDiameterReportsS6ASummaryDataSource() datasource.DataSource {
	return &GetFlowDiameterReportsS6ASummaryDataSource{}
}

// Metadata returns the data source type name.
func (d *GetFlowDiameterReportsS6ASummaryDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_flow_diameter_reports_s6_a_summary"
}

// Schema returns the data source schema.
func (d *GetFlowDiameterReportsS6ASummaryDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load Flow Diameter S6a Report Summary", Attributes: map[string]schema.Attribute{"items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"gsgroup": schema.StringAttribute{MarkdownDescription: "alias of gsgroup", Computed: true}, "s6_a_messages_stats": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"drop": schema.Int64Attribute{Computed: true}, "no_match": schema.Int64Attribute{Computed: true}, "no_rule": schema.Int64Attribute{Computed: true}, "no_session": schema.Int64Attribute{Computed: true}, "other": schema.Int64Attribute{Computed: true}, "s6_a_message": schema.StringAttribute{Computed: true}, "tool_pass": schema.Int64Attribute{Computed: true}}}}, "s6_a_resource_summary": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"num_sessions": schema.Int64Attribute{Computed: true}, "session_avail": schema.Int64Attribute{Computed: true}}}, "s6_a_sessions": schema.Int64Attribute{Computed: true}}}}}}
}

// Read fetches remote state into the data source model.
func (d *GetFlowDiameterReportsS6ASummaryDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetFlowDiameterReportsS6ASummaryDataSourceModel
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
func (d *GetFlowDiameterReportsS6ASummaryDataSource) readListRemote(ctx context.Context, config *GetFlowDiameterReportsS6ASummaryDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/gsGroups/flowOpsReport/flowDiameterS6a/summary"
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_flow_diameter_reports_s6_a_summary", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_flow_diameter_reports_s6_a_summary", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["flowDiameterS6aReportsSummary"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_flow_diameter_reports_s6_a_summary", fmt.Sprintf("Could not decode list page: missing %q array", "flowDiameterS6aReportsSummary"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_flow_diameter_reports_s6_a_summary", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetFlowDiameterReportsS6ASummaryDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
