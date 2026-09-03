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
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*GetUpgradeInfoByTaskIdAndClusterIdDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetUpgradeInfoByTaskIdAndClusterIdDataSource)(nil)
)

// GetUpgradeInfoByTaskIdAndClusterIdDataSource is the generated Terraform data source implementation.
type GetUpgradeInfoByTaskIdAndClusterIdDataSource struct {
	client *client.Client
}

// GetUpgradeInfoByTaskIdAndClusterIdDataSourceModel describes the data source state shape.
type GetUpgradeInfoByTaskIdAndClusterIdDataSourceModel struct {
	ClusterId            types.String  `tfsdk:"cluster_id" json:"clusterId"`
	ClusterLog           types.List    `tfsdk:"cluster_log" json:"clusterLog"`
	EndTime              types.String  `tfsdk:"end_time" json:"endTime"`
	NodeGsUpgradeStatus  types.Dynamic `tfsdk:"node_gs_upgrade_status" json:"nodeGsUpgradeStatus"`
	NodeUpgradeStatus    types.List    `tfsdk:"node_upgrade_status" json:"nodeUpgradeStatus"`
	ObjectCounts         types.Object  `tfsdk:"object_counts" json:"objectCounts"`
	ObjectCountsExpected types.Object  `tfsdk:"object_counts_expected" json:"objectCountsExpected"`
	ObjectDiffReport     types.Object  `tfsdk:"object_diff_report" json:"objectDiffReport"`
	OperationState       types.String  `tfsdk:"operation_state" json:"operationState"`
	OverallUpgradeStatus types.String  `tfsdk:"overall_upgrade_status" json:"overallUpgradeStatus"`
	Stage                types.String  `tfsdk:"stage"`
	StartTime            types.String  `tfsdk:"start_time" json:"startTime"`
	Status               types.String  `tfsdk:"status"`
	TaskId               types.String  `tfsdk:"task_id" json:"taskId"`
	TaskName             types.String  `tfsdk:"task_name" json:"taskName"`
	UpgradeFlow          types.String  `tfsdk:"upgrade_flow" json:"upgradeFlow"`
	UpgradeSummary       types.Object  `tfsdk:"upgrade_summary" json:"upgradeSummary"`
}

// NewGetUpgradeInfoByTaskIdAndClusterIdDataSource returns a new instance of the generated data source.
func NewGetUpgradeInfoByTaskIdAndClusterIdDataSource() datasource.DataSource {
	return &GetUpgradeInfoByTaskIdAndClusterIdDataSource{}
}

// Metadata returns the data source type name.
func (d *GetUpgradeInfoByTaskIdAndClusterIdDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_upgrade_info_by_task_id_and_cluster_id"
}

