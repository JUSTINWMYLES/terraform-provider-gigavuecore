package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadNtpConfigDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadNtpConfigDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadNtpConfigDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadNtpConfigDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadNtpConfigDataSourceMetadata(t *testing.T) {
	d := NewLoadNtpConfigDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_ntp_config" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_ntp_config")
	}
}
