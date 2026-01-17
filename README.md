# K8S Health ❤️ and Readiness 👟

This repository holds all necessary components to demonstrate Liveness and Readiness checks on a pod.

## Components

- Web Server
    - Provide `healthz` endpoint for Liveness check, there is a 33% chance of a failed start
    - Provide `readiness` endpoint for Readiness check
    - Provide `toggle` endpoint for switching off/on the application readiness
- Manifest files for Deployment and Service resources