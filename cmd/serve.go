package cmd

import (
	"log"
	"context"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/99designs/gqlgen/handler"


	"go.opencensus.io/exporter/zipkin"
	"go.opencensus.io/trace"

	openzipkin "github.com/openzipkin/zipkin-go"
	zipkinHTTP "github.com/openzipkin/zipkin-go/reporter/http"
	gateway "github.com/nirajgeorgian/gateway/src/gateway"
)

// Port :- port to listen application on
var Port int

// AccountURI :- used for starting application with one default hash value
var AccountURI string

// JobURI :- database uri connect to
var JobURI string

// MailURI
var MailURI string

// LocalENdpoint :- local endpoint
var LocalEndpoint string

// ZipkinEndpoint string
var ZipkinEndpoint string

func init() {
	serveCmd.Flags().IntVarP(&Port, "port", "p", 8080, "port configuration for this application")
	serveCmd.Flags().StringVarP(&JobURI, "joburi", "j", "localhost:3000", "URI for job service (required)")
	serveCmd.Flags().StringVarP(&AccountURI, "accounturi", "a", "localhost:3001", "URI for account service (required)")
	serveCmd.Flags().StringVarP(&MailURI, "mailuri", "m", "127.0.0.1:3002", "URI for mail service (required)")
	serveCmd.Flags().StringVarP(&LocalEndpoint, "localendpoint", "u", "", "local endopoint URL")
	serveCmd.Flags().StringVarP(&ZipkinEndpoint, "zipkinendpoint", "z", "", "zipkin endopoint URL")


	viper.BindPFlag("port", serveCmd.Flags().Lookup("port"))
	viper.BindPFlag("accounturi", serveCmd.Flags().Lookup("accounturi"))
	viper.BindPFlag("joburi", serveCmd.Flags().Lookup("joburi"))
	viper.BindPFlag("mailuri", serveCmd.Flags().Lookup("mailuri"))
	viper.BindPFlag("localendpoint", serveCmd.Flags().Lookup("localendpoint"))
	viper.BindPFlag("zipkinendpoint", serveCmd.Flags().Lookup("zipkinendpoint"))
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serves the graphql frontend gateway server",
	Long:  `start the frontend gateway server on provided port along with the provided services URI`,
	RunE: func(cmd *cobra.Command, args []string) error {
		port := viper.GetString("port")
		localURL := viper.GetString("localendpoint")
		zipkinURL := viper.GetString("zipkinendpoint")

		localEndpoint, err := openzipkin.NewEndpoint("gateway-svc", localURL)
		if err != nil {
			log.Fatalf("Failed to create the local zipkinEndpoint: %v", err)
		}

		reporter := zipkinHTTP.NewReporter(zipkinURL + "/api/v2/spans")
		ze := zipkin.NewExporter(reporter, localEndpoint)
		trace.RegisterExporter(ze)

		trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})

		ctx, span := trace.StartSpan(context.Background(), "main")
		defer span.End()

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
