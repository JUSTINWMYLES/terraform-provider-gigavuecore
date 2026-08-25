package provider

import "context"
import (
	diag "github.com/hashicorp/terraform-plugin-framework/diag"
	list "github.com/hashicorp/terraform-plugin-framework/list"
	listschema "github.com/hashicorp/terraform-plugin-framework/list/schema"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
)

// Compile-time interface assertion.
var _ list.ListResource = (*LoadAllFmNotificationTargetConfigListResource)(nil)

// LoadAllFmNotificationTargetConfigListResource is the generated Terraform list resource implementation.
type LoadAllFmNotificationTargetConfigListResource struct {
}

// NewLoadAllFmNotificationTargetConfigListResource returns a new instance of the generated list resource.
func NewLoadAllFmNotificationTargetConfigListResource() list.ListResource {
	return &LoadAllFmNotificationTargetConfigListResource{}
}

// Metadata returns the list resource type name.
func (l *LoadAllFmNotificationTargetConfigListResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_all_fm_notification_target_config"
}

// ListResourceConfigSchema returns the list resource config schema.
func (l *LoadAllFmNotificationTargetConfigListResource) ListResourceConfigSchema(_ context.Context, _ list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = listschema.Schema{MarkdownDescription: "Load All Fm Notification Target Config"}
}

// List streams matching resource instances for terraform query.
func (l *LoadAllFmNotificationTargetConfigListResource) List(_ context.Context, _ list.ListRequest, stream *list.ListResultsStream) {
	stream.Results = list.ListResultsStreamDiagnostics(diag.Diagnostics{diag.NewErrorDiagnostic("Generated provider scaffold", "List is not wired to a remote API endpoint.")})
}
