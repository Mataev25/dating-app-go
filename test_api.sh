#!/bin/bash

echo "Starting Dating App API Tests..."

echo "1. Health Check:"
curl -s http://localhost:8080/health
echo ""

echo "2. User Registration:"
curl -s -X POST http://localhost:8080/api/v1/users/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test'$(date +%s)'@example.com",
    "password": "test123", 
    "name": "Test User",
    "age": 25,
    "gender": "male",
    "looking_for": "female"
  }'
echo ""

echo "3. Duplicate Email Test:"
curl -s -X POST http://localhost:8080/api/v1/users/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "duplicate@example.com",
    "password": "pass123",
    "name": "First User", 
    "age": 30,
    "gender": "female",
    "looking_for": "male"
  }'

curl -s -X POST http://localhost:8080/api/v1/users/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "duplicate@example.com", 
    "password": "pass456",
    "name": "Second User",
    "age": 35,
    "gender": "male", 
    "looking_for": "female"
  }'
echo ""

echo "4. Age Validation Test:"
echo "Testing underage user (should fail):"
curl -s -X POST http://localhost:8080/api/v1/users/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "young'$(date +%s)'@exampl5e5.com",
    "password": "young123",
    "name": "Young User",
    "age": 16,
    "gender": "male",
    "looking_for": "female"
  }'
echo ""

echo "All tests executed! Check responses above."
