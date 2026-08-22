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
var _ action.Action = (*RedefineMetadataExporterAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*RedefineMetadataExporterAction)(nil)

// RedefineMetadataExporterAction is the generated Terraform action implementation.
type RedefineMetadataExporterAction struct {
	client *client.Client
}

// RedefineMetadataExporterActionModel describes the action configuration shape.
type RedefineMetadataExporterActionModel struct {
	Alias               types.String  `tfsdk:"alias"`
	ApplicationProfiles types.List    `tfsdk:"application_profiles" json:"applicationProfiles"`
	BodyAlias           types.String  `tfsdk:"body_alias" json:"alias"`
	Cef                 types.Dynamic `tfsdk:"cef"`
	Description         types.String  `tfsdk:"description"`
	Destination         types.Dynamic `tfsdk:"destination"`
	MaxPktSize          types.Int64   `tfsdk:"max_pkt_size" json:"maxPktSize"`
	MobilitySam         types.Dynamic `tfsdk:"mobility_sam" json:"mobilitySam"`
	Monitor             types.Dynamic `tfsdk:"monitor"`
	Netflow             types.Dynamic `tfsdk:"netflow"`
	Snmp                types.Dynamic `tfsdk:"snmp"`
	Source              types.Dynamic `tfsdk:"source"`
	Type                types.String  `tfsdk:"type"`
}

// NewRedefineMetadataExporterAction returns a new instance of the generated action.
func NewRedefineMetadataExporterAction() action.Action {
	return &RedefineMetadataExporterAction{}
}

// Metadata returns the action type name.
func (r *RedefineMetadataExporterAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_redefine_metadata_exporter"
}

// Schema returns the action schema.
func (r *RedefineMetadataExporterAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Redefine a metadata exporter", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "metadata exporter alias", Required: true}, "application_profiles": schema.ListAttribute{MarkdownDescription: "application profile aliases to attach to the exporter", Optional: true, ElementType: types.StringType}, "body_alias": schema.StringAttribute{Required: true}, "cef": schema.DynamicAttribute{Optional: true}, "description": schema.StringAttribute{Optional: true}, "destination": schema.DynamicAttribute{Optional: true}, "max_pkt_size": schema.Int64Attribute{Optional: true}, "mobility_sam": schema.DynamicAttribute{Optional: true}, "monitor": schema.DynamicAttribute{Optional: true}, "netflow": schema.DynamicAttribute{Optional: true}, "snmp": schema.DynamicAttribute{Optional: true}, "source": schema.DynamicAttribute{Optional: true}, "type": schema.StringAttribute{Optional: true}}}
}

// Invoke executes the action against the remote API.
func (r *RedefineMetadataExporterAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config RedefineMetadataExporterActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *RedefineMetadataExporterAction) invokeRemote(ctx context.Context, config *RedefineMetadataExporterActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/apps/metadata/exporters/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(config.Alias.ValueString()))
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_metadata_exporter", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_metadata_exporter", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodPut, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_metadata_exporter", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_metadata_exporter", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_metadata_exporter", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_metadata_exporter", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_metadata_exporter", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_metadata_exporter", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_metadata_exporter", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_metadata_exporter", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_metadata_exporter", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_metadata_exporter", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_redefine_metadata_exporter", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *RedefineMetadataExporterAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
