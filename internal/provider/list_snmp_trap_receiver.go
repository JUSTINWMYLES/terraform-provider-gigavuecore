package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)
import (
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	diag "github.com/hashicorp/terraform-plugin-framework/diag"
	list "github.com/hashicorp/terraform-plugin-framework/list"
	listschema "github.com/hashicorp/terraform-plugin-framework/list/schema"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	types "github.com/hashicorp/terraform-plugin-framework/types"
	tftypes "github.com/hashicorp/terraform-plugin-go/tftypes"
)

// Compile-time interface assertion.
var _ list.ListResource = (*SnmpTrapReceiverListResource)(nil)
var _ list.ListResourceWithConfigure = (*SnmpTrapReceiverListResource)(nil)

// SnmpTrapReceiverListResource is the generated Terraform list resource implementation.
type SnmpTrapReceiverListResource struct {
	client *client.Client
}

// SnmpTrapReceiverListResourceModel describes the gigavuecore_snmp_trap_receiver list filter configuration shape.
type SnmpTrapReceiverListResourceModel struct {
	Alias         types.String `tfsdk:"alias"`
	AuthProtocol  types.String `tfsdk:"auth_protocol"`
	Community     types.String `tfsdk:"community"`
	IpAddress     types.String `tfsdk:"ip_address"`
	Page          types.String `tfsdk:"page"`
	PrivProtocol  types.String `tfsdk:"priv_protocol"`
	SecurityLevel types.String `tfsdk:"security_level"`
	SnmpPort      types.String `tfsdk:"snmp_port"`
	SnmpRetries   types.String `tfsdk:"snmp_retries"`
	SnmpTimeout   types.String `tfsdk:"snmp_timeout"`
	SnmpVersion   types.String `tfsdk:"snmp_version"`
	Sort          types.String `tfsdk:"sort"`
	UserName      types.String `tfsdk:"user_name"`
}

// NewSnmpTrapReceiverListResource returns a new instance of the generated list resource.
func NewSnmpTrapReceiverListResource() list.ListResource {
	return &SnmpTrapReceiverListResource{}
}

// Metadata returns the list resource type name.
func (l *SnmpTrapReceiverListResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_snmp_trap_receiver"
}

// ListResourceConfigSchema returns the list resource config schema.
func (l *SnmpTrapReceiverListResource) ListResourceConfigSchema(_ context.Context, _ list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = listschema.Schema{MarkdownDescription: "Get All External Trap Receiver", Attributes: map[string]listschema.Attribute{"alias": listschema.StringAttribute{MarkdownDescription: "alias", Optional: true}, "auth_protocol": listschema.StringAttribute{MarkdownDescription: "authProtocol", Optional: true}, "community": listschema.StringAttribute{MarkdownDescription: "community", Optional: true}, "ip_address": listschema.StringAttribute{MarkdownDescription: "ipAddress", Optional: true}, "page": listschema.StringAttribute{MarkdownDescription: "page", Optional: true}, "priv_protocol": listschema.StringAttribute{MarkdownDescription: "privProtocol", Optional: true}, "security_level": listschema.StringAttribute{MarkdownDescription: "securityLevel", Optional: true}, "snmp_port": listschema.StringAttribute{MarkdownDescription: "snmpPort", Optional: true}, "snmp_retries": listschema.StringAttribute{MarkdownDescription: "snmpRetries", Optional: true}, "snmp_timeout": listschema.StringAttribute{MarkdownDescription: "snmpTimeout", Optional: true}, "snmp_version": listschema.StringAttribute{MarkdownDescription: "snmpVersion", Optional: true}, "sort": listschema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}, "user_name": listschema.StringAttribute{MarkdownDescription: "userName", Optional: true}}}
}

// List streams matching resource instances for terraform query.
func (l *SnmpTrapReceiverListResource) List(ctx context.Context, req list.ListRequest, stream *list.ListResultsStream) {
	stream.Results = func(push func(list.ListResult) bool) {
		var config SnmpTrapReceiverListResourceModel
		diags := req.Config.Get(ctx, &config)
		if diags.HasError() {
			result := req.NewListResult(ctx)
			result.Diagnostics = diags
			push(result)
			return
		}
		items, diags := l.listRemote(ctx, &config)
		if diags.HasError() {
			result := req.NewListResult(ctx)
			result.Diagnostics = diags
			push(result)
			return
		}
		for _, item := range items {
			result := req.NewListResult(ctx)
			itemMap := map[string]json.RawMessage{}
			if err := json.Unmarshal(item, &itemMap); err != nil {
				result.Diagnostics.AddError("Error listing gigavuecore_snmp_trap_receiver", fmt.Sprintf("Could not decode list item: %s", err))
				if !push(result) {
					return
				}
				continue
			}
			identity := map[string]json.RawMessage{}
			aliasValue, ok := itemMap["alias"]
			if !ok {
				if itemMap["metadata"] != nil {
					metaMap := map[string]json.RawMessage{}
					if json.Unmarshal(itemMap["metadata"], &metaMap) == nil {
						aliasValue, ok = metaMap["alias"]
					}
				}
			}
			if !ok {
				aliasValue, ok = itemMap["id"]
			}
			if !ok {
				result.Diagnostics.AddError("Error listing gigavuecore_snmp_trap_receiver", "List item is missing identity attribute \"alias\".")
				if !push(result) {
					return
				}
				continue
			}
			identity["alias"] = aliasValue
			idJSON, err := json.Marshal(identity)
			if err != nil {
				result.Diagnostics.AddError("Error listing gigavuecore_snmp_trap_receiver", fmt.Sprintf("Could not encode list item identity: %s", err))
				if !push(result) {
					return
				}
				continue
			}
			idVal, err := tftypes.ValueFromJSON(idJSON, req.ResourceIdentitySchema.Type().TerraformType(ctx))
			if err != nil {
				result.Diagnostics.AddError("Error listing gigavuecore_snmp_trap_receiver", fmt.Sprintf("Could not decode list item identity: %s", err))
				if !push(result) {
					return
				}
				continue
			}
			result.Identity.Raw = idVal
			if req.IncludeResource {
				resVal, err := tftypes.ValueFromJSON(item, req.ResourceSchema.Type().TerraformType(ctx))
				if err != nil {
					result.Diagnostics.AddWarning("Error listing gigavuecore_snmp_trap_receiver", fmt.Sprintf("Could not decode list item into the resource schema: %s", err))
				} else {
					result.Resource.Raw = resVal
				}
			}
			if !push(result) {
				return
			}
		}
	}
}

