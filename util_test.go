package matrix_test

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"testing"
)

const (
	valid_matrix = "1,2,3\n4,5,6\n7,8,9"
	invalid_matrix = "1,2,3,4\n8,9\n10,11,12"
	empty_matrix = ""
	not_square = "1,2,3\n4,5,6\n7,8,9\n10,11,12"
	not_number = "x,2,3\n4,5,6\n7,8,a"
	matrix_with_zero = "0,1,2\n3,4,5\n6,7,8"
	matrix_with_big_numbers = "1091555,1091556,1091557\n1292665,1292666,1292667\n1393775,1393776,1393777"
)

// createMultipartFormData creates a multipart form data with a given matrix
func createMultipartFormData(matrix string) (bytes.Buffer, *multipart.Writer) {
	var buffer bytes.Buffer	
	writer := multipart.NewWriter(&buffer)
	var fw io.Writer
	fw, _ = writer.CreateFormFile("file", "file"); 
	fw.Write([]byte(matrix))
	writer.Close()
	return buffer, writer
}

// newRequest returns a new request with a given path and a given matrix
func newRequest(path string, matrix string) *http.Request {
	body, writer := createMultipartFormData(matrix)
	request, _ := http.NewRequest(http.MethodPost, path, &body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	return request
}

// assertResponseBody asserts that a response's body is equal to the expected body
func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}

// assertResponseCode asserts that a response has a specific status code
func assertStatus(t testing.TB, got, want int) {
    t.Helper()
    if got != want {
        t.Errorf("did not get correct status, got %d, want %d", got, want)
    }
}