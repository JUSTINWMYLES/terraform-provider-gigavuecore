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
	_ datasource.DataSource              = (*GetSystemDiagPsuPerSlotDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetSystemDiagPsuPerSlotDataSource)(nil)
)

// GetSystemDiagPsuPerSlotDataSource is the generated Terraform data source implementation.
type GetSystemDiagPsuPerSlotDataSource struct {
	client *client.Client
}

// GetSystemDiagPsuPerSlotDataSourceModel describes the data source state shape.
type GetSystemDiagPsuPerSlotDataSourceModel struct {
	AcOrDc                 types.String `tfsdk:"ac_or_dc" json:"acOrDc"`
	BusyFault              types.String `tfsdk:"busy_fault" json:"busyFault"`
	Capacity               types.String `tfsdk:"capacity"`
	CommunicationFault     types.String `tfsdk:"communication_fault" json:"communicationFault"`
	HardwareRevision       types.String `tfsdk:"hardware_revision" json:"hardwareRevision"`
	HardwareType           types.String `tfsdk:"hardware_type" json:"hardwareType"`
	InputCurrent           types.String `tfsdk:"input_current" json:"inputCurrent"`
	InputOverVoltageFault  types.String `tfsdk:"input_over_voltage_fault" json:"inputOverVoltageFault"`
	InputPower             types.String `tfsdk:"input_power" json:"inputPower"`
	InputUnderVoltageFault types.String `tfsdk:"input_under_voltage_fault" json:"inputUnderVoltageFault"`
	InputVoltage           types.String `tfsdk:"input_voltage" json:"inputVoltage"`
	MaxPermitInputVoltage  types.String `tfsdk:"max_permit_input_voltage" json:"maxPermitInputVoltage"`
	MinPermitInputVoltage  types.String `tfsdk:"min_permit_input_voltage" json:"minPermitInputVoltage"`
	OffFault               types.String `tfsdk:"off_fault" json:"offFault"`
	OutputCurrent          types.String `tfsdk:"output_current" json:"outputCurrent"`
	OutputOverCurrentFault types.String `tfsdk:"output_over_current_fault" json:"outputOverCurrentFault"`
	OutputOverVoltageFault types.String `tfsdk:"output_over_voltage_fault" json:"outputOverVoltageFault"`
	OutputPower            types.String `tfsdk:"output_power" json:"outputPower"`
	OutputVoltage          types.String `tfsdk:"output_voltage" json:"outputVoltage"`
	ProductCode            types.String `tfsdk:"product_code" json:"productCode"`
	PsuFanSpeed            types.String `tfsdk:"psu_fan_speed" json:"psuFanSpeed"`
	SerialNumber           types.String `tfsdk:"serial_number" json:"serialNumber"`
	SlotId                 types.String `tfsdk:"slot_id" json:"slotId"`
	Status                 types.String `tfsdk:"status"`
	Temperature1           types.String `tfsdk:"temperature1"`
	Temperature2           types.String `tfsdk:"temperature2"`
	Temperature3           types.String `tfsdk:"temperature3"`
	TemperatureFault       types.String `tfsdk:"temperature_fault" json:"temperatureFault"`
}

// NewGetSystemDiagPsuPerSlotDataSource returns a new instance of the generated data source.
func NewGetSystemDiagPsuPerSlotDataSource() datasource.DataSource {
	return &GetSystemDiagPsuPerSlotDataSource{}
}

// Metadata returns the data source type name.
func (d *GetSystemDiagPsuPerSlotDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_system_diag_psu_per_slot"
}

