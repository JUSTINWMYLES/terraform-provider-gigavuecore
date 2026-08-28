package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)
import (
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	diag "github.com/hashicorp/terraform-plugin-framework/diag"
	list "github.com/hashicorp/terraform-plugin-framework/list"
	listschema "github.com/hashicorp/terraform-plugin-framework/list/schema"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	types "github.com/hashicorp/terraform-plugin-framework/types"
	tftypes "github.com/hashicorp/terraform-plugin-go/tftypes"
)

// Compile-time interface assertion.
var _ list.ListResource = (*GetMapMigrationResultsListResource)(nil)
var _ list.ListResourceWithConfigure = (*GetMapMigrationResultsListResource)(nil)

// GetMapMigrationResultsListResource is the generated Terraform list resource implementation.
type GetMapMigrationResultsListResource struct {
	client *client.Client
}

// GetMapMigrationResultsListResourceModel describes the gigavuecore_get_map_migration_results list filter configuration shape.
type GetMapMigrationResultsListResourceModel struct {
	ClusterId      types.String `tfsdk:"cluster_id"`
	IsAllMaps      types.Bool   `tfsdk:"is_all_maps"`
	IsDryRun       types.Bool   `tfsdk:"is_dry_run"`
	MapAlias       types.String `tfsdk:"map_alias"`
	MapType        types.String `tfsdk:"map_type"`
	MigrationAlias types.String `tfsdk:"migration_alias"`
	Page           types.String `tfsdk:"page"`
	Sort           types.String `tfsdk:"sort"`
	Status         types.String `tfsdk:"status"`
}

// NewGetMapMigrationResultsListResource returns a new instance of the generated list resource.
func NewGetMapMigrationResultsListResource() list.ListResource {
	return &GetMapMigrationResultsListResource{}
}

// Metadata returns the list resource type name.
func (l *GetMapMigrationResultsListResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_map_migration_results"
}

// ListResourceConfigSchema returns the list resource config schema.
func (l *GetMapMigrationResultsListResource) ListResourceConfigSchema(_ context.Context, _ list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = listschema.Schema{MarkdownDescription: "Query map migration results with filters", Attributes: map[string]listschema.Attribute{"cluster_id": listschema.StringAttribute{MarkdownDescription: "Target Cluster ID", Optional: true}, "is_all_maps": listschema.BoolAttribute{MarkdownDescription: "Include all maps", Optional: true}, "is_dry_run": listschema.BoolAttribute{MarkdownDescription: "Filter by dry run status", Optional: true}, "map_alias": listschema.StringAttribute{MarkdownDescription: "Map alias", Optional: true}, "map_type": listschema.StringAttribute{MarkdownDescription: "Map type", Optional: true}, "migration_alias": listschema.StringAttribute{MarkdownDescription: "Migration alias", Optional: true}, "page": listschema.StringAttribute{MarkdownDescription: "Pagination request string (e.g., \"(1:30)\")", Optional: true}, "sort": listschema.StringAttribute{MarkdownDescription: "Sort request string (e.g., \"(startTime:DESC)\")", Optional: true}, "status": listschema.StringAttribute{MarkdownDescription: "Migration status", Optional: true}}}
}

