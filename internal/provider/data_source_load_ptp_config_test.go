package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadPtpConfigDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadPtpConfigDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadPtpConfigDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadPtpConfigDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadPtpConfigDataSourceMetadata(t *testing.T) {
	d := NewLoadPtpConfigDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_ptp_config" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_ptp_config")
	}
}
