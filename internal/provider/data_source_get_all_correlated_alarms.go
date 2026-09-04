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
	_ datasource.DataSource              = (*GetAllCorrelatedAlarmsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetAllCorrelatedAlarmsDataSource)(nil)
)

// GetAllCorrelatedAlarmsDataSource is the generated Terraform data source implementation.
type GetAllCorrelatedAlarmsDataSource struct {
	client *client.Client
}

// GetAllCorrelatedAlarmsDataSourceModel describes the data source state shape.
type GetAllCorrelatedAlarmsDataSourceModel struct {
	Acknowledged     types.String `tfsdk:"acknowledged"`
	Acknowledgedby   types.String `tfsdk:"acknowledgedby"`
	Alias            types.String `tfsdk:"alias"`
	ClusterId        types.String `tfsdk:"cluster_id" json:"clusterId"`
	DeviceIp         types.String `tfsdk:"device_ip" json:"deviceIp"`
	EndTime          types.String `tfsdk:"end_time" json:"endTime"`
	Hostname         types.String `tfsdk:"hostname"`
	Items            types.List   `tfsdk:"items"`
	Page             types.String `tfsdk:"page"`
	ResourceId       types.String `tfsdk:"resource_id" json:"resourceId"`
	ResourceType     types.String `tfsdk:"resource_type" json:"resourceType"`
	Severity         types.String `tfsdk:"severity"`
	Sort             types.String `tfsdk:"sort"`
	StartTime        types.String `tfsdk:"start_time" json:"startTime"`
	Suppressed       types.String `tfsdk:"suppressed"`
	Type             types.String `tfsdk:"type"`
	Unacknowledgedby types.String `tfsdk:"unacknowledgedby"`
}

// NewGetAllCorrelatedAlarmsDataSource returns a new instance of the generated data source.
func NewGetAllCorrelatedAlarmsDataSource() datasource.DataSource {
	return &GetAllCorrelatedAlarmsDataSource{}
}

// Metadata returns the data source type name.
func (d *GetAllCorrelatedAlarmsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_all_correlated_alarms"
}

// Schema returns the data source schema.
func (d *GetAllCorrelatedAlarmsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load All Correlated Alarms", Attributes: map[string]schema.Attribute{"acknowledged": schema.StringAttribute{MarkdownDescription: "filter by acknowledged alarms", Optional: true}, "acknowledgedby": schema.StringAttribute{MarkdownDescription: "filter alarms that are acknowledged by a specific user", Optional: true}, "alias": schema.StringAttribute{MarkdownDescription: " resource alias name to filter by", Optional: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Cluster ID to filter by", Optional: true}, "device_ip": schema.StringAttribute{MarkdownDescription: "IP address of Device to filter by", Optional: true}, "end_time": schema.StringAttribute{MarkdownDescription: "End Time to filter by. In ISO 8601 format 'yyyy-MM-ddTHH:mm:ssZ'", Optional: true}, "hostname": schema.StringAttribute{MarkdownDescription: "hostname to filter by", Optional: true}, "items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"acknowledged": schema.BoolAttribute{MarkdownDescription: "True if the alarm is acknowledged by the user", Computed: true}, "acknowledged_by": schema.StringAttribute{MarkdownDescription: "User who acknowledged the alarm", Computed: true}, "acknowledged_ts": schema.StringAttribute{MarkdownDescription: "Alarm acknowledged timestamp in ISO 8601 format", Computed: true}, "alias": schema.StringAttribute{MarkdownDescription: "Alias of the Resource ID", Computed: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Cluster ID of the device", Computed: true}, "comment": schema.StringAttribute{MarkdownDescription: "User comments", Computed: true}, "description": schema.StringAttribute{MarkdownDescription: "Alarm description", Computed: true}, "device_ip": schema.StringAttribute{MarkdownDescription: "IP address of the device", Computed: true}, "hostname": schema.StringAttribute{MarkdownDescription: "Hostname of the Resource", Computed: true}, "resource_id": schema.StringAttribute{MarkdownDescription: "Resource ID of the alarm", Computed: true}, "resource_type": schema.StringAttribute{MarkdownDescription: "Resource Type of the alarm", Computed: true}, "severity": schema.StringAttribute{MarkdownDescription: "Severity of the alarm", Computed: true}, "suppressed": schema.BoolAttribute{MarkdownDescription: "True if the alarm is suppressed", Computed: true}, "ts": schema.StringAttribute{MarkdownDescription: "Alarm timestamp in ISO 8601 format", Computed: true}, "type": schema.StringAttribute{MarkdownDescription: "Alarm Type identifier", Computed: true}, "unacknowledged_by": schema.StringAttribute{MarkdownDescription: "User who unacknowledged the alarm", Computed: true}, "unacknowledged_ts": schema.StringAttribute{MarkdownDescription: "Alarm unacknowledged timestamp in ISO 8601 format", Computed: true}}}}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned. Defaults to (1:30) to prevent reading entire DB by accident", Optional: true}, "resource_id": schema.StringAttribute{MarkdownDescription: "Affected entity to filter by", Optional: true}, "resource_type": schema.StringAttribute{MarkdownDescription: "Affected entity type to filter by", Optional: true}, "severity": schema.StringAttribute{MarkdownDescription: "Alarm severity to filter by", Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is DESC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}, "start_time": schema.StringAttribute{MarkdownDescription: "Start Time to filter by. In ISO 8601 format 'yyyy-MM-ddTHH:mm:ssZ'", Optional: true}, "suppressed": schema.StringAttribute{MarkdownDescription: "filter by suppressed alarms", Optional: true}, "type": schema.StringAttribute{MarkdownDescription: "Alarm Type to filter by", Optional: true}, "unacknowledgedby": schema.StringAttribute{MarkdownDescription: "filter alarms that are unacknowledged by a specific user", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetAllCorrelatedAlarmsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetAllCorrelatedAlarmsDataSourceModel
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
func (d *GetAllCorrelatedAlarmsDataSource) readListRemote(ctx context.Context, config *GetAllCorrelatedAlarmsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/alarms/correlated"
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_correlated_alarms", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_correlated_alarms", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["alarms"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_correlated_alarms", fmt.Sprintf("Could not decode list page: missing %q array", "alarms"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_correlated_alarms", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetAllCorrelatedAlarmsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
