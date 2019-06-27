package mrp

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h"
)

func init() {
	h.StockScrap().DeclareModel()

	h.StockScrap().AddFields(map[string]models.FieldDefinition{
		"ProductionId": models.Many2OneField{
			RelationModel: h.MrpProduction(),
			String:        "Manufacturing Order",
			//states={'done': [('readonly', True)]}
		},
		"WorkorderId": models.Many2OneField{
			RelationModel: h.MrpWorkorder(),
			String:        "Work Order",
			//states={'done': [('readonly', True)]}
			Help: "Not to restrict or prefer quants, but informative.",
		},
	})
	h.StockScrap().Methods().OnchangeWorkorderId().DeclareMethod(
		`OnchangeWorkorderId`,
		func(rs m.StockScrapSet) {
			//        if self.workorder_id:
			//            self.location_id = self.workorder_id.production_id.location_src_id.id
		})
	h.StockScrap().Methods().OnchangeProductionId().DeclareMethod(
		`OnchangeProductionId`,
		func(rs m.StockScrapSet) {
			//        if self.production_id:
			//            self.location_id = self.production_id.move_raw_ids.filtered(lambda x: x.state not in (
			//                'done', 'cancel')) and self.production_id.location_src_id.id or self.production_id.location_dest_id.id,
		})
	h.StockScrap().Methods().GetPreferredDomain().DeclareMethod(
		`GetPreferredDomain`,
		func(rs m.StockScrapSet) {
			//        if self.production_id:
			//            if self.product_id in self.production_id.move_raw_ids.mapped('product_id'):
			//                preferred_domain = [
			//                    ('reservation_id', 'in', self.production_id.move_raw_ids.ids)]
			//                preferred_domain2 = [('reservation_id', '=', False)]
			//                preferred_domain3 = [
			//                    '&', ('reservation_id', 'not in', self.production_id.move_raw_ids.ids), ('reservation_id', '!=', False)]
			//                return [preferred_domain, preferred_domain2, preferred_domain3]
			//            elif self.product_id in self.production_id.move_finished_ids.mapped('product_id'):
			//                preferred_domain = [
			//                    ('history_ids', 'in', self.production_id.move_finished_ids.ids)]
			//                preferred_domain2 = [
			//                    ('history_ids', 'not in', self.production_id.move_finished_ids.ids)]
			//                return [preferred_domain, preferred_domain2]
			//        return super(StockScrap, self)._get_preferred_domain()
		})
	h.StockScrap().Methods().PrepareMoveValues().DeclareMethod(
		`PrepareMoveValues`,
		func(rs m.StockScrapSet) {
			//        vals = super(StockScrap, self)._prepare_move_values()
			//        if self.production_id:
			//            vals['origin'] = vals['origin'] or self.production_id.name
			//            if self.product_id in self.production_id.move_finished_ids.mapped('product_id'):
			//                vals.update({'production_id': self.production_id.id})
			//            else:
			//                vals.update(
			//                    {'raw_material_production_id': self.production_id.id})
			//        return vals
		})
	h.StockScrap().Methods().GetOriginMoves().DeclareMethod(
		`GetOriginMoves`,
		func(rs m.StockScrapSet) {
			//        return super(StockScrap, self)._get_origin_moves() or self.production_id and self.production_id.move_raw_ids.filtered(lambda x: x.product_id == self.product_id)
		})
}
