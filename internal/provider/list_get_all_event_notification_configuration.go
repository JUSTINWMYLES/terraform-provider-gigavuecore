package provider

import "context"
import (
	diag "github.com/hashicorp/terraform-plugin-framework/diag"
	list "github.com/hashicorp/terraform-plugin-framework/list"
	listschema "github.com/hashicorp/terraform-plugin-framework/list/schema"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
)

// Compile-time interface assertion.
var _ list.ListResource = (*GetAllEventNotificationConfigurationListResource)(nil)

// GetAllEventNotificationConfigurationListResource is the generated Terraform list resource implementation.
type GetAllEventNotificationConfigurationListResource struct {
}

// NewGetAllEventNotificationConfigurationListResource returns a new instance of the generated list resource.
func NewGetAllEventNotificationConfigurationListResource() list.ListResource {
	return &GetAllEventNotificationConfigurationListResource{}
}

// Metadata returns the list resource type name.
func (l *GetAllEventNotificationConfigurationListResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_all_event_notification_configuration"
}

// ListResourceConfigSchema returns the list resource config schema.
func (l *GetAllEventNotificationConfigurationListResource) ListResourceConfigSchema(_ context.Context, _ list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = listschema.Schema{MarkdownDescription: "Get All Event Notification Configurations"}
}

// List streams matching resource instances for terraform query.
func (l *GetAllEventNotificationConfigurationListResource) List(_ context.Context, _ list.ListRequest, stream *list.ListResultsStream) {
	stream.Results = list.ListResultsStreamDiagnostics(diag.Diagnostics{diag.NewErrorDiagnostic("Generated provider scaffold", "List is not wired to a remote API endpoint.")})
}
