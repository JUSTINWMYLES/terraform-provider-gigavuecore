package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetGigaPortByPortIdDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetGigaPortByPortIdDataSourceSchemaValidation(t *testing.T) {
	d := NewGetGigaPortByPortIdDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetGigaPortByPortIdDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetGigaPortByPortIdDataSourceMetadata(t *testing.T) {
	d := NewGetGigaPortByPortIdDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_giga_port_by_port_id" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_giga_port_by_port_id")
	}
}
