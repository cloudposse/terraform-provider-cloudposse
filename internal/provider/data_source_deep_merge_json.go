package provider

import (
	"context"
	c "github.com/cloudposse/atmos/pkg/convert"
	m "github.com/cloudposse/atmos/pkg/merge"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	jsoniter "github.com/json-iterator/go"
)

func dataSourceDeepMergeJSON() *schema.Resource {
	return &schema.Resource{
		Description: "The `deep_merge_json` data source accepts a list of JSON strings as input and deep merges into a single JSON string as output.",

		ReadContext: dataSourceDeepMergeJSONRead,

		Schema: map[string]*schema.Schema{
			"input": {
				Description: "A list of JSON strings that is deep merged into the `output` attribute.",
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Required:    true,
			},
			"output": {
				Description: "The deep-merged output.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataSourceDeepMergeJSONRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	input := d.Get("input")

	data, err := c.JSONSliceOfInterfaceToSliceOfMaps(input.([]interface{}))
	if err != nil {
		return diag.FromErr(err)
	}

	merged, err := m.Merge(data)
	if err != nil {
		return diag.FromErr(err)
	}

	// Convert result to JSON
	var json = jsoniter.ConfigDefault
	jsonResult, err := json.Marshal(merged)
	if err != nil {
		return diag.FromErr(err)
	}

	err = d.Set("output", string(jsonResult))
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(c.MakeId(jsonResult))

	return nil
}
