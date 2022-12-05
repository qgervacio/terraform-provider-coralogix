package coralogix

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"terraform-provider-coralogix/coralogix/clientset"
	enrichmentv1 "terraform-provider-coralogix/coralogix/clientset/grpc/com/coralogix/enrichment/v1"
)

func dataSourceCoralogixDataSet() *schema.Resource {
	DataSetSchema := datasourceSchemaFromResourceSchema(DataSetSchema())
	DataSetSchema["id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}

	return &schema.Resource{
		ReadContext: dataSourceCoralogixDataSetRead,

		Schema: DataSetSchema,
	}
}

func dataSourceCoralogixDataSetRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Get("id").(string)
	req := &enrichmentv1.GetCustomEnrichmentRequest{Id: wrapperspb.UInt32(strToUint32(id))}
	log.Printf("[INFO] Reading custom-enrichment-data %s", id)
	enrichmentResp, err := meta.(*clientset.ClientSet).DataSet().GetDataSet(ctx, req)
	if err != nil {
		log.Printf("[ERROR] Received error: %#v", err)
		return handleRpcErrorWithID(err, "custom-enrichment-data", id)
	}
	log.Printf("[INFO] Received custom-enrichment-data: %#v", enrichmentResp)

	d.SetId(uint32ToStr(enrichmentResp.GetCustomEnrichment().GetId()))

	return setDataSet(d, enrichmentResp.GetCustomEnrichment())
}
