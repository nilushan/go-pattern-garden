-- Create the users table
CREATE TABLE users (
    id VARCHAR(36) PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL
);

-- Insert some sample data for testing
INSERT INTO users (id, email, name, created_at) VALUES
('a1b2c3d4-e5f6-g7h8-i9j0-k1l2m3n4o5p6', 'diana@example.com', 'Diana Prince', NOW()),
('b2c3d4e5-f6g7-h8i9-j0k1-l2m3n4o5p6q7', 'bruce@example.com', 'Bruce Wayne', NOW());
