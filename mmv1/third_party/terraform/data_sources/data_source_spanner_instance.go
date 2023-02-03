package google

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSpannerInstance() *schema.Resource {

	dsSchema := DatasourceSchemaFromResourceSchema(resourceSpannerInstance().Schema)

	AddRequiredFieldsToSchema(dsSchema, "name")
	AddOptionalFieldsToSchema(dsSchema, "config")       // not sure why this is configurable
	AddOptionalFieldsToSchema(dsSchema, "display_name") // not sure why this is configurable
	AddOptionalFieldsToSchema(dsSchema, "project")

	return &schema.Resource{
		Read:   dataSourceSpannerInstanceRead,
		Schema: dsSchema,
	}
}

func dataSourceSpannerInstanceRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	id, err := ReplaceVars(d, config, "{{project}}/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return resourceSpannerInstanceRead(d, meta)
}