// Schema returns the data source schema.
func (d *GetSystemDiagPsuPerSlotDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "get system diagnostics PSU information per Slot Id", Attributes: map[string]schema.Attribute{"ac_or_dc": schema.StringAttribute{MarkdownDescription: "AC or DC. NA stands for Not Available", Computed: true}, "busy_fault": schema.StringAttribute{MarkdownDescription: "PSU Busy Fault", Computed: true}, "capacity": schema.StringAttribute{MarkdownDescription: "Capacity (in Watts)", Computed: true}, "communication_fault": schema.StringAttribute{MarkdownDescription: "Communication, Logic and Memory Fault", Computed: true}, "hardware_revision": schema.StringAttribute{MarkdownDescription: "Hardware Revision", Computed: true}, "hardware_type": schema.StringAttribute{MarkdownDescription: "Hardware Type", Computed: true}, "input_current": schema.StringAttribute{MarkdownDescription: "Input Current (in Ampere)", Computed: true}, "input_over_voltage_fault": schema.StringAttribute{MarkdownDescription: "Input Over Voltage Fault", Computed: true}, "input_power": schema.StringAttribute{MarkdownDescription: " Input Power (in Watts)", Computed: true}, "input_under_voltage_fault": schema.StringAttribute{MarkdownDescription: "Input Under Voltage Fault", Computed: true}, "input_voltage": schema.StringAttribute{MarkdownDescription: "Input Voltage (in Volts)", Computed: true}, "max_permit_input_voltage": schema.StringAttribute{MarkdownDescription: "Maximum Input Voltage (in Volts)", Computed: true}, "min_permit_input_voltage": schema.StringAttribute{MarkdownDescription: "Minimum Input Voltage (in Volts)", Computed: true}, "off_fault": schema.StringAttribute{MarkdownDescription: "PSU Off Fault", Computed: true}, "output_current": schema.StringAttribute{MarkdownDescription: "Output Current (in Ampere)", Computed: true}, "output_over_current_fault": schema.StringAttribute{MarkdownDescription: "Output Over Current Fault", Computed: true}, "output_over_voltage_fault": schema.StringAttribute{MarkdownDescription: "Output Over Voltage Fault", Computed: true}, "output_power": schema.StringAttribute{MarkdownDescription: " Output Power (in Watts)", Computed: true}, "output_voltage": schema.StringAttribute{MarkdownDescription: "Output Voltage (in Volts)", Computed: true}, "product_code": schema.StringAttribute{MarkdownDescription: "Product Code", Computed: true}, "psu_fan_speed": schema.StringAttribute{MarkdownDescription: " PSU Fan Speed (RPM)", Computed: true}, "serial_number": schema.StringAttribute{MarkdownDescription: "Serial Number", Computed: true}, "slot_id": schema.StringAttribute{MarkdownDescription: "Power Module Slot ID", Required: true}, "status": schema.StringAttribute{MarkdownDescription: "Current PSU Status", Computed: true}, "temperature1": schema.StringAttribute{MarkdownDescription: " Inlet Temperature (in Celsius)", Computed: true}, "temperature2": schema.StringAttribute{MarkdownDescription: " Heat Sink Temperature (in Celsius)", Computed: true}, "temperature3": schema.StringAttribute{MarkdownDescription: " Temperature (in Celsius)", Computed: true}, "temperature_fault": schema.StringAttribute{MarkdownDescription: "Temperature Fault", Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetSystemDiagPsuPerSlotDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetSystemDiagPsuPerSlotDataSourceModel
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
func (d *GetSystemDiagPsuPerSlotDataSource) readRemote(ctx context.Context, config *GetSystemDiagPsuPerSlotDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/system/diag/psu/{slotId}"
	reqPath = strings.ReplaceAll(reqPath, "{slotId}", url.PathEscape(config.SlotId.ValueString()))
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_system_diag_psu_per_slot", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_system_diag_psu_per_slot", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_system_diag_psu_per_slot", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_system_diag_psu_per_slot", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_system_diag_psu_per_slot", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_system_diag_psu_per_slot", "Entity Not Found. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_system_diag_psu_per_slot", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_system_diag_psu_per_slot", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_system_diag_psu_per_slot", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_system_diag_psu_per_slot", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetSystemDiagPsuPerSlotDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
