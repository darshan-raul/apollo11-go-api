package main

import (
  "fmt"
  "strings"
  "net/http"
  "io/ioutil"
)



func make_graphql_call() {

  url := "http://127.0.0.1:8000/graphql"
  method := "POST"

  payload := strings.NewReader(`{"query":"{student { name city}}"}`)

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("Content-Type", "application/json")

  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(body))

}


func main() {

  make_graphql_call()
}
