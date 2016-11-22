package main

import(
    "testing"
    "reflect"
)

// test func isPrime
func TestIsPrime(t *testing.T) {
    cases := []struct {
        in int
        want bool
    }{
        {1, true},
        {2, true},
        {17, true},
        {10, false},
        {-1, false},
    }
    for _, c := range cases {
        got := isPrime(c.in)
        if got != c.want {
            t.Errorf("isPrime(%d) == %t, want %t", c.in, got, c.want)
        }
    }
}

// test func getPrimes
func TestGetPrimes(t *testing.T) {
    cases := []struct {
        start, end int
        want []int
    }{
        {4, 11, []int{5, 7, 11}},
        {-1, 11, []int{1, 2, 3, 5, 7, 11}},
        {0, -1234, []int{}},
    }
    for _, c := range cases {
        got := getPrimes(c.start, c.end)
        if !reflect.DeepEqual(got, c.want) {
            t.Errorf("getPrime(%d, %d) == %v, want %v", c.start, c.end, got, c.want)
        }
    }
}

