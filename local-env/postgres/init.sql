\c "postgres"

CREATE ROLE "aoc-2024" WITH LOGIN PASSWORD 'aoc-2024';

CREATE DATABASE "advent-of-code" WITH OWNER "aoc-2024";

\q
