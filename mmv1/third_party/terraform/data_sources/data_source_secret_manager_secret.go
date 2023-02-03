package google

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSecretManagerSecret() *schema.Resource {

	dsSchema := DatasourceSchemaFromResourceSchema(resourceSecretManagerSecret().Schema)
	AddRequiredFieldsToSchema(dsSchema, "secret_id")
	AddOptionalFieldsToSchema(dsSchema, "project")

	return &schema.Resource{
		Read:   dataSourceSecretManagerSecretRead,
		Schema: dsSchema,
	}
}

func dataSourceSecretManagerSecretRead(d *schema.ResourceData, meta interface{}) error {
	id, err := ReplaceVars(d, meta.(*Config), "projects/{{project}}/secrets/{{secret_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)
	return resourceSecretManagerSecretRead(d, meta)
}
