#!/usr/bin/env bash
set -eo pipefail

# Create mount directory for service
mkdir -p $MNT_DIR

echo "Mounting GCS Fuse."
gcsfuse --implicit-dirs --debug_gcs --debug_fuse $BUCKET $MNT_DIR
echo "Mounting completed."

# Start the application
/app/app serve --http=0.0.0.0:8080

# Exit immediately when one of the background processes terminate.
wait -n