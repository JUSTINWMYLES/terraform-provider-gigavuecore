package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)
import (
	client "github.com/JUSTINWMYLES/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*LoadSyslogConfigDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadSyslogConfigDataSource)(nil)
)

// LoadSyslogConfigDataSource is the generated Terraform data source implementation.
type LoadSyslogConfigDataSource struct {
	client *client.Client
}

// LoadSyslogConfigDataSourceModel describes the data source state shape.
type LoadSyslogConfigDataSourceModel struct {
	ClusterId        types.String `tfsdk:"cluster_id" json:"clusterId"`
	ClusterName      types.String `tfsdk:"cluster_name" json:"clusterName"`
	SyslogConfigList types.List   `tfsdk:"syslog_config_list" json:"syslogConfigList"`
}

// NewLoadSyslogConfigDataSource returns a new instance of the generated data source.
func NewLoadSyslogConfigDataSource() datasource.DataSource {
	return &LoadSyslogConfigDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadSyslogConfigDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_syslog_config"
}

// Schema returns the data source schema.
func (d *LoadSyslogConfigDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load Syslog config", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "cluster_name": schema.StringAttribute{MarkdownDescription: "Name of the cluster", Computed: true}, "syslog_config_list": schema.ListNestedAttribute{MarkdownDescription: "List of the syslog configuration specification for every node in the cluster", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"device_ip": schema.StringAttribute{MarkdownDescription: "IP address of the device", Computed: true}, "log_severity": schema.StringAttribute{MarkdownDescription: "Minimum syslog logging severity level", Computed: true}, "target_hosts": schema.ListNestedAttribute{MarkdownDescription: "List of syslog targets the device is streaming to", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"log_severity": schema.StringAttribute{MarkdownDescription: "Minimum logging severity level the device will stream for this target", Computed: true}, "port": schema.Int64Attribute{MarkdownDescription: "0 represents UDP and non zero is TCP", Computed: true}, "server": schema.StringAttribute{MarkdownDescription: "ipv4 or ipv6 or hostname", Computed: true}, "ssh_enabled": schema.BoolAttribute{MarkdownDescription: "Is Syslog target configured to receive logs via SSH", Computed: true}, "streaming_enabled": schema.BoolAttribute{MarkdownDescription: "Is Syslog Streaming enabled", Computed: true}, "username": schema.StringAttribute{MarkdownDescription: "For syslog over UDP there won't be any user. Only valid for SSH type", Computed: true}}}}}}}}}
}

// Read fetches remote state into the data source model.
func (d *LoadSyslogConfigDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadSyslogConfigDataSourceModel
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
func (d *LoadSyslogConfigDataSource) readRemote(ctx context.Context, config *LoadSyslogConfigDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/system/syslog/config"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_syslog_config", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_syslog_config", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_syslog_config", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_syslog_config", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_syslog_config", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_syslog_config", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_syslog_config", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_syslog_config", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_syslog_config", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_syslog_config", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_load_syslog_config", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_load_syslog_config", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_syslog_config", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_syslog_config", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadSyslogConfigDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
