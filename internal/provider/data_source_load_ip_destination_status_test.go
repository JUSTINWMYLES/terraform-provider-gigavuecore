package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadIpDestinationStatusDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadIpDestinationStatusDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadIpDestinationStatusDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadIpDestinationStatusDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadIpDestinationStatusDataSourceMetadata(t *testing.T) {
	d := NewLoadIpDestinationStatusDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_ip_destination_status" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_ip_destination_status")
	}
}
