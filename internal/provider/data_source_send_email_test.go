package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestSendEmailDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestSendEmailDataSourceSchemaValidation(t *testing.T) {
	d := NewSendEmailDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestSendEmailDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestSendEmailDataSourceMetadata(t *testing.T) {
	d := NewSendEmailDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_send_email" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_send_email")
	}
}
