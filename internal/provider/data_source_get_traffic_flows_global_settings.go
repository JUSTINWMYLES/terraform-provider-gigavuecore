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
	_ datasource.DataSource              = (*GetTrafficFlowsGlobalSettingsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetTrafficFlowsGlobalSettingsDataSource)(nil)
)

// GetTrafficFlowsGlobalSettingsDataSource is the generated Terraform data source implementation.
type GetTrafficFlowsGlobalSettingsDataSource struct {
	client *client.Client
}

// GetTrafficFlowsGlobalSettingsDataSourceModel describes the data source state shape.
type GetTrafficFlowsGlobalSettingsDataSourceModel struct {
	AutoMigrate    types.Bool   `tfsdk:"auto_migrate" json:"autoMigrate"`
	FabricResource types.Object `tfsdk:"fabric_resource" json:"fabricResource"`
	L2Circuit      types.Object `tfsdk:"l2_circuit" json:"l2Circuit"`
}

// NewGetTrafficFlowsGlobalSettingsDataSource returns a new instance of the generated data source.
func NewGetTrafficFlowsGlobalSettingsDataSource() datasource.DataSource {
	return &GetTrafficFlowsGlobalSettingsDataSource{}
}

// Metadata returns the data source type name.
func (d *GetTrafficFlowsGlobalSettingsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_traffic_flows_global_settings"
}

// Schema returns the data source schema.
func (d *GetTrafficFlowsGlobalSettingsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get global settings for Traffic Flows", Attributes: map[string]schema.Attribute{"auto_migrate": schema.BoolAttribute{Computed: true}, "fabric_resource": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"mode": schema.StringAttribute{MarkdownDescription: "['SHARE' or 'NOT_SHARE']: Resource sharing mode. SHARE: same resource can be shared by different fabric maps; NOT_SHARE: not shared. Default: NOT_SHARE.", Computed: true}, "scope": schema.StringAttribute{MarkdownDescription: "['GLOBAL']: Scope of resource pool. GLOBAL: only one global resource pool.", Computed: true}, "type": schema.StringAttribute{MarkdownDescription: "['L2CIRCUIT']: Resource type.", Computed: true}}}, "l2_circuit": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"vlan_ids": schema.StringAttribute{MarkdownDescription: "VLAN id ranges used for L2CIRCUIT resource type.", Computed: true}}}}}
}

// Read fetches remote state into the data source model.
func (d *GetTrafficFlowsGlobalSettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetTrafficFlowsGlobalSettingsDataSourceModel
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
func (d *GetTrafficFlowsGlobalSettingsDataSource) readRemote(ctx context.Context, config *GetTrafficFlowsGlobalSettingsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/traffic-flows/global/settings"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_traffic_flows_global_settings", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_traffic_flows_global_settings", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_traffic_flows_global_settings", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_traffic_flows_global_settings", "Invalid request")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_traffic_flows_global_settings", "Not Authenticated")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_traffic_flows_global_settings", "Access Denied")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_traffic_flows_global_settings", "Entity Not Found")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_traffic_flows_global_settings", "Internal Server Error")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_traffic_flows_global_settings", "Service Unavailable")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_traffic_flows_global_settings", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_traffic_flows_global_settings", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_traffic_flows_global_settings", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_traffic_flows_global_settings", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetTrafficFlowsGlobalSettingsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
