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
	_ datasource.DataSource              = (*GetAllTimestampBriefDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetAllTimestampBriefDataSource)(nil)
)

// GetAllTimestampBriefDataSource is the generated Terraform data source implementation.
type GetAllTimestampBriefDataSource struct {
	client *client.Client
}

// GetAllTimestampBriefDataSourceModel describes the data source state shape.
type GetAllTimestampBriefDataSourceModel struct {
	BoxId types.Int64 `tfsdk:"box_id" json:"boxId"`
	Items types.List  `tfsdk:"items"`
}

// NewGetAllTimestampBriefDataSource returns a new instance of the generated data source.
func NewGetAllTimestampBriefDataSource() datasource.DataSource {
	return &GetAllTimestampBriefDataSource{}
}

// Metadata returns the data source type name.
func (d *GetAllTimestampBriefDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_all_timestamp_brief"
}

// Schema returns the data source schema.
func (d *GetAllTimestampBriefDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Reads the get all timestamp brief data source.", Attributes: map[string]schema.Attribute{"box_id": schema.Int64Attribute{MarkdownDescription: "specify the cluster node by boxId. By default all nodes are selected.", Optional: true}, "items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"admin": schema.BoolAttribute{Computed: true}, "egress_src_id": schema.Int64Attribute{Computed: true}, "egress_timestamp": schema.BoolAttribute{Computed: true}, "ingress_src_id": schema.Int64Attribute{Computed: true}, "ingress_timestamp": schema.BoolAttribute{Computed: true}, "port_id": schema.StringAttribute{Computed: true}}}}}}
}

// Read fetches remote state into the data source model.
func (d *GetAllTimestampBriefDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetAllTimestampBriefDataSourceModel
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
func (d *GetAllTimestampBriefDataSource) readListRemote(ctx context.Context, config *GetAllTimestampBriefDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/timestamp"
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_timestamp_brief", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_timestamp_brief", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["timestamps"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_timestamp_brief", fmt.Sprintf("Could not decode list page: missing %q array", "timestamps"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_timestamp_brief", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetAllTimestampBriefDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
