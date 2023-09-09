package main

import (
	"log"
	"net/http"

	_ "net/http/pprof"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/a-agmon/gql-parquet-api/graph"
	"github.com/a-agmon/gql-parquet-api/pkg/aws"
	"github.com/a-agmon/gql-parquet-api/pkg/data"
)

const defaultPort = "8999"

func main() {
	log.Printf("starting profiling endpoint on port 6060")
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	log.Print("starting server...")
	awsCred, err := aws.GetAWSCredEnv()
	if err != nil {
		log.Fatalf("error getting aws credentials from env: %v", err)
	}
	dataDriver := data.NewDuckDBDriver(awsCred)
	dataStore := data.NewStore(dataDriver)
	resolver := graph.NewResolver(dataStore)
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, nil))
}
