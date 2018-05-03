package main


type KongApi struct {
	Name          	    string  `json:"name"`	
	RequestPath         string  `json:"request_path"`  	
	StripRequestPath 	string	`json:"strip_request_path"`    
	UpstreamUrl 	    string	`json:"upstream_url"`
}
