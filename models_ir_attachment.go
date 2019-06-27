package mrp

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/types"
	"github.com/hexya-erp/pool/h"
)

func init() {

	h.Attachment().AddFields(map[string]models.FieldDefinition{
		"Priority": models.SelectionField{
			Selection: types.Selection{
				"0": "Normal",
				"1": "Low",
				"2": "High",
				"3": "Very High",
			},
			String: "Priority",
			Help:   "Gives the sequence order when displaying a list of tasks.",
		},
	})
}
