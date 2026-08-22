package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)
import (
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
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
	AlarmTypes         types.String `tfsdk:"alarm_types" json:"alarmTypes"`
	Alias              types.String `tfsdk:"alias"`
	ClusterId          types.String `tfsdk:"cluster_id" json:"clusterId"`
	Context            types.Object `tfsdk:"context"`
	Hostname           types.String `tfsdk:"hostname"`
	Page               types.String `tfsdk:"page"`
	ResourceId         types.String `tfsdk:"resource_id" json:"resourceId"`
	SelectedReason     types.String `tfsdk:"selected_reason" json:"selectedReason"`
	Sort               types.String `tfsdk:"sort"`
	SuppressedEntities types.List   `tfsdk:"suppressed_entities" json:"suppressedEntities"`
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
	resp.Schema = schema.Schema{MarkdownDescription: "Get All Alarm Supppression Rules", Attributes: map[string]schema.Attribute{"alarm_types": schema.StringAttribute{MarkdownDescription: "comma-separated list of Alarm Type's to filter", Optional: true}, "alias": schema.StringAttribute{Optional: true}, "cluster_id": schema.StringAttribute{Optional: true}, "context": schema.SingleNestedAttribute{MarkdownDescription: "Gigamon query result context", Computed: true, Attributes: map[string]schema.Attribute{"page_no": schema.Int64Attribute{MarkdownDescription: "page number of the returned result set", Computed: true}, "page_size": schema.Int64Attribute{MarkdownDescription: "page size of the returned result set", Computed: true}, "sort": schema.ListAttribute{MarkdownDescription: "sorting info of the returned result set. list of fields in the array indicate sorting order", Computed: true, ElementType: types.StringType}, "total_items": schema.Int64Attribute{MarkdownDescription: "total number of items in the queried entity type", Computed: true}}}, "hostname": schema.StringAttribute{Optional: true}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned. Defaults to (1:30) to prevent reading entire DB by accident", Optional: true}, "resource_id": schema.StringAttribute{Optional: true}, "selected_reason": schema.StringAttribute{Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is DESC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}, "suppressed_entities": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alarm_types": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "alias": schema.StringAttribute{MarkdownDescription: "Alias", Computed: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Cluster/Node Id", Computed: true}, "created_by": schema.StringAttribute{MarkdownDescription: "Alarm suppression rule created by FM User", Computed: true}, "created_ts": schema.StringAttribute{MarkdownDescription: "Alarm suppression rule creation timestamp in ISO 8601 format", Computed: true}, "enable": schema.BoolAttribute{MarkdownDescription: "FM Alarm Suppression rule state", Computed: true}, "expiry_time": schema.Int64Attribute{MarkdownDescription: "Expiry interval for suppression rule", Computed: true}, "expiry_time_unit": schema.StringAttribute{MarkdownDescription: "Expiry interval unit", Computed: true}, "expiry_ts": schema.StringAttribute{MarkdownDescription: "Expiry Timestamp(UTC) for suppression rule", Computed: true}, "hostname": schema.StringAttribute{MarkdownDescription: "FM Hostname", Computed: true}, "resource_id": schema.StringAttribute{MarkdownDescription: "Id of the resource to be suppressed", Computed: true}, "resource_type": schema.StringAttribute{Computed: true}, "selected_reason": schema.StringAttribute{MarkdownDescription: "Suppression Reason", Computed: true}, "tags": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"tag_key": schema.StringAttribute{MarkdownDescription: "Name of the tag", Computed: true}, "tag_values": schema.ListAttribute{MarkdownDescription: "All possible values of the tag", Computed: true, ElementType: types.StringType}}}}, "updated_by": schema.StringAttribute{MarkdownDescription: "Alarm suppression rule updated by FM User", Computed: true}, "updated_ts": schema.StringAttribute{MarkdownDescription: "Alarm suppression rule updated timestamp in ISO 8601 format", Computed: true}}}}}}
}

// Read fetches remote state into the data source model.
func (d *GetAllAlarmSuppressionMetadataDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetAllAlarmSuppressionMetadataDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	d.readRemote(ctx, &config, resp)
	if resp.Diagnostics.HasError() {
		return
	}
	if config.Page.IsNull() {
		resp.Diagnostics.AddWarning("Single-page result", fmt.Sprintf("This data source reads a single page of a paginated API endpoint and does not aggregate results across pages. The \"page\" argument is unset, so the default page is returned; set it to retrieve a different page."))
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}

// readRemote performs the read HTTP exchange and decodes the response into config. Extracted from Read so the request/response logic is unit-testable without a tfsdk.Config.
func (d *GetAllAlarmSuppressionMetadataDataSource) readRemote(ctx context.Context, config *GetAllAlarmSuppressionMetadataDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/alarms/suppression"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_alarm_suppression_metadata", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !config.Page.IsNull() {
		query.Set("page", config.Page.ValueString())
	}
	if !config.Sort.IsNull() {
		query.Set("sort", config.Sort.ValueString())
	}
	if !config.AlarmTypes.IsNull() {
		query.Set("alarmTypes", config.AlarmTypes.ValueString())
	}
	if !config.ResourceId.IsNull() {
		query.Set("resourceId", config.ResourceId.ValueString())
	}
	if !config.SelectedReason.IsNull() {
		query.Set("selectedReason", config.SelectedReason.ValueString())
	}
	if !config.ClusterId.IsNull() {
		query.Set("clusterId", config.ClusterId.ValueString())
	}
	if !config.Alias.IsNull() {
		query.Set("alias", config.Alias.ValueString())
	}
	if !config.Hostname.IsNull() {
		query.Set("hostname", config.Hostname.ValueString())
	}
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_alarm_suppression_metadata", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_alarm_suppression_metadata", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_alarm_suppression_metadata", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_alarm_suppression_metadata", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_alarm_suppression_metadata", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_alarm_suppression_metadata", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_alarm_suppression_metadata", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_alarm_suppression_metadata", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_alarm_suppression_metadata", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_all_alarm_suppression_metadata", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_alarm_suppression_metadata", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_alarm_suppression_metadata", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
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
