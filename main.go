package main

import (
	admin_models "goblog/admin/models"
	"goblog/config"
	"net/http"
)

func main() {
	admin_models.Post{}.Migrate()
	admin_models.User{}.Migrate()
	admin_models.Category{}.Migrate()

	/* admin_models.Post{
		Title: "Coding with GoLang",
		Slug:  "coding-with-golang",
	}.Add() */

	// post := admin_models.Post{}.Get(1)
	// fmt.Println(post.Title)
	// post1 := admin_models.Post{}.Get("description = ?", "golangCoding")
	// fmt.Println(post1.Slug)

	// fmt.Println(admin_models.Post{}.GetAll())
	// fmt.Println(admin_models.Post{}.GetAll("description = ?", "golangCoding"))

	// postUpd := admin_models.Post{}.Get(1)
	// postUpd.Update("slug", "web")

	// postUpds := admin_models.Post{}.Get(1)
	// postUpds.Updates(admin_models.Post{Title: "Web Programming with Go!", Description: "goWebTest"})

	// postDel := admin_models.Post{}.Get(1)
	// postDel.Delete()

	http.ListenAndServe(":8080", config.Routes())
}
