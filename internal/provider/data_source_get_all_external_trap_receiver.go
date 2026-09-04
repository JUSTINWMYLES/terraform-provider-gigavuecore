package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)
import (
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*GetAllExternalTrapReceiverDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetAllExternalTrapReceiverDataSource)(nil)
)

// GetAllExternalTrapReceiverDataSource is the generated Terraform data source implementation.
type GetAllExternalTrapReceiverDataSource struct {
	client *client.Client
}

// GetAllExternalTrapReceiverDataSourceModel describes the data source state shape.
type GetAllExternalTrapReceiverDataSourceModel struct {
	Alias         types.String `tfsdk:"alias"`
	AuthProtocol  types.String `tfsdk:"auth_protocol" json:"authProtocol"`
	Community     types.String `tfsdk:"community"`
	IpAddress     types.String `tfsdk:"ip_address" json:"ipAddress"`
	Items         types.List   `tfsdk:"items"`
	Page          types.String `tfsdk:"page"`
	PrivProtocol  types.String `tfsdk:"priv_protocol" json:"privProtocol"`
	SecurityLevel types.String `tfsdk:"security_level" json:"securityLevel"`
	SnmpPort      types.String `tfsdk:"snmp_port" json:"snmpPort"`
	SnmpRetries   types.String `tfsdk:"snmp_retries" json:"snmpRetries"`
	SnmpTimeout   types.String `tfsdk:"snmp_timeout" json:"snmpTimeout"`
	SnmpVersion   types.String `tfsdk:"snmp_version" json:"snmpVersion"`
	Sort          types.String `tfsdk:"sort"`
	UserName      types.String `tfsdk:"user_name" json:"userName"`
}

// NewGetAllExternalTrapReceiverDataSource returns a new instance of the generated data source.
func NewGetAllExternalTrapReceiverDataSource() datasource.DataSource {
	return &GetAllExternalTrapReceiverDataSource{}
}

// Metadata returns the data source type name.
func (d *GetAllExternalTrapReceiverDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_all_external_trap_receiver"
}

// Schema returns the data source schema.
func (d *GetAllExternalTrapReceiverDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get All External Trap Receiver", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "alias", Optional: true}, "auth_protocol": schema.StringAttribute{MarkdownDescription: "authProtocol", Optional: true}, "community": schema.StringAttribute{MarkdownDescription: "community", Optional: true}, "ip_address": schema.StringAttribute{MarkdownDescription: "ipAddress", Optional: true}, "items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Alias of the External Trap Receiver", Computed: true}, "auth_password": schema.StringAttribute{MarkdownDescription: "Authentication Password", Computed: true, Sensitive: true}, "auth_protocol": schema.StringAttribute{MarkdownDescription: "Authentication Protocol", Computed: true}, "community": schema.StringAttribute{MarkdownDescription: "Community name for the transaction with the remote system", Computed: true}, "ip_address": schema.StringAttribute{MarkdownDescription: "IP Address of the External Trap Receiver", Computed: true}, "priv_password": schema.StringAttribute{MarkdownDescription: "Private Password", Computed: true, Sensitive: true}, "priv_protocol": schema.StringAttribute{MarkdownDescription: "Private Protocol", Computed: true}, "security_level": schema.StringAttribute{MarkdownDescription: "Security Level", Computed: true}, "snmp_port": schema.Int64Attribute{MarkdownDescription: "Destination port number", Computed: true}, "snmp_retries": schema.Int64Attribute{MarkdownDescription: "Number of retries", Computed: true}, "snmp_timeout": schema.Int64Attribute{MarkdownDescription: "Timeout (in milliseconds)", Computed: true}, "snmp_version": schema.StringAttribute{MarkdownDescription: "SNMP Version", Computed: true}, "user_name": schema.StringAttribute{MarkdownDescription: "Username", Computed: true}}}}, "page": schema.StringAttribute{MarkdownDescription: "page", Optional: true}, "priv_protocol": schema.StringAttribute{MarkdownDescription: "privProtocol", Optional: true}, "security_level": schema.StringAttribute{MarkdownDescription: "securityLevel", Optional: true}, "snmp_port": schema.StringAttribute{MarkdownDescription: "snmpPort", Optional: true}, "snmp_retries": schema.StringAttribute{MarkdownDescription: "snmpRetries", Optional: true}, "snmp_timeout": schema.StringAttribute{MarkdownDescription: "snmpTimeout", Optional: true}, "snmp_version": schema.StringAttribute{MarkdownDescription: "snmpVersion", Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}, "user_name": schema.StringAttribute{MarkdownDescription: "userName", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetAllExternalTrapReceiverDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetAllExternalTrapReceiverDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	d.readListRemote(ctx, &config, resp)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}

// readListRemote performs the paginated read HTTP exchange and decodes the response array into config. Extracted from Read so the request/response logic is unit-testable without a tfsdk.Config.
func (d *GetAllExternalTrapReceiverDataSource) readListRemote(ctx context.Context, config *GetAllExternalTrapReceiverDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
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
		httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
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
		return d.client.Do(httpReq)
	}
	pages, err := client.ListAllPages(ctx, params, fetch, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_external_trap_receiver", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_external_trap_receiver", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["externalTrapReceivers"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_external_trap_receiver", fmt.Sprintf("Could not decode list page: missing %q array", "externalTrapReceivers"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_external_trap_receiver", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetAllExternalTrapReceiverDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Data Source Configure Type", fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData))
		return
	}
	d.client = c
}
