package main

type Post struct {
  Title   string  `json:"title"`
  Id      int     `json:"id"`
}

type Posts struct {
  PostArray []Post  `json:"posts"`
}

type Action struct {
  Type    string `json:"type"`
  Payload string `json:"payload"`
}
