package provider

import "context"
import (
	int64validator "github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	stringvalidator "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	identityschema "github.com/hashicorp/terraform-plugin-framework/resource/identityschema"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	validator "github.com/hashicorp/terraform-plugin-framework/schema/validator"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertions.
var (
	_ resource.Resource             = (*HeaderStripAgingResource)(nil)
	_ resource.ResourceWithIdentity = (*HeaderStripAgingResource)(nil)
)

// HeaderStripAgingResource is the generated Terraform managed resource implementation.
type HeaderStripAgingResource struct {
}

// HeaderStripAgingResourceModel describes the Terraform state and plan shape for HeaderStripAgingResource.
type HeaderStripAgingResourceModel struct {
	AgingInterval types.Int64                    `tfsdk:"aging_interval" json:"agingInterval"`
	BoxId         types.String                   `tfsdk:"box_id" json:"boxId"`
	DstPort       types.Int64                    `tfsdk:"dst_port" json:"dstPort"`
	ProtocolType  types.String                   `tfsdk:"protocol_type" json:"protocolType"`
	Timeouts      *HeaderStripAgingTimeoutsModel `tfsdk:"timeouts"`
}

// HeaderStripAgingTimeoutsModel describes the per-operation timeout configuration (in seconds) for the gigavuecore_header_strip_aging resource.
type HeaderStripAgingTimeoutsModel struct {
	Create types.Int64 `tfsdk:"create"`
	Read   types.Int64 `tfsdk:"read"`
	Update types.Int64 `tfsdk:"update"`
	Delete types.Int64 `tfsdk:"delete"`
}

// Metadata returns the resource type name.
func (r *HeaderStripAgingResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_header_strip_aging"
}

// Schema returns the Terraform schema for this resource.
func (r *HeaderStripAgingResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Configure headerstrip aging", Attributes: map[string]schema.Attribute{"aging_interval": schema.Int64Attribute{MarkdownDescription: "Interval in sec. Valid range 300-1000000. Enter 0 to disable.", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(300, 1000000)}}, "box_id": schema.StringAttribute{MarkdownDescription: "device box id. valid range 1 - 64. all is applicable only for post request.", Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}}, "dst_port": schema.Int64Attribute{MarkdownDescription: "L4 destination port number.Valid value is between 0 to 65535.", Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(0, 65535)}}, "protocol_type": schema.StringAttribute{MarkdownDescription: "protocol type", Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("none", "vxlan")}}}, Blocks: map[string]schema.Block{"timeouts": schema.SingleNestedBlock{Attributes: map[string]schema.Attribute{"create": schema.Int64Attribute{MarkdownDescription: "Maximum time in seconds for the create operation. Defaults to 1200.", Optional: true}, "read": schema.Int64Attribute{MarkdownDescription: "Maximum time in seconds for the read operation. Defaults to 1200.", Optional: true}, "update": schema.Int64Attribute{MarkdownDescription: "Maximum time in seconds for the update operation. Defaults to 1200.", Optional: true}, "delete": schema.Int64Attribute{MarkdownDescription: "Maximum time in seconds for the delete operation. Defaults to 1200.", Optional: true}}, MarkdownDescription: "Per-operation timeouts in seconds."}}}
}

// IdentitySchema returns the resource identity schema shared with the paired list resource.
func (r *HeaderStripAgingResource) IdentitySchema(_ context.Context, _ resource.IdentitySchemaRequest, resp *resource.IdentitySchemaResponse) {
	resp.IdentitySchema = identityschema.Schema{Attributes: map[string]identityschema.Attribute{"box_id": identityschema.StringAttribute{RequiredForImport: true, Description: "device box id. valid range 1 - 64. all is applicable only for post request."}}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *HeaderStripAgingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan HeaderStripAgingResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Create is not wired to a remote API endpoint.")
	if plan.BoxId.IsNull() || plan.BoxId.IsUnknown() {
		plan.BoxId = types.StringValue("generated")
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// Read refreshes the Terraform state with the latest remote values.
func (r *HeaderStripAgingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state HeaderStripAgingResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update modifies the remote resource to match the desired plan.
func (r *HeaderStripAgingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan HeaderStripAgingResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state HeaderStripAgingResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if plan.BoxId.IsNull() || plan.BoxId.IsUnknown() {
		if !state.BoxId.IsNull() && !state.BoxId.IsUnknown() {
			plan.BoxId = state.BoxId
		}
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Update is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// Delete destroys the remote resource.
func (r *HeaderStripAgingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state HeaderStripAgingResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Delete is not wired to a remote API endpoint.")
}
