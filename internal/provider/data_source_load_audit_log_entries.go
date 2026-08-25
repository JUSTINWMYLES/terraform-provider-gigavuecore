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
	_ datasource.DataSource              = (*LoadAuditLogEntriesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadAuditLogEntriesDataSource)(nil)
)

// LoadAuditLogEntriesDataSource is the generated Terraform data source implementation.
type LoadAuditLogEntriesDataSource struct {
	client *client.Client
}

// LoadAuditLogEntriesDataSourceModel describes the data source state shape.
type LoadAuditLogEntriesDataSourceModel struct {
	AuditLogEntries types.Dynamic `tfsdk:"audit_log_entries" json:"auditLogEntries"`
	Context         types.Object  `tfsdk:"context"`
	EndTime         types.String  `tfsdk:"end_time" json:"endTime"`
	Operation       types.String  `tfsdk:"operation"`
	Outcome         types.String  `tfsdk:"outcome"`
	Page            types.String  `tfsdk:"page"`
	Sort            types.String  `tfsdk:"sort"`
	StartTime       types.String  `tfsdk:"start_time" json:"startTime"`
	Target          types.String  `tfsdk:"target"`
	Username        types.String  `tfsdk:"username"`
}

// NewLoadAuditLogEntriesDataSource returns a new instance of the generated data source.
func NewLoadAuditLogEntriesDataSource() datasource.DataSource {
	return &LoadAuditLogEntriesDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadAuditLogEntriesDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_audit_log_entries"
}

// Schema returns the data source schema.
func (d *LoadAuditLogEntriesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load User Action Audit Log Entries", Attributes: map[string]schema.Attribute{"audit_log_entries": schema.DynamicAttribute{Computed: true}, "context": schema.SingleNestedAttribute{MarkdownDescription: "Gigamon query result context", Computed: true, Attributes: map[string]schema.Attribute{"page_no": schema.Int64Attribute{MarkdownDescription: "page number of the returned result set", Computed: true}, "page_size": schema.Int64Attribute{MarkdownDescription: "page size of the returned result set", Computed: true}, "sort": schema.ListAttribute{MarkdownDescription: "sorting info of the returned result set. list of fields in the array indicate sorting order", Computed: true, ElementType: types.StringType}, "total_items": schema.Int64Attribute{MarkdownDescription: "total number of items in the queried entity type", Computed: true}}}, "end_time": schema.StringAttribute{MarkdownDescription: "End Time to filter by. In ISO 8601 format", Optional: true}, "operation": schema.StringAttribute{MarkdownDescription: "User Action type identifier to filter by", Optional: true}, "outcome": schema.StringAttribute{MarkdownDescription: "User Action outcome type to filter by", Optional: true}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned. Defaults to (1:100) to prevent reading entire DB by accident", Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}, "start_time": schema.StringAttribute{MarkdownDescription: "Start Time to filter by. In ISO 8601 format", Optional: true}, "target": schema.StringAttribute{MarkdownDescription: "Action Target identifier to filter by", Optional: true}, "username": schema.StringAttribute{MarkdownDescription: "Username to filter by", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadAuditLogEntriesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadAuditLogEntriesDataSourceModel
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
func (d *LoadAuditLogEntriesDataSource) readRemote(ctx context.Context, config *LoadAuditLogEntriesDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/auditLog"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_audit_log_entries", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !config.Username.IsNull() {
		query.Set("username", config.Username.ValueString())
	}
	if !config.Target.IsNull() {
		query.Set("target", config.Target.ValueString())
	}
	if !config.Operation.IsNull() {
		query.Set("operation", config.Operation.ValueString())
	}
	if !config.Outcome.IsNull() {
		query.Set("outcome", config.Outcome.ValueString())
	}
	if !config.StartTime.IsNull() {
		query.Set("startTime", config.StartTime.ValueString())
	}
	if !config.EndTime.IsNull() {
		query.Set("endTime", config.EndTime.ValueString())
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
		resp.Diagnostics.AddError("Error reading gigavuecore_load_audit_log_entries", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_audit_log_entries", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_audit_log_entries", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_audit_log_entries", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_audit_log_entries", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_audit_log_entries", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_audit_log_entries", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_audit_log_entries", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_audit_log_entries", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_load_audit_log_entries", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_load_audit_log_entries", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_audit_log_entries", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_audit_log_entries", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadAuditLogEntriesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
