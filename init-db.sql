-- Create database
CREATE USER clean-arquitecture-go WITH PASSWORD 'password';

CREATE SCHEMA clean-arquitecture-go;

GRANT ALL PRIVILEGES ON SCHEMA clean-arquitecture-go TO clean-arquitecture-go;

-- Example
-- Create user database
-- CREATE SCHEMA user_microservice;
-- GRANT ALL PRIVILEGES ON SCHEMA user_microservice TO clean-arquitecture-go;
