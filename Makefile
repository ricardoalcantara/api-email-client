version := v0.0.1
commit := $(shell git rev-parse HEAD)
build_os := $(shell lsb_release -d -s)
date_time := $(shell date "+%Y-%m-%d-%H-%M-%S")
timestamp := $(shell date "+%s")
cwd := $(shell pwd)
app_name := api-email-client
build:
	go build \
	-o ${app_name} \
	-ldflags "-X 'github.com/ricardoalcantara/${app_name}/internal/version.GitCommit=${commit}' -X 'github.com/ricardoalcantara/${app_name}/internal/version.BuildOS=${build_os}' -X 'github.com/ricardoalcantara/${app_name}/internal/version.BuildDate=${date_time}' -X 'github.com/ricardoalcantara/${app_name}/internal/version.Version=${version}'" \
	cmd/main.go

build_with_docker:
	docker run -it -v ${cwd}:/app -w /app golang:bullseye make build

build_docker_dev:
	docker build . -t registry.pscanary.com/${app_name}:${timestamp}
	docker push registry.pscanary.com/${app_name}:${timestamp}
	echo registry.pscanary.com/${app_name}:${timestamp}

install_service:
	cp ${app_name}.service /etc/systemd/system/${app_name}.service
	systemctl start ${app_name}.service
	systemctl enable ${app_name}.service

restart_service:
	systemctl restart ${app_name}.service