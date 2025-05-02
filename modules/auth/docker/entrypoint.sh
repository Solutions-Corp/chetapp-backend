set -e

echo "Starting Auth Service..."

uvicorn main:app --host ${HOST:-0.0.0.0} --port ${APP_PORT:-8080} --reload