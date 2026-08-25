package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)
import (
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*GetTrafficFlowsOverviewDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetTrafficFlowsOverviewDataSource)(nil)
)

// GetTrafficFlowsOverviewDataSource is the generated Terraform data source implementation.
type GetTrafficFlowsOverviewDataSource struct {
	client *client.Client
}

// GetTrafficFlowsOverviewDataSourceModel describes the data source state shape.
type GetTrafficFlowsOverviewDataSourceModel struct {
	Failure    types.Int64 `tfsdk:"failure"`
	Green      types.Int64 `tfsdk:"green"`
	Others     types.Int64 `tfsdk:"others"`
	Red        types.Int64 `tfsdk:"red"`
	Success    types.Int64 `tfsdk:"success"`
	Total      types.Int64 `tfsdk:"total"`
	Undeployed types.Int64 `tfsdk:"undeployed"`
	Yellow     types.Int64 `tfsdk:"yellow"`
}

// NewGetTrafficFlowsOverviewDataSource returns a new instance of the generated data source.
func NewGetTrafficFlowsOverviewDataSource() datasource.DataSource {
	return &GetTrafficFlowsOverviewDataSource{}
}

// Metadata returns the data source type name.
func (d *GetTrafficFlowsOverviewDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_traffic_flows_overview"
}

// Schema returns the data source schema.
func (d *GetTrafficFlowsOverviewDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get overview of all Traffic Flows", Attributes: map[string]schema.Attribute{"failure": schema.Int64Attribute{Computed: true}, "green": schema.Int64Attribute{Computed: true}, "others": schema.Int64Attribute{Computed: true}, "red": schema.Int64Attribute{Computed: true}, "success": schema.Int64Attribute{Computed: true}, "total": schema.Int64Attribute{Computed: true}, "undeployed": schema.Int64Attribute{Computed: true}, "yellow": schema.Int64Attribute{Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetTrafficFlowsOverviewDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetTrafficFlowsOverviewDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	d.readRemote(ctx, &config, resp)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}

// readRemote performs the read HTTP exchange and decodes the response into config. Extracted from Read so the request/response logic is unit-testable without a tfsdk.Config.
func (d *GetTrafficFlowsOverviewDataSource) readRemote(ctx context.Context, config *GetTrafficFlowsOverviewDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/traffic-flows/all/overview"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_traffic_flows_overview", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_traffic_flows_overview", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_traffic_flows_overview", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_traffic_flows_overview", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_traffic_flows_overview", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_traffic_flows_overview", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_traffic_flows_overview", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_traffic_flows_overview", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_traffic_flows_overview", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_traffic_flows_overview", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_traffic_flows_overview", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_traffic_flows_overview", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_traffic_flows_overview", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["trafficFlowsOverview"]; ok {
		if m, ok := v.(map[string]any); ok {
			data = m
		}
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_traffic_flows_overview", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetTrafficFlowsOverviewDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
