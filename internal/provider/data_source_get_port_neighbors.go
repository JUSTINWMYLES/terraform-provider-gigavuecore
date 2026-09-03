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
	_ datasource.DataSource              = (*GetPortNeighborsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetPortNeighborsDataSource)(nil)
)

// GetPortNeighborsDataSource is the generated Terraform data source implementation.
type GetPortNeighborsDataSource struct {
	client *client.Client
}

// GetPortNeighborsDataSourceModel describes the data source state shape.
type GetPortNeighborsDataSourceModel struct {
	ClusterId types.String  `tfsdk:"cluster_id" json:"clusterId"`
	Items     types.Dynamic `tfsdk:"items"`
	Page      types.String  `tfsdk:"page"`
	PortId    types.String  `tfsdk:"port_id" json:"portId"`
	SlotId    types.String  `tfsdk:"slot_id" json:"slotId"`
	Sort      types.String  `tfsdk:"sort"`
}

// NewGetPortNeighborsDataSource returns a new instance of the generated data source.
func NewGetPortNeighborsDataSource() datasource.DataSource {
	return &GetPortNeighborsDataSource{}
}

// Metadata returns the data source type name.
func (d *GetPortNeighborsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_port_neighbors"
}

// Schema returns the data source schema.
func (d *GetPortNeighborsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get Port Neighbors", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "items": schema.DynamicAttribute{Computed: true}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned", Optional: true}, "port_id": schema.StringAttribute{MarkdownDescription: "target device Port Id. Ignored if 'slotId' parameter is provided.", Optional: true}, "slot_id": schema.StringAttribute{MarkdownDescription: "target device Slot Id. Takes precedence over 'portId' parameter", Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetPortNeighborsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetPortNeighborsDataSourceModel
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
func (d *GetPortNeighborsDataSource) readListRemote(ctx context.Context, config *GetPortNeighborsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/inventory/ports/neighbors"
	params := url.Values{}
	params.Set("clusterId", config.ClusterId.ValueString())
	if !config.SlotId.IsNull() {
		params.Set("slotId", config.SlotId.ValueString())
	}
	if !config.PortId.IsNull() {
		params.Set("portId", config.PortId.ValueString())
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_port_neighbors", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_port_neighbors", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["portNeighbors"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_port_neighbors", fmt.Sprintf("Could not decode list page: missing %q array", "portNeighbors"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_port_neighbors", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetPortNeighborsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
