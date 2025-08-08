# Requirements Document

## Introduction

This feature implements a comprehensive order management system for a medical platform that handles order creation, payment processing, and inventory management. The system manages the complete order lifecycle from creation through payment to fulfillment, with proper inventory tracking and status updates.

## Requirements

### Requirement 1

**User Story:** As a patient, I want to create an order for medications, so that I can purchase the drugs I need with proper delivery information.

#### Acceptance Criteria

1. WHEN a patient submits order information THEN the system SHALL create a new order record with a unique order number
2. WHEN creating an order THEN the system SHALL capture patient information (user_id, user_name, user_phone)
3. WHEN creating an order THEN the system SHALL capture doctor information (doctor_id, doctor_name) 
4. WHEN creating an order THEN the system SHALL capture delivery address information (address_id, address_detail)
5. WHEN creating an order THEN the system SHALL calculate and store the total amount
6. WHEN an order is created THEN the system SHALL set the initial status to "待支付" (pending payment)
7. WHEN an order is created THEN the system SHALL generate timestamps for creation

### Requirement 2

**User Story:** As a patient, I want to pay for my order using different payment methods, so that I can complete my purchase conveniently.

#### Acceptance Criteria

1. WHEN a patient initiates payment THEN the system SHALL support multiple payment types (微信, 支付宝, 银行卡)
2. WHEN payment is initiated THEN the system SHALL validate the order exists and is in "待支付" status
3. WHEN payment is successful THEN the system SHALL update the order status to "已支付"
4. WHEN payment is successful THEN the system SHALL record the payment time
5. WHEN payment is successful THEN the system SHALL record the payment method used
6. IF payment fails THEN the system SHALL maintain the order in "待支付" status

### Requirement 3

**User Story:** As a system administrator, I want inventory to be automatically updated after successful payment, so that stock levels remain accurate.

#### Acceptance Criteria

1. WHEN payment is successful THEN the system SHALL reduce drug inventory quantities based on order items
2. WHEN reducing inventory THEN the system SHALL validate sufficient stock is available
3. IF insufficient stock is available THEN the system SHALL prevent inventory reduction and notify the user
4. WHEN inventory is reduced THEN the system SHALL update the drug stock quantities atomically
5. WHEN inventory update fails THEN the system SHALL rollback the payment status change

### Requirement 4

**User Story:** As a pharmacy staff member, I want to track order status changes throughout the fulfillment process, so that I can manage order processing efficiently.

#### Acceptance Criteria

1. WHEN payment is completed THEN the system SHALL allow status update to "配药中"
2. WHEN drugs are prepared THEN the system SHALL allow status update to "已发货" with send_time
3. WHEN order is delivered THEN the system SHALL allow status update to "已完成" with finish_time
4. WHEN an order needs to be cancelled THEN the system SHALL allow status update to "已取消" with cancel_time
5. WHEN status changes occur THEN the system SHALL record appropriate timestamps
6. WHEN updating status THEN the system SHALL validate the status transition is valid

### Requirement 5

**User Story:** As a patient, I want to view my order details and status, so that I can track my purchase progress.

#### Acceptance Criteria

1. WHEN a patient requests order information THEN the system SHALL return complete order details
2. WHEN displaying orders THEN the system SHALL show current status and relevant timestamps
3. WHEN displaying orders THEN the system SHALL show patient, doctor, and delivery information
4. WHEN displaying orders THEN the system SHALL show payment information and total amount
5. WHEN a patient has multiple orders THEN the system SHALL support order listing with pagination

### Requirement 6

**User Story:** As a system, I want to ensure data consistency during order processing, so that the system maintains integrity under concurrent operations.

#### Acceptance Criteria

1. WHEN processing payments THEN the system SHALL use database transactions to ensure atomicity
2. WHEN updating inventory THEN the system SHALL prevent race conditions through proper locking
3. WHEN order status changes THEN the system SHALL validate state transitions are legal
4. IF any operation in the order-payment-inventory flow fails THEN the system SHALL rollback all related changes
5. WHEN concurrent orders for the same drug occur THEN the system SHALL handle inventory updates safely