package mrp

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h"
)

func init() {
	h.ProcurementRule().DeclareModel()

	h.ProcurementRule().Methods().GetAction().DeclareMethod(
		`GetAction`,
		func(rs m.ProcurementRuleSet) {
			//        return [('manufacture', _('Manufacture'))] + super(ProcurementRule, self)._get_action()
		})
	h.ProcurementOrder().DeclareModel()

	h.ProcurementOrder().AddFields(map[string]models.FieldDefinition{
		"BomId": models.Many2OneField{
			RelationModel: h.MrpBom(),
			String:        "BoM",
			Index:         true,
			OnDelete:      `cascade`,
		},
		"ProductionId": models.Many2OneField{
			RelationModel: h.MrpProduction(),
			String:        "Manufacturing Order",
		},
	})
	h.ProcurementOrder().Methods().PropagateCancels().DeclareMethod(
		`PropagateCancels`,
		func(rs m.ProcurementOrderSet) {
			//        cancel_man_orders = self.filtered(lambda procurement: procurement.rule_id.action ==
			//                                          'manufacture' and procurement.production_id).mapped('production_id')
			//        if cancel_man_orders:
			//            cancel_man_orders.action_cancel()
			//        return super(ProcurementOrder, self).propagate_cancels()
		})
	h.ProcurementOrder().Methods().Run().DeclareMethod(
		`Run`,
		func(rs m.ProcurementOrderSet) {
			//        self.ensure_one()
			//        if self.rule_id.action == 'manufacture':
			//            # make a manufacturing order for the procurement
			//            return self.make_mo()[self.id]
			//        return super(ProcurementOrder, self)._run()
		})
	h.ProcurementOrder().Methods().Check().DeclareMethod(
		`Check`,
		func(rs m.ProcurementOrderSet) {
			//        return self.production_id.state == 'done' or super(ProcurementOrder, self)._check()
		})
	h.ProcurementOrder().Methods().GetMatchingBom().DeclareMethod(
		` Finds the bill of material for the product from procurement order. `,
		func(rs m.ProcurementOrderSet) {
			//        if self.bom_id:
			//            return self.bom_id
			//        return self.env['mrp.bom'].with_context(
			//            company_id=self.company_id.id, force_company=self.company_id.id
			//        )._bom_find(product=self.product_id, picking_type=self.rule_id.picking_type_id)  # TDE FIXME: context bullshit
		})
	h.ProcurementOrder().Methods().GetDatePlanned().DeclareMethod(
		`GetDatePlanned`,
		func(rs m.ProcurementOrderSet) {
			//        format_date_planned = fields.Datetime.from_string(self.date_planned)
			//        date_planned = format_date_planned - \
			//            relativedelta(days=self.product_id.produce_delay or 0.0)
			//        date_planned = date_planned - \
			//            relativedelta(days=self.company_id.manufacturing_lead)
			//        return date_planned
		})
	h.ProcurementOrder().Methods().PrepareMoVals().DeclareMethod(
		`PrepareMoVals`,
		func(rs m.ProcurementOrderSet, bom interface{}) {
			//        return {
			//            'origin': self.origin,
			//            'product_id': self.product_id.id,
			//            'product_qty': self.product_qty,
			//            'product_uom_id': self.product_uom.id,
			//            'location_src_id': self.rule_id.location_src_id.id or self.location_id.id,
			//            'location_dest_id': self.location_id.id,
			//            'bom_id': bom.id,
			//            'date_planned_start': fields.Datetime.to_string(self._get_date_planned()),
			//            'date_planned_finished': self.date_planned,
			//            'procurement_group_id': self.group_id.id,
			//            'propagate': self.rule_id.propagate,
			//            'picking_type_id': self.rule_id.picking_type_id.id or self.warehouse_id.manu_type_id.id,
			//            'company_id': self.company_id.id,
			//            'procurement_ids': [(6, 0, [self.id])],
			//        }
		})
	h.ProcurementOrder().Methods().MakeMo().DeclareMethod(
		` Create production orders from procurements `,
		func(rs m.ProcurementOrderSet) {
			//        res = {}
			//        Production = self.env['mrp.production']
			//        for procurement in self:
			//            ProductionSudo = Production.sudo().with_context(
			//                force_company=procurement.company_id.id)
			//            bom = procurement._get_matching_bom()
			//            if bom:
			//                # create the MO as SUPERUSER because the current user may not have the rights to do it (mto product launched by a sale for example)
			//                production = ProductionSudo.create(
			//                    procurement._prepare_mo_vals(bom))
			//                res[procurement.id] = production.id
			//                procurement.message_post(
			//                    body=_("Manufacturing Order <em>%s</em> created.") % (production.name))
			//            else:
			//                res[procurement.id] = False
			//                procurement.message_post(
			//                    body=_("No BoM exists for this product!"))
			//        return res
		})
}
