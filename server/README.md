# LeBlanc Go API — Local dev guide

This document explains how to run the LeBlanc Go API locally, which ports it exposes, the expected environment variables, and example requests to interact with the API.

**Port**
- Default HTTP port: `4000` (configurable via `PORT` env var)

**Environment variables**
- `PORT` — HTTP port (default `4000`).
- `MONGO_URI` — MongoDB connection string (e.g. `mongodb://127.0.0.1:27017`).
- `MONGO_DB` — Mongo database name (e.g. `leblanc`).
- `API_PUBLIC_KEY` — optional public API key used by your app.
- `API_PRIVATE_KEY` — optional private API key used by your app.

An example environment template is provided at `server/.env.example`. For local development you can copy it to `server/.env` and fill in your values:

```powershell
Copy-Item .\server\.env.example .\server\.env
# edit server\.env with Notepad or your preferred editor
notepad .\server\.env
```

Important: `server/.env` is in `.gitignore` to avoid committing secrets. Rotate any keys that were exposed publicly.

**Start a local MongoDB (Docker)**
If you don't have a running MongoDB instance, run a local container:

```powershell
docker run -d --name leblanc-mongo -p 27017:27017 `
  -e MONGO_INITDB_DATABASE=leblanc mongo:6.0

# verify listener (Windows)
netstat -ano | findstr 27017
```

Or use your system's `mongod` service and set `MONGO_URI` accordingly.

**Run the API (development)**
From the repository root run:

```powershell
Set-Location server
# loads .env via github.com/joho/godotenv in development
go run main.go
```

To run in background and capture logs:

```powershell
Set-Location server
# Start the server and redirect output to a log file
Start-Process -FilePath 'go' -ArgumentList 'run main.go' -RedirectStandardOutput 'server.log' -RedirectStandardError 'server.log' -NoNewWindow

# To list the process (and note its Id):
Get-Process -Name go

# Follow logs (similar to tail -f):
Get-Content .\server.log -Wait
```

Stop the server (if started as above):

```powershell
# Stop the Go process started above (if any):
Get-Process -Name go | Stop-Process -Force
Remove-Item .\server.pid -ErrorAction SilentlyContinue
```

**API Endpoints**

- `GET /` — health / welcome message

- `GET /drinks` — list all drinks
  - Response: JSON array of `Drink` objects. See models in `internal/models/drink.go`.
  - Example:
    ```bash
    curl http://localhost:4000/drinks
    ```

- `POST /bookings` — create a booking
  - Request JSON matches `Booking` model in `internal/models/booking.go`.
  - Example payload:
    ```json
    {
      "name": "Alice",
      "phone": "555-1234",
      "time": "2025-11-16T12:30:00Z",
      "items": [
        {"drinkId": "652f8f8e8c9f1a2b3c4d5e6f", "qty": 2, "options": {"milk":"oat"}}
      ],
      "channel": "web"
    }
    ```
  - Example curl:
    ```bash
    curl -X POST http://localhost:4000/bookings \
      -H 'Content-Type: application/json' \
      -d @booking.json
    ```

- `POST /reco/from-features` — recommendation from features (HTTP API)
  - Request JSON uses `services.RecoPayload` shape (see `internal/services/reco_score.go`). Minimal example:
    ```json
    {
      "emotion": "calm",
      "colorTone": "warm",
      "context": {"timeOfDay": "morning", "tempPref": "hot"}
    }
    ```
  - Example curl:
    ```bash
    curl -X POST http://localhost:4000/reco/from-features \
      -H 'Content-Type: application/json' \
      -d '{"emotion":"calm","colorTone":"warm","context":{"timeOfDay":"morning","tempPref":"hot"}}'
    ```

- `POST /graphql` and `GET /graphql` — GraphQL endpoint + GraphiQL playground
  - The GraphQL API exposes a `reco` query. Example query:
    ```graphql
    query {
      reco(emotion: "happy", colorTone: "cool", timeOfDay: "afternoon", topK: 3) {
        _id
        name
        price
        score
      }
    }
    ```
  - Use curl (POST) or open http://localhost:4000/graphql in your browser to view GraphiQL.

**Example: full flow**
1. Start Mongo (Docker) if needed.
2. Ensure `server/.env` has `MONGO_URI` pointing at `mongodb://127.0.0.1:27017` and `MONGO_DB=leblanc`.
3. Start the server (`go run main.go`).
4. Load initial `drinks` data into Mongo (if your collection is empty) — you can use `mongoimport` or a small script.
5. Call `POST /reco/from-features` or open `GET /graphql` to get recommendations.

**Troubleshooting**
- `Mongo ping error` (context deadline exceeded / connection refused): Mongo isn't running or `MONGO_URI` is incorrect. Start Mongo or update `MONGO_URI`.
- `bind: address already in use` on port `4000`: another process is running. Find and kill it or change `PORT`.
- If GraphiQL returns errors, check server logs at `server.log` for stack traces.

**Security & production notes**
- Do not store secrets in plaintext in version control. Use environment variables provided by your deployment platform or a secrets manager (AWS Secrets Manager, GCP Secret Manager, Azure Key Vault, Kubernetes Secrets, etc.).
- Rotate `API_PRIVATE_KEY` immediately if it was exposed.

**File structure**
This is the important file layout for the `server/` service:

- `main.go` — application entrypoint and router setup
- `go.mod`, `go.sum` — Go module files
- `.env.example` — example environment variables for local development
- `server.log` — local server log file (created when running in background)
- `internal/`
  - `db/mongo.go` — MongoDB connection initialization
  - `handlers/` — HTTP handlers
    - `drinks.go` — `GET /drinks`
    - `bookings.go` — `POST /bookings`
    - `reco.go` — `POST /reco/from-features`
    - `reco_graphql.go` — GraphQL schema and handler for `/graphql`
  - `models/` — request/response models
    - `drink.go` — `Drink` and `EmotionFit` models
    - `booking.go` — `Booking` and `BookingItem` models
  - `services/`
    - `reco_score.go` — scoring logic used by recommendations

If you want, I can add a `scripts/start-dev.sh` to start Mongo + server and create a small `docker-compose.yml` for convenience. Tell me which you prefer and I will add it.
