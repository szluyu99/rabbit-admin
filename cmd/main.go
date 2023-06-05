package main

import (
	"flag"
	"io"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/szluyu99/rabbit"
	rabbitadmin "github.com/szluyu99/rabbit-admin/internal"
	"gorm.io/gorm"
)

var fileMode os.FileMode = 0666

var (
	serverAddr string
	logFile    string
	dbDriver   string
	dsn        string

	superUserEmail    string
	superUserPassword string
)

func main() {
	flag.StringVar(&serverAddr, "s", ":8765", "listen addr")
	flag.StringVar(&logFile, "l", "", "log file")
	flag.StringVar(&dbDriver, "d", "", "DB Driver, sqlite|mysql")
	flag.StringVar(&dsn, "n", "", "DB DSN")
	flag.StringVar(&superUserEmail, "superuser", "", "Create an super user with email")
	flag.StringVar(&superUserPassword, "password", "", "Super user password")
	flag.Parse()

	// log flag
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var err error
	var lw io.Writer
	if logFile != "" {
		lw, err = os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, fileMode)
		if err != nil {
			log.Fatalf("open %s fail, %v\n", logFile, err)
		}
	}

	// db
	db := rabbit.InitDatabase(dbDriver, dsn, lw)

	// create super user
	if superUserEmail != "" && superUserPassword != "" {
		createSuperUser(db, superUserEmail, superUserPassword)
	}

	// router
	r := gin.New()
	r.Use(gin.LoggerWithWriter(lw), gin.Recovery())

	// rabbit
	rabbit.InitRabbit(db, r)

	// rabbitadmin
	rabbitadmin.InitServer(db, r)

	if serverAddr[0] == ':' || strings.HasPrefix(serverAddr, "0.0.0.0:") {
		log.Printf("Serving HTTP on (http://localhost:%s/) ... \n", strings.Split(serverAddr, ":")[1])
	} else {
		log.Printf("Serving HTTP on (http://%s/) ... \n", serverAddr)
	}

	r.Run(serverAddr)
}

func createSuperUser(db *gorm.DB, email, password string) {
	u, err := rabbit.GetUserByEmail(db, email)
	if err == nil && u != nil {
		rabbit.SetPassword(db, u, password)
		rabbit.Warningln("Super user password changed")
	} else {
		u, err := rabbit.CreateUser(db, email, password)
		if err != nil {
			panic(err)
		}
		u.Activated = true
		u.Enabled = true
		u.IsSuperUser = true
		if err := db.Save(u).Error; err != nil {
			panic(err)
		}
		rabbit.Warningln("Create super user:", email)
		os.Exit(0)
		return
	}
}
