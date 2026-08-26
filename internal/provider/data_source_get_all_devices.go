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
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*GetAllDevicesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetAllDevicesDataSource)(nil)
)

// GetAllDevicesDataSource is the generated Terraform data source implementation.
type GetAllDevicesDataSource struct {
	client *client.Client
}

// GetAllDevicesDataSourceModel describes the data source state shape.
type GetAllDevicesDataSourceModel struct {
	ActivationId types.String `tfsdk:"activation_id" json:"activationId"`
	DeviceModel  types.String `tfsdk:"device_model" json:"deviceModel"`
	Items        types.List   `tfsdk:"items"`
}

// NewGetAllDevicesDataSource returns a new instance of the generated data source.
func NewGetAllDevicesDataSource() datasource.DataSource {
	return &GetAllDevicesDataSource{}
}

// Metadata returns the data source type name.
func (d *GetAllDevicesDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_all_devices"
}

// Schema returns the data source schema.
func (d *GetAllDevicesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get a list of physical devices on which license can be allocated, the selection criteria for devices determined by the query parameters", Attributes: map[string]schema.Attribute{"activation_id": schema.StringAttribute{MarkdownDescription: "activation ID of the license; if provided, devices that are relevant for the SKU / features of the activation are returned", Optional: true}, "device_model": schema.StringAttribute{MarkdownDescription: "model of device, used only if 'activationId' is null", Required: true}, "items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"box_id": schema.Int64Attribute{Computed: true}, "cards": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"box_id": schema.Int64Attribute{Computed: true}, "cluster_name": schema.StringAttribute{Computed: true}, "healthy": schema.BoolAttribute{Computed: true}, "hw_type": schema.StringAttribute{Computed: true}, "serial_number": schema.StringAttribute{Computed: true}, "slot_id": schema.StringAttribute{Computed: true}}}}, "cluster_name": schema.StringAttribute{Computed: true}, "deleted": schema.BoolAttribute{Computed: true}, "healthy": schema.BoolAttribute{Computed: true}, "host_id": schema.StringAttribute{Computed: true}, "host_name": schema.StringAttribute{Computed: true}, "model": schema.StringAttribute{MarkdownDescription: "Gigamon physical device models", Computed: true}, "node_id": schema.StringAttribute{Computed: true}, "product_version": schema.StringAttribute{Computed: true}, "serial_number": schema.StringAttribute{Computed: true}}}}}}
}

// Read fetches remote state into the data source model.
func (d *GetAllDevicesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetAllDevicesDataSourceModel
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
func (d *GetAllDevicesDataSource) readListRemote(ctx context.Context, config *GetAllDevicesDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/licensing/devices"
	params := url.Values{}
	params.Set("deviceModel", config.DeviceModel.ValueString())
	if !config.ActivationId.IsNull() {
		params.Set("activationId", config.ActivationId.ValueString())
	}
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_devices", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_devices", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["devices"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_devices", fmt.Sprintf("Could not decode list page: missing %q array", "devices"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_devices", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetAllDevicesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
