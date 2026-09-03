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
	_ datasource.DataSource              = (*GetGtpPersistenceBackupRestoreByAliasDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetGtpPersistenceBackupRestoreByAliasDataSource)(nil)
)

// GetGtpPersistenceBackupRestoreByAliasDataSource is the generated Terraform data source implementation.
type GetGtpPersistenceBackupRestoreByAliasDataSource struct {
	client *client.Client
}

// GetGtpPersistenceBackupRestoreByAliasDataSourceModel describes the data source state shape.
type GetGtpPersistenceBackupRestoreByAliasDataSourceModel struct {
	Alias       types.String `tfsdk:"alias"`
	BackupInfo  types.Object `tfsdk:"backup_info" json:"backupInfo"`
	ClusterId   types.String `tfsdk:"cluster_id" json:"clusterId"`
	RestoreInfo types.Object `tfsdk:"restore_info" json:"restoreInfo"`
}

// NewGetGtpPersistenceBackupRestoreByAliasDataSource returns a new instance of the generated data source.
func NewGetGtpPersistenceBackupRestoreByAliasDataSource() datasource.DataSource {
	return &GetGtpPersistenceBackupRestoreByAliasDataSource{}
}

// Metadata returns the data source type name.
func (d *GetGtpPersistenceBackupRestoreByAliasDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_gtp_persistence_backup_restore_by_alias"
}

// Schema returns the data source schema.
func (d *GetGtpPersistenceBackupRestoreByAliasDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Find Gtp Persistence backup restore information by alias", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Gs Group alias", Required: true}, "backup_info": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"config_status": schema.StringAttribute{MarkdownDescription: "The status of a backup", Computed: true}, "control_tunnels": schema.Int64Attribute{MarkdownDescription: "The number of control tunnels backed up", Computed: true}, "failed": schema.Int64Attribute{MarkdownDescription: "The number of failed backups", Computed: true}, "filename": schema.StringAttribute{MarkdownDescription: "The internal name of the backup file", Computed: true}, "in_progress": schema.BoolAttribute{Computed: true}, "last_fail_time": schema.StringAttribute{MarkdownDescription: "The timestamp of the last failed backup. In ISO-8601 date format 'yyyy-MM-dd'T'HH:mm:ssZ'", Computed: true}, "last_sucessful_time": schema.StringAttribute{MarkdownDescription: "The timestamp of the last successful backup. In ISO-8601 date format 'yyyy-MM-dd'T'HH:mm:ssZ'", Computed: true}, "sessions": schema.Int64Attribute{MarkdownDescription: "The number of sessions tunnels backed up", Computed: true}, "success": schema.Int64Attribute{MarkdownDescription: "The number of successful backups", Computed: true}, "user_tunnels": schema.Int64Attribute{MarkdownDescription: "The number of user tunnels backed up", Computed: true}}}, "cluster_id": schema.StringAttribute{MarkdownDescription: "id of the defining cluster", Required: true}, "restore_info": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"restore_time": schema.StringAttribute{MarkdownDescription: "The timestamp of the last restore. In ISO-8601 date format 'yyyy-MM-dd'T'HH:mm:ssZ'", Computed: true}, "sessions": schema.Int64Attribute{MarkdownDescription: "The number of sessions restored", Computed: true}, "tunnels": schema.Int64Attribute{MarkdownDescription: "The number of tunnels restored", Computed: true}}}}}
}

// Read fetches remote state into the data source model.
func (d *GetGtpPersistenceBackupRestoreByAliasDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetGtpPersistenceBackupRestoreByAliasDataSourceModel
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
func (d *GetGtpPersistenceBackupRestoreByAliasDataSource) readRemote(ctx context.Context, config *GetGtpPersistenceBackupRestoreByAliasDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/gsGroups/{alias}/gtpPersistence"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(config.Alias.ValueString()))
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_gtp_persistence_backup_restore_by_alias", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_gtp_persistence_backup_restore_by_alias", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_gtp_persistence_backup_restore_by_alias", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_gtp_persistence_backup_restore_by_alias", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_gtp_persistence_backup_restore_by_alias", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_gtp_persistence_backup_restore_by_alias", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_gtp_persistence_backup_restore_by_alias", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_gtp_persistence_backup_restore_by_alias", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_gtp_persistence_backup_restore_by_alias", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_gtp_persistence_backup_restore_by_alias", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_gtp_persistence_backup_restore_by_alias", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_gtp_persistence_backup_restore_by_alias", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_gtp_persistence_backup_restore_by_alias", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["gtpPersistantBackupRestoreInfo"]; ok {
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_gtp_persistence_backup_restore_by_alias", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetGtpPersistenceBackupRestoreByAliasDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
