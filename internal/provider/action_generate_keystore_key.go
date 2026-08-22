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
var _ action.Action = (*GenerateKeystoreKeyAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*GenerateKeystoreKeyAction)(nil)

// GenerateKeystoreKeyAction is the generated Terraform action implementation.
type GenerateKeystoreKeyAction struct {
	client *client.Client
}

// GenerateKeystoreKeyActionModel describes the action configuration shape.
type GenerateKeystoreKeyActionModel struct {
	Alias      types.String `tfsdk:"alias"`
	ClusterId  types.String `tfsdk:"cluster_id"`
	CommonName types.String `tfsdk:"common_name" json:"commonName"`
	Country    types.String `tfsdk:"country"`
	Days       types.Int64  `tfsdk:"days"`
	HashType   types.String `tfsdk:"hash_type" json:"hashType"`
	Keysize    types.Int64  `tfsdk:"keysize"`
	OrgName    types.String `tfsdk:"org_name" json:"orgName"`
	OrgUnit    types.String `tfsdk:"org_unit" json:"orgUnit"`
	State      types.String `tfsdk:"state"`
}

// NewGenerateKeystoreKeyAction returns a new instance of the generated action.
func NewGenerateKeystoreKeyAction() action.Action {
	return &GenerateKeystoreKeyAction{}
}

// Metadata returns the action type name.
func (r *GenerateKeystoreKeyAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_generate_keystore_key"
}

// Schema returns the action schema.
func (r *GenerateKeystoreKeyAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "generate a self-signed key in the keystore", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{Required: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "common_name": schema.StringAttribute{Required: true}, "country": schema.StringAttribute{Optional: true}, "days": schema.Int64Attribute{Optional: true}, "hash_type": schema.StringAttribute{Optional: true}, "keysize": schema.Int64Attribute{MarkdownDescription: "ecdsa sizes: 256, 384, 521; rsa sizes: 1024, 2048, 4096", Optional: true}, "org_name": schema.StringAttribute{Required: true}, "org_unit": schema.StringAttribute{Optional: true}, "state": schema.StringAttribute{Optional: true}}}
}

// Invoke executes the action against the remote API.
func (r *GenerateKeystoreKeyAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config GenerateKeystoreKeyActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *GenerateKeystoreKeyAction) invokeRemote(ctx context.Context, config *GenerateKeystoreKeyActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/apps/keystore/keys/generate"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_generate_keystore_key", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_generate_keystore_key", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_generate_keystore_key", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_generate_keystore_key", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_generate_keystore_key", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_generate_keystore_key", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_generate_keystore_key", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_generate_keystore_key", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_generate_keystore_key", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_generate_keystore_key", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_generate_keystore_key", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_generate_keystore_key", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_generate_keystore_key", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *GenerateKeystoreKeyAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
