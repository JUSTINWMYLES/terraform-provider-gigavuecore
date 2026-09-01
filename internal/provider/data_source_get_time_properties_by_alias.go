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
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*GetTimePropertiesByAliasDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetTimePropertiesByAliasDataSource)(nil)
)

// GetTimePropertiesByAliasDataSource is the generated Terraform data source implementation.
type GetTimePropertiesByAliasDataSource struct {
	client *client.Client
}

// GetTimePropertiesByAliasDataSourceModel describes the data source state shape.
type GetTimePropertiesByAliasDataSourceModel struct {
	Alias              types.String `tfsdk:"alias"`
	BoxId              types.Int64  `tfsdk:"box_id" json:"boxId"`
	CurrentUtcOffset   types.Int64  `tfsdk:"current_utc_offset" json:"currentUtcOffset"`
	FrequencyTraceable types.Int64  `tfsdk:"frequency_traceable" json:"frequencyTraceable"`
	Leap59             types.Int64  `tfsdk:"leap59"`
	Leap61             types.Int64  `tfsdk:"leap61"`
	PtpTimeScale       types.Int64  `tfsdk:"ptp_time_scale" json:"PtpTimeScale"`
	TimeSource         types.String `tfsdk:"time_source" json:"timeSource"`
	TimeTraceable      types.Int64  `tfsdk:"time_traceable" json:"timeTraceable"`
}

// NewGetTimePropertiesByAliasDataSource returns a new instance of the generated data source.
func NewGetTimePropertiesByAliasDataSource() datasource.DataSource {
	return &GetTimePropertiesByAliasDataSource{}
}

// Metadata returns the data source type name.
func (d *GetTimePropertiesByAliasDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_time_properties_by_alias"
}

// Schema returns the data source schema.
func (d *GetTimePropertiesByAliasDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Reads the get time properties by alias data source.", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "alias of the time stamping PTP configuration", Required: true}, "box_id": schema.Int64Attribute{Computed: true}, "current_utc_offset": schema.Int64Attribute{Computed: true}, "frequency_traceable": schema.Int64Attribute{Computed: true}, "leap59": schema.Int64Attribute{Computed: true}, "leap61": schema.Int64Attribute{Computed: true}, "ptp_time_scale": schema.Int64Attribute{Computed: true}, "time_source": schema.StringAttribute{Computed: true}, "time_traceable": schema.Int64Attribute{Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetTimePropertiesByAliasDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetTimePropertiesByAliasDataSourceModel
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
func (d *GetTimePropertiesByAliasDataSource) readRemote(ctx context.Context, config *GetTimePropertiesByAliasDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/ptp/timeProperty/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(config.Alias.ValueString()))
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_time_properties_by_alias", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_time_properties_by_alias", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_time_properties_by_alias", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_time_properties_by_alias", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_time_properties_by_alias", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_time_properties_by_alias", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_time_properties_by_alias", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_time_properties_by_alias", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_time_properties_by_alias", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_time_properties_by_alias", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_time_properties_by_alias", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_time_properties_by_alias", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_time_properties_by_alias", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["timeProperty"]; ok {
		if m, ok := v.(map[string]any); ok {
			data = m
		} else if arr, ok := v.([]any); ok && len(arr) > 0 {
			if m, ok := arr[0].(map[string]any); ok {
				data = m
			}
		}
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_time_properties_by_alias", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetTimePropertiesByAliasDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
