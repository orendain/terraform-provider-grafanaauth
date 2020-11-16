package provider

import (
	"context"
	"net/url"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	gapi "github.com/orendain/grafana-api-golang-client"
)

func New() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("GRAFANA_URL", nil),
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("GRAFANA_USERNAME", nil),
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("GRAFANA_PASSWORD", nil),
			},
			"token": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("GRAFANA_API_TOKEN", nil),
			},
			"organization_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("GRAFANA_ORGANIZATION_ID", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"grafanaauth_api_key": resourceApiKey(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	cfg := gapi.Config{}

	token, ok := d.GetOk("token")
	if ok {
		cfg.APIKey = token.(string)
	}

	username := d.Get("username").(string)
	password := d.Get("password").(string)
	if (username != "") && (password != "") {
		cfg.BasicAuth = url.UserPassword(username, password)
	}

	org, ok := d.GetOk("organization_id")
	if ok {
		cfg.OrgID = int64(org.(int))
	}

	host := d.Get("url").(string)

	c, err := gapi.New(host, cfg)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return c, diags
}
