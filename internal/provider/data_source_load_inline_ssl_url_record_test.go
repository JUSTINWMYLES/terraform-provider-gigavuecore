package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadInlineSslUrlRecordDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadInlineSslUrlRecordDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadInlineSslUrlRecordDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadInlineSslUrlRecordDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadInlineSslUrlRecordDataSourceMetadata(t *testing.T) {
	d := NewLoadInlineSslUrlRecordDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_inline_ssl_url_record" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_inline_ssl_url_record")
	}
}
