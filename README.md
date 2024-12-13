# Simli-Go-SDK

Simli-Go-SDK is a Go library for interacting with the Simli API. This SDK simplifies generating avatar-based content by abstracting the underlying API calls.

---

## **Installation**

Add the library to your Go project using `go get`:
```bash
go get github.com/oyedeletemitope/simli-go-sdk
```

## **Import the Library**
Import the required packages in your Go file:

```go
    "github.com/oyedeletemitope/simli-go-sdk/client"
    "github.com/oyedeletemitope/simli-go-sdk/models"
```
## **Set Up Your Environment**
Use an .env file to store your credentials securely:

```bash
SIMLI_API_KEY=your-simli-api-key
SIMLI_FACE_ID=your-face-id
```

Install the godotenv package to load environment variables:

```bash
go get github.com/joho/godotenv
```
