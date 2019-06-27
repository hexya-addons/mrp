package mrp

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/types"
	"github.com/hexya-erp/pool/h"
	"github.com/hexya-erp/pool/q"
)

func init() {
	h.MrpWorkorder().DeclareModel()

	h.MrpWorkorder().AddFields(map[string]models.FieldDefinition{
		"Name": models.CharField{
			String:   "Work Order",
			Required: true,
			//states={'done': [('readonly', True)], 'cancel': [('readonly', True)]}
		},
		"WorkcenterId": models.Many2OneField{
			RelationModel: h.MrpWorkcenter(),
			String:        "Work Center",
			Required:      true,
			//states={'done': [('readonly', True)], 'cancel': [('readonly', True)]}
		},
		"WorkingState": models.SelectionField{
			Selection: "Workcenter Status",
			Related:   `WorkcenterId.WorkingState`,
			Help:      "Technical: used in views only",
		},
		"ProductionId": models.Many2OneField{
			RelationModel: h.MrpProduction(),
			String:        "Manufacturing Order",
			Index:         true,
			OnDelete:      `cascade`,
			Required:      true,
			//track_visibility='onchange'
			//states={'done': [('readonly', True)], 'cancel': [('readonly', True)]}
		},
		"ProductId": models.Many2OneField{
			RelationModel: h.ProductProduct(),
			String:        "Product",
			Related:       `ProductionId.ProductId`,
			ReadOnly:      true,
			Help:          "Technical: used in views only.",
		},
		"ProductUomId": models.Many2OneField{
			RelationModel: h.ProductUom(),
			String:        "Unit of Measure",
			Related:       `ProductionId.ProductUomId`,
			ReadOnly:      true,
			Help:          "Technical: used in views only.",
		},
		"ProductionAvailability": models.SelectionField{
			Selection: "Stock Availability",
			ReadOnly:  true,
			Related:   `ProductionId.Availability`,
			Stored:    true,
			Help:      "Technical: used in views and domains only.",
		},
		"ProductionState": models.SelectionField{
			Selection: "Production State",
			ReadOnly:  true,
			Related:   `ProductionId.State`,
			Help:      "Technical: used in views only.",
		},
		"ProductTracking": models.SelectionField{
			Selection: "Product Tracking",
			Related:   `ProductionId.ProductId.Tracking`,
			Help:      "Technical: used in views only.",
		},
		"QtyProduction": models.FloatField{
			String:   "Original Production Quantity",
			ReadOnly: true,
			Related:  `ProductionId.ProductQty`,
		},
		"QtyProduced": models.FloatField{
			String:   "Quantity",
			Default:  models.DefaultValue(0),
			ReadOnly: true,
			//digits=dp.get_precision('Product Unit of Measure')
			Help: "The number of products already handled by this work order",
		},
		"QtyProducing": models.FloatField{
			String:  "Currently Produced Quantity",
			Default: models.DefaultValue(1),
			//digits=dp.get_precision('Product Unit of Measure')
			//states={'done': [('readonly', True)], 'cancel': [('readonly', True)]}
		},
		"IsProduced": models.BooleanField{
			Compute: h.MrpWorkorder().Methods().ComputeIsProduced(),
		},
		"State": models.SelectionField{
			Selection: types.Selection{
				"pending":  "Pending",
				"ready":    "Ready",
				"progress": "In Progress",
				"done":     "Finished",
				"cancel":   "Cancelled",
			},
			String:  "Status",
			Default: models.DefaultValue("pending"),
		},
		"DatePlannedStart": models.DateTimeField{
			String: "Scheduled Date Start",
			//states={'done': [('readonly', True)], 'cancel': [('readonly', True)]}
		},
		"DatePlannedFinished": models.DateTimeField{
			String: "Scheduled Date Finished",
			//states={'done': [('readonly', True)], 'cancel': [('readonly', True)]}
		},
		"DateStart": models.DateTimeField{
			String: "Effective Start Date",
			//states={'done': [('readonly', True)], 'cancel': [('readonly', True)]}
		},
		"DateFinished": models.DateTimeField{
			String: "Effective End Date",
			//states={'done': [('readonly', True)], 'cancel': [('readonly', True)]}
		},
		"DurationExpected": models.FloatField{
			String: "Expected Duration",
			//digits=(16, 2)
			//states={'done': [('readonly', True)], 'cancel': [('readonly', True)]}
			Help: "Expected duration (in minutes)",
		},
		"Duration": models.FloatField{
			String:   "Real Duration",
			Compute:  h.MrpWorkorder().Methods().ComputeDuration(),
			ReadOnly: true,
			Stored:   true,
		},
		"DurationUnit": models.FloatField{
			String:   "Duration Per Unit",
			Compute:  h.MrpWorkorder().Methods().ComputeDuration(),
			ReadOnly: true,
			Stored:   true,
		},
		"DurationPercent": models.IntegerField{
			String:  "Duration Deviation (%)",
			Compute: h.MrpWorkorder().Methods().ComputeDuration(),
			//group_operator="avg"
			ReadOnly: true,
			Stored:   true,
		},
		"OperationId": models.Many2OneField{
			RelationModel: h.MrpRoutingWorkcenter(),
			String:        "Operation",
		},
		"Worksheet": models.BinaryField{
			String:   "Worksheet",
			Related:  `OperationId.Worksheet`,
			ReadOnly: true,
		},
		"MoveRawIds": models.One2ManyField{
			RelationModel: h.StockMove(),
			ReverseFK:     "",
			String:        "Moves",
		},
		"MoveLotIds": models.One2ManyField{
			RelationModel: h.StockMoveLots(),
			ReverseFK:     "",
			String:        "Moves to Track",
			Filter:        q.DoneWo().Equals(True),
			Help:          "Inventory moves for which you must scan a lot number at this work order",
		},
		"ActiveMoveLotIds": models.One2ManyField{
			RelationModel: h.StockMoveLots(),
			ReverseFK:     "",
			Filter:        q.DoneWo().Equals(False),
		},
		"FinalLotId": models.Many2OneField{
			RelationModel: h.StockProductionLot(),
			String:        "Current Lot",
			Filter:        q.ProductId().Equals(product_id),
			//states={'done': [('readonly', True)], 'cancel': [('readonly', True)]}
		},
		"TimeIds": models.One2ManyField{
			RelationModel: h.MrpWorkcenterProductivity(),
			ReverseFK:     "",
		},
		"IsUserWorking": models.BooleanField{
			String:  "Is Current User Working",
			Compute: h.MrpWorkorder().Methods().ComputeIsUserWorking(),
			Help:    "Technical field indicating whether the current user is working. ",
		},
		"ProductionMessages": models.HTMLField{
			String:  "Workorder Message",
			Compute: h.MrpWorkorder().Methods().ComputeProductionMessages(),
		},
		"NextWorkOrderId": models.Many2OneField{
			RelationModel: h.MrpWorkorder(),
			String:        "Next Work Order",
		},
		"ScrapIds": models.One2ManyField{
			RelationModel: h.StockScrap(),
			ReverseFK:     "",
		},
		"ScrapCount": models.IntegerField{
			Compute: h.MrpWorkorder().Methods().ComputeScrapMoveCount(),
			String:  "Scrap Move",
		},
		"ProductionDate": models.DateTimeField{
			String:  "Production Date",
			Related: `ProductionId.DatePlannedStart`,
			Stored:  true,
		},
		"Color": models.IntegerField{
			String:  "Color",
			Compute: h.MrpWorkorder().Methods().ComputeColor(),
		},
		"Capacity": models.FloatField{
			String:  "Capacity",
			Default: models.DefaultValue(1),
			Help:    "Number of pieces that can be produced in parallel.",
		},
	})
	h.MrpWorkorder().Methods().ComputeIsProduced().DeclareMethod(
		`ComputeIsProduced`,
		func(rs h.MrpWorkorderSet) h.MrpWorkorderData {
			//        self.is_produced = self.qty_produced >= self.production_id.product_qty
		})
	h.MrpWorkorder().Methods().ComputeDuration().DeclareMethod(
		`ComputeDuration`,
		func(rs h.MrpWorkorderSet) h.MrpWorkorderData {
			//        self.duration = sum(self.time_ids.mapped('duration'))
			//        self.duration_unit = round(
			//            self.duration / max(self.qty_produced, 1), 2)
			//        if self.duration_expected:
			//            self.duration_percent = 100 * \
			//                (self.duration_expected - self.duration) / self.duration_expected
			//        else:
			//            self.duration_percent = 0
		})
	h.MrpWorkorder().Methods().ComputeIsUserWorking().DeclareMethod(
		` Checks whether the current user is working `,
		func(rs h.MrpWorkorderSet) h.MrpWorkorderData {
			//        for order in self:
			//            if order.time_ids.filtered(lambda x: (x.user_id.id == self.env.user.id) and (not x.date_end) and (x.loss_type in ('productive', 'performance'))):
			//                order.is_user_working = True
			//            else:
			//                order.is_user_working = False
		})
	h.MrpWorkorder().Methods().ComputeProductionMessages().DeclareMethod(
		`ComputeProductionMessages`,
		func(rs h.MrpWorkorderSet) h.MrpWorkorderData {
			//        ProductionMessage = self.env['mrp.message']
			//        for workorder in self:
			//            domain = [
			//                ('valid_until', '>=', fields.Date.today()),
			//                '|', ('workcenter_id', '=', False), ('workcenter_id',
			//                                                     '=', workorder.workcenter_id.id),
			//                '|', '|', '|',
			//                ('product_id', '=', workorder.product_id.id),
			//                '&', ('product_id', '=', False), ('product_tmpl_id',
			//                                                  '=', workorder.product_id.product_tmpl_id.id),
			//                ('bom_id', '=', workorder.production_id.bom_id.id),
			//                ('routing_id', '=', workorder.operation_id.routing_id.id)]
			//            messages = ProductionMessage.search(domain).mapped('message')
			//            workorder.production_messages = "<br/>".join(messages)
		})
	h.MrpWorkorder().Methods().ComputeScrapMoveCount().DeclareMethod(
		`ComputeScrapMoveCount`,
		func(rs h.MrpWorkorderSet) h.MrpWorkorderData {
			//        data = self.env['stock.scrap'].read_group([('workorder_id', 'in', self.ids)], [
			//                                                  'workorder_id'], ['workorder_id'])
			//        count_data = dict(
			//            (item['workorder_id'][0], item['workorder_id_count']) for item in data)
			//        for workorder in self:
			//            workorder.scrap_count = count_data.get(workorder.id, 0)
		})
	h.MrpWorkorder().Methods().ComputeColor().DeclareMethod(
		`ComputeColor`,
		func(rs h.MrpWorkorderSet) h.MrpWorkorderData {
			//        late_orders = self.filtered(
			//            lambda x: x.production_id.date_planned_finished and x.date_planned_finished > x.production_id.date_planned_finished)
			//        for order in late_orders:
			//            order.color = 4
			//        for order in (self - late_orders):
			//            order.color = 2
		})
	h.MrpWorkorder().Methods().OnchangeQtyProducing().DeclareMethod(
		` Update stock.move.lot records, according to the new qty currently
        produced. `,
		func(rs m.MrpWorkorderSet) {
			//        moves = self.move_raw_ids.filtered(lambda move: move.state not in (
			//            'done', 'cancel') and move.product_id.tracking != 'none' and move.product_id.id != self.production_id.product_id.id)
			//        for move in moves:
			//            move_lots = self.active_move_lot_ids.filtered(
			//                lambda move_lot: move_lot.move_id == move)
			//            if not move_lots:
			//                continue
			//            new_qty = move.unit_factor * self.qty_producing
			//            if move.product_id.tracking == 'lot':
			//                move_lots[0].quantity = new_qty
			//                move_lots[0].quantity_done = new_qty
			//            elif move.product_id.tracking == 'serial':
			//                # Create extra pseudo record
			//                qty_todo = new_qty - sum(move_lots.mapped('quantity'))
			//                if float_compare(qty_todo, 0.0, precision_rounding=move.product_uom.rounding) > 0:
			//                    while float_compare(qty_todo, 0.0, precision_rounding=move.product_uom.rounding) > 0:
			//                        self.active_move_lot_ids += self.env['stock.move.lots'].new({
			//                            'move_id': move.id,
			//                            'product_id': move.product_id.id,
			//                            'lot_id': False,
			//                            'quantity': min(1.0, qty_todo),
			//                            'quantity_done': min(1.0, qty_todo),
			//                            'workorder_id': self.id,
			//                            'done_wo': False
			//                        })
			//                        qty_todo -= 1
			//                elif float_compare(qty_todo, 0.0, precision_rounding=move.product_uom.rounding) < 0:
			//                    qty_todo = abs(qty_todo)
			//                    for move_lot in move_lots:
			//                        if qty_todo <= 0:
			//                            break
			//                        if not move_lot.lot_id and qty_todo >= move_lot.quantity:
			//                            qty_todo = qty_todo - move_lot.quantity
			//                            self.active_move_lot_ids -= move_lot  # Difference operator
			//                        else:
			//                            move_lot.quantity = move_lot.quantity - qty_todo
			//                            if move_lot.quantity_done - qty_todo > 0:
			//                                move_lot.quantity_done = move_lot.quantity_done - qty_todo
			//                            else:
			//                                move_lot.quantity_done = 0
			//                            qty_todo = 0
		})
	h.MrpWorkorder().Methods().Write().Extend(
		`Write`,
		func(rs m.MrpWorkorderSet, values models.RecordData) {
			//        if ('date_planned_start' in values or 'date_planned_finished' in values) and any(workorder.state == 'done' for workorder in self):
			//            raise UserError(_('You can not change the finished work order.'))
			//        return super(MrpWorkorder, self).write(values)
		})
	h.MrpWorkorder().Methods().GenerateLotIds().DeclareMethod(
		` Generate stock move lots `,
		func(rs m.MrpWorkorderSet) {
			//        self.ensure_one()
			//        MoveLot = self.env['stock.move.lots']
			//        tracked_moves = self.move_raw_ids.filtered(
			//            lambda move: move.state not in ('done', 'cancel') and move.product_id.tracking != 'none' and move.product_id != self.production_id.product_id)
			//        for move in tracked_moves:
			//            qty = move.unit_factor * self.qty_producing
			//            if move.product_id.tracking == 'serial':
			//                while float_compare(qty, 0.0, precision_rounding=move.product_uom.rounding) > 0:
			//                    MoveLot.create({
			//                        'move_id': move.id,
			//                        'quantity': min(1, qty),
			//                        'quantity_done': min(1, qty),
			//                        'production_id': self.production_id.id,
			//                        'workorder_id': self.id,
			//                        'product_id': move.product_id.id,
			//                        'done_wo': False,
			//                    })
			//                    qty -= 1
			//            else:
			//                MoveLot.create({
			//                    'move_id': move.id,
			//                    'quantity': qty,
			//                    'quantity_done': qty,
			//                    'product_id': move.product_id.id,
			//                    'production_id': self.production_id.id,
			//                    'workorder_id': self.id,
			//                    'done_wo': False,
			//                })
		})
	h.MrpWorkorder().Methods().RecordProduction().DeclareMethod(
		`RecordProduction`,
		func(rs m.MrpWorkorderSet) {
			//        self.ensure_one()
			//        if self.qty_producing <= 0:
			//            raise UserError(
			//                _('Please set the quantity you produced in the Current Qty field. It can not be 0!'))
			//        if (self.production_id.product_id.tracking != 'none') and not self.final_lot_id:
			//            raise UserError(
			//                _('You should provide a lot for the final product'))
			//        raw_moves = self.move_raw_ids.filtered(lambda x: (x.has_tracking == 'none') and (
			//            x.state not in ('done', 'cancel')) and x.bom_line_id)
			//        for move in raw_moves:
			//            if move.unit_factor:
			//                rounding = move.product_uom.rounding
			//                move.quantity_done += float_round(
			//                    self.qty_producing * move.unit_factor, precision_rounding=rounding)
			//        for move_lot in self.active_move_lot_ids:
			//            # Check if move_lot already exists
			//            if move_lot.quantity_done <= 0:  # rounding...
			//                move_lot.sudo().unlink()
			//                continue
			//            if not move_lot.lot_id:
			//                raise UserError(_('You should provide a lot for a component'))
			//            # Search other move_lot where it could be added:
			//            lots = self.move_lot_ids.filtered(lambda x: (x.lot_id.id == move_lot.lot_id.id) and (
			//                not x.lot_produced_id) and (not x.done_move))
			//            if lots:
			//                lots[0].quantity_done += move_lot.quantity_done
			//                lots[0].lot_produced_id = self.final_lot_id.id
			//                move_lot.sudo().unlink()
			//            else:
			//                move_lot.lot_produced_id = self.final_lot_id.id
			//                move_lot.done_wo = True
			//        if self.next_work_order_id.state == 'pending':
			//            self.next_work_order_id.state = 'ready'
			//        if self.next_work_order_id and self.final_lot_id and not self.next_work_order_id.final_lot_id:
			//            self.next_work_order_id.final_lot_id = self.final_lot_id.id
			//        self.move_lot_ids.filtered(
			//            lambda move_lot: not move_lot.done_move and not move_lot.lot_produced_id and move_lot.quantity_done > 0
			//        ).write({
			//            'lot_produced_id': self.final_lot_id.id,
			//            'lot_produced_qty': self.qty_producing
			//        })
			//        if not self.next_work_order_id:
			//            production_moves = self.production_id.move_finished_ids.filtered(
			//                lambda x: (x.state not in ('done', 'cancel')))
			//            for production_move in production_moves:
			//                if production_move.product_id.id == self.production_id.product_id.id and production_move.product_id.tracking != 'none':
			//                    move_lot = production_move.move_lot_ids.filtered(
			//                        lambda x: x.lot_id.id == self.final_lot_id.id)
			//                    if move_lot:
			//                        move_lot.quantity += self.qty_producing
			//                        move_lot.quantity_done += self.qty_producing
			//                    else:
			//                        move_lot.create({'move_id': production_move.id,
			//                                         'lot_id': self.final_lot_id.id,
			//                                         'quantity': self.qty_producing,
			//                                         'quantity_done': self.qty_producing,
			//                                         'workorder_id': self.id,
			//                                         })
			//                elif production_move.unit_factor:
			//                    rounding = production_move.product_uom.rounding
			//                    production_move.quantity_done += float_round(
			//                        self.qty_producing * production_move.unit_factor, precision_rounding=rounding)
			//                else:
			//                    production_move.quantity_done += self.qty_producing  # TODO: UoM conversion?
			//        self.qty_produced += self.qty_producing
			//        if self.qty_produced >= self.production_id.product_qty:
			//            self.qty_producing = 0
			//        elif self.production_id.product_id.tracking == 'serial':
			//            self.qty_producing = 1.0
			//            self._generate_lot_ids()
			//        else:
			//            self.qty_producing = self.production_id.product_qty - self.qty_produced
			//            self._generate_lot_ids()
			//        self.final_lot_id = False
			//        if self.qty_produced >= self.production_id.product_qty:
			//            self.button_finish()
			//        return True
		})
	h.MrpWorkorder().Methods().ButtonStart().DeclareMethod(
		`ButtonStart`,
		func(rs m.MrpWorkorderSet) {
			//        timeline = self.env['mrp.workcenter.productivity']
			//        if self.duration < self.duration_expected:
			//            loss_id = self.env['mrp.workcenter.productivity.loss'].search(
			//                [('loss_type', '=', 'productive')], limit=1)
			//            if not len(loss_id):
			//                raise UserError(
			//                    _("You need to define at least one productivity loss in the category 'Productivity'. Create one from the Manufacturing app, menu: Configuration / Productivity Losses."))
			//        else:
			//            loss_id = self.env['mrp.workcenter.productivity.loss'].search(
			//                [('loss_type', '=', 'performance')], limit=1)
			//            if not len(loss_id):
			//                raise UserError(
			//                    _("You need to define at least one productivity loss in the category 'Performance'. Create one from the Manufacturing app, menu: Configuration / Productivity Losses."))
			//        for workorder in self:
			//            if workorder.production_id.state != 'progress':
			//                workorder.production_id.write({
			//                    'state': 'progress',
			//                    'date_start': datetime.now(),
			//                })
			//            timeline.create({
			//                'workorder_id': workorder.id,
			//                'workcenter_id': workorder.workcenter_id.id,
			//                'description': _('Time Tracking: ')+self.env.user.name,
			//                'loss_id': loss_id[0].id,
			//                'date_start': datetime.now(),
			//                'user_id': self.env.user.id
			//            })
			//        return self.write({'state': 'progress',
			//                           'date_start': datetime.now(),
			//                           })
		})
	h.MrpWorkorder().Methods().ButtonFinish().DeclareMethod(
		`ButtonFinish`,
		func(rs m.MrpWorkorderSet) {
			//        self.ensure_one()
			//        self.end_all()
			//        return self.write({'state': 'done', 'date_finished': fields.Datetime.now()})
		})
	h.MrpWorkorder().Methods().EndPrevious().DeclareMethod(
		`
        @param: doall:  This will close all open time lines
on the open work orders when doall = True, otherwise
        only the one of the current user
        `,
		func(rs m.MrpWorkorderSet, doall interface{}) {
			//        timeline_obj = self.env['mrp.workcenter.productivity']
			//        domain = [('workorder_id', 'in', self.ids), ('date_end', '=', False)]
			//        if not doall:
			//            domain.append(('user_id', '=', self.env.user.id))
			//        not_productive_timelines = timeline_obj.browse()
			//        for timeline in timeline_obj.search(domain, limit=None if doall else 1):
			//            wo = timeline.workorder_id
			//            if wo.duration_expected <= wo.duration:
			//                if timeline.loss_type == 'productive':
			//                    not_productive_timelines += timeline
			//                timeline.write({'date_end': fields.Datetime.now()})
			//            else:
			//                maxdate = fields.Datetime.from_string(
			//                    timeline.date_start) + relativedelta(minutes=wo.duration_expected - wo.duration)
			//                enddate = datetime.now()
			//                if maxdate > enddate:
			//                    timeline.write({'date_end': enddate})
			//                else:
			//                    timeline.write({'date_end': maxdate})
			//                    not_productive_timelines += timeline.copy(
			//                        {'date_start': maxdate, 'date_end': enddate})
			//        if not_productive_timelines:
			//            loss_id = self.env['mrp.workcenter.productivity.loss'].search(
			//                [('loss_type', '=', 'performance')], limit=1)
			//            if not len(loss_id):
			//                raise UserError(
			//                    _("You need to define at least one unactive productivity loss in the category 'Performance'. Create one from the Manufacturing app, menu: Configuration / Productivity Losses."))
			//            not_productive_timelines.write({'loss_id': loss_id.id})
			//        return True
		})
	h.MrpWorkorder().Methods().EndAll().DeclareMethod(
		`EndAll`,
		func(rs m.MrpWorkorderSet) {
			//        return self.end_previous(doall=True)
		})
	h.MrpWorkorder().Methods().ButtonPending().DeclareMethod(
		`ButtonPending`,
		func(rs m.MrpWorkorderSet) {
			//        self.end_previous()
			//        return True
		})
	h.MrpWorkorder().Methods().ButtonUnblock().DeclareMethod(
		`ButtonUnblock`,
		func(rs m.MrpWorkorderSet) {
			//        for order in self:
			//            order.workcenter_id.unblock()
			//        return True
		})
	h.MrpWorkorder().Methods().ActionCancel().DeclareMethod(
		`ActionCancel`,
		func(rs m.MrpWorkorderSet) {
			//        return self.write({'state': 'cancel'})
		})
	h.MrpWorkorder().Methods().ButtonDone().DeclareMethod(
		`ButtonDone`,
		func(rs m.MrpWorkorderSet) {
			//        if any([x.state in ('done', 'cancel') for x in self]):
			//            raise UserError(
			//                _('A Manufacturing Order is already done or cancelled!'))
			//        self.end_all()
			//        return self.write({'state': 'done',
			//                           'date_finished': datetime.now()})
		})
	h.MrpWorkorder().Methods().ButtonScrap().DeclareMethod(
		`ButtonScrap`,
		func(rs m.MrpWorkorderSet) {
			//        self.ensure_one()
			//        return {
			//            'name': _('Scrap'),
			//            'view_type': 'form',
			//            'view_mode': 'form',
			//            'res_model': 'stock.scrap',
			//            'view_id': self.env.ref('stock.stock_scrap_form_view2').id,
			//            'type': 'ir.actions.act_window',
			//            'context': {'default_workorder_id': self.id, 'default_production_id': self.production_id.id, 'product_ids': (self.production_id.move_raw_ids.filtered(lambda x: x.state not in ('done', 'cancel')) | self.production_id.move_finished_ids.filtered(lambda x: x.state == 'done')).mapped('product_id').ids},
			//            # 'context': {'product_ids': self.move_raw_ids.filtered(lambda x: x.state not in ('done', 'cancel')).mapped('product_id').ids + [self.production_id.product_id.id]},
			//            'target': 'new',
			//        }
		})
	h.MrpWorkorder().Methods().ActionSeeMoveScrap().DeclareMethod(
		`ActionSeeMoveScrap`,
		func(rs m.MrpWorkorderSet) {
			//        self.ensure_one()
			//        action = self.env.ref('stock.action_stock_scrap').read()[0]
			//        action['domain'] = [('workorder_id', '=', self.id)]
			//        return action
		})
}