// listRemote fetches and decodes the collection pages, returning the items and any diagnostics for the List iterator to surface.
func (l *SnmpTrapReceiverListResource) listRemote(ctx context.Context, config *SnmpTrapReceiverListResourceModel) ([]json.RawMessage, diag.Diagnostics) {
	var diags diag.Diagnostics
	if l.client == nil {
		diags.AddError("Client Not Configured", "The API client was not set on the list resource. The provider Configure method must run before list operations; this is a bug in the generated provider.")
		return nil, diags
	}
	reqPath := "/notification/snmpTrap/snmpTrapReceiver"
	params := url.Values{}
	if !config.Alias.IsNull() {
		params.Set("alias", config.Alias.ValueString())
	}
	if !config.IpAddress.IsNull() {
		params.Set("ipAddress", config.IpAddress.ValueString())
	}
	if !config.SnmpRetries.IsNull() {
		params.Set("snmpRetries", config.SnmpRetries.ValueString())
	}
	if !config.SnmpTimeout.IsNull() {
		params.Set("snmpTimeout", config.SnmpTimeout.ValueString())
	}
	if !config.SnmpPort.IsNull() {
		params.Set("snmpPort", config.SnmpPort.ValueString())
	}
	if !config.SnmpVersion.IsNull() {
		params.Set("snmpVersion", config.SnmpVersion.ValueString())
	}
	if !config.SecurityLevel.IsNull() {
		params.Set("securityLevel", config.SecurityLevel.ValueString())
	}
	if !config.UserName.IsNull() {
		params.Set("userName", config.UserName.ValueString())
	}
	if !config.AuthProtocol.IsNull() {
		params.Set("authProtocol", config.AuthProtocol.ValueString())
	}
	if !config.PrivProtocol.IsNull() {
		params.Set("privProtocol", config.PrivProtocol.ValueString())
	}
	if !config.Community.IsNull() {
		params.Set("community", config.Community.ValueString())
	}
	if !config.Page.IsNull() {
		params.Set("page", config.Page.ValueString())
	}
	if !config.Sort.IsNull() {
		params.Set("sort", config.Sort.ValueString())
	}
	var nextURL string
	fetch := func(ctx context.Context, p url.Values) (*http.Response, error) {
		httpReq, err := l.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
		if err != nil {
			return nil, err
		}
		if nextURL != "" {
			parsed, perr := url.Parse(nextURL)
			if perr != nil {
				return nil, perr
			}
			httpReq.URL = parsed
		} else {
			httpReq.URL.RawQuery = p.Encode()
		}
		return l.client.Do(httpReq)
	}
	pages, err := client.ListAllPages(ctx, params, fetch, nil)
	if err != nil {
		diags.AddError("Error listing gigavuecore_snmp_trap_receiver", fmt.Sprintf("Could not read list response: %s", err))
		return nil, diags
	}
	allItems := []json.RawMessage{}
	for _, page := range pages {
		pageObj := map[string]json.RawMessage{}
		if err := json.Unmarshal(page, &pageObj); err != nil {
			diags.AddError("Error listing gigavuecore_snmp_trap_receiver", fmt.Sprintf("Could not decode list page: %s", err))
			return nil, diags
		}
		rawItems, ok := pageObj["externalTrapReceivers"]
		if !ok {
			diags.AddError("Error listing gigavuecore_snmp_trap_receiver", fmt.Sprintf("Could not decode list page: missing %q array", "externalTrapReceivers"))
			return nil, diags
		}
		items := []json.RawMessage{}
		if err := json.Unmarshal(rawItems, &items); err != nil {
			diags.AddError("Error listing gigavuecore_snmp_trap_receiver", fmt.Sprintf("Could not decode list page: %s", err))
			return nil, diags
		}
		allItems = append(allItems, items...)
	}
	return allItems, diags
}

// Configure stores the API client supplied by the provider.
func (l *SnmpTrapReceiverListResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected List Resource Configure Type", fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData))
		return
	}
	l.client = c
}
