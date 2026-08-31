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
	_ datasource.DataSource              = (*LoadAlertPoliciesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadAlertPoliciesDataSource)(nil)
)

// LoadAlertPoliciesDataSource is the generated Terraform data source implementation.
type LoadAlertPoliciesDataSource struct {
	client *client.Client
}

// LoadAlertPoliciesDataSourceModel describes the data source state shape.
type LoadAlertPoliciesDataSourceModel struct {
	Enabled      types.Bool    `tfsdk:"enabled"`
	Items        types.Dynamic `tfsdk:"items"`
	Page         types.String  `tfsdk:"page"`
	PolicyName   types.String  `tfsdk:"policy_name" json:"policyName"`
	ResourceType types.String  `tfsdk:"resource_type" json:"resourceType"`
	Sort         types.String  `tfsdk:"sort"`
}

// NewLoadAlertPoliciesDataSource returns a new instance of the generated data source.
func NewLoadAlertPoliciesDataSource() datasource.DataSource {
	return &LoadAlertPoliciesDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadAlertPoliciesDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_alert_policies"
}

// Schema returns the data source schema.
func (d *LoadAlertPoliciesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Alert policy listing", Attributes: map[string]schema.Attribute{"enabled": schema.BoolAttribute{Optional: true}, "items": schema.DynamicAttribute{MarkdownDescription: "All Alert Policy Configurations", Computed: true}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned", Optional: true}, "policy_name": schema.StringAttribute{MarkdownDescription: "Name of the alert policy", Optional: true}, "resource_type": schema.StringAttribute{Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadAlertPoliciesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadAlertPoliciesDataSourceModel
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
func (d *LoadAlertPoliciesDataSource) readListRemote(ctx context.Context, config *LoadAlertPoliciesDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/fm/alert-policies"
	params := url.Values{}
	if !config.PolicyName.IsNull() {
		params.Set("policyName", config.PolicyName.ValueString())
	}
	if !config.ResourceType.IsNull() {
		params.Set("resourceType", config.ResourceType.ValueString())
	}
	if !config.Enabled.IsNull() {
		params.Set("enabled", strconv.FormatBool(config.Enabled.ValueBool()))
	}
	if !config.Page.IsNull() {
		params.Set("page", config.Page.ValueString())
	}
	if !config.Sort.IsNull() {
		params.Set("sort", config.Sort.ValueString())
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
		resp.Diagnostics.AddError("Error reading gigavuecore_load_alert_policies", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_alert_policies", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["alertPolicies"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_alert_policies", fmt.Sprintf("Could not decode list page: missing %q array", "alertPolicies"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_alert_policies", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadAlertPoliciesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
