# WASM-Password-Gen

A very simple password generator built to try out WASM with Go

## Getting Started

- Have Go installed
- Run ```bash setup.bash``` in the top dir
- Run ```bash build.bash``` in ```/password-gen```
- Build and run the server with ```go build``` and ```./server``` in ```/server```

## Things to look at next

- The WASM binary that is produced is huge for what it is, should look at *TinyGo* (https://tinygo.org/getting-started/)
- Build something that actually benefits from WASM performance
