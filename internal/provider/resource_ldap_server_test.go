package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestLdapServerResourceSchemaValidation verifies that the generated resource schema is valid.
func TestLdapServerResourceSchemaValidation(t *testing.T) {
	r := &LdapServerResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLdapServerResourceMetadata verifies that the generated resource reports the expected type name.
func TestLdapServerResourceMetadata(t *testing.T) {
	r := &LdapServerResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_ldap_server" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_ldap_server")
	}
}
