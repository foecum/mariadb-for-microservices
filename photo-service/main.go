package main

import (
	"net/http"
	"strconv"

	"github.com/bstaijen/mariadb-for-microservices/photo-service/app/http/routes"
	"github.com/bstaijen/mariadb-for-microservices/photo-service/config"
	"github.com/bstaijen/mariadb-for-microservices/photo-service/database"
	negronilogrus "github.com/meatballhat/negroni-logrus"
	"github.com/urfave/negroni"

	// go-sql-driver/mysql is needed for the database connection
	_ "github.com/go-sql-driver/mysql"

	_ "github.com/joho/godotenv/autoload"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetLevel(log.DebugLevel)

	// Get config
	cnf := config.LoadConfig()

	// Get database // TODO : wait for the DB to go online for max 1 min?
	connection, err := db.OpenConnection(cnf)
	if err != nil {
		log.Fatal(err)
	}
	defer db.CloseConnection(connection)

	// Set the REST API routes
	r := routes.InitRoutes(connection, cnf)
	n := negroni.Classic()
	n.Use(negronilogrus.NewMiddleware())
	n.UseHandler(r)

	// Start and listen on port in cbf.Port
	log.Info("Starting server on port " + strconv.Itoa(cnf.Port))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(cnf.Port), n))
}
