# Go Getter

This Go application leverages the Colly library to efficiently scrape book information from the web and store it securely in an AWS RDS database. Deployed with Terraform, it offers a streamlined setup of AWS resources, ensuring scalability and reliability in the cloud.

## Features

- Web scraping with Go and Colly
- Secure data storage in AWS RDS
- Automated AWS resource provisioning with Terraform
- Scalable and reliable cloud-based architecture

## Getting Started

### Prerequisites

- Go (version 1.15 or newer)
- AWS CLI and AWS account
- Terraform

### Installation

1. Clone the repository:
```bash
git clone https://github.com/uprightsleepy/Go_Getter
```

2. Navigate to the project directory:
```bash
cd GoGetter
```

3. Initialize Terraform to set up AWS resources:
```bash
terraform init
terraform apply
```
