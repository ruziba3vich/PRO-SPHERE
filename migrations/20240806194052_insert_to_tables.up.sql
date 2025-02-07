
CREATE TABLE posts (
    post_id UUID PRIMARY KEY,
    publisher_id VARCHAR(64) NOT NULL,
    post_title VARCHAR(255) NOT NULL,
    post_category VARCHAR(255) NOT NULL,
    post_short_content TEXT NOT NULL,
    post_content TEXT NOT NULL,
    post_featured_image TEXT,
    post_source VARCHAR(255),
    imported_data TEXT,
    views INT DEFAULT 0,
    deleted BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

