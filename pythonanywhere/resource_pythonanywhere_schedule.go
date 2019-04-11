package pythonanywhere

import (
	"github.com/5h4d0wb0y/terraform-provider-pythonanywhere/pythonanywhere/client"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceSchedule() *schema.Resource {
	return &schema.Resource{
		Create: resourceScheduleCreate,
		Read:   resourceScheduleRead,
		Update: resourceScheduleUpdate,
		Delete: resourceScheduleDelete,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Description: "Domain name",
				Required:    true,
				ForceNew:    true,
			},
			"command": {
				Type:        schema.TypeString,
				Description: "Command",
				Required:    true,
				ForceNew:    true,
			},
			"enabled": {
				Type:        schema.TypeBool,
				Description: "Enables",
				Required:    true,
				ForceNew:    true,
			},
			"interval": {
				Type:        schema.TypeString,
				Description: "Interval",
				Required:    true,
				ForceNew:    true,
			},
			"hour": {
				Type:        schema.TypeString,
				Description: "Hour",
				Required:    true,
				ForceNew:    true,
			},
			"minute": {
				Type:        schema.TypeString,
				Description: "Minute",
				Required:    true,
				ForceNew:    true,
			},
		},
	}
}

func resourceScheduleCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pythonanywhere.Client)
	command := d.Get("command").(string)
	enabled := d.Get("enabled").(string)
	interval := d.Get("interval").(string)
	hour := d.Get("hour").(string)
	minute := d.Get("minute").(string)
	client.CreateScheduledTask(command, enabled, interval, hour, minute)

	return resourceFilesRead(d, meta)
}
func resourceScheduleRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pythonanywhere.Client)
	client.ListScheduledTasks()
	return nil
}

func resourceScheduleUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceScheduleRead(d, meta)
}

func resourceScheduleDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pythonanywhere.Client)
	id := d.Get("id").(string)
	client.DeleteScheduledTask(id)
	return nil
}
