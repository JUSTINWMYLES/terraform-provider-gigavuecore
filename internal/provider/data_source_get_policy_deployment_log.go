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
	_ datasource.DataSource              = (*GetPolicyDeploymentLogDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetPolicyDeploymentLogDataSource)(nil)
)

// GetPolicyDeploymentLogDataSource is the generated Terraform data source implementation.
type GetPolicyDeploymentLogDataSource struct {
	client *client.Client
}

// GetPolicyDeploymentLogDataSourceModel describes the data source state shape.
type GetPolicyDeploymentLogDataSourceModel struct {
	Message  types.String `tfsdk:"message"`
	Name     types.String `tfsdk:"name"`
	PolicyId types.String `tfsdk:"policy_id" json:"policyId"`
}

// NewGetPolicyDeploymentLogDataSource returns a new instance of the generated data source.
func NewGetPolicyDeploymentLogDataSource() datasource.DataSource {
	return &GetPolicyDeploymentLogDataSource{}
}

// Metadata returns the data source type name.
func (d *GetPolicyDeploymentLogDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_policy_deployment_log"
}

// Schema returns the data source schema.
func (d *GetPolicyDeploymentLogDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get policy Deployment log", Attributes: map[string]schema.Attribute{"message": schema.StringAttribute{Computed: true}, "name": schema.StringAttribute{MarkdownDescription: "policy name", Required: true}, "policy_id": schema.StringAttribute{Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetPolicyDeploymentLogDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetPolicyDeploymentLogDataSourceModel
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
func (d *GetPolicyDeploymentLogDataSource) readRemote(ctx context.Context, config *GetPolicyDeploymentLogDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/intent/policies/{name}/log"
	reqPath = strings.ReplaceAll(reqPath, "{name}", url.PathEscape(config.Name.ValueString()))
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_policy_deployment_log", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_policy_deployment_log", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_policy_deployment_log", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		apiErr, err := client.NewAPIError(httpResp)
		if err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_policy_deployment_log", fmt.Sprintf("Could not read error response: %s", err))
			return
		}
		resp.Diagnostics.AddError("Error reading gigavuecore_get_policy_deployment_log", apiErr.Error())
		return
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_policy_deployment_log", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_policy_deployment_log", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetPolicyDeploymentLogDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
