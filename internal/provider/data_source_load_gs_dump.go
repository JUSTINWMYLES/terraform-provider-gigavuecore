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
	_ datasource.DataSource              = (*LoadGsDumpDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadGsDumpDataSource)(nil)
)

// LoadGsDumpDataSource is the generated Terraform data source implementation.
type LoadGsDumpDataSource struct {
	client *client.Client
}

// LoadGsDumpDataSourceModel describes the data source state shape.
type LoadGsDumpDataSourceModel struct {
	ClusterId types.String `tfsdk:"cluster_id" json:"clusterId"`
	Items     types.List   `tfsdk:"items"`
}

// NewLoadGsDumpDataSource returns a new instance of the generated data source.
func NewLoadGsDumpDataSource() datasource.DataSource {
	return &LoadGsDumpDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadGsDumpDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_gs_dump"
}

// Schema returns the data source schema.
func (d *LoadGsDumpDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load Available Gigasmart dump Filenames", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster Id", Required: true}, "items": schema.ListNestedAttribute{MarkdownDescription: "List of available Gigasmart Dump files", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"completed_list": schema.StringAttribute{MarkdownDescription: "Eport list for which GigaSmart Dump is completed ", Computed: true}, "eport_list": schema.StringAttribute{MarkdownDescription: "Complete list of Eports for GigaSmart dump generation", Computed: true}, "failed_list": schema.StringAttribute{MarkdownDescription: "Eport list for which GigaSmart Dump failed", Computed: true}, "filename": schema.StringAttribute{MarkdownDescription: "Filename of the Gigasmart dump file", Computed: true}, "gs_exec_status": schema.StringAttribute{MarkdownDescription: "Execution status of  GigaSmart Dump", Computed: true}, "hostname": schema.StringAttribute{MarkdownDescription: "Hostname of the device where Gigasmart dump resides", Computed: true}, "size": schema.Int64Attribute{MarkdownDescription: "Size of file in bytes", Computed: true}, "timestamp": schema.StringAttribute{MarkdownDescription: "File creation date and time in ISO 8601 format", Computed: true}}}}}}
}

// Read fetches remote state into the data source model.
func (d *LoadGsDumpDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadGsDumpDataSourceModel
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
func (d *LoadGsDumpDataSource) readListRemote(ctx context.Context, config *LoadGsDumpDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/system/gsDump"
	params := url.Values{}
	params.Set("clusterId", config.ClusterId.ValueString())
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
		resp.Diagnostics.AddError("Error reading gigavuecore_load_gs_dump", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_gs_dump", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["gsDumpFiles"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_gs_dump", fmt.Sprintf("Could not decode list page: missing %q array", "gsDumpFiles"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_gs_dump", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadGsDumpDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
