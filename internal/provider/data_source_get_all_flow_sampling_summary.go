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
	_ datasource.DataSource              = (*GetAllFlowSamplingSummaryDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetAllFlowSamplingSummaryDataSource)(nil)
)

// GetAllFlowSamplingSummaryDataSource is the generated Terraform data source implementation.
type GetAllFlowSamplingSummaryDataSource struct {
	client *client.Client
}

// GetAllFlowSamplingSummaryDataSourceModel describes the data source state shape.
type GetAllFlowSamplingSummaryDataSourceModel struct {
	Items types.List `tfsdk:"items"`
}

// NewGetAllFlowSamplingSummaryDataSource returns a new instance of the generated data source.
func NewGetAllFlowSamplingSummaryDataSource() datasource.DataSource {
	return &GetAllFlowSamplingSummaryDataSource{}
}

// Metadata returns the data source type name.
func (d *GetAllFlowSamplingSummaryDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_all_flow_sampling_summary"
}

// Schema returns the data source schema.
func (d *GetAllFlowSamplingSummaryDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load all Flow Sampling Report Summary", Attributes: map[string]schema.Attribute{"items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"gsgroup": schema.StringAttribute{MarkdownDescription: "alias of gsgroup", Computed: true}, "num_devices": schema.Int64Attribute{Computed: true}, "num_devices_in_sample": schema.Int64Attribute{Computed: true}}}}}}
}

// Read fetches remote state into the data source model.
func (d *GetAllFlowSamplingSummaryDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetAllFlowSamplingSummaryDataSourceModel
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
func (d *GetAllFlowSamplingSummaryDataSource) readListRemote(ctx context.Context, config *GetAllFlowSamplingSummaryDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/gsGroups/flowOpsReport/flowSampling/summary"
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_flow_sampling_summary", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_flow_sampling_summary", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["flowSamplingReportsSummary"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_flow_sampling_summary", fmt.Sprintf("Could not decode list page: missing %q array", "flowSamplingReportsSummary"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_flow_sampling_summary", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetAllFlowSamplingSummaryDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