// List streams matching resource instances for terraform query.
func (l *GetMapMigrationResultsListResource) List(ctx context.Context, req list.ListRequest, stream *list.ListResultsStream) {
	stream.Results = func(push func(list.ListResult) bool) {
		var config GetMapMigrationResultsListResourceModel
		diags := req.Config.Get(ctx, &config)
		if diags.HasError() {
			result := req.NewListResult(ctx)
			result.Diagnostics = diags
			push(result)
			return
		}
		items, diags := l.listRemote(ctx, &config)
		if diags.HasError() {
			result := req.NewListResult(ctx)
			result.Diagnostics = diags
			push(result)
			return
		}
		for _, item := range items {
			result := req.NewListResult(ctx)
			itemMap := map[string]json.RawMessage{}
			if err := json.Unmarshal(item, &itemMap); err != nil {
				result.Diagnostics.AddError("Error listing gigavuecore_get_map_migration_results", fmt.Sprintf("Could not decode list item: %s", err))
				if !push(result) {
					return
				}
				continue
			}
			identity := map[string]json.RawMessage{}
			aliasValue, ok := itemMap["alias"]
			if !ok {
				if itemMap["metadata"] != nil {
					metaMap := map[string]json.RawMessage{}
					if json.Unmarshal(itemMap["metadata"], &metaMap) == nil {
						aliasValue, ok = metaMap["alias"]
					}
				}
			}
			if !ok {
				aliasValue, ok = itemMap["id"]
			}
			if !ok {
				result.Diagnostics.AddError("Error listing gigavuecore_get_map_migration_results", "List item is missing identity attribute \"alias\".")
				if !push(result) {
					return
				}
				continue
			}
			identity["alias"] = aliasValue
			idJSON, err := json.Marshal(identity)
			if err != nil {
				result.Diagnostics.AddError("Error listing gigavuecore_get_map_migration_results", fmt.Sprintf("Could not encode list item identity: %s", err))
				if !push(result) {
					return
				}
				continue
			}
			idVal, err := tftypes.ValueFromJSON(idJSON, req.ResourceIdentitySchema.Type().TerraformType(ctx))
			if err != nil {
				result.Diagnostics.AddError("Error listing gigavuecore_get_map_migration_results", fmt.Sprintf("Could not decode list item identity: %s", err))
				if !push(result) {
					return
				}
				continue
			}
			result.Identity.Raw = idVal
			if req.IncludeResource {
				resVal, err := tftypes.ValueFromJSON(item, req.ResourceSchema.Type().TerraformType(ctx))
				if err != nil {
					result.Diagnostics.AddWarning("Error listing gigavuecore_get_map_migration_results", fmt.Sprintf("Could not decode list item into the resource schema: %s", err))
				} else {
					result.Resource.Raw = resVal
				}
			}
			if !push(result) {
				return
			}
		}
	}
}

// listRemote fetches and decodes the collection pages, returning the items and any diagnostics for the List iterator to surface.
func (l *GetMapMigrationResultsListResource) listRemote(ctx context.Context, config *GetMapMigrationResultsListResourceModel) ([]json.RawMessage, diag.Diagnostics) {
	var diags diag.Diagnostics
	if l.client == nil {
		diags.AddError("Client Not Configured", "The API client was not set on the list resource. The provider Configure method must run before list operations; this is a bug in the generated provider.")
		return nil, diags
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
		httpReq, err := l.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
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
		return l.client.Do(httpReq)
	}
	pages, err := client.ListAllPages(ctx, params, fetch, nil)
	if err != nil {
		diags.AddError("Error listing gigavuecore_get_map_migration_results", fmt.Sprintf("Could not read list response: %s", err))
		return nil, diags
	}
	allItems := []json.RawMessage{}
	for _, page := range pages {
		pageObj := map[string]json.RawMessage{}
		if err := json.Unmarshal(page, &pageObj); err != nil {
			diags.AddError("Error listing gigavuecore_get_map_migration_results", fmt.Sprintf("Could not decode list page: %s", err))
			return nil, diags
		}
		rawItems, ok := pageObj["mapMigrationResults"]
		if !ok {
			diags.AddError("Error listing gigavuecore_get_map_migration_results", fmt.Sprintf("Could not decode list page: missing %q array", "mapMigrationResults"))
			return nil, diags
		}
		items := []json.RawMessage{}
		if err := json.Unmarshal(rawItems, &items); err != nil {
			diags.AddError("Error listing gigavuecore_get_map_migration_results", fmt.Sprintf("Could not decode list page: %s", err))
			return nil, diags
		}
		allItems = append(allItems, items...)
	}
	return allItems, diags
}

// Configure stores the API client supplied by the provider.
func (l *GetMapMigrationResultsListResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected List Resource Configure Type", fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData))
		return
	}
	l.client = c
}
