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
	_ datasource.DataSource              = (*LoadIpDestinationStatusDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadIpDestinationStatusDataSource)(nil)
)

// LoadIpDestinationStatusDataSource is the generated Terraform data source implementation.
type LoadIpDestinationStatusDataSource struct {
	client *client.Client
}

// LoadIpDestinationStatusDataSourceModel describes the data source state shape.
type LoadIpDestinationStatusDataSourceModel struct {
	ClusterId      types.String `tfsdk:"cluster_id" json:"clusterId"`
	Destination    types.String `tfsdk:"destination"`
	GsGroup        types.String `tfsdk:"gs_group" json:"gsGroup"`
	InterfaceAlias types.String `tfsdk:"interface_alias" json:"interfaceAlias"`
	IpAddress      types.String `tfsdk:"ip_address" json:"ipAddress"`
	Status         types.String `tfsdk:"status"`
	TeId           types.String `tfsdk:"te_id" json:"teId"`
}

// NewLoadIpDestinationStatusDataSource returns a new instance of the generated data source.
func NewLoadIpDestinationStatusDataSource() datasource.DataSource {
	return &LoadIpDestinationStatusDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadIpDestinationStatusDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_ip_destination_status"
}

// Schema returns the data source schema.
func (d *LoadIpDestinationStatusDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load IP destination status of the IP Address specified", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "destination": schema.StringAttribute{MarkdownDescription: "IP Destination", Computed: true}, "gs_group": schema.StringAttribute{MarkdownDescription: "Associated GsGroup", Computed: true}, "interface_alias": schema.StringAttribute{MarkdownDescription: "Interface Alias", Computed: true}, "ip_address": schema.StringAttribute{MarkdownDescription: "IP Address", Required: true}, "status": schema.StringAttribute{MarkdownDescription: "Status", Computed: true}, "te_id": schema.StringAttribute{MarkdownDescription: "Te-ID", Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadIpDestinationStatusDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadIpDestinationStatusDataSourceModel
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
func (d *LoadIpDestinationStatusDataSource) readRemote(ctx context.Context, config *LoadIpDestinationStatusDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/ip/destination/{ipAddress}"
	reqPath = strings.ReplaceAll(reqPath, "{ipAddress}", url.PathEscape(config.IpAddress.ValueString()))
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_ip_destination_status", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_ip_destination_status", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_ip_destination_status", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_ip_destination_status", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_ip_destination_status", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_ip_destination_status", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_ip_destination_status", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_ip_destination_status", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_ip_destination_status", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_ip_destination_status", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_load_ip_destination_status", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_load_ip_destination_status", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_ip_destination_status", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["ipDestinationStatus"]; ok {
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
		resp.Diagnostics.AddError("Error reading gigavuecore_load_ip_destination_status", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadIpDestinationStatusDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
