# Luxury Maxg95 Application

## Overview

This README provides information about the Luxury website, a Go-based backend serving as the foundation for a web service related to interior designs business. The application is designed to handle requests, insert data into a PostgreSQL database, send emails using SMTP. Additionally, it is secured with Let's Encrypt HTTPS and operates under the domain [https://luxurydesign.com.ua](https://luxurydesign.com.ua).

## Features

- **Home Page**: The application provides a home page accessible via a GET request.

- **Insert Request**: Allows the insertion of requests into the PostgreSQL database, ensuring valid input.

- **Email Notification**: After successfully inserting a request, the application sends an email notification using SMTP.

## Configuration

The application uses a configuration struct defined in the config type. Customize the configuration by setting corresponding environment variables.

## Dependencies

The application relies on external packages and modules. Main dependencies are listed in the go.mod file.

## Security

The application includes middleware functions for recovering from panics and setting security headers, enhancing robustness and security.

## TLS Configuration

The server uses TLS with a Let's Encrypt certificate.

## Shutdown

The application gracefully handles shutdown signals, allowing for clean connections closure.