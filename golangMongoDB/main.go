package main

import (
	"fmt"
	"net/http"

	"github.com/Nahin-CDR/golangMongoDB/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	route := httprouter.New()

	uc := controllers.NewUserController(getSession())

	route.GET("/user/:id", uc.GetUser)
	route.POST("/user", uc.CreateUser)
	route.DELETE("/user/:id", uc.DeleteUser)

	http.ListenAndServe("localhost:9000", route)

}

func getSession() *mgo.Session {

	s, err := mgo.Dial("mongodb://user:pass@localhost:27021/")

	if err != nil {
		println("Error Occured")
		panic(err)
	}
	fmt.Println("DB connected !!! ")

	return s
}
