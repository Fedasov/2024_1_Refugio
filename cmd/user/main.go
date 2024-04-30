package main

import (
	"database/sql"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"mail/cmd/configs"
	"net"
	"os"
	"time"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"

	"mail/internal/microservice/interceptors"
	"mail/internal/microservice/user/proto"
	userRepo "mail/internal/microservice/user/repository"
	grpcUser "mail/internal/microservice/user/server"
	userUc "mail/internal/microservice/user/usecase"
)

func main() {
	settingTime()

	db := initializeDatabase()
	defer db.Close()

	userGrpc := initializeUser(db)

	loggerInterceptorAccess := initializationInterceptorLogger()

	startServer(userGrpc, loggerInterceptorAccess)
}

// settingTime setting local time on server
func settingTime() {
	loc, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		fmt.Println("Error in location detection")
	}

	time.Local = loc
}

// initializeDatabase database initialization
func initializeDatabase() *sql.DB {
	db, err := sql.Open("pgx", configs.DSN)
	if err != nil {
		log.Fatalln("Can't parse config", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln("Database is not available", err)
	}

	db.SetMaxOpenConns(10)

	return db
}

// initializeUser initializing user server
func initializeUser(db *sql.DB) *grpcUser.UserServer {
	userRepository := userRepo.NewUserRepository(sqlx.NewDb(db, "pgx"))
	userUseCase := userUc.NewUserUseCase(userRepository)

	return grpcUser.NewUserServer(userUseCase)
}

// initializationInterceptorLogger initializing logger
func initializationInterceptorLogger() *interceptors.Logger {
	f, err := os.OpenFile("logInterEmail.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Failed to create logfile" + "log.txt")
	}

	logrusAccess := interceptors.InitializationAccessLogInterceptor(f)
	loggerAccess := new(interceptors.Logger)
	loggerAccess.Logger = logrusAccess

	return loggerAccess
}

// startServer starting server
func startServer(userGrpc *grpcUser.UserServer, interceptorsLogger *interceptors.Logger) {
	listen, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatalf("Cannot listen port: %s. Err: %s", "8001", err.Error())
	}

	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			interceptorsLogger.AccessLogInterceptor,
			interceptors.PanicRecoveryInterceptor,
		),
	}
	grpcServer := grpc.NewServer(opts...)

	proto.RegisterUserServiceServer(grpcServer, userGrpc)

	fmt.Printf("The server is running in port 8001\n")

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("Cannot listen port: %s. Err: %s", "8001", err.Error())
	}
}
