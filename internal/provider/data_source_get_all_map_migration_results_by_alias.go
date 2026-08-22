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
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*GetAllMapMigrationResultsByAliasDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetAllMapMigrationResultsByAliasDataSource)(nil)
)

// GetAllMapMigrationResultsByAliasDataSource is the generated Terraform data source implementation.
type GetAllMapMigrationResultsByAliasDataSource struct {
	client *client.Client
}

// GetAllMapMigrationResultsByAliasDataSourceModel describes the data source state shape.
type GetAllMapMigrationResultsByAliasDataSourceModel struct {
	Alias                    types.String  `tfsdk:"alias"`
	AliasMismatched          types.Map     `tfsdk:"alias_mismatched" json:"aliasMismatched"`
	Aliases                  types.List    `tfsdk:"aliases"`
	AvgTimePerMap            types.Float64 `tfsdk:"avg_time_per_map" json:"avgTimePerMap"`
	ClusterId                types.String  `tfsdk:"cluster_id" json:"clusterId"`
	ClusterIds               types.List    `tfsdk:"cluster_ids" json:"clusterIds"`
	ClustersWithNoMaps       types.List    `tfsdk:"clusters_with_no_maps" json:"clustersWithNoMaps"`
	DryRun                   types.Bool    `tfsdk:"dry_run" json:"dryRun"`
	DuplicateMaps            types.List    `tfsdk:"duplicate_maps" json:"duplicateMaps"`
	EndTime                  types.Int64   `tfsdk:"end_time" json:"endTime"`
	FailedClusterAndReasons  types.Map     `tfsdk:"failed_cluster_and_reasons" json:"failedClusterAndReasons"`
	FailedMapsAndReasons     types.Map     `tfsdk:"failed_maps_and_reasons" json:"failedMapsAndReasons"`
	FailureReason            types.String  `tfsdk:"failure_reason" json:"failureReason"`
	IsAllMaps                types.Bool    `tfsdk:"is_all_maps" json:"isAllMaps"`
	MapType                  types.String  `tfsdk:"map_type" json:"mapType"`
	MigratedMaps             types.List    `tfsdk:"migrated_maps" json:"migratedMaps"`
	MigrationAlias           types.String  `tfsdk:"migration_alias" json:"migrationAlias"`
	SkippedMigration         types.Map     `tfsdk:"skipped_migration" json:"skippedMigration"`
	StartTime                types.Int64   `tfsdk:"start_time" json:"startTime"`
	Status                   types.String  `tfsdk:"status"`
	TotalDuplicates          types.Int64   `tfsdk:"total_duplicates" json:"totalDuplicates"`
	TotalFailed              types.Int64   `tfsdk:"total_failed" json:"totalFailed"`
	TotalMaps                types.Int64   `tfsdk:"total_maps" json:"totalMaps"`
	TotalPortAliasMismatches types.Int64   `tfsdk:"total_port_alias_mismatches" json:"totalPortAliasMismatches"`
	TotalSkipped             types.Int64   `tfsdk:"total_skipped" json:"totalSkipped"`
	TotalSuccess             types.Int64   `tfsdk:"total_success" json:"totalSuccess"`
	TotalTimeToComplete      types.Int64   `tfsdk:"total_time_to_complete" json:"totalTimeToComplete"`
	TriggeredBy              types.String  `tfsdk:"triggered_by" json:"triggeredBy"`
}

// NewGetAllMapMigrationResultsByAliasDataSource returns a new instance of the generated data source.
func NewGetAllMapMigrationResultsByAliasDataSource() datasource.DataSource {
	return &GetAllMapMigrationResultsByAliasDataSource{}
}

// Metadata returns the data source type name.
func (d *GetAllMapMigrationResultsByAliasDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_all_map_migration_results_by_alias"
}

// Schema returns the data source schema.
func (d *GetAllMapMigrationResultsByAliasDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load map migration result by alias", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Migration alias", Required: true}, "alias_mismatched": schema.MapAttribute{Computed: true, ElementType: types.StringType}, "aliases": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "avg_time_per_map": schema.Float64Attribute{Computed: true}, "cluster_id": schema.StringAttribute{Computed: true}, "cluster_ids": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "clusters_with_no_maps": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "dry_run": schema.BoolAttribute{Computed: true}, "duplicate_maps": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "end_time": schema.Int64Attribute{Computed: true}, "failed_cluster_and_reasons": schema.MapAttribute{Computed: true, ElementType: types.StringType}, "failed_maps_and_reasons": schema.MapAttribute{Computed: true, ElementType: types.StringType}, "failure_reason": schema.StringAttribute{Computed: true}, "is_all_maps": schema.BoolAttribute{Computed: true}, "map_type": schema.StringAttribute{Computed: true}, "migrated_maps": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "migration_alias": schema.StringAttribute{Computed: true}, "skipped_migration": schema.MapAttribute{Computed: true, ElementType: types.StringType}, "start_time": schema.Int64Attribute{Computed: true}, "status": schema.StringAttribute{Computed: true}, "total_duplicates": schema.Int64Attribute{Computed: true}, "total_failed": schema.Int64Attribute{Computed: true}, "total_maps": schema.Int64Attribute{Computed: true}, "total_port_alias_mismatches": schema.Int64Attribute{Computed: true}, "total_skipped": schema.Int64Attribute{Computed: true}, "total_success": schema.Int64Attribute{Computed: true}, "total_time_to_complete": schema.Int64Attribute{Computed: true}, "triggered_by": schema.StringAttribute{Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetAllMapMigrationResultsByAliasDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetAllMapMigrationResultsByAliasDataSourceModel
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
func (d *GetAllMapMigrationResultsByAliasDataSource) readRemote(ctx context.Context, config *GetAllMapMigrationResultsByAliasDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/traffic-flows/migrationResult/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(config.Alias.ValueString()))
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_map_migration_results_by_alias", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_map_migration_results_by_alias", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_map_migration_results_by_alias", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_map_migration_results_by_alias", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_map_migration_results_by_alias", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_map_migration_results_by_alias", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_map_migration_results_by_alias", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_map_migration_results_by_alias", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_map_migration_results_by_alias", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_map_migration_results_by_alias", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_all_map_migration_results_by_alias", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_map_migration_results_by_alias", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_map_migration_results_by_alias", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["mapMigrationResult"]; ok {
		if m, ok := v.(map[string]any); ok {
			data = m
		}
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_map_migration_results_by_alias", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetAllMapMigrationResultsByAliasDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
