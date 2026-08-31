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
	_ datasource.DataSource              = (*LoadFilterTemplateLimitBySlotIdDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadFilterTemplateLimitBySlotIdDataSource)(nil)
)

// LoadFilterTemplateLimitBySlotIdDataSource is the generated Terraform data source implementation.
type LoadFilterTemplateLimitBySlotIdDataSource struct {
	client *client.Client
}

// LoadFilterTemplateLimitBySlotIdDataSourceModel describes the data source state shape.
type LoadFilterTemplateLimitBySlotIdDataSourceModel struct {
	ClusterId    types.String `tfsdk:"cluster_id" json:"clusterId"`
	FilterLimits types.List   `tfsdk:"filter_limits" json:"filterLimits"`
	SlotId       types.String `tfsdk:"slot_id" json:"slotId"`
}

// NewLoadFilterTemplateLimitBySlotIdDataSource returns a new instance of the generated data source.
func NewLoadFilterTemplateLimitBySlotIdDataSource() datasource.DataSource {
	return &LoadFilterTemplateLimitBySlotIdDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadFilterTemplateLimitBySlotIdDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_filter_template_limit_by_slot_id"
}

// Schema returns the data source schema.
func (d *LoadFilterTemplateLimitBySlotIdDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load Filter Template Limits by slot Id", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "id of the defining cluster", Required: true}, "filter_limits": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "alias of filter template", Computed: true}, "limit": schema.Int64Attribute{Computed: true}, "qualifiers": schema.ListAttribute{Computed: true, ElementType: types.StringType}}}}, "slot_id": schema.StringAttribute{MarkdownDescription: "Device card slot Id", Required: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadFilterTemplateLimitBySlotIdDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadFilterTemplateLimitBySlotIdDataSourceModel
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
func (d *LoadFilterTemplateLimitBySlotIdDataSource) readRemote(ctx context.Context, config *LoadFilterTemplateLimitBySlotIdDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/filterTemplates/limits/{slotId}"
	reqPath = strings.ReplaceAll(reqPath, "{slotId}", url.PathEscape(config.SlotId.ValueString()))
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_filter_template_limit_by_slot_id", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_filter_template_limit_by_slot_id", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_filter_template_limit_by_slot_id", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_filter_template_limit_by_slot_id", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_filter_template_limit_by_slot_id", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_filter_template_limit_by_slot_id", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_filter_template_limit_by_slot_id", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_filter_template_limit_by_slot_id", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_filter_template_limit_by_slot_id", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_filter_template_limit_by_slot_id", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_load_filter_template_limit_by_slot_id", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_load_filter_template_limit_by_slot_id", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_filter_template_limit_by_slot_id", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["filterTemplateLimit"]; ok {
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
		resp.Diagnostics.AddError("Error reading gigavuecore_load_filter_template_limit_by_slot_id", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadFilterTemplateLimitBySlotIdDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
