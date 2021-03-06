package main

type errorJson struct {
    Error string            `json:"error"`
}

type movie struct {
    IMDBId string           `json:"IMDBId"`
    Name string             `json:"Name"`
    Year string             `json:"Year"`
    Score string            `json:"Score"`
    Description string      `json:"Description"`
    Poster string           `json:"Poster"`
}

type rating struct{
    Source string
    Value string
}

type omdbMovie struct {
    Title string          `json:"Title"`
    Year string           `json:"Year"`
    Rated string          `json:"Rated"`
    Released string       `json:"Released"`
    Runtime string        `json:"Runtime"`
    Genre string          `json:"Genre"`
    Director string       `json:"Director"`
    Writer string         `json:"Writer"`
    Actors string         `json:"Actors"`
    Plot string           `json:"Plot"`
    Language string       `json:"Language"`
    Country string        `json:"Country"`
    Awards string         `json:"Awards"`
    Poster string         `json:"Poster"`
    Ratings []rating      `json:"Ratings"`
    Metascore string      `json:"Metascore"`
    ImdbRating string     `json:"ImdbRating"`
    imdbVotes string
    ImdbID string         `json:"ImdbID"`
    Type string           `json:"Type"`
    DVD string            `json:"DVD"`
    BoxOffice string      `json:"BoxOffice"`
    Production string     `json:"Production"`
    Website string        `json:"Website"`
    Response bool       `json:"Response"`
}
