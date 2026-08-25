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
	_ resource.Resource                = (*ExporterResource)(nil)
	_ resource.ResourceWithImportState = (*ExporterResource)(nil)
)

// ExporterResource is the generated Terraform managed resource implementation.
type ExporterResource struct {
}

// ExporterResourceModel describes the Terraform state and plan shape for ExporterResource.
type ExporterResourceModel struct {
	Alias             types.String `tfsdk:"alias"`
	Description       types.String `tfsdk:"description"`
	Destination       types.Object `tfsdk:"destination"`
	GsGroupAssociated types.List   `tfsdk:"gs_group_associated" json:"gsGroupAssociated"`
	Source            types.Object `tfsdk:"source"`
	SslProfile        types.String `tfsdk:"ssl_profile" json:"sslProfile"`
	Status            types.String `tfsdk:"status"`
	Tags              types.List   `tfsdk:"tags"`
	TcpProfile        types.String `tfsdk:"tcp_profile" json:"tcpProfile"`
	Type              types.String `tfsdk:"type"`
}

// Metadata returns the resource type name.
func (r *ExporterResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_exporter"
}

// Schema returns the Terraform schema for this resource.
func (r *ExporterResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "new in H 5.8", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "Alias of the exporter", Required: true}, "description": schema.StringAttribute{MarkdownDescription: "Comments if necessary", Optional: true, Computed: true}, "destination": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"l3": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"ip": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"dscp": schema.Int64Attribute{MarkdownDescription: "DSCP Value to use", Optional: true, Computed: true}, "ttl": schema.Int64Attribute{MarkdownDescription: "TTL Value to use", Optional: true, Computed: true}, "ver4": schema.StringAttribute{MarkdownDescription: "IPv4 address. If not specified, it will be determined by the App based on other traffic correlation or external means (auto mode).", Optional: true, Computed: true}, "ver6": schema.StringAttribute{MarkdownDescription: "IPv6 address. If not specified, it will be determined by the App based on other traffic correlation or external means (auto mode).", Optional: true, Computed: true}}}, "protocol": schema.StringAttribute{MarkdownDescription: "Protocol used (when it's auto, it's determined by the App based on context or by discovery)", Optional: true, Computed: true}}}, "l4": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"port": schema.Int64Attribute{MarkdownDescription: "Base port used to export, port is optional for type:gtp-cups", Optional: true, Computed: true}, "protocol": schema.StringAttribute{MarkdownDescription: "Protocol used - TCP or UDP", Required: true}}}}}, "gs_group_associated": schema.ListAttribute{Optional: true, Computed: true, ElementType: types.StringType}, "source": schema.SingleNestedAttribute{Optional: true, Computed: true, Attributes: map[string]schema.Attribute{"interface": schema.StringAttribute{MarkdownDescription: "Alias of IP Interface", Required: true}, "l4_port": schema.Int64Attribute{MarkdownDescription: "Base source port number to use for outgoing connections", Required: true}}}, "ssl_profile": schema.StringAttribute{MarkdownDescription: "SSL profile alias", Optional: true, Computed: true}, "status": schema.StringAttribute{Optional: true, Computed: true}, "tags": schema.ListNestedAttribute{Optional: true, Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{"tag_key": schema.StringAttribute{MarkdownDescription: "Name of the tag", Required: true}, "tag_values": schema.ListAttribute{MarkdownDescription: "All possible values of the tag", Required: true, ElementType: types.StringType}}}}, "tcp_profile": schema.StringAttribute{MarkdownDescription: "TCP profile alias", Optional: true, Computed: true}, "type": schema.StringAttribute{MarkdownDescription: "Type of Apps that export", Optional: true, Computed: true}}}
}

// Create provisions the remote resource and stores the resulting state.
func (r *ExporterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ExporterResourceModel
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
func (r *ExporterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ExporterResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Read is not wired to a remote API endpoint.")
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update modifies the remote resource to match the desired plan.
func (r *ExporterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ExporterResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	var state ExporterResourceModel
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
func (r *ExporterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ExporterResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Delete is not wired to a remote API endpoint.")
}

// ImportState imports an existing remote resource into Terraform state.
func (r *ExporterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("alias"), req.ID)...)
}
