package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)
import (
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*LoadBatteryOptimizationDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadBatteryOptimizationDataSource)(nil)
)

// LoadBatteryOptimizationDataSource is the generated Terraform data source implementation.
type LoadBatteryOptimizationDataSource struct {
	client *client.Client
}

// LoadBatteryOptimizationDataSourceModel describes the data source state shape.
type LoadBatteryOptimizationDataSourceModel struct {
	ClusterId      types.String `tfsdk:"cluster_id" json:"clusterId"`
	CpuHibernation types.Object `tfsdk:"cpu_hibernation" json:"cpuHibernation"`
	MonitorPort    types.List   `tfsdk:"monitor_port" json:"monitorPort"`
	UnusedPort     types.List   `tfsdk:"unused_port" json:"unusedPort"`
}

// NewLoadBatteryOptimizationDataSource returns a new instance of the generated data source.
func NewLoadBatteryOptimizationDataSource() datasource.DataSource {
	return &LoadBatteryOptimizationDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadBatteryOptimizationDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_battery_optimization"
}

// Schema returns the data source schema.
func (d *LoadBatteryOptimizationDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "load all battery optimization profile", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "cpu_hibernation": schema.SingleNestedAttribute{MarkdownDescription: "CPU hibernation battery optimization", Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.BoolAttribute{MarkdownDescription: "Enable/disable cpu hibernation battery optimization", Computed: true}, "sleep_in_mins": schema.Int64Attribute{MarkdownDescription: "Apply cpu hibernation battery optimization for this configured sleep time", Computed: true}, "threshold_level": schema.Int64Attribute{MarkdownDescription: "Apply cpu hibernation battery optimization at this battery level", Computed: true}}}, "monitor_port": schema.ListNestedAttribute{MarkdownDescription: "Monitor port off battery optimization", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"enabled": schema.BoolAttribute{MarkdownDescription: "Enable/disable port battery optimization", Computed: true}, "level": schema.StringAttribute{MarkdownDescription: "Apply port battery optimization at this battery level", Computed: true}, "port_group_id": schema.StringAttribute{MarkdownDescription: "Port group Id", Computed: true}}}}, "unused_port": schema.ListNestedAttribute{MarkdownDescription: "Unused port off battery optimization", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"enabled": schema.BoolAttribute{MarkdownDescription: "Enable/disable port battery optimization", Computed: true}, "level": schema.StringAttribute{MarkdownDescription: "Apply port battery optimization at this battery level", Computed: true}, "port_group_id": schema.StringAttribute{MarkdownDescription: "Port group Id", Computed: true}}}}}}
}

// Read fetches remote state into the data source model.
func (d *LoadBatteryOptimizationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadBatteryOptimizationDataSourceModel
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
func (d *LoadBatteryOptimizationDataSource) readRemote(ctx context.Context, config *LoadBatteryOptimizationDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/gtap/battery/optimization"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_battery_optimization", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_battery_optimization", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_battery_optimization", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_battery_optimization", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_battery_optimization", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_battery_optimization", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_battery_optimization", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_battery_optimization", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_battery_optimization", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_battery_optimization", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_load_battery_optimization", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_load_battery_optimization", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_battery_optimization", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_battery_optimization", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadBatteryOptimizationDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
