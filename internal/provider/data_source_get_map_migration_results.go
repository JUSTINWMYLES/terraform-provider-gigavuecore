package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)
import (
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*GetMapMigrationResultsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetMapMigrationResultsDataSource)(nil)
)

// GetMapMigrationResultsDataSource is the generated Terraform data source implementation.
type GetMapMigrationResultsDataSource struct {
	client *client.Client
}

// GetMapMigrationResultsDataSourceModel describes the data source state shape.
type GetMapMigrationResultsDataSourceModel struct {
	ClusterId           types.String `tfsdk:"cluster_id" json:"clusterId"`
	Context             types.Object `tfsdk:"context"`
	IsAllMaps           types.Bool   `tfsdk:"is_all_maps" json:"isAllMaps"`
	IsDryRun            types.Bool   `tfsdk:"is_dry_run" json:"isDryRun"`
	MapAlias            types.String `tfsdk:"map_alias" json:"mapAlias"`
	MapMigrationResults types.List   `tfsdk:"map_migration_results" json:"mapMigrationResults"`
	MapType             types.String `tfsdk:"map_type" json:"mapType"`
	MigrationAlias      types.String `tfsdk:"migration_alias" json:"migrationAlias"`
	Page                types.String `tfsdk:"page"`
	Sort                types.String `tfsdk:"sort"`
	Status              types.String `tfsdk:"status"`
}

// NewGetMapMigrationResultsDataSource returns a new instance of the generated data source.
func NewGetMapMigrationResultsDataSource() datasource.DataSource {
	return &GetMapMigrationResultsDataSource{}
}

// Metadata returns the data source type name.
func (d *GetMapMigrationResultsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_map_migration_results"
}

// Schema returns the data source schema.
func (d *GetMapMigrationResultsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Query map migration results with filters", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Optional: true}, "context": schema.SingleNestedAttribute{MarkdownDescription: "Gigamon query result context", Computed: true, Attributes: map[string]schema.Attribute{"page_no": schema.Int64Attribute{MarkdownDescription: "page number of the returned result set", Computed: true}, "page_size": schema.Int64Attribute{MarkdownDescription: "page size of the returned result set", Computed: true}, "sort": schema.ListAttribute{MarkdownDescription: "sorting info of the returned result set. list of fields in the array indicate sorting order", Computed: true, ElementType: types.StringType}, "total_items": schema.Int64Attribute{MarkdownDescription: "total number of items in the queried entity type", Computed: true}}}, "is_all_maps": schema.BoolAttribute{MarkdownDescription: "Include all maps", Optional: true}, "is_dry_run": schema.BoolAttribute{MarkdownDescription: "Filter by dry run status", Optional: true}, "map_alias": schema.StringAttribute{MarkdownDescription: "Map alias", Optional: true}, "map_migration_results": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias_mismatched": schema.MapAttribute{Computed: true, ElementType: types.StringType}, "aliases": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "avg_time_per_map": schema.Float64Attribute{Computed: true}, "cluster_id": schema.StringAttribute{Computed: true}, "cluster_ids": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "clusters_with_no_maps": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "dry_run": schema.BoolAttribute{Computed: true}, "duplicate_maps": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "end_time": schema.Int64Attribute{Computed: true}, "failed_cluster_and_reasons": schema.MapAttribute{Computed: true, ElementType: types.StringType}, "failed_maps_and_reasons": schema.MapAttribute{Computed: true, ElementType: types.StringType}, "failure_reason": schema.StringAttribute{Computed: true}, "is_all_maps": schema.BoolAttribute{Computed: true}, "map_type": schema.StringAttribute{Computed: true}, "migrated_maps": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "migration_alias": schema.StringAttribute{Computed: true}, "skipped_migration": schema.MapAttribute{Computed: true, ElementType: types.StringType}, "start_time": schema.Int64Attribute{Computed: true}, "status": schema.StringAttribute{Computed: true}, "total_duplicates": schema.Int64Attribute{Computed: true}, "total_failed": schema.Int64Attribute{Computed: true}, "total_maps": schema.Int64Attribute{Computed: true}, "total_port_alias_mismatches": schema.Int64Attribute{Computed: true}, "total_skipped": schema.Int64Attribute{Computed: true}, "total_success": schema.Int64Attribute{Computed: true}, "total_time_to_complete": schema.Int64Attribute{Computed: true}, "triggered_by": schema.StringAttribute{Computed: true}}}}, "map_type": schema.StringAttribute{MarkdownDescription: "Map type", Optional: true}, "migration_alias": schema.StringAttribute{MarkdownDescription: "Migration alias", Optional: true}, "page": schema.StringAttribute{MarkdownDescription: "Pagination request string (e.g., \"(1:30)\")", Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "Sort request string (e.g., \"(startTime:DESC)\")", Optional: true}, "status": schema.StringAttribute{MarkdownDescription: "Migration status", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetMapMigrationResultsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetMapMigrationResultsDataSourceModel
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
func (d *GetMapMigrationResultsDataSource) readRemote(ctx context.Context, config *GetMapMigrationResultsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/traffic-flows/migrationResult"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_map_migration_results", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !config.MigrationAlias.IsNull() {
		query.Set("migrationAlias", config.MigrationAlias.ValueString())
	}
	if !config.MapType.IsNull() {
		query.Set("mapType", config.MapType.ValueString())
	}
	if !config.Status.IsNull() {
		query.Set("status", config.Status.ValueString())
	}
	if !config.ClusterId.IsNull() {
		query.Set("clusterId", config.ClusterId.ValueString())
	}
	if !config.IsDryRun.IsNull() {
		query.Set("isDryRun", strconv.FormatBool(config.IsDryRun.ValueBool()))
	}
	if !config.IsAllMaps.IsNull() {
		query.Set("isAllMaps", strconv.FormatBool(config.IsAllMaps.ValueBool()))
	}
	if !config.MapAlias.IsNull() {
		query.Set("mapAlias", config.MapAlias.ValueString())
	}
	if !config.Page.IsNull() {
		query.Set("page", config.Page.ValueString())
	}
	if !config.Sort.IsNull() {
		query.Set("sort", config.Sort.ValueString())
	}
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_map_migration_results", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_map_migration_results", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_map_migration_results", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_map_migration_results", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_map_migration_results", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_map_migration_results", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_map_migration_results", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_map_migration_results", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_map_migration_results", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_map_migration_results", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_map_migration_results", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_map_migration_results", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_map_migration_results", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetMapMigrationResultsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
