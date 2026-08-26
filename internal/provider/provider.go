package provider

import "context"
import (
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/list"
	tfframeworkprovider "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

var _ tfframeworkprovider.Provider = (*gigavuecoreProvider)(nil)
var _ tfframeworkprovider.ProviderWithFunctions = (*gigavuecoreProvider)(nil)
var _ tfframeworkprovider.ProviderWithEphemeralResources = (*gigavuecoreProvider)(nil)
var _ tfframeworkprovider.ProviderWithListResources = (*gigavuecoreProvider)(nil)
var _ tfframeworkprovider.ProviderWithActions = (*gigavuecoreProvider)(nil)

// gigavuecoreProvider is the generated Terraform provider implementation.
type gigavuecoreProvider struct {
	configured bool
}

// gigavuecoreProviderModel describes the provider-level configuration shape.
type gigavuecoreProviderModel struct {
	Endpoint                  types.String `tfsdk:"endpoint"`
	Username                  types.String `tfsdk:"username"`
	Password                  types.String `tfsdk:"password"`
	LogFile                   types.String `tfsdk:"log_file"`
	LogCaptureRequestHeaders  types.Bool   `tfsdk:"log_capture_request_headers"`
	LogCaptureRequestBody     types.Bool   `tfsdk:"log_capture_request_body"`
	LogCaptureResponseHeaders types.Bool   `tfsdk:"log_capture_response_headers"`
	LogCaptureResponseBody    types.Bool   `tfsdk:"log_capture_response_body"`
	LogMaxBodyBytes           types.Int64  `tfsdk:"log_max_body_bytes"`
}

// New returns a new instance of the generated provider.
func New() tfframeworkprovider.Provider {
	return &gigavuecoreProvider{}
}

// Metadata returns the provider type name.
func (p *gigavuecoreProvider) Metadata(_ context.Context, _ tfframeworkprovider.MetadataRequest, resp *tfframeworkprovider.MetadataResponse) {
	resp.TypeName = "gigavuecore"
}

// Schema returns the provider configuration schema.
func (p *gigavuecoreProvider) Schema(_ context.Context, _ tfframeworkprovider.SchemaRequest, resp *tfframeworkprovider.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "This provider is generated using [eidos](https://github.com/signalbreak-labs/eidos) from the Gigamon GigaVUE-FM [OpenAPI](https://docs.gigamon.com/ref-api/Content/apiref_514onwards/doc/webapp/release/6.14.00%20GigaVUE-FM%20Core/openapi.fm.yaml) specification.\n\n## Example Usage\n\n```hcl\nprovider \"gigavuecore\" {\n  endpoint = \"https://fm.example.com\"\n  username = \"admin\"\n  password = \"changeme\"\n}\n```", Attributes: map[string]schema.Attribute{"endpoint": schema.StringAttribute{MarkdownDescription: "Overrides the default API base URL derived from the OpenAPI servers. Useful for directing the provider at a test or mock server.", Optional: true}, "username": schema.StringAttribute{MarkdownDescription: "Username for HTTP basic authentication.", Optional: true}, "password": schema.StringAttribute{MarkdownDescription: "Password for HTTP basic authentication.", Optional: true, Sensitive: true}, "log_file": schema.StringAttribute{MarkdownDescription: "Path to a file that receives HTTP request/response trace logs. When unset, trace logging is disabled.", Optional: true}, "log_capture_request_headers": schema.BoolAttribute{MarkdownDescription: "Capture request headers in the trace log. Sensitive headers are redacted.", Optional: true}, "log_capture_request_body": schema.BoolAttribute{MarkdownDescription: "Capture request bodies in the trace log. Disabled by default to avoid writing sensitive payloads to disk.", Optional: true}, "log_capture_response_headers": schema.BoolAttribute{MarkdownDescription: "Capture response headers in the trace log. Sensitive headers are redacted.", Optional: true}, "log_capture_response_body": schema.BoolAttribute{MarkdownDescription: "Capture response bodies in the trace log. Disabled by default to avoid writing sensitive payloads to disk.", Optional: true}, "log_max_body_bytes": schema.Int64Attribute{MarkdownDescription: "Maximum number of body bytes captured per log entry before truncation. Defaults to 4096.", Optional: true}}}
}

// Configure decodes practitioner configuration and marks the provider as configured.
func (p *gigavuecoreProvider) Configure(ctx context.Context, req tfframeworkprovider.ConfigureRequest, resp *tfframeworkprovider.ConfigureResponse) {
	var config gigavuecoreProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	opts := []client.ClientOption{}
	if !config.Endpoint.IsNull() && !config.Endpoint.IsUnknown() {
		opts = append(opts, client.WithBaseURL(config.Endpoint.ValueString()))
	}
	if !config.Username.IsNull() && !config.Username.IsUnknown() {
		opts = append(opts, client.WithSchemeInterceptor("basicAuth", client.BasicAuth(config.Username.ValueString(), config.Password.ValueString())))
	}
	loggingConfig := client.LoggingConfig{}
	if !config.LogFile.IsNull() && !config.LogFile.IsUnknown() {
		loggingConfig.LogFile = config.LogFile.ValueString()
	}
	if !config.LogCaptureRequestHeaders.IsNull() && !config.LogCaptureRequestHeaders.IsUnknown() {
		loggingConfig.CaptureRequestHeaders = config.LogCaptureRequestHeaders.ValueBool()
	}
	if !config.LogCaptureRequestBody.IsNull() && !config.LogCaptureRequestBody.IsUnknown() {
		loggingConfig.CaptureRequestBody = config.LogCaptureRequestBody.ValueBool()
	}
	if !config.LogCaptureResponseHeaders.IsNull() && !config.LogCaptureResponseHeaders.IsUnknown() {
		loggingConfig.CaptureResponseHeaders = config.LogCaptureResponseHeaders.ValueBool()
	}
	if !config.LogCaptureResponseBody.IsNull() && !config.LogCaptureResponseBody.IsUnknown() {
		loggingConfig.CaptureResponseBody = config.LogCaptureResponseBody.ValueBool()
	}
	if !config.LogMaxBodyBytes.IsNull() && !config.LogMaxBodyBytes.IsUnknown() {
		loggingConfig.MaxBodyBytes = int(config.LogMaxBodyBytes.ValueInt64())
	}
	if loggingConfig.LogFile != "" {
		opts = append(opts, client.WithLogging(loggingConfig))
	}
	c := client.New(opts...)
	resp.DataSourceData = c
	resp.ResourceData = c
	resp.EphemeralResourceData = c
	resp.ActionData = c
	resp.ListResourceData = c
	p.configured = true
}

