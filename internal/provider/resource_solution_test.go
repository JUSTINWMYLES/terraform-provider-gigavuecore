package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestSolutionResourceSchemaValidation verifies that the generated resource schema is valid.
func TestSolutionResourceSchemaValidation(t *testing.T) {
	r := &SolutionResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestSolutionResourceMetadata verifies that the generated resource reports the expected type name.
func TestSolutionResourceMetadata(t *testing.T) {
	r := &SolutionResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_solution" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_solution")
	}
}
