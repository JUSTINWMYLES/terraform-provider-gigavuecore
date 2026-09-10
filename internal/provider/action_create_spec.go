package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)
import (
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*CreateSpecAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*CreateSpecAction)(nil)

// CreateSpecAction is the generated Terraform action implementation.
type CreateSpecAction struct {
	client *client.Client
}

// CreateSpecActionModel describes the action configuration shape.
type CreateSpecActionModel struct {
	ActivateStage           types.Bool    `tfsdk:"activate_stage" json:"activateStage"`
	ClusterIds              types.Dynamic `tfsdk:"cluster_ids" json:"clusterIds"`
	ConfigBackup            types.Bool    `tfsdk:"config_backup" json:"configBackup"`
	ExecClusterCounter      types.Int64   `tfsdk:"exec_cluster_counter" json:"execClusterCounter"`
	FetchStage              types.Bool    `tfsdk:"fetch_stage" json:"fetchStage"`
	ImageFileSpecs          types.List    `tfsdk:"image_file_specs" json:"imageFileSpecs"`
	ImageServer             types.String  `tfsdk:"image_server" json:"imageServer"`
	InstallStage            types.Bool    `tfsdk:"install_stage" json:"installStage"`
	NodeIds                 types.Dynamic `tfsdk:"node_ids" json:"nodeIds"`
	PostCheckStage          types.Bool    `tfsdk:"post_check_stage" json:"postCheckStage"`
	Reboot                  types.Bool    `tfsdk:"reboot"`
	SkipNotReachableDevices types.Bool    `tfsdk:"skip_not_reachable_devices" json:"skipNotReachableDevices"`
	Tags                    types.Dynamic `tfsdk:"tags"`
	TaskId                  types.String  `tfsdk:"task_id" json:"taskId"`
	TaskName                types.String  `tfsdk:"task_name" json:"taskName"`
	UpgradeState            types.String  `tfsdk:"upgrade_state" json:"upgradeState"`
	Version                 types.String  `tfsdk:"version"`
}

// NewCreateSpecAction returns a new instance of the generated action.
func NewCreateSpecAction() action.Action {
	return &CreateSpecAction{}
}

// Metadata returns the action type name.
func (r *CreateSpecAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_create_spec"
}

// Schema returns the action schema.
func (r *CreateSpecAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Create Device Upgrade Specification", Attributes: map[string]schema.Attribute{"activate_stage": schema.BoolAttribute{MarkdownDescription: "Activate Stage", Optional: true}, "cluster_ids": schema.DynamicAttribute{MarkdownDescription: "Cluster IDs", Required: true}, "config_backup": schema.BoolAttribute{MarkdownDescription: "Config Backup", Optional: true}, "exec_cluster_counter": schema.Int64Attribute{MarkdownDescription: "Exec Cluster Counter", Optional: true}, "fetch_stage": schema.BoolAttribute{MarkdownDescription: "Fetch Stage", Optional: true}, "image_file_specs": schema.ListNestedAttribute{MarkdownDescription: "Image File Specs", Required: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"device_model": schema.StringAttribute{MarkdownDescription: "Device Model", Required: true}, "device_model_qualifier": schema.StringAttribute{MarkdownDescription: "Device Model Qualifier", Required: true}, "file_path": schema.StringAttribute{MarkdownDescription: "File Path", Required: true}, "file_type": schema.StringAttribute{MarkdownDescription: "File Type", Required: true}, "target_slot": schema.StringAttribute{MarkdownDescription: "Target Slot", Required: true}, "update_uboot": schema.StringAttribute{MarkdownDescription: "Update UBoot", Optional: true}}}}, "image_server": schema.StringAttribute{MarkdownDescription: "Image Server", Required: true}, "install_stage": schema.BoolAttribute{MarkdownDescription: "Install Stage", Optional: true}, "node_ids": schema.DynamicAttribute{MarkdownDescription: "Node IDs", Optional: true}, "post_check_stage": schema.BoolAttribute{MarkdownDescription: "Post Check Stage", Optional: true}, "reboot": schema.BoolAttribute{MarkdownDescription: "Reboot", Optional: true}, "skip_not_reachable_devices": schema.BoolAttribute{MarkdownDescription: "Skip Not Reachable Devices", Optional: true}, "tags": schema.DynamicAttribute{Optional: true}, "task_id": schema.StringAttribute{MarkdownDescription: "Task ID", Required: true}, "task_name": schema.StringAttribute{MarkdownDescription: "Task Name", Optional: true}, "upgrade_state": schema.StringAttribute{MarkdownDescription: "Upgrade State", Optional: true}, "version": schema.StringAttribute{MarkdownDescription: "Version", Optional: true}}}
}

// Invoke executes the action against the remote API.
func (r *CreateSpecAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config CreateSpecActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *CreateSpecAction) invokeRemote(ctx context.Context, config *CreateSpecActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/device/upgrade/orchestration/createSpec"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_create_spec", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_create_spec", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_create_spec", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_create_spec", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_create_spec", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_create_spec", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_create_spec", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_create_spec", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_create_spec", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_create_spec", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_create_spec", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_create_spec", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_create_spec", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *CreateSpecAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Action Configure Type", fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData))
		return
	}
	r.client = c
}
