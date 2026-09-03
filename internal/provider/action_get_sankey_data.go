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
var _ action.Action = (*GetSankeyDataAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*GetSankeyDataAction)(nil)

// GetSankeyDataAction is the generated Terraform action implementation.
type GetSankeyDataAction struct {
	client *client.Client
}

// GetSankeyDataActionModel describes the action configuration shape.
type GetSankeyDataActionModel struct {
	NodesFilter types.List   `tfsdk:"nodes_filter" json:"nodesFilter"`
	TagsFilter  types.Object `tfsdk:"tags_filter" json:"tagsFilter"`
}

// NewGetSankeyDataAction returns a new instance of the generated action.
func NewGetSankeyDataAction() action.Action {
	return &GetSankeyDataAction{}
}

// Metadata returns the action type name.
func (r *GetSankeyDataAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_sankey_data"
}

// Schema returns the action schema.
func (r *GetSankeyDataAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Response for sankey representation", Attributes: map[string]schema.Attribute{"nodes_filter": schema.ListNestedAttribute{Optional: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"placement_tag_value": schema.StringAttribute{Optional: true}, "sort": schema.ListNestedAttribute{Optional: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"field": schema.StringAttribute{Optional: true}, "order": schema.StringAttribute{Optional: true}}}}}}}, "tags_filter": schema.SingleNestedAttribute{Required: true, Attributes: map[string]schema.Attribute{"parent_tags": schema.ListNestedAttribute{MarkdownDescription: "Parent hierarchy information for the requested sankey", Optional: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"tag_key": schema.StringAttribute{Optional: true}, "tag_value": schema.StringAttribute{Optional: true}}}}, "placement_tag_alias": schema.StringAttribute{MarkdownDescription: "Configured placement tag alias", Required: true}, "tag_key": schema.StringAttribute{MarkdownDescription: "Sankey tag key", Required: true}, "tag_value": schema.StringAttribute{MarkdownDescription: "Sankey tag value", Required: true}}}}}
}

// Invoke executes the action against the remote API.
func (r *GetSankeyDataAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config GetSankeyDataActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *GetSankeyDataAction) invokeRemote(ctx context.Context, config *GetSankeyDataActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/topoviz/sankey"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_get_sankey_data", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_get_sankey_data", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_get_sankey_data", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_get_sankey_data", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200 || httpResp.StatusCode == 204) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_get_sankey_data", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_get_sankey_data", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_get_sankey_data", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_get_sankey_data", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_get_sankey_data", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_get_sankey_data", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_get_sankey_data", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_get_sankey_data", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_get_sankey_data", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *GetSankeyDataAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
