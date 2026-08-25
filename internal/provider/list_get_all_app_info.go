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
var _ list.ListResource = (*GetAllAppInfoListResource)(nil)
var _ list.ListResourceWithConfigure = (*GetAllAppInfoListResource)(nil)

// GetAllAppInfoListResource is the generated Terraform list resource implementation.
type GetAllAppInfoListResource struct {
	client *client.Client
}

// GetAllAppInfoListResourceModel describes the gigavuecore_get_all_app_info list filter configuration shape.
type GetAllAppInfoListResourceModel struct {
	EnvId   types.String `tfsdk:"env_id"`
	UnifyId types.String `tfsdk:"unify_id"`
}

// NewGetAllAppInfoListResource returns a new instance of the generated list resource.
func NewGetAllAppInfoListResource() list.ListResource {
	return &GetAllAppInfoListResource{}
}

// Metadata returns the list resource type name.
func (l *GetAllAppInfoListResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_all_app_info"
}

// ListResourceConfigSchema returns the list resource config schema.
func (l *GetAllAppInfoListResource) ListResourceConfigSchema(_ context.Context, _ list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = listschema.Schema{MarkdownDescription: "List all appInfo for unified deployment via environment and connection id", Attributes: map[string]listschema.Attribute{"env_id": listschema.StringAttribute{MarkdownDescription: "unified environment identifier", Required: true}, "unify_id": listschema.StringAttribute{MarkdownDescription: "unified resource identifier", Required: true}}}
}

// List streams matching resource instances for terraform query.
func (l *GetAllAppInfoListResource) List(ctx context.Context, req list.ListRequest, stream *list.ListResultsStream) {
	stream.Results = func(push func(list.ListResult) bool) {
		var config GetAllAppInfoListResourceModel
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
				result.Diagnostics.AddError("Error listing gigavuecore_get_all_app_info", fmt.Sprintf("Could not decode list item: %s", err))
				if !push(result) {
					return
				}
				continue
			}
			identity := map[string]json.RawMessage{}
			envIdValue, ok := itemMap["envId"]
			if !ok {
				if itemMap["metadata"] != nil {
					metaMap := map[string]json.RawMessage{}
					if json.Unmarshal(itemMap["metadata"], &metaMap) == nil {
						envIdValue, ok = metaMap["envId"]
					}
				}
			}
			if !ok {
				envIdValue, ok = itemMap["env_id"]
			}
			if !ok {
				envIdValue, ok = itemMap["id"]
			}
			if !ok {
				result.Diagnostics.AddError("Error listing gigavuecore_get_all_app_info", "List item is missing identity attribute \"env_id\".")
				if !push(result) {
					return
				}
				continue
			}
			identity["env_id"] = envIdValue
			unifyIdValue, ok := itemMap["unifyId"]
			if !ok {
				if itemMap["metadata"] != nil {
					metaMap := map[string]json.RawMessage{}
					if json.Unmarshal(itemMap["metadata"], &metaMap) == nil {
						unifyIdValue, ok = metaMap["unifyId"]
					}
				}
			}
			if !ok {
				unifyIdValue, ok = itemMap["unify_id"]
			}
			if !ok {
				unifyIdValue, ok = itemMap["id"]
			}
			if !ok {
				result.Diagnostics.AddError("Error listing gigavuecore_get_all_app_info", "List item is missing identity attribute \"unify_id\".")
				if !push(result) {
					return
				}
				continue
			}
			identity["unify_id"] = unifyIdValue
			appinfoValue, ok := itemMap["appinfo"]
			if !ok {
				if itemMap["metadata"] != nil {
					metaMap := map[string]json.RawMessage{}
					if json.Unmarshal(itemMap["metadata"], &metaMap) == nil {
						appinfoValue, ok = metaMap["appinfo"]
					}
				}
			}
			if !ok {
				appinfoValue, ok = itemMap["id"]
			}
			if !ok {
				result.Diagnostics.AddError("Error listing gigavuecore_get_all_app_info", "List item is missing identity attribute \"appinfo\".")
				if !push(result) {
					return
				}
				continue
			}
			identity["appinfo"] = appinfoValue
			idJSON, err := json.Marshal(identity)
			if err != nil {
				result.Diagnostics.AddError("Error listing gigavuecore_get_all_app_info", fmt.Sprintf("Could not encode list item identity: %s", err))
				if !push(result) {
					return
				}
				continue
			}
			idVal, err := tftypes.ValueFromJSON(idJSON, req.ResourceIdentitySchema.Type().TerraformType(ctx))
			if err != nil {
				result.Diagnostics.AddError("Error listing gigavuecore_get_all_app_info", fmt.Sprintf("Could not decode list item identity: %s", err))
				if !push(result) {
					return
				}
				continue
			}
			result.Identity.Raw = idVal
			if req.IncludeResource {
				resVal, err := tftypes.ValueFromJSON(item, req.ResourceSchema.Type().TerraformType(ctx))
				if err != nil {
					result.Diagnostics.AddWarning("Error listing gigavuecore_get_all_app_info", fmt.Sprintf("Could not decode list item into the resource schema: %s", err))
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
func (l *GetAllAppInfoListResource) listRemote(ctx context.Context, config *GetAllAppInfoListResourceModel) ([]json.RawMessage, diag.Diagnostics) {
	var diags diag.Diagnostics
	if l.client == nil {
		diags.AddError("Client Not Configured", "The API client was not set on the list resource. The provider Configure method must run before list operations; this is a bug in the generated provider.")
		return nil, diags
	}
	reqPath := "/unifiedResources/env/{envId}/unifyId/{unifyId}/appinfo"
	reqPath = strings.ReplaceAll(reqPath, "{envId}", url.PathEscape(config.EnvId.ValueString()))
	reqPath = strings.ReplaceAll(reqPath, "{unifyId}", url.PathEscape(config.UnifyId.ValueString()))
	params := url.Values{}
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
		diags.AddError("Error listing gigavuecore_get_all_app_info", fmt.Sprintf("Could not read list response: %s", err))
		return nil, diags
	}
	allItems := []json.RawMessage{}
	for _, page := range pages {
		pageObj := map[string]json.RawMessage{}
		if err := json.Unmarshal(page, &pageObj); err != nil {
			diags.AddError("Error listing gigavuecore_get_all_app_info", fmt.Sprintf("Could not decode list page: %s", err))
			return nil, diags
		}
		rawItems, ok := pageObj["env"]
		if !ok {
			diags.AddError("Error listing gigavuecore_get_all_app_info", fmt.Sprintf("Could not decode list page: missing %q array", "env"))
			return nil, diags
		}
		items := []json.RawMessage{}
		if err := json.Unmarshal(rawItems, &items); err != nil {
			diags.AddError("Error listing gigavuecore_get_all_app_info", fmt.Sprintf("Could not decode list page: %s", err))
			return nil, diags
		}
		allItems = append(allItems, items...)
	}
	return allItems, diags
}

// Configure stores the API client supplied by the provider.
func (l *GetAllAppInfoListResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
