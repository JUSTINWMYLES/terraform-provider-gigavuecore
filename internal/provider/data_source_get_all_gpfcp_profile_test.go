package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllGpfcpProfileDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAllGpfcpProfileDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAllGpfcpProfileDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAllGpfcpProfileDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAllGpfcpProfileDataSourceMetadata(t *testing.T) {
	d := NewGetAllGpfcpProfileDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_all_gpfcp_profile" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_all_gpfcp_profile")
	}
}
