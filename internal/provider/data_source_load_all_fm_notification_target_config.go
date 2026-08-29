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
	_ datasource.DataSource              = (*LoadAllFmNotificationTargetConfigDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadAllFmNotificationTargetConfigDataSource)(nil)
)

// LoadAllFmNotificationTargetConfigDataSource is the generated Terraform data source implementation.
type LoadAllFmNotificationTargetConfigDataSource struct {
	client *client.Client
}

// LoadAllFmNotificationTargetConfigDataSourceModel describes the data source state shape.
type LoadAllFmNotificationTargetConfigDataSourceModel struct {
	Items types.List   `tfsdk:"items"`
	Page  types.String `tfsdk:"page"`
	Sort  types.String `tfsdk:"sort"`
}

// NewLoadAllFmNotificationTargetConfigDataSource returns a new instance of the generated data source.
func NewLoadAllFmNotificationTargetConfigDataSource() datasource.DataSource {
	return &LoadAllFmNotificationTargetConfigDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadAllFmNotificationTargetConfigDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_all_fm_notification_target_config"
}

// Schema returns the data source schema.
func (d *LoadAllFmNotificationTargetConfigDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load All Fm Notification Target Config", Attributes: map[string]schema.Attribute{"items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"interface_name": schema.StringAttribute{MarkdownDescription: "Name of any one of the available network interface names on the FM. If this is chosen, FM registers itself as a notification target on the node with Ipv6 address if both FM and the node have Ipv6 address otherwise Ipv4 address is used. If targetAddress is configured then interfaceName has no effect.", Computed: true}, "interface_type": schema.StringAttribute{MarkdownDescription: "Notification target interface type", Computed: true}, "target_address": schema.StringAttribute{MarkdownDescription: "Configure FM's DNS name or static IP address to receive the management or data traffic from the node. The configured address is used by FM to register itself as a notification target on the node", Computed: true}}}}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned", Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadAllFmNotificationTargetConfigDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadAllFmNotificationTargetConfigDataSourceModel
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
func (d *LoadAllFmNotificationTargetConfigDataSource) readListRemote(ctx context.Context, config *LoadAllFmNotificationTargetConfigDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/fm/notifConfig"
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
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_fm_notification_target_config", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_fm_notification_target_config", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["fmNotificationTargetConfigs"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_fm_notification_target_config", fmt.Sprintf("Could not decode list page: missing %q array", "fmNotificationTargetConfigs"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_fm_notification_target_config", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadAllFmNotificationTargetConfigDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
