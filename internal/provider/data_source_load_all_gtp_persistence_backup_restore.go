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
	_ datasource.DataSource              = (*LoadAllGtpPersistenceBackupRestoreDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadAllGtpPersistenceBackupRestoreDataSource)(nil)
)

// LoadAllGtpPersistenceBackupRestoreDataSource is the generated Terraform data source implementation.
type LoadAllGtpPersistenceBackupRestoreDataSource struct {
	client *client.Client
}

// LoadAllGtpPersistenceBackupRestoreDataSourceModel describes the data source state shape.
type LoadAllGtpPersistenceBackupRestoreDataSourceModel struct {
	ClusterId                       types.String `tfsdk:"cluster_id" json:"clusterId"`
	Context                         types.Object `tfsdk:"context"`
	GtpPersistantBackupRestoreInfos types.List   `tfsdk:"gtp_persistant_backup_restore_infos" json:"gtpPersistantBackupRestoreInfos"`
	Page                            types.String `tfsdk:"page"`
	Sort                            types.String `tfsdk:"sort"`
}

// NewLoadAllGtpPersistenceBackupRestoreDataSource returns a new instance of the generated data source.
func NewLoadAllGtpPersistenceBackupRestoreDataSource() datasource.DataSource {
	return &LoadAllGtpPersistenceBackupRestoreDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadAllGtpPersistenceBackupRestoreDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_all_gtp_persistence_backup_restore"
}

// Schema returns the data source schema.
func (d *LoadAllGtpPersistenceBackupRestoreDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load all Gtp Persistence backup restore information", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "context": schema.SingleNestedAttribute{MarkdownDescription: "Gigamon query result context", Computed: true, Attributes: map[string]schema.Attribute{"page_no": schema.Int64Attribute{MarkdownDescription: "page number of the returned result set", Computed: true}, "page_size": schema.Int64Attribute{MarkdownDescription: "page size of the returned result set", Computed: true}, "sort": schema.ListAttribute{MarkdownDescription: "sorting info of the returned result set. list of fields in the array indicate sorting order", Computed: true, ElementType: types.StringType}, "total_items": schema.Int64Attribute{MarkdownDescription: "total number of items in the queried entity type", Computed: true}}}, "gtp_persistant_backup_restore_infos": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Gs Group alias", Computed: true}, "backup_info": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"config_status": schema.StringAttribute{MarkdownDescription: "The status of a backup", Computed: true}, "control_tunnels": schema.Int64Attribute{MarkdownDescription: "The number of control tunnels backed up", Computed: true}, "failed": schema.Int64Attribute{MarkdownDescription: "The number of failed backups", Computed: true}, "filename": schema.StringAttribute{MarkdownDescription: "The internal name of the backup file", Computed: true}, "in_progress": schema.BoolAttribute{Computed: true}, "last_fail_time": schema.StringAttribute{MarkdownDescription: "The timestamp of the last failed backup. In ISO-8601 date format 'yyyy-MM-dd'T'HH:mm:ssZ'", Computed: true}, "last_sucessful_time": schema.StringAttribute{MarkdownDescription: "The timestamp of the last successful backup. In ISO-8601 date format 'yyyy-MM-dd'T'HH:mm:ssZ'", Computed: true}, "sessions": schema.Int64Attribute{MarkdownDescription: "The number of sessions tunnels backed up", Computed: true}, "success": schema.Int64Attribute{MarkdownDescription: "The number of successful backups", Computed: true}, "user_tunnels": schema.Int64Attribute{MarkdownDescription: "The number of user tunnels backed up", Computed: true}}}, "cluster_id": schema.StringAttribute{MarkdownDescription: "id of the defining cluster", Computed: true}, "restore_info": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"restore_time": schema.StringAttribute{MarkdownDescription: "The timestamp of the last restore. In ISO-8601 date format 'yyyy-MM-dd'T'HH:mm:ssZ'", Computed: true}, "sessions": schema.Int64Attribute{MarkdownDescription: "The number of sessions restored", Computed: true}, "tunnels": schema.Int64Attribute{MarkdownDescription: "The number of tunnels restored", Computed: true}}}}}}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned", Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadAllGtpPersistenceBackupRestoreDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadAllGtpPersistenceBackupRestoreDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	d.readRemote(ctx, &config, resp)
	if resp.Diagnostics.HasError() {
		return
	}
	if config.Page.IsNull() {
		resp.Diagnostics.AddWarning("Single-page result", fmt.Sprintf("This data source reads a single page of a paginated API endpoint and does not aggregate results across pages. The \"page\" argument is unset, so the default page is returned; set it to retrieve a different page."))
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}

// readRemote performs the read HTTP exchange and decodes the response into config. Extracted from Read so the request/response logic is unit-testable without a tfsdk.Config.
func (d *LoadAllGtpPersistenceBackupRestoreDataSource) readRemote(ctx context.Context, config *LoadAllGtpPersistenceBackupRestoreDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/gsGroups/gtpPersistence"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_gtp_persistence_backup_restore", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	if !config.Page.IsNull() {
		query.Set("page", config.Page.ValueString())
	}
	if !config.Sort.IsNull() {
		query.Set("sort", config.Sort.ValueString())
	}
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_gtp_persistence_backup_restore", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_gtp_persistence_backup_restore", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_gtp_persistence_backup_restore", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_gtp_persistence_backup_restore", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_gtp_persistence_backup_restore", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_gtp_persistence_backup_restore", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_gtp_persistence_backup_restore", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_gtp_persistence_backup_restore", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_gtp_persistence_backup_restore", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_load_all_gtp_persistence_backup_restore", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_gtp_persistence_backup_restore", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_gtp_persistence_backup_restore", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_gtp_persistence_backup_restore", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadAllGtpPersistenceBackupRestoreDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
