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
	_ datasource.DataSource              = (*GetPtpParentByAliasDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetPtpParentByAliasDataSource)(nil)
)

// GetPtpParentByAliasDataSource is the generated Terraform data source implementation.
type GetPtpParentByAliasDataSource struct {
	client *client.Client
}

// GetPtpParentByAliasDataSourceModel describes the data source state shape.
type GetPtpParentByAliasDataSourceModel struct {
	Alias                      types.String `tfsdk:"alias"`
	BoxId                      types.String `tfsdk:"box_id" json:"boxId"`
	ClockIdentity              types.String `tfsdk:"clock_identity" json:"clockIdentity"`
	ClockPortId                types.Int64  `tfsdk:"clock_port_id" json:"clockPortId"`
	GrandmasterClockIdentity   types.String `tfsdk:"grandmaster_clock_identity" json:"grandmasterClockIdentity"`
	GrandmasterClockQuality    types.Object `tfsdk:"grandmaster_clock_quality" json:"grandmasterClockQuality"`
	GrandmasterPriority1       types.Int64  `tfsdk:"grandmaster_priority1" json:"grandmasterPriority1"`
	GrandmasterPriority2       types.Int64  `tfsdk:"grandmaster_priority2" json:"grandmasterPriority2"`
	ObservedOffset             types.Int64  `tfsdk:"observed_offset" json:"observedOffset"`
	ObservedPhaseChangeRate    types.Int64  `tfsdk:"observed_phase_change_rate" json:"observedPhaseChangeRate"`
	PrimarySourceClockIdentity types.String `tfsdk:"primary_source_clock_identity" json:"primarySourceClockIdentity"`
	PrimarySourceClockQuality  types.Object `tfsdk:"primary_source_clock_quality" json:"primarySourceClockQuality"`
	PrimarySourcePriority1     types.Int64  `tfsdk:"primary_source_priority1" json:"primarySourcePriority1"`
	PrimarySourcePriority2     types.Int64  `tfsdk:"primary_source_priority2" json:"primarySourcePriority2"`
	Stats                      types.Bool   `tfsdk:"stats"`
}

// NewGetPtpParentByAliasDataSource returns a new instance of the generated data source.
func NewGetPtpParentByAliasDataSource() datasource.DataSource {
	return &GetPtpParentByAliasDataSource{}
}

// Metadata returns the data source type name.
func (d *GetPtpParentByAliasDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_ptp_parent_by_alias"
}

// Schema returns the data source schema.
func (d *GetPtpParentByAliasDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Reads the get ptp parent by alias data source.", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "alias of the time stamping PTP configuration", Required: true}, "box_id": schema.StringAttribute{Computed: true}, "clock_identity": schema.StringAttribute{MarkdownDescription: "Parent Clock Identifier as per the IEEE 1588 standard which is represented by an integer 8-octet array", Computed: true}, "clock_port_id": schema.Int64Attribute{MarkdownDescription: "Port Identity of the port on the source that issues the Sync message used in synchronizing this clock", Computed: true}, "grandmaster_clock_identity": schema.StringAttribute{MarkdownDescription: "grandmaster's Clock Identifier as per the IEEE 1588 standard which is represented by an integer 8-octet array. (deprecated: use primarySourceClockIdentity)", Computed: true}, "grandmaster_clock_quality": schema.SingleNestedAttribute{MarkdownDescription: "Clock Quality determines the quality of a clock based on class and accuracy. Used to select master clock in Best Master Clock Algorithm. (deprecated: use primarySourceClockQuality instead of grandmasterClockQuality)", Computed: true, Attributes: map[string]schema.Attribute{"accuracy": schema.StringAttribute{MarkdownDescription: "Indicates the expected accuracy of a clock when it is the grandmaster or in the event it becomes the grandmaster", Computed: true}, "class": schema.Int64Attribute{MarkdownDescription: "Denotes the traceability of the time or frequency distributed by the grandmaster clock", Computed: true}, "offset": schema.Int64Attribute{MarkdownDescription: "Indicates the estimate of the variations of the local clock from a linear timescale when it is not synchronized to another clock using the protocol", Computed: true}}}, "grandmaster_priority1": schema.Int64Attribute{MarkdownDescription: "deprecated: use primarySourcePriority1", Computed: true}, "grandmaster_priority2": schema.Int64Attribute{MarkdownDescription: "deprecated: use primarySourcePriority1", Computed: true}, "observed_offset": schema.Int64Attribute{MarkdownDescription: "Estimate of the parent's clock PTP variance as observed by the receiver clock", Computed: true}, "observed_phase_change_rate": schema.Int64Attribute{MarkdownDescription: "Estimate of the parent's clock phase rate change as observed by the receiver clock", Computed: true}, "primary_source_clock_identity": schema.StringAttribute{MarkdownDescription: "primary source's Clock Identifier as per the IEEE 1588 standard which is represented by an integer 8-octet array", Computed: true}, "primary_source_clock_quality": schema.SingleNestedAttribute{MarkdownDescription: "Clock Quality determines the quality of a clock based on class and accuracy. Used to select source clock in Best Source Clock Algorithm", Computed: true, Attributes: map[string]schema.Attribute{"accuracy": schema.StringAttribute{MarkdownDescription: "Indicates the expected accuracy of a clock when it is the primary source or in the event it becomes the primary source", Computed: true}, "class": schema.Int64Attribute{MarkdownDescription: "Denotes the traceability of the time or frequency distributed by the primary source clock", Computed: true}, "offset": schema.Int64Attribute{MarkdownDescription: "Indicates the estimate of the variations of the local clock from a linear timescale when it is not synchronized to another clock using the protocol", Computed: true}}}, "primary_source_priority1": schema.Int64Attribute{Computed: true}, "primary_source_priority2": schema.Int64Attribute{Computed: true}, "stats": schema.BoolAttribute{MarkdownDescription: "indicates whether the values of observedOffset and observedPhaseChangeRate have been measured and are valid", Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetPtpParentByAliasDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetPtpParentByAliasDataSourceModel
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
func (d *GetPtpParentByAliasDataSource) readRemote(ctx context.Context, config *GetPtpParentByAliasDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/ptp/parentClockState/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(config.Alias.ValueString()))
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_parent_by_alias", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_parent_by_alias", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_parent_by_alias", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_parent_by_alias", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_parent_by_alias", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_parent_by_alias", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_parent_by_alias", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_parent_by_alias", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_parent_by_alias", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_parent_by_alias", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_parent_by_alias", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_parent_by_alias", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_parent_by_alias", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["parent"]; ok {
		if m, ok := v.(map[string]any); ok {
			data = m
		}
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_parent_by_alias", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetPtpParentByAliasDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
