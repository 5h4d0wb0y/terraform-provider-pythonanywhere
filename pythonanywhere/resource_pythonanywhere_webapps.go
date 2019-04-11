package pythonanywhere

import (
	"github.com/5h4d0wb0y/terraform-provider-pythonanywhere/pythonanywhere/client"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceWebApps() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebAppsCreate,
		Read:   resourceWebAppsRead,
		Update: resourceWebAppsUpdate,
		Delete: resourceWebAppsDelete,

		Schema: map[string]*schema.Schema{
			"domain_name": {
				Type:        schema.TypeString,
				Description: "Domain name",
				Required:    true,
				ForceNew:    true,
			},
			"python_version": {
				Type:        schema.TypeString,
				Description: "The python version",
				Required:    true,
				ForceNew:    true,
			},
		},
	}
}

func resourceWebAppsCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pythonanywhere.Client)
	domainName := d.Get("domain_name").(string)
	pythonVersion := d.Get("python_version").(string)
	client.CreateWebapp(domainName, pythonVersion)

	return resourceWebAppsRead(d, meta)
}

func resourceWebAppsRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pythonanywhere.Client)
	client.ListWebapps()
	return nil
}

func resourceWebAppsUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceWebAppsRead(d, meta)
}

func resourceWebAppsDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pythonanywhere.Client)
	domainName := d.Get("domain_name").(string)
	client.DeleteWebapp(domainName)
	return nil
}
