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
var _ action.Action = (*GenerateGigasmartDumpAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*GenerateGigasmartDumpAction)(nil)

// GenerateGigasmartDumpAction is the generated Terraform action implementation.
type GenerateGigasmartDumpAction struct {
	client *client.Client
}

// GenerateGigasmartDumpActionModel describes the action configuration shape.
type GenerateGigasmartDumpActionModel struct {
	ClusterId types.String `tfsdk:"cluster_id" json:"clusterId"`
	EngineIds types.List   `tfsdk:"engine_ids" json:"engineIds"`
	Hostname  types.String `tfsdk:"hostname"`
}

// NewGenerateGigasmartDumpAction returns a new instance of the generated action.
func NewGenerateGigasmartDumpAction() action.Action {
	return &GenerateGigasmartDumpAction{}
}

// Metadata returns the action type name.
func (r *GenerateGigasmartDumpAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_generate_gigasmart_dump"
}

// Schema returns the action schema.
func (r *GenerateGigasmartDumpAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Generate Gigasmart Dump", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Cluster Id for which dump is initiated", Required: true}, "engine_ids": schema.ListAttribute{MarkdownDescription: "list of the Gigasmart engine ports for which dump is initiated", Required: true, ElementType: types.StringType}, "hostname": schema.StringAttribute{MarkdownDescription: "Hostname for which dump is initiated", Required: true}}}
}

// Invoke executes the action against the remote API.
func (r *GenerateGigasmartDumpAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config GenerateGigasmartDumpActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *GenerateGigasmartDumpAction) invokeRemote(ctx context.Context, config *GenerateGigasmartDumpActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/system/gsDump"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_generate_gigasmart_dump", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_generate_gigasmart_dump", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPost, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_generate_gigasmart_dump", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_generate_gigasmart_dump", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_generate_gigasmart_dump", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_generate_gigasmart_dump", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_generate_gigasmart_dump", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_generate_gigasmart_dump", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_generate_gigasmart_dump", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_generate_gigasmart_dump", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_generate_gigasmart_dump", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_generate_gigasmart_dump", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_generate_gigasmart_dump", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *GenerateGigasmartDumpAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
