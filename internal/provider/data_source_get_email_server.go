package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)
import (
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*GetEmailServerDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetEmailServerDataSource)(nil)
)

// GetEmailServerDataSource is the generated Terraform data source implementation.
type GetEmailServerDataSource struct {
	client *client.Client
}

// GetEmailServerDataSourceModel describes the data source state shape.
type GetEmailServerDataSourceModel struct {
	DomainName             types.String `tfsdk:"domain_name" json:"domainName"`
	EnableAutoSupportNotif types.Bool   `tfsdk:"enable_auto_support_notif" json:"enableAutoSupportNotif"`
	EnableSmtpAuth         types.Bool   `tfsdk:"enable_smtp_auth" json:"enableSmtpAuth"`
	IncludeHostname        types.Bool   `tfsdk:"include_hostname" json:"includeHostname"`
	MailHubPort            types.Int64  `tfsdk:"mail_hub_port" json:"mailHubPort"`
	Password               types.String `tfsdk:"password"`
	ReturnAddress          types.String `tfsdk:"return_address" json:"returnAddress"`
	SmtpServer             types.String `tfsdk:"smtp_server" json:"smtpServer"`
	Username               types.String `tfsdk:"username"`
}

// NewGetEmailServerDataSource returns a new instance of the generated data source.
func NewGetEmailServerDataSource() datasource.DataSource {
	return &GetEmailServerDataSource{}
}

// Metadata returns the data source type name.
func (d *GetEmailServerDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_email_server"
}

// Schema returns the data source schema.
func (d *GetEmailServerDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get system email server", Attributes: map[string]schema.Attribute{"domain_name": schema.StringAttribute{MarkdownDescription: "Domain Name", Computed: true}, "enable_auto_support_notif": schema.BoolAttribute{MarkdownDescription: "Enable Auto Support Notifications", Computed: true}, "enable_smtp_auth": schema.BoolAttribute{MarkdownDescription: "Enable SMTP auth", Computed: true}, "include_hostname": schema.BoolAttribute{MarkdownDescription: "Include Hostname", Computed: true}, "mail_hub_port": schema.Int64Attribute{Computed: true}, "password": schema.StringAttribute{MarkdownDescription: "SMTP Password", Computed: true, Sensitive: true}, "return_address": schema.StringAttribute{MarkdownDescription: "Return Address", Computed: true}, "smtp_server": schema.StringAttribute{MarkdownDescription: "SMTP Server", Computed: true}, "username": schema.StringAttribute{MarkdownDescription: "SMTP Username", Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetEmailServerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetEmailServerDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	d.readRemote(ctx, &config, resp)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}

// readRemote performs the read HTTP exchange and decodes the response into config. Extracted from Read so the request/response logic is unit-testable without a tfsdk.Config.
func (d *GetEmailServerDataSource) readRemote(ctx context.Context, config *GetEmailServerDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/system/email/server"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_email_server", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_email_server", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_email_server", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_email_server", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_email_server", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_email_server", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_email_server", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_email_server", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_email_server", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_email_server", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_email_server", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_email_server", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_email_server", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_email_server", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetEmailServerDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
