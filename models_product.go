package mrp

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h"
)

func init() {
	h.ProductTemplate().DeclareModel()

	h.ProductTemplate().AddFields(map[string]models.FieldDefinition{
		"BomIds": models.One2ManyField{
			RelationModel: h.MrpBom(),
			ReverseFK:     "",
			String:        "Bill of Materials",
		},
		"BomCount": models.IntegerField{
			String:  "# Bill of Material",
			Compute: h.ProductTemplate().Methods().ComputeBomCount(),
		},
		"MoCount": models.IntegerField{
			String:  "# Manufacturing Orders",
			Compute: h.ProductTemplate().Methods().ComputeMoCount(),
		},
		"ProduceDelay": models.FloatField{
			String:  "Manufacturing Lead Time",
			Default: models.DefaultValue(0),
			Help: "Average delay in days to produce this product. In the case" +
				"of multi-level BOM, the manufacturing lead times of the" +
				"components will be added.",
		},
	})
	h.ProductTemplate().Methods().ComputeBomCount().DeclareMethod(
		`ComputeBomCount`,
		func(rs h.ProductTemplateSet) h.ProductTemplateData {
			//        read_group_res = self.env['mrp.bom'].read_group(
			//            [('product_tmpl_id', 'in', self.ids)], ['product_tmpl_id'], ['product_tmpl_id'])
			//        mapped_data = dict(
			//            [(data['product_tmpl_id'][0], data['product_tmpl_id_count']) for data in read_group_res])
			//        for product in self:
			//            product.bom_count = mapped_data.get(product.id, 0)
		})
	h.ProductTemplate().Methods().ComputeMoCount().DeclareMethod(
		`ComputeMoCount`,
		func(rs h.ProductTemplateSet) h.ProductTemplateData {
			//        self.mo_count = sum(self.mapped(
			//            'product_variant_ids').mapped('mo_count'))
		})
	h.ProductTemplate().Methods().ActionViewMos().DeclareMethod(
		`ActionViewMos`,
		func(rs m.ProductTemplateSet) {
			//        product_ids = self.mapped('product_variant_ids').ids
			//        action = self.env.ref('mrp.act_product_mrp_production').read()[0]
			//        action['domain'] = [('product_id', 'in', product_ids)]
			//        action['context'] = {}
			//        return action
		})
	h.ProductProduct().DeclareModel()

	h.ProductProduct().AddFields(map[string]models.FieldDefinition{
		"BomCount": models.IntegerField{
			String:  "# Bill of Material",
			Compute: h.ProductProduct().Methods().ComputeBomCount(),
		},
		"MoCount": models.IntegerField{
			String:  "# Manufacturing Orders",
			Compute: h.ProductProduct().Methods().ComputeMoCount(),
		},
	})
	h.ProductProduct().Methods().ComputeBomCount().DeclareMethod(
		`ComputeBomCount`,
		func(rs h.ProductProductSet) h.ProductProductData {
			//        read_group_res = self.env['mrp.bom'].read_group(
			//            [('product_id', 'in', self.ids)], ['product_id'], ['product_id'])
			//        mapped_data = dict(
			//            [(data['product_id'][0], data['product_id_count']) for data in read_group_res])
			//        for product in self:
			//            if product.product_tmpl_id.product_variant_count == 1:
			//                bom_count = mapped_data.get(
			//                    product.id, product.product_tmpl_id.bom_count)
			//            else:
			//                bom_count = mapped_data.get(product.id, 0)
			//            product.bom_count = bom_count
		})
	h.ProductProduct().Methods().ComputeMoCount().DeclareMethod(
		`ComputeMoCount`,
		func(rs h.ProductProductSet) h.ProductProductData {
			//        read_group_res = self.env['mrp.production'].read_group(
			//            [('product_id', 'in', self.ids)], ['product_id'], ['product_id'])
			//        mapped_data = dict(
			//            [(data['product_id'][0], data['product_id_count']) for data in read_group_res])
			//        for product in self:
			//            product.mo_count = mapped_data.get(product.id, 0)
		})
	h.ProductProduct().Methods().ActionViewBom().DeclareMethod(
		`ActionViewBom`,
		func(rs m.ProductProductSet) {
			//        action = self.env.ref('mrp.product_open_bom').read()[0]
			//        template_ids = self.mapped('product_tmpl_id').ids
			//        action['context'] = {
			//            'default_product_tmpl_id': template_ids[0],
			//            'default_product_id': self.ids[0],
			//        }
			//        action['domain'] = ['|', ('product_id', 'in', [self.ids]), '&', (
			//            'product_id', '=', False), ('product_tmpl_id', 'in', template_ids)]
			//        return action
		})
}
