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
	_ datasource.DataSource              = (*GetRfsSyncDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetRfsSyncDataSource)(nil)
)

// GetRfsSyncDataSource is the generated Terraform data source implementation.
type GetRfsSyncDataSource struct {
	client *client.Client
}

// GetRfsSyncDataSourceModel describes the data source state shape.
type GetRfsSyncDataSourceModel struct {
	Alias          types.String `tfsdk:"alias"`
	ClusterId      types.String `tfsdk:"cluster_id" json:"clusterId"`
	LastSync       types.String `tfsdk:"last_sync" json:"lastSync"`
	LastSyncMethod types.String `tfsdk:"last_sync_method" json:"lastSyncMethod"`
	NextSync       types.String `tfsdk:"next_sync" json:"nextSync"`
	RfsAddress     types.String `tfsdk:"rfs_address" json:"rfsAddress"`
	RfsSyncInfo    types.List   `tfsdk:"rfs_sync_info" json:"rfsSyncInfo"`
	SyncPeriod     types.String `tfsdk:"sync_period" json:"syncPeriod"`
}

// NewGetRfsSyncDataSource returns a new instance of the generated data source.
func NewGetRfsSyncDataSource() datasource.DataSource {
	return &GetRfsSyncDataSource{}
}

// Metadata returns the data source type name.
func (d *GetRfsSyncDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_rfs_sync"
}

// Schema returns the data source schema.
func (d *GetRfsSyncDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get HSM Rfs Sync", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "alias of target HSM Group", Required: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "id of the defining cluster", Required: true}, "last_sync": schema.StringAttribute{MarkdownDescription: "Time that the last sync happened", Computed: true}, "last_sync_method": schema.StringAttribute{MarkdownDescription: "How the the last sync happened,either auto or manuel", Computed: true}, "next_sync": schema.StringAttribute{MarkdownDescription: "Time that next sync will happen", Computed: true}, "rfs_address": schema.StringAttribute{MarkdownDescription: "IPv4 address of the RFS server", Computed: true}, "rfs_sync_info": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"key_name": schema.StringAttribute{MarkdownDescription: "Key Name", Computed: true}, "key_token": schema.StringAttribute{MarkdownDescription: "Key Token", Computed: true}}}}, "sync_period": schema.StringAttribute{MarkdownDescription: "Period in hours of when to sync, 0 = no automatic sync", Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetRfsSyncDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetRfsSyncDataSourceModel
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
func (d *GetRfsSyncDataSource) readRemote(ctx context.Context, config *GetRfsSyncDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/apps/hsmGroup/{alias}/rfsSync"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(config.Alias.ValueString()))
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_rfs_sync", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_rfs_sync", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_rfs_sync", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_rfs_sync", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_rfs_sync", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_rfs_sync", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_rfs_sync", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_rfs_sync", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_rfs_sync", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_rfs_sync", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_rfs_sync", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_rfs_sync", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_rfs_sync", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_rfs_sync", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetRfsSyncDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
