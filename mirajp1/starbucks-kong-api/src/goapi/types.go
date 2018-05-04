package main


type KongApi struct {
	Name          	    string  `json:"name"`	
	Uris                string  `json:"uris"`  	
	StripUri 	        string	`json:"strip_uri"`    
	UpstreamUrl 	    string	`json:"upstream_url"`
}
