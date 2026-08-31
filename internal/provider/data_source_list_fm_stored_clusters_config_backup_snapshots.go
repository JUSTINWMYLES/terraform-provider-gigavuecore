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
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*ListFmStoredClustersConfigBackupSnapshotsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*ListFmStoredClustersConfigBackupSnapshotsDataSource)(nil)
)

// ListFmStoredClustersConfigBackupSnapshotsDataSource is the generated Terraform data source implementation.
type ListFmStoredClustersConfigBackupSnapshotsDataSource struct {
	client *client.Client
}

// ListFmStoredClustersConfigBackupSnapshotsDataSourceModel describes the data source state shape.
type ListFmStoredClustersConfigBackupSnapshotsDataSourceModel struct {
	ClusterId types.String `tfsdk:"cluster_id" json:"clusterId"`
	Items     types.List   `tfsdk:"items"`
}

// NewListFmStoredClustersConfigBackupSnapshotsDataSource returns a new instance of the generated data source.
func NewListFmStoredClustersConfigBackupSnapshotsDataSource() datasource.DataSource {
	return &ListFmStoredClustersConfigBackupSnapshotsDataSource{}
}

// Metadata returns the data source type name.
func (d *ListFmStoredClustersConfigBackupSnapshotsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_list_fm_stored_clusters_config_backup_snapshots"
}

// Schema returns the data source schema.
func (d *ListFmStoredClustersConfigBackupSnapshotsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Lists config backup snapshots for multiple clusters", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "if provided, only backups associated with this cluster will be returned", Optional: true}, "items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Id of the cluster", Computed: true}, "cluster_name": schema.StringAttribute{MarkdownDescription: "Name of the cluster", Computed: true}, "entries": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "User-entered alias for this snapshot", Computed: true}, "backup_id": schema.StringAttribute{MarkdownDescription: "Unique identifier of the cluster config snapshot for a given cluster", Computed: true}, "backup_time": schema.StringAttribute{MarkdownDescription: "Configuration file backup time", Computed: true}, "cluster_name": schema.StringAttribute{MarkdownDescription: "Name of the cluster", Computed: true}, "cluster_vip": schema.StringAttribute{MarkdownDescription: "Virtual Ip of the cluster", Computed: true}, "comment": schema.StringAttribute{MarkdownDescription: "User-entered comments for this snapshot", Computed: true}, "device_global_id": schema.StringAttribute{MarkdownDescription: "Device Global Id of the hosting node", Computed: true}, "device_host_name": schema.StringAttribute{MarkdownDescription: "Device HostName of the hosting node", Computed: true}, "device_ip": schema.StringAttribute{MarkdownDescription: "Device Ip of the hosting node", Computed: true}, "do_not_purge": schema.BoolAttribute{MarkdownDescription: "Indicates whether this snapshot should not be auto-aged/purged", Computed: true}, "file_meta_data": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"device_global_id": schema.StringAttribute{MarkdownDescription: "Device Global Id of the hosting node", Computed: true}, "device_host_name": schema.StringAttribute{MarkdownDescription: "Device HostName of the hosting node", Computed: true}, "device_ip": schema.StringAttribute{MarkdownDescription: "Device Ip of the hosting node", Computed: true}, "device_model": schema.StringAttribute{MarkdownDescription: "Gigamon physical device models", Computed: true}, "device_version": schema.StringAttribute{MarkdownDescription: "Device's current Software version", Computed: true}, "do_not_purge": schema.BoolAttribute{MarkdownDescription: "Indicates whether this snapshot should not be auto-aged/purged", Computed: true}, "file_name": schema.StringAttribute{MarkdownDescription: "Config file name", Computed: true}, "member_info": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"box_id": schema.StringAttribute{MarkdownDescription: "Cluster BoxId of the node", Computed: true}, "host_name": schema.StringAttribute{MarkdownDescription: "Device HostName of the hosting node", Computed: true}, "role": schema.StringAttribute{MarkdownDescription: "Role of the node. (deprecated: use roleAlias)", Computed: true}, "role_alias": schema.StringAttribute{MarkdownDescription: "Role of the node", Computed: true}}}}}}, "format": schema.StringAttribute{MarkdownDescription: "format of the Backup file", Computed: true}, "restore_logs": schema.ListAttribute{MarkdownDescription: "restoreLogs associated with this file", Computed: true, ElementType: types.StringType}, "sw_version": schema.StringAttribute{MarkdownDescription: "Software version of the device this snapshot was taken from", Computed: true}, "tags": schema.ListAttribute{MarkdownDescription: "Tags associated with this file", Computed: true, ElementType: types.StringType}}}}}}}}}
}

// Read fetches remote state into the data source model.
func (d *ListFmStoredClustersConfigBackupSnapshotsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ListFmStoredClustersConfigBackupSnapshotsDataSourceModel
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
func (d *ListFmStoredClustersConfigBackupSnapshotsDataSource) readListRemote(ctx context.Context, config *ListFmStoredClustersConfigBackupSnapshotsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/clusterConfig/backup/repo"
	params := url.Values{}
	if !config.ClusterId.IsNull() {
		params.Set("clusterId", config.ClusterId.ValueString())
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
		resp.Diagnostics.AddError("Error reading gigavuecore_list_fm_stored_clusters_config_backup_snapshots", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_list_fm_stored_clusters_config_backup_snapshots", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["clustersConfigBackups"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_list_fm_stored_clusters_config_backup_snapshots", fmt.Sprintf("Could not decode list page: missing %q array", "clustersConfigBackups"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_list_fm_stored_clusters_config_backup_snapshots", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *ListFmStoredClustersConfigBackupSnapshotsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
