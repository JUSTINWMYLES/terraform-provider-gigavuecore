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
	_ datasource.DataSource              = (*LatestImportTagResultDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*LatestImportTagResultDataSource)(nil)
)

// LatestImportTagResultDataSource is the generated Terraform data source implementation.
type LatestImportTagResultDataSource struct {
	client *client.Client
}

// LatestImportTagResultDataSourceModel describes the data source state shape.
type LatestImportTagResultDataSourceModel struct {
	Completed     types.Int64  `tfsdk:"completed"`
	EndTime       types.Int64  `tfsdk:"end_time" json:"endTime"`
	ErrorsCount   types.Int64  `tfsdk:"errors_count" json:"errorsCount"`
	FileType      types.String `tfsdk:"file_type" json:"fileType"`
	OperationType types.String `tfsdk:"operation_type" json:"operationType"`
	SkippedCount  types.Int64  `tfsdk:"skipped_count" json:"skippedCount"`
	StartTime     types.Int64  `tfsdk:"start_time" json:"startTime"`
	Status        types.String `tfsdk:"status"`
	Total         types.Int64  `tfsdk:"total"`
}

// NewLatestImportTagResultDataSource returns a new instance of the generated data source.
func NewLatestImportTagResultDataSource() datasource.DataSource {
	return &LatestImportTagResultDataSource{}
}

// Metadata returns the data source type name.
func (d *LatestImportTagResultDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_latest_import_tag_result"
}

// Schema returns the data source schema.
func (d *LatestImportTagResultDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "get the latest import result all tags", Attributes: map[string]schema.Attribute{"completed": schema.Int64Attribute{MarkdownDescription: "number of operations completed", Computed: true}, "end_time": schema.Int64Attribute{MarkdownDescription: "Operation end timestamp in UTC milliseconds", Computed: true}, "errors_count": schema.Int64Attribute{MarkdownDescription: "number of operations failed", Computed: true}, "file_type": schema.StringAttribute{MarkdownDescription: "type of the file imported", Computed: true}, "operation_type": schema.StringAttribute{Computed: true}, "skipped_count": schema.Int64Attribute{MarkdownDescription: "number of operations skipped", Computed: true}, "start_time": schema.Int64Attribute{MarkdownDescription: "Operation start timestamp in UTC milliseconds", Computed: true}, "status": schema.StringAttribute{MarkdownDescription: "operation status", Computed: true}, "total": schema.Int64Attribute{MarkdownDescription: "total number of operations", Computed: true}}}
}

// Read fetches remote state into the data source model.
func (d *LatestImportTagResultDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LatestImportTagResultDataSourceModel
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
func (d *LatestImportTagResultDataSource) readRemote(ctx context.Context, config *LatestImportTagResultDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/tags/import/result/latest"
	httpReq, err := d.client.NewRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_latest_import_tag_result", fmt.Sprintf("Could not build request: %s", err))
		return
	}
	httpResp, err := d.client.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_latest_import_tag_result", fmt.Sprintf("Could not send request: %s", err))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusNotFound {
		resp.Diagnostics.AddError("Error reading gigavuecore_latest_import_tag_result", "The requested resource was not found.")
		return
	}
	if !(httpResp.StatusCode == 200) {
		switch httpResp.StatusCode {
		case 400:
			resp.Diagnostics.AddError("Error reading gigavuecore_latest_import_tag_result", "Invalid request. See errors payload for details")
			return
		case 401:
			resp.Diagnostics.AddError("Error reading gigavuecore_latest_import_tag_result", "Not Authenticated. See errors payload for details")
			return
		case 403:
			resp.Diagnostics.AddError("Error reading gigavuecore_latest_import_tag_result", "Access Denied. See errors payload for details")
			return
		case 404:
			resp.Diagnostics.AddError("Error reading gigavuecore_latest_import_tag_result", "Entity Not Found. See errors payload for details")
			return
		case 409:
			resp.Diagnostics.AddError("Error reading gigavuecore_latest_import_tag_result", "Entity Already Exists. See errors payload for details")
			return
		case 500:
			resp.Diagnostics.AddError("Error reading gigavuecore_latest_import_tag_result", "Internal Server Error. See errors payload for details")
			return
		case 503:
			resp.Diagnostics.AddError("Error reading gigavuecore_latest_import_tag_result", "Service Unavailable. See errors payload for details")
			return
		default:
			apiErr, err := client.NewAPIError(httpResp)
			if err != nil {
				resp.Diagnostics.AddError("Error reading gigavuecore_latest_import_tag_result", fmt.Sprintf("Could not read error response: %s", err))
				return
			}
			resp.Diagnostics.AddError("Error reading gigavuecore_latest_import_tag_result", apiErr.Error())
			return
		}
	}
	var data map[string]any
	decoder := json.NewDecoder(httpResp.Body)
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil && err != io.EOF {
		resp.Diagnostics.AddError("Error reading gigavuecore_latest_import_tag_result", fmt.Sprintf("Could not decode response body: %s", err))
		return
	}
	if v, ok := data["bulkTagResult"]; ok {
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
		resp.Diagnostics.AddError("Error reading gigavuecore_latest_import_tag_result", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *LatestImportTagResultDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
