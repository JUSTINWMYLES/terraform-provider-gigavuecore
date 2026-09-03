package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)
import (
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*RegisterGigaInsightNodeAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*RegisterGigaInsightNodeAction)(nil)

// RegisterGigaInsightNodeAction is the generated Terraform action implementation.
type RegisterGigaInsightNodeAction struct {
	client *client.Client
}

// RegisterGigaInsightNodeActionModel describes the action configuration shape.
type RegisterGigaInsightNodeActionModel struct {
	Ipv4Address types.String `tfsdk:"ipv4_address" json:"ipv4Address"`
	Ipv6Address types.String `tfsdk:"ipv6_address" json:"ipv6Address"`
}

// NewRegisterGigaInsightNodeAction returns a new instance of the generated action.
func NewRegisterGigaInsightNodeAction() action.Action {
	return &RegisterGigaInsightNodeAction{}
}

// Metadata returns the action type name.
func (r *RegisterGigaInsightNodeAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_register_giga_insight_node"
}

// Schema returns the action schema.
func (r *RegisterGigaInsightNodeAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Register a GigaInsight Node with FM", Attributes: map[string]schema.Attribute{"ipv4_address": schema.StringAttribute{Required: true}, "ipv6_address": schema.StringAttribute{Required: true}}}
}

// Invoke executes the action against the remote API.
func (r *RegisterGigaInsightNodeAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config RegisterGigaInsightNodeActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *RegisterGigaInsightNodeAction) invokeRemote(ctx context.Context, config *RegisterGigaInsightNodeActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/orchestrate/insights/registration"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_register_giga_insight_node", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_register_giga_insight_node", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_register_giga_insight_node", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_register_giga_insight_node", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_register_giga_insight_node", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_register_giga_insight_node", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_register_giga_insight_node", "Access Denied. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_register_giga_insight_node", "Conflict. Node already registered")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_register_giga_insight_node", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_register_giga_insight_node", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_register_giga_insight_node", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_register_giga_insight_node", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *RegisterGigaInsightNodeAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
