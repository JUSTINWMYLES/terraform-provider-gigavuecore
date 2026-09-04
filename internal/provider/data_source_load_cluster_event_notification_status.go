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
	_ datasource.DataSource              = (*LoadClusterEventNotificationStatusDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadClusterEventNotificationStatusDataSource)(nil)
)

// LoadClusterEventNotificationStatusDataSource is the generated Terraform data source implementation.
type LoadClusterEventNotificationStatusDataSource struct {
	client *client.Client
}

// LoadClusterEventNotificationStatusDataSourceModel describes the data source state shape.
type LoadClusterEventNotificationStatusDataSourceModel struct {
	ClusterId            types.String `tfsdk:"cluster_id" json:"clusterId"`
	ClusterName          types.String `tfsdk:"cluster_name" json:"clusterName"`
	EventNotifStatusList types.List   `tfsdk:"event_notif_status_list" json:"eventNotifStatusList"`
}

// NewLoadClusterEventNotificationStatusDataSource returns a new instance of the generated data source.
func NewLoadClusterEventNotificationStatusDataSource() datasource.DataSource {
	return &LoadClusterEventNotificationStatusDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadClusterEventNotificationStatusDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_cluster_event_notification_status"
}

// Schema returns the data source schema.
func (d *LoadClusterEventNotificationStatusDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load Cluster Event Notification Status", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Cluster ID whose event notification status is queried", Required: true}, "cluster_name": schema.StringAttribute{MarkdownDescription: "Name of the cluster", Computed: true}, "event_notif_status_list": schema.ListNestedAttribute{MarkdownDescription: "Status of event notification for every node in the cluster", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"device_ip": schema.StringAttribute{MarkdownDescription: "IP address of cluster member", Computed: true}, "enabled": schema.BoolAttribute{MarkdownDescription: "Should events be streamed from this device", Computed: true}, "event_targets_status": schema.ListNestedAttribute{MarkdownDescription: "Status of every notification target the device is configured to stream to", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"encoding": schema.StringAttribute{MarkdownDescription: "Encoding format the notification target is configured with ", Computed: true}, "error": schema.StringAttribute{MarkdownDescription: "Errors encountered when connecting to notification target", Computed: true}, "protocol": schema.StringAttribute{MarkdownDescription: "Messaging Protocol the notification target is configured with ", Computed: true}, "secured": schema.BoolAttribute{MarkdownDescription: "Is communication channel secured for the notification target", Computed: true}, "state": schema.StringAttribute{MarkdownDescription: "Indicates if the device can reach this target", Computed: true}, "target_address": schema.StringAttribute{MarkdownDescription: "IP address of the notification target", Computed: true}, "target_id": schema.StringAttribute{MarkdownDescription: "ID of the notification target", Computed: true}, "target_port": schema.Int64Attribute{MarkdownDescription: "The port where notification target is listening", Computed: true}, "username": schema.StringAttribute{MarkdownDescription: "Username of the notification target", Computed: true}}}}}}}}}
}

// Read fetches remote state into the data source model.
func (d *LoadClusterEventNotificationStatusDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadClusterEventNotificationStatusDataSourceModel
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
func (d *LoadClusterEventNotificationStatusDataSource) readRemote(ctx context.Context, config *LoadClusterEventNotificationStatusDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/system/eventNotification/targets"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_cluster_event_notification_status", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_cluster_event_notification_status", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_cluster_event_notification_status", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_cluster_event_notification_status", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_cluster_event_notification_status", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_cluster_event_notification_status", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_cluster_event_notification_status", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_cluster_event_notification_status", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_cluster_event_notification_status", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_cluster_event_notification_status", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_load_cluster_event_notification_status", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_load_cluster_event_notification_status", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_cluster_event_notification_status", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_cluster_event_notification_status", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadClusterEventNotificationStatusDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
