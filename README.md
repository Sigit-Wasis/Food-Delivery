# Food Delivery System

A comprehensive food delivery system designed to manage the complete order process from restaurant menus to delivery.

Building a Scalable Food Delivery System with Go
- https://sigitwasis.medium.com/design-a-food-delivery-system-in-go-9d1049be2ab5

## Modules Overview

### 1. User Module
The User Module is responsible for managing the different users in the system, including:
- **Customers**: Place orders, view order history, track delivery status.
- **Restaurants**: Manage restaurant details, menus, availability, and order processing.
- **Delivery Agents**: Assign orders for delivery, update delivery status.

### 2. Order Module
The Order Module is responsible for handling all order-related actions, including:
- **Order Placement**: Customers can place orders from restaurant menus.
- **Order Status**: Tracks order progress from placement to delivery.
- **Order History**: Maintains a log of all past orders placed by customers.

### 3. Restaurant Module
The Restaurant Module helps restaurants manage their operations:
- **Menu Management**: Restaurants can add, update, or remove items from their menu.
- **Availability**: Restaurants can set their availability for accepting orders.
- **Restaurant Details**: Basic details about the restaurant, including location, contact info, etc.

### 4. Delivery Module
The Delivery Module handles the logistics of the delivery process:
- **Order Assignment**: Automatically assigns orders to available delivery agents.
- **Delivery Tracking**: Tracks the status of delivery and updates the customer in real-time.
- **Delivery Completion**: Marks the order as complete once it has been delivered.

### 5. Admin Module
The Admin Module is the control center of the system, allowing for:
- **System Management**: Admins can manage users, restaurants, and delivery agents.
- **Order Approvals**: Admins can approve or reject orders, if necessary.
- **Reporting**: Generate reports on orders, deliveries, and system usage.

## Technologies
- **Frontend**: React, HTML, CSS
- **Backend**: Node.js, Express
- **Database**: MongoDB / PostgreSQL
- **Authentication**: JWT / OAuth

## High-Level Architecture
Microservices or Modular Monolith.

Clean Architecture:

- **Presentation Layer**: HTTP handlers, gRPC endpoints.
- **Service Layer**: Business logic.
- **Repository Layer**: Data access (e.g., PostgreSQL).
- **Entities**: Core domain objects.

## Checklist Clean Code yang Baik
Prinsip	Status	Catatan

| **Aspek**                     | **Status** | **Keterangan**                                           |
|-------------------------------|------------|---------------------------------------------------------|
| **Separation of Concerns**    | ✅         | Sudah terpisah menjadi handler, service, repository, dan model |
| **Konsistensi Error Handling**| ⚠️         | Pastikan semua error ditangani dengan cara yang sama     |
| **Abstraksi dengan Interface**| ⚠️         | Gunakan interface pada service dan repository           |
| **Middleware**                | ⚠️         | Gunakan middleware untuk validasi                       |
| **Unit Testing**              | ⚠️         | Pastikan setiap lapisan memiliki pengujian              |
| **Logging**                   | ✅         | Logging error di server untuk debugging                 |
| **Status Code yang Tepat**    | ✅         | Gunakan status code yang sesuai untuk setiap jenis respons |


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.