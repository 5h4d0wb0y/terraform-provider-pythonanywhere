package pythonanywhere

import (
	"github.com/5h4d0wb0y/terraform-provider-pythonanywhere/pythonanywhere/client"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceConsoles() *schema.Resource {
	return &schema.Resource{
		Create: resourceConsolesCreate,
		Read:   resourceConsolesRead,
		Update: resourceConsolesUpdate,
		Delete: resourceConsolesDelete,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Description: "Console ID",
				Required:    true,
				ForceNew:    true,
			},
			"executable": {
				Type:        schema.TypeString,
				Description: "Executable",
				Required:    true,
				ForceNew:    true,
			},
			"arguments": {
				Type:        schema.TypeString,
				Description: "Arguments",
				Required:    true,
				ForceNew:    true,
			},
			"workingDirectory": {
				Type:        schema.TypeString,
				Description: "Working Directory",
				Required:    true,
				ForceNew:    true,
			},
		},
	}
}

func resourceConsolesCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pythonanywhere.Client)
	executable := d.Get("executable").(string)
	arguments := d.Get("arguments").(string)
	workingDirectory := d.Get("workingDirectory").(string)
	client.CreateConsole(executable, arguments, workingDirectory)

	return resourceFilesRead(d, meta)
}

func resourceConsolesRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pythonanywhere.Client)
	client.ListConsoles()
	return nil
}

func resourceConsolesUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceConsolesRead(d, meta)
}

func resourceConsolesDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pythonanywhere.Client)
	id := d.Get("id").(string)
	client.KillConsole(id)
	return nil
}
