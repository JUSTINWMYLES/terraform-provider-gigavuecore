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
	_ datasource.DataSource              = (*GetAppInfoDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetAppInfoDataSource)(nil)
)

// GetAppInfoDataSource is the generated Terraform data source implementation.
type GetAppInfoDataSource struct {
	client *client.Client
}

// GetAppInfoDataSourceModel describes the data source state shape.
type GetAppInfoDataSourceModel struct {
	Appinfo      types.String `tfsdk:"appinfo"`
	Apps         types.Object `tfsdk:"apps"`
	DpType       types.String `tfsdk:"dp_type" json:"dpType"`
	EnvId        types.String `tfsdk:"env_id" json:"envId"`
	EventDev     types.String `tfsdk:"event_dev" json:"eventDev"`
	MajorVersion types.Int64  `tfsdk:"major_version" json:"majorVersion"`
	MinorVersion types.Int64  `tfsdk:"minor_version" json:"minorVersion"`
	Name         types.String `tfsdk:"name"`
	Path         types.String `tfsdk:"path"`
	PlatformType types.String `tfsdk:"platform_type" json:"platformType"`
	SerialNum    types.Int64  `tfsdk:"serial_num" json:"serialNum"`
	Type         types.String `tfsdk:"type"`
	UnifyId      types.String `tfsdk:"unify_id" json:"unifyId"`
	Vendor       types.String `tfsdk:"vendor"`
}

// NewGetAppInfoDataSource returns a new instance of the generated data source.
func NewGetAppInfoDataSource() datasource.DataSource {
	return &GetAppInfoDataSource{}
}

// Metadata returns the data source type name.
func (d *GetAppInfoDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_app_info"
}

// Schema returns the data source schema.
func (d *GetAppInfoDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get an appInfo for unified deployment via environment and connection id", Attributes: map[string]schema.Attribute{"appinfo": schema.StringAttribute{MarkdownDescription: "appinfo identifier", Required: true}, "apps": schema.SingleNestedAttribute{MarkdownDescription: "Unified Resource app info for specific app", Computed: true, Attributes: map[string]schema.Attribute{"max_nr_aeps": schema.StringAttribute{Computed: true}, "mem_footprints": schema.SingleNestedAttribute{MarkdownDescription: "app specific memory footprint info", Computed: true, Attributes: map[string]schema.Attribute{"base": schema.Int64Attribute{Computed: true}, "scale": schema.SingleNestedAttribute{MarkdownDescription: "app specific memory footprint scale info", Computed: true, Attributes: map[string]schema.Attribute{"item": schema.StringAttribute{Computed: true}, "memory": schema.Int64Attribute{Computed: true}, "shared": schema.BoolAttribute{Computed: true}}}}}, "name": schema.StringAttribute{Computed: true}, "packet_permission": schema.StringAttribute{Computed: true}}}, "dp_type": schema.StringAttribute{Computed: true}, "env_id": schema.StringAttribute{MarkdownDescription: "unified environment identifier", Required: true}, "event_dev": schema.StringAttribute{Computed: true}, "major_version": schema.Int64Attribute{Computed: true}, "minor_version": schema.Int64Attribute{Computed: true}, "name": schema.StringAttribute{Computed: true}, "path": schema.StringAttribute{Computed: true}, "platform_type": schema.StringAttribute{Computed: true}, "serial_num": schema.Int64Attribute{Computed: true}, "type": schema.StringAttribute{Computed: true}, "unify_id": schema.StringAttribute{MarkdownDescription: "unified resource identifier", Required: true}, "vendor": schema.StringAttribute{Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetAppInfoDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetAppInfoDataSourceModel
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
func (d *GetAppInfoDataSource) readRemote(ctx context.Context, config *GetAppInfoDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/unifiedResources/env/{envId}/unifyId/{unifyId}/appinfo/{appinfo}"
	reqPath = strings.ReplaceAll(reqPath, "{envId}", url.PathEscape(config.EnvId.ValueString()))
	reqPath = strings.ReplaceAll(reqPath, "{unifyId}", url.PathEscape(config.UnifyId.ValueString()))
	reqPath = strings.ReplaceAll(reqPath, "{appinfo}", url.PathEscape(config.Appinfo.ValueString()))
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_app_info", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_app_info", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_app_info", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_app_info", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_app_info", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_app_info", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_app_info", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_app_info", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_app_info", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_app_info", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_app_info", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_app_info", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_app_info", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_app_info", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetAppInfoDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
