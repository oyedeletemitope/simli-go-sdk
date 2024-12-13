# Simli-Go-SDK

Simli-Go-SDK is a Go library for interacting with the Simli API, enabling the creation of avatar videos from audio or text inputs. This SDK simplifies the process of generating avatar-based content by abstracting the underlying API calls.

---

## **Features**

- **Audio-to-Video Avatar Generation**:
  Convert audio files into videos of an animated avatar.
- **Text-to-Video Avatar Generation**:
  Generate avatar videos from text using built-in or third-party TTS services.
- **Session Management**:
  Support for initializing and managing audio-to-video sessions.

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
Use an .env file to securely store your credentials:

```bash
SIMLI_API_KEY=your-simli-api-key
SIMLI_FACE_ID=your-face-id
```

Install the godotenv package to load environment variables:

```bash
go get github.com/joho/godotenv
```
