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
	_ datasource.DataSource              = (*GetCopilotSysdumpsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetCopilotSysdumpsDataSource)(nil)
)

// GetCopilotSysdumpsDataSource is the generated Terraform data source implementation.
type GetCopilotSysdumpsDataSource struct {
	client *client.Client
}

// GetCopilotSysdumpsDataSourceModel describes the data source state shape.
type GetCopilotSysdumpsDataSourceModel struct {
	Items types.List `tfsdk:"items"`
}

// NewGetCopilotSysdumpsDataSource returns a new instance of the generated data source.
func NewGetCopilotSysdumpsDataSource() datasource.DataSource {
	return &GetCopilotSysdumpsDataSource{}
}

// Metadata returns the data source type name.
func (d *GetCopilotSysdumpsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_copilot_sysdumps"
}

// Schema returns the data source schema.
func (d *GetCopilotSysdumpsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load Available Sysdump Files", Attributes: map[string]schema.Attribute{"items": schema.ListNestedAttribute{MarkdownDescription: "List of copilot sysdump files", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"filename": schema.StringAttribute{MarkdownDescription: "The sysdump filename", Computed: true}, "size": schema.Int64Attribute{MarkdownDescription: "The sysdump file size in bytes", Computed: true}, "timestamp": schema.StringAttribute{MarkdownDescription: "File creation date and time in ISO 8601 format", Computed: true}}}}}}
}

// Read fetches remote state into the data source model.
func (d *GetCopilotSysdumpsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetCopilotSysdumpsDataSourceModel
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
func (d *GetCopilotSysdumpsDataSource) readListRemote(ctx context.Context, config *GetCopilotSysdumpsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/fm/copilot/sysdump"
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_copilot_sysdumps", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_copilot_sysdumps", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["sysdumpFiles"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_copilot_sysdumps", fmt.Sprintf("Could not decode list page: missing %q array", "sysdumpFiles"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_copilot_sysdumps", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetCopilotSysdumpsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
