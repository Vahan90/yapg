# yapg – Yet Another Pushgateway

**yapg** is a minimal and scalable metrics sink designed to improve upon [Prometheus Pushgateway](https://github.com/prometheus/pushgateway) behavior. It is built for ephemeral job metrics: it accepts raw Prometheus metric lines via a `/push` HTTP endpoint and automatically deletes them after they are scraped from the `/metrics` endpoint.

Backed by [Redis](https://github.com/redis/redis), `yapg` is stateless, scalable, and avoids stale metric buildup—a perfect fit for short-lived jobs and batch pipelines like CronJobs or regular Jobs.

## ✨ Features

* 📥 Push once: Send metrics via POST /push using Prometheus text format.
* 🧹 Self-cleaning: Scraping /metrics serves and deletes pushed metrics.
* ⚡ Backed by Redis: Stateless design, horizontally scalable, minimal memory footprint.
* 🐳 Docker & Compose-ready: Quickly test locally with Docker Compose.
* 🚀 Helm Chart (coming soon): Deploy easily on Kubernetes alongside Redis.

## 📦 Installation

### 🐳 Docker Compose (Local Dev/Test)

You can run yapg locally with Redis using Docker Compose:

```bash
docker-compose up --build
```

This will start two containers:
* `redis` on port `6379`
* `yapg` on port `9091`

## 📤 Pushing metrics

Send Prometheus-formatted metrics to `/push` using a `POST` request:

```bash
curl -X POST http://localhost:9091/push --data-binary @- <<EOF
event_processing_time_seconds{event_id="test123"} 321
EOF
```

## 📥 Scraping Metrics

Scraping the /metrics endpoint:
* Returns all pushed metrics in standard Prometheus format
* Deletes all scraped metrics after serving them

```bash
curl http://localhost:9091/metrics
```

If you scrape again, you'll see an empty response unless new metrics were pushed

## 🧪 Full Test Example

Here’s a complete workflow as shown in `test.sh`:

```bash
#!/bin/bash
set -e

# Push a metric
curl -X POST http://localhost:9091/push --data-binary @- <<EOF
event_processing_time_seconds{event_id="test123"} 321
EOF

# First scrape: shows metric
curl http://localhost:9091/metrics

# Second scrape: shows nothing
curl http://localhost:9091/metrics
```

## 🧰 Configuration

| ENV Variable | Description | Default |
| ------------ | ----------- | ------- |
|   `REDIS_ADDR` | Redis connector URI | `localhost:6379` |

## 📜 Prometheus Integration

In your Prometheus config:

```yaml
scrape_configs:
  - job_name: 'yapg'
    static_configs:
      - targets: ['yapg-host:9091']
```

Ensure the `Prometheus server` and `yapg` instance can communicate.

## 🔧 Development

### Prerequisites

* Go 1.23
* Redis
* Docker / Docker Compose

## 📌 Roadmap

- [X] Redis-backed yapg
- [X] Prometheus `/metrics` endpoint with auto-cleanup
- [ ] Helm chart with bundled Redis
- [ ] Authentication/Authorization for metrics push
- [ ] Multi-tenancy with labels


