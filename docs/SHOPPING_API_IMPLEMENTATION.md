# Shopping API Implementation

This document describes the implementation of the Shopping API for the ESDC backend.

## Overview

The Shopping API provides endpoints for managing products, shopping carts, wishlists, and orders. It follows the existing project architecture with separation of concerns across models, repositories, services, handlers, and routes.

## Project Structure

```
internal/
├── model/
│   ├── product.go        # Product model
│   ├── cart.go          # Cart and CartItem models
│   ├── wishlist.go      # Wishlist and WishlistItem models
│   └── order.go         # Order and OrderItem models
├── repository/
│   ├── product_repository.go    # Product database operations
│   ├── cart_repository.go       # Cart database operations
│   ├── wishlist_repository.go   # Wishlist database operations
│   └── order_repository.go      # Order database operations
├── service/
│   ├── product_service.go       # Product business logic
│   ├── cart_service.go          # Cart business logic (stock validation)
│   ├── wishlist_service.go      # Wishlist business logic
│   └── order_service.go         # Order business logic (checkout)
├── handler/
│   ├── product_handler.go       # Product HTTP handlers
│   ├── cart_handler.go          # Cart HTTP handlers
│   ├── wishlist_handler.go      # Wishlist HTTP handlers
│   └── order_handler.go         # Order HTTP handlers
└── routes/
    ├── product_routes.go        # Product route definitions
    ├── cart_routes.go           # Cart route definitions
    ├── wishlist_routes.go       # Wishlist route definitions
    ├── order_routes.go          # Order route definitions
    └── routes.go                # Main route registration (updated)
```

## API Endpoints

### Public Endpoints (No Authentication Required)

- `GET /api/products` - Get all products with filtering and pagination
- `GET /api/products/:id` - Get a single product by ID

### Protected Endpoints (Require JWT Authentication)

**Cart:**
- `GET /api/cart` - Get user's cart
- `POST /api/cart` - Add item to cart
- `PUT /api/cart/:id` - Update cart item quantity
- `DELETE /api/cart/:id` - Remove item from cart
- `DELETE /api/cart` - Clear entire cart

**Wishlist:**
- `GET /api/wishlist` - Get user's wishlist
- `POST /api/wishlist` - Add item to wishlist
- `DELETE /api/wishlist/:id` - Remove item from wishlist

**Orders:**
- `GET /api/orders` - Get user's orders with pagination
- `GET /api/orders/:id` - Get a specific order by ID
- `POST /api/orders` - Create new order (checkout)

## Database Schema

The following tables are created:

1. **products** - Product catalog
2. **cart** - User shopping carts (with unique constraint on user_id + product_id)
3. **wishlist** - User wishlists (with unique constraint on user_id + product_id)
4. **orders** - Order headers
5. **order_items** - Order line items

See `database/shopping_schema.sql` for the complete SQL schema.

## Features Implemented

### Product Service
- Product listing with filtering by category
- Search functionality (name and description)
- Pagination support
- Individual product retrieval

### Cart Service
- Add products to cart with quantity
- Automatic quantity update if product already in cart
- Stock validation before adding/updating
- Cart item updates and deletions
- Clear entire cart
- Calculate subtotals and total

### Wishlist Service
- Add products to wishlist
- Duplicate detection (409 Conflict response)
- Remove items from wishlist
- View complete wishlist with product details

### Order Service
- Create orders from cart items
- Stock validation during checkout
- Calculate order totals
- Order history with pagination
- Individual order retrieval
- User authorization check (users can only view their own orders)

## Error Handling

All endpoints return consistent JSON responses:

**Success Response:**
```json
{
  "success": true,
  "data": { ... },
  "message": "Optional message"
}
```

**Error Response:**
```json
{
  "success": false,
  "error": "Error type",
  "details": "Detailed error message"
}
```

**HTTP Status Codes:**
- 200 OK - Successful operation
- 400 Bad Request - Invalid input or insufficient stock
- 401 Unauthorized - Missing or invalid JWT token
- 404 Not Found - Resource not found
- 409 Conflict - Duplicate entry (e.g., already in wishlist)
- 500 Internal Server Error - Server-side error

## Setup Instructions

### 1. Database Migration

Run the SQL migration:
```bash
sqlite3 database/db.db < database/shopping_schema.sql
```

Or let GORM auto-migrate (already configured in `internal/initializer/init_db.go`):
- The tables will be created automatically when the application starts

### 2. Build and Run

```bash
# Install dependencies (if not already installed)
go mod tidy

# Run the application
go run main.go
```

### 3. Test the API

Use the provided sample data or create your own products:

**Create a product (you'll need admin access):**
```bash
curl -X POST http://localhost:8080/api/products \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Arduino Uno",
    "description": "Arduino Uno R3 microcontroller board",
    "price": 25.99,
    "image": "https://example.com/arduino-uno.jpg",
    "category": "Hardware",
    "stock": 50
  }'
```

**Get all products:**
```bash
curl http://localhost:8080/api/products
```

**Add to cart:**
```bash
curl -X POST http://localhost:8080/api/cart \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "product_id": 1,
    "quantity": 2
  }'
```

**Get cart:**
```bash
curl http://localhost:8080/api/cart \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

**Create order:**
```bash
curl -X POST http://localhost:8080/api/orders \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "items": [
      {
        "product_id": 1,
        "quantity": 2
      }
    ]
  }'
```

## Authentication

All protected endpoints require a JWT token in the Authorization header:
```
Authorization: Bearer <your_jwt_token>
```

The JWT middleware extracts the `user_id` and makes it available to the handlers via the Gin context.

## Security Considerations

1. **Stock Validation** - Stock is checked before adding to cart and creating orders
2. **User Authorization** - Users can only access their own carts, wishlists, and orders
3. **Duplicate Prevention** - Unique constraints prevent duplicate cart and wishlist entries
4. **Input Validation** - All requests are validated using Gin's binding
5. **Cascade Deletion** - Foreign key constraints ensure data integrity

## Future Enhancements

Potential improvements:
- Admin endpoints to manage products (CRUD operations)
- Update stock after order creation
- Order status updates (pending → processing → completed)
- Payment integration
- Product reviews and ratings
- Advanced filtering (price range, sort by price/popularity)
- Inventory management
- Product images upload
- Order cancellation
- Email notifications

## Dependencies

No new dependencies required. The implementation uses:
- `github.com/gin-gonic/gin` (already installed)
- `gorm.io/gorm` (already installed)
- `gorm.io/driver/sqlite` (already installed)

## Testing

Manual testing can be done using:
- cURL (command line)
- Postman (GUI)
- Thunder Client (VS Code extension)
- Any HTTP client

Sample requests are provided in the Testing section above.

## Notes

- The implementation uses SQLite (as per your existing setup)
- Prices are stored as `float64` (GORM maps to DECIMAL in SQL)
- All timestamps are handled automatically by GORM
- The unique constraints on cart and wishlist prevent duplicates at the database level
- Product routes are registered BEFORE the JWT middleware, making them public
- Cart, wishlist, and order routes are registered AFTER the JWT middleware, making them protected
