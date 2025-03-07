package main

import (
	"context"
	"golang-project/pkg/api"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gorm.io/gorm"
)

func main() {
	log.Println("starting...")
	initTimeZone()
	// configuration := config.New()

	// dsn := configuration.ConnectionString
	// dbSqlServer, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
	// 	Logger: logger.Default.LogMode(logger.Error),
	// })
	// if err != nil {
	// 	log.Println("can't connect to db..")
	// 	return

	// }

	// sqlDb, err := dbSqlServer.DB()
	// if err != nil {
	// 	log.Fatalln(err)
	// 	return
	// }

	// sqlDb.SetConnMaxLifetime(5 * time.Minute)
	// log.Println("db connected successfully..")

	// defer sqlDb.Close()

	// dbSqlServer = setGormUTCTime(dbSqlServer)
	// entity.Migration(dbSqlServer)

	apiMux := api.APIMux(api.APIConfig{DB: nil})

	srv := &http.Server{
		Addr:         "0.0.0.0:5000",
		WriteTimeout: time.Minute * 10,
		ReadTimeout:  time.Minute * 10,
		IdleTimeout:  time.Minute * 10,
		Handler:      apiMux,
	}

	serverError := make(chan error, 1)

	go func() {
		serverError <- srv.ListenAndServe()
	}()

	log.Println("Run server on port 5000")
	log.Println("The service is ready to listen and serve.")

	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-serverError:
		log.Fatal(err)
	case sig := <-gracefulStop:
		log.Println("shutdown", "status", "shutdown started", "signal", sig)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		_ = srv.Shutdown(ctx)
		defer cancel()
	}

	log.Println("The service is shutting down...")

	log.Println("terminated...")

	os.Exit(0)
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}

	time.Local = ict
}

func setGormUTCTime(input *gorm.DB) *gorm.DB {
	return input.Session(&gorm.Session{NowFunc: func() time.Time { return time.Now().UTC() }})
}
