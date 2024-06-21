# Serverless AWS CRUD API using DynamoDB, written in GO and deployed with Terraform

## Overview

AWS Lambda is a serverless computing service that allows you to run code without provisioning or managing servers. It supports multiple programming languages, including Go, which is known for its simplicity and performance. DynamoDB is a fully managed NoSQL database service provided by AWS that offers low latency and high scalability. Terraform is an open-source infrastructure as code software tool that enables you to define and provision data center infrastructure using a declarative configuration language.

This project showcases a straightforward AWS Lambda function developed in Go, which is deployed utilizing Terraform. It serves as a practical example for those looking to get started with AWS Lambda functions using Go and manage their deployment through Terraform's infrastructure as code capabilities.

The Lambda function implements a simple CRUD API for managing notes. It supports the following operations:
- **Create**: Add a new note.
- **Read**: Retrieve all notes or a specific note by ID.
- **Update**: Modify an existing note.
- **Delete**: Remove a note.

The project includes the following components:
- **Go Lambda function**: The Lambda function is written in Go and implements the CRUD API operations.
- **Terraform configuration**: The Terraform configuration files define the AWS resources required for the Lambda function, including the Lambda function itself, an API Gateway, and the necessary IAM roles and policies.
- **Makefile**: The Makefile provides convenient targets for building the Go code, initializing Terraform, creating the deployment plan, and applying the plan to provision the AWS resources.

By following the instructions in this README, you can deploy the Lambda function and API Gateway using Terraform and test the API endpoints.

## Project Structure

The project is structured as follows:
- **`cmd/main.go`**: The main Go file that contains the Lambda function code.
- **`infrastructure`**: The directory containing the Terraform configuration files for provisioning the AWS resources. For more details see [infrastructure/README.md](infrastructure/README.md).
- **`internal`**: The directory containing Go code for the Lambda function.
    - **`internal/api/api.go`**: The file containing the API handlers for the CRUD operations.
    - **`internal/db/db.go`**: The file containing the dynamodb client and methods to interact with the database.
    - **`internal/models/note.go`**: The file containing the data model for the notes.
- **`pkg/logger/logger.go`**: The file containing the logger utility for logging messages.
- **`.editorconfig`**: The configuration file for defining code styles and formatting.
- **`.gitignore`**: The file specifying which files and directories to ignore in version control.
- **`go.mod`**: The Go module file that defines the project's dependencies.
- **`go.sum`**: The Go sum file that contains the expected cryptographic checksums of the content of specific module versions.
- **`LICENSE`**: The license file for the project.
- **`Makefile`**: The file containing targets for building, deploying, and testing the project.
- **`README.md`**: The README file with project information and instructions.

## Prerequisites

Before you begin, ensure you have the following:
- make utility installed on your system.
- Go programming language installed on your system.
- An active AWS account.
- Terraform installed and configured.
- AWS CLI set up with appropriate access credentials.

## How to use

To set up the project:
1. Clone this repository to your local environment.
2. Change to the project directory.
3. Execute `make up` in your terminal. This command compiles the Go files, initializes Terraform, creates the deployment plan, and applies it to provision the necessary AWS resources, including the Lambda function.

## Testing the API endpoints

Once the AWS infrastructure is deployed:
- **Invocation via CURL**: You can invoke the Lambda function using the AWS API Gateway with a command like `curl -X GET https://replace-me.execute-api.eu-central-1.amazonaws.com/dev/notes`, the exact command is also an output of Terraform apply.
- **Postman**: You can use Postman to test the API endpoints. The API Gateway URL is also an output of Terraform apply.
- **AWS API Gateway console**: You can also test the API endpoints via AWS API Gateway console. 
- **AWS Lambda console**: You can view the Lambda function logs in the AWS Lambda console.

## Troubleshooting

If you encounter issues:
- Double-check your AWS CLI configuration to ensure it has the necessary permissions.
- Review the Terraform output for any deployment errors.
- Make sure your Go environment is correctly set up, especially if you're making changes to the Lambda function code.

## Contributing

I welcome contributions to improve this project. Feel free to fork the repository, make your changes, and submit a pull request. If you have suggestions or encounter issues, please open an issue in the repository.

## License

This project is made available under the MIT License. For more details, refer to the [LICENSE](LICENSE) file.
