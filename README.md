# go-gin-gam-rae (고진감래)

<img align="right" width="159px" src="https://user-images.githubusercontent.com/24591259/222955830-f9cbb8f3-b1e8-4cee-8995-bda4d324dc71.png">

`golang`, `gin` 스터디 프로젝트 `고진감래(go-gin-gam-rae)` 입니다. <br>
`golang`으로 무언갈 만들면서 학습을 하는 레포지토리 입니다!

## 🚀 The goal of go-gin-gam-rae are

- api health check
- api gateway
- gin
- grpc
- DB connection
- golang unittest

## 🏃‍♂️ How to run

`env`파일을 복사하고 docker-compose 내용에 맞게 수정해줍니다.

```shell
cp .env_sample .env
```

`docker-compose`로 `MySQL server`, `python grpc server`, `golang grpc server`를 띄웁니다. <br>
proto파일 수정을 하게되면 아래 proto파일 컴파일 과정을 참고하여 새로 컴파일을 합니다.

```shell
docker-compose up -d
```

## 🛠 Golang Proto Buffer Compile

#### 1. protoc 컴파일러 설치

```shell
brew install protobuf
```

#### 2. go path 세팅 (.zshrc 같은곳)

환경설정 할때 `GOPATH` 설정이 잘 안되서 애먹는 경우가 많습니다. <br>
라이브러리를 설치해도 뭔가 잘 안되는 것 같으면 `go env GOPATH` 명령어로 현재 콘솔에서 GOPATH가 잘 지정되었는지 확인해봐야 합니다.

```shell
export GOPATH=$HOME/go #~/go
export PATH=$PATH:$GOPATH/bin
```

#### 3. grpc 관련 패키지 설치

```shell
go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc 
```

#### 4. proto buffer 컴파일

| 컴파일 옵션                     | 설명                    |
|:---------------------------|:----------------------|
| `logtostderr`              | 콘솔에 로그 출력             |
| `paths=source_relative`    | 생성되는 파일경로를 상대경로로      |
| `generate_unbound_methods` | 언바운드 메서드를 생성할지 여부를 지정 |

```shell
# cosmetic protobuffer 컴파일
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/cosmetic/cosmetic.proto

# hello protobuffer 컴파일
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/hello/hello.proto

# cosmetic grpc gateway protobuffer 컴파일    
protoc -I . \
		--grpc-gateway_out=allow_delete_body=true:. \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt generate_unbound_methods=true \
		./proto/cosmetic/cosmetic.proto

# hello grpc gateway protobuffer 컴파일
protoc -I . \
		--grpc-gateway_out=allow_delete_body=true:. \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt generate_unbound_methods=true \
		./proto/hello/hello.proto
```

## 🛠 Python Proto Buffer Compile

```shell
python -m grpc_tools.protoc \
	-I . \
    --python_out=. \
    --grpc_python_out=. \
    ./proto/hello/hello.proto


python -m grpc_tools.protoc \
	-I . \
    --python_out=. \
    --grpc_python_out=. \
    ./proto/cosmetic/cosmetic.proto

# python의 경우 google api 관련 proto 파일도 같이 컴파일을 해줘야 함
# grpc gateway를 사용하지 않기때문에 해당 파일들도 컴파일 해줘야 하는걸로 보임 (뇌피셜) 
python -m grpc_tools.protoc \
	-I . \
    --python_out=. \
    --grpc_python_out=. \
    ./proto/google/api/annotations.proto


python -m grpc_tools.protoc \
	-I . \
    --python_out=. \
    --grpc_python_out=. \
    ./proto/google/api/http.proto

```

## Contributors
<a href="https://github.com/Jay-Chan9yu/go-gin-gam-rae/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=Jay-Chan9yu/go-gin-gam-rae" />
</a>

Made with [contrib.rocks](https://contrib.rocks).

## 🔏 License

[MIT](https://choosealicense.com/licenses/mit/)

