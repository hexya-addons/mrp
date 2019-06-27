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
	
//import math
func init() {
h.MrpProduction().DeclareModel()
h.MrpProduction().AddSQLConstraint("name_uniq", "unique(name, company_id)", "Reference must be unique per Company!")
h.MrpProduction().AddSQLConstraint("qty_positive", "check (product_qty > 0)", "The quantity to produce must be positive!")





h.MrpProduction().Methods().GetDefaultPickingType().DeclareMethod(
`GetDefaultPickingType`,
func(rs m.MrpProductionSet)  {
//        return self.env['stock.picking.type'].search([
//            ('code', '=', 'mrp_operation'),
//            ('warehouse_id.company_id', 'in', [self.env.context.get('company_id', self.env.user.company_id.id), False])],
//            limit=1).id
})
h.MrpProduction().Methods().GetDefaultLocationSrcId().DeclareMethod(
`GetDefaultLocationSrcId`,
func(rs m.MrpProductionSet)  {
//        location = False
//        if self._context.get('default_picking_type_id'):
//            location = self.env['stock.picking.type'].browse(
//                self.env.context['default_picking_type_id']).default_location_src_id
//        if not location:
//            location = self.env.ref(
//                'stock.stock_location_stock', raise_if_not_found=False)
//        return location and location.id or False
})
h.MrpProduction().Methods().GetDefaultLocationDestId().DeclareMethod(
`GetDefaultLocationDestId`,
func(rs m.MrpProductionSet)  {
//        location = False
//        if self._context.get('default_picking_type_id'):
//            location = self.env['stock.picking.type'].browse(
//                self.env.context['default_picking_type_id']).default_location_dest_id
//        if not location:
//            location = self.env.ref(
//                'stock.stock_location_stock', raise_if_not_found=False)
//        return location and location.id or False
})
h.MrpProduction().AddFields(map[string]models.FieldDefinition{
"Name": models.CharField{
String: "Reference",
NoCopy: true,
ReadOnly: true,
Default: func (env models.Environment) interface{} { return odoo._() },
},
"Origin": models.CharField{
String: "Source",
NoCopy: true,
Help: "Reference of the document that generated this production order request.",
},
"ProductId": models.Many2OneField{
RelationModel: h.ProductProduct(),
String: "Product",
Filter: q.Type().In(%!s(<nil>)),
ReadOnly: true,
Required: true,
//states={'confirmed': [('readonly', False)]}
},
"ProductTmplId": models.Many2OneField{
RelationModel: h.ProductTemplate(),
String: "Product Template",
Related: `ProductId.ProductTmplId`,
ReadOnly: true,
},
"ProductQty": models.FloatField{
String: "Quantity To Produce",
Default: models.DefaultValue(1),
//digits=dp.get_precision('Product Unit of Measure')
ReadOnly: true,
Required: true,
//states={'confirmed': [('readonly', False)]}
},
"ProductUomId": models.Many2OneField{
RelationModel: h.ProductUom(),
String: "Product Unit of Measure",
//oldname='product_uom'
ReadOnly: true,
Required: true,
//states={'confirmed': [('readonly', False)]}
},
"PickingTypeId": models.Many2OneField{
RelationModel: h.StockPickingType(),
String: "Picking Type",
Default: models.DefaultValue(_get_default_picking_type),
Required: true,
},
"LocationSrcId": models.Many2OneField{
RelationModel: h.StockLocation(),
String: "Raw Materials Location",
Default: models.DefaultValue(_get_default_location_src_id),
ReadOnly: true,
Required: true,
//states={'confirmed': [('readonly', False)]}
Help: "Location where the system will look for components.",
},
"LocationDestId": models.Many2OneField{
RelationModel: h.StockLocation(),
String: "Finished Products Location",
Default: models.DefaultValue(_get_default_location_dest_id),
ReadOnly: true,
Required: true,
//states={'confirmed': [('readonly', False)]}
Help: "Location where the system will stock the finished products.",
},
"DatePlannedStart": models.DateTimeField{
String: "Deadline Start",
NoCopy: true,
Default: func (env models.Environment) interface{} { return dates.Now() },
Index: true,
Required: true,
//states={'confirmed': [('readonly', False)]}
//oldname="date_planned"
},
"DatePlannedFinished": models.DateTimeField{
String: "Deadline End",
NoCopy: true,
Default: func (env models.Environment) interface{} { return dates.Now() },
Index: true,
//states={'confirmed': [('readonly', False)]}
},
"DateStart": models.DateTimeField{
String: "Start Date",
NoCopy: true,
Index: true,
ReadOnly: true,
},
"DateFinished": models.DateTimeField{
String: "End Date",
NoCopy: true,
Index: true,
ReadOnly: true,
},
"BomId": models.Many2OneField{
RelationModel: h.MrpBom(),
String: "Bill of Material",
ReadOnly: true,
//states={'confirmed': [('readonly', False)]}
Help: "Bill of Materials allow you to define the list of required" + 
"raw materials to make a finished product.",
},
"RoutingId": models.Many2OneField{
RelationModel: h.MrpRouting(),
String: "Routing",
ReadOnly: true,
Compute: h.MrpProduction().Methods().ComputeRouting(),
Stored: true,
Help: "The list of operations (list of work centers) to produce" + 
"the finished product. The routing is mainly used to compute" + 
"work center costs during operations and to plan future" + 
"loads on work centers based on production planning.",
},
"MoveRawIds": models.One2ManyField{
RelationModel: h.StockMove(),
ReverseFK: "",
String: "Raw Materials",
//oldname='move_lines'
NoCopy: true,
//states={'done': [('readonly', True)], 'cancel': [('readonly', True)]}
Filter: q.Scrapped().Equals(False),
},
"MoveFinishedIds": models.One2ManyField{
RelationModel: h.StockMove(),
ReverseFK: "",
String: "Finished Products",
NoCopy: true,
//states={'done': [('readonly', True)], 'cancel': [('readonly', True)]}
Filter: q.Scrapped().Equals(False),
},
"WorkorderIds": models.One2ManyField{
RelationModel: h.MrpWorkorder(),
ReverseFK: "",
String: "Work Orders",
NoCopy: true,
//oldname='workcenter_lines'
ReadOnly: true,
},
"WorkorderCount": models.IntegerField{
String: "# Work Orders",
Compute: h.MrpProduction().Methods().ComputeWorkorderCount(),
},
"WorkorderDoneCount": models.IntegerField{
String: "# Done Work Orders",
Compute: h.MrpProduction().Methods().ComputeWorkorderDoneCount(),
},
"State": models.SelectionField{
Selection: types.Selection{
"confirmed": "Confirmed",
"planned": "Planned",
"progress": "In Progress",
"done": "Done",
"cancel": "Cancelled",
},
String: "State",
NoCopy: true,
Default: models.DefaultValue("confirmed"),
//track_visibility='onchange'
},
"Availability": models.SelectionField{
Selection: types.Selection{
"assigned": "Available",
"partially_available": "Partially Available",
"waiting": "Waiting",
"none": "None",
},
String: "Availability",
Compute: h.MrpProduction().Methods().ComputeAvailability(),
Stored: true,
},
"UnreserveVisible": models.BooleanField{
String: "Inventory Unreserve Visible",
Compute: h.MrpProduction().Methods().ComputeUnreserveVisible(),
Help: "Technical field to check when we can unreserve",
},
"PostVisible": models.BooleanField{
String: "Inventory Post Visible",
Compute: h.MrpProduction().Methods().ComputePostVisible(),
Help: "Technical field to check when we can post",
},
"UserId": models.Many2OneField{
RelationModel: h.User(),
String: "Responsible",
Default: func (env models.Environment) interface{} { return self._uid },
},
"CompanyId": models.Many2OneField{
RelationModel: h.Company(),
String: "Company",
Default: func (env models.Environment) interface{} { return env["res.company"]._company_default_get() },
Required: true,
},
"CheckToDone": models.BooleanField{
Compute: h.MrpProduction().Methods().GetProducedQty(),
String: "Check Produced Qty",
Help: "Technical Field to see if we can show 'Mark as Done' button",
},
"QtyProduced": models.FloatField{
Compute: h.MrpProduction().Methods().GetProducedQty(),
String: "Quantity Produced",
},
"ProcurementGroupId": models.Many2OneField{
RelationModel: h.ProcurementGroup(),
String: "Procurement Group",
NoCopy: true,
},
"ProcurementIds": models.One2ManyField{
RelationModel: h.ProcurementOrder(),
ReverseFK: "",
String: "Related Procurements",
},
"Propagate": models.BooleanField{
String: "Propagate cancel and split",
Help: "If checked, when the previous move of the move (which was" + 
"generated by a next procurement) is cancelled or split," + 
"the move generated by this move will too",
},
"HasMoves": models.BooleanField{
Compute: h.MrpProduction().Methods().HasMoves(),
},
"ScrapIds": models.One2ManyField{
RelationModel: h.StockScrap(),
ReverseFK: "",
String: "Scraps",
},
"ScrapCount": models.IntegerField{
Compute: h.MrpProduction().Methods().ComputeScrapMoveCount(),
String: "Scrap Move",
},
"Priority": models.SelectionField{
Selection: types.Selection{
"0": "Not urgent",
"1": "Normal",
"2": "Urgent",
"3": "Very Urgent",
},
String: "Priority",
ReadOnly: true,
//states={'confirmed': [('readonly', False)]}
Default: models.DefaultValue("1"),
},
})
h.MrpProduction().Methods().ComputeRouting().DeclareMethod(
`ComputeRouting`,
func(rs h.MrpProductionSet) h.MrpProductionData {
//        for production in self:
//            if production.bom_id.routing_id.operation_ids:
//                production.routing_id = production.bom_id.routing_id.id
//            else:
//                production.routing_id = False
})
h.MrpProduction().Methods().ComputeWorkorderCount().DeclareMethod(
`ComputeWorkorderCount`,
func(rs h.MrpProductionSet) h.MrpProductionData {
//        data = self.env['mrp.workorder'].read_group(
//            [('production_id', 'in', self.ids)], ['production_id'], ['production_id'])
//        count_data = dict(
//            (item['production_id'][0], item['production_id_count']) for item in data)
//        for production in self:
//            production.workorder_count = count_data.get(production.id, 0)
})
h.MrpProduction().Methods().ComputeWorkorderDoneCount().DeclareMethod(
`ComputeWorkorderDoneCount`,
func(rs h.MrpProductionSet) h.MrpProductionData {
//        data = self.env['mrp.workorder'].read_group([
//            ('production_id', 'in', self.ids),
//            ('state', '=', 'done')], ['production_id'], ['production_id'])
//        count_data = dict(
//            (item['production_id'][0], item['production_id_count']) for item in data)
//        for production in self:
//            production.workorder_done_count = count_data.get(production.id, 0)
})
h.MrpProduction().Methods().ComputeAvailability().DeclareMethod(
`ComputeAvailability`,
func(rs h.MrpProductionSet) h.MrpProductionData {
//        for order in self:
//            if not order.move_raw_ids:
//                order.availability = 'none'
//                continue
//            if order.bom_id.ready_to_produce == 'all_available':
//                order.availability = any(move.state not in (
//                    'assigned', 'done', 'cancel') for move in order.move_raw_ids) and 'waiting' or 'assigned'
//            else:
//                partial_list = [x.partially_available and x.state in (
//                    'waiting', 'confirmed', 'assigned') for x in order.move_raw_ids]
//                assigned_list = [x.state in (
//                    'assigned', 'done', 'cancel') for x in order.move_raw_ids]
//                order.availability = (all(assigned_list) and 'assigned') or (
//                    any(partial_list) and 'partially_available') or 'waiting'
})
h.MrpProduction().Methods().ComputeUnreserveVisible().DeclareMethod(
`ComputeUnreserveVisible`,
func(rs h.MrpProductionSet) h.MrpProductionData {
//        for order in self:
//            if order.state in ['done', 'cancel'] or not order.move_raw_ids.mapped('reserved_quant_ids'):
//                order.unreserve_visible = False
//            else:
//                order.unreserve_visible = True
})
h.MrpProduction().Methods().ComputePostVisible().DeclareMethod(
`ComputePostVisible`,
func(rs h.MrpProductionSet) h.MrpProductionData {
//        for order in self:
//            order.post_visible = any(order.move_raw_ids.filtered(lambda x: (x.quantity_done) > 0 and (x.state not in ['done', 'cancel']))) or \
//                any(order.move_finished_ids.filtered(lambda x: (
//                    x.quantity_done) > 0 and (x.state not in ['done', 'cancel'])))
})
h.MrpProduction().Methods().GetProducedQty().DeclareMethod(
`GetProducedQty`,
func(rs h.MrpProductionSet) h.MrpProductionData {
//        for production in self:
//            done_moves = production.move_finished_ids.filtered(
//                lambda x: x.state != 'cancel' and x.product_id.id == production.product_id.id)
//            qty_produced = sum(done_moves.mapped('quantity_done'))
//            wo_done = True
//            if any([x.state not in ('done', 'cancel') for x in production.workorder_ids]):
//                wo_done = False
//            production.check_to_done = done_moves and (qty_produced >= production.product_qty) and (
//                production.state not in ('done', 'cancel')) and wo_done
//            production.qty_produced = qty_produced
//        return True
})
h.MrpProduction().Methods().HasMoves().DeclareMethod(
`HasMoves`,
func(rs h.MrpProductionSet) h.MrpProductionData {
//        for mo in self:
//            mo.has_moves = any(mo.move_raw_ids)
})
h.MrpProduction().Methods().ComputeScrapMoveCount().DeclareMethod(
`ComputeScrapMoveCount`,
func(rs h.MrpProductionSet) h.MrpProductionData {
//        data = self.env['stock.scrap'].read_group([('production_id', 'in', self.ids)], [
//                                                  'production_id'], ['production_id'])
//        count_data = dict(
//            (item['production_id'][0], item['production_id_count']) for item in data)
//        for production in self:
//            production.scrap_count = count_data.get(production.id, 0)
})

h.MrpProduction().Methods().OnchangeProductId().DeclareMethod(
` Finds UoM of changed product. `,
func(rs m.MrpProductionSet)  {
//        if not self.product_id:
//            self.bom_id = False
//        else:
//            bom = self.env['mrp.bom']._bom_find(
//                product=self.product_id, picking_type=self.picking_type_id, company_id=self.company_id.id)
//            if bom.type == 'normal':
//                self.bom_id = bom.id
//            else:
//                self.bom_id = False
//            self.product_uom_id = self.product_id.uom_id.id
//            return {'domain': {'product_uom_id': [('category_id', '=', self.product_id.uom_id.category_id.id)]}}
})
h.MrpProduction().Methods().OnchangePickingType().DeclareMethod(
`OnchangePickingType`,
func(rs m.MrpProductionSet)  {
//        location = self.env.ref('stock.stock_location_stock')
//        self.location_src_id = self.picking_type_id.default_location_src_id.id or location.id
//        self.location_dest_id = self.picking_type_id.default_location_dest_id.id or location.id
})
h.MrpProduction().Methods().Write().Extend(
`Write`,
func(rs m.MrpProductionSet, vals models.RecordData)  {
//        res = super(MrpProduction, self).write(vals)
//        if 'date_planned_start' in vals:
//            moves = (self.mapped('move_raw_ids') + self.mapped('move_finished_ids')).filtered(
//                lambda r: r.state not in ['done', 'cancel'])
//            moves.write({
//                'date_expected': vals['date_planned_start'],
//            })
//        return res
})
h.MrpProduction().Methods().Create().Extend(
`Create`,
func(rs m.MrpProductionSet, values models.RecordData)  {
//        if not values.get('name', False) or values['name'] == _('New'):
//            if values.get('picking_type_id'):
//                values['name'] = self.env['stock.picking.type'].browse(
//                    values['picking_type_id']).sequence_id.next_by_id()
//            else:
//                values['name'] = self.env['ir.sequence'].next_by_code(
//                    'mrp.production') or _('New')
//        if not values.get('procurement_group_id'):
//            values['procurement_group_id'] = self.env["procurement.group"].create(
//                {'name': values['name']}).id
//        production = super(MrpProduction, self).create(values)
//        production._generate_moves()
//        return production
})
h.MrpProduction().Methods().Unlink().Extend(
`Unlink`,
func(rs m.MrpProductionSet)  {
//        if any(production.state != 'cancel' for production in self):
//            raise UserError(
//                _('Cannot delete a manufacturing order not in cancel state'))
//        return super(MrpProduction, self).unlink()
})
h.MrpProduction().Methods().GenerateMoves().DeclareMethod(
`GenerateMoves`,
func(rs m.MrpProductionSet)  {
//        for production in self:
//            production._generate_finished_moves()
//            factor = production.product_uom_id._compute_quantity(
//                production.product_qty, production.bom_id.product_uom_id) / production.bom_id.product_qty
//            boms, lines = production.bom_id.explode(
//                production.product_id, factor, picking_type=production.bom_id.picking_type_id)
//            production._generate_raw_moves(lines)
//            # Check for all draft moves whether they are mto or not
//            production._adjust_procure_method()
//            production.move_raw_ids.action_confirm()
//        return True
})
h.MrpProduction().Methods().GenerateFinishedMoves().DeclareMethod(
`GenerateFinishedMoves`,
func(rs m.MrpProductionSet)  {
//        move = self.env['stock.move'].create({
//            'name': self.name,
//            'date': self.date_planned_start,
//            'date_expected': self.date_planned_start,
//            'product_id': self.product_id.id,
//            'product_uom': self.product_uom_id.id,
//            'product_uom_qty': self.product_qty,
//            'location_id': self.product_id.property_stock_production.id,
//            'location_dest_id': self.location_dest_id.id,
//            'move_dest_id': self.procurement_ids and self.procurement_ids[0].move_dest_id.id or False,
//            'procurement_id': self.procurement_ids and self.procurement_ids[0].id or False,
//            'company_id': self.company_id.id,
//            'production_id': self.id,
//            'origin': self.name,
//            'group_id': self.procurement_group_id.id,
//            'propagate': self.propagate,
//        })
//        move.action_confirm()
//        return move
})
h.MrpProduction().Methods().GenerateRawMoves().DeclareMethod(
`GenerateRawMoves`,
func(rs m.MrpProductionSet, exploded_lines interface{})  {
//        self.ensure_one()
//        moves = self.env['stock.move']
//        for bom_line, line_data in exploded_lines:
//            moves += self._generate_raw_move(bom_line, line_data)
//        return moves
})
h.MrpProduction().Methods().GenerateRawMove().DeclareMethod(
`GenerateRawMove`,
func(rs m.MrpProductionSet, bom_line interface{}, line_data interface{})  {
//        quantity = line_data['qty']
//        alt_op = line_data['parent_line'] and line_data['parent_line'].operation_id.id or False
//        if bom_line.child_bom_id and bom_line.child_bom_id.type == 'phantom':
//            return self.env['stock.move']
//        if bom_line.product_id.type not in ['product', 'consu']:
//            return self.env['stock.move']
//        if self.routing_id:
//            routing = self.routing_id
//        else:
//            routing = self.bom_id.routing_id
//        if routing and routing.location_id:
//            source_location = routing.location_id
//        else:
//            source_location = self.location_src_id
//        original_quantity = (self.product_qty - self.qty_produced) or 1.0
//        data = {
//            'sequence': bom_line.sequence,
//            'name': self.name,
//            'date': self.date_planned_start,
//            'date_expected': self.date_planned_start,
//            'bom_line_id': bom_line.id,
//            'product_id': bom_line.product_id.id,
//            'product_uom_qty': quantity,
//            'product_uom': bom_line.product_uom_id.id,
//            'location_id': source_location.id,
//            'location_dest_id': self.product_id.property_stock_production.id,
//            'raw_material_production_id': self.id,
//            'company_id': self.company_id.id,
//            'operation_id': bom_line.operation_id.id or alt_op,
//            'price_unit': bom_line.product_id.standard_price,
//            'procure_method': 'make_to_stock',
//            'origin': self.name,
//            'warehouse_id': source_location.get_warehouse().id,
//            'group_id': self.procurement_group_id.id,
//            'propagate': self.propagate,
//            'unit_factor': quantity / original_quantity,
//        }
//        return self.env['stock.move'].create(data)
})
h.MrpProduction().Methods().AdjustProcureMethod().DeclareMethod(
`AdjustProcureMethod`,
func(rs m.MrpProductionSet)  {
//        try:
//            mto_route = self.env['stock.warehouse']._get_mto_route()
//        except:
//            mto_route = False
//        for move in self.move_raw_ids:
//            product = move.product_id
//            routes = product.route_ids + product.route_from_categ_ids
//            # TODO: optimize with read_group?
//            pull = self.env['procurement.rule'].search([('route_id', 'in', [x.id for x in routes]), ('location_src_id', '=', move.location_id.id),
//                                                        ('location_id', '=', move.location_dest_id.id)], limit=1)
//            if pull and (pull.procure_method == 'make_to_order'):
//                move.procure_method = pull.procure_method
//            elif not pull:  # If there is no make_to_stock rule either
//                if mto_route and mto_route.id in [x.id for x in routes]:
//                    move.procure_method = 'make_to_order'
})
h.MrpProduction().Methods().UpdateRawMove().DeclareMethod(
`UpdateRawMove`,
func(rs m.MrpProductionSet, bom_line interface{}, line_data interface{})  {
//        quantity = line_data['qty']
//        self.ensure_one()
//        move = self.move_raw_ids.filtered(
//            lambda x: x.bom_line_id.id == bom_line.id and x.state not in ('done', 'cancel'))
//        if move:
//            if quantity > 0:
//                move[0].write({'product_uom_qty': quantity})
//            else:
//                if move[0].quantity_done > 0:
//                    raise UserError(
//                        _('Lines need to be deleted, but can not as you still have some quantities to consume in them. '))
//                move[0].action_cancel()
//                move[0].unlink()
//            return move
//        else:
//            self._generate_raw_move(bom_line, line_data)
})
h.MrpProduction().Methods().ActionAssign().DeclareMethod(
`ActionAssign`,
func(rs m.MrpProductionSet)  {
//        for production in self:
//            move_to_assign = production.move_raw_ids.filtered(
//                lambda x: x.state in ('confirmed', 'waiting', 'assigned'))
//            move_to_assign.action_assign()
//        return True
})
h.MrpProduction().Methods().OpenProduceProduct().DeclareMethod(
`OpenProduceProduct`,
func(rs m.MrpProductionSet)  {
//        self.ensure_one()
//        action = self.env.ref('mrp.act_mrp_product_produce').read()[0]
//        return action
})
h.MrpProduction().Methods().ButtonPlan().DeclareMethod(
` Create work orders. And probably do stuff, like things. `,
func(rs m.MrpProductionSet)  {
//        orders_to_plan = self.filtered(
//            lambda order: order.routing_id and order.state == 'confirmed')
//        for order in orders_to_plan:
//            quantity = order.product_uom_id._compute_quantity(
//                order.product_qty, order.bom_id.product_uom_id) / order.bom_id.product_qty
//            boms, lines = order.bom_id.explode(
//                order.product_id, quantity, picking_type=order.bom_id.picking_type_id)
//            order._generate_workorders(boms)
//        return orders_to_plan.write({'state': 'planned'})
})
h.MrpProduction().Methods().GenerateWorkorders().DeclareMethod(
`GenerateWorkorders`,
func(rs m.MrpProductionSet, exploded_boms interface{})  {
//        workorders = self.env['mrp.workorder']
//        original_one = False
//        for bom, bom_data in exploded_boms:
//            # If the routing of the parent BoM and phantom BoM are the same, don't recreate work orders, but use one master routing
//            if bom.routing_id.id and (not bom_data['parent_line'] or bom_data['parent_line'].bom_id.routing_id.id != bom.routing_id.id):
//                temp_workorders = self._workorders_create(bom, bom_data)
//                workorders += temp_workorders
//                if temp_workorders:  # In order to avoid two "ending work orders"
//                    if original_one:
//                        temp_workorders[-1].next_work_order_id = original_one
//                    original_one = temp_workorders[0]
//        return workorders
})
h.MrpProduction().Methods().WorkordersCreate().DeclareMethod(
`
        :param bom: in case of recursive boms: we could
create work orders for child
                    BoMs
        `,
func(rs m.MrpProductionSet, bom interface{}, bom_data interface{})  {
//        workorders = self.env['mrp.workorder']
//        bom_qty = bom_data['qty']
//        if self.product_id.tracking == 'serial':
//            quantity = 1.0
//        else:
//            quantity = self.product_qty - \
//                sum(self.move_finished_ids.mapped('quantity_done'))
//            quantity = quantity if (quantity > 0) else 0
//        for operation in bom.routing_id.operation_ids:
//            # create workorder
//            # TODO: float_round UP
//            cycle_number = math.ceil(
//                bom_qty / operation.workcenter_id.capacity)
//            duration_expected = (operation.workcenter_id.time_start +
//                                 operation.workcenter_id.time_stop +
//                                 cycle_number * operation.time_cycle * 100.0 / operation.workcenter_id.time_efficiency)
//            workorder = workorders.create({
//                'name': operation.name,
//                'production_id': self.id,
//                'workcenter_id': operation.workcenter_id.id,
//                'operation_id': operation.id,
//                'duration_expected': duration_expected,
//                'state': len(workorders) == 0 and 'ready' or 'pending',
//                'qty_producing': quantity,
//                'capacity': operation.workcenter_id.capacity,
//            })
//            if workorders:
//                workorders[-1].next_work_order_id = workorder.id
//            workorders += workorder
//
//            # assign moves; last operation receive all unassigned moves (which case ?)
//            moves_raw = self.move_raw_ids.filtered(
//                lambda move: move.operation_id == operation)
//            if len(workorders) == len(bom.routing_id.operation_ids):
//                moves_raw |= self.move_raw_ids.filtered(
//                    lambda move: not move.operation_id)
//            # TODO: code does nothing, unless maybe by_products?
//            moves_finished = self.move_finished_ids.filtered(
//                lambda move: move.operation_id == operation)
//            moves_raw.mapped('move_lot_ids').write(
//                {'workorder_id': workorder.id})
//            (moves_finished + moves_raw).write({'workorder_id': workorder.id})
//
//            workorder._generate_lot_ids()
//        return workorders
})
h.MrpProduction().Methods().ActionCancel().DeclareMethod(
` Cancels production order, unfinished stock moves and set procurement
        orders in exception `,
func(rs m.MrpProductionSet)  {
//        if any(workorder.state == 'progress' for workorder in self.mapped('workorder_ids')):
//            raise UserError(
//                _('You can not cancel production order, a work order is still in progress.'))
//        ProcurementOrder = self.env['procurement.order']
//        for production in self:
//            production.workorder_ids.filtered(
//                lambda x: x.state != 'cancel').action_cancel()
//
//            finish_moves = production.move_finished_ids.filtered(
//                lambda x: x.state not in ('done', 'cancel'))
//            raw_moves = production.move_raw_ids.filtered(
//                lambda x: x.state not in ('done', 'cancel'))
//            (finish_moves | raw_moves).action_cancel()
//
//            procurements = ProcurementOrder.search(
//                [('move_dest_id', 'in', (finish_moves | raw_moves).ids)])
//            if procurements:
//                procurements.cancel()
//        ProcurementOrder.search([('production_id', 'in', self.ids)]).write(
//            {'state': 'exception'})
//        self.write({'state': 'cancel'})
//        return True
})
h.MrpProduction().Methods().CalPrice().DeclareMethod(
`CalPrice`,
func(rs m.MrpProductionSet, consumed_moves interface{})  {
//        self.ensure_one()
//        return True
})
h.MrpProduction().Methods().PostInventory().DeclareMethod(
`PostInventory`,
func(rs m.MrpProductionSet)  {
//        for order in self:
//            moves_not_to_do = order.move_raw_ids.filtered(
//                lambda x: x.state == 'done')
//            moves_to_do = order.move_raw_ids.filtered(
//                lambda x: x.state not in ('done', 'cancel'))
//            moves_to_do.action_done()
//            moves_to_do = order.move_raw_ids.filtered(
//                lambda x: x.state == 'done') - moves_not_to_do
//            order._cal_price(moves_to_do)
//            moves_to_finish = order.move_finished_ids.filtered(
//                lambda x: x.state not in ('done', 'cancel'))
//            moves_to_finish.action_done()
//
//            for move in moves_to_finish:
//                # Group quants by lots
//                lot_quants = {}
//                raw_lot_quants = {}
//                quants = self.env['stock.quant']
//                if move.has_tracking != 'none':
//                    for quant in move.quant_ids:
//                        lot_quants.setdefault(
//                            quant.lot_id.id, self.env['stock.quant'])
//                        raw_lot_quants.setdefault(
//                            quant.lot_id.id, self.env['stock.quant'])
//                        lot_quants[quant.lot_id.id] |= quant
//                for move_raw in moves_to_do:
//                    if (move.has_tracking != 'none') and (move_raw.has_tracking != 'none'):
//                        for lot in lot_quants:
//                            lots = move_raw.move_lot_ids.filtered(
//                                lambda x: x.lot_produced_id.id == lot).mapped('lot_id')
//                            raw_lot_quants[lot] |= move_raw.quant_ids.filtered(
//                                lambda x: (x.lot_id in lots) and (x.qty > 0.0))
//                    else:
//                        quants |= move_raw.quant_ids.filtered(
//                            lambda x: x.qty > 0.0)
//                if move.has_tracking != 'none':
//                    for lot in lot_quants:
//                        lot_quants[lot].sudo().write(
//                            {'consumed_quant_ids': [(6, 0, [x.id for x in raw_lot_quants[lot] | quants])]})
//                else:
//                    move.quant_ids.sudo().write(
//                        {'consumed_quant_ids': [(6, 0, [x.id for x in quants])]})
//            order.action_assign()
//        return True
})
h.MrpProduction().Methods().ButtonMarkDone().DeclareMethod(
`ButtonMarkDone`,
func(rs m.MrpProductionSet)  {
//        self.ensure_one()
//        for wo in self.workorder_ids:
//            if wo.time_ids.filtered(lambda x: (not x.date_end) and (x.loss_type in ('productive', 'performance'))):
//                raise UserError(_('Work order %s is still running') % wo.name)
//        self.post_inventory()
//        moves_to_cancel = (self.move_raw_ids | self.move_finished_ids).filtered(
//            lambda x: x.state not in ('done', 'cancel'))
//        moves_to_cancel.action_cancel()
//        self.write({'state': 'done', 'date_finished': fields.Datetime.now()})
//        self.env["procurement.order"].search(
//            [('production_id', 'in', self.ids)]).check()
//        return self.write({'state': 'done'})
})
h.MrpProduction().Methods().DoUnreserve().DeclareMethod(
`DoUnreserve`,
func(rs m.MrpProductionSet)  {
//        for production in self:
//            production.move_raw_ids.filtered(
//                lambda x: x.state not in ('done', 'cancel')).do_unreserve()
//        return True
})
h.MrpProduction().Methods().ButtonUnreserve().DeclareMethod(
`ButtonUnreserve`,
func(rs m.MrpProductionSet)  {
//        self.ensure_one()
//        self.do_unreserve()
//        return True
})
h.MrpProduction().Methods().ButtonScrap().DeclareMethod(
`ButtonScrap`,
func(rs m.MrpProductionSet)  {
//        self.ensure_one()
//        return {
//            'name': _('Scrap'),
//            'view_type': 'form',
//            'view_mode': 'form',
//            'res_model': 'stock.scrap',
//            'view_id': self.env.ref('stock.stock_scrap_form_view2').id,
//            'type': 'ir.actions.act_window',
//            'context': {'default_production_id': self.id,
//                        'product_ids': (self.move_raw_ids.filtered(lambda x: x.state not in ('done', 'cancel')) | self.move_finished_ids.filtered(lambda x: x.state == 'done')).mapped('product_id').ids,
//                        },
//            'target': 'new',
//        }
})
h.MrpProduction().Methods().ActionSeeMoveScrap().DeclareMethod(
`ActionSeeMoveScrap`,
func(rs m.MrpProductionSet)  {
//        self.ensure_one()
//        action = self.env.ref('stock.action_stock_scrap').read()[0]
//        action['domain'] = [('production_id', '=', self.id)]
//        return action
})
}