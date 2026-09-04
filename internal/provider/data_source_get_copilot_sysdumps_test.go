package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetCopilotSysdumpsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetCopilotSysdumpsDataSourceSchemaValidation(t *testing.T) {
	d := NewGetCopilotSysdumpsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetCopilotSysdumpsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetCopilotSysdumpsDataSourceMetadata(t *testing.T) {
	d := NewGetCopilotSysdumpsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_copilot_sysdumps" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_copilot_sysdumps")
	}
}
