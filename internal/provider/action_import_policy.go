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
var _ action.Action = (*ImportPolicyAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*ImportPolicyAction)(nil)

// ImportPolicyAction is the generated Terraform action implementation.
type ImportPolicyAction struct {
	client *client.Client
}

// ImportPolicyActionModel describes the action configuration shape.
type ImportPolicyActionModel struct {
	CriteriaBindings             types.Dynamic `tfsdk:"criteria_bindings" json:"criteriaBindings"`
	Name                         types.String  `tfsdk:"name"`
	PacketTransformationBindings types.Dynamic `tfsdk:"packet_transformation_bindings" json:"packetTransformationBindings"`
	Priority                     types.Bool    `tfsdk:"priority"`
	Rules                        types.Dynamic `tfsdk:"rules"`
	SourceBindings               types.Dynamic `tfsdk:"source_bindings" json:"sourceBindings"`
	Sources                      types.List    `tfsdk:"sources"`
	Tags                         types.List    `tfsdk:"tags"`
}

// NewImportPolicyAction returns a new instance of the generated action.
func NewImportPolicyAction() action.Action {
	return &ImportPolicyAction{}
}

// Metadata returns the action type name.
func (r *ImportPolicyAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_import_policy"
}

// Schema returns the action schema.
func (r *ImportPolicyAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Import Policy", Attributes: map[string]schema.Attribute{"criteria_bindings": schema.DynamicAttribute{Optional: true}, "name": schema.StringAttribute{Optional: true}, "packet_transformation_bindings": schema.DynamicAttribute{Optional: true}, "priority": schema.BoolAttribute{Optional: true}, "rules": schema.DynamicAttribute{Optional: true}, "source_bindings": schema.DynamicAttribute{Optional: true}, "sources": schema.ListNestedAttribute{Optional: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"ports": schema.ListAttribute{Optional: true, ElementType: types.StringType}, "type": schema.StringAttribute{Optional: true}}}}, "tags": schema.ListNestedAttribute{Optional: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"key": schema.StringAttribute{Optional: true}, "values": schema.ListAttribute{Optional: true, ElementType: types.StringType}}}}}}
}

// Invoke executes the action against the remote API.
func (r *ImportPolicyAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config ImportPolicyActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *ImportPolicyAction) invokeRemote(ctx context.Context, config *ImportPolicyActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/intent/policies/import"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_import_policy", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_import_policy", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_import_policy", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_import_policy", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201) {
		apiErr, err := client.NewAPIError(httpResp)
		if err != nil {
			resp.Diagnostics.AddError("Error invoking gigavuecore_import_policy", fmt.Sprintf("Could not read error response: %s", err))
			return
		}
		resp.Diagnostics.AddError("Error invoking gigavuecore_import_policy", apiErr.Error())
		return
	}
}

// Configure stores the API client supplied by the provider.
func (r *ImportPolicyAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
