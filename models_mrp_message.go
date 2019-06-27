package mrp

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h"
	"github.com/hexya-erp/pool/q"
)

func init() {
	h.MrpMessage().DeclareModel()

	h.MrpMessage().Methods().DefaultValidUntil().DeclareMethod(
		`DefaultValidUntil`,
		func(rs m.MrpMessageSet) {
			//        return datetime.today() + relativedelta(days=7)
		})
	h.MrpMessage().AddFields(map[string]models.FieldDefinition{
		"Name": models.TextField{
			Compute: h.MrpMessage().Methods().GetNoteFirstLine(),
			Stored:  true,
		},
		"Message": models.HTMLField{
			Required: true,
		},
		"ProductTmplId": models.Many2OneField{
			RelationModel: h.ProductTemplate(),
			String:        "Product Template",
		},
		"ProductId": models.Many2OneField{
			RelationModel: h.ProductProduct(),
			String:        "Product",
		},
		"BomId": models.Many2OneField{
			RelationModel: h.MrpBom(),
			String:        "Bill of Material",
			Filter:        q.ProductId().Equals(product_id).Or().ProductTmplId().ProductVariantIds().Equals(product_id),
		},
		"WorkcenterId": models.Many2OneField{
			RelationModel: h.MrpWorkcenter(),
			String:        "Work Center",
		},
		"ValidUntil": models.DateField{
			String:   "Validity Date",
			Default:  models.DefaultValue(_default_valid_until),
			Required: true,
		},
		"RoutingId": models.Many2OneField{
			RelationModel: h.MrpRouting(),
			String:        "Routing",
		},
	})
	h.MrpMessage().Methods().GetNoteFirstLine().DeclareMethod(
		`GetNoteFirstLine`,
		func(rs h.MrpMessageSet) h.MrpMessageData {
			//        for message in self:
			//            message.name = (message.message and html2plaintext(
			//                message.message) or "").strip().replace('*', '').split("\n")[0]
		})
	h.MrpMessage().Methods().Save().DeclareMethod(
		` Used in a wizard-like form view, manual save button when in edit mode `,
		func(rs m.MrpMessageSet) {
			//        return True
		})
	h.MrpMessage().Methods().Create().Extend(
		`Create`,
		func(rs m.MrpMessageSet, vals models.RecordData) {
			//        print vals
			//        return super(MrpProductionMessage, self).create(vals)
		})
}
