# Word Counter

### What it does?
Count the words in given text (thru rest api) and sorted by highest to lowest (top ten)

### How to run?

1. `go run main.go`
2. make http request `POST http://localhost:8090/count` with json body `{"text" : <YOUR TEXT TO COUNT>}`