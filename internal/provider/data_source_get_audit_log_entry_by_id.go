package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)
import (
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*GetAuditLogEntryByIdDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetAuditLogEntryByIdDataSource)(nil)
)

// GetAuditLogEntryByIdDataSource is the generated Terraform data source implementation.
type GetAuditLogEntryByIdDataSource struct {
	client *client.Client
}

// GetAuditLogEntryByIdDataSourceModel describes the data source state shape.
type GetAuditLogEntryByIdDataSourceModel struct {
	Description   types.String  `tfsdk:"description"`
	EntryId       types.String  `tfsdk:"entry_id" json:"entryId"`
	Operation     types.String  `tfsdk:"operation"`
	Outcome       types.String  `tfsdk:"outcome"`
	Target        types.String  `tfsdk:"target"`
	Ts            types.String  `tfsdk:"ts"`
	TsUtc         types.Int64   `tfsdk:"ts_utc" json:"tsUtc"`
	UpdateDetails types.Dynamic `tfsdk:"update_details" json:"updateDetails"`
	Username      types.String  `tfsdk:"username"`
}

// NewGetAuditLogEntryByIdDataSource returns a new instance of the generated data source.
func NewGetAuditLogEntryByIdDataSource() datasource.DataSource {
	return &GetAuditLogEntryByIdDataSource{}
}

// Metadata returns the data source type name.
func (d *GetAuditLogEntryByIdDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_audit_log_entry_by_id"
}

// Schema returns the data source schema.
func (d *GetAuditLogEntryByIdDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Find Audit Log Entry by ID", Attributes: map[string]schema.Attribute{"description": schema.StringAttribute{MarkdownDescription: "Event description. In case of failure, holds error description", Computed: true}, "entry_id": schema.StringAttribute{MarkdownDescription: "ID of the target Audit Log Entry", Required: true}, "operation": schema.StringAttribute{MarkdownDescription: "Description of the user action", Computed: true}, "outcome": schema.StringAttribute{MarkdownDescription: "Scope of the event", Computed: true}, "target": schema.StringAttribute{MarkdownDescription: "Action target. Device ID for node configuration actions. Component ID for FM-based service configuration", Computed: true}, "ts": schema.StringAttribute{MarkdownDescription: "Event timestamp in ISO 8601 format", Computed: true}, "ts_utc": schema.Int64Attribute{MarkdownDescription: "Timestamp in UTC milliseconds", Computed: true}, "update_details": schema.DynamicAttribute{MarkdownDescription: "Optional extension to hold the update delta details", Computed: true}, "username": schema.StringAttribute{MarkdownDescription: "Action target. Device ID for node configuration actions. Component ID for FM-based service configuration", Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetAuditLogEntryByIdDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetAuditLogEntryByIdDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	d.readRemote(ctx, &config, resp)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}

// readRemote performs the read HTTP exchange and decodes the response into config. Extracted from Read so the request/response logic is unit-testable without a tfsdk.Config.
func (d *GetAuditLogEntryByIdDataSource) readRemote(ctx context.Context, config *GetAuditLogEntryByIdDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/auditLog/{entryId}"
	reqPath = strings.ReplaceAll(reqPath, "{entryId}", url.PathEscape(config.EntryId.ValueString()))
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_audit_log_entry_by_id", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_audit_log_entry_by_id", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_audit_log_entry_by_id", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_audit_log_entry_by_id", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_audit_log_entry_by_id", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_audit_log_entry_by_id", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_audit_log_entry_by_id", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_audit_log_entry_by_id", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_audit_log_entry_by_id", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_audit_log_entry_by_id", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_audit_log_entry_by_id", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_audit_log_entry_by_id", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_audit_log_entry_by_id", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["auditLofEntry"]; ok {
		if m, ok := v.(map[string]any); ok {
			data = m
		}
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_audit_log_entry_by_id", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetAuditLogEntryByIdDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
