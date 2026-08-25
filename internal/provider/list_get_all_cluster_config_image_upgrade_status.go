package provider

import "context"
import (
	diag "github.com/hashicorp/terraform-plugin-framework/diag"
	list "github.com/hashicorp/terraform-plugin-framework/list"
	listschema "github.com/hashicorp/terraform-plugin-framework/list/schema"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
)

// Compile-time interface assertion.
var _ list.ListResource = (*GetAllClusterConfigImageUpgradeStatusListResource)(nil)

// GetAllClusterConfigImageUpgradeStatusListResource is the generated Terraform list resource implementation.
type GetAllClusterConfigImageUpgradeStatusListResource struct {
}

// NewGetAllClusterConfigImageUpgradeStatusListResource returns a new instance of the generated list resource.
func NewGetAllClusterConfigImageUpgradeStatusListResource() list.ListResource {
	return &GetAllClusterConfigImageUpgradeStatusListResource{}
}

// Metadata returns the list resource type name.
func (l *GetAllClusterConfigImageUpgradeStatusListResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_all_cluster_config_image_upgrade_status"
}

// ListResourceConfigSchema returns the list resource config schema.
func (l *GetAllClusterConfigImageUpgradeStatusListResource) ListResourceConfigSchema(_ context.Context, _ list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = listschema.Schema{MarkdownDescription: "get all the cluster configuration imageUpgrade status"}
}

// List streams matching resource instances for terraform query.
func (l *GetAllClusterConfigImageUpgradeStatusListResource) List(_ context.Context, _ list.ListRequest, stream *list.ListResultsStream) {
	stream.Results = list.ListResultsStreamDiagnostics(diag.Diagnostics{diag.NewErrorDiagnostic("Generated provider scaffold", "List is not wired to a remote API endpoint.")})
}
