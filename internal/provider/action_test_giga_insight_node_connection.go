package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)
import (
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*TestGigaInsightNodeConnectionAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*TestGigaInsightNodeConnectionAction)(nil)

// TestGigaInsightNodeConnectionAction is the generated Terraform action implementation.
type TestGigaInsightNodeConnectionAction struct {
	client *client.Client
}

// TestGigaInsightNodeConnectionActionModel describes the action configuration shape.
type TestGigaInsightNodeConnectionActionModel struct {
	BedrockConfig types.Object `tfsdk:"bedrock_config" json:"bedrockConfig"`
	GoogleConfig  types.Object `tfsdk:"google_config" json:"googleConfig"`
	NodeId        types.String `tfsdk:"node_id"`
	OpenAiConfig  types.Object `tfsdk:"open_ai_config" json:"openAIConfig"`
	PrivateConfig types.Object `tfsdk:"private_config" json:"privateConfig"`
	Provider      types.String `tfsdk:"provider_" json:"provider"`
	ProxyUrl      types.String `tfsdk:"proxy_url" json:"proxyUrl"`
}

// NewTestGigaInsightNodeConnectionAction returns a new instance of the generated action.
func NewTestGigaInsightNodeConnectionAction() action.Action {
	return &TestGigaInsightNodeConnectionAction{}
}

// Metadata returns the action type name.
func (r *TestGigaInsightNodeConnectionAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_test_giga_insight_node_connection"
}

// Schema returns the action schema.
func (r *TestGigaInsightNodeConnectionAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Test LLM provider connection via a GigaInsight Node", Attributes: map[string]schema.Attribute{"bedrock_config": schema.SingleNestedAttribute{MarkdownDescription: "AWS Bedrock provider configuration", Optional: true, Attributes: map[string]schema.Attribute{"access_key_id": schema.StringAttribute{MarkdownDescription: "AWS access key ID", Required: true}, "region": schema.StringAttribute{MarkdownDescription: "AWS region", Required: true}, "secret_access_key": schema.StringAttribute{MarkdownDescription: "AWS secret access key", Required: true}}}, "google_config": schema.SingleNestedAttribute{MarkdownDescription: "Google Vertex AI provider configuration", Optional: true, Attributes: map[string]schema.Attribute{"project_id": schema.StringAttribute{MarkdownDescription: "Google Cloud project ID", Required: true}, "region": schema.StringAttribute{MarkdownDescription: "Google Cloud region", Required: true}, "service_account_key_json": schema.DynamicAttribute{MarkdownDescription: "Google service account key JSON", Required: true}}}, "node_id": schema.StringAttribute{Required: true}, "open_ai_config": schema.SingleNestedAttribute{MarkdownDescription: "OpenAI provider configuration", Optional: true, Attributes: map[string]schema.Attribute{"api_key": schema.StringAttribute{MarkdownDescription: "OpenAI API key", Required: true}}}, "private_config": schema.SingleNestedAttribute{MarkdownDescription: "Private (Azure AI Foundry) provider configuration", Optional: true, Attributes: map[string]schema.Attribute{"api_key": schema.StringAttribute{MarkdownDescription: "API key for the private endpoint", Required: true}, "endpoint_url": schema.StringAttribute{MarkdownDescription: "Private LLM endpoint URL", Required: true}}}, "provider_": schema.StringAttribute{MarkdownDescription: "The LLM provider type", Required: true}, "proxy_url": schema.StringAttribute{MarkdownDescription: "Optional proxy URL for the connection", Optional: true}}}
}

// Invoke executes the action against the remote API.
func (r *TestGigaInsightNodeConnectionAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config TestGigaInsightNodeConnectionActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *TestGigaInsightNodeConnectionAction) invokeRemote(ctx context.Context, config *TestGigaInsightNodeConnectionActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/config/insights/nodes/{nodeId}/testconnection"
	reqPath = strings.ReplaceAll(reqPath, "{nodeId}", url.PathEscape(config.NodeId.ValueString()))
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_test_giga_insight_node_connection", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_test_giga_insight_node_connection", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_test_giga_insight_node_connection", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_test_giga_insight_node_connection", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_test_giga_insight_node_connection", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_test_giga_insight_node_connection", "Not Authenticated. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_test_giga_insight_node_connection", "Internal Server Error. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_test_giga_insight_node_connection", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_test_giga_insight_node_connection", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *TestGigaInsightNodeConnectionAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Action Configure Type", fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData))
		return
	}
	r.client = c
}