// Schema returns the data source schema.
func (d *GetUpgradeInfoByTaskIdAndClusterIdDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get Upgrade Info By TaskId and ClusterId", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Cluster ID", Required: true}, "cluster_log": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Cluster ID", Computed: true}, "message": schema.StringAttribute{MarkdownDescription: "Message", Computed: true}, "node_id": schema.StringAttribute{MarkdownDescription: "Node ID", Computed: true}, "time": schema.StringAttribute{MarkdownDescription: "Time", Computed: true}}}}, "end_time": schema.StringAttribute{MarkdownDescription: "End Time", Computed: true}, "node_gs_upgrade_status": schema.DynamicAttribute{Computed: true}, "node_upgrade_status": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"current_step": schema.Int64Attribute{MarkdownDescription: "Current Step", Computed: true}, "current_version": schema.StringAttribute{MarkdownDescription: "Current Version", Computed: true}, "file_name": schema.StringAttribute{MarkdownDescription: "File Name", Computed: true}, "file_server": schema.StringAttribute{MarkdownDescription: "File Server", Computed: true}, "host_name": schema.StringAttribute{MarkdownDescription: "Host Name", Computed: true}, "last_completed_stage": schema.StringAttribute{MarkdownDescription: "Last Completed Stage", Computed: true}, "node_id": schema.StringAttribute{MarkdownDescription: "Node ID", Computed: true}, "node_log": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Cluster ID", Computed: true}, "message": schema.StringAttribute{MarkdownDescription: "Message", Computed: true}, "node_id": schema.StringAttribute{MarkdownDescription: "Node ID", Computed: true}, "time": schema.StringAttribute{MarkdownDescription: "Time", Computed: true}}}}, "num_steps_in_task": schema.Int64Attribute{MarkdownDescription: "Number of Steps In Task", Computed: true}, "prior_version": schema.StringAttribute{MarkdownDescription: "Prior Version", Computed: true}, "status": schema.StringAttribute{MarkdownDescription: "Status", Computed: true}, "step_status": schema.StringAttribute{MarkdownDescription: "Step Status", Computed: true}}}}, "object_counts": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"cards_up": schema.Int64Attribute{MarkdownDescription: "Number of Cards Up", Computed: true}, "circuit_ports_up": schema.Int64Attribute{MarkdownDescription: "Number of Circuit Ports Up", Computed: true}, "gigasmart_ports_up": schema.Int64Attribute{MarkdownDescription: "Number of Giga Smart Ports Up", Computed: true}, "gsops": schema.Int64Attribute{MarkdownDescription: "Number Gsops", Computed: true}, "hybrid_ports_up": schema.Int64Attribute{MarkdownDescription: "Number of Hybrid Ports Up", Computed: true}, "inline_network_ports_up": schema.Int64Attribute{MarkdownDescription: "Number of Inline Network Ports Up", Computed: true}, "inline_tool_ports_up": schema.Int64Attribute{MarkdownDescription: "Number of Inline Tool Ports Up", Computed: true}, "maps": schema.Int64Attribute{MarkdownDescription: "Number of Maps", Computed: true}, "network_ports_up": schema.Int64Attribute{MarkdownDescription: "Number of Network Ports Up", Computed: true}, "nodes": schema.Int64Attribute{MarkdownDescription: "Number of Nodes", Computed: true}, "stack_links": schema.Int64Attribute{MarkdownDescription: "Number of Stack Links", Computed: true}, "stack_ports_up": schema.Int64Attribute{MarkdownDescription: "Number of Stack Ports Up", Computed: true}, "tool_ports_up": schema.Int64Attribute{MarkdownDescription: "Number of Tool Ports Up", Computed: true}}}, "object_counts_expected": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"cards_up": schema.Int64Attribute{MarkdownDescription: "Number of Cards Up", Computed: true}, "circuit_ports_up": schema.Int64Attribute{MarkdownDescription: "Number of Circuit Ports Up", Computed: true}, "gigasmart_ports_up": schema.Int64Attribute{MarkdownDescription: "Number of Giga Smart Ports Up", Computed: true}, "gsops": schema.Int64Attribute{MarkdownDescription: "Number Gsops", Computed: true}, "hybrid_ports_up": schema.Int64Attribute{MarkdownDescription: "Number of Hybrid Ports Up", Computed: true}, "inline_network_ports_up": schema.Int64Attribute{MarkdownDescription: "Number of Inline Network Ports Up", Computed: true}, "inline_tool_ports_up": schema.Int64Attribute{MarkdownDescription: "Number of Inline Tool Ports Up", Computed: true}, "maps": schema.Int64Attribute{MarkdownDescription: "Number of Maps", Computed: true}, "network_ports_up": schema.Int64Attribute{MarkdownDescription: "Number of Network Ports Up", Computed: true}, "nodes": schema.Int64Attribute{MarkdownDescription: "Number of Nodes", Computed: true}, "stack_links": schema.Int64Attribute{MarkdownDescription: "Number of Stack Links", Computed: true}, "stack_ports_up": schema.Int64Attribute{MarkdownDescription: "Number of Stack Ports Up", Computed: true}, "tool_ports_up": schema.Int64Attribute{MarkdownDescription: "Number of Tool Ports Up", Computed: true}}}, "object_diff_report": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"cards_up": schema.ListNestedAttribute{MarkdownDescription: "Cards Up Differences", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"after_state": schema.StringAttribute{MarkdownDescription: "After State", Computed: true}, "object_name": schema.StringAttribute{MarkdownDescription: "Object Name", Computed: true}, "prev_state": schema.StringAttribute{MarkdownDescription: "Previous State", Computed: true}}}}, "circuit_ports_up": schema.ListNestedAttribute{MarkdownDescription: "Circuit Ports Up Differences", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"after_state": schema.StringAttribute{MarkdownDescription: "After State", Computed: true}, "object_name": schema.StringAttribute{MarkdownDescription: "Object Name", Computed: true}, "prev_state": schema.StringAttribute{MarkdownDescription: "Previous State", Computed: true}}}}, "gigasmart_ports_up": schema.ListNestedAttribute{MarkdownDescription: "Giga Smart Ports Up Differences", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"after_state": schema.StringAttribute{MarkdownDescription: "After State", Computed: true}, "object_name": schema.StringAttribute{MarkdownDescription: "Object Name", Computed: true}, "prev_state": schema.StringAttribute{MarkdownDescription: "Previous State", Computed: true}}}}, "gsops": schema.ListNestedAttribute{MarkdownDescription: "Gsops Differences", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"after_state": schema.StringAttribute{MarkdownDescription: "After State", Computed: true}, "object_name": schema.StringAttribute{MarkdownDescription: "Object Name", Computed: true}, "prev_state": schema.StringAttribute{MarkdownDescription: "Previous State", Computed: true}}}}, "hybrid_ports_up": schema.ListNestedAttribute{MarkdownDescription: "Hybrid Ports Up Differences", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"after_state": schema.StringAttribute{MarkdownDescription: "After State", Computed: true}, "object_name": schema.StringAttribute{MarkdownDescription: "Object Name", Computed: true}, "prev_state": schema.StringAttribute{MarkdownDescription: "Previous State", Computed: true}}}}, "inline_network_ports_up": schema.ListNestedAttribute{MarkdownDescription: "Inline Network Ports Up Differences", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"after_state": schema.StringAttribute{MarkdownDescription: "After State", Computed: true}, "object_name": schema.StringAttribute{MarkdownDescription: "Object Name", Computed: true}, "prev_state": schema.StringAttribute{MarkdownDescription: "Previous State", Computed: true}}}}, "inline_tool_ports_up": schema.ListNestedAttribute{MarkdownDescription: "Inline Tool Ports Up Differences", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"after_state": schema.StringAttribute{MarkdownDescription: "After State", Computed: true}, "object_name": schema.StringAttribute{MarkdownDescription: "Object Name", Computed: true}, "prev_state": schema.StringAttribute{MarkdownDescription: "Previous State", Computed: true}}}}, "maps": schema.ListNestedAttribute{MarkdownDescription: "Maps Differences", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"after_state": schema.StringAttribute{MarkdownDescription: "After State", Computed: true}, "object_name": schema.StringAttribute{MarkdownDescription: "Object Name", Computed: true}, "prev_state": schema.StringAttribute{MarkdownDescription: "Previous State", Computed: true}}}}, "network_ports_up": schema.ListNestedAttribute{MarkdownDescription: "Network Ports Up Differences", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"after_state": schema.StringAttribute{MarkdownDescription: "After State", Computed: true}, "object_name": schema.StringAttribute{MarkdownDescription: "Object Name", Computed: true}, "prev_state": schema.StringAttribute{MarkdownDescription: "Previous State", Computed: true}}}}, "nodes": schema.ListNestedAttribute{MarkdownDescription: "Nodes Differences", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"after_state": schema.StringAttribute{MarkdownDescription: "After State", Computed: true}, "object_name": schema.StringAttribute{MarkdownDescription: "Object Name", Computed: true}, "prev_state": schema.StringAttribute{MarkdownDescription: "Previous State", Computed: true}}}}, "stack_links": schema.ListNestedAttribute{MarkdownDescription: "Stack Links Differences", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"after_state": schema.StringAttribute{MarkdownDescription: "After State", Computed: true}, "object_name": schema.StringAttribute{MarkdownDescription: "Object Name", Computed: true}, "prev_state": schema.StringAttribute{MarkdownDescription: "Previous State", Computed: true}}}}, "stack_ports_up": schema.ListNestedAttribute{MarkdownDescription: "Stack Ports Up Differences", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"after_state": schema.StringAttribute{MarkdownDescription: "After State", Computed: true}, "object_name": schema.StringAttribute{MarkdownDescription: "Object Name", Computed: true}, "prev_state": schema.StringAttribute{MarkdownDescription: "Previous State", Computed: true}}}}, "tool_ports_up": schema.ListNestedAttribute{MarkdownDescription: "Tool Ports Up Differences", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"after_state": schema.StringAttribute{MarkdownDescription: "After State", Computed: true}, "object_name": schema.StringAttribute{MarkdownDescription: "Object Name", Computed: true}, "prev_state": schema.StringAttribute{MarkdownDescription: "Previous State", Computed: true}}}}}}, "operation_state": schema.StringAttribute{MarkdownDescription: "Operation State", Computed: true}, "overall_upgrade_status": schema.StringAttribute{MarkdownDescription: "Overall Upgrade Status", Computed: true}, "stage": schema.StringAttribute{MarkdownDescription: "Stage", Computed: true}, "start_time": schema.StringAttribute{MarkdownDescription: "Start Time", Computed: true}, "status": schema.StringAttribute{MarkdownDescription: "Status", Computed: true}, "task_id": schema.StringAttribute{MarkdownDescription: "Task ID", Required: true}, "task_name": schema.StringAttribute{MarkdownDescription: "Task Name", Computed: true}, "upgrade_flow": schema.StringAttribute{MarkdownDescription: "Upgrade Flow", Computed: true}, "upgrade_summary": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"num_activation_preparation_complete": schema.Int64Attribute{MarkdownDescription: "Activation Preparation Complete Count", Computed: true}, "num_config_backup_complete": schema.Int64Attribute{MarkdownDescription: "Config Backup Complete Count", Computed: true}, "num_init_fetch_complete": schema.Int64Attribute{MarkdownDescription: "Init Fetch Complete Count", Computed: true}, "num_initial_validation_complete": schema.Int64Attribute{MarkdownDescription: "CC-Card Sync Complete Count", Computed: true}, "num_install_complete": schema.Int64Attribute{MarkdownDescription: "Install Complete Count", Computed: true}, "num_nodes_success": schema.Int64Attribute{MarkdownDescription: "Nodes Success Count", Computed: true}, "num_post_upgrade_validation_complete": schema.Int64Attribute{MarkdownDescription: "Post Upgrade Validation Complete Count", Computed: true}, "num_reload_complete": schema.Int64Attribute{MarkdownDescription: "Reload Complete Count", Computed: true}, "num_reload_started": schema.Int64Attribute{MarkdownDescription: "Reload Started Count", Computed: true}, "num_uboot_complete": schema.Int64Attribute{MarkdownDescription: "UBoot Complete Count", Computed: true}, "num_upgrade_complete": schema.Int64Attribute{MarkdownDescription: "Upgrade Complete Count", Computed: true}, "num_verification_complete": schema.Int64Attribute{MarkdownDescription: "Verification Complete Count", Computed: true}}}}}
}

