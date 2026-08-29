package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
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
	ClusterId      types.String `tfsdk:"cluster_id" json:"clusterId"`
	IsAllMaps      types.Bool   `tfsdk:"is_all_maps" json:"isAllMaps"`
	IsDryRun       types.Bool   `tfsdk:"is_dry_run" json:"isDryRun"`
	Items          types.List   `tfsdk:"items"`
	MapAlias       types.String `tfsdk:"map_alias" json:"mapAlias"`
	MapType        types.String `tfsdk:"map_type" json:"mapType"`
	MigrationAlias types.String `tfsdk:"migration_alias" json:"migrationAlias"`
	Page           types.String `tfsdk:"page"`
	Sort           types.String `tfsdk:"sort"`
	Status         types.String `tfsdk:"status"`
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
	resp.Schema = schema.Schema{MarkdownDescription: "Query map migration results with filters", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Optional: true}, "is_all_maps": schema.BoolAttribute{MarkdownDescription: "Include all maps", Optional: true}, "is_dry_run": schema.BoolAttribute{MarkdownDescription: "Filter by dry run status", Optional: true}, "items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias_mismatched": schema.MapAttribute{Computed: true, ElementType: types.StringType}, "aliases": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "avg_time_per_map": schema.Float64Attribute{Computed: true}, "cluster_id": schema.StringAttribute{Computed: true}, "cluster_ids": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "clusters_with_no_maps": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "dry_run": schema.BoolAttribute{Computed: true}, "duplicate_maps": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "end_time": schema.Int64Attribute{Computed: true}, "failed_cluster_and_reasons": schema.MapAttribute{Computed: true, ElementType: types.StringType}, "failed_maps_and_reasons": schema.MapAttribute{Computed: true, ElementType: types.StringType}, "failure_reason": schema.StringAttribute{Computed: true}, "is_all_maps": schema.BoolAttribute{Computed: true}, "map_type": schema.StringAttribute{Computed: true}, "migrated_maps": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "migration_alias": schema.StringAttribute{Computed: true}, "skipped_migration": schema.MapAttribute{Computed: true, ElementType: types.StringType}, "start_time": schema.Int64Attribute{Computed: true}, "status": schema.StringAttribute{Computed: true}, "total_duplicates": schema.Int64Attribute{Computed: true}, "total_failed": schema.Int64Attribute{Computed: true}, "total_maps": schema.Int64Attribute{Computed: true}, "total_port_alias_mismatches": schema.Int64Attribute{Computed: true}, "total_skipped": schema.Int64Attribute{Computed: true}, "total_success": schema.Int64Attribute{Computed: true}, "total_time_to_complete": schema.Int64Attribute{Computed: true}, "triggered_by": schema.StringAttribute{Computed: true}}}}, "map_alias": schema.StringAttribute{MarkdownDescription: "Map alias", Optional: true}, "map_type": schema.StringAttribute{MarkdownDescription: "Map type", Optional: true}, "migration_alias": schema.StringAttribute{MarkdownDescription: "Migration alias", Optional: true}, "page": schema.StringAttribute{MarkdownDescription: "Pagination request string (e.g., \"(1:30)\")", Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "Sort request string (e.g., \"(startTime:DESC)\")", Optional: true}, "status": schema.StringAttribute{MarkdownDescription: "Migration status", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetMapMigrationResultsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetMapMigrationResultsDataSourceModel
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
func (d *GetMapMigrationResultsDataSource) readListRemote(ctx context.Context, config *GetMapMigrationResultsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/traffic-flows/migrationResult"
	params := url.Values{}
	if !config.MigrationAlias.IsNull() {
		params.Set("migrationAlias", config.MigrationAlias.ValueString())
	}
	if !config.MapType.IsNull() {
		params.Set("mapType", config.MapType.ValueString())
	}
	if !config.Status.IsNull() {
		params.Set("status", config.Status.ValueString())
	}
	if !config.ClusterId.IsNull() {
		params.Set("clusterId", config.ClusterId.ValueString())
	}
	if !config.IsDryRun.IsNull() {
		params.Set("isDryRun", strconv.FormatBool(config.IsDryRun.ValueBool()))
	}
	if !config.IsAllMaps.IsNull() {
		params.Set("isAllMaps", strconv.FormatBool(config.IsAllMaps.ValueBool()))
	}
	if !config.MapAlias.IsNull() {
		params.Set("mapAlias", config.MapAlias.ValueString())
	}
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_map_migration_results", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_map_migration_results", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["mapMigrationResults"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_map_migration_results", fmt.Sprintf("Could not decode list page: missing %q array", "mapMigrationResults"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
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
