aws cloudformation deploy \
--template-file serverless-api-output.yaml \
--stack-name test-lambda-stack-3 \
--capabilities CAPABILITY_IAM \
--region us-east-1 
