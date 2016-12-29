package main

type Post struct {
  Title   string  `json:"title"`
  Id      int     `json:"id"`
}

type Posts struct {
  PostArray []Post  `json:"posts"`
}

type Var struct {
  Key   string  `json:"key"`
  Value string  `json:"value"`
}

type Vars struct {
  VarArray  []Var   `json:"vars"`
}
