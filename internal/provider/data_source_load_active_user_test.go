package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadActiveUserDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadActiveUserDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadActiveUserDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadActiveUserDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadActiveUserDataSourceMetadata(t *testing.T) {
	d := NewLoadActiveUserDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_active_user" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_active_user")
	}
}
