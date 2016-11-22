# primeserver
This is an HTTP endpoint(“/primes”) that accepts a “GET” request with two URL parameters (start_num and end_num). 
start_num indicates the starting number (inclusive) to start looking for primes 
end_num indicates the ending number (inclusive) to stop looking for primes 
The server calculates and return the prime numbers between start_num and end_num. 
The response is an JSON array of integers. 

## Example  
request: curl –XGET localhost:8080/primes?start_num=4&end_num=11  
response(body): [5, 7, 11]  

## HTTP status codes:  
400 bad request, the parameter is missing or invalid  
404 url not found  
200 OK  

## How to Run:  

    # the current workding directory is primieserver  
    go run primeserver.go  
    curl –XGET localhost:8080/primes?start_num=4&end_num=11  

## How to Test:

    # the current workding directory is primieserver  
    go test *.go  


