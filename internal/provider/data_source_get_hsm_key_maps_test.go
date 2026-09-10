package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetHsmKeyMapsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetHsmKeyMapsDataSourceSchemaValidation(t *testing.T) {
	d := NewGetHsmKeyMapsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetHsmKeyMapsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetHsmKeyMapsDataSourceMetadata(t *testing.T) {
	d := NewGetHsmKeyMapsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_hsm_key_maps" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_hsm_key_maps")
	}
}
