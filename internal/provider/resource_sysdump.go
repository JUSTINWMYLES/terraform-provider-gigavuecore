package provider

import "context"
import (
	timeouts "github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	identityschema "github.com/hashicorp/terraform-plugin-framework/resource/identityschema"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Compile-time interface assertions.
var (
	_ resource.Resource             = (*SysdumpResource)(nil)
	_ resource.ResourceWithIdentity = (*SysdumpResource)(nil)
)

// SysdumpResource is the generated Terraform managed resource implementation.
type SysdumpResource struct {
}

// SysdumpResourceModel describes the Terraform state and plan shape for SysdumpResource.
type SysdumpResourceModel struct {
	Timeouts timeouts.Value `tfsdk:"timeouts"`
}

// Metadata returns the resource type name.
func (r *SysdumpResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_sysdump"
}

// Schema returns the Terraform schema for this resource.
func (r *SysdumpResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Generate New Sysdump File", Blocks: map[string]schema.Block{"timeouts": timeouts.Block(ctx, timeouts.Opts{Create: true, Read: true, Update: true, Delete: true})}}
}

// IdentitySchema returns the resource identity schema shared with the paired list resource.
func (r *SysdumpResource) IdentitySchema(_ context.Context, _ resource.IdentitySchemaRequest, resp *resource.IdentitySchemaResponse) {
	resp.IdentitySchema = identityschema.Schema{Attributes: map[string]identityschema.Attribute{"filename": identityschema.StringAttribute{RequiredForImport: true, Description: "The sysdump filename"}}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *SysdumpResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SysdumpResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Create is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// Read refreshes the Terraform state with the latest remote values.
func (r *SysdumpResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SysdumpResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update modifies the remote resource to match the desired plan.
func (r *SysdumpResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan SysdumpResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state SysdumpResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Update is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// Delete destroys the remote resource.
func (r *SysdumpResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SysdumpResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Delete is not wired to a remote API endpoint.")
}
