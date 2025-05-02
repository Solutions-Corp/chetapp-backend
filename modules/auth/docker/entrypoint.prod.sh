set -e

echo "Generating keys..."

apk add --no-cache openssl

KEY_DIR="/keys"
mkdir -p $KEY_DIR

openssl genpkey -algorithm RSA -out $KEY_DIR/private.pem -pkeyopt rsa_keygen_bits:2048

openssl rsa -pubout -in $KEY_DIR/private.pem -out $KEY_DIR/public.pem

chmod 600 $KEY_DIR/private.pem
chmod 644 $KEY_DIR/public.pem

echo "Keys generated successfully."

echo "Starting Auth Service..."
exec gunicorn -c gunicorn.conf.py main:app
