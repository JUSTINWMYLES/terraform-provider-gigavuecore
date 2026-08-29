package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
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
var _ list.ListResource = (*GetAllClusterCircuitTunnelsOfAnUserFabricMapListResource)(nil)
var _ list.ListResourceWithConfigure = (*GetAllClusterCircuitTunnelsOfAnUserFabricMapListResource)(nil)

// GetAllClusterCircuitTunnelsOfAnUserFabricMapListResource is the generated Terraform list resource implementation.
type GetAllClusterCircuitTunnelsOfAnUserFabricMapListResource struct {
	client *client.Client
}

// GetAllClusterCircuitTunnelsOfAnUserFabricMapListResourceModel describes the gigavuecore_get_all_cluster_circuit_tunnels_of_an_user_fabric_map list filter configuration shape.
type GetAllClusterCircuitTunnelsOfAnUserFabricMapListResourceModel struct {
	Alias types.String `tfsdk:"alias"`
	Mode  types.String `tfsdk:"mode"`
	Type  types.String `tfsdk:"type"`
}

// NewGetAllClusterCircuitTunnelsOfAnUserFabricMapListResource returns a new instance of the generated list resource.
func NewGetAllClusterCircuitTunnelsOfAnUserFabricMapListResource() list.ListResource {
	return &GetAllClusterCircuitTunnelsOfAnUserFabricMapListResource{}
}

// Metadata returns the list resource type name.
func (l *GetAllClusterCircuitTunnelsOfAnUserFabricMapListResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_all_cluster_circuit_tunnels_of_an_user_fabric_map"
}

// ListResourceConfigSchema returns the list resource config schema.
func (l *GetAllClusterCircuitTunnelsOfAnUserFabricMapListResource) ListResourceConfigSchema(_ context.Context, _ list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = listschema.Schema{MarkdownDescription: "Get all cluster circuit tunnel endpoints of a user-defined fabric map", Attributes: map[string]listschema.Attribute{"alias": listschema.StringAttribute{MarkdownDescription: "alias of the fabric map", Required: true}, "mode": listschema.StringAttribute{MarkdownDescription: "tunnel mode: 'encap' or 'decap'.", Optional: true}, "type": listschema.StringAttribute{MarkdownDescription: "tunnel type: 'circuit' or 'vxlan'.", Optional: true}}}
}

// List streams matching resource instances for terraform query.
func (l *GetAllClusterCircuitTunnelsOfAnUserFabricMapListResource) List(ctx context.Context, req list.ListRequest, stream *list.ListResultsStream) {
	stream.Results = func(push func(list.ListResult) bool) {
		var config GetAllClusterCircuitTunnelsOfAnUserFabricMapListResourceModel
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
				result.Diagnostics.AddError("Error listing gigavuecore_get_all_cluster_circuit_tunnels_of_an_user_fabric_map", fmt.Sprintf("Could not decode list item: %s", err))
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
				result.Diagnostics.AddError("Error listing gigavuecore_get_all_cluster_circuit_tunnels_of_an_user_fabric_map", "List item is missing identity attribute \"alias\".")
				if !push(result) {
					return
				}
				continue
			}
			identity["alias"] = aliasValue
			cctAliasValue, ok := itemMap["cctAlias"]
			if !ok {
				if itemMap["metadata"] != nil {
					metaMap := map[string]json.RawMessage{}
					if json.Unmarshal(itemMap["metadata"], &metaMap) == nil {
						cctAliasValue, ok = metaMap["cctAlias"]
					}
				}
			}
			if !ok {
				cctAliasValue, ok = itemMap["cct_alias"]
			}
			if !ok {
				cctAliasValue, ok = itemMap["id"]
			}
			if !ok {
				result.Diagnostics.AddError("Error listing gigavuecore_get_all_cluster_circuit_tunnels_of_an_user_fabric_map", "List item is missing identity attribute \"cct_alias\".")
				if !push(result) {
					return
				}
				continue
			}
			identity["cct_alias"] = cctAliasValue
			idJSON, err := json.Marshal(identity)
			if err != nil {
				result.Diagnostics.AddError("Error listing gigavuecore_get_all_cluster_circuit_tunnels_of_an_user_fabric_map", fmt.Sprintf("Could not encode list item identity: %s", err))
				if !push(result) {
					return
				}
				continue
			}
			idVal, err := tftypes.ValueFromJSON(idJSON, req.ResourceIdentitySchema.Type().TerraformType(ctx))
			if err != nil {
				result.Diagnostics.AddError("Error listing gigavuecore_get_all_cluster_circuit_tunnels_of_an_user_fabric_map", fmt.Sprintf("Could not decode list item identity: %s", err))
				if !push(result) {
					return
				}
				continue
			}
			result.Identity.Raw = idVal
			if req.IncludeResource {
				resVal, err := tftypes.ValueFromJSON(item, req.ResourceSchema.Type().TerraformType(ctx))
				if err != nil {
					result.Diagnostics.AddWarning("Error listing gigavuecore_get_all_cluster_circuit_tunnels_of_an_user_fabric_map", fmt.Sprintf("Could not decode list item into the resource schema: %s", err))
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
func (l *GetAllClusterCircuitTunnelsOfAnUserFabricMapListResource) listRemote(ctx context.Context, config *GetAllClusterCircuitTunnelsOfAnUserFabricMapListResourceModel) ([]json.RawMessage, diag.Diagnostics) {
	var diags diag.Diagnostics
	if l.client == nil {
		diags.AddError("Client Not Configured", "The API client was not set on the list resource. The provider Configure method must run before list operations; this is a bug in the generated provider.")
		return nil, diags
	}
	reqPath := "/fabricMaps/{alias}/clusterCircuitTunnels"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(config.Alias.ValueString()))
	params := url.Values{}
	if !config.Mode.IsNull() {
		params.Set("mode", config.Mode.ValueString())
	}
	if !config.Type.IsNull() {
		params.Set("type", config.Type.ValueString())
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
		diags.AddError("Error listing gigavuecore_get_all_cluster_circuit_tunnels_of_an_user_fabric_map", fmt.Sprintf("Could not read list response: %s", err))
		return nil, diags
	}
	allItems := []json.RawMessage{}
	for _, page := range pages {
		pageObj := map[string]json.RawMessage{}
		if err := json.Unmarshal(page, &pageObj); err != nil {
			diags.AddError("Error listing gigavuecore_get_all_cluster_circuit_tunnels_of_an_user_fabric_map", fmt.Sprintf("Could not decode list page: %s", err))
			return nil, diags
		}
		rawItems, ok := pageObj["clusterMaps"]
		if !ok {
			diags.AddError("Error listing gigavuecore_get_all_cluster_circuit_tunnels_of_an_user_fabric_map", fmt.Sprintf("Could not decode list page: missing %q array", "clusterMaps"))
			return nil, diags
		}
		items := []json.RawMessage{}
		if err := json.Unmarshal(rawItems, &items); err != nil {
			diags.AddError("Error listing gigavuecore_get_all_cluster_circuit_tunnels_of_an_user_fabric_map", fmt.Sprintf("Could not decode list page: %s", err))
			return nil, diags
		}
		allItems = append(allItems, items...)
	}
	return allItems, diags
}

// Configure stores the API client supplied by the provider.
func (l *GetAllClusterCircuitTunnelsOfAnUserFabricMapListResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
