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
var _ list.ListResource = (*LoadIpDestinationStatusesListResource)(nil)
var _ list.ListResourceWithConfigure = (*LoadIpDestinationStatusesListResource)(nil)

// LoadIpDestinationStatusesListResource is the generated Terraform list resource implementation.
type LoadIpDestinationStatusesListResource struct {
	client *client.Client
}

// LoadIpDestinationStatusesListResourceModel describes the gigavuecore_load_ip_destination_statuses list filter configuration shape.
type LoadIpDestinationStatusesListResourceModel struct {
	ClusterId        types.String `tfsdk:"cluster_id"`
	GsGroupAlias     types.String `tfsdk:"gs_group_alias"`
	IpInterfaceAlias types.String `tfsdk:"ip_interface_alias"`
	IpType           types.String `tfsdk:"ip_type"`
}

// NewLoadIpDestinationStatusesListResource returns a new instance of the generated list resource.
func NewLoadIpDestinationStatusesListResource() list.ListResource {
	return &LoadIpDestinationStatusesListResource{}
}

// Metadata returns the list resource type name.
func (l *LoadIpDestinationStatusesListResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_ip_destination_statuses"
}

// ListResourceConfigSchema returns the list resource config schema.
func (l *LoadIpDestinationStatusesListResource) ListResourceConfigSchema(_ context.Context, _ list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = listschema.Schema{MarkdownDescription: "Load IP destination status of all or those matching the IP Interface alias or IP type or GsGroup", Attributes: map[string]listschema.Attribute{"cluster_id": listschema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "gs_group_alias": listschema.StringAttribute{MarkdownDescription: "GsGroup alias", Optional: true}, "ip_interface_alias": listschema.StringAttribute{MarkdownDescription: "IP Interface alias", Optional: true}, "ip_type": listschema.StringAttribute{MarkdownDescription: "IP type", Optional: true}}}
}

// List streams matching resource instances for terraform query.
func (l *LoadIpDestinationStatusesListResource) List(ctx context.Context, req list.ListRequest, stream *list.ListResultsStream) {
	stream.Results = func(push func(list.ListResult) bool) {
		var config LoadIpDestinationStatusesListResourceModel
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
				result.Diagnostics.AddError("Error listing gigavuecore_load_ip_destination_statuses", fmt.Sprintf("Could not decode list item: %s", err))
				if !push(result) {
					return
				}
				continue
			}
			identity := map[string]json.RawMessage{}
			ipAddressValue, ok := itemMap["ipAddress"]
			if !ok {
				if itemMap["metadata"] != nil {
					metaMap := map[string]json.RawMessage{}
					if json.Unmarshal(itemMap["metadata"], &metaMap) == nil {
						ipAddressValue, ok = metaMap["ipAddress"]
					}
				}
			}
			if !ok {
				ipAddressValue, ok = itemMap["ip_address"]
			}
			if !ok {
				ipAddressValue, ok = itemMap["id"]
			}
			if !ok {
				result.Diagnostics.AddError("Error listing gigavuecore_load_ip_destination_statuses", "List item is missing identity attribute \"ip_address\".")
				if !push(result) {
					return
				}
				continue
			}
			identity["ip_address"] = ipAddressValue
			idJSON, err := json.Marshal(identity)
			if err != nil {
				result.Diagnostics.AddError("Error listing gigavuecore_load_ip_destination_statuses", fmt.Sprintf("Could not encode list item identity: %s", err))
				if !push(result) {
					return
				}
				continue
			}
			idVal, err := tftypes.ValueFromJSON(idJSON, req.ResourceIdentitySchema.Type().TerraformType(ctx))
			if err != nil {
				result.Diagnostics.AddError("Error listing gigavuecore_load_ip_destination_statuses", fmt.Sprintf("Could not decode list item identity: %s", err))
				if !push(result) {
					return
				}
				continue
			}
			result.Identity.Raw = idVal
			if req.IncludeResource {
				resVal, err := tftypes.ValueFromJSON(item, req.ResourceSchema.Type().TerraformType(ctx))
				if err != nil {
					result.Diagnostics.AddWarning("Error listing gigavuecore_load_ip_destination_statuses", fmt.Sprintf("Could not decode list item into the resource schema: %s", err))
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
func (l *LoadIpDestinationStatusesListResource) listRemote(ctx context.Context, config *LoadIpDestinationStatusesListResourceModel) ([]json.RawMessage, diag.Diagnostics) {
	var diags diag.Diagnostics
	if l.client == nil {
		diags.AddError("Client Not Configured", "The API client was not set on the list resource. The provider Configure method must run before list operations; this is a bug in the generated provider.")
		return nil, diags
	}
	reqPath := "/ip/destination"
	params := url.Values{}
	params.Set("clusterId", config.ClusterId.ValueString())
	if !config.IpInterfaceAlias.IsNull() {
		params.Set("ipInterfaceAlias", config.IpInterfaceAlias.ValueString())
	}
	if !config.GsGroupAlias.IsNull() {
		params.Set("gsGroupAlias", config.GsGroupAlias.ValueString())
	}
	if !config.IpType.IsNull() {
		params.Set("ipType", config.IpType.ValueString())
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
		diags.AddError("Error listing gigavuecore_load_ip_destination_statuses", fmt.Sprintf("Could not read list response: %s", err))
		return nil, diags
	}
	allItems := []json.RawMessage{}
	for _, page := range pages {
		pageObj := map[string]json.RawMessage{}
		if err := json.Unmarshal(page, &pageObj); err != nil {
			diags.AddError("Error listing gigavuecore_load_ip_destination_statuses", fmt.Sprintf("Could not decode list page: %s", err))
			return nil, diags
		}
		rawItems, ok := pageObj["ipInterfaces"]
		if !ok {
			diags.AddError("Error listing gigavuecore_load_ip_destination_statuses", fmt.Sprintf("Could not decode list page: missing %q array", "ipInterfaces"))
			return nil, diags
		}
		items := []json.RawMessage{}
		if err := json.Unmarshal(rawItems, &items); err != nil {
			diags.AddError("Error listing gigavuecore_load_ip_destination_statuses", fmt.Sprintf("Could not decode list page: %s", err))
			return nil, diags
		}
		allItems = append(allItems, items...)
	}
	return allItems, diags
}

// Configure stores the API client supplied by the provider.
func (l *LoadIpDestinationStatusesListResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
