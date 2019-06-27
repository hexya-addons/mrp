package mrp

	import (
		"net/http"

		"github.com/hexya-erp/hexya/src/controllers"
		"github.com/hexya-erp/hexya/src/models"
		"github.com/hexya-erp/hexya/src/models/types"
		"github.com/hexya-erp/hexya/src/models/types/dates"
		"github.com/hexya-erp/pool/h"
		"github.com/hexya-erp/pool/q"
	)
	
func init() {
h.MrpUnbuild().DeclareModel()




h.MrpUnbuild().Methods().GetDefaultLocationId().DeclareMethod(
`GetDefaultLocationId`,
func(rs m.MrpUnbuildSet)  {
//        return self.env.ref('stock.stock_location_stock', raise_if_not_found=False)
})
h.MrpUnbuild().Methods().GetDefaultLocationDestId().DeclareMethod(
`GetDefaultLocationDestId`,
func(rs m.MrpUnbuildSet)  {
//        return self.env.ref('stock.stock_location_stock', raise_if_not_found=False)
})
h.MrpUnbuild().AddFields(map[string]models.FieldDefinition{
"Name": models.CharField{
String: "Reference",
NoCopy: true,
ReadOnly: true,
Default: func (env models.Environment) interface{} { return odoo._() },
},
"ProductId": models.Many2OneField{
RelationModel: h.ProductProduct(),
String: "Product",
Required: true,
//states={'done': [('readonly', True)]}
},
"ProductQty": models.FloatField{
String: "Quantity",
Required: true,
//states={'done': [('readonly', True)]}
},
"ProductUomId": models.Many2OneField{
RelationModel: h.ProductUom(),
String: "Unit of Measure",
Required: true,
//states={'done': [('readonly', True)]}
},
"BomId": models.Many2OneField{
RelationModel: h.MrpBom(),
String: "Bill of Material",
Filter: q.ProductTmplId().Equals("product_id.product_tmpl_id"),
Required: true,
//states={'done': [('readonly', True)]}
},
"MoId": models.Many2OneField{
RelationModel: h.MrpProduction(),
String: "Manufacturing Order",
Filter: q.ProductId().Equals(product_id).And().State().In(%!s(<nil>)),
//states={'done': [('readonly', True)]}
},
"LotId": models.Many2OneField{
RelationModel: h.StockProductionLot(),
String: "Lot",
Filter: q.ProductId().Equals(product_id),
//states={'done': [('readonly', True)]}
},
"HasTracking": models.SelectionField{
Related: `ProductId.Tracking`,
ReadOnly: true,
},
"LocationId": models.Many2OneField{
RelationModel: h.StockLocation(),
String: "Location",
Default: models.DefaultValue(_get_default_location_id),
Required: true,
//states={'done': [('readonly', True)]}
},
"LocationDestId": models.Many2OneField{
RelationModel: h.StockLocation(),
String: "Destination Location",
Default: models.DefaultValue(_get_default_location_dest_id),
Required: true,
//states={'done': [('readonly', True)]}
},
"ConsumeLineIds": models.One2ManyField{
RelationModel: h.StockMove(),
ReverseFK: "",
ReadOnly: true,
Help: "",
},
"ProduceLineIds": models.One2ManyField{
RelationModel: h.StockMove(),
ReverseFK: "",
ReadOnly: true,
Help: "",
},
"State": models.SelectionField{
Selection: types.Selection{
"draft": "Draft",
"done": "Done",
},
String: "Status",
Default: models.DefaultValue("draft"),
Index: true,
},
})
h.MrpUnbuild().Methods().OnchangeMoId().DeclareMethod(
`OnchangeMoId`,
func(rs m.MrpUnbuildSet)  {
//        if self.mo_id:
//            self.product_id = self.mo_id.product_id.id
//            self.product_qty = self.mo_id.product_qty
})
h.MrpUnbuild().Methods().OnchangeProductId().DeclareMethod(
`OnchangeProductId`,
func(rs m.MrpUnbuildSet)  {
//        if self.product_id:
//            self.bom_id = self.env['mrp.bom']._bom_find(
//                product=self.product_id)
//            self.product_uom_id = self.product_id.uom_id.id
})
h.MrpUnbuild().Methods().CheckQty().DeclareMethod(
`CheckQty`,
func(rs m.MrpUnbuildSet)  {
//        if self.product_qty <= 0:
//            raise ValueError(
//                _('Unbuild Order product quantity has to be strictly positive.'))
})
h.MrpUnbuild().Methods().Create().Extend(
`Create`,
func(rs m.MrpUnbuildSet, vals models.RecordData)  {
//        if not vals.get('name'):
//            vals['name'] = self.env['ir.sequence'].next_by_code(
//                'mrp.unbuild') or _('New')
//        unbuild = super(MrpUnbuild, self).create(vals)
//        return unbuild
})
h.MrpUnbuild().Methods().ActionUnbuild().DeclareMethod(
`ActionUnbuild`,
func(rs m.MrpUnbuildSet)  {
//        self.ensure_one()
//        if self.product_id.tracking != 'none' and not self.lot_id.id:
//            raise UserError(_('Should have a lot for the finished product'))
//        consume_move = self._generate_consume_moves()[0]
//        produce_moves = self._generate_produce_moves()
//        qty = self.product_qty  # Convert to qty on product UoM
//        if self.mo_id:
//            finished_moves = self.mo_id.move_finished_ids.filtered(
//                lambda move: move.product_id == self.mo_id.product_id)
//            domain = [('qty', '>', 0), ('history_ids',
//                                        'in', finished_moves.ids)]
//        else:
//            domain = [('qty', '>', 0)]
//        quants = self.env['stock.quant'].quants_get_preferred_domain(
//            qty, consume_move,
//            domain=domain,
//            preferred_domain_list=[],
//            lot_id=self.lot_id.id)
//        self.env['stock.quant'].quants_reserve(quants, consume_move)
//        if consume_move.has_tracking != 'none':
//            if not quants[0][0]:
//                raise UserError(
//                    _("You don't have in the stock the lot %s.") % (self.lot_id.name))
//            self.env['stock.move.lots'].create({
//                'move_id': consume_move.id,
//                'lot_id': self.lot_id.id,
//                'quantity_done': consume_move.product_uom_qty,
//                'quantity': consume_move.product_uom_qty})
//        else:
//            consume_move.quantity_done = consume_move.product_uom_qty
//        consume_move.move_validate()
//        original_quants = consume_move.quant_ids.mapped('consumed_quant_ids')
//        for produce_move in produce_moves:
//            if produce_move.has_tracking != 'none':
//                original = original_quants.filtered(
//                    lambda quant: quant.product_id == produce_move.product_id)
//                if not original:
//                    raise UserError(
//                        _("You don't have in the stock the required lot/serial number for %s .") % (produce_move.product_id.name))
//                quantity_todo = produce_move.product_qty
//                for quant in original:
//                    if quantity_todo <= 0:
//                        break
//                    move_quantity = min(quantity_todo, quant.qty)
//                    self.env['stock.move.lots'].create({
//                        'move_id': produce_move.id,
//                        'lot_id': quant.lot_id.id,
//                        'quantity_done': produce_move.product_id.uom_id._compute_quantity(move_quantity, produce_move.product_uom),
//                        'quantity': produce_move.product_id.uom_id._compute_quantity(move_quantity, produce_move.product_uom),
//                    })
//                    quantity_todo -= move_quantity
//            else:
//                produce_move.quantity_done = produce_move.product_uom_qty
//        produce_moves.move_validate()
//        produced_quant_ids = produce_moves.mapped(
//            'quant_ids').filtered(lambda quant: quant.qty > 0)
//        consume_move.quant_ids.sudo().write(
//            {'produced_quant_ids': [(6, 0, produced_quant_ids.ids)]})
//        return self.write({'state': 'done'})
})
h.MrpUnbuild().Methods().GenerateConsumeMoves().DeclareMethod(
`GenerateConsumeMoves`,
func(rs m.MrpUnbuildSet)  {
//        moves = self.env['stock.move']
//        for unbuild in self:
//            move = self.env['stock.move'].create({
//                'name': unbuild.name,
//                'date': unbuild.create_date,
//                'product_id': unbuild.product_id.id,
//                'product_uom': unbuild.product_uom_id.id,
//                'product_uom_qty': unbuild.product_qty,
//                'location_id': unbuild.location_id.id,
//                'location_dest_id': unbuild.product_id.property_stock_production.id,
//                'origin': unbuild.name,
//                'consume_unbuild_id': unbuild.id,
//            })
//            move.action_confirm()
//            moves += move
//        return moves
})
h.MrpUnbuild().Methods().GenerateProduceMoves().DeclareMethod(
`GenerateProduceMoves`,
func(rs m.MrpUnbuildSet)  {
//        moves = self.env['stock.move']
//        for unbuild in self:
//            factor = unbuild.product_uom_id._compute_quantity(
//                unbuild.product_qty, unbuild.bom_id.product_uom_id) / unbuild.bom_id.product_qty
//            boms, lines = unbuild.bom_id.explode(
//                unbuild.product_id, factor, picking_type=unbuild.bom_id.picking_type_id)
//            for line, line_data in lines:
//                moves += unbuild._generate_move_from_bom_line(
//                    line, line_data['qty'])
//        return moves
})
h.MrpUnbuild().Methods().GenerateMoveFromBomLine().DeclareMethod(
`GenerateMoveFromBomLine`,
func(rs m.MrpUnbuildSet, bom_line interface{}, quantity interface{})  {
//        return self.env['stock.move'].create({
//            'name': self.name,
//            'date': self.create_date,
//            'bom_line_id': bom_line.id,
//            'product_id': bom_line.product_id.id,
//            'product_uom_qty': quantity,
//            'product_uom': bom_line.product_uom_id.id,
//            'procure_method': 'make_to_stock',
//            'location_dest_id': self.location_dest_id.id,
//            'location_id': self.product_id.property_stock_production.id,
//            'unbuild_id': self.id,
//        })
})
}