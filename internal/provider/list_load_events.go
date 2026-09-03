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
var _ list.ListResource = (*LoadEventsListResource)(nil)
var _ list.ListResourceWithConfigure = (*LoadEventsListResource)(nil)

// LoadEventsListResource is the generated Terraform list resource implementation.
type LoadEventsListResource struct {
	client *client.Client
}

// LoadEventsListResourceModel describes the gigavuecore_load_events list filter configuration shape.
type LoadEventsListResourceModel struct {
	EndTime      types.String `tfsdk:"end_time"`
	Page         types.String `tfsdk:"page"`
	ResourceId   types.String `tfsdk:"resource_id"`
	ResourceType types.String `tfsdk:"resource_type"`
	Scope        types.String `tfsdk:"scope"`
	Severity     types.String `tfsdk:"severity"`
	Sort         types.String `tfsdk:"sort"`
	Source       types.String `tfsdk:"source"`
	StartTime    types.String `tfsdk:"start_time"`
	Type         types.String `tfsdk:"type"`
}

// NewLoadEventsListResource returns a new instance of the generated list resource.
func NewLoadEventsListResource() list.ListResource {
	return &LoadEventsListResource{}
}

// Metadata returns the list resource type name.
func (l *LoadEventsListResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_events"
}

// ListResourceConfigSchema returns the list resource config schema.
func (l *LoadEventsListResource) ListResourceConfigSchema(_ context.Context, _ list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = listschema.Schema{MarkdownDescription: "Load Events", Attributes: map[string]listschema.Attribute{"end_time": listschema.StringAttribute{MarkdownDescription: "End Time to filter by. In ISO 8601 format", Optional: true}, "page": listschema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned. Defaults to (1:100) to prevent reading entire DB by accident", Optional: true}, "resource_id": listschema.StringAttribute{MarkdownDescription: "Affected Entity to filter by", Optional: true}, "resource_type": listschema.StringAttribute{MarkdownDescription: "Affected Entity Type to filter by", Optional: true}, "scope": listschema.StringAttribute{MarkdownDescription: "Event Scope to filter by", Optional: true}, "severity": listschema.StringAttribute{MarkdownDescription: "Event Scope to filter by", Optional: true}, "sort": listschema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}, "source": listschema.StringAttribute{MarkdownDescription: "Event Source identifier to filter by", Optional: true}, "start_time": listschema.StringAttribute{MarkdownDescription: "Start Time to filter by. In ISO 8601 format", Optional: true}, "type": listschema.StringAttribute{MarkdownDescription: "Event Type to filter by", Optional: true}}}
}

// List streams matching resource instances for terraform query.
func (l *LoadEventsListResource) List(ctx context.Context, req list.ListRequest, stream *list.ListResultsStream) {
	stream.Results = func(push func(list.ListResult) bool) {
		var config LoadEventsListResourceModel
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
				result.Diagnostics.AddError("Error listing gigavuecore_load_events", fmt.Sprintf("Could not decode list item: %s", err))
				if !push(result) {
					return
				}
				continue
			}
			identity := map[string]json.RawMessage{}
			eventIdValue, ok := itemMap["eventId"]
			if !ok {
				if itemMap["metadata"] != nil {
					metaMap := map[string]json.RawMessage{}
					if json.Unmarshal(itemMap["metadata"], &metaMap) == nil {
						eventIdValue, ok = metaMap["eventId"]
					}
				}
			}
			if !ok {
				eventIdValue, ok = itemMap["event_id"]
			}
			if !ok {
				eventIdValue, ok = itemMap["id"]
			}
			if !ok {
				result.Diagnostics.AddError("Error listing gigavuecore_load_events", "List item is missing identity attribute \"event_id\".")
				if !push(result) {
					return
				}
				continue
			}
			identity["event_id"] = eventIdValue
			idJSON, err := json.Marshal(identity)
			if err != nil {
				result.Diagnostics.AddError("Error listing gigavuecore_load_events", fmt.Sprintf("Could not encode list item identity: %s", err))
				if !push(result) {
					return
				}
				continue
			}
			idVal, err := tftypes.ValueFromJSON(idJSON, req.ResourceIdentitySchema.Type().TerraformType(ctx))
			if err != nil {
				result.Diagnostics.AddError("Error listing gigavuecore_load_events", fmt.Sprintf("Could not decode list item identity: %s", err))
				if !push(result) {
					return
				}
				continue
			}
			result.Identity.Raw = idVal
			if req.IncludeResource {
				resVal, err := tftypes.ValueFromJSON(item, req.ResourceSchema.Type().TerraformType(ctx))
				if err != nil {
					result.Diagnostics.AddWarning("Error listing gigavuecore_load_events", fmt.Sprintf("Could not decode list item into the resource schema: %s", err))
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
func (l *LoadEventsListResource) listRemote(ctx context.Context, config *LoadEventsListResourceModel) ([]json.RawMessage, diag.Diagnostics) {
	var diags diag.Diagnostics
	if l.client == nil {
		diags.AddError("Client Not Configured", "The API client was not set on the list resource. The provider Configure method must run before list operations; this is a bug in the generated provider.")
		return nil, diags
	}
	reqPath := "/events"
	params := url.Values{}
	if !config.Type.IsNull() {
		params.Set("type", config.Type.ValueString())
	}
	if !config.Source.IsNull() {
		params.Set("source", config.Source.ValueString())
	}
	if !config.Scope.IsNull() {
		params.Set("scope", config.Scope.ValueString())
	}
	if !config.Severity.IsNull() {
		params.Set("severity", config.Severity.ValueString())
	}
	if !config.StartTime.IsNull() {
		params.Set("startTime", config.StartTime.ValueString())
	}
	if !config.EndTime.IsNull() {
		params.Set("endTime", config.EndTime.ValueString())
	}
	if !config.ResourceType.IsNull() {
		params.Set("resourceType", config.ResourceType.ValueString())
	}
	if !config.ResourceId.IsNull() {
		params.Set("resourceId", config.ResourceId.ValueString())
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
		diags.AddError("Error listing gigavuecore_load_events", fmt.Sprintf("Could not read list response: %s", err))
		return nil, diags
	}
	allItems := []json.RawMessage{}
	for _, page := range pages {
		pageObj := map[string]json.RawMessage{}
		if err := json.Unmarshal(page, &pageObj); err != nil {
			diags.AddError("Error listing gigavuecore_load_events", fmt.Sprintf("Could not decode list page: %s", err))
			return nil, diags
		}
		rawItems, ok := pageObj["events"]
		if !ok {
			diags.AddError("Error listing gigavuecore_load_events", fmt.Sprintf("Could not decode list page: missing %q array", "events"))
			return nil, diags
		}
		items := []json.RawMessage{}
		if err := json.Unmarshal(rawItems, &items); err != nil {
			diags.AddError("Error listing gigavuecore_load_events", fmt.Sprintf("Could not decode list page: %s", err))
			return nil, diags
		}
		allItems = append(allItems, items...)
	}
	return allItems, diags
}

// Configure stores the API client supplied by the provider.
func (l *LoadEventsListResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
