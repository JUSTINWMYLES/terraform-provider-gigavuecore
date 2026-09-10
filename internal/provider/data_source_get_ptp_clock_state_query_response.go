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
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*GetPtpClockStateQueryResponseDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetPtpClockStateQueryResponseDataSource)(nil)
)

// GetPtpClockStateQueryResponseDataSource is the generated Terraform data source implementation.
type GetPtpClockStateQueryResponseDataSource struct {
	client *client.Client
}

// GetPtpClockStateQueryResponseDataSourceModel describes the data source state shape.
type GetPtpClockStateQueryResponseDataSourceModel struct {
	Alias            types.String `tfsdk:"alias"`
	BoxId            types.String `tfsdk:"box_id" json:"boxId"`
	ClockIdentity    types.String `tfsdk:"clock_identity" json:"clockIdentity"`
	ClockQuality     types.Object `tfsdk:"clock_quality" json:"clockQuality"`
	Domain           types.Int64  `tfsdk:"domain"`
	LocalClockTime   types.String `tfsdk:"local_clock_time" json:"localClockTime"`
	LocalPriority    types.Int64  `tfsdk:"local_priority" json:"localPriority"`
	MeanPathDelay    types.String `tfsdk:"mean_path_delay" json:"meanPathDelay"`
	Mode             types.String `tfsdk:"mode"`
	OffsetFromMaster types.String `tfsdk:"offset_from_master" json:"offsetFromMaster"`
	OffsetFromSource types.String `tfsdk:"offset_from_source" json:"offsetFromSource"`
	PortPtpCount     types.Int64  `tfsdk:"port_ptp_count" json:"portPtpCount"`
	Priority2        types.Int64  `tfsdk:"priority2"`
	PtpTime          types.String `tfsdk:"ptp_time" json:"PtpTime"`
	StepType         types.String `tfsdk:"step_type" json:"stepType"`
	StepsRemoved     types.Int64  `tfsdk:"steps_removed" json:"stepsRemoved"`
}

// NewGetPtpClockStateQueryResponseDataSource returns a new instance of the generated data source.
func NewGetPtpClockStateQueryResponseDataSource() datasource.DataSource {
	return &GetPtpClockStateQueryResponseDataSource{}
}

// Metadata returns the data source type name.
func (d *GetPtpClockStateQueryResponseDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_ptp_clock_state_query_response"
}

// Schema returns the data source schema.
func (d *GetPtpClockStateQueryResponseDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Reads the get ptp clock state query response data source.", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "alias of the time stamping PTP configuration", Required: true}, "box_id": schema.StringAttribute{MarkdownDescription: "device id", Computed: true}, "clock_identity": schema.StringAttribute{MarkdownDescription: "Clock Identifier as per the IEEE 1588 standard which is represented by an integer 8-octet array", Computed: true}, "clock_quality": schema.SingleNestedAttribute{MarkdownDescription: "Clock Quality determines the quality of a clock based on class and accuracy. Used to select source clock in Best Source Clock Algorithm", Computed: true, Attributes: map[string]schema.Attribute{"accuracy": schema.StringAttribute{MarkdownDescription: "Indicates the expected accuracy of a clock when it is the primary source or in the event it becomes the primary source", Computed: true}, "class": schema.Int64Attribute{MarkdownDescription: "Denotes the traceability of the time or frequency distributed by the primary source clock", Computed: true}, "offset": schema.Int64Attribute{MarkdownDescription: "Indicates the estimate of the variations of the local clock from a linear timescale when it is not synchronized to another clock using the protocol", Computed: true}}}, "domain": schema.Int64Attribute{Computed: true}, "local_clock_time": schema.StringAttribute{MarkdownDescription: "[RFC 3339](https://tools.ietf.org/html/rfc3339) format", Computed: true}, "local_priority": schema.Int64Attribute{Computed: true}, "mean_path_delay": schema.StringAttribute{MarkdownDescription: "Represents current value of the mean propagation time between source and receiver clock as computed by the slave. format hh:mm:ss.secfrac", Computed: true}, "mode": schema.StringAttribute{Computed: true}, "offset_from_master": schema.StringAttribute{MarkdownDescription: "Represents current value of time difference between master and a slave as computed by the slave. format hh:mm:ss. (deprecated: use offsetFromSource)", Computed: true}, "offset_from_source": schema.StringAttribute{MarkdownDescription: "Represents current value of time difference between source and a receiver as computed by the receiver. format hh:mm:ss", Computed: true}, "port_ptp_count": schema.Int64Attribute{MarkdownDescription: "Total number of PTP ports per node", Computed: true}, "priority2": schema.Int64Attribute{Computed: true}, "ptp_time": schema.StringAttribute{Computed: true}, "step_type": schema.StringAttribute{Computed: true}, "steps_removed": schema.Int64Attribute{MarkdownDescription: "Represents the number of communication paths traversed between the local clock and the primary source clock", Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetPtpClockStateQueryResponseDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetPtpClockStateQueryResponseDataSourceModel
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
func (d *GetPtpClockStateQueryResponseDataSource) readRemote(ctx context.Context, config *GetPtpClockStateQueryResponseDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/ptp/clockState/{alias}"
	reqPath = strings.ReplaceAll(reqPath, "{alias}", url.PathEscape(config.Alias.ValueString()))
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_clock_state_query_response", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_clock_state_query_response", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_clock_state_query_response", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_clock_state_query_response", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_clock_state_query_response", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_clock_state_query_response", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_clock_state_query_response", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_clock_state_query_response", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_clock_state_query_response", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_clock_state_query_response", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_clock_state_query_response", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_clock_state_query_response", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_clock_state_query_response", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["clockState"]; ok {
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_ptp_clock_state_query_response", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetPtpClockStateQueryResponseDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
