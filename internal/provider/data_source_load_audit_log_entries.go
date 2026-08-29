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
	EndTime   types.String  `tfsdk:"end_time" json:"endTime"`
	Items     types.Dynamic `tfsdk:"items"`
	Operation types.String  `tfsdk:"operation"`
	Outcome   types.String  `tfsdk:"outcome"`
	Page      types.String  `tfsdk:"page"`
	Sort      types.String  `tfsdk:"sort"`
	StartTime types.String  `tfsdk:"start_time" json:"startTime"`
	Target    types.String  `tfsdk:"target"`
	Username  types.String  `tfsdk:"username"`
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
	resp.Schema = schema.Schema{MarkdownDescription: "Load User Action Audit Log Entries", Attributes: map[string]schema.Attribute{"end_time": schema.StringAttribute{MarkdownDescription: "End Time to filter by. In ISO 8601 format", Optional: true}, "items": schema.DynamicAttribute{Computed: true}, "operation": schema.StringAttribute{MarkdownDescription: "User Action type identifier to filter by", Optional: true}, "outcome": schema.StringAttribute{MarkdownDescription: "User Action outcome type to filter by", Optional: true}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned. Defaults to (1:100) to prevent reading entire DB by accident", Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}, "start_time": schema.StringAttribute{MarkdownDescription: "Start Time to filter by. In ISO 8601 format", Optional: true}, "target": schema.StringAttribute{MarkdownDescription: "Action Target identifier to filter by", Optional: true}, "username": schema.StringAttribute{MarkdownDescription: "Username to filter by", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadAuditLogEntriesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadAuditLogEntriesDataSourceModel
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
func (d *LoadAuditLogEntriesDataSource) readListRemote(ctx context.Context, config *LoadAuditLogEntriesDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/auditLog"
	params := url.Values{}
	if !config.Username.IsNull() {
		params.Set("username", config.Username.ValueString())
	}
	if !config.Target.IsNull() {
		params.Set("target", config.Target.ValueString())
	}
	if !config.Operation.IsNull() {
		params.Set("operation", config.Operation.ValueString())
	}
	if !config.Outcome.IsNull() {
		params.Set("outcome", config.Outcome.ValueString())
	}
	if !config.StartTime.IsNull() {
		params.Set("startTime", config.StartTime.ValueString())
	}
	if !config.EndTime.IsNull() {
		params.Set("endTime", config.EndTime.ValueString())
	}
	if !config.Page.IsNull() {
		params.Set("page", config.Page.ValueString())
	}
	if !config.Sort.IsNull() {
		params.Set("sort", config.Sort.ValueString())
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
		resp.Diagnostics.AddError("Error reading gigavuecore_load_audit_log_entries", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_audit_log_entries", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["auditLogEntries"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_audit_log_entries", fmt.Sprintf("Could not decode list page: missing %q array", "auditLogEntries"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
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
