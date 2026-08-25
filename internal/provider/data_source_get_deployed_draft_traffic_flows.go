package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)
import (
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*GetDeployedDraftTrafficFlowsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetDeployedDraftTrafficFlowsDataSource)(nil)
)

// GetDeployedDraftTrafficFlowsDataSource is the generated Terraform data source implementation.
type GetDeployedDraftTrafficFlowsDataSource struct {
	client *client.Client
}

// GetDeployedDraftTrafficFlowsDataSourceModel describes the data source state shape.
type GetDeployedDraftTrafficFlowsDataSourceModel struct {
	Alias              types.String  `tfsdk:"alias"`
	Comment            types.String  `tfsdk:"comment"`
	ConfigStatus       types.String  `tfsdk:"config_status" json:"configStatus"`
	CreatedBy          types.String  `tfsdk:"created_by" json:"createdBy"`
	CreatedTime        types.Float64 `tfsdk:"created_time" json:"createdTime"`
	DeployedBy         types.String  `tfsdk:"deployed_by" json:"deployedBy"`
	DeployedTime       types.Float64 `tfsdk:"deployed_time" json:"deployedTime"`
	DeploymentType     types.String  `tfsdk:"deployment_type" json:"deploymentType"`
	Enable             types.Bool    `tfsdk:"enable"`
	ErrorMessage       types.String  `tfsdk:"error_message" json:"errorMessage"`
	Flows              types.Dynamic `tfsdk:"flows"`
	HasDraft           types.Bool    `tfsdk:"has_draft" json:"hasDraft"`
	HealthState        types.String  `tfsdk:"health_state" json:"healthState"`
	HealthStateReasons types.List    `tfsdk:"health_state_reasons" json:"healthStateReasons"`
	PolicyId           types.String  `tfsdk:"policy_id" json:"policyId"`
	SourcesAndRules    types.Dynamic `tfsdk:"sources_and_rules" json:"sourcesAndRules"`
	Tags               types.List    `tfsdk:"tags"`
	UpdatedBy          types.String  `tfsdk:"updated_by" json:"updatedBy"`
	UpdatedTime        types.Float64 `tfsdk:"updated_time" json:"updatedTime"`
}

// NewGetDeployedDraftTrafficFlowsDataSource returns a new instance of the generated data source.
func NewGetDeployedDraftTrafficFlowsDataSource() datasource.DataSource {
	return &GetDeployedDraftTrafficFlowsDataSource{}
}

// Metadata returns the data source type name.
func (d *GetDeployedDraftTrafficFlowsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_deployed_draft_traffic_flows"
}

// Schema returns the data source schema.
func (d *GetDeployedDraftTrafficFlowsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get deployed draft Traffic Flows by alias", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "unique map alias", Required: true}, "comment": schema.StringAttribute{Computed: true}, "config_status": schema.StringAttribute{MarkdownDescription: "Configuration status of this Traffic flow.", Computed: true}, "created_by": schema.StringAttribute{MarkdownDescription: "created by of the traffic flow", Computed: true}, "created_time": schema.Float64Attribute{MarkdownDescription: "created name of the traffic flow", Computed: true}, "deployed_by": schema.StringAttribute{MarkdownDescription: "deployed by of the traffic flow", Computed: true}, "deployed_time": schema.Float64Attribute{MarkdownDescription: "deployed name of the traffic flow", Computed: true}, "deployment_type": schema.StringAttribute{MarkdownDescription: "'regular' maps are from network/hybrid ports to tool/hybrid/gigastream; 'inline' maps are from inline ports to inline/tool/hybrid/gigastream; 'firstLevel' are from network/hybrid ports to vPorts/tool/hybrid/gigastream; 'secondLevel' maps are from vPorts to tool/hybrid/gigastream; 'inlineFirstLevel is from inline network to vport; 'inlineSecondLevel' is from vport to inline tool; 'transitLevel' is from vport to vport", Computed: true}, "enable": schema.BoolAttribute{MarkdownDescription: "enable/disable map, applicable only to first level maps", Computed: true}, "error_message": schema.StringAttribute{MarkdownDescription: "In case of configuration failure, this message provides details about the possible cause of the failure.", Computed: true}, "flows": schema.DynamicAttribute{Computed: true}, "has_draft": schema.BoolAttribute{Computed: true}, "health_state": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "health_state_reasons": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"message": schema.StringAttribute{MarkdownDescription: "Read-only. Describes the reason for component's health state", Computed: true}, "severity": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "traffic_health_state_computation_type": schema.StringAttribute{MarkdownDescription: "Traffic Health State Computation Type", Computed: true}}}}, "policy_id": schema.StringAttribute{MarkdownDescription: "unique id", Computed: true}, "sources_and_rules": schema.DynamicAttribute{Computed: true}, "tags": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"tag_key": schema.StringAttribute{MarkdownDescription: "Name of the tag", Computed: true}, "tag_values": schema.ListAttribute{MarkdownDescription: "All possible values of the tag", Computed: true, ElementType: types.StringType}}}}, "updated_by": schema.StringAttribute{MarkdownDescription: "updated by of the traffic flow", Computed: true}, "updated_time": schema.Float64Attribute{MarkdownDescription: "updated time of the traffic flow", Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetDeployedDraftTrafficFlowsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetDeployedDraftTrafficFlowsDataSourceModel
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
func (d *GetDeployedDraftTrafficFlowsDataSource) readRemote(ctx context.Context, config *GetDeployedDraftTrafficFlowsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/traffic-flows/{alias}/deployedDraft"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(config.Alias.ValueString()))
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_deployed_draft_traffic_flows", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_deployed_draft_traffic_flows", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_deployed_draft_traffic_flows", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_deployed_draft_traffic_flows", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_deployed_draft_traffic_flows", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_deployed_draft_traffic_flows", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_deployed_draft_traffic_flows", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_deployed_draft_traffic_flows", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_deployed_draft_traffic_flows", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_deployed_draft_traffic_flows", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_deployed_draft_traffic_flows", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_deployed_draft_traffic_flows", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_deployed_draft_traffic_flows", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["trafficFlow"]; ok {
		if m, ok := v.(map[string]any); ok {
			data = m
		}
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_deployed_draft_traffic_flows", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetDeployedDraftTrafficFlowsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
