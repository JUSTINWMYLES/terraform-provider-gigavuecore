package provider

import "context"
import (
	path "github.com/hashicorp/terraform-plugin-framework/path"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertions.
var (
	_ resource.Resource                = (*InterfaceResource)(nil)
	_ resource.ResourceWithImportState = (*InterfaceResource)(nil)
)

// InterfaceResource is the generated Terraform managed resource implementation.
type InterfaceResource struct {
}

// InterfaceResourceModel describes the Terraform state and plan shape for InterfaceResource.
type InterfaceResourceModel struct {
	Alias            types.String `tfsdk:"alias"`
	Attach           types.List   `tfsdk:"attach"`
	Comment          types.String `tfsdk:"comment"`
	Gateway          types.String `tfsdk:"gateway"`
	GsGroups         types.List   `tfsdk:"gs_groups" json:"gsGroups"`
	HwAddress        types.String `tfsdk:"hw_address" json:"hwAddress"`
	IpAddress        types.String `tfsdk:"ip_address" json:"ipAddress"`
	IpMask           types.String `tfsdk:"ip_mask" json:"ipMask"`
	IpType           types.String `tfsdk:"ip_type" json:"ipType"`
	Mtu              types.Int64  `tfsdk:"mtu"`
	NetflowExporters types.List   `tfsdk:"netflow_exporters" json:"netflowExporters"`
	Tags             types.List   `tfsdk:"tags"`
}

// Metadata returns the resource type name.
func (r *InterfaceResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_interface"
}

// Schema returns the Terraform schema for this resource.
func (r *InterfaceResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Load Ip Interface by alias", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "ip interface name", Required: true}, "attach": schema.ListAttribute{MarkdownDescription: "network ports ,tool ports or circuit ports", Optional: true, Computed: true, ElementType: types.StringType}, "comment": schema.StringAttribute{Optional: true, Computed: true}, "gateway": schema.StringAttribute{MarkdownDescription: "gateway ipv4 or ipv6 address", Optional: true, Computed: true}, "gs_groups": schema.ListAttribute{MarkdownDescription: "Gs Groups associated with the IP Interface", Optional: true, Computed: true, ElementType: types.StringType}, "hw_address": schema.StringAttribute{Optional: true, Computed: true}, "ip_address": schema.StringAttribute{MarkdownDescription: "ipv4/ipv6 address", Optional: true, Computed: true}, "ip_mask": schema.StringAttribute{MarkdownDescription: "ipAddress netmask required with ipAddress", Optional: true, Computed: true}, "ip_type": schema.StringAttribute{Optional: true, Computed: true}, "mtu": schema.Int64Attribute{Optional: true, Computed: true}, "netflow_exporters": schema.ListAttribute{MarkdownDescription: "Netflow Exporters associated with the IP Interface", Optional: true, Computed: true, ElementType: types.StringType}, "tags": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"tag_key": schema.StringAttribute{MarkdownDescription: "Name of the tag", Required: true}, "tag_values": schema.ListAttribute{MarkdownDescription: "All possible values of the tag", Required: true, ElementType: types.StringType}}}}}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *InterfaceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Create is not wired to a remote API endpoint.")
	if plan.Alias.IsNull() || plan.Alias.IsUnknown() {
		plan.Alias = types.StringValue("generated")
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// Read refreshes the Terraform state with the latest remote values.
func (r *InterfaceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update modifies the remote resource to match the desired plan.
func (r *InterfaceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan InterfaceResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state InterfaceResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if plan.Alias.IsNull() || plan.Alias.IsUnknown() {
		if !state.Alias.IsNull() && !state.Alias.IsUnknown() {
			plan.Alias = state.Alias
		}
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Update is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// Delete destroys the remote resource.
func (r *InterfaceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfaceResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Delete is not wired to a remote API endpoint.")
}

// ImportState imports an existing remote resource into Terraform state.
func (r *InterfaceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("alias"), req.ID)...)
}
