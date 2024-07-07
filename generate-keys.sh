mkdir keys

openssl ecparam -name prime256v1 -genkey -noout -out ./keys/ecdsa_p256_private_key.pem
openssl ec -in ./keys/ecdsa_p256_private_key.pem -pubout -out ./keys/ecdsa_p256_public_key.pem

openssl genpkey -algorithm RSA -out ./keys/private_key.pem -pkeyopt rsa_keygen_bits:2048
openssl rsa -pubout -in ./keys/private_key.pem -out ./keys/public_key.pem
