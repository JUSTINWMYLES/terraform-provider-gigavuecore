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
	_ datasource.DataSource              = (*GetAllAlarmAutoSuppressionRulesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetAllAlarmAutoSuppressionRulesDataSource)(nil)
)

// GetAllAlarmAutoSuppressionRulesDataSource is the generated Terraform data source implementation.
type GetAllAlarmAutoSuppressionRulesDataSource struct {
	client *client.Client
}

// GetAllAlarmAutoSuppressionRulesDataSourceModel describes the data source state shape.
type GetAllAlarmAutoSuppressionRulesDataSourceModel struct {
	AutoSuppressionRules types.List `tfsdk:"auto_suppression_rules" json:"autoSuppressionRules"`
	Enable               types.Bool `tfsdk:"enable"`
}

// NewGetAllAlarmAutoSuppressionRulesDataSource returns a new instance of the generated data source.
func NewGetAllAlarmAutoSuppressionRulesDataSource() datasource.DataSource {
	return &GetAllAlarmAutoSuppressionRulesDataSource{}
}

// Metadata returns the data source type name.
func (d *GetAllAlarmAutoSuppressionRulesDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_all_alarm_auto_suppression_rules"
}

// Schema returns the data source schema.
func (d *GetAllAlarmAutoSuppressionRulesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get All Alarm Auto Supppression Rules", Attributes: map[string]schema.Attribute{"auto_suppression_rules": schema.ListNestedAttribute{MarkdownDescription: "Alarm Auto Suppression Rules", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"enable": schema.BoolAttribute{Computed: true}, "type": schema.StringAttribute{MarkdownDescription: "Type of operation on Cluster/Node", Computed: true}}}}, "enable": schema.BoolAttribute{MarkdownDescription: "Enable/Disable auto alarm suppression rules for operation Image Upgrade,Config Restore,Device Reboot and Cluster Operation", Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetAllAlarmAutoSuppressionRulesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetAllAlarmAutoSuppressionRulesDataSourceModel
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
func (d *GetAllAlarmAutoSuppressionRulesDataSource) readRemote(ctx context.Context, config *GetAllAlarmAutoSuppressionRulesDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/alarms/autosuppression"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_alarm_auto_suppression_rules", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_alarm_auto_suppression_rules", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_alarm_auto_suppression_rules", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_alarm_auto_suppression_rules", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_alarm_auto_suppression_rules", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_alarm_auto_suppression_rules", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_alarm_auto_suppression_rules", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_alarm_auto_suppression_rules", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_alarm_auto_suppression_rules", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_alarm_auto_suppression_rules", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_all_alarm_auto_suppression_rules", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_alarm_auto_suppression_rules", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_alarm_auto_suppression_rules", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_alarm_auto_suppression_rules", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetAllAlarmAutoSuppressionRulesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
