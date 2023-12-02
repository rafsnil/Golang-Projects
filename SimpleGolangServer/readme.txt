A simple Golang server with a few basic APIs

NEW THINGS LEARNED:

1) http.Handle VS http.handleFunc
* http.Handle works with an existing http.Handler implementation, 
while http.HandleFunc takes a function and internally converts it to an http.Handler implementation.

* http.Handle is slightly more flexible as it can handle any http.Handler 
implementation, while http.HandleFunc is limited to functions with the http.Handler signature.

* http.HandleFunc is generally more concise and easier to use for simple handlers, 
while http.Handle is more suitable for complex handlers or when you need to reuse an 
existing http.Handler implementation.

2) HANDLERS 
Any handler function usually have 2 parts.
                    response writer        request
func helloHandler(w http.ResponseWriter, r *http.Request)

* REQUEST is received from the FE 
* RESPONSE is what the API sends back to the FE after 
processing the request

3) HTTP ERROR
http.Error is a function in the net/http package of the Go standard library. 
It is used to generate an HTTP error response.

http.Error(w http.ResponseWriter, msg string, code int)
w: The http.ResponseWriter to write the error response to.
msg: The error message to send in the response body.
code: The HTTP status code to send in the response header.

http.Error(w, "Not Found", http.StatusNotFound)
This code will send an HTTP 404 (Not Found) error response with the 
message "Not Found" in the response body.

4) "r.URL.Path" used to get the path of the request 

5) "r.Method" is used to get the method of the request, e.g. GET/POST/UPDATE/DELETE

6) http.ListenAndServe(addr string, handler http.Handler)
EX-1: Server listens on port 8080
http.ListenAndServe(":8080", nil)

EX-2: Server listen on 8080 and executes the handler
func myHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, world!")
}
http.ListenAndServe(":8080", http.HandlerFunc(myHandler))

7) The log.Fatal(err) function in Go logs an error message to the 
standard error output and then immediately exits the program with 
a non-zero exit code. This indicates that the program has 
encountered a serious error and cannot continue running.

EX: This example defines a main function that attempts to open 
a file named myfile.txt. If the file cannot be opened, the 
function logs the error using log.Fatal() and exits the program.

func main() {
    err := os.Open("myfile.txt")
    if err != nil {
        log.Fatal(err)
    }

    // Do something with the opened file
}