// Read fetches remote state into the data source model.
func (d *GetUpgradeInfoByTaskIdAndClusterIdDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetUpgradeInfoByTaskIdAndClusterIdDataSourceModel
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
func (d *GetUpgradeInfoByTaskIdAndClusterIdDataSource) readRemote(ctx context.Context, config *GetUpgradeInfoByTaskIdAndClusterIdDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/device/upgrade/orchestration/upgradeInfoByTaskId/{taskId}/{clusterId}"
	reqPath = strings.ReplaceAll(reqPath, "{taskId}", url.PathEscape(config.TaskId.ValueString()))
	reqPath = strings.ReplaceAll(reqPath, "{clusterId}", url.PathEscape(config.ClusterId.ValueString()))
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_info_by_task_id_and_cluster_id", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_info_by_task_id_and_cluster_id", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_info_by_task_id_and_cluster_id", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_info_by_task_id_and_cluster_id", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_info_by_task_id_and_cluster_id", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_info_by_task_id_and_cluster_id", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_info_by_task_id_and_cluster_id", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_info_by_task_id_and_cluster_id", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_info_by_task_id_and_cluster_id", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_info_by_task_id_and_cluster_id", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_info_by_task_id_and_cluster_id", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_info_by_task_id_and_cluster_id", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_info_by_task_id_and_cluster_id", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["clusterImageUpgradeStatus"]; ok {
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_upgrade_info_by_task_id_and_cluster_id", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetUpgradeInfoByTaskIdAndClusterIdDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
