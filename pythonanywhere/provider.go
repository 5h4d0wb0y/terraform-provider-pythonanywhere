package pythonanywhere

import (
	"github.com/5h4d0wb0y/terraform-provider-pythonanywhere/pythonanywhere/client"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Default:     "",
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("PYTHONANYWHERE_USERNAME", nil),
				Description: "Username for PythonAnywhere",
			},
			"api_token": {
				Type:        schema.TypeString,
				Default:     "",
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("PYTHONANYWHERE_API_TOKEN", nil),
				Description: "The token key for PythonAnywhere",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"webapps":  resourceWebApps(),
			"consoles": resourceConsoles(),
			"files":    resourceFiles(),
			"schedule": resourceSchedule(),
		},
		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	username := d.Get("username").(string)
	token := d.Get("api_token").(string)

	return pythonanywhere.NewClientWith(username, token)
}
