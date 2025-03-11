package routers

import (
	"go-pos/controllers"
	
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// Category routes
	beego.Router("/api/categories", &controllers.CategoryController{}, "get:GetAll;post:Create")
	beego.Router("/api/categories/:id", &controllers.CategoryController{}, "get:Get;put:Update;delete:Delete")
	
	// Item routes
	beego.Router("/api/items", &controllers.ItemController{}, "get:GetAll;post:Create")
	beego.Router("/api/items/:id", &controllers.ItemController{}, "get:Get;put:Update;delete:Delete")
	
	// ItemBatch routes
	beego.Router("/api/item-batches", &controllers.ItemBatchController{}, "get:GetAll;post:Create")
	beego.Router("/api/item-batches/:id", &controllers.ItemBatchController{}, "get:Get;put:Update;delete:Delete")
	
	// Member routes
	beego.Router("/api/members", &controllers.MemberController{}, "get:GetAll;post:Create")
	beego.Router("/api/members/:id", &controllers.MemberController{}, "get:Get;put:Update;delete:Delete")
	
	// MemberPoint routes
	beego.Router("/api/member-points", &controllers.MemberController{}, "get:GetAllPoints")
	beego.Router("/api/member-points/:id", &controllers.MemberController{}, "get:GetPoint;put:UpdatePoint")
	beego.Router("/api/member-points", &controllers.MemberPointController{}, "post:Create")
	
	// SalesBasket routes
	beego.Router("/api/sales", &controllers.SalesBasketController{}, "get:GetAll;post:Create")
	beego.Router("/api/sales/:id", &controllers.SalesBasketController{}, "get:Get;put:Update;delete:Delete")
	
	// SalesItem routes
	beego.Router("/api/sales-items", &controllers.SalesItemController{}, "get:GetAll;post:Create")
	beego.Router("/api/sales-items/:id", &controllers.SalesItemController{}, "get:Get;put:Update;delete:Delete")
	
	// User routes
	beego.Router("/api/users", &controllers.UserController{}, "get:GetAll;post:Create")
	beego.Router("/api/users/:id", &controllers.UserController{}, "get:Get;put:Update;delete:Delete")
	
	// UserLog routes
	beego.Router("/api/user-logs", &controllers.UserLogController{}, "get:GetAll;post:Create")
	beego.Router("/api/user-logs/:id", &controllers.UserLogController{}, "get:Get;put:Update;delete:Delete")
	
	// UserMember routes
	beego.Router("/api/user-members", &controllers.UserMemberController{}, "get:GetAll;post:Create")
	beego.Router("/api/user-members/:id", &controllers.UserMemberController{}, "get:Get;put:Update;delete:Delete")
	
	// Authentication routes
	beego.Router("/api/auth/login", &controllers.AuthController{}, "post:Login")
	beego.Router("/api/auth/logout", &controllers.AuthController{}, "post:Logout")
}