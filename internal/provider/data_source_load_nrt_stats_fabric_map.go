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
	_ datasource.DataSource              = (*LoadNrtStatsFabricMapDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadNrtStatsFabricMapDataSource)(nil)
)

// LoadNrtStatsFabricMapDataSource is the generated Terraform data source implementation.
type LoadNrtStatsFabricMapDataSource struct {
	client *client.Client
}

// LoadNrtStatsFabricMapDataSourceModel describes the data source state shape.
type LoadNrtStatsFabricMapDataSourceModel struct {
	FabricMapAlias types.String `tfsdk:"fabric_map_alias" json:"fabricMapAlias"`
	Items          types.List   `tfsdk:"items"`
}

// NewLoadNrtStatsFabricMapDataSource returns a new instance of the generated data source.
func NewLoadNrtStatsFabricMapDataSource() datasource.DataSource {
	return &LoadNrtStatsFabricMapDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadNrtStatsFabricMapDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_nrt_stats_fabric_map"
}

// Schema returns the data source schema.
func (d *LoadNrtStatsFabricMapDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get NRT stats for the user-defined fabric map", Attributes: map[string]schema.Attribute{"fabric_map_alias": schema.StringAttribute{MarkdownDescription: "alias of the fabric map", Required: true}, "items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "unique nrt alias", Computed: true}, "cluster_name": schema.StringAttribute{MarkdownDescription: "cluster name", Computed: true}, "component_type": schema.StringAttribute{Computed: true}, "health_state": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state; 'grey' indicates undeployed state;", Computed: true}, "health_state_reasons": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"message": schema.StringAttribute{MarkdownDescription: "Read-only. Describes the reason for component's health state", Computed: true}, "severity": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "traffic_health_state_computation_type": schema.StringAttribute{MarkdownDescription: "Traffic Health State Computation Type", Computed: true}}}}, "nrt_alias": schema.StringAttribute{MarkdownDescription: "unique nrt alias", Computed: true}, "reason": schema.StringAttribute{Computed: true}, "stats_data": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{Computed: true}, "current_interval_start": schema.StringAttribute{Computed: true}, "current_offset_amount": schema.Int64Attribute{Computed: true}, "fabric_map_alias": schema.StringAttribute{MarkdownDescription: "unique fabric map alias", Computed: true}, "flex_dir": schema.StringAttribute{MarkdownDescription: "only applicable for flexInline maps", Computed: true}, "fstype": schema.StringAttribute{Computed: true}, "interval_duration": schema.StringAttribute{Computed: true}, "next_interval_start": schema.StringAttribute{Computed: true}, "octets": schema.StringAttribute{Computed: true}, "offset_shift_per_interval": schema.Int64Attribute{Computed: true}, "packets": schema.StringAttribute{Computed: true}, "rotational_time_rem": schema.StringAttribute{Computed: true}, "rules": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"accepted": schema.Int64Attribute{MarkdownDescription: "SecondLevel/flowSample only", Computed: true}, "accepted_bytes": schema.Int64Attribute{MarkdownDescription: "SecondLevel/flowSample only", Computed: true}, "accepted_pkts": schema.Int64Attribute{MarkdownDescription: "SecondLevel/flowSample only", Computed: true}, "entries": schema.Int64Attribute{MarkdownDescription: "SecondLevel/whitelist only", Computed: true}, "fs_interval_type": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"current_interval": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"interval_start": schema.StringAttribute{Computed: true}, "sample_range": schema.StringAttribute{Computed: true}}}, "future_intervals": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"interval_start": schema.StringAttribute{Computed: true}, "sample_range": schema.StringAttribute{Computed: true}}}}, "prior_intervals": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"interval_start": schema.StringAttribute{Computed: true}, "sample_range": schema.StringAttribute{Computed: true}}}}}}, "ip_can_bearer": schema.Int64Attribute{MarkdownDescription: "SecondLevel/flowFilter, flowWhitelist only", Computed: true}, "map_alias": schema.StringAttribute{MarkdownDescription: "Map Alias", Computed: true}, "matched": schema.Int64Attribute{MarkdownDescription: "SecondLevel/flowSample only", Computed: true}, "matched_bytes": schema.Int64Attribute{MarkdownDescription: "SecondLevel/whitelist only", Computed: true}, "matched_pkts": schema.Int64Attribute{MarkdownDescription: "SecondLevel/whitelist only", Computed: true}, "octets": schema.Int64Attribute{Computed: true}, "packets": schema.Int64Attribute{Computed: true}, "rejected": schema.Int64Attribute{MarkdownDescription: "SecondLevel/flowSample only", Computed: true}, "rejected_bytes": schema.Int64Attribute{MarkdownDescription: "SecondLevel/flowSample only", Computed: true}, "rejected_pkts": schema.Int64Attribute{MarkdownDescription: "SecondLevel/flowSample only", Computed: true}, "rule_id": schema.Int64Attribute{MarkdownDescription: "Map Rule Id", Computed: true}, "rule_type": schema.StringAttribute{Computed: true}, "rules": schema.Int64Attribute{MarkdownDescription: "SecondLevel/flowSample only", Computed: true}}}}, "sub_type": schema.StringAttribute{MarkdownDescription: "'byRule' is applicable to all map types; 'collector' is applicable to 'regular', 'inline', 'flexInline' and 'secondLevel' maps; 'passAll' is applicable to 'regular' and 'inline' maps; 'flowFilter', 'flowSample', 'flowWhitelist', 'flowSampleSip', 'flowWhitelistSip', 'flowSampleDiameter', 'flowWhitelistDiameter', 'flowSampleOverlap', 'flowWhitelistOverlap', 'flowSample5g','flowWhitelist5g','flowSample5gOverlap' and 'flowWhitelist5gOverlap' are applicable to 'secondLevel' maps", Computed: true}, "timestamp": schema.StringAttribute{MarkdownDescription: "Counter collection date and time in ISO 8601 format", Computed: true}, "total_drop_octets": schema.Int64Attribute{Computed: true}, "total_drop_packets": schema.Int64Attribute{Computed: true}, "total_pass_octets": schema.Int64Attribute{Computed: true}, "total_pass_packets": schema.Int64Attribute{Computed: true}, "type": schema.StringAttribute{MarkdownDescription: "'regular' maps are from network/hybrid ports to tool/hybrid/gigastream; 'inline' maps are from inline ports to inline/tool/hybrid/gigastream; 'firstLevel' are from network/hybrid ports to vPorts/tool/hybrid/gigastream; 'secondLevel' maps are from vPorts to tool/hybrid/gigastream; 'inlineFirstLevel is from inline network to vport; 'inlineSecondLevel' is from vport to inline tool; 'flexInline' is from inline network to tools; 'transitLevel' is from vport to vport", Computed: true}}}, "traffic_health_state": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates traffic healthy state; 'red'  indicates traffic critical state; 'grey' indicates traffic not monitored state;", Computed: true}, "traffic_health_state_reasons": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"message": schema.StringAttribute{MarkdownDescription: "Read-only. Describes the reason for component's health state", Computed: true}, "severity": schema.StringAttribute{MarkdownDescription: "Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;", Computed: true}, "traffic_health_state_computation_type": schema.StringAttribute{MarkdownDescription: "Traffic Health State Computation Type", Computed: true}}}}, "update_time": schema.StringAttribute{MarkdownDescription: "Last Updated time of the fabric map", Computed: true}}}}}}
}

// Read fetches remote state into the data source model.
func (d *LoadNrtStatsFabricMapDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadNrtStatsFabricMapDataSourceModel
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
func (d *LoadNrtStatsFabricMapDataSource) readListRemote(ctx context.Context, config *LoadNrtStatsFabricMapDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/fabricMaps/nrtStats/{fabricMapAlias}"
	reqPath = strings.ReplaceAll(reqPath, "{fabricMapAlias}", url.PathEscape(config.FabricMapAlias.ValueString()))
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
		resp.Diagnostics.AddError("Error reading gigavuecore_load_nrt_stats_fabric_map", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_nrt_stats_fabric_map", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["fmNearRealTimeCachedStatisticsList"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_nrt_stats_fabric_map", fmt.Sprintf("Could not decode list page: missing %q array", "fmNearRealTimeCachedStatisticsList"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_nrt_stats_fabric_map", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadNrtStatsFabricMapDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
