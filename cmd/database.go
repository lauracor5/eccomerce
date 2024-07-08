package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
)

const AppName = "EDcommerce"

func newDBConnection() (*pgxpool.Pool, error) {
	minDefault := 3
	maxDefault := 100
	minDefaultFinal := 0
	maxDefaultFinal := 0

	minConn := os.Getenv("DB_MIN_CONN")
	maxConn := os.Getenv("DB_MAX_CONN")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("DB_SSL_MODE")

	if minConn != "" {
		minConnectionNumber, err := strconv.Atoi(minConn) //parseo de string a entero
		if err != nil {
			log.Println("warning DB_MIN_CONN has not valid, we will set min connections to", minDefault)
		} else {
			if minConnectionNumber >= minDefault && minConnectionNumber <= maxDefault {
				minDefaultFinal = minConnectionNumber
			}

		}

	}

	if maxConn != "" {
		maxConnectionNumber, err := strconv.Atoi(maxConn)
		if err != nil {
			log.Println("warning DB_MAX_CONN has not valid, we will set max connections to", minDefault)
		} else {
			if maxConnectionNumber >= minDefault && maxConnectionNumber <= maxDefault {
				maxDefaultFinal = maxConnectionNumber
			}
		}
	}

	if minDefaultFinal > maxDefaultFinal {
		minDefaultFinal, _ = strconv.Atoi(minConn)

	}

	urlConnection := makeDns(user, pass, host, port, dbName, sslMode, minDefaultFinal, maxDefaultFinal)
	config, err := pgxpool.ParseConfig(urlConnection)
	if err != nil {
		return nil, fmt.Errorf("%s %w", "pgxpool.ParseConfig()", err)
	}
	config.ConnConfig.RuntimeParams["application_name"] = AppName

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("%s %w", "pgxpool.NewWithConfig()", err)
	}

	return pool, nil

}

func makeDns(user, pass, host, port, dbName, sslMode string, minConnection, maxConnection int) string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s pool_min_conns=%d pool_max_conns=%d",
		user,
		pass,
		host,
		port,
		dbName,
		sslMode,
		minConnection,
		maxConnection)
}
