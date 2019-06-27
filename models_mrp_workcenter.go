package mrp

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/types"
	"github.com/hexya-erp/hexya/src/models/types/dates"
	"github.com/hexya-erp/pool/h"
)

//import datetime
func init() {
	h.MrpWorkcenter().DeclareModel()

	h.MrpWorkcenter().AddFields(map[string]models.FieldDefinition{
		"Note": models.TextField{
			String: "Description",
			Help:   "Description of the Work Center.",
		},
		"Capacity": models.FloatField{
			String:  "Capacity",
			Default: models.DefaultValue(1),
			//oldname='capacity_per_cycle'
			Help: "Number of pieces that can be produced in parallel.",
		},
		"Sequence": models.IntegerField{
			String:   "Sequence",
			Default:  models.DefaultValue(1),
			Required: true,
			Help:     "Gives the sequence order when displaying a list of work centers.",
		},
		"Color": models.IntegerField{
			String: "Color",
		},
		"TimeStart": models.FloatField{
			String: "Time before prod.",
			Help:   "Time in minutes for the setup.",
		},
		"TimeStop": models.FloatField{
			String: "Time after prod.",
			Help:   "Time in minutes for the cleaning.",
		},
		"ResourceId": models.Many2OneField{
			RelationModel: h.ResourceResource(),
			String:        "Resource",
			OnDelete:      `cascade`,
			Required:      true,
		},
		"RoutingLineIds": models.One2ManyField{
			RelationModel: h.MrpRoutingWorkcenter(),
			ReverseFK:     "",
			String:        "Routing Lines",
		},
		"OrderIds": models.One2ManyField{
			RelationModel: h.MrpWorkorder(),
			ReverseFK:     "",
			String:        "Orders",
		},
		"WorkorderCount": models.IntegerField{
			String:  "# Work Orders",
			Compute: h.MrpWorkcenter().Methods().ComputeWorkorderCount(),
		},
		"WorkorderReadyCount": models.IntegerField{
			String:  "# Read Work Orders",
			Compute: h.MrpWorkcenter().Methods().ComputeWorkorderCount(),
		},
		"WorkorderProgressCount": models.IntegerField{
			String:  "Total Running Orders",
			Compute: h.MrpWorkcenter().Methods().ComputeWorkorderCount(),
		},
		"WorkorderPendingCount": models.IntegerField{
			String:  "Total Running Orders",
			Compute: h.MrpWorkcenter().Methods().ComputeWorkorderCount(),
		},
		"WorkorderLateCount": models.IntegerField{
			String:  "Total Late Orders",
			Compute: h.MrpWorkcenter().Methods().ComputeWorkorderCount(),
		},
		"TimeIds": models.One2ManyField{
			RelationModel: h.MrpWorkcenterProductivity(),
			ReverseFK:     "",
			String:        "Time Logs",
		},
		"WorkingState": models.SelectionField{
			Selection: types.Selection{
				"normal":  "Normal",
				"blocked": "Blocked",
				"done":    "In Progress",
			},
			String:  "Status",
			Compute: h.MrpWorkcenter().Methods().ComputeWorkingState(),
			Stored:  true,
		},
		"BlockedTime": models.FloatField{
			String:  "Blocked Time",
			Compute: h.MrpWorkcenter().Methods().ComputeBlockedTime(),
			Help:    "Blocked hour(s) over the last month",
			//digits=(16, 2)
		},
		"ProductiveTime": models.FloatField{
			String:  "Productive Time",
			Compute: h.MrpWorkcenter().Methods().ComputeProductiveTime(),
			Help:    "Productive hour(s) over the last month",
			//digits=(16, 2)
		},
		"Oee": models.FloatField{
			Compute: h.MrpWorkcenter().Methods().ComputeOee(),
			Help:    "Overall Equipment Effectiveness, based on the last month",
		},
		"OeeTarget": models.FloatField{
			String:  "OEE Target",
			Help:    "OEE Target in percentage",
			Default: models.DefaultValue(90),
		},
		"Performance": models.IntegerField{
			String:  "Performance",
			Compute: h.MrpWorkcenter().Methods().ComputePerformance(),
			Help:    "Performance over the last month",
		},
		"WorkcenterLoad": models.FloatField{
			String:  "Work Center Load",
			Compute: h.MrpWorkcenter().Methods().ComputeWorkorderCount(),
		},
	})
	h.MrpWorkcenter().Methods().ComputeWorkorderCount().DeclareMethod(
		`ComputeWorkorderCount`,
		func(rs h.MrpWorkcenterSet) h.MrpWorkcenterData {
			//        MrpWorkorder = self.env['mrp.workorder']
			//        result = {wid: {} for wid in self.ids}
			//        result_duration_expected = {wid: 0 for wid in self.ids}
			//        data = MrpWorkorder.read_group([('workcenter_id', 'in', self.ids), ('state', 'in', ('pending', 'ready')), (
			//            'date_planned_start', '<', datetime.datetime.now().strftime('%Y-%m-%d'))], ['workcenter_id'], ['workcenter_id'])
			//        count_data = dict(
			//            (item['workcenter_id'][0], item['workcenter_id_count']) for item in data)
			//        res = MrpWorkorder.read_group(
			//            [('workcenter_id', 'in', self.ids)],
			//            ['workcenter_id', 'state', 'duration_expected'], [
			//                'workcenter_id', 'state'],
			//            lazy=False)
			//        for res_group in res:
			//            result[res_group['workcenter_id'][0]
			//                   ][res_group['state']] = res_group['__count']
			//            if res_group['state'] in ('pending', 'ready', 'progress'):
			//                result_duration_expected[res_group['workcenter_id']
			//                                         [0]] += res_group['duration_expected']
			//        for workcenter in self:
			//            workcenter.workorder_count = sum(
			//                count for state, count in result[workcenter.id].items() if state not in ('done', 'cancel'))
			//            workcenter.workorder_pending_count = result[workcenter.id].get(
			//                'pending', 0)
			//            workcenter.workcenter_load = result_duration_expected[workcenter.id]
			//            workcenter.workorder_ready_count = result[workcenter.id].get(
			//                'ready', 0)
			//            workcenter.workorder_progress_count = result[workcenter.id].get(
			//                'progress', 0)
			//            workcenter.workorder_late_count = count_data.get(workcenter.id, 0)
		})
	h.MrpWorkcenter().Methods().ComputeWorkingState().DeclareMethod(
		`ComputeWorkingState`,
		func(rs h.MrpWorkcenterSet) h.MrpWorkcenterData {
			//        for workcenter in self:
			//            # We search for a productivity line associated to this workcenter having no `date_end`.
			//            # If we do not find one, the workcenter is not currently being used. If we find one, according
			//            # to its `type_loss`, the workcenter is either being used or blocked.
			//            time_log = self.env['mrp.workcenter.productivity'].search([
			//                ('workcenter_id', '=', workcenter.id),
			//                ('date_end', '=', False)
			//            ], limit=1)
			//            if not time_log:
			//                # the workcenter is not being used
			//                workcenter.working_state = 'normal'
			//            elif time_log.loss_type in ('productive', 'performance'):
			//                # the productivity line has a `loss_type` that means the workcenter is being used
			//                workcenter.working_state = 'done'
			//            else:
			//                # the workcenter is blocked
			//                workcenter.working_state = 'blocked'
		})
	h.MrpWorkcenter().Methods().ComputeBlockedTime().DeclareMethod(
		`ComputeBlockedTime`,
		func(rs h.MrpWorkcenterSet) h.MrpWorkcenterData {
			//        data = self.env['mrp.workcenter.productivity'].read_group([
			//            ('date_start', '>=', fields.Datetime.to_string(
			//                datetime.datetime.now() - relativedelta.relativedelta(months=1))),
			//            ('workcenter_id', 'in', self.ids),
			//            ('date_end', '!=', False),
			//            ('loss_type', '!=', 'productive')],
			//            ['duration', 'workcenter_id'], ['workcenter_id'], lazy=False)
			//        count_data = dict(
			//            (item['workcenter_id'][0], item['duration']) for item in data)
			//        for workcenter in self:
			//            workcenter.blocked_time = count_data.get(workcenter.id, 0.0) / 60.0
		})
	h.MrpWorkcenter().Methods().ComputeProductiveTime().DeclareMethod(
		`ComputeProductiveTime`,
		func(rs h.MrpWorkcenterSet) h.MrpWorkcenterData {
			//        data = self.env['mrp.workcenter.productivity'].read_group([
			//            ('date_start', '>=', fields.Datetime.to_string(
			//                datetime.datetime.now() - relativedelta.relativedelta(months=1))),
			//            ('workcenter_id', 'in', self.ids),
			//            ('date_end', '!=', False),
			//            ('loss_type', '=', 'productive')],
			//            ['duration', 'workcenter_id'], ['workcenter_id'], lazy=False)
			//        count_data = dict(
			//            (item['workcenter_id'][0], item['duration']) for item in data)
			//        for workcenter in self:
			//            workcenter.productive_time = count_data.get(
			//                workcenter.id, 0.0) / 60.0
		})
	h.MrpWorkcenter().Methods().ComputeOee().DeclareMethod(
		`ComputeOee`,
		func(rs h.MrpWorkcenterSet) h.MrpWorkcenterData {
			//        for order in self:
			//            if order.productive_time:
			//                order.oee = round(order.productive_time * 100.0 /
			//                                  (order.productive_time + order.blocked_time), 2)
			//            else:
			//                order.oee = 0.0
		})
	h.MrpWorkcenter().Methods().ComputePerformance().DeclareMethod(
		`ComputePerformance`,
		func(rs h.MrpWorkcenterSet) h.MrpWorkcenterData {
			//        wo_data = self.env['mrp.workorder'].read_group([
			//            ('date_start', '>=', fields.Datetime.to_string(
			//                datetime.datetime.now() - relativedelta.relativedelta(months=1))),
			//            ('workcenter_id', 'in', self.ids),
			//            ('state', '=', 'done')], ['duration_expected', 'workcenter_id', 'duration'], ['workcenter_id'], lazy=False)
			//        duration_expected = dict(
			//            (data['workcenter_id'][0], data['duration_expected']) for data in wo_data)
			//        duration = dict((data['workcenter_id'][0], data['duration'])
			//                        for data in wo_data)
			//        for workcenter in self:
			//            if duration.get(workcenter.id):
			//                workcenter.performance = 100 * \
			//                    duration_expected.get(
			//                        workcenter.id, 0.0) / duration[workcenter.id]
			//            else:
			//                workcenter.performance = 0.0
		})
	h.MrpWorkcenter().Methods().CheckCapacity().DeclareMethod(
		`CheckCapacity`,
		func(rs m.MrpWorkcenterSet) {
			//        if any(workcenter.capacity <= 0.0 for workcenter in self):
			//            raise exceptions.UserError(
			//                _('The capacity must be strictly positive.'))
		})
	h.MrpWorkcenter().Methods().Unblock().DeclareMethod(
		`Unblock`,
		func(rs m.MrpWorkcenterSet) {
			//        self.ensure_one()
			//        if self.working_state != 'blocked':
			//            raise exceptions.UserError(_("It has been unblocked already. "))
			//        times = self.env['mrp.workcenter.productivity'].search(
			//            [('workcenter_id', '=', self.id), ('date_end', '=', False)])
			//        times.write({'date_end': fields.Datetime.now()})
			//        return {'type': 'ir.actions.client', 'tag': 'reload'}
		})
	h.MrpWorkcenterProductivityLoss().DeclareModel()

	h.MrpWorkcenterProductivityLoss().AddFields(map[string]models.FieldDefinition{
		"Name": models.CharField{
			String:   "Reason",
			Required: true,
		},
		"Sequence": models.IntegerField{
			String:  "Sequence",
			Default: models.DefaultValue(1),
		},
		"Manual": models.BooleanField{
			String:  "Is a Blocking Reason",
			Default: models.DefaultValue(true),
		},
		"LossType": models.SelectionField{
			Selection: types.Selection{
				"availability": "Availability",
				"performance":  "Performance",
				"quality":      "Quality",
				"productive":   "Productive",
			},
			String:   "Effectiveness Category",
			Default:  models.DefaultValue("availability"),
			Required: true,
		},
	})
	h.MrpWorkcenterProductivity().DeclareModel()

	h.MrpWorkcenterProductivity().AddFields(map[string]models.FieldDefinition{
		"WorkcenterId": models.Many2OneField{
			RelationModel: h.MrpWorkcenter(),
			String:        "Work Center",
			Required:      true,
		},
		"WorkorderId": models.Many2OneField{
			RelationModel: h.MrpWorkorder(),
			String:        "Work Order",
		},
		"UserId": models.Many2OneField{
			RelationModel: h.User(),
			String:        "User",
			Default:       func(env models.Environment) interface{} { return env.uid },
		},
		"LossId": models.Many2OneField{
			RelationModel: h.MrpWorkcenterProductivityLoss(),
			String:        "Loss Reason",
			OnDelete:      `restrict`,
			Required:      true,
		},
		"LossType": models.SelectionField{
			Selection: "Effectiveness",
			Related:   `LossId.LossType`,
			Stored:    true,
		},
		"Description": models.TextField{
			String: "Description",
		},
		"DateStart": models.DateTimeField{
			String:   "Start Date",
			Default:  func(env models.Environment) interface{} { return dates.Now() },
			Required: true,
		},
		"DateEnd": models.DateTimeField{
			String: "End Date",
		},
		"Duration": models.FloatField{
			String:  "Duration",
			Compute: h.MrpWorkcenterProductivity().Methods().ComputeDuration(),
			Stored:  true,
		},
	})
	h.MrpWorkcenterProductivity().Methods().ComputeDuration().DeclareMethod(
		`ComputeDuration`,
		func(rs h.MrpWorkcenterProductivitySet) h.MrpWorkcenterProductivityData {
			//        for blocktime in self:
			//            if blocktime.date_end:
			//                diff = fields.Datetime.from_string(
			//                    blocktime.date_end) - fields.Datetime.from_string(blocktime.date_start)
			//                blocktime.duration = round(diff.total_seconds() / 60.0, 2)
			//            else:
			//                blocktime.duration = 0.0
		})
	h.MrpWorkcenterProductivity().Methods().ButtonBlock().DeclareMethod(
		`ButtonBlock`,
		func(rs m.MrpWorkcenterProductivitySet) {
			//        self.ensure_one()
			//        self.workcenter_id.order_ids.end_all()
			//        return {'type': 'ir.actions.client', 'tag': 'reload'}
		})
}
