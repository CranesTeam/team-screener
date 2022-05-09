# create user

curl --location --request POST 'http://localhost:8089/auth/sing-up' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Ioan Ioanovich",
    "username": "test_user_ioan",
    "password": "pass"
}'