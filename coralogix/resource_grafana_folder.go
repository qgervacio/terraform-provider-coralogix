package coralogix

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"terraform-provider-coralogix/coralogix/clientset"

	gapi "github.com/grafana/grafana-api-golang-client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func resourceGrafanaFolder() *schema.Resource {
	return &schema.Resource{

		Description: `
* [Official documentation](https://grafana.com/docs/grafana/latest/dashboards/manage-dashboards/)
* [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/folder/)
`,

		CreateContext: CreateFolder,
		DeleteContext: DeleteFolder,
		ReadContext:   ReadFolder,
		UpdateContext: UpdateFolder,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: "Unique identifier.",
			},
			"title": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The title of the folder.",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The full URL of the folder.",
			},
			"prevent_destroy_if_not_empty": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Prevent deletion of the folder if it is not empty (contains dashboards or alert rules).",
			},
		},
	}
}

func CreateFolder(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var err error
	var folder gapi.Folder
	folder.Title = d.Get("title").(string)
	if uid, ok := d.GetOk("uid"); ok {
		folder.UID = uid.(string)
	}

	log.Printf("[INFO] Creating grafana-folder: %#v", folder)
	resp, err := meta.(*clientset.ClientSet).Grafana().CreateGrafanaFolder(ctx, folder)
	if err != nil {
		return diag.Errorf("failed to create folder: %s", err)
	}
	log.Printf("[INFO] Received grafana-folder: %#v", resp)

	flattenGrafanaFolder(*resp, d, meta)

	return ReadFolder(ctx, d, meta)
}

func expandGrafanaFolder(d *schema.ResourceData) gapi.Folder {
	var folder gapi.Folder
	if v, ok := d.GetOk("id"); ok {
		folder.ID = v.(int64)
	}
	if v, ok := d.GetOk("title"); ok {
		folder.Title = v.(string)
	}
	if v, ok := d.GetOk("uid"); ok {
		folder.UID = v.(string)
	}
	if v, ok := d.GetOk("url"); ok {
		folder.URL = v.(string)
	}
	return folder
}

func flattenGrafanaFolder(folder gapi.Folder, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	d.SetId(fmt.Sprintf("%d", folder.ID))

	if err := d.Set("title", folder.Title); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("uid", folder.UID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("url", strings.TrimRight(meta.(*clientset.ClientSet).Grafana().GetTargetURL(), "/")+folder.URL); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func UpdateFolder(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	folder := expandGrafanaFolder(d)
	resp, err := meta.(*clientset.ClientSet).Grafana().UpdateGrafanaFolder(ctx, folder)
	if err != nil {
		log.Printf("[ERROR] Received error: %#v", err)
		return diag.FromErr(err)
	}
	flattenGrafanaFolder(*resp, d, meta)

	return ReadFolder(ctx, d, meta)
}

// SplitOrgResourceID splits into two parts (org ID and resource ID) the ID of an org-scoped resource
func SplitOrgResourceID(id string) (int64, string) {
	if strings.ContainsRune(id, ':') {
		parts := strings.SplitN(id, ":", 2)
		orgID, _ := strconv.ParseInt(parts[0], 10, 64)
		return orgID, parts[1]
	}

	return 0, id
}

func ReadFolder(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	folder, err := meta.(*clientset.ClientSet).Grafana().GetGrafanaFolder(ctx, d.Id())
	if err != nil {
		log.Printf("[ERROR] Received error: %#v", err)
		return diag.FromErr(err)
	}
	log.Printf("[INFO] Received grafana-folder: %#v", folder)

	return flattenGrafanaFolder(*folder, d, meta)
}

func DeleteFolder(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	folder := expandGrafanaFolder(d)
	err := meta.(*clientset.ClientSet).Grafana().DeleteGrafanaFolder(ctx, folder.UID)
	if err != nil {
		log.Printf("[ERROR] Received error: %#v", err)
		if status.Code(err) == codes.NotFound {
			d.SetId("")
			return diag.Diagnostics{diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  fmt.Sprintf("grafana-dashboard %q is in state, but no longer exists in Coralogix backend", d.Id()),
				Detail:   fmt.Sprintf("%s will be recreated when you apply", d.Id()),
			}}
		}
		return diag.Errorf(formatRpcErrors(err, fmt.Sprintf("/grafana/api/folders/%s", folder.UID), fmt.Sprintf("%#v", folder)))
	}

	d.SetId("")
	return nil
}
