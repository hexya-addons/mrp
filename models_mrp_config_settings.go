package mrp

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/types"
	"github.com/hexya-erp/pool/h"
)

func init() {
	h.MrpConfigSettings().DeclareModel()

	h.MrpConfigSettings().AddFields(map[string]models.FieldDefinition{
		"CompanyId": models.Many2OneField{
			RelationModel: h.Company(),
			String:        "Company",
			Required:      true,
			Default:       func(env models.Environment) interface{} { return env.Uid().company_id },
		},
		"ManufacturingLead": models.FloatField{
			Related: `CompanyId.ManufacturingLead`,
			String:  "Manufacturing Lead Time *",
		},
		"GroupProductVariant": models.SelectionField{
			Selection: types.Selection{
				"": "No variants on products",
				"": "Products can have several attributes, defining variants (Example: size, color,...)",
			},
			String: "Product Variants",
			Help: "Work with product variant allows you to define some variant" +
				"of the same products, an ease the product management in" +
				"the ecommerce for example",
			//implied_group='product.group_product_variant'
		},
		"ModuleMrpByproduct": models.SelectionField{
			Selection: types.Selection{
				"": "No by-products in bills of materials (A + B --> C)",
				"": "Bills of materials may produce residual products (A + B --> C + D)",
			},
			String: "By-Products",
			Help: "You can configure by-products in the bill of material." +
				"Without this module: A + B + C -> D." +
				"With this module: A + B + C -> D + E." +
				"-This installs the module mrp_byproduct.",
		},
		"ModuleMrpMps": models.SelectionField{
			Selection: types.Selection{
				"": "No need for Master Production Schedule as products have short lead times",
				"": "Use Master Production Schedule in order to create procurements based on forecasts",
			},
			String: "Master Production Schedule",
		},
		"ModuleMrpPlm": models.SelectionField{
			Selection: types.Selection{
				"": "No product lifecycle management",
				"": "Manage engineering changes, versions and documents",
			},
			String: "PLM",
		},
		"ModuleMrpMaintenance": models.SelectionField{
			Selection: types.Selection{
				"": "No maintenance machine and work centers",
				"": "Preventive and Corrective maintenance management",
			},
			String: "Maintenance",
		},
		"ModuleQualityMrp": models.SelectionField{
			Selection: types.Selection{
				"": "No quality control",
				"": "Manage quality control points, checks and measures",
			},
			String: "Quality",
		},
		"GroupMrpRoutings": models.SelectionField{
			Selection: types.Selection{
				"": "Manage production by manufacturing orders",
				"": "Manage production by work orders",
			},
			String: "Routings & Planning",
			//implied_group='mrp.group_mrp_routings'
			Help: "Work Order Operations allow you to create and manage the" +
				"manufacturing operations that should be followed within" +
				"your work centers in order to produce a product. They are" +
				"attached to bills of materials that will define the required" +
				"raw materials.",
		},
	})
}
