package clients_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"workerpool-application/internal/clients"

	"github.com/stretchr/testify/assert"
)

func TestDoReq(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`hello this is server responding`))
	}))
	defer ts.Close()
	fmt.Println(ts.URL)
	client := clients.NewClient(http.DefaultClient)
	resp, err := client.DoReq(ts.URL)
	assert.NoError(t, err)
	assert.Equal(t, string(resp), "hello this is server responding")
}
