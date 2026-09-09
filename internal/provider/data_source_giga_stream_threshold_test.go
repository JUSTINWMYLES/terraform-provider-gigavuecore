package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGigaStreamThresholdDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGigaStreamThresholdDataSourceSchemaValidation(t *testing.T) {
	d := NewGigaStreamThresholdDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGigaStreamThresholdDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGigaStreamThresholdDataSourceMetadata(t *testing.T) {
	d := NewGigaStreamThresholdDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_giga_stream_threshold" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_giga_stream_threshold")
	}
}
