package pythonanywhere

import (
	"github.com/5h4d0wb0y/terraform-provider-pythonanywhere/pythonanywhere/client"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFiles() *schema.Resource {
	return &schema.Resource{
		Create: resourceFilesCreate,
		Read:   resourceFilesRead,
		Update: resourceFilesUpdate,
		Delete: resourceFilesDelete,

		Schema: map[string]*schema.Schema{
			"path": {
				Type:        schema.TypeString,
				Description: "Path",
				Required:    true,
				ForceNew:    true,
			},
		},
	}
}

func resourceFilesCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pythonanywhere.Client)
	path := d.Get("path").(string)
	client.ShareFile(path)

	return resourceFilesRead(d, meta)
}

func resourceFilesRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pythonanywhere.Client)
	path := d.Get("path").(string)
	client.ListFolder(path)
	return nil
}

func resourceFilesUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceFilesRead(d, meta)
}

func resourceFilesDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pythonanywhere.Client)
	path := d.Get("path").(string)
	client.DeleteFile(path)
	return nil
}
