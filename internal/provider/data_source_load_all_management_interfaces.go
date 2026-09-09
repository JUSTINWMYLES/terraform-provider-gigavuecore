package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)
import (
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*LoadAllManagementInterfacesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadAllManagementInterfacesDataSource)(nil)
)

// LoadAllManagementInterfacesDataSource is the generated Terraform data source implementation.
type LoadAllManagementInterfacesDataSource struct {
	client *client.Client
}

// LoadAllManagementInterfacesDataSourceModel describes the data source state shape.
type LoadAllManagementInterfacesDataSourceModel struct {
	BoxId         types.Int64  `tfsdk:"box_id" json:"boxId"`
	ClusterId     types.String `tfsdk:"cluster_id" json:"clusterId"`
	InterfaceName types.String `tfsdk:"interface_name" json:"interfaceName"`
	Items         types.List   `tfsdk:"items"`
}

// NewLoadAllManagementInterfacesDataSource returns a new instance of the generated data source.
func NewLoadAllManagementInterfacesDataSource() datasource.DataSource {
	return &LoadAllManagementInterfacesDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadAllManagementInterfacesDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_all_management_interfaces"
}

// Schema returns the data source schema.
func (d *LoadAllManagementInterfacesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "since FM 5.8", Attributes: map[string]schema.Attribute{"box_id": schema.Int64Attribute{MarkdownDescription: "boxId", Optional: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "interface_name": schema.StringAttribute{MarkdownDescription: "Interface Name", Optional: true}, "items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"box_id": schema.Int64Attribute{Computed: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "id of the defining cluster", Computed: true}, "discovery_protocol": schema.StringAttribute{Computed: true}, "g_arp": schema.BoolAttribute{MarkdownDescription: "Enable or disable Gratuitous ARP", Computed: true}, "interface_name": schema.StringAttribute{MarkdownDescription: "Management Interface Name", Computed: true}}}}}}
}

// Read fetches remote state into the data source model.
func (d *LoadAllManagementInterfacesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadAllManagementInterfacesDataSourceModel
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
func (d *LoadAllManagementInterfacesDataSource) readListRemote(ctx context.Context, config *LoadAllManagementInterfacesDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/system/interfaces/mgmt"
	params := url.Values{}
	params.Set("clusterId", config.ClusterId.ValueString())
	if !config.BoxId.IsNull() {
		params.Set("boxId", strconv.FormatInt(config.BoxId.ValueInt64(), 10))
	}
	if !config.InterfaceName.IsNull() {
		params.Set("interfaceName", config.InterfaceName.ValueString())
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
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_management_interfaces", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_management_interfaces", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["mgmtInterfaces"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_management_interfaces", fmt.Sprintf("Could not decode list page: missing %q array", "mgmtInterfaces"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_management_interfaces", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadAllManagementInterfacesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
