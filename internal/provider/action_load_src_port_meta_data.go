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
var _ action.Action = (*LoadSrcPortMetaDataAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*LoadSrcPortMetaDataAction)(nil)

// LoadSrcPortMetaDataAction is the generated Terraform action implementation.
type LoadSrcPortMetaDataAction struct {
	client *client.Client
}

// LoadSrcPortMetaDataActionModel describes the action configuration shape.
type LoadSrcPortMetaDataActionModel struct {
	SourceDetails types.List `tfsdk:"source_details" json:"sourceDetails"`
}

// NewLoadSrcPortMetaDataAction returns a new instance of the generated action.
func NewLoadSrcPortMetaDataAction() action.Action {
	return &LoadSrcPortMetaDataAction{}
}

// Metadata returns the action type name.
func (r *LoadSrcPortMetaDataAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_src_port_meta_data"
}

// Schema returns the action schema.
func (r *LoadSrcPortMetaDataAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Load source port metadata for traffic flows", Attributes: map[string]schema.Attribute{"source_details": schema.ListNestedAttribute{Optional: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{Required: true}, "components": schema.ListNestedAttribute{Optional: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"ids": schema.ListAttribute{Optional: true, ElementType: types.StringType}, "type": schema.StringAttribute{Optional: true}}}}}}}}}
}

// Invoke executes the action against the remote API.
func (r *LoadSrcPortMetaDataAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config LoadSrcPortMetaDataActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *LoadSrcPortMetaDataAction) invokeRemote(ctx context.Context, config *LoadSrcPortMetaDataActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/utils/port/metaData/traffic-flows"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_load_src_port_meta_data", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_load_src_port_meta_data", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_load_src_port_meta_data", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_load_src_port_meta_data", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_load_src_port_meta_data", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_load_src_port_meta_data", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_load_src_port_meta_data", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_load_src_port_meta_data", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_load_src_port_meta_data", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_load_src_port_meta_data", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_load_src_port_meta_data", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_load_src_port_meta_data", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_load_src_port_meta_data", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *LoadSrcPortMetaDataAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
