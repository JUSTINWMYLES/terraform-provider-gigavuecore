package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)
import (
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*ResourceSelectionDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*ResourceSelectionDataSource)(nil)
)

// ResourceSelectionDataSource is the generated Terraform data source implementation.
type ResourceSelectionDataSource struct {
	client *client.Client
}

// ResourceSelectionDataSourceModel describes the data source state shape.
type ResourceSelectionDataSourceModel struct {
	Items            types.List   `tfsdk:"items"`
	ParentResourceId types.String `tfsdk:"parent_resource_id" json:"parentResourceId"`
	ResourceType     types.String `tfsdk:"resource_type" json:"resourceType"`
	SkipUsed         types.Bool   `tfsdk:"skip_used" json:"skipUsed"`
}

// NewResourceSelectionDataSource returns a new instance of the generated data source.
func NewResourceSelectionDataSource() datasource.DataSource {
	return &ResourceSelectionDataSource{}
}

// Metadata returns the data source type name.
func (d *ResourceSelectionDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_resource_selection"
}

// Schema returns the data source schema.
func (d *ResourceSelectionDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get all resource details of given resource type that don't have any alert policies associated with.", Attributes: map[string]schema.Attribute{"items": schema.ListNestedAttribute{MarkdownDescription: "List of resources of given resource type", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"parent_resource_id": schema.StringAttribute{MarkdownDescription: "Parent resourceId", Computed: true}, "resource_ids": schema.ListAttribute{MarkdownDescription: "Collection of resource entities", Computed: true, ElementType: types.StringType}}}}, "parent_resource_id": schema.StringAttribute{MarkdownDescription: "Parent resource ID. If the rsource type is port, then providing the clusterId as the parent, fetches the ports of given clusterId.", Optional: true}, "resource_type": schema.StringAttribute{MarkdownDescription: "Type of the resource", Required: true}, "skip_used": schema.BoolAttribute{MarkdownDescription: "Flag to control whether all resources to be returned or only resources that aren't paty of any alert policies. The default value is true.", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *ResourceSelectionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ResourceSelectionDataSourceModel
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
func (d *ResourceSelectionDataSource) readListRemote(ctx context.Context, config *ResourceSelectionDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/fm/alert-policies/resources/{resourceType}"
	reqPath = strings.ReplaceAll(reqPath, "{resourceType}", url.PathEscape(config.ResourceType.ValueString()))
	params := url.Values{}
	if !config.ParentResourceId.IsNull() {
		params.Set("parentResourceId", config.ParentResourceId.ValueString())
	}
	if !config.SkipUsed.IsNull() {
		params.Set("skipUsed", strconv.FormatBool(config.SkipUsed.ValueBool()))
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
		resp.Diagnostics.AddError("Error reading gigavuecore_resource_selection", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_resource_selection", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["resources"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_resource_selection", fmt.Sprintf("Could not decode list page: missing %q array", "resources"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_resource_selection", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *ResourceSelectionDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
