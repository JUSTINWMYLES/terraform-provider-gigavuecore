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
	_ datasource.DataSource              = (*LoadAllVlanResourcesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadAllVlanResourcesDataSource)(nil)
)

// LoadAllVlanResourcesDataSource is the generated Terraform data source implementation.
type LoadAllVlanResourcesDataSource struct {
	client *client.Client
}

// LoadAllVlanResourcesDataSourceModel describes the data source state shape.
type LoadAllVlanResourcesDataSourceModel struct {
	AdvVlanManipCount              types.Int64  `tfsdk:"adv_vlan_manip_count" json:"advVlanManipCount"`
	AvailableVlanIds               types.List   `tfsdk:"available_vlan_ids" json:"availableVlanIds"`
	AvailableVlanIdsCommaSeparated types.String `tfsdk:"available_vlan_ids_comma_separated" json:"availableVlanIdsCommaSeparated"`
	CircuitTunnelCount             types.Int64  `tfsdk:"circuit_tunnel_count" json:"circuitTunnelCount"`
	ClusterId                      types.String `tfsdk:"cluster_id" json:"clusterId"`
	FlexInlineCount                types.Int64  `tfsdk:"flex_inline_count" json:"flexInlineCount"`
	GsMobilityCount                types.Int64  `tfsdk:"gs_mobility_count" json:"gsMobilityCount"`
	IpInterfaceCount               types.Int64  `tfsdk:"ip_interface_count" json:"ipInterfaceCount"`
	OtherFeatureCount              types.Int64  `tfsdk:"other_feature_count" json:"otherFeatureCount"`
	PVlanCount                     types.Int64  `tfsdk:"p_vlan_count" json:"pVlanCount"`
	SecureTunnelCount              types.Int64  `tfsdk:"secure_tunnel_count" json:"secureTunnelCount"`
	TotalCount                     types.Int64  `tfsdk:"total_count" json:"totalCount"`
	UnusedCount                    types.Int64  `tfsdk:"unused_count" json:"unusedCount"`
	UsedCount                      types.Int64  `tfsdk:"used_count" json:"usedCount"`
	UsedVlanIds                    types.List   `tfsdk:"used_vlan_ids" json:"usedVlanIds"`
	UsedVlanIdsCommaSeparated      types.String `tfsdk:"used_vlan_ids_comma_separated" json:"usedVlanIdsCommaSeparated"`
}

// NewLoadAllVlanResourcesDataSource returns a new instance of the generated data source.
func NewLoadAllVlanResourcesDataSource() datasource.DataSource {
	return &LoadAllVlanResourcesDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadAllVlanResourcesDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_all_vlan_resources"
}

// Schema returns the data source schema.
func (d *LoadAllVlanResourcesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load vlan resource details", Attributes: map[string]schema.Attribute{"adv_vlan_manip_count": schema.Int64Attribute{MarkdownDescription: "Total number of Vlan Ids used by Advanced VLAN Manipulation on this box", Computed: true}, "available_vlan_ids": schema.ListNestedAttribute{MarkdownDescription: "VlanIds that are available to be used on this box", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"value": schema.Int64Attribute{MarkdownDescription: "Represents Vlan ID value", Computed: true}, "value_max": schema.Int64Attribute{MarkdownDescription: "Represents maximum Vlan ID in a range and If present, value must be greater than 'value'", Computed: true}}}}, "available_vlan_ids_comma_separated": schema.StringAttribute{MarkdownDescription: "comma separated VLAN string", Computed: true}, "circuit_tunnel_count": schema.Int64Attribute{MarkdownDescription: "Total number of Vlan Ids used by Circuit Tunnel on this box", Computed: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "flex_inline_count": schema.Int64Attribute{MarkdownDescription: "Total number of Vlan Ids used by Flexible Inline Maps on this box", Computed: true}, "gs_mobility_count": schema.Int64Attribute{MarkdownDescription: "Total number of Vlan Ids used by Mobility on this box", Computed: true}, "ip_interface_count": schema.Int64Attribute{MarkdownDescription: "Total number of Vlan Ids used by IP Interface on this box", Computed: true}, "other_feature_count": schema.Int64Attribute{MarkdownDescription: "Total number of Vlan Ids used by Other Features on this box", Computed: true}, "p_vlan_count": schema.Int64Attribute{MarkdownDescription: "Total number of Port Vlan Ids used on this box", Computed: true}, "secure_tunnel_count": schema.Int64Attribute{MarkdownDescription: "Total number of Vlan Ids used by Secure Tunnel on this box", Computed: true}, "total_count": schema.Int64Attribute{MarkdownDescription: "Total number of Vlan Ids on this box", Computed: true}, "unused_count": schema.Int64Attribute{MarkdownDescription: "Number of Vlan Ids available as unused on this box", Computed: true}, "used_count": schema.Int64Attribute{MarkdownDescription: "Number of Vlan Ids being used on this box", Computed: true}, "used_vlan_ids": schema.ListNestedAttribute{MarkdownDescription: "VlanIds that are being used on this box", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"value": schema.Int64Attribute{MarkdownDescription: "Represents Vlan ID value", Computed: true}, "value_max": schema.Int64Attribute{MarkdownDescription: "Represents maximum Vlan ID in a range and If present, value must be greater than 'value'", Computed: true}}}}, "used_vlan_ids_comma_separated": schema.StringAttribute{MarkdownDescription: "comma separated VLAN string", Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadAllVlanResourcesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadAllVlanResourcesDataSourceModel
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
func (d *LoadAllVlanResourcesDataSource) readRemote(ctx context.Context, config *LoadAllVlanResourcesDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/vlanResources"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_vlan_resources", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_vlan_resources", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_vlan_resources", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_vlan_resources", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_vlan_resources", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_vlan_resources", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_vlan_resources", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_vlan_resources", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_vlan_resources", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_vlan_resources", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_load_all_vlan_resources", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_vlan_resources", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_vlan_resources", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["vlanResourcesDef"]; ok {
		if m, ok := v.(map[string]any); ok {
			data = m
		} else if arr, ok := v.([]any); ok && len(arr) > 0 {
			if m, ok := arr[0].(map[string]any); ok {
				data = m
			}
		}
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_vlan_resources", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadAllVlanResourcesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
