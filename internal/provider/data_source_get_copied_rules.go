package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)
import (
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*GetCopiedRulesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetCopiedRulesDataSource)(nil)
)

// GetCopiedRulesDataSource is the generated Terraform data source implementation.
type GetCopiedRulesDataSource struct {
	client *client.Client
}

// GetCopiedRulesDataSourceModel describes the data source state shape.
type GetCopiedRulesDataSourceModel struct {
	Context      types.Object `tfsdk:"context"`
	CopiedRules  types.Object `tfsdk:"copied_rules" json:"copiedRules"`
	RuleCategory types.String `tfsdk:"rule_category" json:"ruleCategory"`
	RuleType     types.String `tfsdk:"rule_type" json:"ruleType"`
}

// NewGetCopiedRulesDataSource returns a new instance of the generated data source.
func NewGetCopiedRulesDataSource() datasource.DataSource {
	return &GetCopiedRulesDataSource{}
}

// Metadata returns the data source type name.
func (d *GetCopiedRulesDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_copied_rules"
}

// Schema returns the data source schema.
func (d *GetCopiedRulesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Retrieve previously copied rules from the user-specific clipboard filtered by rule category and type.\nReturns the complete set of copied rules matching the specified category and type.", Attributes: map[string]schema.Attribute{"context": schema.SingleNestedAttribute{MarkdownDescription: "Gigamon query result context", Computed: true, Attributes: map[string]schema.Attribute{"page_no": schema.Int64Attribute{MarkdownDescription: "page number of the returned result set", Computed: true}, "page_size": schema.Int64Attribute{MarkdownDescription: "page size of the returned result set", Computed: true}, "sort": schema.ListAttribute{MarkdownDescription: "sorting info of the returned result set. list of fields in the array indicate sorting order", Computed: true, ElementType: types.StringType}, "total_items": schema.Int64Attribute{MarkdownDescription: "total number of items in the queried entity type", Computed: true}}}, "copied_rules": schema.SingleNestedAttribute{Computed: true, Attributes: map[string]schema.Attribute{"created_time": schema.Int64Attribute{MarkdownDescription: "Timestamp when rules were copied", Computed: true}, "policy_id": schema.StringAttribute{MarkdownDescription: "MongoDB document ID", Computed: true}, "rule_category": schema.StringAttribute{MarkdownDescription: "Category of rules to copy/paste", Computed: true}, "rule_set": schema.DynamicAttribute{MarkdownDescription: "Set of copied rules", Computed: true}, "rule_type": schema.StringAttribute{MarkdownDescription: "Type of application rule", Computed: true}, "username": schema.StringAttribute{MarkdownDescription: "Username of the user who copied the rules", Computed: true}}}, "rule_category": schema.StringAttribute{MarkdownDescription: "Rule category (SOURCE or APPLICATION)", Required: true}, "rule_type": schema.StringAttribute{MarkdownDescription: "Rule type (MapSubType enum value)", Required: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetCopiedRulesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetCopiedRulesDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	d.readRemote(ctx, &config, resp)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}

// readRemote performs the read HTTP exchange and decodes the response into config. Extracted from Read so the request/response logic is unit-testable without a tfsdk.Config.
func (d *GetCopiedRulesDataSource) readRemote(ctx context.Context, config *GetCopiedRulesDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/traffic-flows/copyRules/ruleCategory/{ruleCategory}/ruleType/{ruleType}"
	reqPath = strings.ReplaceAll(reqPath, "{ruleCategory}", url.PathEscape(config.RuleCategory.ValueString()))
	reqPath = strings.ReplaceAll(reqPath, "{ruleType}", url.PathEscape(config.RuleType.ValueString()))
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_copied_rules", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_copied_rules", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_copied_rules", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_copied_rules", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_copied_rules", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_copied_rules", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_copied_rules", "Entity Not Found. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_copied_rules", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_copied_rules", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_copied_rules", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_copied_rules", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_copied_rules", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_copied_rules", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetCopiedRulesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
