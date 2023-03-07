# GoLang CRUD operation App 🚀
## Using Gin - Gonic framework, MongoDB Atlas & Keploy for testing 👀

- ✨By Ayush Sahu✨

- Star🌟 this repository and fork🍴 to save it to your repos for easy access.

Connect with me : [LinkedIn] [Twitter]

### Running keploy in record mode to capture test cases

```sh
export KEPLOY_MODE=record
go run main.go
```

Preview of the project running 🚀 :

<img width="1142" alt="Screenshot 2023-03-07 at 3 33 20 PM" src="https://user-images.githubusercontent.com/102364606/223393822-5563b5bf-a76d-4b31-b32e-fe535cab1589.png">

Preview of Keploy capturing test cases 🔍 : 

<img width="1440" alt="Screenshot 2023-03-07 at 3 32 41 PM" src="https://user-images.githubusercontent.com/102364606/223393918-69b05994-224c-4974-a3be-9de815185ba4.png">

### Running keploy in test mode to get test coverage

```sh
export KEPLOY_MODE="test"
go test -v -coverpkg=./... -covermode=atomic  ./...
```

Preview of automatic testing with keploy 🔥 : 

<img width="866" alt="Screenshot 2023-03-07 at 3 33 36 PM" src="https://user-images.githubusercontent.com/102364606/223393946-c6076c51-1f68-4fe4-a926-3271a82ff2da.png">
<img width="887" alt="Screenshot 2023-03-07 at 3 32 46 PM" src="https://user-images.githubusercontent.com/102364606/223393962-5bd59ad6-1137-423d-8a68-58d60eb6b3a7.png">

So cool! 😎

P.s. Got less coverage as I didn't capture all functionality (CRUD) of the app.

Used API testing platform : [Postman] 10.11.1

Used IDE : Visual Studio Code 1.76.0 

OS : MacOS Ventura 13.2

Download Visual Studio Code on your machine at [VScode]

## For more info visit official Keploy documentation here [Keploy-Offical]

[//]: # 

[LinkedIn]: <https://www.linkedin.com/in/ayushsahu77/>
[Twitter]: <https://twitter.com/Ayush_7477/>
[Postman]: <https://www.postman.com//>
[VScode]:<https://code.visualstudio.com//>
[Keploy-Offical]: <https://github.com/keploy/>
