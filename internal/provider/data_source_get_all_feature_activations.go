package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)
import (
	client "github.com/gigavuecore/terraform-provider-gigavuecore/internal/client"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var (
	_ datasource.DataSource              = (*GetAllFeatureActivationsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*GetAllFeatureActivationsDataSource)(nil)
)

// GetAllFeatureActivationsDataSource is the generated Terraform data source implementation.
type GetAllFeatureActivationsDataSource struct {
	client *client.Client
}

// GetAllFeatureActivationsDataSourceModel describes the data source state shape.
type GetAllFeatureActivationsDataSourceModel struct {
	Items types.List   `tfsdk:"items"`
	Sku   types.String `tfsdk:"sku"`
}

// NewGetAllFeatureActivationsDataSource returns a new instance of the generated data source.
func NewGetAllFeatureActivationsDataSource() datasource.DataSource {
	return &GetAllFeatureActivationsDataSource{}
}

// Metadata returns the data source type name.
func (d *GetAllFeatureActivationsDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_all_feature_activations"
}

// Schema returns the data source schema.
func (d *GetAllFeatureActivationsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get all the pooled license feature activations (floating and VBL), including the bindings associated with each floating activation", Attributes: map[string]schema.Attribute{"items": schema.ListNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"aid": schema.StringAttribute{MarkdownDescription: "activation id as encoded in the license", Computed: true}, "binding_errors": schema.BoolAttribute{MarkdownDescription: "true if there are heartbeat misses (from FM to device) or any operational error on license bindings to device", Computed: true}, "bindings": schema.ListNestedAttribute{MarkdownDescription: "provides information on the chassis/cards where licenses from this floating pool are assigned to", Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"box_id": schema.Int64Attribute{Computed: true}, "cluster_name": schema.StringAttribute{Computed: true}, "license_key": schema.StringAttribute{MarkdownDescription: "encoded string of a license key that looks like LK2-SMT_HC3-7YF0-86GT-1CEG-LMEU-4KTL-EUPJ-4WHF-L2A4-L3 when decoded", Computed: true}, "slot_id": schema.Int64Attribute{Computed: true}, "target_type": schema.StringAttribute{Computed: true}}}}, "bundle": schema.StringAttribute{MarkdownDescription: "this field is unused in floating licenses", Computed: true}, "description": schema.StringAttribute{MarkdownDescription: "a description of this subset of features", Computed: true}, "eid": schema.StringAttribute{MarkdownDescription: "entitlement id as encoded in the license", Computed: true}, "end_date": schema.StringAttribute{MarkdownDescription: "end date of the license, when it transitions out of ACTIVE state", Computed: true}, "grace_days": schema.Int64Attribute{MarkdownDescription: "number of days of grace period beyond the license end date that allows the functionality provided by the license to be used as if the license was still active", Computed: true}, "license_status": schema.StringAttribute{MarkdownDescription: "status of the license", Computed: true}, "license_type": schema.StringAttribute{MarkdownDescription: "type of the license", Computed: true}, "num_licenses": schema.Int64Attribute{MarkdownDescription: "licenses provided by the SKU is multiplied by this factor", Computed: true}, "registration_date": schema.StringAttribute{MarkdownDescription: "date when the license was imported into FM", Computed: true}, "revocation_code": schema.StringAttribute{MarkdownDescription: "explains why the license was revoked", Computed: true}, "revocation_date": schema.StringAttribute{MarkdownDescription: "date when the license was revoked", Computed: true}, "revoked": schema.BoolAttribute{MarkdownDescription: "was the license revoked", Computed: true}, "sku": schema.StringAttribute{MarkdownDescription: "the license SKU (Stock Keeping Unit) or product code that uniquely identifies the license type in Gigamon catalog", Computed: true}, "start_date": schema.StringAttribute{MarkdownDescription: "start date of the license, when it transitions into ACTIVE state", Computed: true}, "total_volume": schema.StringAttribute{MarkdownDescription: "unused for floating licenses", Computed: true}}}}, "sku": schema.StringAttribute{MarkdownDescription: "Get pooled License activations (floating and VBL) by SKU", Optional: true}}}
}

// Read fetches remote state into the data source model.
func (d *GetAllFeatureActivationsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GetAllFeatureActivationsDataSourceModel
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
func (d *GetAllFeatureActivationsDataSource) readListRemote(ctx context.Context, config *GetAllFeatureActivationsDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/licensing/activations"
	params := url.Values{}
	if !config.Sku.IsNull() {
		params.Set("sku", config.Sku.ValueString())
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
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_feature_activations", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageObj := map[string]any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageObj); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_feature_activations", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		pageItems, ok := pageObj["activations"].([]any)
		if !ok {
			resp.Diagnostics.AddError("Error reading gigavuecore_get_all_feature_activations", fmt.Sprintf("Could not decode list page: missing %q array", "activations"))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_get_all_feature_activations", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *GetAllFeatureActivationsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
