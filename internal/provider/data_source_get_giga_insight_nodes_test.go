package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetGigaInsightNodesDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetGigaInsightNodesDataSourceSchemaValidation(t *testing.T) {
	d := NewGetGigaInsightNodesDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetGigaInsightNodesDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetGigaInsightNodesDataSourceMetadata(t *testing.T) {
	d := NewGetGigaInsightNodesDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_giga_insight_nodes" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_giga_insight_nodes")
	}
}
