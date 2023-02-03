package google

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGoogleCloudBuildTrigger() *schema.Resource {

	dsSchema := DatasourceSchemaFromResourceSchema(resourceCloudBuildTrigger().Schema)

	AddRequiredFieldsToSchema(dsSchema, "trigger_id", "location")
	AddOptionalFieldsToSchema(dsSchema, "project")

	return &schema.Resource{
		Read:   dataSourceGoogleCloudBuildTriggerRead,
		Schema: dsSchema,
	}

}

func dataSourceGoogleCloudBuildTriggerRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	id, err := ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/triggers/{{trigger_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}

	id = strings.ReplaceAll(id, "/locations/global/", "/")

	d.SetId(id)
	return resourceCloudBuildTriggerRead(d, meta)
}
