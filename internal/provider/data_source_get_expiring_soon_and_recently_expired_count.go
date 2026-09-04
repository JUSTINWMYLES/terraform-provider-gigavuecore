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
	_ datasource.DataSource              = (*GetExpiringSoonAndRecentlyExpiredCountDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetExpiringSoonAndRecentlyExpiredCountDataSource)(nil)
)

// GetExpiringSoonAndRecentlyExpiredCountDataSource is the generated Terraform data source implementation.
type GetExpiringSoonAndRecentlyExpiredCountDataSource struct {
	client *client.Client
}

// GetExpiringSoonAndRecentlyExpiredCountDataSourceModel describes the data source state shape.
type GetExpiringSoonAndRecentlyExpiredCountDataSourceModel struct {
	ExpiringGraceActivations types.String `tfsdk:"expiring_grace_activations" json:"expiringGraceActivations"`
	ExpiringGraceLicenses    types.String `tfsdk:"expiring_grace_licenses" json:"expiringGraceLicenses"`
	ExpiringSoonActivations  types.String `tfsdk:"expiring_soon_activations" json:"expiringSoonActivations"`
	ExpiringSoonLicenses     types.String `tfsdk:"expiring_soon_licenses" json:"expiringSoonLicenses"`
}

// NewGetExpiringSoonAndRecentlyExpiredCountDataSource returns a new instance of the generated data source.
func NewGetExpiringSoonAndRecentlyExpiredCountDataSource() datasource.DataSource {
	return &GetExpiringSoonAndRecentlyExpiredCountDataSource{}
}

// Metadata returns the data source type name.
func (d *GetExpiringSoonAndRecentlyExpiredCountDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_expiring_soon_and_recently_expired_count"
}

// Schema returns the data source schema.
func (d *GetExpiringSoonAndRecentlyExpiredCountDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Returns number of floating licenses that are expiring soon or are recently expired", Attributes: map[string]schema.Attribute{"expiring_grace_activations": schema.StringAttribute{Computed: true}, "expiring_grace_licenses": schema.StringAttribute{Computed: true}, "expiring_soon_activations": schema.StringAttribute{Computed: true}, "expiring_soon_licenses": schema.StringAttribute{Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetExpiringSoonAndRecentlyExpiredCountDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetExpiringSoonAndRecentlyExpiredCountDataSourceModel
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
func (d *GetExpiringSoonAndRecentlyExpiredCountDataSource) readRemote(ctx context.Context, config *GetExpiringSoonAndRecentlyExpiredCountDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/licensing/activations/expiringCount"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_expiring_soon_and_recently_expired_count", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_expiring_soon_and_recently_expired_count", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_expiring_soon_and_recently_expired_count", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_expiring_soon_and_recently_expired_count", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_expiring_soon_and_recently_expired_count", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_expiring_soon_and_recently_expired_count", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_expiring_soon_and_recently_expired_count", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_expiring_soon_and_recently_expired_count", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_expiring_soon_and_recently_expired_count", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_expiring_soon_and_recently_expired_count", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_expiring_soon_and_recently_expired_count", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_expiring_soon_and_recently_expired_count", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_expiring_soon_and_recently_expired_count", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_expiring_soon_and_recently_expired_count", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetExpiringSoonAndRecentlyExpiredCountDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
