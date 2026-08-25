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
var _ action.Action = (*UpdateHeaderStripAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*UpdateHeaderStripAction)(nil)

// UpdateHeaderStripAction is the generated Terraform action implementation.
type UpdateHeaderStripAction struct {
	client *client.Client
}

// UpdateHeaderStripActionModel describes the action configuration shape.
type UpdateHeaderStripActionModel struct {
	AddLabels    types.List   `tfsdk:"add_labels" json:"addLabels"`
	BodyBoxId    types.String `tfsdk:"body_box_id" json:"boxId"`
	BoxId        types.String `tfsdk:"box_id"`
	ClusterId    types.String `tfsdk:"cluster_id"`
	DeleteLabels types.List   `tfsdk:"delete_labels" json:"deleteLabels"`
}

// NewUpdateHeaderStripAction returns a new instance of the generated action.
func NewUpdateHeaderStripAction() action.Action {
	return &UpdateHeaderStripAction{}
}

// Metadata returns the action type name.
func (r *UpdateHeaderStripAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_update_header_strip"
}

// Schema returns the action schema.
func (r *UpdateHeaderStripAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Update header strip for target box", Attributes: map[string]schema.Attribute{"add_labels": schema.ListAttribute{MarkdownDescription: "mpls ids, valid and required. Range can be specified. Example:1..200 ", Optional: true, ElementType: types.StringType}, "body_box_id": schema.StringAttribute{MarkdownDescription: "device box id. valid range 1 - 64.", Required: true}, "box_id": schema.StringAttribute{MarkdownDescription: "device box id. valid range 1 - 64.", Required: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target cluster ID.", Required: true}, "delete_labels": schema.ListAttribute{MarkdownDescription: "mpls ids, valid and required. Range can be specified. Example:1..200 ", Optional: true, ElementType: types.StringType}}}
}

// Invoke executes the action against the remote API.
func (r *UpdateHeaderStripAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config UpdateHeaderStripActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *UpdateHeaderStripAction) invokeRemote(ctx context.Context, config *UpdateHeaderStripActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/headerStrip/{boxId}/updateLabels"
	reqPath = strings.ReplaceAll(reqPath, "{boxId}", url.PathEscape(config.BoxId.ValueString()))
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_header_strip", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_header_strip", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_header_strip", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_header_strip", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_header_strip", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_header_strip", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_header_strip", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_header_strip", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_header_strip", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_header_strip", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_header_strip", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_update_header_strip", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_header_strip", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *UpdateHeaderStripAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
