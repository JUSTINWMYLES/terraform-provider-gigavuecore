package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllExternalTrapReceiverDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAllExternalTrapReceiverDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAllExternalTrapReceiverDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAllExternalTrapReceiverDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAllExternalTrapReceiverDataSourceMetadata(t *testing.T) {
	d := NewGetAllExternalTrapReceiverDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_all_external_trap_receiver" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_all_external_trap_receiver")
	}
}
