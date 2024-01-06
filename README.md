# SQLEase

## Overview
SQLEase is a powerful SQL golang ORM designed to simplify and streamline the process of creating and managing database queries. It offers a robust solution for developers looking to integrate SQL databases into their Go applications with ease and efficiency.

## Features
- Simplified database query management.
- Streamlined integration with SQL databases.
- Easy-to-use ORM functionalities for Go applications.

## Getting Started

### Prerequisites
- Go (version 1.x or newer)
- MySQL or any SQL server set up and running.

### Installation
To install SQLEase, clone the repository and set it up in your Go workspace:


### Database Configuration

#### Setting Up Database Connection
To configure the database connection for SQLEase, you need to create a specific file with connection details:

1. **Create Connection File**: Inside the `database` folder, create a file named `connection_info.go`.

2. **Add Connection Details**:
   Paste the following code into `connection_info.go`:
   ```go
   package database

   // DBConnectionDetails struct holds the database connection parameters
   type DBConnectionDetails struct {
       Username string
       Password string
       Hostname string
       Port     int
       Database string
   }

   // ConnectionString function returns database connection details
   func ConnectionString() DBConnectionDetails {
       return DBConnectionDetails{
           // Replace with your actual database connection details
           Username: "your-username",
           Password: "your-password",
           Hostname: "your-hostname",
           Port:     your-port-number, // Example: 3306
           Database: "your-database-name",
       }
   }
