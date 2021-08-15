Install golang 1.16 and great. How install see here https://golang.org/doc/install

Run:
`git clone git@github.com:Ivnkara/ocr-test.git && cd ocr-test`

Create .env file and add this is fields:

`API_KEY=xxxxxxxxx # API key for ocr space service
APP_LANG=rus      # Lang for ocr
PORT=:8080        # Port for running web-server```

After run:
```go run main.go```

Web server have interface:
```curl 'http://localhost:8080/from-url?url=https://url.to/your/picture1&url=https://url.to/your/picture2'```
Or local picture
```curl -X POST 'http://localhost:8080/from-file' -F "file=@/home/user/Pictures/photo_01.jpg"```
