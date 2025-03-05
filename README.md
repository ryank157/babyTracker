# Baby Tracker

## Project Overview

Baby Tracker is a comprehensive baby monitoring application integrating Home Assistant, Go backend, HTMX frontend, and local LLM technologies.

## Project Structure

- `babyTracker/`: No description
- `backend/`: Go Web Application Backend
  - `cmd/`: Application entry points
  - `internal/`: Internal packages and core logic
    - `database/`: Database access layer
    - `llm/`: LLM integration and functionalities
    - `models/`: Data models and definitions
    - `services/`: Business logic and services
  - `migrations/`: Database migration scripts
- `docs/`: Project Documentation
- `frontend/`: HTMX Web Frontend
  - `static/`: Static web assets
  - `templates/`: HTML templates
- `homeassistant/`: Home Assistant Integration
  - `custom_components/`: Custom Home Assistant components
    - `baby_tracker/`: Custom Home Assistant component for baby tracking
- `scripts/`: Utility Scripts

## Key Components

### Homeassistant
Custom Home Assistant integration for system-wide monitoring

### Docs
Comprehensive project documentation

### Scripts
Utility scripts for setup and deployment

### Backend
Go-based web application backend with database and LLM integration

### Frontend
HTMX-powered web interface for user interactions

## Development

### Prerequisites
- Go 1.24+
- Home Assistant
- PostgreSQL
- Docker (optional)

### Setup
1. Clone the repository
2. Copy `.env.example` to `.env` and configure
3. Run `./scripts/setup.sh`

## License
MIT License

## Contributing
Please read `docs/DEVELOPMENT.md` for details on our code of conduct and the process for submitting pull requests.
