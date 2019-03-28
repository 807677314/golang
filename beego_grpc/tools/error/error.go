package error


type Result struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}


var SUCCESS = Result{200, "success"}
