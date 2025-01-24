# Car Rental API Project

This repository contains the implementation of a Car Rental API using Golang and PostgreSQL. The project is divided into two phases, **Car Rental - v1** and **Car Rental - v2**, each with specific features and requirements.

---

## Project Setup

### Prerequisites
- Go (version 1.20 or later)
- PostgreSQL (version 14 or later)
- Git (to clone the repository)

### Installation
1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd car-rental-api
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Configure the database:
   - Set up a PostgreSQL database.
   - Update the database connection string in the `.env` file.

4. Run the application:
   ```bash
   go run main.go
   ```

### Branch Structure
- The **main** branch combines features from both v1 and v2.
- To check a specific version, switch to the respective branch:
  - `v1` branch: Contains only the features from Car Rental - v1.
  - `v2` branch: Contains only the features from Car Rental - v2.

---

## Features

### Car Rental - v1
This phase focuses on the basic API implementation:

#### API Endpoints
1. **Customer Management:**
   - Create, Read, Update, and Delete (CRUD) operations for customers.
2. **Car Management:**
   - CRUD operations for cars.
3. **Booking Management:**
   - CRUD operations for bookings.

#### Additional Requirements
- Create an Entity-Relationship Diagram (ERD) using dummy data.
- Push the project to a GitHub or GitLab repository.

---

### Car Rental - v2
This phase enhances the functionality of the API:

#### New Features
1. **Membership Program:**
   - Customers can opt for a membership program that offers discounts on rentals.
   - Membership tiers:
     - Bronze: 4% discount
     - Silver: 7% discount
     - Gold: 15% discount
   - Discount formula:
     ```
     Discount = (Days_of_Rent * Daily_Car_Rent) * Membership_Discount
     ```
   - Customers who don’t join the membership program won’t receive discounts.

2. **Rent with a Driver:**
   - Customers can choose to rent with or without a driver.
   - Driver data and daily costs are stored in the database.
   - Incentive for drivers:
     ```
     Incentive = (Days_of_Rent * Daily_Car_Rent) * 5%
     ```

#### Database Changes
- Update the ERD to accommodate the new features.
- Add tables for driver data and membership programs.

#### API Changes
1. Extend the Customer and Booking APIs to support memberships, discounts, driver costs, and booking types (Car Only / Car & Driver).
2. Add new endpoints for:
   - **Membership Management:**
     - CRUD operations for membership data.
   - **Driver Management:**
     - CRUD operations for driver data.
     - Incentive calculations.

---

## Dummy Data
Sample data is provided to populate the database for testing purposes. ERD diagrams are created to illustrate relationships between entities.

- [Car Rental V1 Dummy Data](https://docs.google.com/spreadsheets/d/1uhNxSx6k4vvfuz5fuomrHWDSgrcVQkup0mAvnO5quSU/edit?usp=sharing)
- [Car Rental V2 Dummy Data](https://docs.google.com/spreadsheets/d/1a0buC_HyCduG_tXnEwZ4IcQxbKOIfkQbBXaDluu8R8k/edit?usp=sharing)

---

## Contribution
Contributions are welcome! Please follow these steps:
1. Fork the repository.
2. Create a new branch:
   ```bash
   git checkout -b feature-name
   ```
3. Commit your changes:
   ```bash
   git commit -m "Add feature"
   ```
4. Push to the branch:
   ```bash
   git push origin feature-name
   ```
5. Create a pull request.

---

## License
This project is licensed under the MIT License. See the LICENSE file for details.

