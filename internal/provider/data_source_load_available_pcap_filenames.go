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
	_ datasource.DataSource              = (*LoadAvailablePcapFilenamesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadAvailablePcapFilenamesDataSource)(nil)
)

// LoadAvailablePcapFilenamesDataSource is the generated Terraform data source implementation.
type LoadAvailablePcapFilenamesDataSource struct {
	client *client.Client
}

// LoadAvailablePcapFilenamesDataSourceModel describes the data source state shape.
type LoadAvailablePcapFilenamesDataSourceModel struct {
	BoxId     types.String `tfsdk:"box_id" json:"boxId"`
	ClusterId types.String `tfsdk:"cluster_id" json:"clusterId"`
	Items     types.List   `tfsdk:"items"`
}

// NewLoadAvailablePcapFilenamesDataSource returns a new instance of the generated data source.
func NewLoadAvailablePcapFilenamesDataSource() datasource.DataSource {
	return &LoadAvailablePcapFilenamesDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadAvailablePcapFilenamesDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_available_pcap_filenames"
}

// Schema returns the data source schema.
func (d *LoadAvailablePcapFilenamesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load Available PCAP Filenames", Attributes: map[string]schema.Attribute{"box_id": schema.StringAttribute{MarkdownDescription: "Box ID range from 1 to 64(inclusive). all is applicable", Optional: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "items": schema.ListNestedAttribute{MarkdownDescription: "List of available pcap files", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"box_id": schema.StringAttribute{MarkdownDescription: "The chassis Box ID value", Computed: true}, "capture_files": schema.ListNestedAttribute{MarkdownDescription: "List of available pcap files", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"filename": schema.StringAttribute{MarkdownDescription: "Name of the PCAP file", Computed: true}, "size": schema.Int64Attribute{MarkdownDescription: "Size of file in bytes", Computed: true}, "timestamp": schema.StringAttribute{MarkdownDescription: "File creation date and time in ISO 8601 format", Computed: true}}}}}}}}}
}

// Read fetches remote state into the data source model.
func (d *LoadAvailablePcapFilenamesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadAvailablePcapFilenamesDataSourceModel
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
func (d *LoadAvailablePcapFilenamesDataSource) readListRemote(ctx context.Context, config *LoadAvailablePcapFilenamesDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/pcap/file"
	params := url.Values{}
	params.Set("clusterId", config.ClusterId.ValueString())
	if !config.BoxId.IsNull() {
		params.Set("boxId", config.BoxId.ValueString())
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
		resp.Diagnostics.AddError("Error reading gigavuecore_load_available_pcap_filenames", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_available_pcap_filenames", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["pcapFiles"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_available_pcap_filenames", fmt.Sprintf("Could not decode list page: missing %q array", "pcapFiles"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_available_pcap_filenames", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadAvailablePcapFilenamesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
