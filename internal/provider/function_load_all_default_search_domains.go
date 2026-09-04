package provider

import "context"
import "github.com/hashicorp/terraform-plugin-framework/function"

var _ function.Function = (*LoadAllDefaultSearchDomainsFunction)(nil)

// LoadAllDefaultSearchDomainsFunction is the generated Terraform provider-defined function implementation.
type LoadAllDefaultSearchDomainsFunction struct {
	SourceOperation string
}

// Metadata returns the function name.
func (f *LoadAllDefaultSearchDomainsFunction) Metadata(_ context.Context, _ function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "load_all_default_search_domains"
}

// Definition returns the function signature.
func (f *LoadAllDefaultSearchDomainsFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{Summary: "Load All Default Search Domains", Description: "Load all default search domains", Return: function.DynamicReturn{}}
}

// Run executes the function logic.
func (f *LoadAllDefaultSearchDomainsFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	resp.Error = function.NewFuncError("Run is not wired to a remote API endpoint.")
}
