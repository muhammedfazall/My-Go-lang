-- E-commerce Schema Design (DDL)

-- 1. SWITCH to the database (if not already connected)
\c ecom_db;

-- 2. CREATE the USERS table (One side of the 1:N relationship)
CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    created_at TIMESTAMP
);

-- 3. CREATE the PRODUCTS table (One of the M:N entities)
CREATE TABLE products (
    product_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price NUMERIC(10, 2) NOT NULL,
    stock_quantity INT NOT NULL
);

-- 4. CREATE the ORDERS table (The 'Many' side, establishing 1:N with USERS)
CREATE TABLE orders (
    order_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    order_date TIMESTAMP NOT NULL,
    status VARCHAR(50),
    -- Define the Foreign Key to the users table
    FOREIGN KEY (user_id) REFERENCES users (user_id)
);

-- 5. CREATE the ORDER_ITEMS table (The Junction Table, resolving M:N)
CREATE TABLE order_items (
    order_item_id SERIAL PRIMARY KEY,
    order_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity INT NOT NULL,
    price_at_sale NUMERIC(10, 2) NOT NULL,
    
    -- Define Foreign Key to orders table
    FOREIGN KEY (order_id) REFERENCES orders (order_id),
    
    -- Define Foreign Key to products table
    FOREIGN KEY (product_id) REFERENCES products (product_id)
);