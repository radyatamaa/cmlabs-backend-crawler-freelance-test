# technical-test-cognotiv


### Clean Architecture
This project has  4 Domain layer :

 * Models Layer
 * Repository Layer
 * Usecase Layer  
 * Delivery Layer

#### The diagram:

![golang clean architecture](https://github.com/bxcodec/go-clean-arch/raw/master/clean-arch.png)

The explanation about this project's structure  can read from this medium's post : https://medium.com/@imantumorang/golang-clean-archithecture-efd6d7c43047

### How To Run This Project

```bash
#move to directory
cd $GOPATH/src/github.com/radyatamaa

# Clone into YOUR $GOPATH/src
git clone https://github.com/radyatama/cmlabs-backend-crawler-freelance-test.git

#move to project
cd technical-test-cognotiv

# Run app 
go run main.go

# Run worker
go run ./cmd/worker/main.go

# if you not installed yet redis and mysql
docker compose -f "docker-compose.yml" up -d --build

# Open at browser this url
http://localhost:8082/swagger/index.html
```


### Swagger UI:

http://localhost:8082/swagger/index.html

### More about app details:
open file tutorial test apps.pdf