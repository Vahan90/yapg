#!/bin/bash

set -e

# Push a single metric line to yapg
echo "Pushing metrics..."
curl -X POST http://localhost:9091/push --data-binary @- <<EOF
event_processing_time_seconds{event_id="test123"} 321
EOF

echo "Waiting 1 second..."
sleep 1

# Scrape from /metrics (this also deletes the pushed metrics)
echo "Scraping metrics..."
curl http://localhost:9091/metrics

echo "Waiting 1 second..."
sleep 1

# Scrape again to confirm metrics were deleted
echo "Scraping again to confirm cleanup..."
curl http://localhost:9091/metrics
