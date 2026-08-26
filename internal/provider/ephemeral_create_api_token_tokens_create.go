package provider

import "context"
import (
	ephemeral "github.com/hashicorp/terraform-plugin-framework/ephemeral"
	ephemeralschema "github.com/hashicorp/terraform-plugin-framework/ephemeral/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion.
var _ ephemeral.EphemeralResource = (*CreateApiTokenTokensCreateEphemeralResource)(nil)

// CreateApiTokenTokensCreateEphemeralResource is the generated Terraform ephemeral resource implementation.
type CreateApiTokenTokensCreateEphemeralResource struct {
}

// CreateApiTokenTokensCreateEphemeralResourceModel describes the ephemeral resource config and result shape.
type CreateApiTokenTokensCreateEphemeralResourceModel struct {
	AuthenticationType types.String  `tfsdk:"authentication_type" json:"authenticationType"`
	CreatedBy          types.String  `tfsdk:"created_by" json:"createdBy"`
	CreatedTs          types.Dynamic `tfsdk:"created_ts" json:"createdTs"`
	ExpiryTime         types.String  `tfsdk:"expiry_time" json:"expiryTime"`
	ExpiryTs           types.String  `tfsdk:"expiry_ts" json:"expiryTs"`
	Groups             types.List    `tfsdk:"groups"`
	Token              types.String  `tfsdk:"token"`
	TokenId            types.String  `tfsdk:"token_id" json:"tokenId"`
	TokenName          types.String  `tfsdk:"token_name" json:"tokenName"`
	UsageCount         types.Dynamic `tfsdk:"usage_count" json:"usageCount"`
	Username           types.String  `tfsdk:"username"`
}

// NewCreateApiTokenTokensCreateEphemeralResource returns a new instance of the generated ephemeral resource.
func NewCreateApiTokenTokensCreateEphemeralResource() ephemeral.EphemeralResource {
	return &CreateApiTokenTokensCreateEphemeralResource{}
}

// Metadata returns the ephemeral resource type name.
func (e *CreateApiTokenTokensCreateEphemeralResource) Metadata(_ context.Context, _ ephemeral.MetadataRequest, resp *ephemeral.MetadataResponse) {
	resp.TypeName = "gigavuecore_create_api_token_tokens_create"
}

// Schema returns the ephemeral resource schema.
func (e *CreateApiTokenTokensCreateEphemeralResource) Schema(_ context.Context, _ ephemeral.SchemaRequest, resp *ephemeral.SchemaResponse) {
	resp.Schema = ephemeralschema.Schema{MarkdownDescription: "Create token for FM Users", Attributes: map[string]ephemeralschema.Attribute{"authentication_type": ephemeralschema.StringAttribute{MarkdownDescription: "Authentication Type local/radius/tacacs+/external", Optional: true, Computed: true}, "created_by": ephemeralschema.StringAttribute{MarkdownDescription: "FM user who has created the token", Optional: true, Computed: true}, "created_ts": ephemeralschema.DynamicAttribute{MarkdownDescription: "Token creation timestamp", Optional: true, Computed: true}, "expiry_time": ephemeralschema.StringAttribute{MarkdownDescription: "Expiry Time in number of days. Default would be 30 and Maximum of 105 days.", Optional: true, Computed: true}, "expiry_ts": ephemeralschema.StringAttribute{MarkdownDescription: "Expiry Timestamp", Optional: true, Computed: true}, "groups": ephemeralschema.ListAttribute{MarkdownDescription: "FM User Groups", Optional: true, Computed: true, ElementType: types.StringType}, "token": ephemeralschema.StringAttribute{MarkdownDescription: "FM generated JWT token for FM REST API access", Optional: true, Computed: true, Sensitive: true}, "token_id": ephemeralschema.StringAttribute{MarkdownDescription: "Random alpha-numeric 64 digit string", Optional: true, Computed: true}, "token_name": ephemeralschema.StringAttribute{MarkdownDescription: "User defined token name", Optional: true, Computed: true}, "usage_count": ephemeralschema.DynamicAttribute{MarkdownDescription: "Token Tracking count. Usagecount keep track of number times token was used to access API's ", Optional: true, Computed: true}, "username": ephemeralschema.StringAttribute{MarkdownDescription: "FM Username", Optional: true, Computed: true}}}
}

// Open generates a new ephemeral resource value.
func (e *CreateApiTokenTokensCreateEphemeralResource) Open(ctx context.Context, req ephemeral.OpenRequest, resp *ephemeral.OpenResponse) {
	var data CreateApiTokenTokensCreateEphemeralResourceModel
	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Open is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.Result.Set(ctx, &data)...)
}
