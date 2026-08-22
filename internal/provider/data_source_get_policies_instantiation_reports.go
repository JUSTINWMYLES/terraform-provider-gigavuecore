package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)
import (
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
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
	AvPolicyTriggerReports types.List   `tfsdk:"av_policy_trigger_reports" json:"avPolicyTriggerReports"`
	Context                types.Object `tfsdk:"context"`
	Page                   types.String `tfsdk:"page"`
	PolicyId               types.String `tfsdk:"policy_id" json:"policyId"`
	Since                  types.String `tfsdk:"since"`
	Sort                   types.String `tfsdk:"sort"`
	TimeRange              types.String `tfsdk:"time_range" json:"timeRange"`
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
	resp.Schema = schema.Schema{MarkdownDescription: "Get Active Visibility Trigger Reports", Attributes: map[string]schema.Attribute{"av_policy_trigger_reports": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"action_status": schema.ListNestedAttribute{MarkdownDescription: "Execution report for all associated Policy Actions. On Policy 'success', could be left out since success of every action is implied.", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"action_id": schema.StringAttribute{MarkdownDescription: "Action ID in the parent Policy", Computed: true}, "action_type": schema.StringAttribute{MarkdownDescription: "Action Type (Source Action Template). Included for readability to provide Action context", Computed: true}, "failure_reasons": schema.ListAttribute{MarkdownDescription: "List of the reasons that resulted in the failed Action execution. Empty or omitted on success", Computed: true, ElementType: types.StringType}, "outcome": schema.StringAttribute{Computed: true}}}}, "condition_status": schema.ListNestedAttribute{MarkdownDescription: "Triggering Conditions report. If none of the conditions are parameterizable, this status field can be omitted", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"condition_id": schema.StringAttribute{MarkdownDescription: "Condition ID in the parent Policy", Computed: true}, "params": schema.ListNestedAttribute{MarkdownDescription: "key/value criteria parameters under which this Condition was satisfied for a given Policy triggering instance. This map should have an entry for each Parameter required for the corresponding condition type. If corresponding Condition does not take parameters, this field can be omitted", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"key": schema.StringAttribute{Computed: true}, "value": schema.StringAttribute{Computed: true}}}}}}}, "outcome": schema.StringAttribute{MarkdownDescription: " 'success' is reported when all actions completed successfully, otherwise 'failure' is reported", Computed: true}, "policy_id": schema.StringAttribute{MarkdownDescription: "reference to the triggering Policy", Computed: true}, "policy_name": schema.StringAttribute{MarkdownDescription: "triggering Policy name. Included for readability to provide Policy context", Computed: true}, "trigger_time": schema.StringAttribute{MarkdownDescription: "policy trigger time. In ISO-8601 date format 'yyyy-MM-dd'T'HH:mm:ssZ'", Computed: true}}}}, "context": schema.SingleNestedAttribute{MarkdownDescription: "Gigamon query result context", Computed: true, Attributes: map[string]schema.Attribute{"page_no": schema.Int64Attribute{MarkdownDescription: "page number of the returned result set", Computed: true}, "page_size": schema.Int64Attribute{MarkdownDescription: "page size of the returned result set", Computed: true}, "sort": schema.ListAttribute{MarkdownDescription: "sorting info of the returned result set. list of fields in the array indicate sorting order", Computed: true, ElementType: types.StringType}, "total_items": schema.Int64Attribute{MarkdownDescription: "total number of items in the queried entity type", Computed: true}}}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned", Optional: true}, "policy_id": schema.StringAttribute{MarkdownDescription: "Target Policy Id. If left out, reports for all AV Policies are returned", Optional: true}, "since": schema.StringAttribute{MarkdownDescription: "Reference timestamps to include trigger reports from. In ISO-8601 date format 'yyyy-MM-ddTHH:mm:ssZ'. Mutually exclusive with 'timeRange' query parameter", Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}, "time_range": schema.StringAttribute{MarkdownDescription: "Time window for which to include trigger reports. Mutually exclusive with 'since' query paramater", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetPoliciesInstantiationReportsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetPoliciesInstantiationReportsDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	d.readRemote(ctx, &config, resp)
	if resp.Diagnostics.HasError() {
		return
	}
	if config.Page.IsNull() {
		resp.Diagnostics.AddWarning("Single-page result", fmt.Sprintf("This data source reads a single page of a paginated API endpoint and does not aggregate results across pages. The \"page\" argument is unset, so the default page is returned; set it to retrieve a different page."))
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}

// readRemote performs the read HTTP exchange and decodes the response into config. Extracted from Read so the request/response logic is unit-testable without a tfsdk.Config.
func (d *GetPoliciesInstantiationReportsDataSource) readRemote(ctx context.Context, config *GetPoliciesInstantiationReportsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/avisi/policiesInstantiations"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_policies_instantiation_reports", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !config.PolicyId.IsNull() {
		query.Set("policyId", config.PolicyId.ValueString())
	}
	if !config.TimeRange.IsNull() {
		query.Set("timeRange", config.TimeRange.ValueString())
	}
	if !config.Since.IsNull() {
		query.Set("since", config.Since.ValueString())
	}
	if !config.Page.IsNull() {
		query.Set("page", config.Page.ValueString())
	}
	if !config.Sort.IsNull() {
		query.Set("sort", config.Sort.ValueString())
	}
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_policies_instantiation_reports", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_policies_instantiation_reports", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_policies_instantiation_reports", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_policies_instantiation_reports", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_policies_instantiation_reports", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_policies_instantiation_reports", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_policies_instantiation_reports", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_policies_instantiation_reports", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_policies_instantiation_reports", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_policies_instantiation_reports", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_policies_instantiation_reports", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_policies_instantiation_reports", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
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
