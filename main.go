package main

import( "fmt"
	"io"
	"net/http"
	"os"
	)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {

         // the FormFile function takes in the POST input id file
         file, header, err := r.FormFile("file")

         if err != nil {
                 fmt.Fprintln(w, err)
                 return
         }


         defer file.Close()

         out, err := os.Create("/tmp/uploadedfile")
         if err != nil {
                 fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privilege")
                 return
         }

         defer out.Close()

         // write the content from POST to the file
         _, err = io.Copy(out, file)
         if err != nil {
                 fmt.Fprintln(w, err)
         }

         fmt.Fprintf(w, "File uploaded successfully : ")
         fmt.Fprintf(w, header.Filename)
 }
func main() {
 	fmt.Printf("hello, world\n")
	http.HandleFunc("/", hello)
	http.ListenAndServe(":9080", nil)
}
