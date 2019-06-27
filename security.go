package mrp

import (
	"github.com/hexya-addons/base"
	"github.com/hexya-erp/pool/h"
)

//vars

var (
	//User
	GroupMrpUser *security.Group
	//Manager
	GroupMrpManager *security.Group
	//Manage Work Order Operations
	GroupMrpRoutings *security.Group
)


//rights
func init() {
	h.MrpWorkcenterProductivityLoss().Methods().AllowAllToGroup(GroupMrpManager)
	h.MrpWorkcenterProductivityLoss().Methods().Load().AllowGroup(GroupMrpUser)
	h.MrpWorkcenterProductivity().Methods().AllowAllToGroup(GroupMrpUser)
	h.MrpWorkcenter().Methods().Load().AllowGroup(GroupMrpUser)
	h.MrpRouting().Methods().Load().AllowGroup(GroupMrpUser)
	h.MrpRoutingWorkcenter().Methods().Load().AllowGroup(GroupMrpUser)
	h.MrpBom().Methods().Load().AllowGroup(GroupMrpUser)
	h.MrpBomLine().Methods().Load().AllowGroup(GroupMrpUser)
	h.MrpProduction().Methods().AllowAllToGroup(GroupMrpUser)
	h.ProcurementOrder().Methods().AllowAllToGroup(GroupMrpUser)
	h.MrpWorkcenter().Methods().AllowAllToGroup(GroupMrpManager)
	h.Resource.ModelResourceResource().Methods().AllowAllToGroup(GroupMrpManager)
	h.MrpRouting().Methods().AllowAllToGroup(GroupMrpManager)
	h.MrpRoutingWorkcenter().Methods().AllowAllToGroup(GroupMrpManager)
	h.MrpBom().Methods().AllowAllToGroup(GroupMrpManager)
	h.MrpBomLine().Methods().AllowAllToGroup(GroupMrpManager)
	h.Stock.ModelStockLocation().Methods().Load().AllowGroup(GroupMrpUser)
	h.Stock.ModelStockMove().Methods().Load().AllowGroup(GroupMrpUser)
	h.Stock.ModelStockMove().Methods().Write().AllowGroup(GroupMrpUser)
	h.Stock.ModelStockMove().Methods().Create().AllowGroup(GroupMrpUser)
	h.Stock.ModelStockPicking().Methods().AllowAllToGroup(GroupMrpUser)
	h.Stock.ModelStockWarehouse().Methods().Load().AllowGroup(GroupMrpUser)
	h.ProcurementOrder().Methods().AllowAllToGroup(base.GroupUser)
	h.MrpProduction().Methods().Load().AllowGroup(GroupStockUser)
	h.Product.ModelProductProduct().Methods().Load().AllowGroup(GroupMrpUser)
	h.Product.ModelProductTemplate().Methods().Load().AllowGroup(GroupMrpUser)
	h.Product.ModelProductUom().Methods().Load().AllowGroup(GroupMrpUser)
	h.Product.ModelProductSupplierinfo().Methods().AllowAllToGroup(GroupMrpUser)
	h.Base.ModelResPartner().Methods().Load().AllowGroup(GroupMrpUser)
	h.MrpWorkorder().Methods().AllowAllToGroup(GroupMrpUser)
	h.MrpWorkorder().Methods().AllowAllToGroup(GroupMrpManager)
	h.Resource.ModelResourceCalendarLeaves().Methods().AllowAllToGroup(GroupMrpUser)
	h.Resource.ModelResourceCalendarLeaves().Methods().Load().AllowGroup(GroupMrpManager)
	h.Resource.ModelResourceCalendarAttendance().Methods().AllowAllToGroup(GroupMrpUser)
	h.Resource.ModelResourceCalendarAttendance().Methods().AllowAllToGroup(GroupMrpManager)
	h.Product.ModelProductUomCateg().Methods().Load().AllowGroup(GroupMrpUser)
	h.Resource.ModelResourceResource().Methods().Load().AllowGroup(GroupMrpUser)
	h.Resource.ModelResourceResource().Methods().AllowAllToGroup(GroupMrpManager)
	h.Product.ModelProductSupplierinfo().Methods().Load().AllowGroup(GroupMrpManager)
	h.MrpProduction().Methods().Load().AllowGroup(GroupMrpManager)
	h.Stock.ModelStockMove().Methods().Load().AllowGroup(GroupMrpManager)
	h.Stock.ModelStockProductionLot().Methods().AllowAllToGroup(GroupMrpUser)
	h.Stock.ModelStockWarehouseOrderpoint().Methods().Load().AllowGroup(GroupMrpUser)
	h.Stock.ModelStockPicking().Methods().Load().AllowGroup(GroupMrpManager)
	h.MrpBom().Methods().Load().AllowGroup(GroupStockUser)
	h.MrpBomLine().Methods().Load().AllowGroup(GroupStockUser)
	h.Product.ModelProductUomCateg().Methods().AllowAllToGroup(GroupMrpManager)
	h.Product.ModelProductUom().Methods().AllowAllToGroup(GroupMrpManager)
	h.Product.ModelProductCategory().Methods().AllowAllToGroup(GroupMrpManager)
	h.Product.ModelProductTemplate().Methods().AllowAllToGroup(GroupMrpManager)
	h.Product.ModelProductProduct().Methods().AllowAllToGroup(GroupMrpManager)
	h.Product.ModelProductPackaging().Methods().AllowAllToGroup(GroupMrpManager)
	h.Product.ModelProductPricelist().Methods().AllowAllToGroup(GroupMrpManager)
	h.Base.ModelResPartner().Methods().Load().AllowGroup(GroupMrpManager)
	h.Base.ModelResPartner().Methods().Write().AllowGroup(GroupMrpManager)
	h.Base.ModelResPartner().Methods().Create().AllowGroup(GroupMrpManager)
	h.Product.ModelProductPricelistItem().Methods().AllowAllToGroup(GroupMrpManager)
	h.Resource.ModelResourceCalendar().Methods().Load().AllowGroup(GroupMrpUser)
	h.MrpUnbuild().Methods().Load().AllowGroup(GroupMrpUser)
	h.MrpUnbuild().Methods().AllowAllToGroup(GroupMrpManager)
	h.MrpMessage().Methods().Load().AllowGroup(GroupMrpUser)
	h.MrpMessage().Methods().AllowAllToGroup(GroupMrpManager)
	h.StockMoveLots().Methods().AllowAllToGroup(GroupMrpUser)
	h.StockMoveLots().Methods().AllowAllToGroup(GroupMrpManager)
}
