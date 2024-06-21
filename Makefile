.SILENT: build init plan apply apply_auto destroy destroy_auto fmt clean tf-doc

build:
	cd cmd/ && GOARCH=arm64 GOOS=linux go build -o ../bin/bootstrap main.go

init:
	cd infrastructure/ && \
	terraform init

plan: build
	cd infrastructure/ && \
	terraform plan -var-file env_vars/dev.tfvars

apply: build
	cd infrastructure/ && \
	terraform apply -var-file env_vars/dev.tfvars

apply_auto: build
	cd infrastructure/ && \
	terraform apply -var-file env_vars/dev.tfvars -auto-approve

destroy:
	cd infrastructure/ && \
	terraform apply -destroy

destroy_auto:
	cd infrastructure/ && \
	terraform apply -destroy -auto-approve

fmt:
	go fmt ./...
	cd infrastructure/ && \
	terraform fmt -recursive

clean:
	rm -rf bin/
	cd infrastructure/ && \
	rm -rf .terraform/ .terraform* terraform* lambda.zip

tf-doc:
	terraform-docs markdown infrastructure/ > infrastructure/README.md

up: build init apply_auto

down: destroy_auto clean
