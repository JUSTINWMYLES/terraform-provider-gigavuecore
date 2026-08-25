package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)
import (
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*GetAllUnlicensedChassisAndCardsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetAllUnlicensedChassisAndCardsDataSource)(nil)
)

// GetAllUnlicensedChassisAndCardsDataSource is the generated Terraform data source implementation.
type GetAllUnlicensedChassisAndCardsDataSource struct {
	client *client.Client
}

// GetAllUnlicensedChassisAndCardsDataSourceModel describes the data source state shape.
type GetAllUnlicensedChassisAndCardsDataSourceModel struct {
	UnlicensedCards   types.List `tfsdk:"unlicensed_cards" json:"unlicensedCards"`
	UnlicensedChassis types.List `tfsdk:"unlicensed_chassis" json:"unlicensedChassis"`
	UseDb             types.Bool `tfsdk:"use_db" json:"useDb"`
}

// NewGetAllUnlicensedChassisAndCardsDataSource returns a new instance of the generated data source.
func NewGetAllUnlicensedChassisAndCardsDataSource() datasource.DataSource {
	return &GetAllUnlicensedChassisAndCardsDataSource{}
}

// Metadata returns the data source type name.
func (d *GetAllUnlicensedChassisAndCardsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_all_unlicensed_chassis_and_cards"
}

// Schema returns the data source schema.
func (d *GetAllUnlicensedChassisAndCardsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get all the chassis and cards that need GVOS and GVOS-module license respectively", Attributes: map[string]schema.Attribute{"unlicensed_cards": schema.ListAttribute{MarkdownDescription: "list of cards that are missing GVOS-module license", Computed: true, ElementType: types.StringType}, "unlicensed_chassis": schema.ListAttribute{MarkdownDescription: "list of chassis that are missing GVOS license", Computed: true, ElementType: types.StringType}, "use_db": schema.BoolAttribute{MarkdownDescription: "if true, use FM database as information provider, else retrieve from node", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetAllUnlicensedChassisAndCardsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetAllUnlicensedChassisAndCardsDataSourceModel
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
func (d *GetAllUnlicensedChassisAndCardsDataSource) readRemote(ctx context.Context, config *GetAllUnlicensedChassisAndCardsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/licensing/module/missing"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_unlicensed_chassis_and_cards", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	if !config.UseDb.IsNull() {
		query.Set("useDb", strconv.FormatBool(config.UseDb.ValueBool()))
	}
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_unlicensed_chassis_and_cards", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_unlicensed_chassis_and_cards", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_unlicensed_chassis_and_cards", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_unlicensed_chassis_and_cards", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_unlicensed_chassis_and_cards", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_unlicensed_chassis_and_cards", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_unlicensed_chassis_and_cards", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_unlicensed_chassis_and_cards", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_unlicensed_chassis_and_cards", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_all_unlicensed_chassis_and_cards", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_unlicensed_chassis_and_cards", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_unlicensed_chassis_and_cards", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_unlicensed_chassis_and_cards", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetAllUnlicensedChassisAndCardsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
