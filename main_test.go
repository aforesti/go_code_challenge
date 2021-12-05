package matrix_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aforesti/codechallenge/matrix"
)

// TestServer tests the following:
// - returns 404 on invalid path
// - validates only numbers on the matrix
// - validates square matrix
// - validates empty matrix
// - validates invalid matrix
// - validates matrix was uploaded correctly
// - returns the same matrix when using /echo
// - returns the inverted matrix when using /invert
// - returns the flattened matrix when using /flatten
// - returns the sum of the matrix when using /sum
// - returns the product of the matrix when using /multiply
// - when multiplying by zero, always returns 0
func TestServer(t *testing.T) { 
	t.Run("Echo", func(t *testing.T) {		
		request := newRequest("/echo", valid_matrix)
		response := httptest.NewRecorder()

		matrix.Server(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "1,2,3\n4,5,6\n7,8,9\n")	
	})

	t.Run("Invert", func(t *testing.T) {
		request := newRequest("/invert", valid_matrix)
		response := httptest.NewRecorder()

		matrix.Server(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "1,4,7\n2,5,8\n3,6,9\n")		
	})	
	
	t.Run("Flatten", func(t *testing.T) {
		request := newRequest("/flatten", valid_matrix)
		response := httptest.NewRecorder()

		matrix.Server(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "1,2,3,4,5,6,7,8,9")		
	})	

	t.Run("Sum", func(t *testing.T) {
		request := newRequest("/sum", valid_matrix)
		response := httptest.NewRecorder()

		matrix.Server(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "45")		
	})

	t.Run("Multiply", func(t *testing.T) {
		request := newRequest("/multiply", valid_matrix)
		response := httptest.NewRecorder()

		matrix.Server(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "362880")		
	})

	t.Run("returns 404 on invalid path", func(t *testing.T) {
		request := newRequest("/invalid", valid_matrix)
		response := httptest.NewRecorder()

		matrix.Server(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)		
	})

	t.Run("validates only numbers on the matrix", func(t *testing.T) {
		request := newRequest("/echo", not_number)
		response := httptest.NewRecorder()

		matrix.Server(response, request)

		assertStatus(t, response.Code, http.StatusBadRequest)
		assertResponseBody(t, response.Body.String(), "error: matrix contains non-numeric values")
	})

	t.Run("validates square matrix", func(t *testing.T) {
		request := newRequest("/echo", not_square)
		response := httptest.NewRecorder()

		matrix.Server(response, request)

		assertStatus(t, response.Code, http.StatusBadRequest)
		assertResponseBody(t, response.Body.String(), "error: matrix is not square")
	})

	t.Run("validates empty matrix", func(t *testing.T) {
		request := newRequest("/echo", empty_matrix)
		response := httptest.NewRecorder()

		matrix.Server(response, request)

		assertStatus(t, response.Code, http.StatusBadRequest)
		assertResponseBody(t, response.Body.String(), "error: matrix is empty")
	})

	t.Run("validates invalid matrix", func(t *testing.T) {
		request := newRequest("/echo", invalid_matrix)
		response := httptest.NewRecorder()

		matrix.Server(response, request)

		assertStatus(t, response.Code, http.StatusBadRequest)
		assertResponseBody(t, response.Body.String(), "error: not a valid matrix")
	})

	t.Run("when multiplying by zero, always returns 0", func(t *testing.T) {
		request := newRequest("/multiply", matrix_with_zero)
		response := httptest.NewRecorder()

		matrix.Server(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "0")
	})

	t.Run("sum matrix with big numbers", func(t *testing.T) {
		request := newRequest("/sum", matrix_with_big_numbers)
		response := httptest.NewRecorder()

		matrix.Server(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "11333994")
	})

	t.Run("multiply matrix with big numbers", func(t *testing.T) {
		request := newRequest("/multiply", matrix_with_big_numbers)
		response := httptest.NewRecorder()

		matrix.Server(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "7606344435967868270093159864770981491376593066866640000")
	})

	t.Run("request without cvs file", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/echo", nil)
		response := httptest.NewRecorder()

		matrix.Server(response, request)

		assertStatus(t, response.Code, http.StatusBadRequest)
		assertResponseBody(t, response.Body.String(), "error: Could not open the matrix. Did you upload a valid csv file?")
	})

}

func benchmark(path string, b *testing.B) {
	request := newRequest(path, valid_matrix)
	response := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {		
		matrix.Server(response, request)
	}
}

func BenchmarkEcho(b *testing.B) { benchmark("/echo", b) }
func BenchmarkInvert(b *testing.B) { benchmark("/invert", b) }
func BenchmarkFlatten(b *testing.B) { benchmark("/flatten", b) }
func BenchmarkSum(b *testing.B) { benchmark("/sum", b) }
func BenchmarkMultiply(b *testing.B) { benchmark("/multiply", b) }
