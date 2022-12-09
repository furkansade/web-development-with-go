package config

import (
	"github.com/julienschmidt/httprouter"
	admin "goblog/admin/controllers" //admin - site ayrimi icin
	site "goblog/site/controllers"
	"net/http"
)

func Routes() *httprouter.Router {
	r := httprouter.New()
	// ADMIN
	// Blog Posts
	r.GET("/admin", admin.Dashboard{}.Index)
	r.GET("/admin/new-post", admin.Dashboard{}.NewItem)
	r.POST("/admin/add", admin.Dashboard{}.Add)
	r.GET("/admin/delete/:id", admin.Dashboard{}.Delete)
	r.GET("/admin/edit/:id", admin.Dashboard{}.Edit)
	r.POST("/admin/update/:id", admin.Dashboard{}.Update)

	// User Operations
	r.GET("/admin/login", admin.Userops{}.Index)
	r.POST("/admin/do_login", admin.Userops{}.Login)
	r.GET("/admin/logout", admin.Userops{}.Logout)

	// Categories
	r.GET("/admin/categories", admin.Categories{}.Index)
	r.POST("/admin/categories/add", admin.Categories{}.Add)
	r.GET("/admin/categories/delete/:id", admin.Categories{}.Delete)

	// Site
	r.GET("/", site.Homepage{}.Index)
	r.GET("/posts/:slug", site.Homepage{}.Detail)

	// SERVE FILES
	r.ServeFiles("/admin/assets/*filepath", http.Dir("admin/assets"))
	// admin/assets gelirse admin/assets dizinine yonlendir
	r.ServeFiles("/uploads/*filepath", http.Dir("uploads"))
	r.ServeFiles("/assets/*filepath", http.Dir("site/assets"))
	return r
}
