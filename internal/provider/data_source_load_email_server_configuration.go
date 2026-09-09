package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)
import (
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*LoadEmailServerConfigurationDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadEmailServerConfigurationDataSource)(nil)
)

// LoadEmailServerConfigurationDataSource is the generated Terraform data source implementation.
type LoadEmailServerConfigurationDataSource struct {
	client *client.Client
}

// LoadEmailServerConfigurationDataSourceModel describes the data source state shape.
type LoadEmailServerConfigurationDataSourceModel struct {
	EmailHost      types.String `tfsdk:"email_host" json:"emailHost"`
	EnableSmtpAuth types.Bool   `tfsdk:"enable_smtp_auth" json:"enableSmtpAuth"`
	From           types.String `tfsdk:"from"`
	Password       types.String `tfsdk:"password"`
	Port           types.Int64  `tfsdk:"port"`
	UserName       types.String `tfsdk:"user_name" json:"userName"`
}

// NewLoadEmailServerConfigurationDataSource returns a new instance of the generated data source.
func NewLoadEmailServerConfigurationDataSource() datasource.DataSource {
	return &LoadEmailServerConfigurationDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadEmailServerConfigurationDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_email_server_configuration"
}

// Schema returns the data source schema.
func (d *LoadEmailServerConfigurationDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load Email Server Configuration", Attributes: map[string]schema.Attribute{"email_host": schema.StringAttribute{MarkdownDescription: "Address of the SMTP Server", Computed: true}, "enable_smtp_auth": schema.BoolAttribute{MarkdownDescription: "Enable/Disable SMTP Authentication", Computed: true}, "from": schema.StringAttribute{MarkdownDescription: "From address for the email", Computed: true}, "password": schema.StringAttribute{MarkdownDescription: "Password for the SMTP Server", Computed: true, Sensitive: true}, "port": schema.Int64Attribute{MarkdownDescription: "SMTP Server Port", Computed: true}, "user_name": schema.StringAttribute{MarkdownDescription: "Username for the SMTP Server", Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadEmailServerConfigurationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadEmailServerConfigurationDataSourceModel
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
func (d *LoadEmailServerConfigurationDataSource) readRemote(ctx context.Context, config *LoadEmailServerConfigurationDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/notification/email/emailServer"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_email_server_configuration", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_email_server_configuration", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_email_server_configuration", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_email_server_configuration", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_email_server_configuration", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_email_server_configuration", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_email_server_configuration", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_email_server_configuration", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_email_server_configuration", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_email_server_configuration", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_load_email_server_configuration", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_load_email_server_configuration", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_email_server_configuration", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_email_server_configuration", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadEmailServerConfigurationDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
