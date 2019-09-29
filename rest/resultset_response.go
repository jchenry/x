package rest

//ResultSetMetadata -
type ResultSetMetadata struct {
	Count  int `json:"count"`
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

//Metadata -
type Metadata struct {
	ResultSet ResultSetMetadata `json:"resultset"`
}

//ResultSetResponse -
type ResultSetResponse struct {
	Metadata Metadata    `json:"metadata"`
	Results  interface{} `json:"results"`
}
