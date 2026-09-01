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
	_ datasource.DataSource              = (*GetDeactivableVblsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetDeactivableVblsDataSource)(nil)
)

// GetDeactivableVblsDataSource is the generated Terraform data source implementation.
type GetDeactivableVblsDataSource struct {
	client *client.Client
}

// GetDeactivableVblsDataSourceModel describes the data source state shape.
type GetDeactivableVblsDataSourceModel struct {
	Items types.List `tfsdk:"items"`
}

// NewGetDeactivableVblsDataSource returns a new instance of the generated data source.
func NewGetDeactivableVblsDataSource() datasource.DataSource {
	return &GetDeactivableVblsDataSource{}
}

// Metadata returns the data source type name.
func (d *GetDeactivableVblsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_deactivable_vbls"
}

// Schema returns the data source schema.
func (d *GetDeactivableVblsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get all VBL activations that can be deactivated (and are not already deactivated)", Attributes: map[string]schema.Attribute{"items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"aid": schema.StringAttribute{MarkdownDescription: "activation id as encoded in the license", Computed: true}, "end_date": schema.Int64Attribute{MarkdownDescription: "end date of the license (format YYYYMMDD)", Computed: true}, "license_status": schema.StringAttribute{MarkdownDescription: "status of the license", Computed: true}, "sku": schema.StringAttribute{MarkdownDescription: "the license SKU (Stock Keeping Unit) or product code that uniquely identifies the license type in Gigamon catalog", Computed: true}, "start_date": schema.Int64Attribute{MarkdownDescription: "start date of the license (format YYYYMMDD)", Computed: true}, "total_volume": schema.StringAttribute{MarkdownDescription: "ingress traffic volume licensed per day (in GigaBytes)", Computed: true}}}}}}
}

// Read fetches remote state into the data source model.
func (d *GetDeactivableVblsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetDeactivableVblsDataSourceModel
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
func (d *GetDeactivableVblsDataSource) readListRemote(ctx context.Context, config *GetDeactivableVblsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/licensing/volumes/deactivable"
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_deactivable_vbls", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_deactivable_vbls", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["vblActivations"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_deactivable_vbls", fmt.Sprintf("Could not decode list page: missing %q array", "vblActivations"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_deactivable_vbls", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetDeactivableVblsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
