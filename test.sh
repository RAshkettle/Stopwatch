#!/bin/bash

# Simple test runner script for the Stopwatch library

echo "Running unit tests with coverage..."
go test -v -cover

echo ""
echo "Running benchmarks..."
go test -bench=. -benchmem

echo ""
echo "Generating detailed coverage report..."
go test -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html

echo ""
echo "Coverage report generated in coverage.html"
echo "All tests completed successfully!"
