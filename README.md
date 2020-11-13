# golang-line-chatbot

## Prerequisite
- GO
- Serverless
- AWS Account

## Install Env
- npm install -g serverless
- export AWS_ACCESS_KEY_ID=<your_aws_access_key_id>
- export AWS_SECRET_ACCESS_KEY=<your_aws_secret_access_key>

## How to get aws access key
- Open the AWS Console
- Click on your username near the top right and select My Security Credentials
- Click on Users in the sidebar
- Click on your username
- Click on the Security Credentials tab
- Click Create Access Key
- Click Show User Security Credentials

## Create the project
- sls create --template aws-go --path myService

## Build the project
- make deploy
- execute Url

## Check your AWS
- check lambda
- check api gateway
- check cloudFormation
