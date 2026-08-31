package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)
import (
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*PasteRulesAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*PasteRulesAction)(nil)

// PasteRulesAction is the generated Terraform action implementation.
type PasteRulesAction struct {
	client *client.Client
}

// PasteRulesActionModel describes the action configuration shape.
type PasteRulesActionModel struct {
	Configs      types.List   `tfsdk:"configs"`
	Option       types.String `tfsdk:"option"`
	RuleCategory types.String `tfsdk:"rule_category" json:"ruleCategory"`
	RuleType     types.String `tfsdk:"rule_type" json:"ruleType"`
}

// NewPasteRulesAction returns a new instance of the generated action.
func NewPasteRulesAction() action.Action {
	return &PasteRulesAction{}
}

// Metadata returns the action type name.
func (r *PasteRulesAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_paste_rules"
}

// Schema returns the action schema.
func (r *PasteRulesAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Paste rules from the user-specific clipboard into one or more target traffic flow policies.\nSupports duplicate detection with two options:\n- DETECT_DUPLICATES: Returns duplicate rules without saving (409 Conflict)\n- PROCEED_WITHOUT_DUPLICATES: Saves only non-duplicate rules\nUses optimistic locking for concurrent conflict detection.", Attributes: map[string]schema.Attribute{"configs": schema.ListNestedAttribute{MarkdownDescription: "List of paste configurations", Required: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"flow_alias": schema.StringAttribute{MarkdownDescription: "Flow alias (required for APPLICATION category)", Optional: true}, "policy_id_or_alias": schema.StringAttribute{MarkdownDescription: "Policy ID or alias to which to paste rules", Required: true}, "policy_updated_time": schema.Int64Attribute{MarkdownDescription: "Timestamp of the policy version for optimistic locking", Required: true}, "source_and_rule_alias": schema.StringAttribute{MarkdownDescription: "Source and rule alias (required for SOURCE category)", Optional: true}, "sub_flow_alias": schema.StringAttribute{MarkdownDescription: "Sub-flow alias (required for APPLICATION category)", Optional: true}}}}, "option": schema.StringAttribute{MarkdownDescription: "Option for handling duplicate rules during paste operation", Optional: true}, "rule_category": schema.StringAttribute{MarkdownDescription: "Category of rules to copy/paste", Required: true}, "rule_type": schema.StringAttribute{MarkdownDescription: "Type of application rule", Required: true}}}
}

// Invoke executes the action against the remote API.
func (r *PasteRulesAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config PasteRulesActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *PasteRulesAction) invokeRemote(ctx context.Context, config *PasteRulesActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/traffic-flows/pasteRules"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_paste_rules", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_paste_rules", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_paste_rules", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_paste_rules", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_paste_rules", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_paste_rules", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_paste_rules", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_paste_rules", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_paste_rules", "Conflict - Duplicate rules detected (when using DETECT_DUPLICATES option) or policy has been modified")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_paste_rules", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_paste_rules", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_paste_rules", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_paste_rules", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *PasteRulesAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
