package cmd

import (
	"log"
	"time"
	"context"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/99designs/gqlgen/handler"

	gateway "github.com/nirajgeorgian/gateway/src/gateway"
)

// Port :- port to listen application on
var Port int

// AccountURI :- used for starting application with one default hash value
var AccountURI string

// JobURI :- database uri connect to
var JobURI string

func init() {
	serveCmd.Flags().IntVarP(&Port, "port", "p", 8080, "port configuration for this application")
	serveCmd.Flags().StringVarP(&AccountURI, "accounturi", "a", "localhost:3001", "URI for account service (required)")
	serveCmd.Flags().StringVarP(&JobURI, "joburi", "j", "localhost:3000", "URI for job service (required)")

	// serveCmd.MarkFlagRequired("secretkey")
	// serveCmd.MarkFlagRequired("databaseuri")

	viper.BindPFlag("port", serveCmd.Flags().Lookup("port"))
	viper.BindPFlag("accounturi", serveCmd.Flags().Lookup("accounturi"))
	viper.BindPFlag("joburi", serveCmd.Flags().Lookup("joburi"))
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serves the graphql frontend gateway server",
	Long:  `start the frontend gateway server on provided port along with the provided services URI`,
	RunE: func(cmd *cobra.Command, args []string) error {
		port := viper.GetString("port")

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		s, err := gateway.NewGraphQLServer(ctx)
		if err != nil {
			log.Fatal(err)
		}

		http.Handle("/graphql", handler.GraphQL(s.ToExecutableSchema()))
		http.Handle("/playground", handler.Playground("gateway", "/graphql"))

		log.Fatal(http.ListenAndServe(":"+port, nil))

		return nil
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
}
