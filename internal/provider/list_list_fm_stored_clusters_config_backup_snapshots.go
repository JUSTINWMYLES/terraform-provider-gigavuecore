package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)
import (
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	diag "github.com/hashicorp/terraform-plugin-framework/diag"
	list "github.com/hashicorp/terraform-plugin-framework/list"
	listschema "github.com/hashicorp/terraform-plugin-framework/list/schema"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	types "github.com/hashicorp/terraform-plugin-framework/types"
	tftypes "github.com/hashicorp/terraform-plugin-go/tftypes"
)

// Compile-time interface assertion.
var _ list.ListResource = (*ListFmStoredClustersConfigBackupSnapshotsListResource)(nil)
var _ list.ListResourceWithConfigure = (*ListFmStoredClustersConfigBackupSnapshotsListResource)(nil)

// ListFmStoredClustersConfigBackupSnapshotsListResource is the generated Terraform list resource implementation.
type ListFmStoredClustersConfigBackupSnapshotsListResource struct {
	client *client.Client
}

// ListFmStoredClustersConfigBackupSnapshotsListResourceModel describes the gigavuecore_list_fm_stored_clusters_config_backup_snapshots list filter configuration shape.
type ListFmStoredClustersConfigBackupSnapshotsListResourceModel struct {
	ClusterId types.String `tfsdk:"cluster_id"`
}

// NewListFmStoredClustersConfigBackupSnapshotsListResource returns a new instance of the generated list resource.
func NewListFmStoredClustersConfigBackupSnapshotsListResource() list.ListResource {
	return &ListFmStoredClustersConfigBackupSnapshotsListResource{}
}

// Metadata returns the list resource type name.
func (l *ListFmStoredClustersConfigBackupSnapshotsListResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_list_fm_stored_clusters_config_backup_snapshots"
}

// ListResourceConfigSchema returns the list resource config schema.
func (l *ListFmStoredClustersConfigBackupSnapshotsListResource) ListResourceConfigSchema(_ context.Context, _ list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = listschema.Schema{MarkdownDescription: "Lists config backup snapshots for multiple clusters", Attributes: map[string]listschema.Attribute{"cluster_id": listschema.StringAttribute{MarkdownDescription: "if provided, only backups associated with this cluster will be returned", Optional: true}}}
}

// List streams matching resource instances for terraform query.
func (l *ListFmStoredClustersConfigBackupSnapshotsListResource) List(ctx context.Context, req list.ListRequest, stream *list.ListResultsStream) {
	stream.Results = func(push func(list.ListResult) bool) {
		var config ListFmStoredClustersConfigBackupSnapshotsListResourceModel
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
				result.Diagnostics.AddError("Error listing gigavuecore_list_fm_stored_clusters_config_backup_snapshots", fmt.Sprintf("Could not decode list item: %s", err))
				if !push(result) {
					return
				}
				continue
			}
			identity := map[string]json.RawMessage{}
			clusterIdValue, ok := itemMap["clusterId"]
			if !ok {
				if itemMap["metadata"] != nil {
					metaMap := map[string]json.RawMessage{}
					if json.Unmarshal(itemMap["metadata"], &metaMap) == nil {
						clusterIdValue, ok = metaMap["clusterId"]
					}
				}
			}
			if !ok {
				clusterIdValue, ok = itemMap["cluster_id"]
			}
			if !ok {
				clusterIdValue, ok = itemMap["id"]
			}
			if !ok {
				result.Diagnostics.AddError("Error listing gigavuecore_list_fm_stored_clusters_config_backup_snapshots", "List item is missing identity attribute \"cluster_id\".")
				if !push(result) {
					return
				}
				continue
			}
			identity["cluster_id"] = clusterIdValue
			backupIdValue, ok := itemMap["backupId"]
			if !ok {
				if itemMap["metadata"] != nil {
					metaMap := map[string]json.RawMessage{}
					if json.Unmarshal(itemMap["metadata"], &metaMap) == nil {
						backupIdValue, ok = metaMap["backupId"]
					}
				}
			}
			if !ok {
				backupIdValue, ok = itemMap["backup_id"]
			}
			if !ok {
				backupIdValue, ok = itemMap["id"]
			}
			if !ok {
				result.Diagnostics.AddError("Error listing gigavuecore_list_fm_stored_clusters_config_backup_snapshots", "List item is missing identity attribute \"backup_id\".")
				if !push(result) {
					return
				}
				continue
			}
			identity["backup_id"] = backupIdValue
			idJSON, err := json.Marshal(identity)
			if err != nil {
				result.Diagnostics.AddError("Error listing gigavuecore_list_fm_stored_clusters_config_backup_snapshots", fmt.Sprintf("Could not encode list item identity: %s", err))
				if !push(result) {
					return
				}
				continue
			}
			idVal, err := tftypes.ValueFromJSON(idJSON, req.ResourceIdentitySchema.Type().TerraformType(ctx))
			if err != nil {
				result.Diagnostics.AddError("Error listing gigavuecore_list_fm_stored_clusters_config_backup_snapshots", fmt.Sprintf("Could not decode list item identity: %s", err))
				if !push(result) {
					return
				}
				continue
			}
			result.Identity.Raw = idVal
			if req.IncludeResource {
				resVal, err := tftypes.ValueFromJSON(item, req.ResourceSchema.Type().TerraformType(ctx))
				if err != nil {
					result.Diagnostics.AddWarning("Error listing gigavuecore_list_fm_stored_clusters_config_backup_snapshots", fmt.Sprintf("Could not decode list item into the resource schema: %s", err))
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
func (l *ListFmStoredClustersConfigBackupSnapshotsListResource) listRemote(ctx context.Context, config *ListFmStoredClustersConfigBackupSnapshotsListResourceModel) ([]json.RawMessage, diag.Diagnostics) {
	var diags diag.Diagnostics
	if l.client == nil {
		diags.AddError("Client Not Configured", "The API client was not set on the list resource. The provider Configure method must run before list operations; this is a bug in the generated provider.")
		return nil, diags
	}
	reqPath := "/clusterConfig/backup/repo"
	params := url.Values{}
	if !config.ClusterId.IsNull() {
		params.Set("clusterId", config.ClusterId.ValueString())
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
		diags.AddError("Error listing gigavuecore_list_fm_stored_clusters_config_backup_snapshots", fmt.Sprintf("Could not read list response: %s", err))
		return nil, diags
	}
	allItems := []json.RawMessage{}
	for _, page := range pages {
		pageObj := map[string]json.RawMessage{}
		if err := json.Unmarshal(page, &pageObj); err != nil {
			diags.AddError("Error listing gigavuecore_list_fm_stored_clusters_config_backup_snapshots", fmt.Sprintf("Could not decode list page: %s", err))
			return nil, diags
		}
		rawItems, ok := pageObj["clustersConfigBackups"]
		if !ok {
			diags.AddError("Error listing gigavuecore_list_fm_stored_clusters_config_backup_snapshots", fmt.Sprintf("Could not decode list page: missing %q array", "clustersConfigBackups"))
			return nil, diags
		}
		items := []json.RawMessage{}
		if err := json.Unmarshal(rawItems, &items); err != nil {
			diags.AddError("Error listing gigavuecore_list_fm_stored_clusters_config_backup_snapshots", fmt.Sprintf("Could not decode list page: %s", err))
			return nil, diags
		}
		allItems = append(allItems, items...)
	}
	return allItems, diags
}

// Configure stores the API client supplied by the provider.
func (l *ListFmStoredClustersConfigBackupSnapshotsListResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
