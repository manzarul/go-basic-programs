# go-basic-programs
## prerequisite to run go program
- install the GO language , ref : https://go.dev/doc/install
## Go program can be run as follow
- go run {path-of-go-file} // this step will interally compile and run the program
- OR you can first compile it then it will generate {file-name}.exe file and you can directly run this exe file.
- go build {path-of-go-file}
- then {file-name}.exe // to run the .exe file
### to test http decode you can use below curl
curl --location --request POST 'http://localhost:8080/decode' \
--header 'Content-Type: application/json' \
--data-raw '{"firstname":"John","lastname":"Doe","age":25}'

And

curl --location --request GET 'http://localhost:8080/encode'
