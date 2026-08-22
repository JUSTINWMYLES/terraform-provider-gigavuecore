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
	_ datasource.DataSource              = (*LoadAllCriticalNotificationsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadAllCriticalNotificationsDataSource)(nil)
)

// LoadAllCriticalNotificationsDataSource is the generated Terraform data source implementation.
type LoadAllCriticalNotificationsDataSource struct {
	client *client.Client
}

// LoadAllCriticalNotificationsDataSourceModel describes the data source state shape.
type LoadAllCriticalNotificationsDataSourceModel struct {
	Items types.List `tfsdk:"items"`
}

// NewLoadAllCriticalNotificationsDataSource returns a new instance of the generated data source.
func NewLoadAllCriticalNotificationsDataSource() datasource.DataSource {
	return &LoadAllCriticalNotificationsDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadAllCriticalNotificationsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_all_critical_notifications"
}

// Schema returns the data source schema.
func (d *LoadAllCriticalNotificationsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load All Critical Notifications", Attributes: map[string]schema.Attribute{"items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"banner_config_data": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"current_usage_percent": schema.Int64Attribute{Computed: true}, "disk_partition": schema.StringAttribute{MarkdownDescription: "Disk Partition", Computed: true}, "host_name": schema.StringAttribute{MarkdownDescription: "Host Name", Computed: true}, "role": schema.StringAttribute{MarkdownDescription: "Node Role", Computed: true}, "threshold_percent": schema.Int64Attribute{Computed: true}, "total_gb": schema.Int64Attribute{Computed: true}}}, "description": schema.StringAttribute{MarkdownDescription: "Notification description", Computed: true}, "severity": schema.StringAttribute{MarkdownDescription: "Notification Severity", Computed: true}, "type": schema.StringAttribute{MarkdownDescription: "Notification type", Computed: true}}}}}}
}

// Read fetches remote state into the data source model.
func (d *LoadAllCriticalNotificationsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadAllCriticalNotificationsDataSourceModel
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
func (d *LoadAllCriticalNotificationsDataSource) readListRemote(ctx context.Context, config *LoadAllCriticalNotificationsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/fmNotify/message"
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
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_critical_notifications", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_critical_notifications", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["message"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_critical_notifications", fmt.Sprintf("Could not decode list page: missing %q array", "message"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_critical_notifications", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadAllCriticalNotificationsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
