package provider

import "context"
import (
	diag "github.com/hashicorp/terraform-plugin-framework/diag"
	list "github.com/hashicorp/terraform-plugin-framework/list"
	listschema "github.com/hashicorp/terraform-plugin-framework/list/schema"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
)

// Compile-time interface assertion.
var _ list.ListResource = (*GetAllClusterConfigImageUpgradeStatusByTaskGroupIdListResource)(nil)

// GetAllClusterConfigImageUpgradeStatusByTaskGroupIdListResource is the generated Terraform list resource implementation.
type GetAllClusterConfigImageUpgradeStatusByTaskGroupIdListResource struct {
}

// NewGetAllClusterConfigImageUpgradeStatusByTaskGroupIdListResource returns a new instance of the generated list resource.
func NewGetAllClusterConfigImageUpgradeStatusByTaskGroupIdListResource() list.ListResource {
	return &GetAllClusterConfigImageUpgradeStatusByTaskGroupIdListResource{}
}

// Metadata returns the list resource type name.
func (l *GetAllClusterConfigImageUpgradeStatusByTaskGroupIdListResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_all_cluster_config_image_upgrade_status_by_task_group_id"
}

// ListResourceConfigSchema returns the list resource config schema.
func (l *GetAllClusterConfigImageUpgradeStatusByTaskGroupIdListResource) ListResourceConfigSchema(_ context.Context, _ list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = listschema.Schema{MarkdownDescription: "get all the cluster configuration imageUpgrade status by TaskGroupId"}
}

// List streams matching resource instances for terraform query.
func (l *GetAllClusterConfigImageUpgradeStatusByTaskGroupIdListResource) List(_ context.Context, _ list.ListRequest, stream *list.ListResultsStream) {
	stream.Results = list.ListResultsStreamDiagnostics(diag.Diagnostics{diag.NewErrorDiagnostic("Generated provider scaffold", "List is not wired to a remote API endpoint.")})
}
