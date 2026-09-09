package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)
import (
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*LoadAllMetadataExportersDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadAllMetadataExportersDataSource)(nil)
)

// LoadAllMetadataExportersDataSource is the generated Terraform data source implementation.
type LoadAllMetadataExportersDataSource struct {
	client *client.Client
}

// LoadAllMetadataExportersDataSourceModel describes the data source state shape.
type LoadAllMetadataExportersDataSourceModel struct {
	Items types.List   `tfsdk:"items"`
	Page  types.String `tfsdk:"page"`
	Sort  types.String `tfsdk:"sort"`
}

// NewLoadAllMetadataExportersDataSource returns a new instance of the generated data source.
func NewLoadAllMetadataExportersDataSource() datasource.DataSource {
	return &LoadAllMetadataExportersDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadAllMetadataExportersDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_all_metadata_exporters"
}

// Schema returns the data source schema.
func (d *LoadAllMetadataExportersDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load All Metadata Exporters", Attributes: map[string]schema.Attribute{"items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{Computed: true}, "application_profiles": schema.ListAttribute{MarkdownDescription: "application profile aliases to attach to the exporter", Computed: true, ElementType: types.StringType}, "cef": schema.SingleNestedAttribute{MarkdownDescription: "cef attributes", Computed: true, Attributes: map[string]schema.Attribute{"active_timeout": schema.Int64Attribute{MarkdownDescription: "in seconds", Computed: true}, "inactive_timeout": schema.Int64Attribute{MarkdownDescription: "in seconds", Computed: true}}}, "description": schema.StringAttribute{Computed: true}, "destination": schema.SingleNestedAttribute{MarkdownDescription: "destination attributes", Computed: true, Attributes: map[string]schema.Attribute{"dscp": schema.Int64Attribute{Computed: true}, "ipv4_address": schema.StringAttribute{MarkdownDescription: "ipv4 address", Computed: true}, "l4_port_dst": schema.Int64Attribute{Computed: true}, "l4_port_src": schema.Int64Attribute{Computed: true}, "l4_protocol": schema.StringAttribute{Computed: true}, "ttl": schema.Int64Attribute{Computed: true}}}, "max_pkt_size": schema.Int64Attribute{Computed: true}, "mobility_sam": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"encoding": schema.StringAttribute{Computed: true}, "encoding_format": schema.StringAttribute{Computed: true}, "event_enable": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"modify": schema.BoolAttribute{Computed: true}, "update": schema.BoolAttribute{Computed: true}}}, "trigger": schema.StringAttribute{Computed: true}}}, "monitor": schema.SingleNestedAttribute{MarkdownDescription: "monitor attributes", Computed: true, Attributes: map[string]schema.Attribute{"timeout": schema.Int64Attribute{MarkdownDescription: "how often to export in seconds", Computed: true}}}, "netflow": schema.SingleNestedAttribute{MarkdownDescription: "netflow attributes", Computed: true, Attributes: map[string]schema.Attribute{"active_timeout": schema.Int64Attribute{MarkdownDescription: "in seconds", Computed: true}, "inactive_timeout": schema.Int64Attribute{MarkdownDescription: "in seconds", Computed: true}, "template_refresh": schema.Int64Attribute{MarkdownDescription: "template refresh interval in seconds", Computed: true}, "template_type": schema.StringAttribute{Computed: true}, "version": schema.StringAttribute{Computed: true}}}, "snmp": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"enabled": schema.BoolAttribute{MarkdownDescription: "snmp reverse lookup enable/disable", Computed: true}}}, "source": schema.SingleNestedAttribute{MarkdownDescription: "source tunnel port", Computed: true, Attributes: map[string]schema.Attribute{"ip_interface": schema.StringAttribute{Computed: true}}}, "type": schema.StringAttribute{Computed: true}}}}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned", Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadAllMetadataExportersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadAllMetadataExportersDataSourceModel
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
func (d *LoadAllMetadataExportersDataSource) readListRemote(ctx context.Context, config *LoadAllMetadataExportersDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/apps/metadata/exporters"
	params := url.Values{}
	if !config.Page.IsNull() {
		params.Set("page", config.Page.ValueString())
	}
	if !config.Sort.IsNull() {
		params.Set("sort", config.Sort.ValueString())
	}
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
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_metadata_exporters", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_metadata_exporters", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["metadataExporters"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_metadata_exporters", fmt.Sprintf("Could not decode list page: missing %q array", "metadataExporters"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_metadata_exporters", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadAllMetadataExportersDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
