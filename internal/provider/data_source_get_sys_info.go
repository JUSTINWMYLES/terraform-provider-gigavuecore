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
	_ datasource.DataSource              = (*GetSysInfoDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetSysInfoDataSource)(nil)
)

// GetSysInfoDataSource is the generated Terraform data source implementation.
type GetSysInfoDataSource struct {
	client *client.Client
}

// GetSysInfoDataSourceModel describes the data source state shape.
type GetSysInfoDataSourceModel struct {
	AdditionalDns        types.String `tfsdk:"additional_dns" json:"additionalDns"`
	AdditionalDomainName types.String `tfsdk:"additional_domain_name" json:"additionalDomainName"`
	BootImage            types.String `tfsdk:"boot_image" json:"bootImage"`
	BuildDate            types.String `tfsdk:"build_date" json:"buildDate"`
	BuildId              types.String `tfsdk:"build_id" json:"buildId"`
	DefaultGateway       types.String `tfsdk:"default_gateway" json:"defaultGateway"`
	DefaultIpv6Gateway   types.String `tfsdk:"default_ipv6_gateway" json:"defaultIpv6Gateway"`
	DnsName              types.String `tfsdk:"dns_name" json:"dnsName"`
	DomainName           types.String `tfsdk:"domain_name" json:"domainName"`
	FmTime               types.String `tfsdk:"fm_time" json:"fmTime"`
	FmUptime             types.String `tfsdk:"fm_uptime" json:"fmUptime"`
	HostId               types.String `tfsdk:"host_id" json:"hostId"`
	Hostname             types.String `tfsdk:"hostname"`
	Mac                  types.String `tfsdk:"mac"`
	NtpIp                types.String `tfsdk:"ntp_ip" json:"ntpIp"`
	Platform             types.String `tfsdk:"platform"`
	PrimaryDns           types.String `tfsdk:"primary_dns" json:"primaryDns"`
	PrimaryIp            types.String `tfsdk:"primary_ip" json:"primaryIp"`
	PrimaryIpMask        types.String `tfsdk:"primary_ip_mask" json:"primaryIpMask"`
	PrimaryIpMaskLen     types.String `tfsdk:"primary_ip_mask_len" json:"primaryIpMaskLen"`
	PrimaryIpv6          types.String `tfsdk:"primary_ipv6" json:"primaryIpv6"`
	PrimaryIpv6Mask      types.String `tfsdk:"primary_ipv6_mask" json:"primaryIpv6Mask"`
	ProductName          types.String `tfsdk:"product_name" json:"productName"`
	SystemTimestamp      types.Int64  `tfsdk:"system_timestamp" json:"systemTimestamp"`
	SystemTimestampUtc   types.String `tfsdk:"system_timestamp_utc" json:"systemTimestampUtc"`
	Version              types.String `tfsdk:"version"`
}

// NewGetSysInfoDataSource returns a new instance of the generated data source.
func NewGetSysInfoDataSource() datasource.DataSource {
	return &GetSysInfoDataSource{}
}

// Metadata returns the data source type name.
func (d *GetSysInfoDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_sys_info"
}

// Schema returns the data source schema.
func (d *GetSysInfoDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Cms build info for FmHa", Attributes: map[string]schema.Attribute{"additional_dns": schema.StringAttribute{MarkdownDescription: "Additional Dns of build", Computed: true}, "additional_domain_name": schema.StringAttribute{MarkdownDescription: "Additional Domain Name of build", Computed: true}, "boot_image": schema.StringAttribute{MarkdownDescription: "Boot Image of build", Computed: true}, "build_date": schema.StringAttribute{MarkdownDescription: "Build Date of build", Computed: true}, "build_id": schema.StringAttribute{MarkdownDescription: "Build Id of build", Computed: true}, "default_gateway": schema.StringAttribute{MarkdownDescription: "Default Gateway of build", Computed: true}, "default_ipv6_gateway": schema.StringAttribute{MarkdownDescription: "Default Ipv6 Gateway of build", Computed: true}, "dns_name": schema.StringAttribute{MarkdownDescription: "DNS Name of build", Computed: true}, "domain_name": schema.StringAttribute{MarkdownDescription: "Domain Name of build", Computed: true}, "fm_time": schema.StringAttribute{MarkdownDescription: "FM time of build", Computed: true}, "fm_uptime": schema.StringAttribute{MarkdownDescription: "FM Up time of build", Computed: true}, "host_id": schema.StringAttribute{MarkdownDescription: "Host Id of build", Computed: true}, "hostname": schema.StringAttribute{MarkdownDescription: "DNS name or IP Address of build", Computed: true}, "mac": schema.StringAttribute{MarkdownDescription: "Mac of build", Computed: true}, "ntp_ip": schema.StringAttribute{MarkdownDescription: "ntp Ip of build", Computed: true}, "platform": schema.StringAttribute{MarkdownDescription: "Platform of build", Computed: true}, "primary_dns": schema.StringAttribute{MarkdownDescription: "Primary Dns of build", Computed: true}, "primary_ip": schema.StringAttribute{MarkdownDescription: "Primary Ip of build", Computed: true}, "primary_ip_mask": schema.StringAttribute{MarkdownDescription: "Primary Ip Mask of build", Computed: true}, "primary_ip_mask_len": schema.StringAttribute{MarkdownDescription: "Primary Ip Mask Len of build", Computed: true}, "primary_ipv6": schema.StringAttribute{MarkdownDescription: "Primary Ipv6 of build", Computed: true}, "primary_ipv6_mask": schema.StringAttribute{MarkdownDescription: "Primary Ipv6 Mask of build", Computed: true}, "product_name": schema.StringAttribute{MarkdownDescription: "Product Name of build", Computed: true}, "system_timestamp": schema.Int64Attribute{MarkdownDescription: "System time stamp of build", Computed: true}, "system_timestamp_utc": schema.StringAttribute{MarkdownDescription: "System time stamp Utc of build", Computed: true}, "version": schema.StringAttribute{MarkdownDescription: "Version of build", Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetSysInfoDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetSysInfoDataSourceModel
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
func (d *GetSysInfoDataSource) readRemote(ctx context.Context, config *GetSysInfoDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/fmHa/sys/info"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_sys_info", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_sys_info", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_sys_info", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_sys_info", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_sys_info", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_sys_info", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_sys_info", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_sys_info", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_sys_info", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_get_sys_info", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_get_sys_info", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_get_sys_info", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_sys_info", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	err = applyJSONToModel(&config, data)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_sys_info", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetSysInfoDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
