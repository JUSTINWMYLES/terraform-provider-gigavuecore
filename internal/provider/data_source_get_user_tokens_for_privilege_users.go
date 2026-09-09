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
	_ datasource.DataSource              = (*GetUserTokensForPrivilegeUsersDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetUserTokensForPrivilegeUsersDataSource)(nil)
)

// GetUserTokensForPrivilegeUsersDataSource is the generated Terraform data source implementation.
type GetUserTokensForPrivilegeUsersDataSource struct {
	client *client.Client
}

// GetUserTokensForPrivilegeUsersDataSourceModel describes the data source state shape.
type GetUserTokensForPrivilegeUsersDataSourceModel struct {
	Items types.List `tfsdk:"items"`
}

// NewGetUserTokensForPrivilegeUsersDataSource returns a new instance of the generated data source.
func NewGetUserTokensForPrivilegeUsersDataSource() datasource.DataSource {
	return &GetUserTokensForPrivilegeUsersDataSource{}
}

// Metadata returns the data source type name.
func (d *GetUserTokensForPrivilegeUsersDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_user_tokens_for_privilege_users"
}

// Schema returns the data source schema.
func (d *GetUserTokensForPrivilegeUsersDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Returns API tokens of the other users. Users with FM Security Management role with write access can access API.", Attributes: map[string]schema.Attribute{"items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"authentication_type": schema.StringAttribute{MarkdownDescription: "Authentication Type local/radius/tacacs+/external", Computed: true}, "created_by": schema.StringAttribute{MarkdownDescription: "FM user who has created the token", Computed: true}, "created_ts": schema.Int64Attribute{MarkdownDescription: "Token creation timestamp", Computed: true}, "expiry_time": schema.StringAttribute{MarkdownDescription: "Expiry Time in number of days. Default would be 30 and Maximum of 105 days.", Computed: true}, "expiry_ts": schema.StringAttribute{MarkdownDescription: "Expiry Timestamp", Computed: true}, "groups": schema.ListAttribute{MarkdownDescription: "FM User Groups", Computed: true, ElementType: types.StringType}, "token": schema.StringAttribute{MarkdownDescription: "FM generated JWT token for FM REST API access", Computed: true, Sensitive: true}, "token_id": schema.StringAttribute{MarkdownDescription: "Random alpha-numeric 64 digit string", Computed: true}, "token_name": schema.StringAttribute{MarkdownDescription: "User defined token name", Computed: true}, "usage_count": schema.Int64Attribute{MarkdownDescription: "Token Tracking count. Usagecount keep track of number times token was used to access API's ", Computed: true}, "username": schema.StringAttribute{MarkdownDescription: "FM Username", Computed: true}}}}}}
}

// Read fetches remote state into the data source model.
func (d *GetUserTokensForPrivilegeUsersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetUserTokensForPrivilegeUsersDataSourceModel
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
func (d *GetUserTokensForPrivilegeUsersDataSource) readListRemote(ctx context.Context, config *GetUserTokensForPrivilegeUsersDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/tokens/manage"
	params := url.Values{}
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_user_tokens_for_privilege_users", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_user_tokens_for_privilege_users", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["fmUserTokenEntities"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_user_tokens_for_privilege_users", fmt.Sprintf("Could not decode list page: missing %q array", "fmUserTokenEntities"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_user_tokens_for_privilege_users", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetUserTokensForPrivilegeUsersDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
