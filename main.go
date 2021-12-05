// Package matrix provides a server for matrix operations.
package matrix

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"strings"
)

var (
	empty_matrix = fmt.Errorf("matrix is empty")
	not_square = fmt.Errorf("matrix is not square")
	not_numeric = fmt.Errorf("matrix contains non-numeric values")
)

// validateMatrix checks if the matrix is valid.
func validateMatrix(matrix [][]string) error {
	// Check if the matrix is empty
	if len(matrix) == 0 {
		return empty_matrix
	}

	// check if matrix is a square matrix
	if len(matrix) != len(matrix[0]) {
		return not_square
	}

	// check if matrix have only numbers
	for _, row := range matrix {
		for _, col := range row {
			_, err := strconv.Atoi(col)
			if err != nil {
				return not_numeric
			}
		}
	}

	return nil
}

// getCsvMatrix returns a matrix from the csv file in the request body.
func getCsvMatrix(r *http.Request) ([][]string, error) {
	file, _, err := r.FormFile("file")
	if err != nil {		
		return nil, fmt.Errorf("Could not open the matrix. Did you upload a valid csv file?")
	}
	defer file.Close()
	matrix, err := csv.NewReader(file).ReadAll()
	if err != nil {		
		return nil, fmt.Errorf("not a valid matrix")
	}
	
	return matrix, validateMatrix(matrix)
}

// echo returns the matrix as a string in matrix format.
func echo(matrix [][]string) (response string) {
	for _, row := range matrix {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	return response
}

// invert rows and colums of the matrix.
func invert(matrix [][]string) (response string) {
	newMatrix := make([][]string, len(matrix))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			newMatrix[j] = append(newMatrix[j], matrix[i][j])
		}
	}

	return echo(newMatrix)
}

// flatten the matrix as a 1 line string, with values separated by commas.
func flatten(matrix [][]string) (response string) {
	for _, row := range matrix {
		response = fmt.Sprintf("%s%s,", response, strings.Join(row, ","))
	}
	
	return strings.TrimSuffix(response, ",")
}

// sum all integers in the matrix.
func sum(matrix [][]string) string {
	var result int64

	for _, row := range matrix {
		for _, col := range row {
			i, _ := strconv.ParseInt(col, 10, 64)
			result += i
		} 
	}
	return strconv.FormatInt(result, 10)
}

// multiply returns the product of the matrix.
func multiply(matrix [][]string) string {
	result := big.NewInt(1)

	for _, row := range matrix {
		for _, col := range row {
			i, _ := strconv.ParseInt(col, 10, 64)
			if i == 0 {
				return "0"
			}
			result = result.Mul(result, big.NewInt(i))
		} 
	}
	return result.Text(10)
}

// Server handles the request and returns the response.
func Server(w http.ResponseWriter, r *http.Request) {
	matrix, err := getCsvMatrix(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("error: %s", err.Error())))
		return
	}

	var response string
	switch r.URL.Path {
	case "/echo":
		response = echo(matrix)
	case "/invert":
		response = invert(matrix)
	case "/flatten":
		response = flatten(matrix)
	case "/sum":
		response = sum(matrix)
	case "/multiply":
		response = multiply(matrix)
	default:
		w.WriteHeader(http.StatusNotFound)
		response = "not found"
	}

	fmt.Fprint(w, response)
}

// main is the entry point of the program.
func main() {
    log.Print("Starting server on localhost:5000") 
	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(Server)))
}