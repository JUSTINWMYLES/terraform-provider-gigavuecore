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
	_ datasource.DataSource              = (*GetAllForeignMastersDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetAllForeignMastersDataSource)(nil)
)

// GetAllForeignMastersDataSource is the generated Terraform data source implementation.
type GetAllForeignMastersDataSource struct {
	client *client.Client
}

// GetAllForeignMastersDataSourceModel describes the data source state shape.
type GetAllForeignMastersDataSourceModel struct {
	BoxId types.Int64 `tfsdk:"box_id" json:"boxId"`
	Items types.List  `tfsdk:"items"`
}

// NewGetAllForeignMastersDataSource returns a new instance of the generated data source.
func NewGetAllForeignMastersDataSource() datasource.DataSource {
	return &GetAllForeignMastersDataSource{}
}

// Metadata returns the data source type name.
func (d *GetAllForeignMastersDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_all_foreign_masters"
}

// Schema returns the data source schema.
func (d *GetAllForeignMastersDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "deprecated: use GET /ptp/portState/foreignSource", Attributes: map[string]schema.Attribute{"box_id": schema.Int64Attribute{MarkdownDescription: "specify the cluster node by boxId. By default all nodes are selected.", Optional: true}, "items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"best_master": schema.Int64Attribute{Computed: true}, "foreign_masters": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"announce_msg_count": schema.Int64Attribute{MarkdownDescription: "Number of Announce messages from the foreign master that have been received within a time window", Computed: true}, "clock_address": schema.StringAttribute{Computed: true}, "grandmaster_clock_identity": schema.StringAttribute{MarkdownDescription: "grandmaster's Clock Identifier as per the IEEE 1588 standard which is represented by an integer 8-octet array", Computed: true}, "grandmaster_priority1": schema.Int64Attribute{Computed: true}, "grandmaster_priority2": schema.Int64Attribute{Computed: true}, "last_delay_response": schema.StringAttribute{Computed: true}, "last_follow_up": schema.StringAttribute{Computed: true}, "last_sync": schema.StringAttribute{Computed: true}, "mtsd_scaled_avar": schema.Int64Attribute{Computed: true}, "port_identity": schema.StringAttribute{Computed: true}, "port_module_number": schema.Int64Attribute{Computed: true}, "port_number": schema.Int64Attribute{Computed: true}, "ptp_protocol": schema.StringAttribute{Computed: true}, "quality": schema.SingleNestedAttribute{MarkdownDescription: "Clock Quality determines the quality of a clock based on class and accuracy. Used to select master clock in Best Master Clock Algorithm. (deprecated: use primarySourceClockQuality instead of grandmasterClockQuality)", Computed: true, Attributes: map[string]schema.Attribute{"accuracy": schema.StringAttribute{MarkdownDescription: "Indicates the expected accuracy of a clock when it is the grandmaster or in the event it becomes the grandmaster", Computed: true}, "class": schema.Int64Attribute{MarkdownDescription: "Denotes the traceability of the time or frequency distributed by the grandmaster clock", Computed: true}, "offset": schema.Int64Attribute{MarkdownDescription: "Indicates the estimate of the variations of the local clock from a linear timescale when it is not synchronized to another clock using the protocol", Computed: true}}}, "steps_removed": schema.Int64Attribute{Computed: true}}}}, "port_id": schema.StringAttribute{Computed: true}}}}}}
}

// Read fetches remote state into the data source model.
func (d *GetAllForeignMastersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetAllForeignMastersDataSourceModel
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
func (d *GetAllForeignMastersDataSource) readListRemote(ctx context.Context, config *GetAllForeignMastersDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/ptp/portState/foreignMaster"
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_foreign_masters", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_foreign_masters", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["foreignMasters"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_foreign_masters", fmt.Sprintf("Could not decode list page: missing %q array", "foreignMasters"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_foreign_masters", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetAllForeignMastersDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
