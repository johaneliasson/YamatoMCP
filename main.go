package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"flag"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

var (
	// YamatoBearerToken is the bearer token for the Yamato API, read from a file
	YamatoBearerToken = GetYamatoBearerToken()
	httpAddr          = flag.String("http", "", "use HTTP at this address")
)

func main() {
	// Create a server with a single tool.
	server := mcp.NewServer(&mcp.Implementation{Name: "yamato-mcp", Version: "v1.0.0"}, nil)

	mcp.AddTool(server, &mcp.Tool{Name: "yamato_job_id", Description: "Fetch data for a given Yamato Job ID"}, GetYamatoDataForJobID)

	mcp.AddTool(server, &mcp.Tool{Name: "yamato_job_definition", Description: "Provides build history for given Yamato job definition"}, GetYamatoJobDefinitionHistory)

	flag.Parse()

	if *httpAddr != "" {
		log.Printf("Starting HTTP server on %s", *httpAddr)

		// Create a single transport instance that handles all connections
		sessionID := generateSessionID()
		transport := mcp.StreamableServerTransport{
			SessionID: sessionID,
		}

		// Start the server in a goroutine
		go func() {
			if err := server.Run(context.Background(), &transport); err != nil {
				log.Printf("Server error: %v", err)
			}
		}()

		// Register the HTTP handler
		http.Handle("/mcp", &transport)
		log.Fatal(http.ListenAndServe(*httpAddr, nil))
	} else {
		log.Printf("Starting MCP server on stdin/stdout")
		if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
			log.Fatal(err)
		}
	}
}

func GetYamatoBearerToken() string {
	refreshCmd := exec.Command("yamato", "auth", "show")
	err := refreshCmd.Run()
	if err != nil {
		return ""
	}
	token, err := os.ReadFile(os.ExpandEnv("$HOME/.yamato/token"))
	if err != nil {
		log.Fatalf("Error reading auth token: %v", err)
	}
	return string(token)

}

type YamatoJobParams struct {
	YamatoJobId string `json:"YamatoJobId" jsonschema:"the Yamato Job ID to get data for, e.g. 51154833"`
}

type YamatoJobDefinitionParams struct {
	JobDefinition string `json:"jobDefinition" jsonschema:"the jobdefinition to get history for, e.g. .yamato/utr.yml#build_utr_win"`
	PageSize      string `json:"pageSize" jsonschema:"the number of results to return per page, default is 50"`
	Project       string `json:"project" jsonschema:"the project ID to filter by, default is 3"`
	Status        string `json:"status" jsonschema:"the status to filter by, default is 'completed'"`
}

func generateSessionID() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func GetYamatoJobDefinitionHistory(ctx context.Context, req *mcp.CallToolRequest, args YamatoJobDefinitionParams) (*mcp.CallToolResult, any, error) {
	// Get History of Yamato jobs using the Yamato API for a given job definition
	client := http.Client{Timeout: 5 * time.Second}
	//url_yamato_history := "https://yamato-api.cds.internal.unity3d.com/jobs" + url.QueryEscape("pageSize=50&filter=project eq 3 and filename eq '"+params.Arguments.JobDefinition+"'") //ie .yamato/utr.yml#build_utr_win
	//url_yamato_history := "https://yamato-api.cds.internal.unity3d.com/jobs" + url.QueryEscape("pageSize=50&filter=project eq 3 and filename eq '.yamato/utr.yml#build_utr_win'") //ie .yamato/utr.yml#build_utr_win

	u, _ := url.Parse("https://yamato-api.cds.internal.unity3d.com/jobs")
	q := u.Query()
	q.Set("pageSize", "50")
	q.Set("filter", "project eq 3 and filename eq '"+args.JobDefinition+"'") //ie .yamato/utr.yml#build_utr_win
	u.RawQuery = q.Encode()

	l := log.New(os.Stderr, "", 0)
	l.Println("url:" + u.String())

	request, err := http.NewRequest(http.MethodGet, u.String(), nil)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+YamatoBearerToken)
	response, err := client.Do(request)
	if err != nil {
		return nil, nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Error closing response body: %v", err)
		}
	}(response.Body)

	// Unmarshal the response body into JSON
	var responseBody map[string]any
	err = json.NewDecoder(response.Body).Decode(&responseBody)

	YamatoJobDefinitionHistory, err := json.Marshal(responseBody)

	return &mcp.CallToolResult{
		Content: []mcp.Content{&mcp.TextContent{Text: string(YamatoJobDefinitionHistory)}},
	}, nil, err
}

func GetYamatoDataForJobID(ctx context.Context, req *mcp.CallToolRequest, args YamatoJobDefinitionParams) (*mcp.CallToolResult, any, error) {

	// Get Yamato job data using the Yamato API for a given Job ID
	client := http.Client{Timeout: 5 * time.Second}
	url := "https://yamato-api.cds.internal.unity3d.com/jobs/" + args.JobDefinition
	request, err := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+YamatoBearerToken)

	response, err := client.Do(request)
	if err != nil {
		return nil, nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Error closing response body: %v", err)
		}
	}(response.Body)

	// Unmarshal the response body into JSON
	var responseBody map[string]any
	err = json.NewDecoder(response.Body).Decode(&responseBody)

	YamatoJobInfoJSON, err := json.Marshal(responseBody)

	return &mcp.CallToolResult{
		Content: []mcp.Content{&mcp.TextContent{Text: string(YamatoJobInfoJSON)}},
	}, nil, err
}
