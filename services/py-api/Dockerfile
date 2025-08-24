# --- builder ---
FROM python:3.12-slim AS builder
ENV PYTHONDONTWRITEBYTECODE=1 PYTHONUNBUFFERED=1
WORKDIR /w
COPY requirements.txt ./
RUN apt-get update && apt-get install -y --no-install-recommends build-essential     && pip install --upgrade pip     && pip wheel --wheel-dir /wheels -r requirements.txt     && apt-get purge -y build-essential && rm -rf /var/lib/apt/lists/*

# --- runtime ---
FROM python:3.12-slim AS runtime
ENV PYTHONDONTWRITEBYTECODE=1 PYTHONUNBUFFERED=1
WORKDIR /app
COPY --from=builder /wheels /wheels
COPY requirements.txt /wheels/../requirements.txt
RUN pip install --no-cache-dir --find-links=/wheels -r /wheels/../requirements.txt     && rm -rf /wheels
COPY app/ ./app/
# Security hardening
RUN useradd -r -u 10001 -g root appuser && mkdir -p /app/tmp && chown -R appuser:root /app
USER 10001
EXPOSE 8080
HEALTHCHECK --interval=30s --timeout=3s --retries=3 CMD curl -fsS http://localhost:8080/healthz || exit 1
CMD ["python", "-m", "uvicorn", "app.main:app", "--host", "0.0.0.0", "--port", "8080"]
