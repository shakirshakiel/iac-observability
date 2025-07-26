package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"go.elastic.co/apm/v2"
)

func main() {
	// Start a transaction for the entire main function execution
	// os.Setenv("ELASTIC_APM_ENVIRONMENT", "development")
	// os.Setenv("ELASTIC_APM_TRANSACTION_SAMPLE_RATE", "0.0")
	tx := apm.DefaultTracer().StartTransaction("observability2.0", "request")

	// Simulate some work
	fmt.Println("Starting custom operation...")
	time.Sleep(100 * time.Millisecond) // Simulate some processing

	// Create a span for a sub-operation
	span := tx.StartSpan("example_resource", "create", nil)
	time.Sleep(50 * time.Millisecond) // Simulate sub-task work

	req, err := http.NewRequest("GET", "http://localhost:7071/api/health", nil)
	req.Header.Set("traceparent", fmt.Sprintf("00-%s-%s-00", tx.TraceContext().Trace.String(), tx.TraceContext().Span.String()))
	req.Header.Set("tracestate", tx.TraceContext().State.String())
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	span.End()

	// // Simulate an error
	// err := fmt.Errorf("simulated error during main execution")
	// apm.CaptureError(context.Background(), err).Send()

	// Flush pending transactions for short-lived processes
	tx.Result = "success"
	tx.Context.SetLabel("tx_resource_id", "123")
	tx.End()
	apm.DefaultTracer().Flush(nil)
}
