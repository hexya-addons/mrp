package mrp

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h"
)

func init() {
	h.StockPickingType().DeclareModel()

	h.StockPickingType().AddFields(map[string]models.FieldDefinition{
		"Code": models.SelectionField{
			//selection_add=[('mrp_operation', 'Manufacturing Operation')]
		},
		"CountMoTodo": models.IntegerField{
			Compute: h.StockPickingType().Methods().GetMoCount(),
		},
		"CountMoWaiting": models.IntegerField{
			Compute: h.StockPickingType().Methods().GetMoCount(),
		},
		"CountMoLate": models.IntegerField{
			Compute: h.StockPickingType().Methods().GetMoCount(),
		},
	})
	h.StockPickingType().Methods().GetMoCount().DeclareMethod(
		`GetMoCount`,
		func(rs h.StockPickingTypeSet) h.StockPickingTypeData {
			//        mrp_picking_types = self.filtered(
			//            lambda picking: picking.code == 'mrp_operation')
			//        if not mrp_picking_types:
			//            return
			//        domains = {
			//            'count_mo_waiting': [('availability', '=', 'waiting')],
			//            'count_mo_todo': [('state', 'in', ('confirmed', 'planned', 'progress'))],
			//            'count_mo_late': [('date_planned_start', '<', fields.Date.today()), ('state', '=', 'confirmed')],
			//        }
			//        for field in domains:
			//            data = self.env['mrp.production'].read_group(domains[field] +
			//                                                         [('state', 'not in', ('done', 'cancel')),
			//                                                          ('picking_type_id', 'in', self.ids)],
			//                                                         ['picking_type_id'], ['picking_type_id'])
			//            count = dict(map(lambda x: (
			//                x['picking_type_id'] and x['picking_type_id'][0], x['picking_type_id_count']), data))
			//            for record in mrp_picking_types:
			//                record[field] = count.get(record.id, 0)
		})
}
