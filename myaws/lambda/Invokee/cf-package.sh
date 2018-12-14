aws cloudformation package    \
--template-file serverless-api.yml \
--output-template-file serverless-api-output.yaml \
--s3-bucket lambda-deploy-virginia 
