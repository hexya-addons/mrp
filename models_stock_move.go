package mrp

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h"
	"github.com/hexya-erp/pool/q"
)

func init() {
	h.StockMoveLots().DeclareModel()

	h.StockMoveLots().AddFields(map[string]models.FieldDefinition{
		"MoveId": models.Many2OneField{
			RelationModel: h.StockMove(),
			String:        "Move",
		},
		"WorkorderId": models.Many2OneField{
			RelationModel: h.MrpWorkorder(),
			String:        "Work Order",
		},
		"ProductionId": models.Many2OneField{
			RelationModel: h.MrpProduction(),
			String:        "Production Order",
		},
		"LotId": models.Many2OneField{
			RelationModel: h.StockProductionLot(),
			String:        "Lot",
			Filter:        q.ProductId().Equals(product_id),
		},
		"LotProducedId": models.Many2OneField{
			RelationModel: h.StockProductionLot(),
			String:        "Finished Lot",
		},
		"LotProducedQty": models.FloatField{
			String: "Quantity Finished Product",
			//digits=dp.get_precision('Product Unit of Measure')
			Help: "Informative, not used in matching",
		},
		"Quantity": models.FloatField{
			String:  "To Do",
			Default: models.DefaultValue(1),
			//digits=dp.get_precision('Product Unit of Measure')
		},
		"QuantityDone": models.FloatField{
			String: "Done",
			//digits=dp.get_precision('Product Unit of Measure')
		},
		"ProductId": models.Many2OneField{
			RelationModel: h.ProductProduct(),
			String:        "Product",
			ReadOnly:      true,
			Related:       `MoveId.ProductId`,
			Stored:        true,
		},
		"DoneWo": models.BooleanField{
			String:  "Done for Work Order",
			Default: models.DefaultValue(true),
			Help:    "Technical Field which is False when temporarily filled in in work order",
		},
		"DoneMove": models.BooleanField{
			String:  "Move Done",
			Related: `MoveId.IsDone`,
			Stored:  true,
		},
		"PlusVisible": models.BooleanField{
			String:  "Plus Visible",
			Compute: h.StockMoveLots().Methods().ComputePlus(),
		},
	})
	h.StockMoveLots().Methods().CheckLotId().DeclareMethod(
		`CheckLotId`,
		func(rs m.StockMoveLotsSet) {
			//        if self.move_id.product_id.tracking == 'serial':
			//            lots = set([])
			//            for move_lot in self.move_id.active_move_lot_ids.filtered(lambda r: not r.lot_produced_id and r.lot_id):
			//                if move_lot.lot_id in lots:
			//                    raise exceptions.UserError(
			//                        _('You cannot use the same serial number in two different lines.'))
			//                if float_compare(move_lot.quantity_done, 1.0, precision_rounding=move_lot.product_id.uom_id.rounding) == 1:
			//                    raise exceptions.UserError(
			//                        _('You can only produce 1.0 %s for products with unique serial number.') % move_lot.product_id.uom_id.name)
			//                lots.add(move_lot.lot_id)
		})
	h.StockMoveLots().Methods().ComputePlus().DeclareMethod(
		`ComputePlus`,
		func(rs h.StockMoveLotsSet) h.StockMoveLotsData {
			//        for movelot in self:
			//            if movelot.move_id.product_id.tracking == 'serial':
			//                movelot.plus_visible = (movelot.quantity_done <= 0.0)
			//            else:
			//                movelot.plus_visible = (movelot.quantity == 0.0) or (
			//                    movelot.quantity_done < movelot.quantity)
		})
	h.StockMoveLots().Methods().DoPlus().DeclareMethod(
		`DoPlus`,
		func(rs m.StockMoveLotsSet) {
			//        self.ensure_one()
			//        self.quantity_done = self.quantity_done + 1
			//        return self.move_id.split_move_lot()
		})
	h.StockMoveLots().Methods().DoMinus().DeclareMethod(
		`DoMinus`,
		func(rs m.StockMoveLotsSet) {
			//        self.ensure_one()
			//        self.quantity_done = self.quantity_done - 1
			//        return self.move_id.split_move_lot()
		})
	h.StockMoveLots().Methods().Write().Extend(
		`Write`,
		func(rs m.StockMoveLotsSet, vals models.RecordData) {
			//        if 'lot_id' in vals:
			//            for movelot in self:
			//                movelot.move_id.production_id.move_raw_ids.mapped('move_lot_ids')\
			//                    .filtered(lambda r: r.done_wo and not r.done_move and r.lot_produced_id == movelot.lot_id)\
			//                    .write({'lot_produced_id': vals['lot_id']})
			//        return super(StockMoveLots, self).write(vals)
		})
	h.StockMove().DeclareModel()

	h.StockMove().AddFields(map[string]models.FieldDefinition{
		"ProductionId": models.Many2OneField{
			RelationModel: h.MrpProduction(),
			String:        "Production Order for finished products",
		},
		"RawMaterialProductionId": models.Many2OneField{
			RelationModel: h.MrpProduction(),
			String:        "Production Order for raw materials",
		},
		"UnbuildId": models.Many2OneField{
			RelationModel: h.MrpUnbuild(),
			String:        "Unbuild Order",
		},
		"ConsumeUnbuildId": models.Many2OneField{
			RelationModel: h.MrpUnbuild(),
			String:        "Consume Unbuild Order",
		},
		"OperationId": models.Many2OneField{
			RelationModel: h.MrpRoutingWorkcenter(),
			String:        "Operation To Consume",
		},
		"WorkorderId": models.Many2OneField{
			RelationModel: h.MrpWorkorder(),
			String:        "Work Order To Consume",
		},
		"HasTracking": models.SelectionField{
			Related: `ProductId.Tracking`,
			String:  "Product with Tracking",
		},
		"QuantityAvailable": models.FloatField{
			String:  "Quantity Available",
			Compute: h.StockMove().Methods().QtyAvailable(),
			//digits=dp.get_precision('Product Unit of Measure')
		},
		"QuantityDoneStore": models.FloatField{
			String: "Quantity done store",
			//digits=0
		},
		"QuantityDone": models.FloatField{
			String:  "Quantity",
			Compute: h.StockMove().Methods().QtyDoneCompute(),
			//inverse='_qty_done_set'
			//digits=dp.get_precision('Product Unit of Measure')
		},
		"MoveLotIds": models.One2ManyField{
			RelationModel: h.StockMoveLots(),
			ReverseFK:     "",
			String:        "Lots",
		},
		"ActiveMoveLotIds": models.One2ManyField{
			RelationModel: h.StockMoveLots(),
			ReverseFK:     "",
			Filter:        q.DoneWo().Equals(True),
			String:        "Lots",
		},
		"BomLineId": models.Many2OneField{
			RelationModel: h.MrpBomLine(),
			String:        "BoM Line",
		},
		"UnitFactor": models.FloatField{
			String: "Unit Factor",
		},
		"IsDone": models.BooleanField{
			String:  "Done",
			Compute: h.StockMove().Methods().ComputeIsDone(),
			Stored:  true,
			Help:    "Technical Field to order moves",
		},
	})
	h.StockMove().Methods().QtyAvailable().DeclareMethod(
		`QtyAvailable`,
		func(rs h.StockMoveSet) h.StockMoveData {
			//        for move in self:
			//            # For consumables, state is available so availability = qty to do
			//            if move.state == 'assigned':
			//                move.quantity_available = move.product_uom_qty
			//            elif move.product_id.uom_id and move.product_uom:
			//                move.quantity_available = move.product_id.uom_id._compute_quantity(
			//                    move.reserved_availability, move.product_uom)
		})
	h.StockMove().Methods().QtyDoneCompute().DeclareMethod(
		`QtyDoneCompute`,
		func(rs h.StockMoveSet) h.StockMoveData {
			//        for move in self:
			//            if move.has_tracking != 'none' or move.sudo().move_lot_ids.mapped('lot_id'):
			//                move.quantity_done = sum(move.move_lot_ids.filtered(lambda x: x.done_wo).mapped(
			//                    'quantity_done'))  # TODO: change with active_move_lot_ids?
			//            else:
			//                move.quantity_done = move.quantity_done_store
		})
	h.StockMove().Methods().QtyDoneSet().DeclareMethod(
		`QtyDoneSet`,
		func(rs m.StockMoveSet) {
			//        for move in self:
			//            if move.has_tracking == 'none':
			//                move.quantity_done_store = move.quantity_done
		})
	h.StockMove().Methods().ComputeIsDone().DeclareMethod(
		`ComputeIsDone`,
		func(rs h.StockMoveSet) h.StockMoveData {
			//        for move in self:
			//            move.is_done = (move.state in ('done', 'cancel'))
		})
	h.StockMove().Methods().ActionAssign().DeclareMethod(
		`ActionAssign`,
		func(rs m.StockMoveSet, no_prepare interface{}) {
			//        res = super(StockMove, self).action_assign(no_prepare=no_prepare)
			//        self.check_move_lots()
			//        return res
		})
	h.StockMove().Methods().PropagateCancel().DeclareMethod(
		`PropagateCancel`,
		func(rs m.StockMoveSet) {
			//        self.ensure_one()
			//        if not self.move_dest_id.raw_material_production_id:
			//            super(StockMove, self)._propagate_cancel()
			//        elif self.move_dest_id.state == 'waiting':
			//            # If waiting, the chain will be broken and we are not sure if we can still wait for it (=> could take from stock instead)
			//            self.move_dest_id.write({'state': 'confirmed'})
		})
	h.StockMove().Methods().ActionCancel().DeclareMethod(
		`ActionCancel`,
		func(rs m.StockMoveSet) {
			//        if any(move.quantity_done for move in self):
			//            raise exceptions.UserError(
			//                _('You cannot cancel a move move having already consumed material'))
			//        return super(StockMove, self).action_cancel()
		})
	h.StockMove().Methods().CheckMoveLots().DeclareMethod(
		`CheckMoveLots`,
		func(rs m.StockMoveSet) {
			//        moves_todo = self.filtered(
			//            lambda x: x.raw_material_production_id and x.state not in ('done', 'cancel'))
			//        return moves_todo.create_lots()
		})
	h.StockMove().Methods().CreateLots().DeclareMethod(
		`CreateLots`,
		func(rs m.StockMoveSet) {
			//        lots = self.env['stock.move.lots']
			//        for move in self:
			//            unlink_move_lots = move.move_lot_ids.filtered(
			//                lambda x: (x.quantity_done == 0) and x.done_wo)
			//            unlink_move_lots.sudo().unlink()
			//            group_new_quant = {}
			//            old_move_lot = {}
			//            for movelot in move.move_lot_ids:
			//                key = (movelot.lot_id.id or False)
			//                old_move_lot.setdefault(key, []).append(movelot)
			//            for quant in move.reserved_quant_ids:
			//                key = (quant.lot_id.id or False)
			//                quantity = move.product_id.uom_id._compute_quantity(
			//                    quant.qty, move.product_uom)
			//                if group_new_quant.get(key):
			//                    group_new_quant[key] += quantity
			//                else:
			//                    group_new_quant[key] = quantity
			//            for key in group_new_quant:
			//                quantity = group_new_quant[key]
			//                if old_move_lot.get(key):
			//                    if old_move_lot[key][0].quantity == quantity:
			//                        continue
			//                    else:
			//                        old_move_lot[key][0].quantity = quantity
			//                else:
			//                    vals = {
			//                        'move_id': move.id,
			//                        'product_id': move.product_id.id,
			//                        'workorder_id': move.workorder_id.id,
			//                        'production_id': move.raw_material_production_id.id,
			//                        'quantity': quantity,
			//                        'lot_id': key,
			//                    }
			//                    lots.create(vals)
			//        return True
		})
	h.StockMove().Methods().CreateExtraMove().DeclareMethod(
		` Creates an extra move if necessary depending on extra
quantities than foreseen or extra moves`,
		func(rs m.StockMoveSet) {
			//        self.ensure_one()
			//        quantity_to_split = 0
			//        uom_qty_to_split = 0
			//        extra_move = self.env['stock.move']
			//        rounding = self.product_uom.rounding
			//        link_procurement = False
			//        if self.procurement_id and self.production_id and float_compare(self.production_id.qty_produced, self.procurement_id.product_qty, precision_rounding=rounding) > 0:
			//            done_moves_total = sum(self.production_id.move_finished_ids.filtered(
			//                lambda x: x.product_id == self.product_id and x.state == 'done').mapped('product_uom_qty'))
			//            # If you depassed the quantity before, you don't need to split anymore, but adapt the quantities
			//            if float_compare(done_moves_total, self.procurement_id.product_qty, precision_rounding=rounding) >= 0:
			//                quantity_to_split = 0
			//                if float_compare(self.product_uom_qty, self.quantity_done, precision_rounding=rounding) < 0:
			//                    # TODO: could change qty on move_dest_id also (in case of 2-step in/out)
			//                    self.product_uom_qty = self.quantity_done
			//            else:
			//                quantity_to_split = done_moves_total + \
			//                    self.quantity_done - self.procurement_id.product_qty
			//                # self.product_uom_qty - (self.procurement_id.product_qty + done_moves_total)
			//                uom_qty_to_split = self.product_uom_qty - \
			//                    (self.quantity_done - quantity_to_split)
			//                if float_compare(uom_qty_to_split, quantity_to_split, precision_rounding=rounding) < 0:
			//                    uom_qty_to_split = quantity_to_split
			//                self.product_uom_qty = self.quantity_done - quantity_to_split
			//        # You split also simply  when the quantity done is bigger than foreseen
			//        elif float_compare(self.quantity_done, self.product_uom_qty, precision_rounding=rounding) > 0:
			//            quantity_to_split = self.quantity_done - self.product_uom_qty
			//            # + no need to change existing self.product_uom_qty
			//            uom_qty_to_split = quantity_to_split
			//            link_procurement = True
			//        if quantity_to_split:
			//            extra_move = self.copy(default={'quantity_done': quantity_to_split, 'product_uom_qty': uom_qty_to_split, 'production_id': self.production_id.id,
			//                                            'raw_material_production_id': self.raw_material_production_id.id,
			//                                            'procurement_id': link_procurement and self.procurement_id.id or False})
			//            extra_move.action_confirm()
			//            if self.has_tracking != 'none':
			//                qty_todo = self.quantity_done - quantity_to_split
			//                for movelot in self.move_lot_ids.filtered(lambda x: x.done_wo):
			//                    if movelot.quantity_done and movelot.done_wo:
			//                        if float_compare(qty_todo, movelot.quantity_done, precision_rounding=rounding) >= 0:
			//                            qty_todo -= movelot.quantity_done
			//                        elif float_compare(qty_todo, 0, precision_rounding=rounding) > 0:
			//                            # split
			//                            remaining = movelot.quantity_done - qty_todo
			//                            movelot.quantity_done = qty_todo
			//                            movelot.copy(
			//                                default={'move_id': extra_move.id, 'quantity_done': remaining})
			//                            qty_todo = 0
			//                        else:
			//                            movelot.move_id = extra_move.id
			//            else:
			//                self.quantity_done -= quantity_to_split
			//        return extra_move
		})
	h.StockMove().Methods().MoveValidate().DeclareMethod(
		` Validate moves based on a production order. `,
		func(rs m.StockMoveSet) {
			//        moves = self._filter_closed_moves()
			//        quant_obj = self.env['stock.quant']
			//        moves_todo = self.env['stock.move']
			//        moves_to_unreserve = self.env['stock.move']
			//        for move in moves:
			//            # Here, the `quantity_done` was already rounded to the product UOM by the `do_produce` wizard. However,
			//            # it is possible that the user changed the value before posting the inventory by a value that should be
			//            # rounded according to the move's UOM. In this specific case, we chose to round up the value, because it
			//            # is what is expected by the user (if i consumed/produced a little more, the whole UOM unit should be
			//            # consumed/produced and the moves are split correctly).
			//            rounding = move.product_uom.rounding
			//            move.quantity_done = float_round(
			//                move.quantity_done, precision_rounding=rounding, rounding_method='UP')
			//            if move.quantity_done <= 0:
			//                continue
			//            moves_todo |= move
			//            moves_todo |= move._create_extra_move()
			//        for move in moves_todo:
			//            rounding = move.product_uom.rounding
			//            if float_compare(move.quantity_done, move.product_uom_qty, precision_rounding=rounding) < 0:
			//                # Need to do some kind of conversion here
			//                qty_split = move.product_uom._compute_quantity(
			//                    move.product_uom_qty - move.quantity_done, move.product_id.uom_id)
			//                new_move = move.split(qty_split)
			//                # If you were already putting stock.move.lots on the next one in the work order, transfer those to the new move
			//                move.move_lot_ids.filtered(
			//                    lambda x: not x.done_wo or x.quantity_done == 0.0).write({'move_id': new_move})
			//                self.browse(new_move).quantity_done = 0.0
			//            main_domain = [('qty', '>', 0)]
			//            preferred_domain = [('reservation_id', '=', move.id)]
			//            fallback_domain = [('reservation_id', '=', False)]
			//            fallback_domain2 = [
			//                '&', ('reservation_id', '!=', move.id), ('reservation_id', '!=', False)]
			//            preferred_domain_list = [preferred_domain] + \
			//                [fallback_domain] + [fallback_domain2]
			//            if move.has_tracking == 'none':
			//                quants = quant_obj.quants_get_preferred_domain(
			//                    move.product_qty, move, domain=main_domain, preferred_domain_list=preferred_domain_list)
			//                self.env['stock.quant'].quants_move(
			//                    quants, move, move.location_dest_id, owner_id=move.restrict_partner_id.id)
			//            else:
			//                for movelot in move.active_move_lot_ids:
			//                    if float_compare(movelot.quantity_done, 0, precision_rounding=rounding) > 0:
			//                        if not movelot.lot_id:
			//                            raise UserError(
			//                                _('You need to supply a lot/serial number.'))
			//                        qty = move.product_uom._compute_quantity(
			//                            movelot.quantity_done, move.product_id.uom_id)
			//                        quants = quant_obj.quants_get_preferred_domain(
			//                            qty, move, lot_id=movelot.lot_id.id, domain=main_domain, preferred_domain_list=preferred_domain_list)
			//                        self.env['stock.quant'].quants_move(
			//                            quants, move, move.location_dest_id, lot_id=movelot.lot_id.id, owner_id=move.restrict_partner_id.id)
			//            moves_to_unreserve |= move
			//            # Next move in production order
			//            if move.move_dest_id and move.move_dest_id.state not in ('done', 'cancel'):
			//                move.move_dest_id.action_assign()
			//        moves_to_unreserve.quants_unreserve()
			//        moves_todo.write({'state': 'done', 'date': fields.Datetime.now()})
			//        return moves_todo
		})
	h.StockMove().Methods().ActionDone().DeclareMethod(
		`ActionDone`,
		func(rs m.StockMoveSet) {
			//        production_moves = self.filtered(lambda move: (
			//            move.production_id or move.raw_material_production_id) and not move.scrapped)
			//        production_moves.move_validate()
			//        return super(StockMove, self-production_moves).action_done()
		})
	h.StockMove().Methods().SplitMoveLot().DeclareMethod(
		`SplitMoveLot`,
		func(rs m.StockMoveSet) {
			//        ctx = dict(self.env.context)
			//        self.ensure_one()
			//        view = self.env.ref('mrp.view_stock_move_lots')
			//        serial = (self.has_tracking == 'serial')
			//        only_create = False  # Check picking type in theory
			//        show_reserved = any([x for x in self.move_lot_ids if x.quantity > 0.0])
			//        ctx.update({
			//            'serial': serial,
			//            'only_create': only_create,
			//            'create_lots': True,
			//            'state_done': self.is_done,
			//            'show_reserved': show_reserved,
			//        })
			//        if ctx.get('w_production'):
			//            action = self.env.ref('mrp.act_mrp_product_produce').read()[0]
			//            action['context'] = ctx
			//            return action
			//        result = {
			//            'name': _('Register Lots'),
			//            'type': 'ir.actions.act_window',
			//            'view_type': 'form',
			//            'view_mode': 'form',
			//            'res_model': 'stock.move',
			//            'views': [(view.id, 'form')],
			//            'view_id': view.id,
			//            'target': 'new',
			//            'res_id': self.id,
			//            'context': ctx,
			//        }
			//        return result
		})
	h.StockMove().Methods().Save().DeclareMethod(
		`Save`,
		func(rs m.StockMoveSet) {
			//        return True
		})
	h.StockMove().Methods().ActionConfirm().DeclareMethod(
		`ActionConfirm`,
		func(rs m.StockMoveSet) {
			//        moves = self.env['stock.move']
			//        for move in self:
			//            moves |= move.action_explode()
			//        return super(StockMove, moves).action_confirm()
		})
	h.StockMove().Methods().ActionExplode().DeclareMethod(
		` Explodes pickings `,
		func(rs m.StockMoveSet) {
			//        if not self.picking_type_id:
			//            return self
			//        bom = self.env['mrp.bom'].sudo()._bom_find(
			//            product=self.product_id, company_id=self.company_id.id)
			//        if not bom or bom.type != 'phantom':
			//            return self
			//        phantom_moves = self.env['stock.move']
			//        processed_moves = self.env['stock.move']
			//        factor = self.product_uom._compute_quantity(
			//            self.product_uom_qty, bom.product_uom_id) / bom.product_qty
			//        boms, lines = bom.sudo().explode(self.product_id, factor,
			//                                         picking_type=bom.picking_type_id)
			//        for bom_line, line_data in lines:
			//            phantom_moves += self._generate_move_phantom(
			//                bom_line, line_data['qty'])
			//        for new_move in phantom_moves:
			//            processed_moves |= new_move.action_explode()
			//        if not self.split_from and self.procurement_id:
			//            # Check if procurements have been made to wait for
			//            moves = self.procurement_id.move_ids
			//            if len(moves) == 1:
			//                self.procurement_id.write({'state': 'done'})
			//        if processed_moves and self.state == 'assigned':
			//            # Set the state of resulting moves according to 'assigned' as the original move is assigned
			//            processed_moves.write({'state': 'assigned'})
			//        self.sudo().unlink()
			//        return processed_moves
		})
	h.StockMove().Methods().PropagateSplit().DeclareMethod(
		`PropagateSplit`,
		func(rs m.StockMoveSet, new_move interface{}, qty interface{}) {
			//        if not self.move_dest_id.raw_material_production_id:
			//            super(StockMove, self)._propagate_split(new_move, qty)
		})
	h.StockMove().Methods().GenerateMovePhantom().DeclareMethod(
		`GenerateMovePhantom`,
		func(rs m.StockMoveSet, bom_line interface{}, quantity interface{}) {
			//        if bom_line.product_id.type in ['product', 'consu']:
			//            return self.copy(default={
			//                'picking_id': self.picking_id.id if self.picking_id else False,
			//                'product_id': bom_line.product_id.id,
			//                'product_uom': bom_line.product_uom_id.id,
			//                'product_uom_qty': quantity,
			//                'state': 'draft',  # will be confirmed below
			//                'name': self.name,
			//                'procurement_id': self.procurement_id.id,
			//                # Needed in order to keep sale connection, but will be removed by unlink
			//                'split_from': self.id,
			//            })
			//        return self.env['stock.move']
		})
	h.StockLocationPath().DeclareModel()

	h.StockLocationPath().Methods().PrepareMoveCopyValues().DeclareMethod(
		`PrepareMoveCopyValues`,
		func(rs m.StockLocationPathSet, move_to_copy interface{}, new_date interface{}) {
			//        new_move_vals = super(PushedFlow, self)._prepare_move_copy_values(
			//            move_to_copy, new_date)
			//        new_move_vals['production_id'] = False
			//        return new_move_vals
		})
}
