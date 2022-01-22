package main

import (
  "fmt"
  "strings"
  "net/http"
  "io/ioutil"
  "github.com/gofiber/fiber/v2"
  "log"
)


func main() {

  app := fiber.New()


  app.Get("/graphql", func(c *fiber.Ctx) error {
      msg := fmt.Sprintf("✋ %s", c.Params("*"))
      url := "http://127.0.0.1:8000/graphql"
      method := "POST"

      payload := strings.NewReader(`{"query":"{student { name city}}"}`)

      client := &http.Client {
      }
      req, err := http.NewRequest(method, url, payload)

      if err != nil {
        fmt.Println(err)
        return err
      }
      req.Header.Add("Content-Type", "application/json")

      res, err := client.Do(req)
      if err != nil {
        fmt.Println(err)
        return err
      }
      defer res.Body.Close()

      body, err := ioutil.ReadAll(res.Body)
      if err != nil {
        fmt.Println(err)
        return err
      }
      
      fmt.Println(string(body)+msg)
      return c.SendString(string(body)) // => ✋ register
  })


  log.Fatal(app.Listen(":3000"))

  // make_graphql_call()
}
