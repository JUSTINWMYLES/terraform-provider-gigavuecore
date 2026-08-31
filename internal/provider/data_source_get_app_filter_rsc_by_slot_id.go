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
	_ datasource.DataSource              = (*GetAppFilterRscBySlotIdDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetAppFilterRscBySlotIdDataSource)(nil)
)

// GetAppFilterRscBySlotIdDataSource is the generated Terraform data source implementation.
type GetAppFilterRscBySlotIdDataSource struct {
	client *client.Client
}

// GetAppFilterRscBySlotIdDataSourceModel describes the data source state shape.
type GetAppFilterRscBySlotIdDataSourceModel struct {
	AppFilterResources types.List   `tfsdk:"app_filter_resources" json:"appFilterResources"`
	ClusterId          types.String `tfsdk:"cluster_id" json:"clusterId"`
	SlotId             types.String `tfsdk:"slot_id" json:"slotId"`
}

// NewGetAppFilterRscBySlotIdDataSource returns a new instance of the generated data source.
func NewGetAppFilterRscBySlotIdDataSource() datasource.DataSource {
	return &GetAppFilterRscBySlotIdDataSource{}
}

// Metadata returns the data source type name.
func (d *GetAppFilterRscBySlotIdDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_app_filter_rsc_by_slot_id"
}

// Schema returns the data source schema.
func (d *GetAppFilterRscBySlotIdDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load Application Filter Resource by slotId", Attributes: map[string]schema.Attribute{"app_filter_resources": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"app": schema.StringAttribute{MarkdownDescription: "vfp client application", Computed: true}, "free_pool_used": schema.Int64Attribute{MarkdownDescription: "vfp resource used from free pool", Computed: true}, "rsvd": schema.Int64Attribute{MarkdownDescription: "vfp resource reserved for app", Computed: true}, "rsvd_used": schema.Int64Attribute{MarkdownDescription: "resource used from 'rsvd'", Computed: true}}}}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "slot_id": schema.StringAttribute{MarkdownDescription: "Device slotId", Required: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetAppFilterRscBySlotIdDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetAppFilterRscBySlotIdDataSourceModel
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
func (d *GetAppFilterRscBySlotIdDataSource) readRemote(ctx context.Context, config *GetAppFilterRscBySlotIdDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/appFilterRsc/{slotId}"
	reqPath = strings.ReplaceAll(reqPath, "{slotId}", url.PathEscape(config.SlotId.ValueString()))
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_app_filter_rsc_by_slot_id", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_app_filter_rsc_by_slot_id", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_app_filter_rsc_by_slot_id", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_app_filter_rsc_by_slot_id", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_app_filter_rsc_by_slot_id", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_app_filter_rsc_by_slot_id", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_app_filter_rsc_by_slot_id", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_app_filter_rsc_by_slot_id", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_app_filter_rsc_by_slot_id", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_app_filter_rsc_by_slot_id", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_app_filter_rsc_by_slot_id", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_app_filter_rsc_by_slot_id", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_app_filter_rsc_by_slot_id", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["appFilterRsc"]; ok {
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_app_filter_rsc_by_slot_id", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetAppFilterRscBySlotIdDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
