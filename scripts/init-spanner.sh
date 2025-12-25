#!/bin/bash
set -e

PROJECT_ID="test-project"
INSTANCE_ID="test-instance"
DATABASE_ID="test-database"

# ローカルSpanner Emulatorを常に使用
export SPANNER_EMULATOR_HOST="${SPANNER_EMULATOR_HOST:-localhost:9010}"

echo "Using Spanner Emulator at: ${SPANNER_EMULATOR_HOST}"
echo "Waiting for Spanner emulator to be ready..."
sleep 2

# エミュレータ使用時は認証を無効化
gcloud config set auth/disable_credentials true
gcloud config set project ${PROJECT_ID}
gcloud config set api_endpoint_overrides/spanner ${SPANNER_EMULATOR_HOST}

echo "Creating Spanner instance..."
gcloud spanner instances create ${INSTANCE_ID} \
  --config=emulator-config \
  --description="Test Instance" \
  --nodes=1

echo "Creating Spanner database with schema..."
gcloud spanner databases create ${DATABASE_ID} \
  --instance=${INSTANCE_ID} \
  --ddl-file=/schema/spanner.sql

echo "Spanner initialization completed successfully!"gcloud spanner instances create
