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
var _ action.Action = (*RedefineGsGroupHealthCheckAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*RedefineGsGroupHealthCheckAction)(nil)

// RedefineGsGroupHealthCheckAction is the generated Terraform action implementation.
type RedefineGsGroupHealthCheckAction struct {
	client *client.Client
}

// RedefineGsGroupHealthCheckActionModel describes the action configuration shape.
type RedefineGsGroupHealthCheckActionModel struct {
	Action        types.String `tfsdk:"action"`
	Alias         types.String `tfsdk:"alias"`
	DstPort       types.Int64  `tfsdk:"dst_port" json:"dstPort"`
	Enabled       types.Bool   `tfsdk:"enabled"`
	Interval      types.Int64  `tfsdk:"interval"`
	Protocol      types.String `tfsdk:"protocol"`
	RcvPort       types.Int64  `tfsdk:"rcv_port" json:"rcvPort"`
	Retries       types.Int64  `tfsdk:"retries"`
	RoundTripTime types.Int64  `tfsdk:"round_trip_time" json:"roundTripTime"`
	SrcPort       types.Int64  `tfsdk:"src_port" json:"srcPort"`
}

// NewRedefineGsGroupHealthCheckAction returns a new instance of the generated action.
func NewRedefineGsGroupHealthCheckAction() action.Action {
	return &RedefineGsGroupHealthCheckAction{}
}

// Metadata returns the action type name.
func (r *RedefineGsGroupHealthCheckAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_redefine_gs_group_health_check"
}

// Schema returns the action schema.
func (r *RedefineGsGroupHealthCheckAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Redefine GS Group's health check", Attributes: map[string]schema.Attribute{"action": schema.StringAttribute{Optional: true}, "alias": schema.StringAttribute{MarkdownDescription: "alias of the target GS Group", Required: true}, "dst_port": schema.Int64Attribute{Optional: true}, "enabled": schema.BoolAttribute{Optional: true}, "interval": schema.Int64Attribute{Optional: true}, "protocol": schema.StringAttribute{Optional: true}, "rcv_port": schema.Int64Attribute{Optional: true}, "retries": schema.Int64Attribute{Optional: true}, "round_trip_time": schema.Int64Attribute{Optional: true}, "src_port": schema.Int64Attribute{Optional: true}}}
}

// Invoke executes the action against the remote API.
func (r *RedefineGsGroupHealthCheckAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config RedefineGsGroupHealthCheckActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *RedefineGsGroupHealthCheckAction) invokeRemote(ctx context.Context, config *RedefineGsGroupHealthCheckActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/gsGroups/{alias}/params/healthCheck"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(config.Alias.ValueString()))
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_health_check", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_health_check", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_health_check", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_health_check", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_health_check", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_health_check", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_health_check", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_health_check", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_health_check", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_health_check", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_health_check", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_health_check", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_gs_group_health_check", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *RedefineGsGroupHealthCheckAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