// DataSources returns the data sources registered with this provider.
func (p *gigavuecoreProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{func() datasource.DataSource {
		return &GetAcmeCertificateDetailsDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllSystemAcmeCertificateDetailsDataSource{}
	}, func() datasource.DataSource {
		return &LoadSystemAcmeCertificateDetailsDataSource{}
	}, func() datasource.DataSource {
		return &GetAcmeServerDetailsDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllAlarmsDataSource{}
	}, func() datasource.DataSource {
		return &GetAllAlarmAutoSuppressionRulesDataSource{}
	}, func() datasource.DataSource {
		return &GetAllCorrelatedAlarmsDataSource{}
	}, func() datasource.DataSource {
		return &GetAlarmsCountByParameterDataSource{}
	}, func() datasource.DataSource {
		return &GetCorrelatedAlarmsCountByParameterDataSource{}
	}, func() datasource.DataSource {
		return &GetAllAlarmSuppressionMetadataDataSource{}
	}, func() datasource.DataSource {
		return &GetAlarmByIdDataSource{}
	}, func() datasource.DataSource {
		return &GetAllAppFilterRscDataSource{}
	}, func() datasource.DataSource {
		return &GetAppFilterRscBySlotIdDataSource{}
	}, func() datasource.DataSource {
		return &LoadDiameterWhitelistEntryDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllElbsDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllEnhancedSlicingsDataSource{}
	}, func() datasource.DataSource {
		return &GetAllAppsExporterDataSource{}
	}, func() datasource.DataSource {
		return &GetAllExporterGroupDataSource{}
	}, func() datasource.DataSource {
		return &LoadGtpWhitelistEntryDataSource{}
	}, func() datasource.DataSource {
		return &GetAllHsmsDataSource{}
	}, func() datasource.DataSource {
		return &GetHsmKeyMapsDataSource{}
	}, func() datasource.DataSource {
		return &GetRfsSyncDataSource{}
	}, func() datasource.DataSource {
		return &GetHsmGroupStatusReportDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllIcapProfilesDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllIcapServersDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllIcapServerGroupsDataSource{}
	}, func() datasource.DataSource {
		return &GetInlineSslConfigDataSource{}
	}, func() datasource.DataSource {
		return &LoadInlineSslCertValidStatusDataSource{}
	}, func() datasource.DataSource {
		return &LoadInlineSslCertValidRecordDataSource{}
	}, func() datasource.DataSource {
		return &LoadInlineSslUrlStatusDataSource{}
	}, func() datasource.DataSource {
		return &LoadInlineSslUrlRecordDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllInlineSslProfilesDataSource{}
	}, func() datasource.DataSource {
		return &FetchInlineSslProfileAllListInfoDataSource{}
	}, func() datasource.DataSource {
		return &FetchInlineSslProfileListInfoDataSource{}
	}, func() datasource.DataSource {
		return &DownloadListDataSource{}
	}, func() datasource.DataSource {
		return &GetInlineSslSigningDataSource{}
	}, func() datasource.DataSource {
		return &GetTrustStoreDataSource{}
	}, func() datasource.DataSource {
		return &GetTrustStoreCertDataSource{}
	}, func() datasource.DataSource {
		return &GetInlineSslTrustStoreFileDataSource{}
	}, func() datasource.DataSource {
		return &GetSslDecryptionKeysStoreStateDataSource{}
	}, func() datasource.DataSource {
		return &GetKeystoreKeysDataSource{}
	}, func() datasource.DataSource {
		return &GetKeystoreKeyCertDataSource{}
	}, func() datasource.DataSource {
		return &GetKeystorePreferenceDataSource{}
	}, func() datasource.DataSource {
		return &GetKeystoreCountersDataSource{}
	}, func() datasource.DataSource {
		return &GetAllAppsListenerDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllMetadataApplicationProfilesDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllMetadataCacheDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllMetadataExportersDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllMetadataNetworkProfilesDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllMetadataApplicationTemplatesDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllNetflowExportersDataSource{}
	}, func() datasource.DataSource {
		return &GetNetflowExporterFilterDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllNetflowMonitorsDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllNetflowRecordsDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllPortThrottlesDataSource{}
	}, func() datasource.DataSource {
		return &GetAllProxyServerProfilesDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllSessionAwareApfProfilesDataSource{}
	}, func() datasource.DataSource {
		return &LoadSipWhitelistEntryDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllSslDecryptionEndpointsDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllSslDecryptionKeyMapsDataSource{}
	}, func() datasource.DataSource {
		return &GetAllSslProfilesDataSource{}
	}, func() datasource.DataSource {
		return &GetSslClientTrustStoreListDataSource{}
	}, func() datasource.DataSource {
		return &GetSslClientTrustStoreDataSource{}
	}, func() datasource.DataSource {
		return &GetSslClientTrustStoreCertificateDataSource{}
	}, func() datasource.DataSource {
		return &GetSslClientTrustStoreFileDataSource{}
	}, func() datasource.DataSource {
		return &GetAllTcpProfilesDataSource{}
	}, func() datasource.DataSource {
		return &GetTrustStoreListDataSource{}
	}, func() datasource.DataSource {
		return &GetTrustStoreCertificateDataSource{}
	}, func() datasource.DataSource {
		return &GetSslTrustStoreFileDataSource{}
	}, func() datasource.DataSource {
		return &GetAllTunnelApplicationDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllAppVizSolutionsDataSource{}
	}, func() datasource.DataSource {
		return &LoadAuditLogEntriesDataSource{}
	}, func() datasource.DataSource {
		return &GetAuditLogEntryByIdDataSource{}
	}, func() datasource.DataSource {
		return &GetAvisiPoliciesDataSource{}
	}, func() datasource.DataSource {
		return &GetPoliciesInstantiationReportsDataSource{}
	}, func() datasource.DataSource {
		return &GetAvisiActionTemplDataSource{}
	}, func() datasource.DataSource {
		return &GetAvisiConditionTemplDataSource{}
	}, func() datasource.DataSource {
		return &GetReportInfoDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllCircuitTunnelGlobalDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllCircuitTunnelsDataSource{}
	}, func() datasource.DataSource {
		return &ListDeviceConfigFilesDataSource{}
	}, func() datasource.DataSource {
		return &DownloadBackupConfigFileOfFormatTextDataSource{}
	}, func() datasource.DataSource {
		return &DownloadConfigFileDataSource{}
	}, func() datasource.DataSource {
		return &GenerateRunningTextBackupConfigDataSource{}
	}, func() datasource.DataSource {
		return &ListFmStoredClustersConfigBackupSnapshotsDataSource{}
	}, func() datasource.DataSource {
		return &ListFmStoredClusterConfigBackupFileByTaskGroupIdDataSource{}
	}, func() datasource.DataSource {
		return &ListFmStoredClusterConfigBackupSnapshotsDataSource{}
	}, func() datasource.DataSource {
		return &ListFmStoredClusterConfigBackupSnapshotDataSource{}
	}, func() datasource.DataSource {
		return &ListBulkReplicateConfigFilesDataSource{}
	}, func() datasource.DataSource {
		return &DownloadBulkReplicateConfigFileDataSource{}
	}, func() datasource.DataSource {
		return &ListBulkReplicateConfigRestoreLogsDataSource{}
	}, func() datasource.DataSource {
		return &ListBulkReplicateConfigRestoreLogsByConfigFilenameDataSource{}
	}, func() datasource.DataSource {
		return &ListBulkReplicateConfigRestoreLogDataSource{}
	}, func() datasource.DataSource {
		return &GetClusterConfigStateDataSource{}
	}, func() datasource.DataSource {
		return &GsCardsDataSource{}
	}, func() datasource.DataSource {
		return &GetAllClusterConfigImageUpgradeStatusDataSource{}
	}, func() datasource.DataSource {
		return &GetAllClusterConfigImageUpgradeStatusByTaskGroupIdDataSource{}
	}, func() datasource.DataSource {
		return &GetClusterConfigImageUpgradeStatusByTaskGroupIdDataSource{}
	}, func() datasource.DataSource {
		return &GetClusterConfigImageUpgradeStatusByTaskIdDataSource{}
	}, func() datasource.DataSource {
		return &GetDeviceLocatorLedDetailsDataSource{}
	}, func() datasource.DataSource {
		return &GetUpgradeExecutionTimelineDataSource{}
	}, func() datasource.DataSource {
		return &GetUpgradeInfoByTaskIdDataSource{}
	}, func() datasource.DataSource {
		return &GetUpgradeInfoByTaskIdAndClusterIdDataSource{}
	}, func() datasource.DataSource {
		return &GetUpgradeJobInfoDataSource{}
	}, func() datasource.DataSource {
		return &GetUpgradeJobsDataSource{}
	}, func() datasource.DataSource {
		return &GetUpgradeSummaryDataSource{}
	}, func() datasource.DataSource {
		return &LoadEventsDataSource{}
	}, func() datasource.DataSource {
		return &GetEventsCountBySeverityDataSource{}
	}, func() datasource.DataSource {
		return &GetEventByIdDataSource{}
	}, func() datasource.DataSource {
		return &LoadFabricAdvHashDataSource{}
	}, func() datasource.DataSource {
		return &GetAllFabricMapsDataSource{}
	}, func() datasource.DataSource {
		return &LoadNrtStatsFabricMapDataSource{}
	}, func() datasource.DataSource {
		return &GetAllClusterCircuitTunnelsOfAnUserFabricMapDataSource{}
	}, func() datasource.DataSource {
		return &GetClusterCircuitTunnelOfAnUserFabricMapDataSource{}
	}, func() datasource.DataSource {
		return &GetAllClusterMapsOfAnUserFabricMapDataSource{}
	}, func() datasource.DataSource {
		return &GetClusterMapOfAnUserFabricMapDataSource{}
	}, func() datasource.DataSource {
		return &GetAllInternalFabricMapsDataSource{}
	}, func() datasource.DataSource {
		return &GetInternalFabricMapByAliasDataSource{}
	}, func() datasource.DataSource {
		return &GetAllClusterCircuitTunnelsOfAnInternalFabricMapDataSource{}
	}, func() datasource.DataSource {
		return &GetClusterCircuitTunnelOfAnInternalFabricMapDataSource{}
	}, func() datasource.DataSource {
		return &GetAllClusterMapsOfAnInternalFabricMapDataSource{}
	}, func() datasource.DataSource {
		return &GetClusterMapOfAnInternalFabricMapDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllFabricPortGroupsDataSource{}
	}, func() datasource.DataSource {
		return &GetFabricResourceConfigDataSource{}
	}, func() datasource.DataSource {
		return &GetL2CircuitResourceConfigDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllFilterResourcesDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllFilterResourcesBySlotIdDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllFilterTemplatesDataSource{}
	}, func() datasource.DataSource {
		return &LoadFilterTemplateLimitsDataSource{}
	}, func() datasource.DataSource {
		return &LoadFilterTemplateLimitBySlotIdDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllPlatformFilterLimitsDataSource{}
	}, func() datasource.DataSource {
		return &GetAllFlexInlineSolutionsDataSource{}
	}, func() datasource.DataSource {
		return &GetAllFlexInlineSslAppsDataSource{}
	}, func() datasource.DataSource {
		return &GetFlexInlineVlanConfigDataSource{}
	}, func() datasource.DataSource {
		return &GetAllGigaFlexInlineNetworkGroupDataSource{}
	}, func() datasource.DataSource {
		return &LoadNrtStatsFlexInlineSolutionDataSource{}
	}, func() datasource.DataSource {
		return &GetFlexInlineConfigDataSource{}
	}, func() datasource.DataSource {
		return &GetFlexInlineNetworkConfigDataSource{}
	}, func() datasource.DataSource {
		return &GetFlexInlineMapConfigDataSource{}
	}, func() datasource.DataSource {
		return &GetFlexInlineMapConfigStatusDataSource{}
	}, func() datasource.DataSource {
		return &GetFlowFilteringDeltaReportSummaryDataSource{}
	}, func() datasource.DataSource {
		return &LoadAlertPoliciesDataSource{}
	}, func() datasource.DataSource {
		return &ResourceSelectionDataSource{}
	}, func() datasource.DataSource {
		return &TestConnectionDataSource{}
	}, func() datasource.DataSource {
		return &GetCopilotSysdumpsDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllFmNotificationTargetConfigDataSource{}
	}, func() datasource.DataSource {
		return &GetAllNtpServerDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllFmTemplatesHierarchialConfigDataSource{}
	}, func() datasource.DataSource {
		return &GigaStreamThresholdDataSource{}
	}, func() datasource.DataSource {
		return &PortPacketThresholdDataSource{}
	}, func() datasource.DataSource {
		return &LoadsGlobalFmTemplateWithValuesDataSource{}
	}, func() datasource.DataSource {
		return &FmHaDetailsDataSource{}
	}, func() datasource.DataSource {
		return &GetClusterStateDataSource{}
	}, func() datasource.DataSource {
		return &GetFmHaRoleDataSource{}
	}, func() datasource.DataSource {
		return &GetFmHastatusDataSource{}
	}, func() datasource.DataSource {
		return &GetSysInfoDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllCriticalNotificationsDataSource{}
	}, func() datasource.DataSource {
		return &CreateApiTokenDataSource{}
	}, func() datasource.DataSource {
		return &LoadFmBackupArchiveServersDataSource{}
	}, func() datasource.DataSource {
		return &LoadFmBackupArchiveFilesDataSource{}
	}, func() datasource.DataSource {
		return &LoadGdpNodeReportsDataSource{}
	}, func() datasource.DataSource {
		return &GetGigaInsightNodesDataSource{}
	}, func() datasource.DataSource {
		return &GetEportsInterfacesDataSource{}
	}, func() datasource.DataSource {
		return &GetAllArpEntriesDataSource{}
	}, func() datasource.DataSource {
		return &GetArpEntriesDataSource{}
	}, func() datasource.DataSource {
		return &GetPingResultDataSource{}
	}, func() datasource.DataSource {
		return &GetAllGpfcpProfileDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllUserGroupsDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllGsGroupsDataSource{}
	}, func() datasource.DataSource {
		return &GetFlowDiameterReportsS6ASummaryDataSource{}
	}, func() datasource.DataSource {
		return &GetAllFlowFilteringSummaryDataSource{}
	}, func() datasource.DataSource {
		return &GetAllFlowSamplingSummaryDataSource{}
	}, func() datasource.DataSource {
		return &GetAllFlowSipSummaryDataSource{}
	}, func() datasource.DataSource {
		return &GetAllSslDecryptionSummaryDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllGtpPersistenceBackupRestoreDataSource{}
	}, func() datasource.DataSource {
		return &GetGtpBackupFilesDataSource{}
	}, func() datasource.DataSource {
		return &GetPortThrottleReportDataSource{}
	}, func() datasource.DataSource {
		return &GetFlowOpsReportDataSource{}
	}, func() datasource.DataSource {
		return &GetFlowDiameterReportS6ASummaryDataSource{}
	}, func() datasource.DataSource {
		return &GetFlowFilteringSummaryDataSource{}
	}, func() datasource.DataSource {
		return &GetFlowSamplingSummaryDataSource{}
	}, func() datasource.DataSource {
		return &GetFlowSipSummaryDataSource{}
	}, func() datasource.DataSource {
		return &GetSslDecryptionStatisticsDataSource{}
	}, func() datasource.DataSource {
		return &GetSslDecryptionSummaryDataSource{}
	}, func() datasource.DataSource {
		return &GetGtpPersistenceBackupRestoreByAliasDataSource{}
	}, func() datasource.DataSource {
		return &GetPortThrottleReportByAliasDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllGsopsDataSource{}
	}, func() datasource.DataSource {
		return &GetAllGtaProfilesDataSource{}
	}, func() datasource.DataSource {
		return &LoadBatteryOptimizationDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllGtapPortGroupDataSource{}
	}, func() datasource.DataSource {
		return &LoadGtapPortGroupDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllHeaderStripDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllHeaderStripAgingDataSource{}
	}, func() datasource.DataSource {
		return &GetAllicapClientDataSource{}
	}, func() datasource.DataSource {
		return &LoadNrtStatsIcapSolutionDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllImageRepoImagesDataSource{}
	}, func() datasource.DataSource {
		return &DownloadImageFileDataSource{}
	}, func() datasource.DataSource {
		return &GetImageServerRepoByVersionDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllImageServersDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllHeartbeatPacketsDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllHeartbeatProfilesDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllIbPathwaysDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllNegativeHeartbeatProfilesDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllInlineNetworkGroupsDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllInlineNetworkLagsDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllInlineNetworksDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllRedundancyProfileDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllInlineSerialToolGroupsDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllInlineToolGroupsDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllInlineToolsDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllAppVisibilitySolutionsDataSource{}
	}, func() datasource.DataSource {
		return &LoadAppVisibilitySolutionsByAliasWithConfigObjectsDataSource{}
	}, func() datasource.DataSource {
		return &GetAllIntentMobilityDataSource{}
	}, func() datasource.DataSource {
		return &LoadMobilityControlNodeDataSource{}
	}, func() datasource.DataSource {
		return &LoadMobilityControlNodeConfigsDataSource{}
	}, func() datasource.DataSource {
		return &LoadMobilityGtpNodeDataSource{}
	}, func() datasource.DataSource {
		return &LoadMobilityGtpNodeConfigsDataSource{}
	}, func() datasource.DataSource {
		return &LoadMobilitySamNodeDataSource{}
	}, func() datasource.DataSource {
		return &LoadMobilitySamNodeConfigsDataSource{}
	}, func() datasource.DataSource {
		return &LoadMobilityUserNodeDataSource{}
	}, func() datasource.DataSource {
		return &LoadMobilityUserNodeConfigsDataSource{}
	}, func() datasource.DataSource {
		return &GetPolicyDeploymentStatusDataSource{}
	}, func() datasource.DataSource {
		return &ExportPolicyDataSource{}
	}, func() datasource.DataSource {
		return &GetGsopInfoDataSource{}
	}, func() datasource.DataSource {
		return &GetToolsDataSource{}
	}, func() datasource.DataSource {
		return &GetPolicyDeploymentLogDataSource{}
	}, func() datasource.DataSource {
		return &GetPolicyFabricMapsDataSource{}
	}, func() datasource.DataSource {
		return &LoadGigaChassisDataSource{}
	}, func() datasource.DataSource {
		return &LoadCardsDetailsDataSource{}
	}, func() datasource.DataSource {
		return &LoadCardDetailsDataSource{}
	}, func() datasource.DataSource {
		return &LoadGigaPortsDataSource{}
	}, func() datasource.DataSource {
		return &GetPortNeighborsDataSource{}
	}, func() datasource.DataSource {
		return &GetGigaPortByPortIdDataSource{}
	}, func() datasource.DataSource {
		return &LoadIpDestinationStatusesDataSource{}
	}, func() datasource.DataSource {
		return &LoadIpDestinationStatusDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllIpInterfacesDataSource{}
	}, func() datasource.DataSource {
		return &GetAllIpInterfaceConfigsDataSource{}
	}, func() datasource.DataSource {
		return &GetEligibleRenewersDataSource{}
	}, func() datasource.DataSource {
		return &GetAllFeatureActivationsDataSource{}
	}, func() datasource.DataSource {
		return &GetExpiringCountDataSource{}
	}, func() datasource.DataSource {
		return &GetExpiringSoonAndRecentlyExpiredCountDataSource{}
	}, func() datasource.DataSource {
		return &GetAllFloatingFeatureActivationsDataSource{}
	}, func() datasource.DataSource {
		return &GetAllVblFeatureActivationsDataSource{}
	}, func() datasource.DataSource {
		return &LoadLicensingConfigDataSource{}
	}, func() datasource.DataSource {
		return &IsEmailServerConfiguredDataSource{}
	}, func() datasource.DataSource {
		return &GetAllDevicesDataSource{}
	}, func() datasource.DataSource {
		return &GetSerialNumberOfCardDataSource{}
	}, func() datasource.DataSource {
		return &GetSerialNumbersOfChassisRequiringGvosLicenseDataSource{}
	}, func() datasource.DataSource {
		return &GetEmsEntitlementsDataSource{}
	}, func() datasource.DataSource {
		return &GetEmsRegistrationStatusDataSource{}
	}, func() datasource.DataSource {
		return &GetEmsPreferencesDataSource{}
	}, func() datasource.DataSource {
		return &LoadFmLicensingSummaryDataSource{}
	}, func() datasource.DataSource {
		return &LoadFmLicensesDataSource{}
	}, func() datasource.DataSource {
		return &GetAllAllocationMapsDataSource{}
	}, func() datasource.DataSource {
		return &AlertableFloatingLicenseExpiriesDataSource{}
	}, func() datasource.DataSource {
		return &AlertableNodeLockedLicenseExpiriesDataSource{}
	}, func() datasource.DataSource {
		return &GetAllClusterLicensesDataSource{}
	}, func() datasource.DataSource {
		return &GetAllClusterLicensesLicensingModuleAllFlatDataSource{}
	}, func() datasource.DataSource {
		return &GetCurrentPhysicalAppsDataSource{}
	}, func() datasource.DataSource {
		return &ExpiryCountFlAndNllDataSource{}
	}, func() datasource.DataSource {
		return &GetAllUnlicensedChassisAndCardsDataSource{}
	}, func() datasource.DataSource {
		return &ExpiryNotifAndEmailAndCountDataSource{}
	}, func() datasource.DataSource {
		return &RefreshAllClusterLicensesDataSource{}
	}, func() datasource.DataSource {
		return &GetLicenseSkusDataSource{}
	}, func() datasource.DataSource {
		return &GetAppsInLastNPeriodsDataSource{}
	}, func() datasource.DataSource {
		return &GetBundlesInPeriodDataSource{}
	}, func() datasource.DataSource {
		return &GetCurrentAllowanceDataSource{}
	}, func() datasource.DataSource {
		return &GetCurrentAppTiersDataSource{}
	}, func() datasource.DataSource {
		return &GetCurrentAppsDataSource{}
	}, func() datasource.DataSource {
		return &GetCurrentMonSessionsToAppTierMapDataSource{}
	}, func() datasource.DataSource {
		return &GetCurrentPosIdsDataSource{}
	}, func() datasource.DataSource {
		return &GetCustomerNameDataSource{}
	}, func() datasource.DataSource {
		return &GetDeactivableVblsDataSource{}
	}, func() datasource.DataSource {
		return &GetDetailedReportForLastNPeriodsDataSource{}
	}, func() datasource.DataSource {
		return &DisplayPeriodsDataSource{}
	}, func() datasource.DataSource {
		return &SendEmailDataSource{}
	}, func() datasource.DataSource {
		return &SendEmail2DataSource{}
	}, func() datasource.DataSource {
		return &GetHighestVblBundleDataSource{}
	}, func() datasource.DataSource {
		return &GetLastNPeriodsSummaryVolumesDataSource{}
	}, func() datasource.DataSource {
		return &GetLastNPeriodsSummaryWithUnitDataSource{}
	}, func() datasource.DataSource {
		return &GetLatestPeriodVolumeDetailedAndPeriodsDataSource{}
	}, func() datasource.DataSource {
		return &GetPeriodVolumesDataSource{}
	}, func() datasource.DataSource {
		return &GetPeriodVolumeByDateDataSource{}
	}, func() datasource.DataSource {
		return &GetPeriodVolumeDetailedByDateDataSource{}
	}, func() datasource.DataSource {
		return &GetProcessedVolumesDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllMapChainsDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllMapAliasChainsDataSource{}
	}, func() datasource.DataSource {
		return &LoadMapAliasChainDataSource{}
	}, func() datasource.DataSource {
		return &GetMapGroupsDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllMapTemplatesDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllMapsDataSource{}
	}, func() datasource.DataSource {
		return &LoadClusterPortFilterCountersDataSource{}
	}, func() datasource.DataSource {
		return &GetAllPtpCountersDataSource{}
	}, func() datasource.DataSource {
		return &GetPtpCountersByAliasDataSource{}
	}, func() datasource.DataSource {
		return &GetAllPtpPortsCountersDataSource{}
	}, func() datasource.DataSource {
		return &GetPtpPortsCountersByPortIdDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllDeviceCredentialsDataSource{}
	}, func() datasource.DataSource {
		return &GetManagedClustersDataSource{}
	}, func() datasource.DataSource {
		return &GetManagedDevicesDataSource{}
	}, func() datasource.DataSource {
		return &LoadPermittedAddressesDataSource{}
	}, func() datasource.DataSource {
		return &LoadEmailServerConfigurationDataSource{}
	}, func() datasource.DataSource {
		return &GetEventNotificationCategoryDataSource{}
	}, func() datasource.DataSource {
		return &GetAllEventNotificationConfigurationDataSource{}
	}, func() datasource.DataSource {
		return &GetEventNotificationTemplateDataSource{}
	}, func() datasource.DataSource {
		return &GetAllExternalTrapReceiverDataSource{}
	}, func() datasource.DataSource {
		return &LoadAvailablePcapFilenamesDataSource{}
	}, func() datasource.DataSource {
		return &LoadPcapFileDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllGigastreamDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllPortConfigDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllPortFilterDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllPortGroupDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllPortPairDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllStackLinksDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllToolPortMirrorDataSource{}
	}, func() datasource.DataSource {
		return &GetAllPtpClockStatesDataSource{}
	}, func() datasource.DataSource {
		return &GetPtpClockStateQueryResponseDataSource{}
	}, func() datasource.DataSource {
		return &GetAllPtpParentDataSource{}
	}, func() datasource.DataSource {
		return &GetPtpParentByAliasDataSource{}
	}, func() datasource.DataSource {
		return &GetAllPtpPortsDataSource{}
	}, func() datasource.DataSource {
		return &GetAllForeignMastersDataSource{}
	}, func() datasource.DataSource {
		return &GetAllForeignSourcesDataSource{}
	}, func() datasource.DataSource {
		return &GetPtpPortByPortIdDataSource{}
	}, func() datasource.DataSource {
		return &GetForeignMastersDataSource{}
	}, func() datasource.DataSource {
		return &GetForeignSourcesDataSource{}
	}, func() datasource.DataSource {
		return &GetAllTimeStampingPtpConfigsDataSource{}
	}, func() datasource.DataSource {
		return &GetAllTimePropertiesDataSource{}
	}, func() datasource.DataSource {
		return &GetTimePropertiesByAliasDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllRolesDataSource{}
	}, func() datasource.DataSource {
		return &GetAllSffpProfileDataSource{}
	}, func() datasource.DataSource {
		return &LoadSnmpThrottleConfigDataSource{}
	}, func() datasource.DataSource {
		return &LoadSpineLinkAllDataSource{}
	}, func() datasource.DataSource {
		return &GetCryptoCaDataSource{}
	}, func() datasource.DataSource {
		return &GetAllExternalExportTargetServerDataSource{}
	}, func() datasource.DataSource {
		return &ValidateExternalExportServerDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllNameServersDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllNameServersGroupedByInterfaceDataSource{}
	}, func() datasource.DataSource {
		return &LoadAaaAuthConfigDataSource{}
	}, func() datasource.DataSource {
		return &LoadRemoteAuthSystemConfigDataSource{}
	}, func() datasource.DataSource {
		return &GetSystemArpEntriesDataSource{}
	}, func() datasource.DataSource {
		return &GetSystemArpRefreshIntervalDataSource{}
	}, func() datasource.DataSource {
		return &GetActiveConfigDataSource{}
	}, func() datasource.DataSource {
		return &DownloadNodeSystemConfigFileDataSource{}
	}, func() datasource.DataSource {
		return &GetNodeSystemConfigFilesDataSource{}
	}, func() datasource.DataSource {
		return &GetNodeSystemConfigTextFileDataSource{}
	}, func() datasource.DataSource {
		return &DownloadNodeSystemConfigTextFileDataSource{}
	}, func() datasource.DataSource {
		return &GetUserSelectedCiphersDataSource{}
	}, func() datasource.DataSource {
		return &GetCryptoStatusDataSource{}
	}, func() datasource.DataSource {
		return &GetFmCryptoStatusDataSource{}
	}, func() datasource.DataSource {
		return &GetSupportedCiphersDataSource{}
	}, func() datasource.DataSource {
		return &GetSystemDiagDataSource{}
	}, func() datasource.DataSource {
		return &GetSystemDiagPsuDataSource{}
	}, func() datasource.DataSource {
		return &GetSystemDiagPsuPerSlotDataSource{}
	}, func() datasource.DataSource {
		return &GetEmailNotifConfigSpecDataSource{}
	}, func() datasource.DataSource {
		return &GetEmailServerDataSource{}
	}, func() datasource.DataSource {
		return &LoadClusterEventNotificationStatusDataSource{}
	}, func() datasource.DataSource {
		return &LoadGsDumpDataSource{}
	}, func() datasource.DataSource {
		return &LoadGsDumpFileDataSource{}
	}, func() datasource.DataSource {
		return &GetSystemHostBannerDataSource{}
	}, func() datasource.DataSource {
		return &GetSystemHostnameDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllSystemInterfacesDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllManagementInterfacesDataSource{}
	}, func() datasource.DataSource {
		return &LoadManagementInterfaceNeighborsDataSource{}
	}, func() datasource.DataSource {
		return &LoadIpv6NeighborsDataSource{}
	}, func() datasource.DataSource {
		return &GetLicensedFeaturesFromNodeDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllLocalUsersDataSource{}
	}, func() datasource.DataSource {
		return &GetSystemNdpDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllRadiusServersDataSource{}
	}, func() datasource.DataSource {
		return &LoadSecurityConfigDataSource{}
	}, func() datasource.DataSource {
		return &LoadSnmpServerConfigDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllSnmpNotifTargetsDataSource{}
	}, func() datasource.DataSource {
		return &LoadSnmpThrottleStatusDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllV3UsersDataSource{}
	}, func() datasource.DataSource {
		return &GetSupportedSnmpV3UserProtocolsDataSource{}
	}, func() datasource.DataSource {
		return &GetSshCiphersDataSource{}
	}, func() datasource.DataSource {
		return &LoadSshSupportedParametersDataSource{}
	}, func() datasource.DataSource {
		return &LoadSysdumpDataSource{}
	}, func() datasource.DataSource {
		return &LoadSysdumpFileDataSource{}
	}, func() datasource.DataSource {
		return &LoadDeviceSyslogDataSource{}
	}, func() datasource.DataSource {
		return &LoadSyslogConfigDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllTacacsServersDataSource{}
	}, func() datasource.DataSource {
		return &GetNodeSystemTimeDataSource{}
	}, func() datasource.DataSource {
		return &LoadNtpConfigDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllNtpServersDataSource{}
	}, func() datasource.DataSource {
		return &LoadPpsSourceDataSource{}
	}, func() datasource.DataSource {
		return &LoadPtpConfigDataSource{}
	}, func() datasource.DataSource {
		return &GetCommonCiphersDataSource{}
	}, func() datasource.DataSource {
		return &GetSystemWebConfigurationDataSource{}
	}, func() datasource.DataSource {
		return &GetSystemWebProxyDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllUserTagsDataSource{}
	}, func() datasource.DataSource {
		return &LatestImportTagResultDataSource{}
	}, func() datasource.DataSource {
		return &GetAllTimestampBriefDataSource{}
	}, func() datasource.DataSource {
		return &GetUserTokensDataSource{}
	}, func() datasource.DataSource {
		return &GetUserTokensForPrivilegeUsersDataSource{}
	}, func() datasource.DataSource {
		return &LoadTopologyDataSource{}
	}, func() datasource.DataSource {
		return &LoadFabricPathDataSource{}
	}, func() datasource.DataSource {
		return &LoadTopologyVizConfigDataSource{}
	}, func() datasource.DataSource {
		return &LoadTopologyVizEndPointsDataSource{}
	}, func() datasource.DataSource {
		return &GetTopovizExternalDevicesDataSource{}
	}, func() datasource.DataSource {
		return &GetTopovizLinksDataSource{}
	}, func() datasource.DataSource {
		return &LoadTopologyVizNodesDataSource{}
	}, func() datasource.DataSource {
		return &GetAllTrafficFlowsDataSource{}
	}, func() datasource.DataSource {
		return &GetTrafficFlowsOverviewDataSource{}
	}, func() datasource.DataSource {
		return &GetCopiedRulesDataSource{}
	}, func() datasource.DataSource {
		return &GetTrafficFlowsGlobalSettingsDataSource{}
	}, func() datasource.DataSource {
		return &GetMapChainDataSource{}
	}, func() datasource.DataSource {
		return &GetMapGroupDataSource{}
	}, func() datasource.DataSource {
		return &GetMapMigrationResultsDataSource{}
	}, func() datasource.DataSource {
		return &GetAllMapMigrationResultsDataSource{}
	}, func() datasource.DataSource {
		return &GetAllMapMigrationResultsByAliasDataSource{}
	}, func() datasource.DataSource {
		return &GetPolicyNrtStatsDataSource{}
	}, func() datasource.DataSource {
		return &GetOverlapComponentsDataSource{}
	}, func() datasource.DataSource {
		return &GetConnectedLinksDataSource{}
	}, func() datasource.DataSource {
		return &GetConnectedNodesDataSource{}
	}, func() datasource.DataSource {
		return &GetConnectedLinksTrafficFlowsTroubleshootFlowsTopologyDataSource{}
	}, func() datasource.DataSource {
		return &GetTroubleshootSummaryDataSource{}
	}, func() datasource.DataSource {
		return &GetTroubleshootClusterConfigDataSource{}
	}, func() datasource.DataSource {
		return &GetDeployedDraftTrafficFlowsDataSource{}
	}, func() datasource.DataSource {
		return &GetTrafficPolicyGraphTunnelInterfaceMappingsDataSource{}
	}, func() datasource.DataSource {
		return &GetAllTrafficPolicyGraphDataSource{}
	}, func() datasource.DataSource {
		return &GetTrafficPolicyGraphStatusDataSource{}
	}, func() datasource.DataSource {
		return &QueryFabricMapTimeSeriesDataSource{}
	}, func() datasource.DataSource {
		return &QueryGsEngineTimeSeriesDataSource{}
	}, func() datasource.DataSource {
		return &QueryGsGroupTimeSeriesDataSource{}
	}, func() datasource.DataSource {
		return &QueryGsopTimeSeriesDataSource{}
	}, func() datasource.DataSource {
		return &QueryMapTimeSeriesDataSource{}
	}, func() datasource.DataSource {
		return &QueryNetflowExporterTimeSeriesDataSource{}
	}, func() datasource.DataSource {
		return &QueryNetflowMonitorTimeSeriesDataSource{}
	}, func() datasource.DataSource {
		return &QueryPortGroupLbTimeSeriesDataSource{}
	}, func() datasource.DataSource {
		return &QueryPortTimeSeriesDataSource{}
	}, func() datasource.DataSource {
		return &QueryForApplicationsDataSource{}
	}, func() datasource.DataSource {
		return &QueryDbForTrafficFlowsDataSource{}
	}, func() datasource.DataSource {
		return &QueryTunneledPortTimeSeriesDataSource{}
	}, func() datasource.DataSource {
		return &QueryVportTimeSeriesDataSource{}
	}, func() datasource.DataSource {
		return &LoadTunnelEndpointEntriesDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllTunnelEndpointsDataSource{}
	}, func() datasource.DataSource {
		return &GetAllTunnelLogicalGroupsDataSource{}
	}, func() datasource.DataSource {
		return &GetTunnelLogicalGroupStatsDataSource{}
	}, func() datasource.DataSource {
		return &ListEnvDataSource{}
	}, func() datasource.DataSource {
		return &GetEnvDataSource{}
	}, func() datasource.DataSource {
		return &ListConnectionsDataSource{}
	}, func() datasource.DataSource {
		return &GetMonitoringDomainDataSource{}
	}, func() datasource.DataSource {
		return &GetConnectionStatusDataSource{}
	}, func() datasource.DataSource {
		return &GetDeployDataSource{}
	}, func() datasource.DataSource {
		return &GetAllAppInfoDataSource{}
	}, func() datasource.DataSource {
		return &GetAppInfoDataSource{}
	}, func() datasource.DataSource {
		return &GetAppIntelTemplateDataSource{}
	}, func() datasource.DataSource {
		return &GetInterfacesDataSource{}
	}, func() datasource.DataSource {
		return &GetNetworkInterfacesDataSource{}
	}, func() datasource.DataSource {
		return &GetNodeInterfacesDataSource{}
	}, func() datasource.DataSource {
		return &GetNodeSysStatsDataSource{}
	}, func() datasource.DataSource {
		return &GetNodeVersionDataSource{}
	}, func() datasource.DataSource {
		return &GetNodesDataSource{}
	}, func() datasource.DataSource {
		return &GetNodeStatusDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllUsersDataSource{}
	}, func() datasource.DataSource {
		return &LoadActiveUserDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllVlanResourcesDataSource{}
	}, func() datasource.DataSource {
		return &LoadAllVportsDataSource{}
	}}
}

