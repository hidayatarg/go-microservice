### full form 
```go
func GetUser(resp http.ResponseWriter, req *http.Request){
	userIdParam := req.URL.Query().Get("user_id")
	userId, err := strconv.ParseInt(userIdParam, 10, 10)
	if err != nil {
		
	}
}

// can be written
userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 10)
if err != nil {
// return bad request
return
}
```