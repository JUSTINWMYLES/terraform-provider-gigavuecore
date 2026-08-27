package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
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
var _ list.ListResource = (*LoadAllAlarmsListResource)(nil)
var _ list.ListResourceWithConfigure = (*LoadAllAlarmsListResource)(nil)

// LoadAllAlarmsListResource is the generated Terraform list resource implementation.
type LoadAllAlarmsListResource struct {
	client *client.Client
}

// LoadAllAlarmsListResourceModel describes the gigavuecore_load_all_alarms list filter configuration shape.
type LoadAllAlarmsListResourceModel struct {
	Acknowledged     types.String `tfsdk:"acknowledged"`
	Acknowledgedby   types.String `tfsdk:"acknowledgedby"`
	Alias            types.String `tfsdk:"alias"`
	ClusterId        types.String `tfsdk:"cluster_id"`
	DeviceIp         types.String `tfsdk:"device_ip"`
	EndTime          types.String `tfsdk:"end_time"`
	Hostname         types.String `tfsdk:"hostname"`
	Page             types.String `tfsdk:"page"`
	ResourceId       types.String `tfsdk:"resource_id"`
	ResourceType     types.String `tfsdk:"resource_type"`
	Severity         types.String `tfsdk:"severity"`
	Sort             types.String `tfsdk:"sort"`
	StartTime        types.String `tfsdk:"start_time"`
	Suppressed       types.String `tfsdk:"suppressed"`
	Type             types.String `tfsdk:"type"`
	Unacknowledgedby types.String `tfsdk:"unacknowledgedby"`
}

// NewLoadAllAlarmsListResource returns a new instance of the generated list resource.
func NewLoadAllAlarmsListResource() list.ListResource {
	return &LoadAllAlarmsListResource{}
}

// Metadata returns the list resource type name.
func (l *LoadAllAlarmsListResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_all_alarms"
}

