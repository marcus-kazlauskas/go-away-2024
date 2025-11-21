\c "postgres"

CREATE USER "aoc-2024" WITH PASSWORD 'aoc-2024';

CREATE DATABASE "advent-of-code";

GRANT ALL PRIVILEGES ON DATABASE "advent-of-code" TO "aoc-2024";

\c "advent-of-code"

CREATE EXTENSION IF NOT EXISTS "uuid-ossp"
