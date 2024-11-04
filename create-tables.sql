DROP TABLE IF EXISTS comments;
DROP TABLE IF EXISTS articles;
DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE categories (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    description VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE articles (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    author_id INT NOT NULL,
    category_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (author_id) REFERENCES users(id),
    FOREIGN KEY (category_id) REFERENCES categories(id)
);

CREATE TABLE comments (
    id INT AUTO_INCREMENT PRIMARY KEY,
    article_id INT NOT NULL,
    user_id INT NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (article_id) REFERENCES articles(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Insert dummy users
INSERT INTO users (username, email, password, created_at, updated_at) VALUES
('user1', 'user1@example.com', 'password1', NOW(), NOW()),
('user2', 'user2@example.com', 'password2', NOW(), NOW()),
('user3', 'user3@example.com', 'password3', NOW(), NOW());

-- Insert dummy categories
INSERT INTO categories (name, description, created_at, updated_at) VALUES
('Technology', 'Articles about technology', NOW(), NOW()),
('Health', 'Articles about health and wellness', NOW(), NOW()),
('Lifestyle', 'Articles about lifestyle and hobbies', NOW(), NOW());

-- Insert dummy articles
INSERT INTO articles (title, content, author_id, category_id, created_at, updated_at) VALUES
('First Tech Article', 'Content of the first tech article', 1, 1, NOW(), NOW()),
('Second Health Article', 'Content of the second health article', 2, 2, NOW(), NOW()),
('Third Lifestyle Article', 'Content of the third lifestyle article', 3, 3, NOW(), NOW());

-- Insert dummy comments
INSERT INTO comments (article_id, user_id, content, created_at, updated_at) VALUES
(1, 2, 'Great article on technology!', NOW(), NOW()),
(2, 3, 'Very informative health tips.', NOW(), NOW()),
(3, 1, 'I love this lifestyle piece!', NOW(), NOW());
