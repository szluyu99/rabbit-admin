package main

import (
	"flag"
	"io"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/restsend/rabbit"
	rabbitadmin "github.com/szluyu99/rabbit-admin/internal"
)

var fileMode os.FileMode = 0666

var (
	serverAddr string
	logFile    string
	dbDriver   string
	dsn        string
)

func main() {
	flag.StringVar(&serverAddr, "s", ":8765", "listen addr")
	flag.StringVar(&logFile, "l", "", "log file")
	flag.StringVar(&dbDriver, "d", "", "DB Driver, sqlite|mysql")
	flag.StringVar(&dsn, "n", "", "DB DSN")
	flag.Parse()

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
