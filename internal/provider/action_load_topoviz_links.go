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
var _ action.Action = (*LoadTopovizLinksAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*LoadTopovizLinksAction)(nil)

// LoadTopovizLinksAction is the generated Terraform action implementation.
type LoadTopovizLinksAction struct {
	client *client.Client
}

// LoadTopovizLinksActionModel describes the action configuration shape.
type LoadTopovizLinksActionModel struct {
	DestinationHierarchicalTagsFilter types.List   `tfsdk:"destination_hierarchical_tags_filter" json:"destinationHierarchicalTagsFilter"`
	DestinationPlacementTagFilter     types.Object `tfsdk:"destination_placement_tag_filter" json:"destinationPlacementTagFilter"`
	DestinationSubGroupFilter         types.Object `tfsdk:"destination_sub_group_filter" json:"destinationSubGroupFilter"`
	Endpoint1GlobalNodeId             types.String `tfsdk:"endpoint1_global_node_id" json:"endpoint1GlobalNodeId"`
	Endpoint2GlobalNodeId             types.String `tfsdk:"endpoint2_global_node_id" json:"endpoint2GlobalNodeId"`
	LinksFilter                       types.Object `tfsdk:"links_filter" json:"linksFilter"`
	SourceHierarchicalTagsFilter      types.List   `tfsdk:"source_hierarchical_tags_filter" json:"sourceHierarchicalTagsFilter"`
	SourcePlacementTagFilter          types.Object `tfsdk:"source_placement_tag_filter" json:"sourcePlacementTagFilter"`
	SourceSubGroupFilter              types.Object `tfsdk:"source_sub_group_filter" json:"sourceSubGroupFilter"`
}

// NewLoadTopovizLinksAction returns a new instance of the generated action.
func NewLoadTopovizLinksAction() action.Action {
	return &LoadTopovizLinksAction{}
}

// Metadata returns the action type name.
func (r *LoadTopovizLinksAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_topoviz_links"
}

// Schema returns the action schema.
func (r *LoadTopovizLinksAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "load links from given nodes or nodeGroups", Attributes: map[string]schema.Attribute{"destination_hierarchical_tags_filter": schema.ListNestedAttribute{MarkdownDescription: "List of tag key and tag values which specifies the hierarchy of destination node group", Optional: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"tag_key": schema.StringAttribute{Optional: true}, "tag_value": schema.StringAttribute{Optional: true}}}}, "destination_placement_tag_filter": schema.SingleNestedAttribute{MarkdownDescription: "Tag key and tag values which specifies the placement of destination node group in sankey view", Optional: true, Attributes: map[string]schema.Attribute{"tag_key": schema.StringAttribute{Optional: true}, "tag_value": schema.StringAttribute{Optional: true}}}, "destination_sub_group_filter": schema.SingleNestedAttribute{MarkdownDescription: "Property name and value based on which the nodes are subGrouped", Optional: true, Attributes: map[string]schema.Attribute{"tag_key": schema.StringAttribute{Optional: true}, "tag_value": schema.StringAttribute{Optional: true}}}, "endpoint1_global_node_id": schema.StringAttribute{Optional: true}, "endpoint2_global_node_id": schema.StringAttribute{Optional: true}, "links_filter": schema.SingleNestedAttribute{MarkdownDescription: "Property name and values based on which the link are to be filtered", Optional: true, Attributes: map[string]schema.Attribute{"link_type": schema.ListAttribute{MarkdownDescription: "Type of topology links - cascade, circuit, stack, tool, network", Optional: true, ElementType: types.StringType}}}, "source_hierarchical_tags_filter": schema.ListNestedAttribute{MarkdownDescription: "List of tag key and tag values which specifies the hierarchy of source node group", Optional: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"tag_key": schema.StringAttribute{Optional: true}, "tag_value": schema.StringAttribute{Optional: true}}}}, "source_placement_tag_filter": schema.SingleNestedAttribute{MarkdownDescription: "Tag key and tag values which specifies the placement of source node group in sankey view", Optional: true, Attributes: map[string]schema.Attribute{"tag_key": schema.StringAttribute{Optional: true}, "tag_value": schema.StringAttribute{Optional: true}}}, "source_sub_group_filter": schema.SingleNestedAttribute{MarkdownDescription: "Property name and value based on which the nodes are subGrouped", Optional: true, Attributes: map[string]schema.Attribute{"tag_key": schema.StringAttribute{Optional: true}, "tag_value": schema.StringAttribute{Optional: true}}}}}
}

// Invoke executes the action against the remote API.
func (r *LoadTopovizLinksAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config LoadTopovizLinksActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *LoadTopovizLinksAction) invokeRemote(ctx context.Context, config *LoadTopovizLinksActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/topoviz/links"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_load_topoviz_links", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_load_topoviz_links", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_load_topoviz_links", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_load_topoviz_links", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200 || httpResp.StatusCode == 204) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_load_topoviz_links", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_load_topoviz_links", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_load_topoviz_links", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_load_topoviz_links", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_load_topoviz_links", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_load_topoviz_links", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_load_topoviz_links", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_load_topoviz_links", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_load_topoviz_links", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *LoadTopovizLinksAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
