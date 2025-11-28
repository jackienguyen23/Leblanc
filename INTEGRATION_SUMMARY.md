# Frontend-Backend Connection Summary

## ✅ What Was Done

### Backend (Go Server)
1. **Fixed Port Configuration**: Updated to use port 4000 consistently
2. **Added GraphQL Support**:
   - Created GraphQL schema (`internal/graph/schema.graphql`)
   - Implemented resolvers for all operations (`internal/graph/resolver.go`)
   - Added GraphQL HTTP handler (`internal/graph/handler.go`)
   - Integrated GraphQL endpoint at `/graphql` in `main.go`
3. **Enhanced Services**: Added `ScoreDrinks` function for recommendation scoring
4. **Maintained REST API**: All existing REST endpoints still work
5. **Fly.io Deployment**: Added `Dockerfile` and `fly.toml` to deploy the Go API (internal port 8080, secrets for Mongo/admin/email, CORS allows the Vercel frontend)

### Frontend (Vue.js)
1. **Fixed API Base URL**: Updated `.env` to use port 4000
2. **Installed GraphQL Client**: Added `graphql` and `graphql-request` packages
3. **Created GraphQL Integration** (`src/graphql.js`):
   - Defined queries for drinks, bookings, users
   - Defined mutations for auth, bookings, recommendations
   - Created wrapper functions for all operations
4. **Updated API Service** (`src/api.js`):
   - Unified interface that works with both REST and GraphQL
   - Toggle between REST/GraphQL via `VITE_USE_GRAPHQL` env variable
   - Backward compatible with existing code
5. **Created Demo Component**: `ApiDemo.vue` to test both APIs
6. **Vercel Hosting**: Added SPA rewrite config (`vercel.json`) and env defaults (`VITE_API_BASE` -> Fly API) for Vercel deployments
7. **Booking Email Confirmation**: Booking form now sends a confirmation email via EmailJS (set `VITE_EMAILJS_BOOKING_TEMPLATE_ID`)

## 🚀 How to Use

### Start Backend
```bash
cd server
go run main.go
```
Server runs on `http://localhost:4000`

### Start Frontend
```bash
cd website/LeBlanc\ web
npm install  # First time only
npm run dev
```
Frontend runs on the Vite dev server (usually `http://localhost:5173`)

## Production Deployments

- API on Fly.io: `https://server-wandering-tree-4946.fly.dev` (deploy with `flyctl deploy --config server/fly.toml --dockerfile server/Dockerfile` and set secrets `MONGO_URI`, `MONGO_DB`, `ADMIN_NAME`, `ADMIN_EMAIL`, `ADMIN_PASSWORD`, `FRONTEND_VERIFY_URL`, `EMAIL_REQUIRE_MX`).
- Frontend on Vercel: `https://le-blanc-web.vercel.app` (deploy from `website/LeBlanc web` with `vercel --prod`; envs `VITE_API_BASE`, `VITE_EMAILJS_*`, `VITE_ADMIN_EMAIL` should point at the Fly API and email service).

## 📋 API Endpoints

### REST API
- `GET /drinks` - List all drinks
- `POST /auth/register` - Register user
- `POST /auth/login` - Login user
- `POST /bookings` - Create booking
- `POST /reco/from-features` - Get recommendations

### GraphQL API
- `POST /graphql` - Single endpoint for all operations
  - Queries: `drinks`, `drink(id)`, `users`, `bookings`
  - Mutations: `register`, `login`, `createBooking`, `recommendFromFeatures`

## 🔄 Switching Between REST and GraphQL

Edit `website/LeBlanc web/.env`:
```env
VITE_USE_GRAPHQL=false  # Use REST API
VITE_USE_GRAPHQL=true   # Use GraphQL
```

The same code works with both! Example:
```javascript
import { getDrinks } from '@/api'
const drinks = await getDrinks()  // Uses REST or GraphQL based on config
```

## 📁 New Files Created

### Backend
- `server/internal/graph/schema.graphql` - GraphQL schema
- `server/internal/graph/resolver.go` - GraphQL resolvers
- `server/internal/graph/handler.go` - GraphQL HTTP handler
- `server/gqlgen.yml` - GraphQL generator config (optional for gqlgen)

### Frontend
- `website/LeBlanc web/src/graphql.js` - GraphQL client & queries
- `website/LeBlanc web/src/views/ApiDemo.vue` - Demo component

### Documentation
- `API_INTEGRATION.md` - Complete integration guide

## 📝 Modified Files

### Backend
- `server/main.go` - Added GraphQL endpoint
- `server/internal/services/reco_score.go` - Added ScoreDrinks function

### Frontend
- `website/LeBlanc web/.env` - Updated port & added GraphQL flag
- `website/LeBlanc web/src/api.js` - Unified REST/GraphQL interface
- `website/LeBlanc web/package.json` - Added graphql dependencies

## 🎯 Key Features

✅ **Dual API Support** - Choose REST or GraphQL
✅ **Zero Code Changes** - Switch APIs via environment variable
✅ **Backward Compatible** - Existing code continues to work
✅ **Type Safety** - GraphQL schema provides validation
✅ **CORS Enabled** - Cross-origin requests supported
✅ **Unified Interface** - Same functions for both APIs

## 🧪 Testing

### Test REST API
```bash
curl http://localhost:4000/drinks
```

### Test GraphQL
```bash
curl -X POST http://localhost:4000/graphql \
  -H "Content-Type: application/json" \
  -d '{"query": "{ drinks { _id name price } }"}'
```

### Test in Browser
Visit the demo page (add route to `ApiDemo.vue` in your router)

## 📚 Next Steps

1. **Add Authentication Tokens**: Implement JWT for secure API access
2. **Add GraphQL Subscriptions**: Real-time updates for bookings
3. **Add Error Handling**: Better error messages and retry logic
4. **Add Loading States**: Show loading indicators in UI
5. **Add Caching**: Cache GraphQL queries for better performance
6. **Add Tests**: Unit and integration tests for both APIs

## 💡 Tips

- Use GraphQL for complex data requirements
- Use REST for simple CRUD operations
- GraphQL reduces over-fetching of data
- REST is easier to cache at HTTP level
- Both APIs access the same MongoDB database
- No data duplication - choose based on your needs!
