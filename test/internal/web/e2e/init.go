package e2e

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/karmek-k/basic-service-go/internal/web"
)

type TestServer struct {
	server *httptest.Server
}

type GenericJSON map[string]interface{}

func InitTestServer() *TestServer {
	return &TestServer{
		server: httptest.NewServer(web.CreateCalculatorHandler()),
	}
}

func (ts *TestServer) Close() {
	ts.server.Close()
}

func (ts *TestServer) TestGet(t *testing.T, uri string) *GenericJSON {
	res, err := http.Get(ts.server.URL + uri)
	if err != nil {
		t.Fatal(err)
	}

	defer res.Body.Close()

	return ts.TestResponse(t, res)
}

func (ts *TestServer) TestResponse(t *testing.T, res *http.Response) *GenericJSON {
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected code %v, got %v", http.StatusOK, res.Status)
	}

	data, err := io.ReadAll(res.Body)

	if err != nil {
		t.Fatal(err)
	}

	var parsed GenericJSON
	err = json.Unmarshal(data, &parsed)

	if err != nil {
		t.Fatal("unmarshaling data failed")
	}

	return &parsed
}
