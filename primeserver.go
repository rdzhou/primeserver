package main

import (
    "fmt"
)

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


func main() {
    fmt.Printf("%v\n", getPrimes(1, 100))
    fmt.Printf("%v\n", getPrimes(100, 1))
}