// Resources returns the managed resources registered with this provider.
func (p *gigavuecoreProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{func() resource.Resource {
		return &AdvHashResource{}
	}, func() resource.Resource {
		return &AlertPolicyResource{}
	}, func() resource.Resource {
		return &AppVisibilityResource{}
	}, func() resource.Resource {
		return &ApplicationProfileResource{}
	}, func() resource.Resource {
		return &ArchiveServerResource{}
	}, func() resource.Resource {
		return &CacheResource{}
	}, func() resource.Resource {
		return &CircuitTunnelResource{}
	}, func() resource.Resource {
		return &ConfigResource{}
	}, func() resource.Resource {
		return &ConnectionResource{}
	}, func() resource.Resource {
		return &ElbResource{}
	}, func() resource.Resource {
		return &EndpointResource{}
	}, func() resource.Resource {
		return &EngineResource{}
	}, func() resource.Resource {
		return &EnhancedSlicingResource{}
	}, func() resource.Resource {
		return &ExportTargetResource{}
	}, func() resource.Resource {
		return &ExporterResource{}
	}, func() resource.Resource {
		return &ExporterGroupResource{}
	}, func() resource.Resource {
		return &FabricMapResource{}
	}, func() resource.Resource {
		return &FabricPortGroupResource{}
	}, func() resource.Resource {
		return &FilterTemplateResource{}
	}, func() resource.Resource {
		return &FlexInlineResource{}
	}, func() resource.Resource {
		return &GigastreamResource{}
	}, func() resource.Resource {
		return &GpfcpProfileResource{}
	}, func() resource.Resource {
		return &GroupResource{}
	}, func() resource.Resource {
		return &GsGroupResource{}
	}, func() resource.Resource {
		return &GsopResource{}
	}, func() resource.Resource {
		return &GtaProfileResource{}
	}, func() resource.Resource {
		return &HbPacketResource{}
	}, func() resource.Resource {
		return &HbProfileResource{}
	}, func() resource.Resource {
		return &HeaderStripResource{}
	}, func() resource.Resource {
		return &HsmResource{}
	}, func() resource.Resource {
		return &HsmGroupResource{}
	}, func() resource.Resource {
		return &IbPathwayResource{}
	}, func() resource.Resource {
		return &IcapResource{}
	}, func() resource.Resource {
		return &ImageServerResource{}
	}, func() resource.Resource {
		return &InlineSslAppResource{}
	}, func() resource.Resource {
		return &InterfaceResource{}
	}, func() resource.Resource {
		return &KeyResource{}
	}, func() resource.Resource {
		return &KeyMapResource{}
	}, func() resource.Resource {
		return &ListenerResource{}
	}, func() resource.Resource {
		return &LocalUserResource{}
	}, func() resource.Resource {
		return &MapResource{}
	}, func() resource.Resource {
		return &MapGroupResource{}
	}, func() resource.Resource {
		return &MapTemplateResource{}
	}, func() resource.Resource {
		return &MobilityResource{}
	}, func() resource.Resource {
		return &MonitorResource{}
	}, func() resource.Resource {
		return &NegativeHbProfileResource{}
	}, func() resource.Resource {
		return &NetworkResource{}
	}, func() resource.Resource {
		return &NetworkGroupResource{}
	}, func() resource.Resource {
		return &NetworkLagResource{}
	}, func() resource.Resource {
		return &NetworkProfileResource{}
	}, func() resource.Resource {
		return &NodeResource{}
	}, func() resource.Resource {
		return &NotifMetaConfigResource{}
	}, func() resource.Resource {
		return &NotifTargetResource{}
	}, func() resource.Resource {
		return &PolicyResource{}
	}, func() resource.Resource {
		return &PortFilterResource{}
	}, func() resource.Resource {
		return &PortGroupResource{}
	}, func() resource.Resource {
		return &PortPairResource{}
	}, func() resource.Resource {
		return &PortThrottleResource{}
	}, func() resource.Resource {
		return &ProfileResource{}
	}, func() resource.Resource {
		return &PtpConfigResource{}
	}, func() resource.Resource {
		return &RadiusServerResource{}
	}, func() resource.Resource {
		return &RecordResource{}
	}, func() resource.Resource {
		return &RedundancyProfileResource{}
	}, func() resource.Resource {
		return &RoleResource{}
	}, func() resource.Resource {
		return &SaApfProfileResource{}
	}, func() resource.Resource {
		return &SerialToolGroupResource{}
	}, func() resource.Resource {
		return &ServerResource{}
	}, func() resource.Resource {
		return &ServerGroupResource{}
	}, func() resource.Resource {
		return &SffpProfileResource{}
	}, func() resource.Resource {
		return &SnmpTrapReceiverResource{}
	}, func() resource.Resource {
		return &SolutionResource{}
	}, func() resource.Resource {
		return &SourceRuleResource{}
	}, func() resource.Resource {
		return &SpineLinkResource{}
	}, func() resource.Resource {
		return &StackLinkResource{}
	}, func() resource.Resource {
		return &SubFlowResource{}
	}, func() resource.Resource {
		return &SysdumpResource{}
	}, func() resource.Resource {
		return &TacacsServerResource{}
	}, func() resource.Resource {
		return &TagResource{}
	}, func() resource.Resource {
		return &ToolResource{}
	}, func() resource.Resource {
		return &ToolGroupResource{}
	}, func() resource.Resource {
		return &ToolPortMirrorResource{}
	}, func() resource.Resource {
		return &TrafficFlowResource{}
	}, func() resource.Resource {
		return &TrafficPolicyGraphResource{}
	}, func() resource.Resource {
		return &TunnelApplicationResource{}
	}, func() resource.Resource {
		return &TunnelLbEndpointResource{}
	}, func() resource.Resource {
		return &V3UserResource{}
	}, func() resource.Resource {
		return &VportResource{}
	}, func() resource.Resource {
		return &GlobalResource{}
	}, func() resource.Resource {
		return &FmNtpServerResource{}
	}, func() resource.Resource {
		return &NtpServerResource{}
	}, func() resource.Resource {
		return &L2GreGroupResource{}
	}, func() resource.Resource {
		return &VxlanGroupResource{}
	}, func() resource.Resource {
		return &HeaderStripAgingResource{}
	}, func() resource.Resource {
		return &ProxyServerProfileResource{}
	}, func() resource.Resource {
		return &SslProfileResource{}
	}, func() resource.Resource {
		return &TcpProfileResource{}
	}, func() resource.Resource {
		return &NodeCredentialsResource{}
	}, func() resource.Resource {
		return &DiameterWhitelistResource{}
	}, func() resource.Resource {
		return &EmailRecipientResource{}
	}, func() resource.Resource {
		return &GtpWhitelistResource{}
	}, func() resource.Resource {
		return &InlineNetworkGroupResource{}
	}, func() resource.Resource {
		return &InlineSslProfileResource{}
	}, func() resource.Resource {
		return &LdapServerResource{}
	}, func() resource.Resource {
		return &MetadataExporterResource{}
	}, func() resource.Resource {
		return &NetflowExporterResource{}
	}, func() resource.Resource {
		return &FmTemplateResource{}
	}, func() resource.Resource {
		return &PcapProfileResource{}
	}, func() resource.Resource {
		return &IntentPolicyResource{}
	}, func() resource.Resource {
		return &SipWhitelistResource{}
	}, func() resource.Resource {
		return &UserResource{}
	}, func() resource.Resource {
		return &UserRoleResource{}
	}, func() resource.Resource {
		return &NotifConfigResource{}
	}, func() resource.Resource {
		return &ActivationResource{}
	}, func() resource.Resource {
		return &MapChainResource{}
	}, func() resource.Resource {
		return &PortConfigResource{}
	}, func() resource.Resource {
		return &CopilotConfigResource{}
	}, func() resource.Resource {
		return &FileResource{}
	}}
}

