package main

import (
    "net/http"
    "strconv"
    "encoding/json"
)

// Test if the input n is prime
// return true if n is prime
func isPrime(n int) bool {
    if n <= 0 {
        return false
    }
    for i:=2; i<=n/2; i++ {
        if n % i == 0 {
            return false
        }
    }
    return true
}

// get the prime numbers in [start, end]
// start and end are both included
// if start is not positive, it will be adjusted to 1
// if end is not positive, an empty slice will be returned
func getPrimes(start, end int) []int {
    ret := []int{}
    if start <= 0 {
        start = 1
    }
    if end <= 0 {
        return ret
    }
    for ; start <= end; start++ {
        if isPrime(start) {
            ret = append(ret, start)
        }
    }
    return ret
}

// handler function for the prime server
// if the parameters are valid, StatusOK and the json array of prime numbers will
// be sent
// otherwise, StatusBadRequest will be sent
func handler(w http.ResponseWriter, r *http.Request) {
    if len(r.URL.Query()["start_num"]) > 0 {
        if len(r.URL.Query()["end_num"]) > 0 {
            if start_num, err := strconv.Atoi(r.URL.Query()["start_num"][0]); err == nil {
                if end_num, err := strconv.Atoi(r.URL.Query()["end_num"][0]); err == nil {
                    if jData, err := json.Marshal(getPrimes(start_num, end_num)); err == nil {
                        w.Header().Set("Content-Type", "application/json")
                        w.Write(jData)
                        return                      
                    }
                }
            }
        }
    }
    w.WriteHeader(http.StatusBadRequest)
}

func main() {
    http.HandleFunc("/primes", handler)
    http.ListenAndServe(":8080", nil)
}

