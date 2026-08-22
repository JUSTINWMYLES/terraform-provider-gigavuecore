package provider

import "context"
import (
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertions.
var _ resource.Resource = (*TcpProfileResource)(nil)

// TcpProfileResource is the generated Terraform managed resource implementation.
type TcpProfileResource struct {
}

// TcpProfileResourceModel describes the Terraform state and plan shape for TcpProfileResource.
type TcpProfileResourceModel struct {
	Alias          types.String `tfsdk:"alias"`
	Id             types.String `tfsdk:"id"`
	KeepAliveTimer types.Int64  `tfsdk:"keep_alive_timer" json:"keepAliveTimer"`
	SelectiveAck   types.String `tfsdk:"selective_ack" json:"selectiveAck"`
	SynRetries     types.Int64  `tfsdk:"syn_retries" json:"synRetries"`
}

// Metadata returns the resource type name.
func (r *TcpProfileResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_tcp_profile"
}

// Schema returns the Terraform schema for this resource.
func (r *TcpProfileResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get Apps TCP Profile for given alias", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{Required: true}, "id": schema.StringAttribute{Computed: true}, "keep_alive_timer": schema.Int64Attribute{Optional: true, Computed: true}, "selective_ack": schema.StringAttribute{Optional: true, Computed: true}, "syn_retries": schema.Int64Attribute{Optional: true, Computed: true}}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *TcpProfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan TcpProfileResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Create is not wired to a remote API endpoint.")
	if plan.Id.IsNull() || plan.Id.IsUnknown() {
		plan.Id = types.StringValue("generated")
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// Read refreshes the Terraform state with the latest remote values.
func (r *TcpProfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state TcpProfileResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update modifies the remote resource to match the desired plan.
func (r *TcpProfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan TcpProfileResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state TcpProfileResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if plan.Id.IsNull() || plan.Id.IsUnknown() {
		if !state.Id.IsNull() && !state.Id.IsUnknown() {
			plan.Id = state.Id
		}
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Update is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// Delete destroys the remote resource.
func (r *TcpProfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state TcpProfileResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Delete is not wired to a remote API endpoint.")
}
