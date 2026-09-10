package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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
	_ datasource.DataSource              = (*GetAllAppInfoDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetAllAppInfoDataSource)(nil)
)

// GetAllAppInfoDataSource is the generated Terraform data source implementation.
type GetAllAppInfoDataSource struct {
	client *client.Client
}

// GetAllAppInfoDataSourceModel describes the data source state shape.
type GetAllAppInfoDataSourceModel struct {
	EnvId   types.String `tfsdk:"env_id" json:"envId"`
	Items   types.Set    `tfsdk:"items"`
	UnifyId types.String `tfsdk:"unify_id" json:"unifyId"`
}

// NewGetAllAppInfoDataSource returns a new instance of the generated data source.
func NewGetAllAppInfoDataSource() datasource.DataSource {
	return &GetAllAppInfoDataSource{}
}

// Metadata returns the data source type name.
func (d *GetAllAppInfoDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_all_app_info"
}

// Schema returns the data source schema.
func (d *GetAllAppInfoDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "List all appInfo for unified deployment via environment and connection id", Attributes: map[string]schema.Attribute{"env_id": schema.StringAttribute{MarkdownDescription: "unified environment identifier", Required: true}, "items": schema.SetNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"apps": schema.SingleNestedAttribute{MarkdownDescription: "Unified Resource app info for specific app", Computed: true, Attributes: map[string]schema.Attribute{"max_nr_aeps": schema.StringAttribute{Computed: true}, "mem_footprints": schema.SingleNestedAttribute{MarkdownDescription: "app specific memory footprint info", Computed: true, Attributes: map[string]schema.Attribute{"base": schema.Int64Attribute{Computed: true}, "scale": schema.SingleNestedAttribute{MarkdownDescription: "app specific memory footprint scale info", Computed: true, Attributes: map[string]schema.Attribute{"item": schema.StringAttribute{Computed: true}, "memory": schema.Int64Attribute{Computed: true}, "shared": schema.BoolAttribute{Computed: true}}}}}, "name": schema.StringAttribute{Computed: true}, "packet_permission": schema.StringAttribute{Computed: true}}}, "dp_type": schema.StringAttribute{Computed: true}, "event_dev": schema.StringAttribute{Computed: true}, "major_version": schema.Int64Attribute{Computed: true}, "minor_version": schema.Int64Attribute{Computed: true}, "name": schema.StringAttribute{Computed: true}, "path": schema.StringAttribute{Computed: true}, "platform_type": schema.StringAttribute{Computed: true}, "serial_num": schema.Int64Attribute{Computed: true}, "type": schema.StringAttribute{Computed: true}, "vendor": schema.StringAttribute{Computed: true}}}}, "unify_id": schema.StringAttribute{MarkdownDescription: "unified resource identifier", Required: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetAllAppInfoDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetAllAppInfoDataSourceModel
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
func (d *GetAllAppInfoDataSource) readListRemote(ctx context.Context, config *GetAllAppInfoDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/unifiedResources/env/{envId}/unifyId/{unifyId}/appinfo"
	reqPath = strings.ReplaceAll(reqPath, "{envId}", url.PathEscape(config.EnvId.ValueString()))
	reqPath = strings.ReplaceAll(reqPath, "{unifyId}", url.PathEscape(config.UnifyId.ValueString()))
	params := url.Values{}
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_app_info", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_app_info", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["env"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_app_info", fmt.Sprintf("Could not decode list page: missing %q array", "env"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_app_info", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetAllAppInfoDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
