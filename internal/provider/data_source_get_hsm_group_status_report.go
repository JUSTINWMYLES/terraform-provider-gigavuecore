package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)
import (
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*GetHsmGroupStatusReportDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetHsmGroupStatusReportDataSource)(nil)
)

// GetHsmGroupStatusReportDataSource is the generated Terraform data source implementation.
type GetHsmGroupStatusReportDataSource struct {
	client *client.Client
}

// GetHsmGroupStatusReportDataSourceModel describes the data source state shape.
type GetHsmGroupStatusReportDataSourceModel struct {
	Alias        types.String `tfsdk:"alias"`
	Report       types.String `tfsdk:"report"`
	SessionStats types.List   `tfsdk:"session_stats" json:"sessionStats"`
	Type         types.String `tfsdk:"type"`
}

// NewGetHsmGroupStatusReportDataSource returns a new instance of the generated data source.
func NewGetHsmGroupStatusReportDataSource() datasource.DataSource {
	return &GetHsmGroupStatusReportDataSource{}
}

// Metadata returns the data source type name.
func (d *GetHsmGroupStatusReportDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_hsm_group_status_report"
}

// Schema returns the data source schema.
func (d *GetHsmGroupStatusReportDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load HSM Group status report", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "alias of target HSM Group", Optional: true}, "report": schema.StringAttribute{MarkdownDescription: "status report", Computed: true}, "session_stats": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"engine_id": schema.StringAttribute{Computed: true}, "hsm_avg_rtt": schema.Int64Attribute{Computed: true}, "hsm_late": schema.Int64Attribute{Computed: true}, "hsm_req_sent": schema.Int64Attribute{Computed: true}, "hsm_resp_rcvd": schema.Int64Attribute{Computed: true}, "hsmc_avg_pkcs11_rtt": schema.Int64Attribute{Computed: true}, "hsmc_avg_rtt": schema.Int64Attribute{Computed: true}, "hsmc_avgq_delay": schema.Int64Attribute{Computed: true}, "hsmc_err_pkcs11": schema.Int64Attribute{Computed: true}, "hsmc_err_send_resp": schema.Int64Attribute{Computed: true}, "hsmc_errq_full": schema.Int64Attribute{Computed: true}, "hsmc_max_pkcs11_rtt": schema.Int64Attribute{Computed: true}, "hsmc_max_rtt": schema.Int64Attribute{Computed: true}, "hsmc_maxq_delay": schema.Int64Attribute{Computed: true}, "hsmc_min_pkcs11_rtt": schema.Int64Attribute{Computed: true}, "hsmc_min_rtt": schema.Int64Attribute{Computed: true}, "hsmc_minq_delay": schema.Int64Attribute{Computed: true}, "hsmc_req_rcvd": schema.Int64Attribute{Computed: true}, "hsmc_resp_sent": schema.Int64Attribute{Computed: true}}}}, "type": schema.StringAttribute{MarkdownDescription: "report type, valid only if alias is specified", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetHsmGroupStatusReportDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetHsmGroupStatusReportDataSourceModel
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
func (d *GetHsmGroupStatusReportDataSource) readRemote(ctx context.Context, config *GetHsmGroupStatusReportDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/apps/hsmGroupStatusReport"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_hsm_group_status_report", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !config.Type.IsNull() {
		query.Set("type", config.Type.ValueString())
	}
	if !config.Alias.IsNull() {
		query.Set("alias", config.Alias.ValueString())
	}
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_hsm_group_status_report", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_hsm_group_status_report", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_hsm_group_status_report", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_hsm_group_status_report", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_hsm_group_status_report", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_hsm_group_status_report", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_hsm_group_status_report", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_hsm_group_status_report", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_hsm_group_status_report", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_hsm_group_status_report", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_hsm_group_status_report", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_hsm_group_status_report", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_hsm_group_status_report", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetHsmGroupStatusReportDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
