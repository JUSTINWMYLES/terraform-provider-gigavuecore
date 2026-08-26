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
	_ datasource.DataSource              = (*ExpiryNotifAndEmailAndCountDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*ExpiryNotifAndEmailAndCountDataSource)(nil)
)

// ExpiryNotifAndEmailAndCountDataSource is the generated Terraform data source implementation.
type ExpiryNotifAndEmailAndCountDataSource struct {
	client *client.Client
}

// ExpiryNotifAndEmailAndCountDataSourceModel describes the data source state shape.
type ExpiryNotifAndEmailAndCountDataSourceModel struct {
	Items types.List `tfsdk:"items"`
}

// NewExpiryNotifAndEmailAndCountDataSource returns a new instance of the generated data source.
func NewExpiryNotifAndEmailAndCountDataSource() datasource.DataSource {
	return &ExpiryNotifAndEmailAndCountDataSource{}
}

// Metadata returns the data source type name.
func (d *ExpiryNotifAndEmailAndCountDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "gigavuecore_expiry_notif_and_email_and_count"
}

// Schema returns the data source schema.
func (d *ExpiryNotifAndEmailAndCountDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Returns expiry count for floating (90 day, expiring soon, expired) then same for node-locked licenses, after generating events and sending email for those", Attributes: map[string]schema.Attribute{"items": schema.ListAttribute{MarkdownDescription: "first three numbers are for floating licenses, second three are for node-locked licenses; the numbers signify count expiring in 90 days, expiring soon (within notification period), and expired (beyond end date)", Computed: true, ElementType: types.Int64Type}}}
}

// Read fetches remote state into the data source model.
func (d *ExpiryNotifAndEmailAndCountDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ExpiryNotifAndEmailAndCountDataSourceModel
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
func (d *ExpiryNotifAndEmailAndCountDataSource) readListRemote(ctx context.Context, config *ExpiryNotifAndEmailAndCountDataSourceModel, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("Client Not Configured", "The API client was not set on the resource. The provider Configure method must run before resource operations; this is a bug in the generated provider.")
		return
	}
	reqPath := "/licensing/module/notif"
	params := url.Values{}
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
		resp.Diagnostics.AddError("Error reading gigavuecore_expiry_notif_and_email_and_count", fmt.Sprintf("Could not read list response: %s", err))
		return
	}
	items := []any{}
	for _, page := range pages {
		pageItems := []any{}
		dec := json.NewDecoder(bytes.NewReader(page))
		dec.UseNumber()
		if err := dec.Decode(&pageItems); err != nil {
			resp.Diagnostics.AddError("Error reading gigavuecore_expiry_notif_and_email_and_count", fmt.Sprintf("Could not decode list page: %s", err))
			return
		}
		items = append(items, pageItems...)
	}
	if err := applyJSONToModel(&config, map[string]any{"items": items}); err != nil {
		resp.Diagnostics.AddError("Error reading gigavuecore_expiry_notif_and_email_and_count", fmt.Sprintf("Could not map response to state: %s", err))
		return
	}
}

// Configure stores the API client supplied by the provider.
func (d *ExpiryNotifAndEmailAndCountDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
