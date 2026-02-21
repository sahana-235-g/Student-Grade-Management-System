Capstone Project 3: Student Grade Management System API

A complete Go (Golang) REST API for managing student grades similar to university portals (Canvas / Blackboard).
The system supports Admin, Teacher, and Student roles, course & grade management, GPA calculation, and statistics.

1. Design and Implementation

1.1 Approach

RESTful API using Gin framework

MongoDB for storage

Authentication

Role-Based Access Control (RBAC)

Layered architecture: controllers → services → config → database

Middleware for JWT + role authorization

Seed script for sample users, courses, enrollments, grades

1.2 Architecture
Client (React / Postman)
        ↓
Gin Router (CORS, Routes)
        ↓
Middleware (JWT Auth → RBAC)
        ↓
Controllers (auth, user, course, grade, stats)
        ↓
Services (GPA, Course Avg) / Config (MongoDB)
        ↓
MongoDB (users, courses, enrollments, grades)

1.3 Key Design Decisions
Decision	Reason
Gin Framework	Lightweight, fast, simple REST APIs
MongoDB	Flexible schema, natural fit for documents
JWT Auth	Stateless, secure API authentication
RBAC Middleware	Easy role-based permission enforcement
Separate Services	Clean controller code, reusable logic

1.4 Project Structure
student-grade-api/
├── cmd/main.go                # Entry point
├── config/config.go           # MongoDB connection
├── controllers/               # All HTTP controllers
├── middleware/                # JWT, Role-based Auth
├── models/                    # User, Course, Grade, Enrollment
├── services/                  # GPA, Course Average logic
├── seed/seed.go               # Seed sample data
├── utils/                     # JWT & password hashing
├── openapi.yaml               # API Docs
├── postman_collection.json    # Postman Collection
├── PROMPTS.md                 # Prompts used for AI
└── student-grade-frontend/    # Optional React UI

2. Database Schema and Relationships
Collection	Purpose	Key Fields
users	Stores admin / teacher / student	_id, name, email, password, role
courses	Stores courses	_id, name, code, teacher_id
enrollments	Student-course mapping	_id, student_id, course_id
grades	Course grades	_id, student_id, course_id, grade
Relationships

User (teacher) → teaches → Courses

User (student) → enrolled in → Courses

User + Course → has → Grade

3. Role Permissions Matrix
Endpoint	Method	Admin	Teacher	Student
/auth/register	POST	✓	✓	✓
/auth/login	POST	✓	✓	✓
/api/users	GET	✓	—	—
/api/users	POST	✓	—	—
/api/courses	GET	✓	✓	✓
/api/courses	POST	✓	✓	—
/api/courses/:id/enroll	POST	✓	✓	—
/api/grades	POST	—	✓	—
/api/grades	PUT	—	✓	—
/api/grades/:studentId	GET	✓	✓	✓
/api/stats/gpa/:studentId	GET	✓	✓	✓
/api/stats/course-avg/:courseId

4. Seed Data Script

Location: seed/seed.go

Usage:

Uncomment seed.RunSeed() in cmd/main.go

Run the project once

Comment it again to avoid duplication

5. API Documentation

OpenAPI 3.0 Spec: openapi.yaml

Postman Collection: postman_collection.json

All protected endpoints require:

Authorization: Bearer <token>

6. Setup Guide
Prerequisites

Go 1.21+
MongoDB
(Optional) Node.js 18+ for frontend

7. Deliverables Checklist
Deliverable	Status
Complete REST API with proper structure	✓
Database schema with relationships	✓
Seed data script	✓
API documentation (OpenAPI / Postman)	✓
README file with permissions & setup	✓
8. AI Tools and Prompts

All prompts used are listed in:
📄 PROMPTS.md
