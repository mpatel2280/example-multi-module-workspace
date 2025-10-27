-- Migration: Create projects table with seed data
-- Description: Creates the projects table and populates it with sample data
-- Created: 2025-10-27

-- Drop Database
DROP DATABASE IF EXISTS testdb;

-- Create Database
CREATE DATABASE IF NOT EXISTS testdb;

-- Use Database
USE testdb;

-- Create the projects table
CREATE TABLE IF NOT EXISTS projects (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    tech_stack VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Insert seed data
INSERT INTO projects (name, tech_stack) VALUES
    ('E-Commerce Platform', 'Go, PostgreSQL, React'),
    ('Real-time Chat Application', 'Go, WebSocket, Vue.js'),
    ('Data Analytics Dashboard', 'Go, MySQL, Angular'),
    ('Mobile API Gateway', 'Go, gRPC, Kubernetes'),
    ('Content Management System', 'Go, MongoDB, Next.js'),
    ('IoT Device Manager', 'Go, MQTT, Flutter'),
    ('Microservices Framework', 'Go, Docker, Kubernetes'),
    ('Machine Learning Pipeline', 'Go, Python, TensorFlow'),
    ('Payment Processing System', 'Go, MySQL, React'),
    ('Social Media Backend', 'Go, Redis, PostgreSQL');

