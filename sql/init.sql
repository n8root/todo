DROP TABLE IF EXISTS tasks;

CREATE TABLE tasks(
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    completed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO tasks (title, description, completed)
    VALUES 
    ('Test1', 'Test description1', false),
    ('Test2', 'Test description2', false),
    ('Test3', 'Test description3', false);
