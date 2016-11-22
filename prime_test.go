package main

import(
    "testing"
    "reflect"
    "net/http"
    "net/http/httptest"
    "io/ioutil"
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

// test the primeserver
func TestServer(t *testing.T) {
    server := httptest.NewServer(http.HandlerFunc(handler))
    defer server.Close()

    cases := []struct {
        param string
        status_code int
        response_body string
    }{
        {"?start_num=4&end_num=11", http.StatusOK, "[5,7,11]"},
        {"?start_num=-1&end_num=11", 200, "[1,2,3,5,7,11]"},
        {"?start_num=a&end_num=b", http.StatusBadRequest , "[]"},
        {"?start_num=1", http.StatusBadRequest, "[]"},
    }

    for _, c := range cases {
        url := server.URL + "/primes" + c.param

        resp, err := http.Get(url)
        if err != nil {
            t.Fatal(err)
        }

        // check status code
        if resp.StatusCode != c.status_code {
            t.Fatalf("TestServer, param(%s), expected status_code %d got %d\n", c.param, c.status_code,
             resp.StatusCode)
        }

        // check response body if the request is valid
        if http.StatusBadRequest == c.status_code {
            continue
        }
        actual, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            t.Fatal(err)
        }
        if c.response_body != string(actual) {
            t.Fatalf("TestServer, param(%s), expected response body '%s' got '%s'\n", c.param, c.response_body,
             string(actual))
        }
    }
}

