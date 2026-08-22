package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)
import (
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*GetForeignMastersDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetForeignMastersDataSource)(nil)
)

// GetForeignMastersDataSource is the generated Terraform data source implementation.
type GetForeignMastersDataSource struct {
	client *client.Client
}

// GetForeignMastersDataSourceModel describes the data source state shape.
type GetForeignMastersDataSourceModel struct {
	BestMaster     types.Int64  `tfsdk:"best_master" json:"bestMaster"`
	ForeignMasters types.List   `tfsdk:"foreign_masters" json:"foreignMasters"`
	PortId         types.String `tfsdk:"port_id" json:"portId"`
}

// NewGetForeignMastersDataSource returns a new instance of the generated data source.
func NewGetForeignMastersDataSource() datasource.DataSource {
	return &GetForeignMastersDataSource{}
}

// Metadata returns the data source type name.
func (d *GetForeignMastersDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_foreign_masters"
}

// Schema returns the data source schema.
func (d *GetForeignMastersDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "deprecated: use GET /ptp/portState/{portId}/foreignSource", Attributes: map[string]schema.Attribute{"best_master": schema.Int64Attribute{Computed: true}, "foreign_masters": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"announce_msg_count": schema.Int64Attribute{MarkdownDescription: "Number of Announce messages from the foreign master that have been received within a time window", Computed: true}, "clock_address": schema.StringAttribute{Computed: true}, "grandmaster_clock_identity": schema.StringAttribute{MarkdownDescription: "grandmaster's Clock Identifier as per the IEEE 1588 standard which is represented by an integer 8-octet array", Computed: true}, "grandmaster_priority1": schema.Int64Attribute{Computed: true}, "grandmaster_priority2": schema.Int64Attribute{Computed: true}, "last_delay_response": schema.StringAttribute{Computed: true}, "last_follow_up": schema.StringAttribute{Computed: true}, "last_sync": schema.StringAttribute{Computed: true}, "mtsd_scaled_avar": schema.Int64Attribute{Computed: true}, "port_identity": schema.StringAttribute{Computed: true}, "port_module_number": schema.Int64Attribute{Computed: true}, "port_number": schema.Int64Attribute{Computed: true}, "ptp_protocol": schema.StringAttribute{Computed: true}, "quality": schema.SingleNestedAttribute{MarkdownDescription: "Clock Quality determines the quality of a clock based on class and accuracy. Used to select master clock in Best Master Clock Algorithm. (deprecated: use primarySourceClockQuality instead of grandmasterClockQuality)", Computed: true, Attributes: map[string]schema.Attribute{"accuracy": schema.StringAttribute{MarkdownDescription: "Indicates the expected accuracy of a clock when it is the grandmaster or in the event it becomes the grandmaster", Computed: true}, "class": schema.Int64Attribute{MarkdownDescription: "Denotes the traceability of the time or frequency distributed by the grandmaster clock", Computed: true}, "offset": schema.Int64Attribute{MarkdownDescription: "Indicates the estimate of the variations of the local clock from a linear timescale when it is not synchronized to another clock using the protocol", Computed: true}}}, "steps_removed": schema.Int64Attribute{Computed: true}}}}, "port_id": schema.StringAttribute{MarkdownDescription: "id of the target device Port (format: boxId_slotId_port, example: 1_1_c1)", Required: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetForeignMastersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetForeignMastersDataSourceModel
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
func (d *GetForeignMastersDataSource) readRemote(ctx context.Context, config *GetForeignMastersDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/ptp/portState/{portId}/foreignMaster"
	reqPath = strings.ReplaceAll(reqPath, "{portId}", url.PathEscape(config.PortId.ValueString()))
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_foreign_masters", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_foreign_masters", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_foreign_masters", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_foreign_masters", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_foreign_masters", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_foreign_masters", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_foreign_masters", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_foreign_masters", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_foreign_masters", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_foreign_masters", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_foreign_masters", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_foreign_masters", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_foreign_masters", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_foreign_masters", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetForeignMastersDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
