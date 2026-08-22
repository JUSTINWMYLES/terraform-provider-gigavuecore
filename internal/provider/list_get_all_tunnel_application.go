package provider

import "context"
import (
	diag "github.com/hashicorp/terraform-plugin-framework/diag"
	list "github.com/hashicorp/terraform-plugin-framework/list"
	listschema "github.com/hashicorp/terraform-plugin-framework/list/schema"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
)

// Compile-time interface assertion.
var _ list.ListResource = (*GetAllTunnelApplicationListResource)(nil)

// GetAllTunnelApplicationListResource is the generated Terraform list resource implementation.
type GetAllTunnelApplicationListResource struct {
}

// NewGetAllTunnelApplicationListResource returns a new instance of the generated list resource.
func NewGetAllTunnelApplicationListResource() list.ListResource {
	return &GetAllTunnelApplicationListResource{}
}

// Metadata returns the list resource type name.
func (l *GetAllTunnelApplicationListResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_all_tunnel_application"
}

// ListResourceConfigSchema returns the list resource config schema.
func (l *GetAllTunnelApplicationListResource) ListResourceConfigSchema(_ context.Context, _ list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = listschema.Schema{MarkdownDescription: "Get all Tunnel Apps"}
}

// List streams matching resource instances for terraform query.
func (l *GetAllTunnelApplicationListResource) List(_ context.Context, _ list.ListRequest, stream *list.ListResultsStream) {
	stream.Results = list.ListResultsStreamDiagnostics(diag.Diagnostics{diag.NewErrorDiagnostic("Generated provider scaffold", "List is not wired to a remote API endpoint.")})
}
