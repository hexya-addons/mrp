package mrp

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h"
)

func init() {

	h.Company().AddFields(map[string]models.FieldDefinition{
		"ManufacturingLead": models.FloatField{
			String:   "Manufacturing Lead Time",
			Default:  models.DefaultValue(0),
			Required: true,
			Help:     "Security days for each manufacturing operation.",
		},
	})
}
