package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)
import (
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*GetAllForeignSourcesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetAllForeignSourcesDataSource)(nil)
)

// GetAllForeignSourcesDataSource is the generated Terraform data source implementation.
type GetAllForeignSourcesDataSource struct {
	client *client.Client
}

// GetAllForeignSourcesDataSourceModel describes the data source state shape.
type GetAllForeignSourcesDataSourceModel struct {
	BoxId types.Int64 `tfsdk:"box_id" json:"boxId"`
	Items types.List  `tfsdk:"items"`
}

// NewGetAllForeignSourcesDataSource returns a new instance of the generated data source.
func NewGetAllForeignSourcesDataSource() datasource.DataSource {
	return &GetAllForeignSourcesDataSource{}
}

// Metadata returns the data source type name.
func (d *GetAllForeignSourcesDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_all_foreign_sources"
}

// Schema returns the data source schema.
func (d *GetAllForeignSourcesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Reads the get all foreign sources data source.", Attributes: map[string]schema.Attribute{"box_id": schema.Int64Attribute{MarkdownDescription: "specify the cluster node by boxId. By default all nodes are selected.", Optional: true}, "items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"best_source": schema.Int64Attribute{Computed: true}, "foreign_sources": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"announce_msg_count": schema.Int64Attribute{MarkdownDescription: "Number of Announce messages from the foreign source that have been received within a time window", Computed: true}, "clock_address": schema.StringAttribute{Computed: true}, "last_delay_response": schema.StringAttribute{Computed: true}, "last_follow_up": schema.StringAttribute{Computed: true}, "last_sync": schema.StringAttribute{Computed: true}, "mtsd_scaled_avar": schema.Int64Attribute{Computed: true}, "port_identity": schema.StringAttribute{Computed: true}, "port_module_number": schema.Int64Attribute{Computed: true}, "port_number": schema.Int64Attribute{Computed: true}, "primary_source_clock_identity": schema.StringAttribute{MarkdownDescription: "primary source's Clock Identifier as per the IEEE 1588 standard which is represented by an integer 8-octet array", Computed: true}, "primary_source_priority1": schema.Int64Attribute{Computed: true}, "primary_source_priority2": schema.Int64Attribute{Computed: true}, "ptp_protocol": schema.StringAttribute{Computed: true}, "quality": schema.SingleNestedAttribute{MarkdownDescription: "Clock Quality determines the quality of a clock based on class and accuracy. Used to select source clock in Best Source Clock Algorithm", Computed: true, Attributes: map[string]schema.Attribute{"accuracy": schema.StringAttribute{MarkdownDescription: "Indicates the expected accuracy of a clock when it is the primary source or in the event it becomes the primary source", Computed: true}, "class": schema.Int64Attribute{MarkdownDescription: "Denotes the traceability of the time or frequency distributed by the primary source clock", Computed: true}, "offset": schema.Int64Attribute{MarkdownDescription: "Indicates the estimate of the variations of the local clock from a linear timescale when it is not synchronized to another clock using the protocol", Computed: true}}}, "steps_removed": schema.Int64Attribute{Computed: true}}}}, "port_id": schema.StringAttribute{Computed: true}}}}}}
}

// Read fetches remote state into the data source model.
func (d *GetAllForeignSourcesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetAllForeignSourcesDataSourceModel
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
func (d *GetAllForeignSourcesDataSource) readListRemote(ctx context.Context, config *GetAllForeignSourcesDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/ptp/portState/foreignSource"
	params := url.Values{}
	if !config.BoxId.IsNull() {
		params.Set("boxId", strconv.FormatInt(config.BoxId.ValueInt64(), 10))
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_foreign_sources", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_foreign_sources", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["foreignSources"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_foreign_sources", fmt.Sprintf("Could not decode list page: missing %q array", "foreignSources"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_foreign_sources", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetAllForeignSourcesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
