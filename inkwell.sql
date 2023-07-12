-- 创建用户表
CREATE TABLE users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(50) NOT NULL,
    email VARCHAR(255) NOT NULL,
    nickname VARCHAR(50) NOT NULL,
    avatar VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by INT,
    updated_by INT,
    is_deleted BOOLEAN DEFAULT 0,
    is_admin BOOLEAN DEFAULT 0
);

-- 创建标签表
CREATE TABLE tags (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by INT,
    updated_by INT,
    is_deleted BOOLEAN DEFAULT 0
);

-- 创建分类表
CREATE TABLE categories (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    cover_img_url VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by INT,
    updated_by INT,
    is_deleted BOOLEAN DEFAULT 0
);

-- 创建文章表
CREATE TABLE articles (
    id INT PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    cover_img_url VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by INT,
    updated_by INT,
    is_deleted BOOLEAN DEFAULT 0
);

-- 创建文章标签关联表
CREATE TABLE article_tags (
    article_id INT,
    tag_id INT,
    PRIMARY KEY (article_id, tag_id),
    FOREIGN KEY (article_id) REFERENCES articles (id) ON DELETE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES tags (id) ON DELETE CASCADE
);

-- 创建文章分类关联表
CREATE TABLE article_categories (
    article_id INT,
    category_id INT,
    PRIMARY KEY (article_id, category_id),
    FOREIGN KEY (article_id) REFERENCES articles (id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES categories (id) ON DELETE CASCADE
);


INSERT INTO users (username, password, email, nickname, avatar, created_by, updated_by)
VALUES
    ('user1', 'password1', 'user1@example.com', 'User 1', 'avatar1.jpg', 1, 1),
    ('user2', 'password2', 'user2@example.com', 'User 2', 'avatar2.jpg', 2, 2),
    ('user3', 'password3', 'user3@example.com', 'User 3', 'avatar3.jpg', 3, 3);


INSERT INTO tags (name, description, created_by, updated_by)
VALUES
    ('tag1', 'Tag 1', 1, 1),
    ('tag2', 'Tag 2', 2, 2),
    ('tag3', 'Tag 3', 3, 3);


INSERT INTO categories (name, description, cover_img_url, created_by, updated_by)
VALUES
    ('category1', 'Category 1', 'cover1.jpg', 1, 1),
    ('category2', 'Category 2', 'cover2.jpg', 2, 2),
    ('category3', 'Category 3', 'cover3.jpg', 3, 3);


INSERT INTO articles (title, description, cover_img_url, content, created_by, updated_by)
VALUES
    ('Article 1', 'Description for Article 1', 'cover1.jpg', 'Content for Article 1', 1, 1),
    ('Article 2', 'Description for Article 2', 'cover2.jpg', 'Content for Article 2', 2, 2),
    ('Article 3', 'Description for Article 3', 'cover3.jpg', 'Content for Article 3', 3, 3);


INSERT INTO article_tags (article_id, tag_id)
VALUES
    (1, 1),
    (1, 2),
    (2, 2),
    (3, 3);


INSERT INTO article_categories (article_id, category_id)
VALUES
    (1, 1),
    (1, 2),
    (2, 2),
    (3, 3);