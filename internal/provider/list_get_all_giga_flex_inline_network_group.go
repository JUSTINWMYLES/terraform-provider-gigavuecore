package provider

import "context"
import (
	diag "github.com/hashicorp/terraform-plugin-framework/diag"
	list "github.com/hashicorp/terraform-plugin-framework/list"
	listschema "github.com/hashicorp/terraform-plugin-framework/list/schema"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
)

// Compile-time interface assertion.
var _ list.ListResource = (*GetAllGigaFlexInlineNetworkGroupListResource)(nil)

// GetAllGigaFlexInlineNetworkGroupListResource is the generated Terraform list resource implementation.
type GetAllGigaFlexInlineNetworkGroupListResource struct {
}

// NewGetAllGigaFlexInlineNetworkGroupListResource returns a new instance of the generated list resource.
func NewGetAllGigaFlexInlineNetworkGroupListResource() list.ListResource {
	return &GetAllGigaFlexInlineNetworkGroupListResource{}
}

// Metadata returns the list resource type name.
func (l *GetAllGigaFlexInlineNetworkGroupListResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_all_giga_flex_inline_network_group"
}

// ListResourceConfigSchema returns the list resource config schema.
func (l *GetAllGigaFlexInlineNetworkGroupListResource) ListResourceConfigSchema(_ context.Context, _ list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = listschema.Schema{Attributes: map[string]listschema.Attribute{"cluster_id": listschema.StringAttribute{MarkdownDescription: "if provided, Network groups only for that cluster is returned", Optional: true}}}
}

// List streams matching resource instances for terraform query.
func (l *GetAllGigaFlexInlineNetworkGroupListResource) List(_ context.Context, _ list.ListRequest, stream *list.ListResultsStream) {
	stream.Results = list.ListResultsStreamDiagnostics(diag.Diagnostics{diag.NewErrorDiagnostic("Generated provider scaffold", "List is not wired to a remote API endpoint.")})
}
