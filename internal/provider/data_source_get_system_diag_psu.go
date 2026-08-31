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
	_ datasource.DataSource              = (*GetSystemDiagPsuDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetSystemDiagPsuDataSource)(nil)
)

// GetSystemDiagPsuDataSource is the generated Terraform data source implementation.
type GetSystemDiagPsuDataSource struct {
	client *client.Client
}

// GetSystemDiagPsuDataSourceModel describes the data source state shape.
type GetSystemDiagPsuDataSourceModel struct {
	Items types.List `tfsdk:"items"`
}

// NewGetSystemDiagPsuDataSource returns a new instance of the generated data source.
func NewGetSystemDiagPsuDataSource() datasource.DataSource {
	return &GetSystemDiagPsuDataSource{}
}

// Metadata returns the data source type name.
func (d *GetSystemDiagPsuDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_system_diag_psu"
}

// Schema returns the data source schema.
func (d *GetSystemDiagPsuDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "get system diagnostics PSU information", Attributes: map[string]schema.Attribute{"items": schema.ListNestedAttribute{MarkdownDescription: "System Diag PSU Information", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"ac_or_dc": schema.StringAttribute{MarkdownDescription: "AC or DC. NA stands for Not Available", Computed: true}, "busy_fault": schema.StringAttribute{MarkdownDescription: "PSU Busy Fault", Computed: true}, "capacity": schema.StringAttribute{MarkdownDescription: "Capacity (in Watts)", Computed: true}, "communication_fault": schema.StringAttribute{MarkdownDescription: "Communication, Logic and Memory Fault", Computed: true}, "hardware_revision": schema.StringAttribute{MarkdownDescription: "Hardware Revision", Computed: true}, "hardware_type": schema.StringAttribute{MarkdownDescription: "Hardware Type", Computed: true}, "input_current": schema.StringAttribute{MarkdownDescription: "Input Current (in Ampere)", Computed: true}, "input_over_voltage_fault": schema.StringAttribute{MarkdownDescription: "Input Over Voltage Fault", Computed: true}, "input_power": schema.StringAttribute{MarkdownDescription: " Input Power (in Watts)", Computed: true}, "input_under_voltage_fault": schema.StringAttribute{MarkdownDescription: "Input Under Voltage Fault", Computed: true}, "input_voltage": schema.StringAttribute{MarkdownDescription: "Input Voltage (in Volts)", Computed: true}, "max_permit_input_voltage": schema.StringAttribute{MarkdownDescription: "Maximum Input Voltage (in Volts)", Computed: true}, "min_permit_input_voltage": schema.StringAttribute{MarkdownDescription: "Minimum Input Voltage (in Volts)", Computed: true}, "off_fault": schema.StringAttribute{MarkdownDescription: "PSU Off Fault", Computed: true}, "output_current": schema.StringAttribute{MarkdownDescription: "Output Current (in Ampere)", Computed: true}, "output_over_current_fault": schema.StringAttribute{MarkdownDescription: "Output Over Current Fault", Computed: true}, "output_over_voltage_fault": schema.StringAttribute{MarkdownDescription: "Output Over Voltage Fault", Computed: true}, "output_power": schema.StringAttribute{MarkdownDescription: " Output Power (in Watts)", Computed: true}, "output_voltage": schema.StringAttribute{MarkdownDescription: "Output Voltage (in Volts)", Computed: true}, "product_code": schema.StringAttribute{MarkdownDescription: "Product Code", Computed: true}, "psu_fan_speed": schema.StringAttribute{MarkdownDescription: " PSU Fan Speed (RPM)", Computed: true}, "serial_number": schema.StringAttribute{MarkdownDescription: "Serial Number", Computed: true}, "slot_id": schema.StringAttribute{MarkdownDescription: "Power Module Slot ID", Computed: true}, "status": schema.StringAttribute{MarkdownDescription: "Current PSU Status", Computed: true}, "temperature1": schema.StringAttribute{MarkdownDescription: " Inlet Temperature (in Celsius)", Computed: true}, "temperature2": schema.StringAttribute{MarkdownDescription: " Heat Sink Temperature (in Celsius)", Computed: true}, "temperature3": schema.StringAttribute{MarkdownDescription: " Temperature (in Celsius)", Computed: true}, "temperature_fault": schema.StringAttribute{MarkdownDescription: "Temperature Fault", Computed: true}}}}}}
}

// Read fetches remote state into the data source model.
func (d *GetSystemDiagPsuDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetSystemDiagPsuDataSourceModel
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
func (d *GetSystemDiagPsuDataSource) readListRemote(ctx context.Context, config *GetSystemDiagPsuDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/system/diag/psu"
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_system_diag_psu", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_system_diag_psu", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["systemPsuDiagDetails"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_system_diag_psu", fmt.Sprintf("Could not decode list page: missing %q array", "systemPsuDiagDetails"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_system_diag_psu", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetSystemDiagPsuDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
