# League Backend Challenge

## Authors

- [André Foresti](https://www.github.com/aforesti)

## Run Locally

Unzip the package and go to the project directory

Run the following command to start the server

```bash
  go run .
```

## Running Unit Tests

To run tests, run the following command

```bash
  go test -cover
```

## Usage/Examples

Given an uploaded csv file

```
1,2,3
4,5,6
7,8,9
```

With the server runing you may call one of the following endpoint:

POST /echo

- Return the matrix as a string in matrix format.

```
1,2,3
4,5,6
7,8,9
```

POST /invert

- Return the matrix as a string in matrix format where the columns and rows are inverted

```
1,4,7
2,5,8
3,6,9
```

POST /flatten

- Return the matrix as a 1 line string, with values separated by commas.

```
1,2,3,4,5,6,7,8,9
```

POST /sum

- Return the sum of the integers in the matrix

```
45
```

POST /multiply

- Return the product of the integers in the matrix

```
362880
```

Example of a request:

```
curl -F 'file=@./matrix.csv' "localhost:5000/echo"
```

## Feedback

If you have any feedback, please reach out at aforest@gmail.com
# go_code_challenge