// ListResourceConfigSchema returns the list resource config schema.
func (l *LoadAllAlarmsListResource) ListResourceConfigSchema(_ context.Context, _ list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = listschema.Schema{MarkdownDescription: "Load All Alarms", Attributes: map[string]listschema.Attribute{"acknowledged": listschema.StringAttribute{MarkdownDescription: "filter by acknowledged alarms", Optional: true}, "acknowledgedby": listschema.StringAttribute{MarkdownDescription: "filter alarms that are acknowledged by a specific user", Optional: true}, "alias": listschema.StringAttribute{MarkdownDescription: " resource alias name to filter by", Optional: true}, "cluster_id": listschema.StringAttribute{MarkdownDescription: "Cluster ID to filter by", Optional: true}, "device_ip": listschema.StringAttribute{MarkdownDescription: "IP address of Device to filter by", Optional: true}, "end_time": listschema.StringAttribute{MarkdownDescription: "Filter End timestamps to include reports ending by. In ISO-8601 date format 'yyyy-MM-ddTHH:mm:ssZ'", Optional: true}, "hostname": listschema.StringAttribute{MarkdownDescription: "hostname to filter by", Optional: true}, "page": listschema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned. Defaults to (1:30) to prevent reading entire DB by accident", Optional: true}, "resource_id": listschema.StringAttribute{MarkdownDescription: "Affected entity to filter by", Optional: true}, "resource_type": listschema.StringAttribute{MarkdownDescription: "Affected entity type to filter by", Optional: true}, "severity": listschema.StringAttribute{MarkdownDescription: "Alarm severity to filter by", Optional: true}, "sort": listschema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is DESC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}, "start_time": listschema.StringAttribute{MarkdownDescription: "Filter Start timestamps to include reports starting from. In ISO-8601 date format 'yyyy-MM-ddTHH:mm:ssZ'", Optional: true}, "suppressed": listschema.StringAttribute{MarkdownDescription: "filter by suppressed alarms", Optional: true}, "type": listschema.StringAttribute{MarkdownDescription: "Alarm Type to filter by", Optional: true}, "unacknowledgedby": listschema.StringAttribute{MarkdownDescription: "filter alarms that are unacknowledged by a specific user", Optional: true}}}
}

// List streams matching resource instances for terraform query.
func (l *LoadAllAlarmsListResource) List(ctx context.Context, req list.ListRequest, stream *list.ListResultsStream) {
	stream.Results = func(push func(list.ListResult) bool) {
		var config LoadAllAlarmsListResourceModel
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
				result.Diagnostics.AddError("Error listing gigavuecore_load_all_alarms", fmt.Sprintf("Could not decode list item: %s", err))
				if !push(result) {
					return
				}
				continue
			}
			identity := map[string]json.RawMessage{}
			alarmIdValue, ok := itemMap["alarmId"]
			if !ok {
				if itemMap["metadata"] != nil {
					metaMap := map[string]json.RawMessage{}
					if json.Unmarshal(itemMap["metadata"], &metaMap) == nil {
						alarmIdValue, ok = metaMap["alarmId"]
					}
				}
			}
			if !ok {
				alarmIdValue, ok = itemMap["alarm_id"]
			}
			if !ok {
				alarmIdValue, ok = itemMap["id"]
			}
			if !ok {
				result.Diagnostics.AddError("Error listing gigavuecore_load_all_alarms", "List item is missing identity attribute \"alarm_id\".")
				if !push(result) {
					return
				}
				continue
			}
			identity["alarm_id"] = alarmIdValue
			idJSON, err := json.Marshal(identity)
			if err != nil {
				result.Diagnostics.AddError("Error listing gigavuecore_load_all_alarms", fmt.Sprintf("Could not encode list item identity: %s", err))
				if !push(result) {
					return
				}
				continue
			}
			idVal, err := tftypes.ValueFromJSON(idJSON, req.ResourceIdentitySchema.Type().TerraformType(ctx))
			if err != nil {
				result.Diagnostics.AddError("Error listing gigavuecore_load_all_alarms", fmt.Sprintf("Could not decode list item identity: %s", err))
				if !push(result) {
					return
				}
				continue
			}
			result.Identity.Raw = idVal
			if req.IncludeResource {
				resVal, err := tftypes.ValueFromJSON(item, req.ResourceSchema.Type().TerraformType(ctx))
				if err != nil {
					result.Diagnostics.AddWarning("Error listing gigavuecore_load_all_alarms", fmt.Sprintf("Could not decode list item into the resource schema: %s", err))
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
func (l *LoadAllAlarmsListResource) listRemote(ctx context.Context, config *LoadAllAlarmsListResourceModel) ([]json.RawMessage, diag.Diagnostics) {
	var diags diag.Diagnostics
	if l.client == nil {
		diags.AddError("Client Not Configured", "The API client was not set on the list resource. The provider Configure method must run before list operations; this is a bug in the generated provider.")
		return nil, diags
	}
	reqPath := "/alarms"
	params := url.Values{}
	if !config.Page.IsNull() {
		params.Set("page", config.Page.ValueString())
	}
	if !config.Sort.IsNull() {
		params.Set("sort", config.Sort.ValueString())
	}
	if !config.Type.IsNull() {
		params.Set("type", config.Type.ValueString())
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
	if !config.ClusterId.IsNull() {
		params.Set("clusterId", config.ClusterId.ValueString())
	}
	if !config.Hostname.IsNull() {
		params.Set("hostname", config.Hostname.ValueString())
	}
	if !config.Alias.IsNull() {
		params.Set("alias", config.Alias.ValueString())
	}
	if !config.DeviceIp.IsNull() {
		params.Set("deviceIp", config.DeviceIp.ValueString())
	}
	if !config.Acknowledged.IsNull() {
		params.Set("acknowledged", config.Acknowledged.ValueString())
	}
	if !config.Acknowledgedby.IsNull() {
		params.Set("acknowledgedby", config.Acknowledgedby.ValueString())
	}
	if !config.Unacknowledgedby.IsNull() {
		params.Set("unacknowledgedby", config.Unacknowledgedby.ValueString())
	}
	if !config.Suppressed.IsNull() {
		params.Set("suppressed", config.Suppressed.ValueString())
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
		diags.AddError("Error listing gigavuecore_load_all_alarms", fmt.Sprintf("Could not read list response: %s", err))
		return nil, diags
	}
	allItems := []json.RawMessage{}
	for _, page := range pages {
		pageObj := map[string]json.RawMessage{}
		if err := json.Unmarshal(page, &pageObj); err != nil {
			diags.AddError("Error listing gigavuecore_load_all_alarms", fmt.Sprintf("Could not decode list page: %s", err))
			return nil, diags
		}
		rawItems, ok := pageObj["alarms"]
		if !ok {
			diags.AddError("Error listing gigavuecore_load_all_alarms", fmt.Sprintf("Could not decode list page: missing %q array", "alarms"))
			return nil, diags
		}
		items := []json.RawMessage{}
		if err := json.Unmarshal(rawItems, &items); err != nil {
			diags.AddError("Error listing gigavuecore_load_all_alarms", fmt.Sprintf("Could not decode list page: %s", err))
			return nil, diags
		}
		allItems = append(allItems, items...)
	}
	return allItems, diags
}

// Configure stores the API client supplied by the provider.
func (l *LoadAllAlarmsListResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