// Actions returns the actions registered with this provider.
func (p *gigavuecoreProvider) Actions(_ context.Context) []func() action.Action {
	return []func() action.Action{func() action.Action {
		return &DeleteAcmecertificateDetailsOfFmAction{}
	}, func() action.Action {
		return &IssueAcmeCertificateAction{}
	}, func() action.Action {
		return &DeleteAcmecertificateDetailsOfDeviceAction{}
	}, func() action.Action {
		return &ConfigureAcmeCertificateDetailsOfDeviceAction{}
	}, func() action.Action {
		return &UploadAcmeServerConfigDetailsAction{}
	}, func() action.Action {
		return &DeleteAcmeServerDetailsAction{}
	}, func() action.Action {
		return &ConfigureAcmeCertificateDetailsOfFmAction{}
	}, func() action.Action {
		return &EnableAlarmAutoSuppressionAction{}
	}, func() action.Action {
		return &AcknowledgeMultipleAlarmsAction{}
	}, func() action.Action {
		return &DeleteMultipleAlarmsAction{}
	}, func() action.Action {
		return &UnacknowledgeMultipleAlarmsAction{}
	}, func() action.Action {
		return &EnableAlarmSuppressionAction{}
	}, func() action.Action {
		return &EditAlarmSuppressionAction{}
	}, func() action.Action {
		return &DeleteAlarmSuppressionAction{}
	}, func() action.Action {
		return &UpdateAlarmAction{}
	}, func() action.Action {
		return &DeleteAlarmAction{}
	}, func() action.Action {
		return &CreateDiameterWhitelistEntriesAction{}
	}, func() action.Action {
		return &DeleteDiameterWhitelistEntryAction{}
	}, func() action.Action {
		return &UploadDiameterWhitelistEntriesFromFileAction{}
	}, func() action.Action {
		return &UploadDiameterWhitelistEntriesFromRemoteAction{}
	}, func() action.Action {
		return &DeleteAllElbAction{}
	}, func() action.Action {
		return &DeleteAllEnhancedSlicingAction{}
	}, func() action.Action {
		return &DeleteAllAppsExporterAction{}
	}, func() action.Action {
		return &DeleteAllAppsExporterGroupAction{}
	}, func() action.Action {
		return &CreateGtpWhitelistEntriesAction{}
	}, func() action.Action {
		return &DeleteGtpWhitelistEntryAction{}
	}, func() action.Action {
		return &DeleteAllGtpWhitelistEntriesAction{}
	}, func() action.Action {
		return &UploadGtpWhitelistEntriesFromFileAction{}
	}, func() action.Action {
		return &UploadGtpWhitelistEntriesFromUrlAction{}
	}, func() action.Action {
		return &DeleteAllHsmAction{}
	}, func() action.Action {
		return &DeleteAllHsmGroupAction{}
	}, func() action.Action {
		return &AddHsmToHsmGroupAction{}
	}, func() action.Action {
		return &DeleteHsmFromHsmGroupAction{}
	}, func() action.Action {
		return &FetchKeyHandlerRemoteAction{}
	}, func() action.Action {
		return &UploadKeyHandlerFromLocalAction{}
	}, func() action.Action {
		return &AddKeyMapToHsmGroupAction{}
	}, func() action.Action {
		return &DelKeyMapFromHsmGroupAction{}
	}, func() action.Action {
		return &FetchKeyMapAction{}
	}, func() action.Action {
		return &UploadKeyMapFromLocalAction{}
	}, func() action.Action {
		return &ActionRfsSyncNowAction{}
	}, func() action.Action {
		return &SetRfsSyncServerAction{}
	}, func() action.Action {
		return &PatchRtfsSyncServerAction{}
	}, func() action.Action {
		return &DeleteAllIcapProfilesAction{}
	}, func() action.Action {
		return &DeleteAllIcapServersAction{}
	}, func() action.Action {
		return &DeleteAllIcapServerGroupsAction{}
	}, func() action.Action {
		return &UpdateInlineSslConfigAction{}
	}, func() action.Action {
		return &ClearInlineSslCertValidCacheAction{}
	}, func() action.Action {
		return &ClearInlineSslUrlCacheAction{}
	}, func() action.Action {
		return &DeleteAllInlineSslProfilesAction{}
	}, func() action.Action {
		return &ReplaceInlineSslProfileAction{}
	}, func() action.Action {
		return &AddProfileDecryptPortMapAction{}
	}, func() action.Action {
		return &DeleteProfileDecryptPortMapAction{}
	}, func() action.Action {
		return &DeleteProfileDecryptPortMapByIdAction{}
	}, func() action.Action {
		return &ReplaceProfileDecryptPortMapAction{}
	}, func() action.Action {
		return &AddProfileKeyMapAction{}
	}, func() action.Action {
		return &DeleteProfileKeyMapsAction{}
	}, func() action.Action {
		return &DeleteProfileKeyMapByIdAction{}
	}, func() action.Action {
		return &ReplaceProfileKeyMapsAction{}
	}, func() action.Action {
		return &FetchInlineSslProfileListAction{}
	}, func() action.Action {
		return &DeleteInlineSslProfileListAction{}
	}, func() action.Action {
		return &CreateInlineSslProfileRuleAction{}
	}, func() action.Action {
		return &DeleteInlineSslProfileRuleAction{}
	}, func() action.Action {
		return &ReplaceAllInlineSslProfileRulesAction{}
	}, func() action.Action {
		return &UploadInlineSslSigningAction{}
	}, func() action.Action {
		return &DeleteSigningAction{}
	}, func() action.Action {
		return &UploadInlineSslTrustStoreAction{}
	}, func() action.Action {
		return &AddInlineSslTrustStoreCertAction{}
	}, func() action.Action {
		return &AppendInlineSslTrustStoreFileAction{}
	}, func() action.Action {
		return &DeleteTrustStoreCertAction{}
	}, func() action.Action {
		return &ResetInlineSslTrustStoreAction{}
	}, func() action.Action {
		return &UploadInlineSslTrustStoreFileAction{}
	}, func() action.Action {
		return &UnlockSslDecryptionKeyStoreAction{}
	}, func() action.Action {
		return &SetSslDecryptionKeyStorePasswordAction{}
	}, func() action.Action {
		return &ResetSslDecryptionKeysAction{}
	}, func() action.Action {
		return &DeleteKeystoreKeysAction{}
	}, func() action.Action {
		return &UploadKeystoreKeyAction{}
	}, func() action.Action {
		return &GenerateKeystoreKeyAction{}
	}, func() action.Action {
		return &UpdateKeystorePreferenceAction{}
	}, func() action.Action {
		return &CreateRiaKeystoreKeyAction{}
	}, func() action.Action {
		return &ClearKeystoreCountersAction{}
	}, func() action.Action {
		return &DeleteAllAppsListenerAction{}
	}, func() action.Action {
		return &RedefineMetadataExporterAction{}
	}, func() action.Action {
		return &CreateNetflowExporterFilterAction{}
	}, func() action.Action {
		return &RedefineNetflowExporterFilterAction{}
	}, func() action.Action {
		return &DeleteNetflowExporterFilterAction{}
	}, func() action.Action {
		return &AddNetflowExporterFilterRuleAction{}
	}, func() action.Action {
		return &RemoveNetflowExporterFilterRuleAction{}
	}, func() action.Action {
		return &DeleteAllPortThrottlesAction{}
	}, func() action.Action {
		return &DeleteAllAppsProxyServerProfileAction{}
	}, func() action.Action {
		return &AddSessionAwareApfFieldAction{}
	}, func() action.Action {
		return &CreateSipWhitelistEntriesAction{}
	}, func() action.Action {
		return &DeleteSipWhitelistEntryAction{}
	}, func() action.Action {
		return &DeleteAllSipWhitelistEntriesAction{}
	}, func() action.Action {
		return &UploadSipWhitelistEntriesFromFileAction{}
	}, func() action.Action {
		return &UploadSipWhitelistEntriesFromRemoteAction{}
	}, func() action.Action {
		return &DeleteSslDecryptionEndpointsAction{}
	}, func() action.Action {
		return &AddSslDecryptionKeyMappingAction{}
	}, func() action.Action {
		return &DeleteSslDecryptionKeyMappingAction{}
	}, func() action.Action {
		return &DeleteAllAppsSslProfileAction{}
	}, func() action.Action {
		return &CreateSslClientTrustStoreAction{}
	}, func() action.Action {
		return &DeleteClientTrustStoreAction{}
	}, func() action.Action {
		return &AppendSslClientTrustStoreCertAction{}
	}, func() action.Action {
		return &ReplaceSslClientTrustStoreCertAction{}
	}, func() action.Action {
		return &AppendSslClientTrustStoreFileAction{}
	}, func() action.Action {
		return &DeleteClientTrustStoreCertificateAction{}
	}, func() action.Action {
		return &UploadSslClientTrustStoreFileAction{}
	}, func() action.Action {
		return &DeleteAllAppsTcpProfileAction{}
	}, func() action.Action {
		return &UploadSslTrustStoreAction{}
	}, func() action.Action {
		return &AddSslTrustStoreCertAction{}
	}, func() action.Action {
		return &AppendSslTrustStoreFileAction{}
	}, func() action.Action {
		return &DeleteTrustStoreCertificateAction{}
	}, func() action.Action {
		return &ResetSslTrustStoreAction{}
	}, func() action.Action {
		return &UploadSslTrustStoreFileAction{}
	}, func() action.Action {
		return &PurgeAuditLogAction{}
	}, func() action.Action {
		return &DownloadAuditArchiveAction{}
	}, func() action.Action {
		return &BulkPolicyEnableAction{}
	}, func() action.Action {
		return &AddAvPolicyActionAction{}
	}, func() action.Action {
		return &DeleteAvisiPolicyActionAction{}
	}, func() action.Action {
		return &AddAvPolicyConditionAction{}
	}, func() action.Action {
		return &DeleteAvisiPolicyConditionAction{}
	}, func() action.Action {
		return &DeleteFabricMapsAction{}
	}, func() action.Action {
		return &DeleteAllTrafficFlowsAction{}
	}, func() action.Action {
		return &DeleteAllFlowRulesAction{}
	}, func() action.Action {
		return &DeleteAllSourceRulesAction{}
	}, func() action.Action {
		return &DeleteReportAction{}
	}, func() action.Action {
		return &GenerateReportAction{}
	}, func() action.Action {
		return &UpdateCircuitTunnelGlobalAction{}
	}, func() action.Action {
		return &DeleteAllCircuitTunnelL2GreAction{}
	}, func() action.Action {
		return &DeleteCircuitTunnelVxlanGroupsAction{}
	}, func() action.Action {
		return &DeleteAllCircuitTunnelsAction{}
	}, func() action.Action {
		return &CreateClusterConfigAction{}
	}, func() action.Action {
		return &BackupClustersConfigAction{}
	}, func() action.Action {
		return &BackupClusterConfigAction{}
	}, func() action.Action {
		return &ActivetDeviceConfigFileAction{}
	}, func() action.Action {
		return &DeleteDeviceConfigFileAction{}
	}, func() action.Action {
		return &UploadDeviceConfigFileAction{}
	}, func() action.Action {
		return &UpdateClusterConfigBackupAction{}
	}, func() action.Action {
		return &DeleteAllStoredConfigBackupSnapshotsAction{}
	}, func() action.Action {
		return &RestoreCluaterConfigBackupSnapshotAction{}
	}, func() action.Action {
		return &DeleteStoredClusterConfigBackupSnapshotAction{}
	}, func() action.Action {
		return &DeleteBulkReplicateConfigRestoreLogAction{}
	}, func() action.Action {
		return &UpgradeClusterGsCardImageAction{}
	}, func() action.Action {
		return &UpgradeClusterNodesImageAction{}
	}, func() action.Action {
		return &DeleteClusterConfigImageUpgradeStatusByTaskGroupIdAction{}
	}, func() action.Action {
		return &DeleteClusterConfigImageUpgradeStatusByTaskIdAction{}
	}, func() action.Action {
		return &ClusterConfigSwitchIPprotocolSpecAction{}
	}, func() action.Action {
		return &ClusterConfigAddMemberAction{}
	}, func() action.Action {
		return &ClusterConfigEditLeaderPreferenceAction{}
	}, func() action.Action {
		return &ClusterConfigDeleteMemberAction{}
	}, func() action.Action {
		return &RebootClusterNodesAction{}
	}, func() action.Action {
		return &UpdateStackingModeAction{}
	}, func() action.Action {
		return &UpdateGigaInsightNodePromptUpgradeAction{}
	}, func() action.Action {
		return &TestGigaInsightNodeConnectionAction{}
	}, func() action.Action {
		return &UpdateDeviceLocatorLedDetailsAction{}
	}, func() action.Action {
		return &CreateSpecAction{}
	}, func() action.Action {
		return &ResetDeviceUpgradeAction{}
	}, func() action.Action {
		return &ResumeDeviceUpgradeAction{}
	}, func() action.Action {
		return &RetryDeviceUpgradeAction{}
	}, func() action.Action {
		return &DeleteUpgradeJobsAction{}
	}, func() action.Action {
		return &ValidateDeviceUpgradeSpecAction{}
	}, func() action.Action {
		return &DeleteByScrollIdAction{}
	}, func() action.Action {
		return &PurgeEventAction{}
	}, func() action.Action {
		return &DownloadEventArchiveAction{}
	}, func() action.Action {
		return &RedefineFabricAdvHashAction{}
	}, func() action.Action {
		return &AuditFabricMapsAction{}
	}, func() action.Action {
		return &ConvertToFabricMapAction{}
	}, func() action.Action {
		return &RegisterNrtStatsFabricMapAction{}
	}, func() action.Action {
		return &ClearNrtStatsFabricMapAction{}
	}, func() action.Action {
		return &ReApplyAllFabricMapAction{}
	}, func() action.Action {
		return &ReplaceFabricMapWithLastUpdateTimestampValidationAction{}
	}, func() action.Action {
		return &UpdateFabricMapWithLastUpdateTimestampValidationAction{}
	}, func() action.Action {
		return &DeleteFabricMapWithLastUpdateTimestampValidationAction{}
	}, func() action.Action {
		return &UpdateFabricResourceConfigAction{}
	}, func() action.Action {
		return &UpdateL2CircuitResourceConfigAction{}
	}, func() action.Action {
		return &DeleteFilterTemplatesAction{}
	}, func() action.Action {
		return &ClearNrtStatsCountersAction{}
	}, func() action.Action {
		return &UpdateFlexInlineNrtStatsConfigAction{}
	}, func() action.Action {
		return &UpdateFlexInlineNetworkConfigAction{}
	}, func() action.Action {
		return &CreateFlexInlineConfigAction{}
	}, func() action.Action {
		return &ReplaceFlexInlineMapConfigAction{}
	}, func() action.Action {
		return &DeleteFlexInlineMapConfigAction{}
	}, func() action.Action {
		return &UpdateAlertPoliciesAction{}
	}, func() action.Action {
		return &DeleteAlertPoliciesAction{}
	}, func() action.Action {
		return &DeleteAllCopilotSysdumpFileAction{}
	}, func() action.Action {
		return &DeleteFmNtpServerAction{}
	}, func() action.Action {
		return &EnableFmntpServiceAction{}
	}, func() action.Action {
		return &ResetGlobalTemplateDefaultsAction{}
	}, func() action.Action {
		return &UploadCertificateAction{}
	}, func() action.Action {
		return &DeleteCertificateAction{}
	}, func() action.Action {
		return &AuditFmTemplateAction{}
	}, func() action.Action {
		return &ResetTemplateValueToGlobalAction{}
	}, func() action.Action {
		return &AddFmTemplateChildAction{}
	}, func() action.Action {
		return &DeleteFmTemplateChildAction{}
	}, func() action.Action {
		return &CreateHaGroupAction{}
	}, func() action.Action {
		return &FmBuildInfoAction{}
	}, func() action.Action {
		return &ServiceStateChangeAction{}
	}, func() action.Action {
		return &ChangeTunnelAuthModeAction{}
	}, func() action.Action {
		return &HaGroupNameAction{}
	}, func() action.Action {
		return &JoinHaGroupAction{}
	}, func() action.Action {
		return &LeaveHaInstanceAction{}
	}, func() action.Action {
		return &ReloadFmInstancesAction{}
	}, func() action.Action {
		return &ReloadFmInstanceAction{}
	}, func() action.Action {
		return &RemoveFmInstanceAction{}
	}, func() action.Action {
		return &ResetFmInstanceAction{}
	}, func() action.Action {
		return &UpdateFmInstanceAction{}
	}, func() action.Action {
		return &ConfigDiskUsageThresholdAction{}
	}, func() action.Action {
		return &CreateApiTokenFmSystemApiRateLimitingAction{}
	}, func() action.Action {
		return &BackupFmConfigAction{}
	}, func() action.Action {
		return &UpgradeFmImageAction{}
	}, func() action.Action {
		return &RebootFmAction{}
	}, func() action.Action {
		return &RestoreFmConfigAction{}
	}, func() action.Action {
		return &DeleteAllGsEngineInterfaceAction{}
	}, func() action.Action {
		return &DeleteAllIpsAction{}
	}, func() action.Action {
		return &DeleteAllGtpBackupFileAction{}
	}, func() action.Action {
		return &DeleteGtpBackupFileAction{}
	}, func() action.Action {
		return &UploadFlowOpsReportAction{}
	}, func() action.Action {
		return &RedefineGsGroupParamsAction{}
	}, func() action.Action {
		return &UpdateGsGroupParamsAction{}
	}, func() action.Action {
		return &RedefineGsGroupTcpParamsAction{}
	}, func() action.Action {
		return &RedefineGsGroupDedupParamsAction{}
	}, func() action.Action {
		return &RedefineGsGroupDiameterPacketAction{}
	}, func() action.Action {
		return &RedefineGsGroupDiameterS6ASessionAction{}
	}, func() action.Action {
		return &RedefineGsGroupDiameterWhitelistAction{}
	}, func() action.Action {
		return &RedefineGsGroupEflowParamsAction{}
	}, func() action.Action {
		return &RedefineGsGroupEngineWatchdogTimerAction{}
	}, func() action.Action {
		return &RedefineGsGroupGsGroupErspan3ParamsAction{}
	}, func() action.Action {
		return &RedefineGsGroupGsGroupFlowMaskParamsAction{}
	}, func() action.Action {
		return &RedefineGsGroupFlowSamplingParamsAction{}
	}, func() action.Action {
		return &RedefineGsGroupGenericSessionTimeoutAction{}
	}, func() action.Action {
		return &RedefineGsGroupGpfcpProfilesParamsAction{}
	}, func() action.Action {
		return &RedefineGsGroupGtaProfilesParamsAction{}
	}, func() action.Action {
		return &RedefineGsGroupParamsGtpControlSamplingAction{}
	}, func() action.Action {
		return &RedefineGsGroupGtpFlowParamsAction{}
	}, func() action.Action {
		return &RedefineGsGroupGtpGpfcpDelayParamsAction{}
	}, func() action.Action {
		return &RedefineGsGroupParamsGtpPersistenceParamsAction{}
	}, func() action.Action {
		return &RedefineGsGroupParamsGtpRandomSamplingAction{}
	}, func() action.Action {
		return &RedefineGsGroupGtpWhitelistParamsAction{}
	}, func() action.Action {
		return &RedefineGsGroupHealthCheckAction{}
	}, func() action.Action {
		return &RedefineGsGroupHsmGroupParamsAction{}
	}, func() action.Action {
		return &RedefineGsGroupIpFragParamsAction{}
	}, func() action.Action {
		return &RedefineGsGroupLoadBalanceParamsAction{}
	}, func() action.Action {
		return &RedefineGsGroupNetflowParamsAction{}
	}, func() action.Action {
		return &RedefineGsGroupNodeRoleParamsAction{}
	}, func() action.Action {
		return &RedefineGsGroupPortThrottleSipParamsAction{}
	}, func() action.Action {
		return &RedefineGsGroupGsGroupResourceParamsAction{}
	}, func() action.Action {
		return &RedefineGsGroupRtpPortsAction{}
	}, func() action.Action {
		return &RedefineGsGroupGsGroupSaApfParamsAction{}
	}, func() action.Action {
		return &RedefineGsGroupSessionLoggingAction{}
	}, func() action.Action {
		return &RedefineGsGroupSffpProfilesParamsAction{}
	}, func() action.Action {
		return &RedefineGsGroupSipMediaAction{}
	}, func() action.Action {
		return &RedefineGsGroupSipPortsAction{}
	}, func() action.Action {
		return &RedefineGsGroupSipSessionAction{}
	}, func() action.Action {
		return &RedefineGsGroupSipTcpIdleTimeoutAction{}
	}, func() action.Action {
		return &RedefineGsGroupSipWhitelistParamsAction{}
	}, func() action.Action {
		return &RedefineGsGroupSslDecryptParamsAction{}
	}, func() action.Action {
		return &RedefineGsGroupGsGroupSystemParamsAction{}
	}, func() action.Action {
		return &RedefineGsGroupXpktMatchAction{}
	}, func() action.Action {
		return &UpdateGsopAppsAction{}
	}, func() action.Action {
		return &DeleteAllGtaProfilesAction{}
	}, func() action.Action {
		return &UpdateBatteryOptimizationAction{}
	}, func() action.Action {
		return &DeleteAllHeaderStripAction{}
	}, func() action.Action {
		return &UpdateHeaderStripAction{}
	}, func() action.Action {
		return &ClearIcapNrtStatsCountersAction{}
	}, func() action.Action {
		return &UpdateIcapNrtStatsConfigAction{}
	}, func() action.Action {
		return &UploadImageFileAction{}
	}, func() action.Action {
		return &DeleteImageFileAction{}
	}, func() action.Action {
		return &UploadHeartbeatPacketFromFileAction{}
	}, func() action.Action {
		return &UploadNegativeHeartbeatPacketFromFileAction{}
	}, func() action.Action {
		return &RedefineInlineNetworkGroupAction{}
	}, func() action.Action {
		return &ClearHbStatisticsInlineToolAction{}
	}, func() action.Action {
		return &ClearNhbStatisticsInlineToolAction{}
	}, func() action.Action {
		return &ClearHbStatisticsInlineToolByAliasAction{}
	}, func() action.Action {
		return &ClearNhbStatisticsInlineToolByAliasAction{}
	}, func() action.Action {
		return &RecoverInlineToolAction{}
	}, func() action.Action {
		return &ClearPolicyAction{}
	}, func() action.Action {
		return &DeployPolicyAction{}
	}, func() action.Action {
		return &ImportPolicyAction{}
	}, func() action.Action {
		return &ValidatePolicyAction{}
	}, func() action.Action {
		return &ValidateUpdatedPolicyAction{}
	}, func() action.Action {
		return &ClearPolicyStatusAction{}
	}, func() action.Action {
		return &UndeployPolicyAction{}
	}, func() action.Action {
		return &UpdateDeployedPolicyAction{}
	}, func() action.Action {
		return &ConfigureAllChassisCardAction{}
	}, func() action.Action {
		return &UnconfigureAllChassisCardsAction{}
	}, func() action.Action {
		return &UpdateDeviceCardAction{}
	}, func() action.Action {
		return &ConfigureChassisCardAction{}
	}, func() action.Action {
		return &UnconfigureChassisCardAction{}
	}, func() action.Action {
		return &ConfigureDeviceChassisAction{}
	}, func() action.Action {
		return &ReconfigureDeviceChassisAction{}
	}, func() action.Action {
		return &UnconfigureDeviceChassisAction{}
	}, func() action.Action {
		return &UpdateDevicePortAction{}
	}, func() action.Action {
		return &CreateIpInterfaceSolutionAction{}
	}, func() action.Action {
		return &PatchIpInterfaceSolutionAction{}
	}, func() action.Action {
		return &DeleteAllIpInterfaceSolutionsAction{}
	}, func() action.Action {
		return &CreateBindingsAction{}
	}, func() action.Action {
		return &TransferBindingsAction{}
	}, func() action.Action {
		return &DeleteBindingsAction{}
	}, func() action.Action {
		return &ReplaceOldFmBindingsAction{}
	}, func() action.Action {
		return &RenewBindingsAction{}
	}, func() action.Action {
		return &DeactivateByAidAction{}
	}, func() action.Action {
		return &ConfigureExpiryAlertEnableAction{}
	}, func() action.Action {
		return &ConfigureVolumeUsageAlertEnableAction{}
	}, func() action.Action {
		return &UpdateEmailReceiversAction{}
	}, func() action.Action {
		return &SetEmsPreferencesAction{}
	}, func() action.Action {
		return &ReclaimEmsActivationsAction{}
	}, func() action.Action {
		return &RegisterWithEmsAction{}
	}, func() action.Action {
		return &UnregisterFromEmsAction{}
	}, func() action.Action {
		return &ImportLicenseAction{}
	}, func() action.Action {
		return &ImportRevocationAction{}
	}, func() action.Action {
		return &AddFmLicenseAction{}
	}, func() action.Action {
		return &RevokeFmLicenseAction{}
	}, func() action.Action {
		return &RemoveFmLicenseAction{}
	}, func() action.Action {
		return &RevokeLicenseKeyAction{}
	}, func() action.Action {
		return &RevokeDefaultVblAction{}
	}, func() action.Action {
		return &GetAggregateLinksAction{}
	}, func() action.Action {
		return &UpdateMapPriorityAction{}
	}, func() action.Action {
		return &DeleteMapGroupsAction{}
	}, func() action.Action {
		return &RemoveAllMapTemplateRulesAction{}
	}, func() action.Action {
		return &RemoveMapTemplateRuleAction{}
	}, func() action.Action {
		return &AddMapTemplateRuleAction{}
	}, func() action.Action {
		return &UpdateMapTemplateRuleAction{}
	}, func() action.Action {
		return &DeleteMapsAction{}
	}, func() action.Action {
		return &RemoveAllMapApRulesAction{}
	}, func() action.Action {
		return &RemoveMapApRuleAction{}
	}, func() action.Action {
		return &AddMapApRuleAction{}
	}, func() action.Action {
		return &UpdateMapApRuleAction{}
	}, func() action.Action {
		return &RemoveAllMapFlowRulesAction{}
	}, func() action.Action {
		return &RemoveMapFlowRuleAction{}
	}, func() action.Action {
		return &AddMapFlowRuleAction{}
	}, func() action.Action {
		return &UpdateMapFlowRuleAction{}
	}, func() action.Action {
		return &RemoveAllMapFlowSample5GOverlapRulesAction{}
	}, func() action.Action {
		return &AddMapFlowSample5GOverlapRuleAction{}
	}, func() action.Action {
		return &UpdateMapFlowSample5GOverlapRuleAction{}
	}, func() action.Action {
		return &RemoveMapFlowSample5GOverlapRuleAction{}
	}, func() action.Action {
		return &RemoveAllMapFlowSample5GRulesAction{}
	}, func() action.Action {
		return &AddMapFlowSample5GRuleAction{}
	}, func() action.Action {
		return &UpdateMapFlowSample5GRuleAction{}
	}, func() action.Action {
		return &RemoveMapFlowSample5GRuleAction{}
	}, func() action.Action {
		return &RemoveAllMapFlowSampleDiameterRulesAction{}
	}, func() action.Action {
		return &AddMapFlowSampleDiameterRuleAction{}
	}, func() action.Action {
		return &UpdateMapFlowSampleDiameterRuleAction{}
	}, func() action.Action {
		return &RemoveMapFlowSampleDiameterRuleAction{}
	}, func() action.Action {
		return &RemoveAllMapFlowSampleOverlapRulesAction{}
	}, func() action.Action {
		return &AddMapFlowSampleOverlapRuleAction{}
	}, func() action.Action {
		return &UpdateMapFlowSampleOverlapRuleAction{}
	}, func() action.Action {
		return &RemoveMapFlowSampleOverlapRuleAction{}
	}, func() action.Action {
		return &RemoveAllMapFlowSampleRulesAction{}
	}, func() action.Action {
		return &AddMapFlowSampleRuleAction{}
	}, func() action.Action {
		return &UpdateMapFlowSampleRuleAction{}
	}, func() action.Action {
		return &RemoveMapFlowSampleRuleAction{}
	}, func() action.Action {
		return &RemoveAllMapFlowSampleSipRulesAction{}
	}, func() action.Action {
		return &AddMapFlowSampleSipRuleAction{}
	}, func() action.Action {
		return &UpdateMapFlowSampleSipRuleAction{}
	}, func() action.Action {
		return &RemoveMapFlowSampleSipRuleAction{}
	}, func() action.Action {
		return &RemoveAllMapFlowWhitelistOverlapRulesAction{}
	}, func() action.Action {
		return &AddMapFlowWhitelistOverlapRuleAction{}
	}, func() action.Action {
		return &UpdateMapFlowWhitelistOverlapRuleAction{}
	}, func() action.Action {
		return &RemoveMapFlowWhitelistOverlapRuleAction{}
	}, func() action.Action {
		return &RemoveAllMapFlowWhitelistRulesAction{}
	}, func() action.Action {
		return &AddMapFlowWhitelistRuleAction{}
	}, func() action.Action {
		return &UpdateMapFlowWhitelistRuleAction{}
	}, func() action.Action {
		return &RemoveMapFlowWhitelistRuleAction{}
	}, func() action.Action {
		return &RemoveAllMapGsRulesAction{}
	}, func() action.Action {
		return &RemoveMapGsRuleAction{}
	}, func() action.Action {
		return &AddMapGsRuleAction{}
	}, func() action.Action {
		return &UpdateMapGsRuleAction{}
	}, func() action.Action {
		return &RemoveAllMapRulesAction{}
	}, func() action.Action {
		return &AddMapRulesFromTemplateAction{}
	}, func() action.Action {
		return &RemoveMapRuleAction{}
	}, func() action.Action {
		return &AddMapRuleAction{}
	}, func() action.Action {
		return &UpdateMapRuleAction{}
	}, func() action.Action {
		return &DeleteClusterPortFilterCountersAction{}
	}, func() action.Action {
		return &ClearAllPtpCountersAction{}
	}, func() action.Action {
		return &ClearPtpCountersByAliasAction{}
	}, func() action.Action {
		return &ClearPtpPortsCountersByPortIdAction{}
	}, func() action.Action {
		return &RedefineDeviceCredentialsAction{}
	}, func() action.Action {
		return &AddDevicesToDomainAction{}
	}, func() action.Action {
		return &RediscoverDomainClustersAction{}
	}, func() action.Action {
		return &RemoveClustersFromDomainAction{}
	}, func() action.Action {
		return &ClearTaskStatusAction{}
	}, func() action.Action {
		return &CommunicationAction{}
	}, func() action.Action {
		return &ConfigurePermittedAddressesAction{}
	}, func() action.Action {
		return &DeletePermittedAddressAction{}
	}, func() action.Action {
		return &ConfigureEmailServerAction{}
	}, func() action.Action {
		return &VerifyEmailServerAction{}
	}, func() action.Action {
		return &DisableBatchEventNotificationConfigurationAction{}
	}, func() action.Action {
		return &EnableBatchEventNotificationConfigurationAction{}
	}, func() action.Action {
		return &CreateEventNotificationConfigurationAction{}
	}, func() action.Action {
		return &CollectHeartbeatfromGigaInsightNodeAction{}
	}, func() action.Action {
		return &RegisterGigaInsightNodeAction{}
	}, func() action.Action {
		return &DeletePcapFileAction{}
	}, func() action.Action {
		return &RemoveHashToolMappingAction{}
	}, func() action.Action {
		return &UpdatePortConfigAction{}
	}, func() action.Action {
		return &RemovePortFilterRuleAction{}
	}, func() action.Action {
		return &AddPortFilterRuleAction{}
	}, func() action.Action {
		return &UpdatePortFilterRuleAction{}
	}, func() action.Action {
		return &DeleteAllSxIpsAction{}
	}, func() action.Action {
		return &RedefineSnmpThrottleConfigAction{}
	}, func() action.Action {
		return &UpdateSnmpThrottleConfigAction{}
	}, func() action.Action {
		return &DeleteSpineLinkAllAction{}
	}, func() action.Action {
		return &RemoveCertificateFromCaListAction{}
	}, func() action.Action {
		return &UploadCertificateSslNodeCertificateUploadNameAction{}
	}, func() action.Action {
		return &FmServerAaaConfigAction{}
	}, func() action.Action {
		return &AddNameServerAction{}
	}, func() action.Action {
		return &BulkUpdateNameserversAction{}
	}, func() action.Action {
		return &DeleteNameServerAction{}
	}, func() action.Action {
		return &AddSearchDomainAction{}
	}, func() action.Action {
		return &BulkUpdateSearchDomainAction{}
	}, func() action.Action {
		return &DeleteSearchDomainAction{}
	}, func() action.Action {
		return &RedefineAaaAuthConfigAction{}
	}, func() action.Action {
		return &UpdateAaaAuthConfigAction{}
	}, func() action.Action {
		return &ResetUserAccountLocksAction{}
	}, func() action.Action {
		return &RedefineRemoteAuthSystemConfigAction{}
	}, func() action.Action {
		return &UpdateRemoteAuthSystemConfigAction{}
	}, func() action.Action {
		return &RedefineLdapSystemConfigAction{}
	}, func() action.Action {
		return &DeleteLdapUserGroupMappingAction{}
	}, func() action.Action {
		return &RedefineRadiusSystemConfigAction{}
	}, func() action.Action {
		return &RedefineTacacsSystemConfigAction{}
	}, func() action.Action {
		return &ClearSystemArpEntriesAction{}
	}, func() action.Action {
		return &RedefineSystemArpRefreshIntervalAction{}
	}, func() action.Action {
		return &CreateNodeSystemConfigFileAction{}
	}, func() action.Action {
		return &UploadSystemConfigFileAction{}
	}, func() action.Action {
		return &FetchNodeSystemConfigFileAction{}
	}, func() action.Action {
		return &DeleteNodeSystemConfigFileAction{}
	}, func() action.Action {
		return &RevertNodeSystemConfigFileAction{}
	}, func() action.Action {
		return &SwitchNodeSystemConfigFileAction{}
	}, func() action.Action {
		return &GenerateNodeSystemConfigTextFileAction{}
	}, func() action.Action {
		return &UploadTextConfigFileAction{}
	}, func() action.Action {
		return &RestoreNodeSystemConfigTextFileAction{}
	}, func() action.Action {
		return &DeleteTextConfigFileAction{}
	}, func() action.Action {
		return &DeleteTrafficConfigAction{}
	}, func() action.Action {
		return &SetUserSelectedCiphersAction{}
	}, func() action.Action {
		return &EnableCryptoModeAction{}
	}, func() action.Action {
		return &EnableFipsModeAction{}
	}, func() action.Action {
		return &RunIntegrityCheckAction{}
	}, func() action.Action {
		return &UpdateEmailNotifConfigSpecAction{}
	}, func() action.Action {
		return &ModifyEmailServerAction{}
	}, func() action.Action {
		return &ModifyEventNotificationConfigAction{}
	}, func() action.Action {
		return &GenerateGigasmartDumpAction{}
	}, func() action.Action {
		return &DeleteGsDumpFileAction{}
	}, func() action.Action {
		return &UpdateSystemHostBannerAction{}
	}, func() action.Action {
		return &UpdateSystemHostnameAction{}
	}, func() action.Action {
		return &UpdateManagementInterfaceAction{}
	}, func() action.Action {
		return &ClearIpv6NeighborEntriesAction{}
	}, func() action.Action {
		return &InstallNodeLockedLicenseKeyOnNodeAction{}
	}, func() action.Action {
		return &RedefineSystemNdpAction{}
	}, func() action.Action {
		return &RedefineSecurityConfigAction{}
	}, func() action.Action {
		return &UpdateSecurityConfigAction{}
	}, func() action.Action {
		return &RedefineSnmpServerCommunityConfigAction{}
	}, func() action.Action {
		return &UpdateSnmpServerCommunityConfigAction{}
	}, func() action.Action {
		return &RedefineSnmpServerNofifyConfigAction{}
	}, func() action.Action {
		return &UpdateSnmpServerNofifyConfigAction{}
	}, func() action.Action {
		return &RedefineSnmpServerSystemConfigAction{}
	}, func() action.Action {
		return &UpdateSnmpServerSystemConfigAction{}
	}, func() action.Action {
		return &RedefineSnmpThrottleConfigSystemSnmpThrottleAction{}
	}, func() action.Action {
		return &UpdateSnmpThrottleConfigSystemSnmpThrottleAction{}
	}, func() action.Action {
		return &UpdateSshCiphersAction{}
	}, func() action.Action {
		return &GenerateSysdumpAction{}
	}, func() action.Action {
		return &DeleteSysdumpFileAction{}
	}, func() action.Action {
		return &PurgeSyslogAction{}
	}, func() action.Action {
		return &DownloadSyslogArchiveAction{}
	}, func() action.Action {
		return &RedefineSyslogConfigAction{}
	}, func() action.Action {
		return &UpdateSyslogConfigAction{}
	}, func() action.Action {
		return &UpdateNodeSystemTimeAction{}
	}, func() action.Action {
		return &RedefineNtpConfigAction{}
	}, func() action.Action {
		return &UpdateNtpConfigAction{}
	}, func() action.Action {
		return &AddNtpAuthKeyAction{}
	}, func() action.Action {
		return &DeleteNtpAuthKeyAction{}
	}, func() action.Action {
		return &RedefineNtpServerAction{}
	}, func() action.Action {
		return &NtpSyncAction{}
	}, func() action.Action {
		return &RedefinePpsSourceAction{}
	}, func() action.Action {
		return &RedefinePtpConfigAction{}
	}, func() action.Action {
		return &UbootInstallAction{}
	}, func() action.Action {
		return &UpdateSystemWebConfigurationAction{}
	}, func() action.Action {
		return &UpdateSystemWebProxyAction{}
	}, func() action.Action {
		return &AuditTagDetailsAction{}
	}, func() action.Action {
		return &ImportTagsAction{}
	}, func() action.Action {
		return &DownloadTagResourcesAction{}
	}, func() action.Action {
		return &ImportTagResourcesAction{}
	}, func() action.Action {
		return &RevokeApiTokensForPrivilegeUsersAction{}
	}, func() action.Action {
		return &RevokeAllApiTokensAction{}
	}, func() action.Action {
		return &CreateManualTopologyEntityAction{}
	}, func() action.Action {
		return &UpdateManualTopologyEntityAction{}
	}, func() action.Action {
		return &DeleteManualTopologyEntitiesAction{}
	}, func() action.Action {
		return &ImportManualTopologyEntityAction{}
	}, func() action.Action {
		return &DeleteManualTopologyAction{}
	}, func() action.Action {
		return &CreateTopologyLinkAction{}
	}, func() action.Action {
		return &UpdateTopologyLinkAction{}
	}, func() action.Action {
		return &DeleteTopologyLinkAction{}
	}, func() action.Action {
		return &CreateManualTopologyNodeAction{}
	}, func() action.Action {
		return &UpdateManualTopologyNodeAction{}
	}, func() action.Action {
		return &DeleteManualTopologyNodeAction{}
	}, func() action.Action {
		return &UpdateTopologyVizConfigsAction{}
	}, func() action.Action {
		return &GetSmartSankeyDataAction{}
	}, func() action.Action {
		return &GetHierarchialDataAction{}
	}, func() action.Action {
		return &LoadTopovizLinksAction{}
	}, func() action.Action {
		return &ResetTopologyVizAction{}
	}, func() action.Action {
		return &GetSankeyDataAction{}
	}, func() action.Action {
		return &AuditTrafficFlowGeneratedFMapAction{}
	}, func() action.Action {
		return &CopyRulesAction{}
	}, func() action.Action {
		return &DeleteCopiedRulesAction{}
	}, func() action.Action {
		return &DeployTrafficFlowsAction{}
	}, func() action.Action {
		return &AddTrafficFlowsDeployedDraftAction{}
	}, func() action.Action {
		return &UpdateTrafficFlowsGlobalSettingsAction{}
	}, func() action.Action {
		return &UpdateMapChainAction{}
	}, func() action.Action {
		return &UpdateMapChainPriorityAction{}
	}, func() action.Action {
		return &UpdateMapGroupAction{}
	}, func() action.Action {
		return &DeleteMapGroupTrafficFlowsMapGroupsPolicyAliasSourceRulesSourceRulesAliasAction{}
	}, func() action.Action {
		return &MigrateAppIntelToTrafficFlowsAction{}
	}, func() action.Action {
		return &RollbackTrafficFlowsAction{}
	}, func() action.Action {
		return &MigrateFabricMapsToTrafficFlows1Action{}
	}, func() action.Action {
		return &MigrateAllFabricMapsToTrafficFlowsAction{}
	}, func() action.Action {
		return &MigrateAllFlowMapToTrafficFlowsInMultipleClustersAction{}
	}, func() action.Action {
		return &MigrateFabricMapsToTrafficFlowsAction{}
	}, func() action.Action {
		return &MigrateAllFlowMapToTrafficFlowsInClusterAction{}
	}, func() action.Action {
		return &RollbackTrafficFlowsMigrationAction{}
	}, func() action.Action {
		return &RegisterOrDeRegisterNrtAction{}
	}, func() action.Action {
		return &PasteRulesAction{}
	}, func() action.Action {
		return &ReApplyTrafficFlowGeneratedFMapAction{}
	}, func() action.Action {
		return &GetPolicyRulesAction{}
	}, func() action.Action {
		return &AddTrafficFlowsDraftAction{}
	}, func() action.Action {
		return &BuildFlowPathAction{}
	}, func() action.Action {
		return &UndeployTrafficFlowsAction{}
	}, func() action.Action {
		return &DeleteDeployedDraftTrafficFlowsAction{}
	}, func() action.Action {
		return &UpdateTrafficFlowsDeployedDraftAction{}
	}, func() action.Action {
		return &AddFlowRulesAction{}
	}, func() action.Action {
		return &AddTemplateRulesAction{}
	}, func() action.Action {
		return &UpdateTrafficFlowsDraftAction{}
	}, func() action.Action {
		return &AddSourceRulesAction{}
	}, func() action.Action {
		return &DeployOrSaveTrafficPolicyGraphAction{}
	}, func() action.Action {
		return &DeployTrafficPolicyGraphAction{}
	}, func() action.Action {
		return &PurgeStatsAction{}
	}, func() action.Action {
		return &DeleteAllTunnelEndpointAction{}
	}, func() action.Action {
		return &ModifyTunnelLogicalGroupAction{}
	}, func() action.Action {
		return &CreateEnvAction{}
	}, func() action.Action {
		return &UpdateEnvAction{}
	}, func() action.Action {
		return &DeleteEnvAction{}
	}, func() action.Action {
		return &ReloadInventoryAction{}
	}, func() action.Action {
		return &CreateDeployAction{}
	}, func() action.Action {
		return &UpdateDeployAction{}
	}, func() action.Action {
		return &DeleteDeployAction{}
	}, func() action.Action {
		return &ApplyGvTapPrecryptionPolicyConfigAction{}
	}, func() action.Action {
		return &ClearGvTapPrecryptionPolicyConfigurationAction{}
	}, func() action.Action {
		return &ApplyGvTapPrefilteringPolicyConfigAction{}
	}, func() action.Action {
		return &ClearGvTapTrafficPolicyConfigurationAction{}
	}, func() action.Action {
		return &ApplyPrecryptionPolicyConfigAction{}
	}, func() action.Action {
		return &ClearPrecryptionPolicyConfigurationAction{}
	}, func() action.Action {
		return &ApplyPrefilteringPolicyConfigAction{}
	}, func() action.Action {
		return &ClearTrafficPolicyConfigurationAction{}
	}, func() action.Action {
		return &EditUserDetailsAction{}
	}, func() action.Action {
		return &DeleteUserAction{}
	}, func() action.Action {
		return &FmUnlockUserConfigAction{}
	}, func() action.Action {
		return &LoadAliasesMetaDataAction{}
	}, func() action.Action {
		return &LoadSrcPortMetaDataAction{}
	}}
}

// Functions returns the provider-defined functions registered with this provider.
func (p *gigavuecoreProvider) Functions(_ context.Context) []func() function.Function {
	return []func() function.Function{func() function.Function {
		return &QueryRawDataFunction{SourceOperation: "Query_Raw_Data"}
	}, func() function.Function {
		return &LoadResultsBasedOnScrollIdFunction{SourceOperation: "Load_results_based_on_ScrollId"}
	}, func() function.Function {
		return &LoadAllDefaultSearchDomainsFunction{SourceOperation: "loadAllDefaultSearchDomains"}
	}}
}

// EphemeralResources returns the ephemeral resources registered with this provider.
func (p *gigavuecoreProvider) EphemeralResources(_ context.Context) []func() ephemeral.EphemeralResource {
	return []func() ephemeral.EphemeralResource{func() ephemeral.EphemeralResource {
		return &CreateApiTokenTokensCreateEphemeralResource{}
	}, func() ephemeral.EphemeralResource {
		return &RegenerateSecretEphemeralResource{}
	}}
}

// ListResources returns the list resources registered with this provider.
func (p *gigavuecoreProvider) ListResources(_ context.Context) []func() list.ListResource {
	return nil
}
