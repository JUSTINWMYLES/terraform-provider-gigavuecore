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
	_ datasource.DataSource              = (*LoadEventsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadEventsDataSource)(nil)
)

// LoadEventsDataSource is the generated Terraform data source implementation.
type LoadEventsDataSource struct {
	client *client.Client
}

// LoadEventsDataSourceModel describes the data source state shape.
type LoadEventsDataSourceModel struct {
	Context      types.Object `tfsdk:"context"`
	EndTime      types.String `tfsdk:"end_time" json:"endTime"`
	Events       types.List   `tfsdk:"events"`
	Page         types.String `tfsdk:"page"`
	ResourceId   types.String `tfsdk:"resource_id" json:"resourceId"`
	ResourceType types.String `tfsdk:"resource_type" json:"resourceType"`
	Scope        types.String `tfsdk:"scope"`
	Severity     types.String `tfsdk:"severity"`
	Sort         types.String `tfsdk:"sort"`
	Source       types.String `tfsdk:"source"`
	StartTime    types.String `tfsdk:"start_time" json:"startTime"`
	Type         types.String `tfsdk:"type"`
}

// NewLoadEventsDataSource returns a new instance of the generated data source.
func NewLoadEventsDataSource() datasource.DataSource {
	return &LoadEventsDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadEventsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_events"
}

// Schema returns the data source schema.
func (d *LoadEventsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load Events", Attributes: map[string]schema.Attribute{"context": schema.SingleNestedAttribute{MarkdownDescription: "Gigamon query result context", Computed: true, Attributes: map[string]schema.Attribute{"page_no": schema.Int64Attribute{MarkdownDescription: "page number of the returned result set", Computed: true}, "page_size": schema.Int64Attribute{MarkdownDescription: "page size of the returned result set", Computed: true}, "sort": schema.ListAttribute{MarkdownDescription: "sorting info of the returned result set. list of fields in the array indicate sorting order", Computed: true, ElementType: types.StringType}, "total_items": schema.Int64Attribute{MarkdownDescription: "total number of items in the queried entity type", Computed: true}}}, "end_time": schema.StringAttribute{MarkdownDescription: "End Time to filter by. In ISO 8601 format", Optional: true}, "events": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"description": schema.StringAttribute{MarkdownDescription: "Event description", Computed: true}, "resource_id": schema.StringAttribute{MarkdownDescription: "Affected Entity of the event", Computed: true}, "resource_type": schema.StringAttribute{MarkdownDescription: "Affected Entity Type of the event", Computed: true}, "scope": schema.StringAttribute{MarkdownDescription: "Scope of the event", Computed: true}, "severity": schema.StringAttribute{MarkdownDescription: "Severity of the event", Computed: true}, "source": schema.StringAttribute{MarkdownDescription: "Event Source. Device ID for node-originated events like Traps or Syslogs. Component ID for FM-originated events", Computed: true}, "ts": schema.StringAttribute{MarkdownDescription: "Event timestamp in ISO 8601 format", Computed: true}, "ts_utc": schema.Int64Attribute{MarkdownDescription: "Timestamp in UTC milliseconds", Computed: true}, "type": schema.StringAttribute{MarkdownDescription: "Event Type identifier", Computed: true}}}}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned. Defaults to (1:100) to prevent reading entire DB by accident", Optional: true}, "resource_id": schema.StringAttribute{MarkdownDescription: "Affected Entity to filter by", Optional: true}, "resource_type": schema.StringAttribute{MarkdownDescription: "Affected Entity Type to filter by", Optional: true}, "scope": schema.StringAttribute{MarkdownDescription: "Event Scope to filter by", Optional: true}, "severity": schema.StringAttribute{MarkdownDescription: "Event Scope to filter by", Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}, "source": schema.StringAttribute{MarkdownDescription: "Event Source identifier to filter by", Optional: true}, "start_time": schema.StringAttribute{MarkdownDescription: "Start Time to filter by. In ISO 8601 format", Optional: true}, "type": schema.StringAttribute{MarkdownDescription: "Event Type to filter by", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadEventsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadEventsDataSourceModel
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
func (d *LoadEventsDataSource) readRemote(ctx context.Context, config *LoadEventsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/events"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_events", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !config.Type.IsNull() {
		query.Set("type", config.Type.ValueString())
	}
	if !config.Source.IsNull() {
		query.Set("source", config.Source.ValueString())
	}
	if !config.Scope.IsNull() {
		query.Set("scope", config.Scope.ValueString())
	}
	if !config.Severity.IsNull() {
		query.Set("severity", config.Severity.ValueString())
	}
	if !config.StartTime.IsNull() {
		query.Set("startTime", config.StartTime.ValueString())
	}
	if !config.EndTime.IsNull() {
		query.Set("endTime", config.EndTime.ValueString())
	}
	if !config.ResourceType.IsNull() {
		query.Set("resourceType", config.ResourceType.ValueString())
	}
	if !config.ResourceId.IsNull() {
		query.Set("resourceId", config.ResourceId.ValueString())
	}
	if !config.Page.IsNull() {
		query.Set("page", config.Page.ValueString())
	}
	if !config.Sort.IsNull() {
		query.Set("sort", config.Sort.ValueString())
	}
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_events", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_events", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_events", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_events", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_events", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_events", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_events", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_events", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_events", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_load_events", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_load_events", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_events", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_events", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadEventsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
