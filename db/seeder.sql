-- Create user Table
CREATE TABLE "user" (
    id UUID PRIMARY KEY,                      
    name VARCHAR(255) NOT NULL,            
    email VARCHAR(255) UNIQUE NOT NULL,     
    address TEXT,                             
    created_at TIMESTAMP DEFAULT NOW(),       
    updated_at TIMESTAMP DEFAULT NOW()     
);

CREATE TABLE magic_link (
    id UUID PRIMARY KEY,
    email TEXT NOT NULL,
    token TEXT UNIQUE NOT NULL,
    expired_at TIMESTAMP NOT NULL
);

INSERT INTO "user" (id, name, email, address)
VALUES
    ('5c6f4f3b-5e3e-4b47-a8f5-18b417623a1f', 'John Doe', 'john.doe@example.com', '123 Main St, Springfield'),
    ('2a1d1d58-b6b5-44bc-baf8-6dbe034c4034', 'Jane Smith', 'jane.smith@example.com', '456 Elm St, Metropolis'),
    ('f5d6b0a7-7cc5-4d65-9259-f7ed7c9b5c9a', 'Alice Johnson', 'alice.johnson@example.com', '789 Oak St, Gotham'),
    ('7e92e4b1-86a4-4230-a3ff-d6e4098dbbef', 'Bob Brown', 'bob.brown@example.com', '321 Pine St, Star City'),
    ('c3d6579c-4c71-4263-b634-c8a49f9942a3', 'Charlie White', 'charlie.white@example.com', '654 Birch St, Central City');