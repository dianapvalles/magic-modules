package google

import (
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGoogleComputeNetworkEndpointGroup() *schema.Resource {
	// Generate datasource schema from resource
	dsSchema := DatasourceSchemaFromResourceSchema(resourceComputeNetworkEndpointGroup().Schema)

	// Set 'Optional' schema elements
	AddOptionalFieldsToSchema(dsSchema, "name")
	AddOptionalFieldsToSchema(dsSchema, "zone")
	AddOptionalFieldsToSchema(dsSchema, "project")
	AddOptionalFieldsToSchema(dsSchema, "self_link")

	return &schema.Resource{
		Read:   dataSourceComputeNetworkEndpointGroupRead,
		Schema: dsSchema,
	}
}

func dataSourceComputeNetworkEndpointGroupRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	if name, ok := d.GetOk("name"); ok {
		project, err := getProject(d, config)
		if err != nil {
			return err
		}
		zone, err := GetZone(d, config)
		if err != nil {
			return err
		}
		d.SetId(fmt.Sprintf("projects/%s/zones/%s/networkEndpointGroups/%s", project, zone, name.(string)))
	} else if selfLink, ok := d.GetOk("self_link"); ok {
		parsed, err := ParseNetworkEndpointGroupFieldValue(selfLink.(string), d, config)
		if err != nil {
			return err
		}
		if err := d.Set("name", parsed.Name); err != nil {
			return fmt.Errorf("Error setting name: %s", err)
		}
		if err := d.Set("zone", parsed.Zone); err != nil {
			return fmt.Errorf("Error setting zone: %s", err)
		}
		if err := d.Set("project", parsed.Project); err != nil {
			return fmt.Errorf("Error setting project: %s", err)
		}
		d.SetId(fmt.Sprintf("projects/%s/zones/%s/networkEndpointGroups/%s", parsed.Project, parsed.Zone, parsed.Name))
	} else {
		return errors.New("Must provide either `self_link` or `zone/name`")
	}

	return resourceComputeNetworkEndpointGroupRead(d, meta)
}
