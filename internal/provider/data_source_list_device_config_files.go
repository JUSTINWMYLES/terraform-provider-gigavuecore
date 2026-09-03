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
	_ datasource.DataSource              = (*ListDeviceConfigFilesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*ListDeviceConfigFilesDataSource)(nil)
)

// ListDeviceConfigFilesDataSource is the generated Terraform data source implementation.
type ListDeviceConfigFilesDataSource struct {
	client *client.Client
}

// ListDeviceConfigFilesDataSourceModel describes the data source state shape.
type ListDeviceConfigFilesDataSourceModel struct {
	BoxId       types.String `tfsdk:"box_id" json:"boxId"`
	ConfigFiles types.List   `tfsdk:"config_files" json:"configFiles"`
	DeviceModel types.String `tfsdk:"device_model" json:"deviceModel"`
	NodeId      types.String `tfsdk:"node_id" json:"nodeId"`
	SwVersion   types.String `tfsdk:"sw_version" json:"swVersion"`
}

// NewListDeviceConfigFilesDataSource returns a new instance of the generated data source.
func NewListDeviceConfigFilesDataSource() datasource.DataSource {
	return &ListDeviceConfigFilesDataSource{}
}

// Metadata returns the data source type name.
func (d *ListDeviceConfigFilesDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_list_device_config_files"
}

// Schema returns the data source schema.
func (d *ListDeviceConfigFilesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "List existing configuration files on device", Attributes: map[string]schema.Attribute{"box_id": schema.StringAttribute{MarkdownDescription: "Cluster BoxId of the node", Computed: true}, "config_files": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"description": schema.StringAttribute{MarkdownDescription: "Description of the device configuration file", Computed: true}, "file_name": schema.StringAttribute{MarkdownDescription: "Config file name", Computed: true}, "file_size": schema.Int64Attribute{MarkdownDescription: "Config file size", Computed: true}, "last_modified": schema.StringAttribute{MarkdownDescription: "Last modification timestamp", Computed: true}, "next_boot_file": schema.BoolAttribute{MarkdownDescription: "indicates whether this config file should become active after reboot", Computed: true}, "running": schema.BoolAttribute{MarkdownDescription: "indicates whether this is the currently active config file", Computed: true}}}}, "device_model": schema.StringAttribute{MarkdownDescription: "Gigamon physical device models", Computed: true}, "node_id": schema.StringAttribute{MarkdownDescription: "Node ID", Required: true}, "sw_version": schema.StringAttribute{MarkdownDescription: "Device's current Software version", Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *ListDeviceConfigFilesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ListDeviceConfigFilesDataSourceModel
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
func (d *ListDeviceConfigFilesDataSource) readRemote(ctx context.Context, config *ListDeviceConfigFilesDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/clusterConfig/backup/device/{nodeId}"
	reqPath = strings.ReplaceAll(reqPath, "{nodeId}", url.PathEscape(config.NodeId.ValueString()))
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_list_device_config_files", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_list_device_config_files", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_list_device_config_files", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_list_device_config_files", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_list_device_config_files", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_list_device_config_files", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_list_device_config_files", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_list_device_config_files", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_list_device_config_files", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_list_device_config_files", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_list_device_config_files", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_list_device_config_files", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_list_device_config_files", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["configFiles"]; ok {
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
		resp.Diagnostics.AddError("Error reading gigavuecore_list_device_config_files", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *ListDeviceConfigFilesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
