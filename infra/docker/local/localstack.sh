# -- > List SNS Topics
echo Listing topics ...
echo $(aws --endpoint-url=http://localhost:4566 sns list-topics --profile test-profile --region us-east-1 --output table | cat)

# -- > Send SNS status update event
echo Sending SNS status update event ...
echo $(aws --endpoint-url=http://localhost:4566 sns publish --topic-arn arn:aws:sns:us-east-1:000000000000:order_status_update --message "{\"order_id\": 2, \"status\": \"FINISHED\"}" --profile test-profile --region us-east-1)

# -- > Send SNS payment rollback event
echo Sending SNS payment rollback event ...
echo $(aws --endpoint-url=http://localhost:4566 sns publish --topic-arn arn:aws:sns:us-east-1:000000000000:payment_status_update --message "{\"order_id\": 3, \"status\": \"CANCELED\"}" --profile test-profile --region us-east-1)