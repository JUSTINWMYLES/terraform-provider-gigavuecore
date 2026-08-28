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
var _ action.Action = (*UpdateKeystorePreferenceAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*UpdateKeystorePreferenceAction)(nil)

// UpdateKeystorePreferenceAction is the generated Terraform action implementation.
type UpdateKeystorePreferenceAction struct {
	client *client.Client
}

// UpdateKeystorePreferenceActionModel describes the action configuration shape.
type UpdateKeystorePreferenceActionModel struct {
	AutoDelete    types.Bool   `tfsdk:"auto_delete" json:"autoDelete"`
	AutoEnable    types.Bool   `tfsdk:"auto_enable" json:"autoEnable"`
	AutoPurge     types.Bool   `tfsdk:"auto_purge" json:"autoPurge"`
	BodyClusterId types.String `tfsdk:"body_cluster_id" json:"clusterId"`
	ClusterId     types.String `tfsdk:"cluster_id"`
	MaxKeys       types.Int64  `tfsdk:"max_keys" json:"maxKeys"`
	RetentionTime types.Int64  `tfsdk:"retention_time" json:"retentionTime"`
}

// NewUpdateKeystorePreferenceAction returns a new instance of the generated action.
func NewUpdateKeystorePreferenceAction() action.Action {
	return &UpdateKeystorePreferenceAction{}
}

// Metadata returns the action type name.
func (r *UpdateKeystorePreferenceAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_update_keystore_preference"
}

// Schema returns the action schema.
func (r *UpdateKeystorePreferenceAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Set keystore preference", Attributes: map[string]schema.Attribute{"auto_delete": schema.BoolAttribute{Optional: true}, "auto_enable": schema.BoolAttribute{Optional: true}, "auto_purge": schema.BoolAttribute{Optional: true}, "body_cluster_id": schema.StringAttribute{Required: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "max_keys": schema.Int64Attribute{Optional: true}, "retention_time": schema.Int64Attribute{Optional: true}}}
}

// Invoke executes the action against the remote API.
func (r *UpdateKeystorePreferenceAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config UpdateKeystorePreferenceActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *UpdateKeystorePreferenceAction) invokeRemote(ctx context.Context, config *UpdateKeystorePreferenceActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/apps/keystore/preference"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_keystore_preference", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_keystore_preference", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_keystore_preference", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.BodyClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_update_keystore_preference", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_keystore_preference", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_keystore_preference", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_keystore_preference", "Access Denied. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_keystore_preference", "Entity Already Exists. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_update_keystore_preference", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_update_keystore_preference", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *UpdateKeystorePreferenceAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
