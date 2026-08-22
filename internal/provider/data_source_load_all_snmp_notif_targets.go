package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)
import (
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*LoadAllSnmpNotifTargetsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LoadAllSnmpNotifTargetsDataSource)(nil)
)

// LoadAllSnmpNotifTargetsDataSource is the generated Terraform data source implementation.
type LoadAllSnmpNotifTargetsDataSource struct {
	client *client.Client
}

// LoadAllSnmpNotifTargetsDataSourceModel describes the data source state shape.
type LoadAllSnmpNotifTargetsDataSourceModel struct {
	ClusterId    types.String `tfsdk:"cluster_id" json:"clusterId"`
	Context      types.Object `tfsdk:"context"`
	NotifTargets types.List   `tfsdk:"notif_targets" json:"notifTargets"`
	Page         types.String `tfsdk:"page"`
	Sort         types.String `tfsdk:"sort"`
}

// NewLoadAllSnmpNotifTargetsDataSource returns a new instance of the generated data source.
func NewLoadAllSnmpNotifTargetsDataSource() datasource.DataSource {
	return &LoadAllSnmpNotifTargetsDataSource{}
}

// Metadata returns the data source type name.
func (d *LoadAllSnmpNotifTargetsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_all_snmp_notif_targets"
}

// Schema returns the data source schema.
func (d *LoadAllSnmpNotifTargetsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load all SNMP Notification Targets", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "context": schema.SingleNestedAttribute{MarkdownDescription: "Gigamon query result context", Computed: true, Attributes: map[string]schema.Attribute{"page_no": schema.Int64Attribute{MarkdownDescription: "page number of the returned result set", Computed: true}, "page_size": schema.Int64Attribute{MarkdownDescription: "page size of the returned result set", Computed: true}, "sort": schema.ListAttribute{MarkdownDescription: "sorting info of the returned result set. list of fields in the array indicate sorting order", Computed: true, ElementType: types.StringType}, "total_items": schema.Int64Attribute{MarkdownDescription: "total number of items in the queried entity type", Computed: true}}}, "notif_targets": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"enabled": schema.BoolAttribute{MarkdownDescription: "temporarily enable/disable the notification destination", Computed: true}, "host": schema.StringAttribute{MarkdownDescription: "ipv4 or ipv6 or domain name", Computed: true}, "notify_config": schema.SingleNestedAttribute{MarkdownDescription: "Notification Target configuration for specific Notification type (Trap/Inform)", Computed: true, Attributes: map[string]schema.Attribute{"auth_key": schema.StringAttribute{MarkdownDescription: "authentication password. required with 'v3user'", Computed: true}, "auth_protocol": schema.StringAttribute{MarkdownDescription: "authentication hash algorithm. required with 'v3user'", Computed: true}, "community": schema.StringAttribute{MarkdownDescription: "required when when 'version' is 'v2c'", Computed: true}, "engine_id": schema.StringAttribute{MarkdownDescription: "remote engineID. only valid with notifyType 'inform' and 'version' v3", Computed: true}, "port": schema.Int64Attribute{Computed: true}, "priv_key": schema.StringAttribute{MarkdownDescription: "privacy password", Computed: true}, "priv_protocol": schema.StringAttribute{MarkdownDescription: "privacy encryption", Computed: true}, "v3_user": schema.StringAttribute{MarkdownDescription: "required when when 'version' is 'v3'", Computed: true}, "version": schema.StringAttribute{MarkdownDescription: "SNMP version to use. v1 is only valid for traps. for v3, user name should be provided", Computed: true}}}, "notify_type": schema.StringAttribute{MarkdownDescription: "SNMP notification type to use", Computed: true}}}}, "page": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned", Optional: true}, "sort": schema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *LoadAllSnmpNotifTargetsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LoadAllSnmpNotifTargetsDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	d.readRemote(ctx, &config, resp)
	if resp.Diagnostics.HasError() {
		return
	}
	if config.Page.IsNull() {
		resp.Diagnostics.AddWarning("Single-page result", fmt.Sprintf("This data source reads a single page of a paginated API endpoint and does not aggregate results across pages. The \"page\" argument is unset, so the default page is returned; set it to retrieve a different page."))
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}

// readRemote performs the read HTTP exchange and decodes the response into config. Extracted from Read so the request/response logic is unit-testable without a tfsdk.Config.
func (d *LoadAllSnmpNotifTargetsDataSource) readRemote(ctx context.Context, config *LoadAllSnmpNotifTargetsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/system/snmp/notifTargets"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_snmp_notif_targets", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	query := httpReq.URL.Query()
	query.Set("clusterId", config.ClusterId.ValueString())
	if !config.Page.IsNull() {
		query.Set("page", config.Page.ValueString())
	}
	if !config.Sort.IsNull() {
		query.Set("sort", config.Sort.ValueString())
	}
	httpReq.URL.RawQuery = query.Encode()
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_snmp_notif_targets", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_snmp_notif_targets", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_snmp_notif_targets", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_snmp_notif_targets", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_snmp_notif_targets", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_snmp_notif_targets", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_snmp_notif_targets", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_snmp_notif_targets", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_snmp_notif_targets", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_load_all_snmp_notif_targets", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_load_all_snmp_notif_targets", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_snmp_notif_targets", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_load_all_snmp_notif_targets", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LoadAllSnmpNotifTargetsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
