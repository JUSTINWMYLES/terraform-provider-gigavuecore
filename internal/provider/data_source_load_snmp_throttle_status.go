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
	_ datasource.DataSource              = (*LoadSnmpThrottleStatusDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadSnmpThrottleStatusDataSource)(nil)
)

// LoadSnmpThrottleStatusDataSource is the generated Terraform data source implementation.
type LoadSnmpThrottleStatusDataSource struct {
	client *client.Client
}

// LoadSnmpThrottleStatusDataSourceModel describes the data source state shape.
type LoadSnmpThrottleStatusDataSourceModel struct {
	Items types.List `tfsdk:"items"`
}

// NewLoadSnmpThrottleStatusDataSource returns a new instance of the generated data source.
func NewLoadSnmpThrottleStatusDataSource() datasource.DataSource {
	return &LoadSnmpThrottleStatusDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadSnmpThrottleStatusDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_snmp_throttle_status"
}

// Schema returns the data source schema.
func (d *LoadSnmpThrottleStatusDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load Snmp Throttle Status", Attributes: map[string]schema.Attribute{"items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"elapsed_time": schema.Int64Attribute{MarkdownDescription: "Elapsed Time", Computed: true}, "entity_id": schema.StringAttribute{MarkdownDescription: "Entity ID", Computed: true}, "event": schema.SetAttribute{MarkdownDescription: "The set of notification event types", Computed: true, ElementType: types.StringType}, "interval": schema.Int64Attribute{MarkdownDescription: "Interval (in seconds)", Computed: true}, "last_time_triggered": schema.StringAttribute{MarkdownDescription: "Date and time of last time the event has been triggered", Computed: true}, "trigger_count": schema.Int64Attribute{MarkdownDescription: "Trigger count", Computed: true}}}}}}
}

// Read fetches remote state into the data source model.
func (d *LoadSnmpThrottleStatusDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadSnmpThrottleStatusDataSourceModel
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
func (d *LoadSnmpThrottleStatusDataSource) readListRemote(ctx context.Context, config *LoadSnmpThrottleStatusDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/system/snmp/throttleStatus"
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
		resp.Diagnostics.AddError("Error reading gigavuecore_load_snmp_throttle_status", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_snmp_throttle_status", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["snmpThrottleStatus"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_snmp_throttle_status", fmt.Sprintf("Could not decode list page: missing %q array", "snmpThrottleStatus"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_snmp_throttle_status", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadSnmpThrottleStatusDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
