package provider

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	gapi "github.com/orendain/grafana-api-golang-client"
)

func resourceApiKey() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceApiKeyCreate,
		ReadContext:   resourceApiKeyRead,
		DeleteContext: resourceApiKeyDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"role": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"seconds_to_live": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"key": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
			"expiration": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceApiKeyCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	name := d.Get("name").(string)
	role := d.Get("role").(string)
	ttl := d.Get("seconds_to_live").(int)

	c := m.(*gapi.Client)
	request := gapi.CreateApiKeyRequest{Name: name, Role: role, SecondsToLive: int64(ttl)}
	response, err := c.CreateApiKey(request)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("key", response.Key)

	// Fill the true resource's state after a create by performing a read.
	resourceApiKeyRead(ctx, d, m)

	return diags
}

func resourceApiKeyRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(*gapi.Client)
	response, err := c.GetApiKeys(true)
	if err != nil {
		return diag.FromErr(err)
	}

	name := d.Get("name").(string)
	for _, key := range response {
		if name == key.Name {
			d.SetId(strconv.FormatInt(key.Id, 10))
			d.Set("name", key.Name)
			d.Set("role", key.Role)

			if !key.Expiration.IsZero() {
				d.Set("expiration", key.Expiration.String())
			}

			return diags
		}
	}

	// Resource was not found via the client. Have Terraform destroy it.
	d.SetId("")

	return diags
}

func resourceApiKeyDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	id, err := strconv.ParseInt(d.Id(), 10, 32)
	if err != nil {
		return diag.FromErr(err)
	}

	c := m.(*gapi.Client)
	_, err = c.DeleteApiKey(id)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
