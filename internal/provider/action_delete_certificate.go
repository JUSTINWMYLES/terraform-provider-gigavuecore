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
var _ action.Action = (*DeleteCertificateAction)(nil)

// Compile-time interface assertion for action.ActionWithConfigure.
var _ action.ActionWithConfigure = (*DeleteCertificateAction)(nil)

// DeleteCertificateAction is the generated Terraform action implementation.
type DeleteCertificateAction struct {
	client *client.Client
}

// DeleteCertificateActionModel describes the action configuration shape.
type DeleteCertificateActionModel struct {
	Config           types.Object `tfsdk:"config"`
	ConfigLevel      types.String `tfsdk:"config_level" json:"configLevel"`
	ConfigLevelValue types.List   `tfsdk:"config_level_value" json:"configLevelValue"`
	ConfigType       types.String `tfsdk:"config_type" json:"configType"`
	Modifiable       types.Bool   `tfsdk:"modifiable"`
	RefCount         types.Int64  `tfsdk:"ref_count" json:"refCount"`
	TemplateName     types.String `tfsdk:"template_name" json:"templateName"`
	UpdateTime       types.String `tfsdk:"update_time" json:"updateTime"`
}

// NewDeleteCertificateAction returns a new instance of the generated action.
func NewDeleteCertificateAction() action.Action {
	return &DeleteCertificateAction{}
}

// Metadata returns the action type name.
func (r *DeleteCertificateAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_delete_certificate"
}

// Schema returns the action schema.
func (r *DeleteCertificateAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Push global delete configuration for the selected certificate to all the devices", Attributes: map[string]schema.Attribute{"config": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"device_ssl_certificate_configs": schema.ListNestedAttribute{MarkdownDescription: "Available when ConfigType is SSL_CERTIFICATE_TEMPLATE", Optional: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"issuer": schema.StringAttribute{MarkdownDescription: "issuer details of the certificate", Optional: true}, "not_after": schema.StringAttribute{MarkdownDescription: "date and time when certificate stops being valid ([rfc3339](https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14))", Optional: true}, "not_before": schema.StringAttribute{MarkdownDescription: "date and time when certificate starts being valid ([rfc3339](https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14))", Optional: true}, "operation_type": schema.StringAttribute{Required: true}, "signature_algorithm": schema.StringAttribute{Optional: true}, "subject": schema.StringAttribute{MarkdownDescription: "subject name of the certificate", Optional: true}, "trusted_ca": schema.SingleNestedAttribute{Required: true, Attributes: map[string]schema.Attribute{"name": schema.StringAttribute{MarkdownDescription: "name of the certificate", Required: true}}}, "upload_spec": schema.SingleNestedAttribute{Required: true, Attributes: map[string]schema.Attribute{"info": schema.SingleNestedAttribute{Optional: true, Attributes: map[string]schema.Attribute{"comment": schema.StringAttribute{MarkdownDescription: "a short description of the certificate", Optional: true}, "name": schema.StringAttribute{MarkdownDescription: "name of the certificate", Required: true}, "passphrase": schema.StringAttribute{MarkdownDescription: "used to decrypt pkcs12 and private keys", Optional: true}, "type": schema.StringAttribute{MarkdownDescription: "type of the certificate", Required: true}}}, "pem": schema.StringAttribute{MarkdownDescription: "contents of the certificate in pem format", Optional: true}}}}}}}}, "config_level": schema.StringAttribute{MarkdownDescription: "Scope of the applied FM template", Optional: true}, "config_level_value": schema.ListAttribute{Optional: true, ElementType: types.StringType}, "config_type": schema.StringAttribute{MarkdownDescription: "Configuration Type of the FM template", Optional: true}, "modifiable": schema.BoolAttribute{Optional: true}, "ref_count": schema.Int64Attribute{Optional: true}, "template_name": schema.StringAttribute{Optional: true}, "update_time": schema.StringAttribute{Optional: true}}}
}

// Invoke executes the action against the remote API.
func (r *DeleteCertificateAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config DeleteCertificateActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.invokeRemote(ctx, &config, resp)
}

// invokeRemote performs the invoke HTTP exchange and surfaces any error via diagnostics. Extracted from Invoke so the request/response logic is unit-testable without a tfsdk.Config.
func (r *DeleteCertificateAction) invokeRemote(ctx context.Context, config *DeleteCertificateActionModel, resp *action.InvokeResponse) {
	if r.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/fm/templates/ssl"
	body, err := modelToJSONMap(&config)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_delete_certificate", fmt.Sprintf("Could not build request body: %s", err))
		return
	}
	payload, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_delete_certificate", fmt.Sprintf("Could not encode request body: %s", err))
		return
	}
	httpReq, err := r.client.NewRequest(ctx, http.MethodDelete, reqPath, bytes.NewReader(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_delete_certificate", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := r.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking gigavuecore_delete_certificate", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if !(httpResp.StatusCode == 201) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_certificate", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_certificate", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_certificate", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_certificate", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_certificate", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_certificate", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_certificate", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error invoking gigavuecore_delete_certificate", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error invoking gigavuecore_delete_certificate", apiErr.Error())
			return
		}
	}
}

// Configure stores the API client supplied by the provider.
func (r *DeleteCertificateAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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
