{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": { "Service": "ec2.amazonaws.com"},
      "Action": "sts:AssumeRole"
    }
  ]
}
aws iam create-role --role-name dynamoAccess --assume-role-policy-document file://ec2-role-trust-policy.json

{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": ["s3:*"],
      "Resource": ["*"]
    }
  ]
}

aws iam put-role-policy --role-name dynamoAccess --policy-name dynamo-perm --policy-document file://ec2-role-access-policy.json

aws iam create-instance-profile --instance-profile-name dynamoAccess-profile

aws iam add-role-to-instance-profile --instance-profile-name dynamoAccess-profile --role-name dynamoAccess

aws ec2 associate-iam-instance-profile arn="arn:aws:iam::631795995269:instance-profile/dynamoAccess-profile",name="dynamoAccess-profile"  --instance_id i-02864a50eee3e22c1
