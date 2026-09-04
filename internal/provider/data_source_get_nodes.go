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
	_ datasource.DataSource              = (*GetNodesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetNodesDataSource)(nil)
)

// GetNodesDataSource is the generated Terraform data source implementation.
type GetNodesDataSource struct {
	client *client.Client
}

// GetNodesDataSourceModel describes the data source state shape.
type GetNodesDataSourceModel struct {
	EnvId   types.String `tfsdk:"env_id" json:"envId"`
	Items   types.List   `tfsdk:"items"`
	UnifyId types.String `tfsdk:"unify_id" json:"unifyId"`
}

// NewGetNodesDataSource returns a new instance of the generated data source.
func NewGetNodesDataSource() datasource.DataSource {
	return &GetNodesDataSource{}
}

// Metadata returns the data source type name.
func (d *GetNodesDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_nodes"
}

// Schema returns the data source schema.
func (d *GetNodesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get unified resource deployment nodes by environment and connection id", Attributes: map[string]schema.Attribute{"env_id": schema.StringAttribute{MarkdownDescription: "unified environment identifier", Required: true}, "items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"domain_id": schema.SingleNestedAttribute{MarkdownDescription: "Unified Environment", Computed: true, Attributes: map[string]schema.Attribute{"name": schema.StringAttribute{Computed: true}, "type": schema.StringAttribute{MarkdownDescription: "Unified Environment platform types", Computed: true}}}, "env": schema.SingleNestedAttribute{MarkdownDescription: "Unified Environment", Computed: true, Attributes: map[string]schema.Attribute{"name": schema.StringAttribute{Computed: true}, "type": schema.StringAttribute{MarkdownDescription: "Unified Environment platform types", Computed: true}}}, "fabric_node_id": schema.StringAttribute{Computed: true}, "interfaces": schema.SingleNestedAttribute{MarkdownDescription: "Unified Resource interfaces", Computed: true, Attributes: map[string]schema.Attribute{"env": schema.SetNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"fabric_node_id": schema.StringAttribute{Computed: true}, "mgmt_ip": schema.StringAttribute{Computed: true}, "network_id": schema.StringAttribute{Computed: true}, "v_series_node_interface": schema.SingleNestedAttribute{MarkdownDescription: "Cloud V Series node interface object", Computed: true, Attributes: map[string]schema.Attribute{"ip_address": schema.StringAttribute{Computed: true}, "mac_address": schema.StringAttribute{Computed: true}, "name": schema.StringAttribute{Computed: true}, "public_ip_address": schema.StringAttribute{Computed: true}, "subnet_cidr": schema.StringAttribute{Computed: true}, "subnet_id": schema.StringAttribute{Computed: true}}}, "virt_domain_id": schema.SingleNestedAttribute{MarkdownDescription: "Virtual Domain Id for cloud resources", Computed: true, Attributes: map[string]schema.Attribute{"name": schema.StringAttribute{Computed: true}, "type": schema.StringAttribute{Computed: true}}}}}}}}, "mgmt_ip": schema.StringAttribute{Computed: true}, "unification_conn_id": schema.StringAttribute{Computed: true}}}}, "unify_id": schema.StringAttribute{MarkdownDescription: "unified resource identifier", Required: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetNodesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetNodesDataSourceModel
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
func (d *GetNodesDataSource) readListRemote(ctx context.Context, config *GetNodesDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/unifiedResources/env/{envId}/unifyId/{unifyId}/nodes"
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_nodes", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_nodes", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["nodes"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_nodes", fmt.Sprintf("Could not decode list page: missing %q array", "nodes"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_nodes", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetNodesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
