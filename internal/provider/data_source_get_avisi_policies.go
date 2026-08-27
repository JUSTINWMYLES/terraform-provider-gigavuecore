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
	_ datasource.DataSource              = (*GetAvisiPoliciesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetAvisiPoliciesDataSource)(nil)
)

// GetAvisiPoliciesDataSource is the generated Terraform data source implementation.
type GetAvisiPoliciesDataSource struct {
	client *client.Client
}

// GetAvisiPoliciesDataSourceModel describes the data source state shape.
type GetAvisiPoliciesDataSourceModel struct {
	Items types.List `tfsdk:"items"`
}

// NewGetAvisiPoliciesDataSource returns a new instance of the generated data source.
func NewGetAvisiPoliciesDataSource() datasource.DataSource {
	return &GetAvisiPoliciesDataSource{}
}

// Metadata returns the data source type name.
func (d *GetAvisiPoliciesDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_avisi_policies"
}

// Schema returns the data source schema.
func (d *GetAvisiPoliciesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get Active Visibility Policies", Attributes: map[string]schema.Attribute{"items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"description": schema.StringAttribute{Computed: true}, "enabled": schema.BoolAttribute{Computed: true}, "name": schema.StringAttribute{MarkdownDescription: "AV Policy name", Computed: true}, "then_do": schema.ListNestedAttribute{MarkdownDescription: "AND-joined list of instantiated actions to execute when this policy triggers", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"action": schema.StringAttribute{MarkdownDescription: "reference to a pre-defined action template ", Computed: true}, "params": schema.ListNestedAttribute{MarkdownDescription: "key/value criteria parameters to instantiate the referenced template", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"key": schema.StringAttribute{Computed: true}, "value": schema.StringAttribute{Computed: true}}}}}}}, "when_condition": schema.ListNestedAttribute{MarkdownDescription: "AND-joined list of instantiated criteria conditions to trigger this policy", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"condition": schema.StringAttribute{MarkdownDescription: "reference to a pre-defined condition template ", Computed: true}, "params": schema.ListNestedAttribute{MarkdownDescription: "key/value criteria parameters to instantiate the referenced template", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"key": schema.StringAttribute{Computed: true}, "value": schema.StringAttribute{Computed: true}}}}}}}}}}}}
}

// Read fetches remote state into the data source model.
func (d *GetAvisiPoliciesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetAvisiPoliciesDataSourceModel
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
func (d *GetAvisiPoliciesDataSource) readListRemote(ctx context.Context, config *GetAvisiPoliciesDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/avisi/policies"
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_avisi_policies", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_avisi_policies", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["avPolicies"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_avisi_policies", fmt.Sprintf("Could not decode list page: missing %q array", "avPolicies"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_avisi_policies", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetAvisiPoliciesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
