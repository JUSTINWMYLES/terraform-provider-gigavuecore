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
	_ datasource.DataSource              = (*GetConnectedLinksTrafficFlowsTroubleshootFlowsTopologyDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetConnectedLinksTrafficFlowsTroubleshootFlowsTopologyDataSource)(nil)
)

// GetConnectedLinksTrafficFlowsTroubleshootFlowsTopologyDataSource is the generated Terraform data source implementation.
type GetConnectedLinksTrafficFlowsTroubleshootFlowsTopologyDataSource struct {
	client *client.Client
}

// GetConnectedLinksTrafficFlowsTroubleshootFlowsTopologyDataSourceModel describes the data source state shape.
type GetConnectedLinksTrafficFlowsTroubleshootFlowsTopologyDataSourceModel struct {
	ClusterName types.String  `tfsdk:"cluster_name" json:"clusterName"`
	Items       types.Dynamic `tfsdk:"items"`
	StartPoint  types.String  `tfsdk:"start_point"`
}

// NewGetConnectedLinksTrafficFlowsTroubleshootFlowsTopologyDataSource returns a new instance of the generated data source.
func NewGetConnectedLinksTrafficFlowsTroubleshootFlowsTopologyDataSource() datasource.DataSource {
	return &GetConnectedLinksTrafficFlowsTroubleshootFlowsTopologyDataSource{}
}

// Metadata returns the data source type name.
func (d *GetConnectedLinksTrafficFlowsTroubleshootFlowsTopologyDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_connected_links_traffic_flows_troubleshoot_flows_topology"
}

// Schema returns the data source schema.
func (d *GetConnectedLinksTrafficFlowsTroubleshootFlowsTopologyDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Retrieve traversed nodes with levels for a cluster and port", Attributes: map[string]schema.Attribute{"cluster_name": schema.StringAttribute{MarkdownDescription: "Target Cluster Name", Optional: true}, "items": schema.DynamicAttribute{Computed: true}, "start_point": schema.StringAttribute{MarkdownDescription: "Start port", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetConnectedLinksTrafficFlowsTroubleshootFlowsTopologyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetConnectedLinksTrafficFlowsTroubleshootFlowsTopologyDataSourceModel
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
func (d *GetConnectedLinksTrafficFlowsTroubleshootFlowsTopologyDataSource) readListRemote(ctx context.Context, config *GetConnectedLinksTrafficFlowsTroubleshootFlowsTopologyDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/traffic-flows/troubleshoot/flows/topology"
	params := url.Values{}
	if !config.ClusterName.IsNull() {
		params.Set("clusterName", config.ClusterName.ValueString())
	}
	if !config.StartPoint.IsNull() {
		params.Set("start_point", config.StartPoint.ValueString())
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_connected_links_traffic_flows_troubleshoot_flows_topology", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageItems := []any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageItems); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_connected_links_traffic_flows_troubleshoot_flows_topology", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_connected_links_traffic_flows_troubleshoot_flows_topology", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetConnectedLinksTrafficFlowsTroubleshootFlowsTopologyDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
