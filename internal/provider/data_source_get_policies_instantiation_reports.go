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
	_ datasource.DataSource              = (*GetPoliciesInstantiationReportsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetPoliciesInstantiationReportsDataSource)(nil)
)

// GetPoliciesInstantiationReportsDataSource is the generated Terraform data source implementation.
type GetPoliciesInstantiationReportsDataSource struct {
	client *client.Client
}

// GetPoliciesInstantiationReportsDataSourceModel describes the data source state shape.
type GetPoliciesInstantiationReportsDataSourceModel struct {
	Items     types.List   `tfsdk:"items"`
	Page      types.String `tfsdk:"page"`
	PolicyId  types.String `tfsdk:"policy_id" json:"policyId"`
	Since     types.String `tfsdk:"since"`
	Sort      types.String `tfsdk:"sort"`
	TimeRange types.String `tfsdk:"time_range" json:"timeRange"`
}

// NewGetPoliciesInstantiationReportsDataSource returns a new instance of the generated data source.
func NewGetPoliciesInstantiationReportsDataSource() datasource.DataSource {
	return &GetPoliciesInstantiationReportsDataSource{}
}

// Metadata returns the data source type name.
func (d *GetPoliciesInstantiationReportsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_policies_instantiation_reports"
}

// Schema returns the data source schema.
func (d *GetPoliciesInstantiationReportsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get Active Visibility Trigger Reports", Attributes: map[string]schema.Attribute{"items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"action_status": schema.ListNestedAttribute{MarkdownDescription: "Execution report for all associated Policy Actions. On Policy 'success', could be left out since success of every action is implied.", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"action_id": schema.StringAttribute{MarkdownDescription: "Action ID in the parent Policy", Computed: true}, "action_type": schema.StringAttribute{MarkdownDescription: "Action Type (Source Action Template). Included for readability to provide Action context", Computed: true}, "failure_reasons": schema.ListAttribute{MarkdownDescription: "List of the reasons that resulted in the failed Action execution. Empty or omitted on success", Computed: true, ElementType: types.StringType}, "outcome": schema.StringAttribute{Computed: true}}}}, "condition_status": schema.ListNestedAttribute{MarkdownDescription: "Triggering Conditions report. If none of the conditions are parameterizable, this status field can be omitted", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"condition_id": schema.StringAttribute{MarkdownDescription: "Condition ID in the parent Policy", Computed: true}, "params": schema.ListNestedAttribute{MarkdownDescription: "key/value criteria parameters under which this Condition was satisfied for a given Policy triggering instance. This map should have an entry for each Parameter required for the corresponding condition type. If corresponding Condition does not take parameters, this field can be omitted", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"key": schema.StringAttribute{Computed: true}, "value": schema.StringAttribute{Computed: true}}}}}}}, "outcome": schema.StringAttribute{MarkdownDescription: " 'success' is reported when all actions completed successfully, otherwise 'failure' is reported", Computed: true}, "policy_id": schema.StringAttribute{MarkdownDescription: "reference to the triggering Policy", Computed: true}, "policy_name": schema.StringAttribute{MarkdownDescription: "triggering Policy name. Included for readability to provide Policy context", Computed: true}, "trigger_time": schema.StringAttribute{MarkdownDescription: "policy trigger time. In ISO-8601 date format 'yyyy-MM-dd'T'HH:mm:ssZ'", Computed: true}}}}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned", Optional: true}, "policy_id": schema.StringAttribute{MarkdownDescription: "Target Policy Id. If left out, reports for all AV Policies are returned", Optional: true}, "since": schema.StringAttribute{MarkdownDescription: "Reference timestamps to include trigger reports from. In ISO-8601 date format 'yyyy-MM-ddTHH:mm:ssZ'. Mutually exclusive with 'timeRange' query parameter", Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}, "time_range": schema.StringAttribute{MarkdownDescription: "Time window for which to include trigger reports. Mutually exclusive with 'since' query paramater", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetPoliciesInstantiationReportsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetPoliciesInstantiationReportsDataSourceModel
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
func (d *GetPoliciesInstantiationReportsDataSource) readListRemote(ctx context.Context, config *GetPoliciesInstantiationReportsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/avisi/policiesInstantiations"
	params := url.Values{}
	if !config.PolicyId.IsNull() {
		params.Set("policyId", config.PolicyId.ValueString())
	}
	if !config.TimeRange.IsNull() {
		params.Set("timeRange", config.TimeRange.ValueString())
	}
	if !config.Since.IsNull() {
		params.Set("since", config.Since.ValueString())
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_policies_instantiation_reports", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_policies_instantiation_reports", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["avPolicyTriggerReports"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_policies_instantiation_reports", fmt.Sprintf("Could not decode list page: missing %q array", "avPolicyTriggerReports"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_policies_instantiation_reports", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetPoliciesInstantiationReportsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
