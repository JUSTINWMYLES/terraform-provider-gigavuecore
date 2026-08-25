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
	_ datasource.DataSource              = (*GetAllClusterConfigImageUpgradeStatusDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetAllClusterConfigImageUpgradeStatusDataSource)(nil)
)

// GetAllClusterConfigImageUpgradeStatusDataSource is the generated Terraform data source implementation.
type GetAllClusterConfigImageUpgradeStatusDataSource struct {
	client *client.Client
}

// GetAllClusterConfigImageUpgradeStatusDataSourceModel describes the data source state shape.
type GetAllClusterConfigImageUpgradeStatusDataSourceModel struct {
	ClustersStatus types.List   `tfsdk:"clusters_status" json:"clustersStatus"`
	Context        types.Object `tfsdk:"context"`
}

// NewGetAllClusterConfigImageUpgradeStatusDataSource returns a new instance of the generated data source.
func NewGetAllClusterConfigImageUpgradeStatusDataSource() datasource.DataSource {
	return &GetAllClusterConfigImageUpgradeStatusDataSource{}
}

// Metadata returns the data source type name.
func (d *GetAllClusterConfigImageUpgradeStatusDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_all_cluster_config_image_upgrade_status"
}

// Schema returns the data source schema.
func (d *GetAllClusterConfigImageUpgradeStatusDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "get all the cluster configuration imageUpgrade status", Attributes: map[string]schema.Attribute{"clusters_status": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Id of the Target cluster", Computed: true}, "cluster_log": schema.ListNestedAttribute{MarkdownDescription: "list of the cluster Log", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Id of the Target cluster", Computed: true}, "message": schema.StringAttribute{MarkdownDescription: "Read-only. Describes the reason for image upgrade state", Computed: true}, "node_id": schema.StringAttribute{MarkdownDescription: "unique ID representing device. This will be available only for nodeLog", Computed: true}, "time": schema.StringAttribute{MarkdownDescription: "In ISO-8601 date format 'yyyy-MM-dd'T'HH:mm:ssZ'", Computed: true}}}}, "end_time": schema.StringAttribute{MarkdownDescription: "End Time in UTC format", Computed: true}, "node_upgrade_status": schema.ListNestedAttribute{MarkdownDescription: "list of the nodeUpgrade status", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"current_step": schema.Int64Attribute{MarkdownDescription: "current step number of task", Computed: true}, "current_version": schema.Int64Attribute{MarkdownDescription: "Current Device Version ", Computed: true}, "file_name": schema.StringAttribute{MarkdownDescription: "Image filename", Computed: true}, "file_server": schema.StringAttribute{MarkdownDescription: "IpAddress of the Destination Image file server ", Computed: true}, "hostname": schema.StringAttribute{MarkdownDescription: "hostname or IP address", Computed: true}, "node_id": schema.StringAttribute{MarkdownDescription: "unique ID representing device. This will be available in nodeLog", Computed: true}, "node_log": schema.ListNestedAttribute{MarkdownDescription: "list of the node Log", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Id of the Target cluster", Computed: true}, "message": schema.StringAttribute{MarkdownDescription: "Read-only. Describes the reason for image upgrade state", Computed: true}, "node_id": schema.StringAttribute{MarkdownDescription: "unique ID representing device. This will be available only for nodeLog", Computed: true}, "time": schema.StringAttribute{MarkdownDescription: "In ISO-8601 date format 'yyyy-MM-dd'T'HH:mm:ssZ'", Computed: true}}}}, "num_steps_in_task": schema.Int64Attribute{MarkdownDescription: "total number of steps in task", Computed: true}, "prior_version": schema.Int64Attribute{MarkdownDescription: "Device Version before upgrade", Computed: true}, "status": schema.StringAttribute{MarkdownDescription: "node upgrade status", Computed: true}, "step_status": schema.StringAttribute{MarkdownDescription: "Describes the status of current step ", Computed: true}}}}, "object_counts": schema.SingleNestedAttribute{MarkdownDescription: "Cluster objectCounts", Computed: true, Attributes: map[string]schema.Attribute{"cards_up": schema.Int64Attribute{MarkdownDescription: "total number of cards in up state", Computed: true}, "circuit_ports_up": schema.Int64Attribute{MarkdownDescription: "total number of circuit ports in up state", Computed: true}, "gigasmart_ports_up": schema.Int64Attribute{MarkdownDescription: "total number of gigasmart ports in up state", Computed: true}, "gsops": schema.Int64Attribute{MarkdownDescription: "total number of gigasmart operations", Computed: true}, "hybrid_ports_up": schema.Int64Attribute{MarkdownDescription: "total number of hybrid ports in up state", Computed: true}, "inline_network_ports_up": schema.Int64Attribute{MarkdownDescription: "total number of inlineNetwork ports in up state ", Computed: true}, "inline_tool_ports_up": schema.Int64Attribute{MarkdownDescription: "total number of inlineTool ports in up state ", Computed: true}, "maps": schema.Int64Attribute{MarkdownDescription: "total number of maps", Computed: true}, "network_ports_up": schema.Int64Attribute{MarkdownDescription: "total number of network ports in up state", Computed: true}, "nodes": schema.Int64Attribute{MarkdownDescription: "total number of nodes", Computed: true}, "stack_links": schema.Int64Attribute{MarkdownDescription: "total number of stacklinks", Computed: true}, "stack_ports_up": schema.Int64Attribute{MarkdownDescription: "total number of stack ports in up state", Computed: true}, "tool_ports_up": schema.Int64Attribute{MarkdownDescription: "total number of tool ports in up state", Computed: true}}}, "object_counts_expected": schema.SingleNestedAttribute{MarkdownDescription: "Cluster objectCounts Expected", Computed: true, Attributes: map[string]schema.Attribute{"cards_up": schema.Int64Attribute{MarkdownDescription: "total number of cards in up state", Computed: true}, "circuit_ports_up": schema.Int64Attribute{MarkdownDescription: "total number of circuit ports in up state", Computed: true}, "gigasmart_ports_up": schema.Int64Attribute{MarkdownDescription: "total number of gigasmart ports in up state", Computed: true}, "gsops": schema.Int64Attribute{MarkdownDescription: "total number of gigasmart operations", Computed: true}, "hybrid_ports_up": schema.Int64Attribute{MarkdownDescription: "total number of hybrid ports in up state", Computed: true}, "inline_network_ports_up": schema.Int64Attribute{MarkdownDescription: "total number of inlineNetwork ports in up state ", Computed: true}, "inline_tool_ports_up": schema.Int64Attribute{MarkdownDescription: "total number of inlineTool ports in up state ", Computed: true}, "maps": schema.Int64Attribute{MarkdownDescription: "total number of maps", Computed: true}, "network_ports_up": schema.Int64Attribute{MarkdownDescription: "total number of network ports in up state", Computed: true}, "nodes": schema.Int64Attribute{MarkdownDescription: "total number of nodes", Computed: true}, "stack_links": schema.Int64Attribute{MarkdownDescription: "total number of stacklinks", Computed: true}, "stack_ports_up": schema.Int64Attribute{MarkdownDescription: "total number of stack ports in up state", Computed: true}, "tool_ports_up": schema.Int64Attribute{MarkdownDescription: "total number of tool ports in up state", Computed: true}}}, "object_diff_report": schema.SingleNestedAttribute{MarkdownDescription: "Cluster object Diff report", Computed: true, Attributes: map[string]schema.Attribute{"cards_up_differences": schema.ListNestedAttribute{MarkdownDescription: "Diff in cards in up state", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"after_state": schema.StringAttribute{MarkdownDescription: "next state of the object", Computed: true}, "object_name": schema.StringAttribute{MarkdownDescription: "alias of the object", Computed: true}, "prev_state": schema.StringAttribute{MarkdownDescription: "previous state of the object", Computed: true}}}}, "circuit_ports_up_differences": schema.ListNestedAttribute{MarkdownDescription: "Diff in circuit ports in up state", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"after_state": schema.StringAttribute{MarkdownDescription: "next state of the object", Computed: true}, "object_name": schema.StringAttribute{MarkdownDescription: "alias of the object", Computed: true}, "prev_state": schema.StringAttribute{MarkdownDescription: "previous state of the object", Computed: true}}}}, "gigasmart_ports_up_differences": schema.ListNestedAttribute{MarkdownDescription: "Diff in gigasmart ports in up state", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"after_state": schema.StringAttribute{MarkdownDescription: "next state of the object", Computed: true}, "object_name": schema.StringAttribute{MarkdownDescription: "alias of the object", Computed: true}, "prev_state": schema.StringAttribute{MarkdownDescription: "previous state of the object", Computed: true}}}}, "gsop_differences": schema.ListNestedAttribute{MarkdownDescription: "Diff in gigasmart operations", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"after_state": schema.StringAttribute{MarkdownDescription: "next state of the object", Computed: true}, "object_name": schema.StringAttribute{MarkdownDescription: "alias of the object", Computed: true}, "prev_state": schema.StringAttribute{MarkdownDescription: "previous state of the object", Computed: true}}}}, "hybrid_ports_up_differences": schema.ListNestedAttribute{MarkdownDescription: "Diff in hybrid ports in up state", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"after_state": schema.StringAttribute{MarkdownDescription: "next state of the object", Computed: true}, "object_name": schema.StringAttribute{MarkdownDescription: "alias of the object", Computed: true}, "prev_state": schema.StringAttribute{MarkdownDescription: "previous state of the object", Computed: true}}}}, "inline_network_ports_up_differences": schema.ListNestedAttribute{MarkdownDescription: "Diff in inlineNetwork ports in up state ", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"after_state": schema.StringAttribute{MarkdownDescription: "next state of the object", Computed: true}, "object_name": schema.StringAttribute{MarkdownDescription: "alias of the object", Computed: true}, "prev_state": schema.StringAttribute{MarkdownDescription: "previous state of the object", Computed: true}}}}, "inline_tool_ports_up_differences": schema.ListNestedAttribute{MarkdownDescription: "Diff in inlineTool ports in up state ", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"after_state": schema.StringAttribute{MarkdownDescription: "next state of the object", Computed: true}, "object_name": schema.StringAttribute{MarkdownDescription: "alias of the object", Computed: true}, "prev_state": schema.StringAttribute{MarkdownDescription: "previous state of the object", Computed: true}}}}, "maps_differences": schema.ListNestedAttribute{MarkdownDescription: "Diff in maps", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"after_state": schema.StringAttribute{MarkdownDescription: "next state of the object", Computed: true}, "object_name": schema.StringAttribute{MarkdownDescription: "alias of the object", Computed: true}, "prev_state": schema.StringAttribute{MarkdownDescription: "previous state of the object", Computed: true}}}}, "network_ports_up_differences": schema.ListNestedAttribute{MarkdownDescription: "Diff in network ports in up state", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"after_state": schema.StringAttribute{MarkdownDescription: "next state of the object", Computed: true}, "object_name": schema.StringAttribute{MarkdownDescription: "alias of the object", Computed: true}, "prev_state": schema.StringAttribute{MarkdownDescription: "previous state of the object", Computed: true}}}}, "node_differences": schema.ListNestedAttribute{MarkdownDescription: "Diff in nodes", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"after_state": schema.StringAttribute{MarkdownDescription: "next state of the object", Computed: true}, "object_name": schema.StringAttribute{MarkdownDescription: "alias of the object", Computed: true}, "prev_state": schema.StringAttribute{MarkdownDescription: "previous state of the object", Computed: true}}}}, "stack_links_differences": schema.ListNestedAttribute{MarkdownDescription: "Diff in stacklinks", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"after_state": schema.StringAttribute{MarkdownDescription: "next state of the object", Computed: true}, "object_name": schema.StringAttribute{MarkdownDescription: "alias of the object", Computed: true}, "prev_state": schema.StringAttribute{MarkdownDescription: "previous state of the object", Computed: true}}}}, "stack_ports_up_differences": schema.ListNestedAttribute{MarkdownDescription: "Diff in stack ports in up state", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"after_state": schema.StringAttribute{MarkdownDescription: "next state of the object", Computed: true}, "object_name": schema.StringAttribute{MarkdownDescription: "alias of the object", Computed: true}, "prev_state": schema.StringAttribute{MarkdownDescription: "previous state of the object", Computed: true}}}}, "tool_ports_up_differences": schema.ListNestedAttribute{MarkdownDescription: "Diff in tool ports in up state", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"after_state": schema.StringAttribute{MarkdownDescription: "next state of the object", Computed: true}, "object_name": schema.StringAttribute{MarkdownDescription: "alias of the object", Computed: true}, "prev_state": schema.StringAttribute{MarkdownDescription: "previous state of the object", Computed: true}}}}}}, "start_time": schema.StringAttribute{MarkdownDescription: "Start Time in UTC format", Computed: true}, "status": schema.StringAttribute{MarkdownDescription: "Cluster Image Upgrade Status", Computed: true}, "task_id": schema.StringAttribute{MarkdownDescription: "taskId of the task", Computed: true}, "task_name": schema.StringAttribute{MarkdownDescription: "name of the task", Computed: true}, "upgrade_summary": schema.SingleNestedAttribute{MarkdownDescription: "Cluster upgrade Summary", Computed: true, Attributes: map[string]schema.Attribute{"num_initial_validation_complete": schema.Int64Attribute{MarkdownDescription: "total number of items in the initial validation complete state", Computed: true}, "num_install_complete": schema.Int64Attribute{MarkdownDescription: "total number of items in the installation complete state", Computed: true}, "num_nodes_success": schema.Int64Attribute{MarkdownDescription: "total number of items in the nodes success state ", Computed: true}, "num_reload_complete": schema.Int64Attribute{MarkdownDescription: "total number of items in the reload complete state", Computed: true}, "num_uboot_complete": schema.Int64Attribute{MarkdownDescription: "total number of items in the Uboot complete state", Computed: true}, "num_upgrade_complete": schema.Int64Attribute{MarkdownDescription: "total number of items in the upgrade complete state", Computed: true}, "num_verification_complete": schema.Int64Attribute{MarkdownDescription: "total number of items in the verification complete state", Computed: true}}}}}}, "context": schema.SingleNestedAttribute{MarkdownDescription: "Gigamon query result context", Computed: true, Attributes: map[string]schema.Attribute{"page_no": schema.Int64Attribute{MarkdownDescription: "page number of the returned result set", Computed: true}, "page_size": schema.Int64Attribute{MarkdownDescription: "page size of the returned result set", Computed: true}, "sort": schema.ListAttribute{MarkdownDescription: "sorting info of the returned result set. list of fields in the array indicate sorting order", Computed: true, ElementType: types.StringType}, "total_items": schema.Int64Attribute{MarkdownDescription: "total number of items in the queried entity type", Computed: true}}}}}
}

// Read fetches remote state into the data source model.
func (d *GetAllClusterConfigImageUpgradeStatusDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetAllClusterConfigImageUpgradeStatusDataSourceModel
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
func (d *GetAllClusterConfigImageUpgradeStatusDataSource) readRemote(ctx context.Context, config *GetAllClusterConfigImageUpgradeStatusDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/clusterConfig/imageUpgrade/status"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_config_image_upgrade_status", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_config_image_upgrade_status", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_config_image_upgrade_status", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_config_image_upgrade_status", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_config_image_upgrade_status", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_config_image_upgrade_status", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_config_image_upgrade_status", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_config_image_upgrade_status", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_config_image_upgrade_status", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_config_image_upgrade_status", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_config_image_upgrade_status", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_config_image_upgrade_status", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_config_image_upgrade_status", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_cluster_config_image_upgrade_status", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetAllClusterConfigImageUpgradeStatusDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
