package main

import (
	"UTS/controllers"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/rooms", controllers.GetAllRooms).Methods("GET")
	router.HandleFunc("/participants", controllers.InsertRoom).Methods("POST")
	router.HandleFunc("/detail_rooms", controllers.GetDetailRooms).Methods("GET")
	// router.HandleFunc("/v1/users", controllers.InsertUser).Methods("POST")
	// router.HandleFunc("/v2/users", controllers.InsertUserGORM).Methods("POST")
	// router.HandleFunc("/v1/users", controllers.UpdateUser).Methods("PUT")
	// router.HandleFunc("/v2/users", controllers.UpdateUserGORM).Methods("PUT")
	// router.HandleFunc("/v1/users", controllers.DeleteUser).Methods("DELETE")
	// router.HandleFunc("/v2/users", controllers.DeleteUserGORM).Methods("DELETE")
	// router.HandleFunc("/products", controllers.GetAllProducts).Methods("GET")
	// router.HandleFunc("/products", controllers.InsertProduct).Methods("POST")
	// router.HandleFunc("/products", controllers.UpdateProduct).Methods("PUT")
	// router.HandleFunc("/products", controllers.DeleteProduct).Methods("DELETE")
	// router.HandleFunc("/transactions", controllers.GetAllTransactions).Methods("GET")
	// router.HandleFunc("/transactions", controllers.InsertTransaction).Methods("POST")
	// router.HandleFunc("/transactions", controllers.UpdateTransaction).Methods("PUT")
	// router.HandleFunc("/transactions", controllers.DeleteTransaction).Methods("DELETE")
	// router.HandleFunc("/v1/user_transactions", controllers.GetUserTransactions).Methods("GET")
	// router.HandleFunc("/v2/user_transactions", controllers.GetUserTransactionsGORM).Methods("GET")
	// router.HandleFunc("/user_login", controllers.Login).Methods("POST")

	http.Handle("/", router)
	fmt.Println("Connected to port 8888")
	log.Println("Connected to port 8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}
