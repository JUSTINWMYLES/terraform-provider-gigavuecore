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
	_ datasource.DataSource              = (*LoadActiveUserDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadActiveUserDataSource)(nil)
)

// LoadActiveUserDataSource is the generated Terraform data source implementation.
type LoadActiveUserDataSource struct {
	client *client.Client
}

// LoadActiveUserDataSourceModel describes the data source state shape.
type LoadActiveUserDataSourceModel struct {
	AuthMethod        types.String `tfsdk:"auth_method" json:"authMethod"`
	Enabled           types.Bool   `tfsdk:"enabled"`
	FmAuthorizedRoles types.List   `tfsdk:"fm_authorized_roles" json:"fmAuthorizedRoles"`
	FullName          types.String `tfsdk:"full_name" json:"fullName"`
	Groups            types.List   `tfsdk:"groups"`
	Password          types.String `tfsdk:"password"`
	Username          types.String `tfsdk:"username"`
}

// NewLoadActiveUserDataSource returns a new instance of the generated data source.
func NewLoadActiveUserDataSource() datasource.DataSource {
	return &LoadActiveUserDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadActiveUserDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_active_user"
}

// Schema returns the data source schema.
func (d *LoadActiveUserDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load Active User Details", Attributes: map[string]schema.Attribute{"auth_method": schema.StringAttribute{MarkdownDescription: "authentication method", Computed: true}, "enabled": schema.BoolAttribute{Computed: true}, "fm_authorized_roles": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "full_name": schema.StringAttribute{MarkdownDescription: "user's full name", Computed: true}, "groups": schema.ListAttribute{Computed: true, ElementType: types.StringType}, "password": schema.StringAttribute{MarkdownDescription: "password", Computed: true, Sensitive: true}, "username": schema.StringAttribute{MarkdownDescription: "username", Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadActiveUserDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadActiveUserDataSourceModel
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
func (d *LoadActiveUserDataSource) readRemote(ctx context.Context, config *LoadActiveUserDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/user/active"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_active_user", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_active_user", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_active_user", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_active_user", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_active_user", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_active_user", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_active_user", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_active_user", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_active_user", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_active_user", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_load_active_user", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_load_active_user", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_active_user", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["activeuser"]; ok {
		if m, ok := v.(map[string]any); ok {
			data = m
		} else if arr, ok := v.([]any); ok && len(arr) > 0 {
			if m, ok := arr[0].(map[string]any); ok {
				data = m
			}
		}
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_active_user", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadActiveUserDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
