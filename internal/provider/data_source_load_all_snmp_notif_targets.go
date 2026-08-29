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
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*LoadAllSnmpNotifTargetsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadAllSnmpNotifTargetsDataSource)(nil)
)

// LoadAllSnmpNotifTargetsDataSource is the generated Terraform data source implementation.
type LoadAllSnmpNotifTargetsDataSource struct {
	client *client.Client
}

// LoadAllSnmpNotifTargetsDataSourceModel describes the data source state shape.
type LoadAllSnmpNotifTargetsDataSourceModel struct {
	ClusterId types.String `tfsdk:"cluster_id" json:"clusterId"`
	Items     types.List   `tfsdk:"items"`
	Page      types.String `tfsdk:"page"`
	Sort      types.String `tfsdk:"sort"`
}

// NewLoadAllSnmpNotifTargetsDataSource returns a new instance of the generated data source.
func NewLoadAllSnmpNotifTargetsDataSource() datasource.DataSource {
	return &LoadAllSnmpNotifTargetsDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadAllSnmpNotifTargetsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_all_snmp_notif_targets"
}

// Schema returns the data source schema.
func (d *LoadAllSnmpNotifTargetsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load all SNMP Notification Targets", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"enabled": schema.BoolAttribute{MarkdownDescription: "temporarily enable/disable the notification destination", Computed: true}, "host": schema.StringAttribute{MarkdownDescription: "ipv4 or ipv6 or domain name", Computed: true}, "notify_config": schema.SingleNestedAttribute{MarkdownDescription: "Notification Target configuration for specific Notification type (Trap/Inform)", Computed: true, Attributes: map[string]schema.Attribute{"auth_key": schema.StringAttribute{MarkdownDescription: "authentication password. required with 'v3user'", Computed: true}, "auth_protocol": schema.StringAttribute{MarkdownDescription: "authentication hash algorithm. required with 'v3user'", Computed: true}, "community": schema.StringAttribute{MarkdownDescription: "required when when 'version' is 'v2c'", Computed: true}, "engine_id": schema.StringAttribute{MarkdownDescription: "remote engineID. only valid with notifyType 'inform' and 'version' v3", Computed: true}, "port": schema.Int64Attribute{Computed: true}, "priv_key": schema.StringAttribute{MarkdownDescription: "privacy password", Computed: true}, "priv_protocol": schema.StringAttribute{MarkdownDescription: "privacy encryption", Computed: true}, "v3_user": schema.StringAttribute{MarkdownDescription: "required when when 'version' is 'v3'", Computed: true}, "version": schema.StringAttribute{MarkdownDescription: "SNMP version to use. v1 is only valid for traps. for v3, user name should be provided", Computed: true}}}, "notify_type": schema.StringAttribute{MarkdownDescription: "SNMP notification type to use", Computed: true}}}}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned", Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadAllSnmpNotifTargetsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadAllSnmpNotifTargetsDataSourceModel
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
func (d *LoadAllSnmpNotifTargetsDataSource) readListRemote(ctx context.Context, config *LoadAllSnmpNotifTargetsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/system/snmp/notifTargets"
	params := url.Values{}
	params.Set("clusterId", config.ClusterId.ValueString())
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
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_snmp_notif_targets", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_snmp_notif_targets", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["notifTargets"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_snmp_notif_targets", fmt.Sprintf("Could not decode list page: missing %q array", "notifTargets"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_snmp_notif_targets", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadAllSnmpNotifTargetsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
