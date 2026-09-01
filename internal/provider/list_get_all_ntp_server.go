package provider

import "context"
import (
	diag "github.com/hashicorp/terraform-plugin-framework/diag"
	list "github.com/hashicorp/terraform-plugin-framework/list"
	listschema "github.com/hashicorp/terraform-plugin-framework/list/schema"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
)

// Compile-time interface assertion.
var _ list.ListResource = (*GetAllNtpServerListResource)(nil)

// GetAllNtpServerListResource is the generated Terraform list resource implementation.
type GetAllNtpServerListResource struct {
}

// NewGetAllNtpServerListResource returns a new instance of the generated list resource.
func NewGetAllNtpServerListResource() list.ListResource {
	return &GetAllNtpServerListResource{}
}

// Metadata returns the list resource type name.
func (l *GetAllNtpServerListResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_all_ntp_server"
}

// ListResourceConfigSchema returns the list resource config schema.
func (l *GetAllNtpServerListResource) ListResourceConfigSchema(_ context.Context, _ list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = listschema.Schema{MarkdownDescription: "Get all NTP Servers in FM", Attributes: map[string]listschema.Attribute{"auth_status": listschema.StringAttribute{MarkdownDescription: "Auth status of NTP", Optional: true}, "fm_ip": listschema.StringAttribute{MarkdownDescription: "FMHighAvailability node IpAddress/domainName", Optional: true}, "page": listschema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned", Optional: true}, "server_host": listschema.StringAttribute{MarkdownDescription: "NTP server/host address", Optional: true}, "server_status": listschema.StringAttribute{MarkdownDescription: "NTP server status", Optional: true}}}
}

// List streams matching resource instances for terraform query.
func (l *GetAllNtpServerListResource) List(_ context.Context, _ list.ListRequest, stream *list.ListResultsStream) {
	stream.Results = list.ListResultsStreamDiagnostics(diag.Diagnostics{diag.NewErrorDiagnostic("Generated provider scaffold", "List is not wired to a remote API endpoint.")})
}
