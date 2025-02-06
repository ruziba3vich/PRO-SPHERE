-- Language Enum Type
-- First check if the enum type exists
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'langs') THEN
        CREATE TYPE langs AS ENUM (
            'uz','ru','en','kk','tk','ky','tj','az','am','by',
            'kg','md','ge','ro','uk','hy','sr','mn','cs','pl','th','vi'
        );
    END IF;
END
$$;

-- Feed Categories Table
CREATE TABLE feed_categories (
    id SERIAL PRIMARY KEY,
    icon_url VARCHAR(2083),
    icon_id VARCHAR(64)
);  

-- Feed Category Translations Table
CREATE TABLE feed_category_translations (
    id SERIAL PRIMARY KEY,
    feed_category_id INT REFERENCES feed_categories(id) ON DELETE CASCADE,
    lang langs NOT NULL,                       -- e.g., 'uz', 'ru', 'en'
    name VARCHAR(64) NOT NULL,
    UNIQUE (feed_category_id, lang)           -- Prevent duplicate translations
);

-- Feeds Table (Base URL Only)
CREATE TABLE feeds (
    id SERIAL PRIMARY KEY,
    priority INT,
    max_items INT,
    base_url VARCHAR(2083) NOT NULL,          -- e.g., 'https://www.gazeta.uz'
    logo_url VARCHAR(2083),
    logo_url_id VARCHAR(64),
    last_refreshed TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- Feed Translations Table
CREATE TABLE feed_translations (
    id SERIAL PRIMARY KEY,
    feed_id INT REFERENCES feeds(id) ON DELETE CASCADE,
    lang langs NOT NULL,                      -- Language of the feed
    title VARCHAR(255) NOT NULL,              -- e.g., "Gazeta UZ News"
    description TEXT,
    UNIQUE (feed_id, lang)                    -- Prevent duplicate translations
);

-- Feed Contents Table
CREATE TABLE feed_contents (  
    id SERIAL PRIMARY KEY,
    feed_id INT REFERENCES feeds(id) ON DELETE CASCADE,
    category_id INT REFERENCES feed_categories(id) ON DELETE CASCADE, -- society
    lang langs NOT NULL,
    link VARCHAR(2083) NOT NULL,              -- e.g., 'https://www.gazeta.uz/uz/rss/?section=society'
    feed_translation_id INT REFERENCES feed_translations(id) ON DELETE CASCADE, -- Reference feed_translations
    UNIQUE (feed_translation_id, category_id) -- Prevent duplicate links for the same translation, category, and language
);

-- Feed Items Table
CREATE TABLE feed_items (
    id SERIAL PRIMARY KEY,
    feed_id INT REFERENCES feeds(id) ON DELETE CASCADE,
    image_url VARCHAR(2083),
    published_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- Feed Item Translations Table
CREATE TABLE feed_item_translations (
    id SERIAL PRIMARY KEY,
    feed_item_id INT REFERENCES feed_items(id) ON DELETE CASCADE,
    lang langs NOT NULL,                          -- Language code
    title VARCHAR(255) NOT NULL,
    description TEXT,
    link VARCHAR(2083),                           -- Link to the original content
    UNIQUE (feed_item_id, lang)                   -- Ensure unique translations
);
