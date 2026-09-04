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
	_ datasource.DataSource              = (*GetAllMapMigrationResultsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetAllMapMigrationResultsDataSource)(nil)
)

// GetAllMapMigrationResultsDataSource is the generated Terraform data source implementation.
type GetAllMapMigrationResultsDataSource struct {
	client *client.Client
}

// GetAllMapMigrationResultsDataSourceModel describes the data source state shape.
type GetAllMapMigrationResultsDataSourceModel struct {
	Items types.List `tfsdk:"items"`
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
	resp.Schema = schema.Schema{MarkdownDescription: "Load all map migration results", Attributes: map[string]schema.Attribute{"items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias_mismatched": schema.MapAttribute{Computed: true, ElementType: types.StringType}, "aliases": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "avg_time_per_map": schema.Float64Attribute{Computed: true}, "cluster_id": schema.StringAttribute{Computed: true}, "cluster_ids": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "clusters_with_no_maps": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "dry_run": schema.BoolAttribute{Computed: true}, "duplicate_maps": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "end_time": schema.Int64Attribute{Computed: true}, "failed_cluster_and_reasons": schema.MapAttribute{Computed: true, ElementType: types.StringType}, "failed_maps_and_reasons": schema.MapAttribute{Computed: true, ElementType: types.StringType}, "failure_reason": schema.StringAttribute{Computed: true}, "is_all_maps": schema.BoolAttribute{Computed: true}, "map_type": schema.StringAttribute{Computed: true}, "migrated_maps": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "migration_alias": schema.StringAttribute{Computed: true}, "skipped_migration": schema.MapAttribute{Computed: true, ElementType: types.StringType}, "start_time": schema.Int64Attribute{Computed: true}, "status": schema.StringAttribute{Computed: true}, "total_duplicates": schema.Int64Attribute{Computed: true}, "total_failed": schema.Int64Attribute{Computed: true}, "total_maps": schema.Int64Attribute{Computed: true}, "total_port_alias_mismatches": schema.Int64Attribute{Computed: true}, "total_skipped": schema.Int64Attribute{Computed: true}, "total_success": schema.Int64Attribute{Computed: true}, "total_time_to_complete": schema.Int64Attribute{Computed: true}, "triggered_by": schema.StringAttribute{Computed: true}}}}}}
}

// Read fetches remote state into the data source model.
func (d *GetAllMapMigrationResultsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetAllMapMigrationResultsDataSourceModel
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
func (d *GetAllMapMigrationResultsDataSource) readListRemote(ctx context.Context, config *GetAllMapMigrationResultsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/traffic-flows/migrationResult/all"
	params := url.Values{}
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_map_migration_results", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_map_migration_results", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["mapMigrationResults"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_map_migration_results", fmt.Sprintf("Could not decode list page: missing %q array", "mapMigrationResults"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
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
