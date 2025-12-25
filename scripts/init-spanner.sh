#!/bin/bash
set -e

PROJECT_ID="test-project"
INSTANCE_ID="test-instance"
DATABASE_ID="test-database"

echo "Waiting for Spanner emulator to be ready..."
sleep 2

echo "Creating Spanner instance..."
gcloud spanner instances create ${INSTANCE_ID} \
  --config=emulator-config \
  --description="Test Instance" \
  --nodes=1

echo "Creating Spanner database with schema..."
gcloud spanner databases create ${DATABASE_ID} \
  --instance=${INSTANCE_ID} \
  --ddl-file=/schema/spanner.sql

echo "Spanner initialization completed successfully!"
