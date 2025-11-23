# LeBlanc API Integration Guide

This project now supports both **REST API** and **GraphQL** for frontend-backend communication.

## Backend (Go Server)

### Available Endpoints

#### REST API Endpoints
- `GET /` - Health check
- `GET /drinks` - Get all drinks
- `POST /reco/from-features` - Get drink recommendations
- `POST /bookings` - Create a booking
- `POST /auth/register` - Register new user
- `POST /auth/login` - Login user

#### GraphQL Endpoint
- `POST /graphql` - GraphQL endpoint for queries and mutations
- `GET /graphql` - GraphQL info

### Starting the Backend

```bash
cd server
go run main.go
```

Server will start on `http://localhost:4000`

### GraphQL Schema

The GraphQL API supports the following operations:

**Queries:**
- `drinks` - Get all drinks
- `drink(id)` - Get a specific drink
- `users` - Get all users
- `bookings` - Get all bookings

**Mutations:**
- `createBooking` - Create a new booking
- `register` - Register a new user
- `login` - Login a user
- `recommendFromFeatures` - Get drink recommendations based on emotion fit

### Example GraphQL Queries

**Get all drinks:**
```graphql
query {
  drinks {
    _id
    name
    price
    tags
    caffeine
    temp
    sweetness
    desc
    emotionFit {
      calm
      happy
      stressed
      sad
      adventurous
    }
  }
}
```

**Create a booking:**
```graphql
mutation {
  createBooking(input: {
    name: "John Doe"
    phone: "123-456-7890"
    time: "2025-11-23T10:00:00Z"
    items: [
      {
        drinkId: "507f1f77bcf86cd799439011"
        qty: 2
        options: "{\"size\": \"large\"}"
      }
    ]
    channel: "web"
  }) {
    _id
    name
    phone
    time
  }
}
```

**Get recommendations:**
```graphql
mutation {
  recommendFromFeatures(
    emotionFit: {
      calm: 0.8
      happy: 0.6
      stressed: 0.2
      sad: 0.1
      adventurous: 0.5
    }
    caffeine: "med"
    temp: "iced"
    sweetness: 7
  ) {
    drinkId
    score
  }
}
```

## Frontend (Vue.js)

### Configuration

The frontend can use either REST API or GraphQL. Configure in `.env`:

```env
VITE_API_BASE=http://localhost:4000
VITE_USE_GRAPHQL=false  # Set to 'true' to use GraphQL
```

### Starting the Frontend

```bash
cd website/LeBlanc\ web
npm install
npm run dev
```

### API Usage in Vue Components

The unified API automatically switches between REST and GraphQL based on `VITE_USE_GRAPHQL`:

```javascript
import { getDrinks, createBooking, registerUser, loginUser, recoFromFeatures } from '@/api'

// Get drinks (works with both REST and GraphQL)
const drinks = await getDrinks()

// Create booking
const booking = await createBooking({
  name: 'John Doe',
  phone: '123-456-7890',
  time: new Date().toISOString(),
  items: [
    { drinkId: 'drink_id_here', qty: 2, options: { size: 'large' } }
  ],
  channel: 'web'
})

// Register user
const result = await registerUser({
  name: 'username',
  email: 'user@example.com',
  password: 'password123'
})

// Login
const loginResult = await loginUser({
  name: 'username',
  password: 'password123'
})

// Get recommendations
const recommendations = await recoFromFeatures({
  emotionFit: {
    calm: 0.8,
    happy: 0.6,
    stressed: 0.2,
    sad: 0.1,
    adventurous: 0.5
  },
  caffeine: 'med',
  temp: 'iced',
  sweetness: 7
})
```

### Using GraphQL Directly

If you want to use GraphQL directly regardless of the configuration:

```javascript
import { 
  getDrinksGraphQL, 
  createBookingGraphQL,
  registerUserGraphQL,
  loginUserGraphQL,
  recoFromFeaturesGraphQL
} from '@/api'

// Use GraphQL functions directly
const drinks = await getDrinksGraphQL()
```

### Using REST API Directly

If you want to use REST API directly:

```javascript
import { 
  getDrinksREST, 
  createBookingREST,
  registerUserREST,
  loginUserREST,
  recoFromFeaturesREST
} from '@/api'

// Use REST functions directly
const drinks = await getDrinksREST()
```

## Architecture

### Backend Structure
```
server/
├── main.go                 # Entry point with REST & GraphQL routes
├── internal/
│   ├── db/
│   │   └── mongo.go       # MongoDB connection
│   ├── graph/
│   │   ├── schema.graphql # GraphQL schema definition
│   │   ├── resolver.go    # GraphQL resolvers
│   │   └── handler.go     # GraphQL HTTP handler
│   ├── handlers/
│   │   ├── drinks.go      # REST handlers
│   │   ├── users.go
│   │   ├── bookings.go
│   │   └── reco.go
│   ├── models/
│   │   ├── drink.go
│   │   ├── user.go
│   │   └── booking.go
│   └── services/
│       └── reco_score.go  # Recommendation scoring logic
```

### Frontend Structure
```
website/LeBlanc web/src/
├── api.js                 # Unified API (REST + GraphQL)
├── graphql.js             # GraphQL client & queries
├── views/
│   ├── Home.vue
│   ├── Menu.vue
│   ├── Booking.vue
│   ├── MoodBooker.vue
│   ├── Login.vue
│   └── Register.vue
```

## Features

✅ **Dual API Support** - Use REST or GraphQL
✅ **Unified Interface** - Same function calls, different backends
✅ **CORS Enabled** - Cross-origin requests supported
✅ **MongoDB Integration** - Persistent data storage
✅ **Authentication** - User registration and login
✅ **Drink Recommendations** - AI-powered mood-based suggestions
✅ **Booking System** - Create and manage bookings

## Testing GraphQL

You can test GraphQL queries using tools like:
- **GraphQL Playground** - Visit `http://localhost:4000/graphql`
- **Postman** - Send POST requests to `http://localhost:4000/graphql`
- **curl**:
```bash
curl -X POST http://localhost:4000/graphql \
  -H "Content-Type: application/json" \
  -d '{"query": "{ drinks { _id name price } }"}'
```

## Switching Between REST and GraphQL

To switch from REST to GraphQL:
1. Update `.env` in the frontend:
   ```env
   VITE_USE_GRAPHQL=true
   ```
2. Restart the frontend dev server
3. All API calls will now use GraphQL

No code changes needed! The unified API handles the switch automatically.

## Benefits of GraphQL

- **Flexible Queries** - Request only the data you need
- **Single Endpoint** - One URL for all operations
- **Strongly Typed** - Schema validation built-in
- **Real-time Updates** - Easy to add subscriptions later
- **Better Performance** - Reduce over-fetching and under-fetching

## Benefits of REST

- **Simple & Familiar** - Standard HTTP methods
- **Caching** - Built-in HTTP caching support
- **Debugging** - Easy to test with browser tools
- **Stateless** - Each request is independent
