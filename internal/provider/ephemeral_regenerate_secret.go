package provider

import "context"
import (
	ephemeral "github.com/hashicorp/terraform-plugin-framework/ephemeral"
	ephemeralschema "github.com/hashicorp/terraform-plugin-framework/ephemeral/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var _ ephemeral.EphemeralResource = (*RegenerateSecretEphemeralResource)(nil)

// RegenerateSecretEphemeralResource is the generated Terraform ephemeral resource implementation.
type RegenerateSecretEphemeralResource struct {
}

// RegenerateSecretEphemeralResourceModel describes the ephemeral resource config and result shape.
type RegenerateSecretEphemeralResourceModel struct {
	Context                types.Object `tfsdk:"context"`
	FmApiTokenUserEntities types.List   `tfsdk:"fm_api_token_user_entities" json:"fmApiTokenUserEntities"`
}

// NewRegenerateSecretEphemeralResource returns a new instance of the generated ephemeral resource.
func NewRegenerateSecretEphemeralResource() ephemeral.EphemeralResource {
	return &RegenerateSecretEphemeralResource{}
}

// Metadata returns the ephemeral resource type name.
func (e *RegenerateSecretEphemeralResource) Metadata(_ context.Context, _ ephemeral.MetadataRequest, resp *ephemeral.MetadataResponse) {
	resp.TypeName = "gigavuecore_regenerate_secret"
}

// Schema returns the ephemeral resource schema.
func (e *RegenerateSecretEphemeralResource) Schema(_ context.Context, _ ephemeral.SchemaRequest, resp *ephemeral.SchemaResponse) {
	resp.Schema = ephemeralschema.Schema{MarkdownDescription: "Revoke other user tokens of FM Users. Users with FM Security Management role with write access can access API.", Attributes: map[string]ephemeralschema.Attribute{"context": ephemeralschema.SingleNestedAttribute{MarkdownDescription: "Gigamon query result context", Optional: true, Attributes: map[string]ephemeralschema.Attribute{"page_no": ephemeralschema.Int64Attribute{MarkdownDescription: "page number of the returned result set", Optional: true}, "page_size": ephemeralschema.Int64Attribute{MarkdownDescription: "page size of the returned result set", Optional: true}, "sort": ephemeralschema.ListAttribute{MarkdownDescription: "sorting info of the returned result set. list of fields in the array indicate sorting order", Optional: true, ElementType: types.StringType}, "total_items": ephemeralschema.Int64Attribute{MarkdownDescription: "total number of items in the queried entity type", Required: true}}}, "fm_api_token_user_entities": ephemeralschema.ListNestedAttribute{Required: true, NestedObject: ephemeralschema.NestedAttributeObject{Attributes: map[string]ephemeralschema.Attribute{"authentication_type": ephemeralschema.StringAttribute{MarkdownDescription: "Authentication Type local/radius/tacacs+/external", Optional: true}, "created_by": ephemeralschema.StringAttribute{MarkdownDescription: "FM user who has created the token", Optional: true}, "created_ts": ephemeralschema.Int64Attribute{MarkdownDescription: "Token creation timestamp", Optional: true}, "expiry_time": ephemeralschema.StringAttribute{MarkdownDescription: "Expiry Time in number of days. Default would be 30 and Maximum of 105 days.", Required: true}, "expiry_ts": ephemeralschema.StringAttribute{MarkdownDescription: "Expiry Timestamp", Optional: true}, "groups": ephemeralschema.ListAttribute{MarkdownDescription: "FM User Groups", Required: true, ElementType: types.StringType}, "token": ephemeralschema.StringAttribute{MarkdownDescription: "FM generated JWT token for FM REST API access", Optional: true, Sensitive: true}, "token_id": ephemeralschema.StringAttribute{MarkdownDescription: "Random alpha-numeric 64 digit string", Optional: true}, "token_name": ephemeralschema.StringAttribute{MarkdownDescription: "User defined token name", Required: true}, "usage_count": ephemeralschema.Int64Attribute{MarkdownDescription: "Token Tracking count. Usagecount keep track of number times token was used to access API's ", Optional: true}, "username": ephemeralschema.StringAttribute{MarkdownDescription: "FM Username", Required: true}}}}}}
}

// Open generates a new ephemeral resource value.
func (e *RegenerateSecretEphemeralResource) Open(ctx context.Context, req ephemeral.OpenRequest, resp *ephemeral.OpenResponse) {
	var data RegenerateSecretEphemeralResourceModel
	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Open is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.Result.Set(ctx, &data)...)
}
