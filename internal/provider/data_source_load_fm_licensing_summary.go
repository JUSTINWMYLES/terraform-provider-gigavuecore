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
	_ datasource.DataSource              = (*LoadFmLicensingSummaryDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadFmLicensingSummaryDataSource)(nil)
)

// LoadFmLicensingSummaryDataSource is the generated Terraform data source implementation.
type LoadFmLicensingSummaryDataSource struct {
	client *client.Client
}

// LoadFmLicensingSummaryDataSourceModel describes the data source state shape.
type LoadFmLicensingSummaryDataSourceModel struct {
	Aws        types.Object `tfsdk:"aws"`
	Azure      types.Object `tfsdk:"azure"`
	BaseBundle types.Object `tfsdk:"base_bundle" json:"baseBundle"`
	Features   types.Object `tfsdk:"features"`
	OpenStack  types.Object `tfsdk:"open_stack" json:"openStack"`
	TargetFmId types.String `tfsdk:"target_fm_id" json:"targetFmId"`
	Vmm        types.Object `tfsdk:"vmm"`
}

// NewLoadFmLicensingSummaryDataSource returns a new instance of the generated data source.
func NewLoadFmLicensingSummaryDataSource() datasource.DataSource {
	return &LoadFmLicensingSummaryDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadFmLicensingSummaryDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_fm_licensing_summary"
}

// Schema returns the data source schema.
func (d *LoadFmLicensingSummaryDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load FM Licensing Summary", Attributes: map[string]schema.Attribute{"aws": schema.SingleNestedAttribute{MarkdownDescription: "AWS License summary", Computed: true, Attributes: map[string]schema.Attribute{"licensed": schema.BoolAttribute{Computed: true}, "v_nics": schema.Int64Attribute{MarkdownDescription: "number of licensed vNics", Computed: true}}}, "azure": schema.SingleNestedAttribute{MarkdownDescription: "Azure License summary", Computed: true, Attributes: map[string]schema.Attribute{"licensed": schema.BoolAttribute{Computed: true}, "v_nics": schema.Int64Attribute{MarkdownDescription: "number of licensed vNics", Computed: true}}}, "base_bundle": schema.SingleNestedAttribute{MarkdownDescription: "Gigamon FM License Base Bundle summary", Computed: true, Attributes: map[string]schema.Attribute{"nodes": schema.Int64Attribute{MarkdownDescription: "number of licensed physical nodes", Computed: true}, "type": schema.StringAttribute{Computed: true}}}, "features": schema.SingleNestedAttribute{MarkdownDescription: "Gigamon FM License Features", Computed: true, Attributes: map[string]schema.Attribute{"dashboard_customization": schema.SingleNestedAttribute{MarkdownDescription: "Gigamon FM License Dashboard Customization Feature summary", Computed: true, Attributes: map[string]schema.Attribute{"type": schema.StringAttribute{Computed: true}}}, "reporting": schema.SingleNestedAttribute{MarkdownDescription: "Gigamon FM License Reporting Feature summary", Computed: true, Attributes: map[string]schema.Attribute{"type": schema.StringAttribute{Computed: true}}}, "traffic_analyzer": schema.SingleNestedAttribute{MarkdownDescription: "Gigamon FM License Traffic Analyzer Feature summary", Computed: true, Attributes: map[string]schema.Attribute{"type": schema.StringAttribute{Computed: true}}}, "trending": schema.SingleNestedAttribute{MarkdownDescription: "Gigamon FM License Trending Feature summary", Computed: true, Attributes: map[string]schema.Attribute{"historical_days": schema.Int64Attribute{Computed: true}, "type": schema.StringAttribute{Computed: true}}}}}, "open_stack": schema.SingleNestedAttribute{MarkdownDescription: "OpenStack License summary", Computed: true, Attributes: map[string]schema.Attribute{"licensed": schema.BoolAttribute{Computed: true}, "v_nics": schema.Int64Attribute{MarkdownDescription: "number of licensed vNics", Computed: true}}}, "target_fm_id": schema.StringAttribute{Computed: true}, "vmm": schema.SingleNestedAttribute{MarkdownDescription: "VMM License summary", Computed: true, Attributes: map[string]schema.Attribute{"historical_days": schema.Int64Attribute{Computed: true}, "licensed": schema.BoolAttribute{Computed: true}, "nodes": schema.Int64Attribute{MarkdownDescription: "number of licensed GVM nodes", Computed: true}}}}}
}

// Read fetches remote state into the data source model.
func (d *LoadFmLicensingSummaryDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadFmLicensingSummaryDataSourceModel
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
func (d *LoadFmLicensingSummaryDataSource) readRemote(ctx context.Context, config *LoadFmLicensingSummaryDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/licensing/fm"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_fm_licensing_summary", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_fm_licensing_summary", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_fm_licensing_summary", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_fm_licensing_summary", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_fm_licensing_summary", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_fm_licensing_summary", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_fm_licensing_summary", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_fm_licensing_summary", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_fm_licensing_summary", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_fm_licensing_summary", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_load_fm_licensing_summary", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_load_fm_licensing_summary", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_fm_licensing_summary", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["licensingSummary"]; ok {
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
		resp.Diagnostics.AddError("Error reading gigavuecore_load_fm_licensing_summary", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadFmLicensingSummaryDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
