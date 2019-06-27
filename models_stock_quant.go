package mrp

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h"
)

func init() {
	h.StockQuant().DeclareModel()

	h.StockQuant().AddFields(map[string]models.FieldDefinition{
		"ConsumedQuantIds": models.Many2ManyField{
			RelationModel:    h.StockQuant(),
			M2MLinkModelName: "",
			M2MOurField:      "",
			M2MTheirField:    "",
		},
		"ProducedQuantIds": models.Many2ManyField{
			RelationModel:    h.StockQuant(),
			M2MLinkModelName: "",
			M2MOurField:      "",
			M2MTheirField:    "",
		},
	})
	h.StockQuant().Methods().PrepareHistory().DeclareMethod(
		`PrepareHistory`,
		func(rs m.StockQuantSet) {
			//        vals = super(StockQuant, self)._prepare_history()
			//        vals['consumed_quant_ids'] = [(4, quant.id)
			//                                      for quant in self.consumed_quant_ids]
			//        vals['produced_quant_ids'] = [(4, quant.id)
			//                                      for quant in self.produced_quant_ids]
			//        return vals
		})
}
