curl --location --request POST 'http://localhost:9999/addCard' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": 5435,
    "issuer": "Visa",
    "currency": "USD",
    "virtual": false
}'

curl --location --request POST 'http://localhost:9999/addCard' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": 5435,
    "issuer": "Visa",
    "currency": "USD",
    "virtual": true
}'

curl --location --request POST 'http://localhost:9999/addCard' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": 5435,
    "issuer": "Visa",
    "currency": "USD",
    "virtual": true
}'

curl --location --request POST 'http://localhost:9999/addCard' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": 5436,
    "issuer": "Visa",
    "currency": "USD",
    "virtual": false
}'
