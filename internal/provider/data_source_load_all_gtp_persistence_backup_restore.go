package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)
import (
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
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
	ClusterId types.String `tfsdk:"cluster_id" json:"clusterId"`
	Items     types.List   `tfsdk:"items"`
	Page      types.String `tfsdk:"page"`
	Sort      types.String `tfsdk:"sort"`
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
	resp.Schema = schema.Schema{MarkdownDescription: "Load all Gtp Persistence backup restore information", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Gs Group alias", Computed: true}, "backup_info": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"config_status": schema.StringAttribute{MarkdownDescription: "The status of a backup", Computed: true}, "control_tunnels": schema.Int64Attribute{MarkdownDescription: "The number of control tunnels backed up", Computed: true}, "failed": schema.Int64Attribute{MarkdownDescription: "The number of failed backups", Computed: true}, "filename": schema.StringAttribute{MarkdownDescription: "The internal name of the backup file", Computed: true}, "in_progress": schema.BoolAttribute{Computed: true}, "last_fail_time": schema.StringAttribute{MarkdownDescription: "The timestamp of the last failed backup. In ISO-8601 date format 'yyyy-MM-dd'T'HH:mm:ssZ'", Computed: true}, "last_sucessful_time": schema.StringAttribute{MarkdownDescription: "The timestamp of the last successful backup. In ISO-8601 date format 'yyyy-MM-dd'T'HH:mm:ssZ'", Computed: true}, "sessions": schema.Int64Attribute{MarkdownDescription: "The number of sessions tunnels backed up", Computed: true}, "success": schema.Int64Attribute{MarkdownDescription: "The number of successful backups", Computed: true}, "user_tunnels": schema.Int64Attribute{MarkdownDescription: "The number of user tunnels backed up", Computed: true}}}, "cluster_id": schema.StringAttribute{MarkdownDescription: "id of the defining cluster", Computed: true}, "restore_info": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"restore_time": schema.StringAttribute{MarkdownDescription: "The timestamp of the last restore. In ISO-8601 date format 'yyyy-MM-dd'T'HH:mm:ssZ'", Computed: true}, "sessions": schema.Int64Attribute{MarkdownDescription: "The number of sessions restored", Computed: true}, "tunnels": schema.Int64Attribute{MarkdownDescription: "The number of tunnels restored", Computed: true}}}}}}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned", Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadAllGtpPersistenceBackupRestoreDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadAllGtpPersistenceBackupRestoreDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	d.readListRemote(ctx, &config, resp)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}

// readListRemote performs the paginated read HTTP exchange and decodes the response array into config. Extracted from Read so the request/response logic is unit-testable without a tfsdk.Config.
func (d *LoadAllGtpPersistenceBackupRestoreDataSource) readListRemote(ctx context.Context, config *LoadAllGtpPersistenceBackupRestoreDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/gsGroups/gtpPersistence"
	params := url.Values{}
	params.Set("clusterId", config.ClusterId.ValueString())
	if !config.Page.IsNull() {
		params.Set("page", config.Page.ValueString())
	}
	if !config.Sort.IsNull() {
		params.Set("sort", config.Sort.ValueString())
	}
	var nextURL string
	fetch := func(ctx context.Context, p url.Values) (*http.Response, error) {
		httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
		if err != nil {
			return nil, err
		}
		if nextURL != "" {
			parsed, perr := url.Parse(nextURL)
			if perr != nil {
				return nil, perr
			}
			httpReq.URL = parsed
		} else {
			httpReq.URL.RawQuery = p.Encode()
		}
		return d.client.Do(httpReq)
	}
	pages, err := client.ListAllPages(ctx, params, fetch, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_gtp_persistence_backup_restore", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_gtp_persistence_backup_restore", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["gtpPersistantBackupRestoreInfos"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_gtp_persistence_backup_restore", fmt.Sprintf("Could not decode list page: missing %q array", "gtpPersistantBackupRestoreInfos"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
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
