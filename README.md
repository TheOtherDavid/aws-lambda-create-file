# aws-lambda-create-file
Simple test to create a file in an AWS Lamba function tmp directory.
With this test isolating my bug, I discovered that for local development
For local development, I discovered that "./tmp/newfile.csv" works quite well, but that file path fails in the Lambda function. I was required to change the path to "/tmp/newfile.csv".
