# aws-lambda-create-file
Simple test to create a file in an AWS Lamba function tmp directory.
With this repo isolating my bug from a larger repo, I discovered that for local development, setting the FILE_PATH env variable to "./tmp/newfile.csv" works quite well, but that file path fails in the Lambda function. I was required to change the path to "/tmp/newfile.csv" to successfully create the file.
