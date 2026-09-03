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
	_ datasource.DataSource              = (*GetAllAlarmSuppressionMetadataDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetAllAlarmSuppressionMetadataDataSource)(nil)
)

// GetAllAlarmSuppressionMetadataDataSource is the generated Terraform data source implementation.
type GetAllAlarmSuppressionMetadataDataSource struct {
	client *client.Client
}

// GetAllAlarmSuppressionMetadataDataSourceModel describes the data source state shape.
type GetAllAlarmSuppressionMetadataDataSourceModel struct {
	AlarmTypes     types.String `tfsdk:"alarm_types" json:"alarmTypes"`
	Alias          types.String `tfsdk:"alias"`
	ClusterId      types.String `tfsdk:"cluster_id" json:"clusterId"`
	Hostname       types.String `tfsdk:"hostname"`
	Items          types.List   `tfsdk:"items"`
	Page           types.String `tfsdk:"page"`
	ResourceId     types.String `tfsdk:"resource_id" json:"resourceId"`
	SelectedReason types.String `tfsdk:"selected_reason" json:"selectedReason"`
	Sort           types.String `tfsdk:"sort"`
}

// NewGetAllAlarmSuppressionMetadataDataSource returns a new instance of the generated data source.
func NewGetAllAlarmSuppressionMetadataDataSource() datasource.DataSource {
	return &GetAllAlarmSuppressionMetadataDataSource{}
}

// Metadata returns the data source type name.
func (d *GetAllAlarmSuppressionMetadataDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_all_alarm_suppression_metadata"
}

// Schema returns the data source schema.
func (d *GetAllAlarmSuppressionMetadataDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get All Alarm Supppression Rules", Attributes: map[string]schema.Attribute{"alarm_types": schema.StringAttribute{MarkdownDescription: "comma-separated list of Alarm Type's to filter", Optional: true}, "alias": schema.StringAttribute{MarkdownDescription: "Alias", Optional: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Cluster Id", Optional: true}, "hostname": schema.StringAttribute{MarkdownDescription: "Hostname", Optional: true}, "items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alarm_types": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "alias": schema.StringAttribute{MarkdownDescription: "Alias", Computed: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Cluster/Node Id", Computed: true}, "created_by": schema.StringAttribute{MarkdownDescription: "Alarm suppression rule created by FM User", Computed: true}, "created_ts": schema.StringAttribute{MarkdownDescription: "Alarm suppression rule creation timestamp in ISO 8601 format", Computed: true}, "enable": schema.BoolAttribute{MarkdownDescription: "FM Alarm Suppression rule state", Computed: true}, "expiry_time": schema.Int64Attribute{MarkdownDescription: "Expiry interval for suppression rule", Computed: true}, "expiry_time_unit": schema.StringAttribute{MarkdownDescription: "Expiry interval unit", Computed: true}, "expiry_ts": schema.StringAttribute{MarkdownDescription: "Expiry Timestamp(UTC) for suppression rule", Computed: true}, "hostname": schema.StringAttribute{MarkdownDescription: "FM Hostname", Computed: true}, "resource_id": schema.StringAttribute{MarkdownDescription: "Id of the resource to be suppressed", Computed: true}, "resource_type": schema.StringAttribute{Computed: true}, "selected_reason": schema.StringAttribute{MarkdownDescription: "Suppression Reason", Computed: true}, "tags": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"tag_key": schema.StringAttribute{MarkdownDescription: "Name of the tag", Computed: true}, "tag_values": schema.ListAttribute{MarkdownDescription: "All possible values of the tag", Computed: true, ElementType: types.StringType}}}}, "updated_by": schema.StringAttribute{MarkdownDescription: "Alarm suppression rule updated by FM User", Computed: true}, "updated_ts": schema.StringAttribute{MarkdownDescription: "Alarm suppression rule updated timestamp in ISO 8601 format", Computed: true}}}}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned. Defaults to (1:30) to prevent reading entire DB by accident", Optional: true}, "resource_id": schema.StringAttribute{MarkdownDescription: "Suppressed resource Id", Optional: true}, "selected_reason": schema.StringAttribute{MarkdownDescription: "Suppression Reason", Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is DESC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetAllAlarmSuppressionMetadataDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetAllAlarmSuppressionMetadataDataSourceModel
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
func (d *GetAllAlarmSuppressionMetadataDataSource) readListRemote(ctx context.Context, config *GetAllAlarmSuppressionMetadataDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/alarms/suppression"
	params := url.Values{}
	if !config.Page.IsNull() {
		params.Set("page", config.Page.ValueString())
	}
	if !config.Sort.IsNull() {
		params.Set("sort", config.Sort.ValueString())
	}
	if !config.AlarmTypes.IsNull() {
		params.Set("alarmTypes", config.AlarmTypes.ValueString())
	}
	if !config.ResourceId.IsNull() {
		params.Set("resourceId", config.ResourceId.ValueString())
	}
	if !config.SelectedReason.IsNull() {
		params.Set("selectedReason", config.SelectedReason.ValueString())
	}
	if !config.ClusterId.IsNull() {
		params.Set("clusterId", config.ClusterId.ValueString())
	}
	if !config.Alias.IsNull() {
		params.Set("alias", config.Alias.ValueString())
	}
	if !config.Hostname.IsNull() {
		params.Set("hostname", config.Hostname.ValueString())
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_alarm_suppression_metadata", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_alarm_suppression_metadata", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["suppressedEntities"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_alarm_suppression_metadata", fmt.Sprintf("Could not decode list page: missing %q array", "suppressedEntities"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_alarm_suppression_metadata", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetAllAlarmSuppressionMetadataDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
