package mrp

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/types"
	"github.com/hexya-erp/pool/h"
)

func init() {
	h.MrpRouting().DeclareModel()

	h.MrpRouting().AddFields(map[string]models.FieldDefinition{
		"Name": models.CharField{
			String:   "Routing Name",
			Required: true,
		},
		"Active": models.BooleanField{
			String:  "Active",
			Default: models.DefaultValue(true),
			Help: "If the active field is set to False, it will allow you" +
				"to hide the routing without removing it.",
		},
		"Code": models.CharField{
			String:   "Reference",
			NoCopy:   true,
			Default:  func(env models.Environment) interface{} { return odoo._() },
			ReadOnly: true,
		},
		"Note": models.TextField{
			String: "Description",
		},
		"OperationIds": models.One2ManyField{
			RelationModel: h.MrpRoutingWorkcenter(),
			ReverseFK:     "",
			String:        "Operations",
			NoCopy:        false,
			//oldname='workcenter_lines'
		},
		"LocationId": models.Many2OneField{
			RelationModel: h.StockLocation(),
			String:        "Production Location",
			Help: "Keep empty if you produce at the location where you find" +
				"the raw materials.Set a location if you produce at a fixed" +
				"location. This can be a partner location if you subcontract" +
				"the manufacturing operations.",
		},
		"CompanyId": models.Many2OneField{
			RelationModel: h.Company(),
			String:        "Company",
			Default:       func(env models.Environment) interface{} { return env["res.company"]._company_default_get() },
		},
	})
	h.MrpRouting().Methods().Create().Extend(
		`Create`,
		func(rs m.MrpRoutingSet, vals models.RecordData) {
			//        if 'code' not in vals or vals['code'] == _('New'):
			//            vals['code'] = self.env['ir.sequence'].next_by_code(
			//                'mrp.routing') or _('New')
			//        return super(MrpRouting, self).create(vals)
		})
	h.MrpRoutingWorkcenter().DeclareModel()

	h.MrpRoutingWorkcenter().AddFields(map[string]models.FieldDefinition{
		"Name": models.CharField{
			String:   "Operation",
			Required: true,
		},
		"WorkcenterId": models.Many2OneField{
			RelationModel: h.MrpWorkcenter(),
			String:        "Work Center",
			Required:      true,
		},
		"Sequence": models.IntegerField{
			String:  "Sequence",
			Default: models.DefaultValue(100),
			Help:    "Gives the sequence order when displaying a list of routing Work Centers.",
		},
		"RoutingId": models.Many2OneField{
			RelationModel: h.MrpRouting(),
			String:        "Parent Routing",
			Index:         true,
			OnDelete:      `cascade`,
			Required:      true,
			Help: "The routing contains all the Work Centers used and for" +
				"how long. This will create work orders afterwardswhich" +
				"alters the execution of the manufacturing order. ",
		},
		"Note": models.TextField{
			String: "Description",
		},
		"CompanyId": models.Many2OneField{
			RelationModel: h.Company(),
			String:        "Company",
			ReadOnly:      true,
			Related:       `RoutingId.CompanyId`,
			Stored:        true,
		},
		"Worksheet": models.BinaryField{
			String: "worksheet",
		},
		"TimeMode": models.SelectionField{
			Selection: types.Selection{
				"auto":   "Compute based on real time",
				"manual": "Set duration manually",
			},
			String:  "Duration Computation",
			Default: models.DefaultValue("auto"),
		},
		"TimeModeBatch": models.IntegerField{
			String:  "Based on",
			Default: models.DefaultValue(10),
		},
		"TimeCycleManual": models.FloatField{
			String:  "Manual Duration",
			Default: models.DefaultValue(60),
			Help: "Time in minutes. Is the time used in manual mode, or the" +
				"first time supposed in real time when there are not any work orders yet.",
		},
		"TimeCycle": models.FloatField{
			String:  "Duration",
			Compute: h.MrpRoutingWorkcenter().Methods().ComputeTimeCycle(),
		},
		"WorkorderCount": models.IntegerField{
			String:  "# Work Orders",
			Compute: h.MrpRoutingWorkcenter().Methods().ComputeWorkorderCount(),
		},
		"Batch": models.SelectionField{
			Selection: types.Selection{
				"no":  "Once all products are processed",
				"yes": "Once a minimum number of products is processed",
			},
			String: "Next Operation",
			Help: "Will determine if the next work order will be planned after" +
				"the previous one or after the first Quantity To Process" +
				"of the previous one.",
			Default:  models.DefaultValue("no"),
			Required: true,
		},
		"BatchSize": models.FloatField{
			String:  "Quantity to Process",
			Default: models.DefaultValue(1),
		},
		"WorkorderIds": models.One2ManyField{
			RelationModel: h.MrpWorkorder(),
			ReverseFK:     "",
			String:        "Work Orders",
		},
	})
	h.MrpRoutingWorkcenter().Methods().ComputeTimeCycle().DeclareMethod(
		`ComputeTimeCycle`,
		func(rs h.MrpRoutingWorkcenterSet) h.MrpRoutingWorkcenterData {
			//        manual_ops = self.filtered(
			//            lambda operation: operation.time_mode == 'manual')
			//        for operation in manual_ops:
			//            operation.time_cycle = operation.time_cycle_manual
			//        for operation in self - manual_ops:
			//            data = self.env['mrp.workorder'].read_group([
			//                ('operation_id', '=', operation.id),
			//                ('state', '=', 'done')], ['operation_id', 'duration', 'qty_produced'], ['operation_id'],
			//                limit=operation.time_mode_batch)
			//            count_data = dict(
			//                (item['operation_id'][0], (item['duration'], item['qty_produced'])) for item in data)
			//            if count_data.get(operation.id) and count_data[operation.id][1]:
			//                operation.time_cycle = count_data[operation.id][0] / \
			//                    count_data[operation.id][1]
			//            else:
			//                operation.time_cycle = operation.time_cycle_manual
		})
	h.MrpRoutingWorkcenter().Methods().ComputeWorkorderCount().DeclareMethod(
		`ComputeWorkorderCount`,
		func(rs h.MrpRoutingWorkcenterSet) h.MrpRoutingWorkcenterData {
			//        data = self.env['mrp.workorder'].read_group([
			//            ('operation_id', 'in', self.ids),
			//            ('state', '=', 'done')], ['operation_id'], ['operation_id'])
			//        count_data = dict(
			//            (item['operation_id'][0], item['operation_id_count']) for item in data)
			//        for operation in self:
			//            operation.workorder_count = count_data.get(operation.id, 0)
		})
}
