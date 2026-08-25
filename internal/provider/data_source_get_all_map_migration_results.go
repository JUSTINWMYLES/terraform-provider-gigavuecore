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
	_ datasource.DataSource              = (*GetAllMapMigrationResultsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetAllMapMigrationResultsDataSource)(nil)
)

// GetAllMapMigrationResultsDataSource is the generated Terraform data source implementation.
type GetAllMapMigrationResultsDataSource struct {
	client *client.Client
}

// GetAllMapMigrationResultsDataSourceModel describes the data source state shape.
type GetAllMapMigrationResultsDataSourceModel struct {
	Context             types.Object `tfsdk:"context"`
	MapMigrationResults types.List   `tfsdk:"map_migration_results" json:"mapMigrationResults"`
}

// NewGetAllMapMigrationResultsDataSource returns a new instance of the generated data source.
func NewGetAllMapMigrationResultsDataSource() datasource.DataSource {
	return &GetAllMapMigrationResultsDataSource{}
}

// Metadata returns the data source type name.
func (d *GetAllMapMigrationResultsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_all_map_migration_results"
}

// Schema returns the data source schema.
func (d *GetAllMapMigrationResultsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load all map migration results", Attributes: map[string]schema.Attribute{"context": schema.SingleNestedAttribute{MarkdownDescription: "Gigamon query result context", Computed: true, Attributes: map[string]schema.Attribute{"page_no": schema.Int64Attribute{MarkdownDescription: "page number of the returned result set", Computed: true}, "page_size": schema.Int64Attribute{MarkdownDescription: "page size of the returned result set", Computed: true}, "sort": schema.ListAttribute{MarkdownDescription: "sorting info of the returned result set. list of fields in the array indicate sorting order", Computed: true, ElementType: types.StringType}, "total_items": schema.Int64Attribute{MarkdownDescription: "total number of items in the queried entity type", Computed: true}}}, "map_migration_results": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias_mismatched": schema.MapAttribute{Computed: true, ElementType: types.StringType}, "aliases": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "avg_time_per_map": schema.Float64Attribute{Computed: true}, "cluster_id": schema.StringAttribute{Computed: true}, "cluster_ids": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "clusters_with_no_maps": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "dry_run": schema.BoolAttribute{Computed: true}, "duplicate_maps": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "end_time": schema.Int64Attribute{Computed: true}, "failed_cluster_and_reasons": schema.MapAttribute{Computed: true, ElementType: types.StringType}, "failed_maps_and_reasons": schema.MapAttribute{Computed: true, ElementType: types.StringType}, "failure_reason": schema.StringAttribute{Computed: true}, "is_all_maps": schema.BoolAttribute{Computed: true}, "map_type": schema.StringAttribute{Computed: true}, "migrated_maps": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "migration_alias": schema.StringAttribute{Computed: true}, "skipped_migration": schema.MapAttribute{Computed: true, ElementType: types.StringType}, "start_time": schema.Int64Attribute{Computed: true}, "status": schema.StringAttribute{Computed: true}, "total_duplicates": schema.Int64Attribute{Computed: true}, "total_failed": schema.Int64Attribute{Computed: true}, "total_maps": schema.Int64Attribute{Computed: true}, "total_port_alias_mismatches": schema.Int64Attribute{Computed: true}, "total_skipped": schema.Int64Attribute{Computed: true}, "total_success": schema.Int64Attribute{Computed: true}, "total_time_to_complete": schema.Int64Attribute{Computed: true}, "triggered_by": schema.StringAttribute{Computed: true}}}}}}
}

// Read fetches remote state into the data source model.
func (d *GetAllMapMigrationResultsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetAllMapMigrationResultsDataSourceModel
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
func (d *GetAllMapMigrationResultsDataSource) readRemote(ctx context.Context, config *GetAllMapMigrationResultsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/traffic-flows/migrationResult/all"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_map_migration_results", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_map_migration_results", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_map_migration_results", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_map_migration_results", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_map_migration_results", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_map_migration_results", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_map_migration_results", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_map_migration_results", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_map_migration_results", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_map_migration_results", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_all_map_migration_results", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_map_migration_results", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_map_migration_results", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_map_migration_results", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetAllMapMigrationResultsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
