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
	_ datasource.DataSource              = (*ListFmStoredClusterConfigBackupSnapshotDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*ListFmStoredClusterConfigBackupSnapshotDataSource)(nil)
)

// ListFmStoredClusterConfigBackupSnapshotDataSource is the generated Terraform data source implementation.
type ListFmStoredClusterConfigBackupSnapshotDataSource struct {
	client *client.Client
}

// ListFmStoredClusterConfigBackupSnapshotDataSourceModel describes the data source state shape.
type ListFmStoredClusterConfigBackupSnapshotDataSourceModel struct {
	Alias          types.String `tfsdk:"alias"`
	BackupId       types.String `tfsdk:"backup_id" json:"backupId"`
	BackupTime     types.String `tfsdk:"backup_time" json:"backupTime"`
	ClusterId      types.String `tfsdk:"cluster_id" json:"clusterId"`
	ClusterName    types.String `tfsdk:"cluster_name" json:"clusterName"`
	ClusterVip     types.String `tfsdk:"cluster_vip" json:"clusterVip"`
	Comment        types.String `tfsdk:"comment"`
	DeviceGlobalId types.String `tfsdk:"device_global_id" json:"deviceGlobalId"`
	DeviceHostName types.String `tfsdk:"device_host_name" json:"deviceHostName"`
	DeviceIp       types.String `tfsdk:"device_ip" json:"deviceIp"`
	DoNotPurge     types.Bool   `tfsdk:"do_not_purge" json:"doNotPurge"`
	FileMetaData   types.Object `tfsdk:"file_meta_data" json:"fileMetaData"`
	Format         types.String `tfsdk:"format"`
	RestoreLogs    types.List   `tfsdk:"restore_logs" json:"restoreLogs"`
	SwVersion      types.String `tfsdk:"sw_version" json:"swVersion"`
	Tags           types.List   `tfsdk:"tags"`
}

// NewListFmStoredClusterConfigBackupSnapshotDataSource returns a new instance of the generated data source.
func NewListFmStoredClusterConfigBackupSnapshotDataSource() datasource.DataSource {
	return &ListFmStoredClusterConfigBackupSnapshotDataSource{}
}

// Metadata returns the data source type name.
func (d *ListFmStoredClusterConfigBackupSnapshotDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_list_fm_stored_cluster_config_backup_snapshot"
}

// Schema returns the data source schema.
func (d *ListFmStoredClusterConfigBackupSnapshotDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Lists a labeled config backup snapshot", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "User-entered alias for this snapshot", Computed: true}, "backup_id": schema.StringAttribute{MarkdownDescription: "Unique identifier of the cluster config snapshot for a given cluster", Required: true}, "backup_time": schema.StringAttribute{MarkdownDescription: "Configuration file backup time", Computed: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster Id", Required: true}, "cluster_name": schema.StringAttribute{MarkdownDescription: "Name of the cluster", Computed: true}, "cluster_vip": schema.StringAttribute{MarkdownDescription: "Virtual Ip of the cluster", Computed: true}, "comment": schema.StringAttribute{MarkdownDescription: "User-entered comments for this snapshot", Computed: true}, "device_global_id": schema.StringAttribute{MarkdownDescription: "Device Global Id of the hosting node", Computed: true}, "device_host_name": schema.StringAttribute{MarkdownDescription: "Device HostName of the hosting node", Computed: true}, "device_ip": schema.StringAttribute{MarkdownDescription: "Device Ip of the hosting node", Computed: true}, "do_not_purge": schema.BoolAttribute{MarkdownDescription: "Indicates whether this snapshot should not be auto-aged/purged", Computed: true}, "file_meta_data": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"device_global_id": schema.StringAttribute{MarkdownDescription: "Device Global Id of the hosting node", Computed: true}, "device_host_name": schema.StringAttribute{MarkdownDescription: "Device HostName of the hosting node", Computed: true}, "device_ip": schema.StringAttribute{MarkdownDescription: "Device Ip of the hosting node", Computed: true}, "device_model": schema.StringAttribute{MarkdownDescription: "Gigamon physical device models", Computed: true}, "device_version": schema.StringAttribute{MarkdownDescription: "Device's current Software version", Computed: true}, "do_not_purge": schema.BoolAttribute{MarkdownDescription: "Indicates whether this snapshot should not be auto-aged/purged", Computed: true}, "file_name": schema.StringAttribute{MarkdownDescription: "Config file name", Computed: true}, "member_info": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"box_id": schema.StringAttribute{MarkdownDescription: "Cluster BoxId of the node", Computed: true}, "host_name": schema.StringAttribute{MarkdownDescription: "Device HostName of the hosting node", Computed: true}, "role": schema.StringAttribute{MarkdownDescription: "Role of the node. (deprecated: use roleAlias)", Computed: true}, "role_alias": schema.StringAttribute{MarkdownDescription: "Role of the node", Computed: true}}}}}}, "format": schema.StringAttribute{MarkdownDescription: "format of the Backup file", Computed: true}, "restore_logs": schema.ListAttribute{MarkdownDescription: "restoreLogs associated with this file", Computed: true, ElementType: types.StringType}, "sw_version": schema.StringAttribute{MarkdownDescription: "Software version of the device this snapshot was taken from", Computed: true}, "tags": schema.ListAttribute{MarkdownDescription: "Tags associated with this file", Computed: true, ElementType: types.StringType}}}
}

// Read fetches remote state into the data source model.
func (d *ListFmStoredClusterConfigBackupSnapshotDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ListFmStoredClusterConfigBackupSnapshotDataSourceModel
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
func (d *ListFmStoredClusterConfigBackupSnapshotDataSource) readRemote(ctx context.Context, config *ListFmStoredClusterConfigBackupSnapshotDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/clusterConfig/backup/repo/{clusterId}/{backupId}"
	reqPath = strings.ReplaceAll(reqPath, "{clusterId}", url.PathEscape(config.ClusterId.ValueString()))
	reqPath = strings.ReplaceAll(reqPath, "{backupId}", url.PathEscape(config.BackupId.ValueString()))
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_list_fm_stored_cluster_config_backup_snapshot", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_list_fm_stored_cluster_config_backup_snapshot", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_list_fm_stored_cluster_config_backup_snapshot", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_list_fm_stored_cluster_config_backup_snapshot", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_list_fm_stored_cluster_config_backup_snapshot", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_list_fm_stored_cluster_config_backup_snapshot", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_list_fm_stored_cluster_config_backup_snapshot", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_list_fm_stored_cluster_config_backup_snapshot", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_list_fm_stored_cluster_config_backup_snapshot", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_list_fm_stored_cluster_config_backup_snapshot", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_list_fm_stored_cluster_config_backup_snapshot", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_list_fm_stored_cluster_config_backup_snapshot", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_list_fm_stored_cluster_config_backup_snapshot", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_list_fm_stored_cluster_config_backup_snapshot", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *ListFmStoredClusterConfigBackupSnapshotDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
