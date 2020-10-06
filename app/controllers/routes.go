package controllers

import (
	"Rest/app/middlewares"
	u "Rest/app/utils"
	"net/http"
)

func (a *App) InitializeRoutes() {
	a.Router.HandleFunc("/",a.home)
	// Login Route
	a.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(a.Login)).Methods("POST")

	//Users routes
	a.Router.HandleFunc("/createUser", middlewares.SetMiddlewareJSON(a.CreateUser)).Methods("POST")
	a.Router.HandleFunc("/getUsers", middlewares.SetMiddlewareJSON(a.GetUsers)).Methods("GET")
	a.Router.HandleFunc("/getUser/{id}", middlewares.SetMiddlewareJSON(a.GetUser)).Methods("GET")
	a.Router.HandleFunc("/updateUser/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(a.UpdateUser))).Methods("PUT")
	a.Router.HandleFunc("/deleteUser/{id}", middlewares.SetMiddlewareAuthentication(a.DeleteUser)).Methods("DELETE")

	//Posts routes
	a.Router.HandleFunc("/createProduct", middlewares.SetMiddlewareJSON(a.CreateProduct)).Methods("POST")
	a.Router.HandleFunc("/getProducts", middlewares.SetMiddlewareJSON(a.GetProducts)).Methods("GET")
	a.Router.HandleFunc("/getProduct/{id}", middlewares.SetMiddlewareJSON(a.GetProduct)).Methods("GET")
	a.Router.HandleFunc("/updateProduct/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(a.UpdateProduct))).Methods("PUT")
	a.Router.HandleFunc("/deleteProduct/{id}", middlewares.SetMiddlewareAuthentication(a.DeleteProduct)).Methods("DELETE")
}
func (a* App ) home(w http.ResponseWriter,req *http.Request){
	u.JSON(w,http.StatusOK,"Welcome")
}