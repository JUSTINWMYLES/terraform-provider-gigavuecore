package provider

import "context"
import (
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertions.
var _ resource.Resource = (*FileResource)(nil)

// FileResource is the generated Terraform managed resource implementation.
type FileResource struct {
}

// FileResourceModel describes the Terraform state and plan shape for FileResource.
type FileResourceModel struct {
	ClusterId   types.String       `tfsdk:"cluster_id" json:"clusterId"`
	File        types.String       `tfsdk:"file"`
	Id          types.String       `tfsdk:"id"`
	Result      types.String       `tfsdk:"result"`
	ResultAlias types.String       `tfsdk:"result_alias" json:"resultAlias"`
	Timeouts    *FileTimeoutsModel `tfsdk:"timeouts"`
}

// FileTimeoutsModel describes the per-operation timeout configuration (in seconds) for the gigavuecore_file resource.
type FileTimeoutsModel struct {
	Create types.Int64 `tfsdk:"create"`
	Read   types.Int64 `tfsdk:"read"`
	Update types.Int64 `tfsdk:"update"`
	Delete types.Int64 `tfsdk:"delete"`
}

// Metadata returns the resource type name.
func (r *FileResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_file"
}

// Schema returns the Terraform schema for this resource.
func (r *FileResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Upload a nodecryptlist or decryptlist from local file.", Attributes: map[string]schema.Attribute{"cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "file": schema.StringAttribute{MarkdownDescription: "file to upload to device", Required: true}, "id": schema.StringAttribute{Computed: true}, "result": schema.StringAttribute{MarkdownDescription: "(deprecated: use resultAlias instead)", Computed: true}, "result_alias": schema.StringAttribute{Computed: true}}, Blocks: map[string]schema.Block{"timeouts": schema.SingleNestedBlock{Attributes: map[string]schema.Attribute{"create": schema.Int64Attribute{MarkdownDescription: "Maximum time in seconds for the create operation. Defaults to 1200.", Optional: true}, "read": schema.Int64Attribute{MarkdownDescription: "Maximum time in seconds for the read operation. Defaults to 1200.", Optional: true}, "update": schema.Int64Attribute{MarkdownDescription: "Maximum time in seconds for the update operation. Defaults to 1200.", Optional: true}, "delete": schema.Int64Attribute{MarkdownDescription: "Maximum time in seconds for the delete operation. Defaults to 1200.", Optional: true}}, MarkdownDescription: "Per-operation timeouts in seconds."}}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *FileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan FileResourceModel
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
func (r *FileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state FileResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update modifies the remote resource to match the desired plan.
func (r *FileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan FileResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state FileResourceModel
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
func (r *FileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state FileResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Delete is not wired to a remote API endpoint.")
}
