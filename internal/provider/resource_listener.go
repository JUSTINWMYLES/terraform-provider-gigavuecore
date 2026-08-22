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
	_ resource.Resource                = (*ListenerResource)(nil)
	_ resource.ResourceWithImportState = (*ListenerResource)(nil)
)

// ListenerResource is the generated Terraform managed resource implementation.
type ListenerResource struct {
}

// ListenerResourceModel describes the Terraform state and plan shape for ListenerResource.
type ListenerResourceModel struct {
	Alias             types.String `tfsdk:"alias"`
	Description       types.String `tfsdk:"description"`
	GsGroupAssociated types.List   `tfsdk:"gs_group_associated" json:"gsGroupAssociated"`
	IpInterface       types.List   `tfsdk:"ip_interface" json:"ipInterface"`
	L3                types.Object `tfsdk:"l3"`
	L4                types.Object `tfsdk:"l4"`
	Mode              types.String `tfsdk:"mode"`
	SslProfile        types.String `tfsdk:"ssl_profile" json:"sslProfile"`
	Status            types.String `tfsdk:"status"`
	Tags              types.List   `tfsdk:"tags"`
	TcpProfile        types.String `tfsdk:"tcp_profile" json:"tcpProfile"`
	Type              types.String `tfsdk:"type"`
}

// Metadata returns the resource type name.
func (r *ListenerResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_listener"
}

// Schema returns the Terraform schema for this resource.
func (r *ListenerResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Get Apps Listener for given alias for GS as Service", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Alias of the listener", Required: true}, "description": schema.StringAttribute{MarkdownDescription: "Comments if necessary", Optional: true, Computed: true}, "gs_group_associated": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType}, "ip_interface": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType}, "l3": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"dscp": schema.Int64Attribute{MarkdownDescription: "DSCP Value to use", Optional: true, Computed: true}, "protocol": schema.StringAttribute{MarkdownDescription: "Protocol used - ipv4 , ipv6 or both", Optional: true, Computed: true}, "ttl": schema.Int64Attribute{MarkdownDescription: "TTL Value to use", Optional: true, Computed: true}}}, "l4": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"port_list": schema.ListAttribute{MarkdownDescription: "Port lists to listen on", Required: true, ElementType: types.Int64Type}, "protocol": schema.StringAttribute{MarkdownDescription: "Protocol used - TCP or UDP", Required: true}}}, "mode": schema.StringAttribute{MarkdownDescription: "Listen to IP interface or promiscuous mode", Optional: true, Computed: true}, "ssl_profile": schema.StringAttribute{MarkdownDescription: "SSL profile", Optional: true, Computed: true}, "status": schema.StringAttribute{Optional: true, Computed: true}, "tags": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"tag_key": schema.StringAttribute{MarkdownDescription: "Name of the tag", Required: true}, "tag_values": schema.ListAttribute{MarkdownDescription: "All possible values of the tag", Required: true, ElementType: types.StringType}}}}, "tcp_profile": schema.StringAttribute{MarkdownDescription: "TCP profile", Optional: true, Computed: true}, "type": schema.StringAttribute{MarkdownDescription: "Type of Apps that listen", Optional: true, Computed: true}}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *ListenerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ListenerResourceModel
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
func (r *ListenerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ListenerResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update modifies the remote resource to match the desired plan.
func (r *ListenerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ListenerResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state ListenerResourceModel
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
func (r *ListenerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ListenerResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Delete is not wired to a remote API endpoint.")
}

// ImportState imports an existing remote resource into Terraform state.
func (r *ListenerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("alias"), req.ID)...)
}
