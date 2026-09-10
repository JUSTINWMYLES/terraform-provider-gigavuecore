package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetSshCiphersDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetSshCiphersDataSourceSchemaValidation(t *testing.T) {
	d := NewGetSshCiphersDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetSshCiphersDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetSshCiphersDataSourceMetadata(t *testing.T) {
	d := NewGetSshCiphersDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_ssh_ciphers" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_ssh_ciphers")
	}
}
