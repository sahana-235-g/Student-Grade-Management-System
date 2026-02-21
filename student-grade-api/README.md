# Capstone Project 3: Student Grade Management System API

**Go (Golang) REST API** for managing student grades similar to university portals (Canvas / Blackboard). The system supports **Admin**, **Teacher**, and **Student** roles, course and grade management, and provides academic performance data including **GPA** and **grade statistics**.

---

## 1. Design and Implementation

### 1.1 Approach

- **RESTful API** using **Gin** framework in Go.
- **MongoDB** for persistence with clear collections and relationships.
- **JWT** for authentication; **role-based access control (RBAC)** for authorization.
- **Layered structure:** `controllers` → `services` / `config`; `middleware` for auth and RBAC; `models` for data; `seed` for sample data.

### 1.2 Architecture

```
Client (React / Postman)
        ↓
   Gin Router (CORS, routes)
        ↓
   Middleware: JWT Auth → RBAC (RequireRole)
        ↓
   Controllers (auth, user, course, grade, stats)
        ↓
   Services (GPA, course average) / Config (MongoDB)
        ↓
   MongoDB (users, courses, enrollments, grades)
```

### 1.3 Key Design Decisions

| Decision | Reason |
|----------|--------|
| Gin | Lightweight, fast, common for Go APIs. |
| MongoDB | Flexible schema; documents map well to User, Course, Grade, Enrollment. |
| JWT in Authorization header | Stateless auth; easy to use from frontends and Postman. |
| RBAC middleware | Central place to enforce role per route; clear permission matrix. |
| Separate services for GPA/stats | Reusable logic; controllers stay thin. |

### 1.4 Project Structure

```
student-grade-api/
├── cmd/main.go           # Entry point, route groups, CORS, middleware wiring
├── config/config.go      # MongoDB connection and collection access
├── controllers/          # HTTP handlers (auth, user, course, grade, stats)
├── middleware/           # JWT auth, RequireRole RBAC
├── models/               # User, Course, Grade, Enrollment structs
├── services/             # GPA calculation, course average
├── seed/seed.go          # Seed script (users, courses, enrollments, grades)
├── utils/                # JWT generation/validation, password hashing
├── openapi.yaml          # OpenAPI 3.0 spec
├── postman_collection.json
├── PROMPTS.md            # AI prompts used (if any)
└── student-grade-frontend/   # Optional React UI
```

---

## 2. Database Schema and Relationships

| Collection    | Purpose                    | Main Fields / Relationships      |
|---------------|----------------------------|-----------------------------------|
| **users**     | All users (admin/teacher/student) | `_id`, name, email, password, role |
| **courses**   | Courses                    | `_id`, name, code, credits, `teacher_id` → users._id |
| **enrollments** | Student–course enrollment | `_id`, `student_id` → users._id, `course_id` → courses._id |
| **grades**    | Grade per student per course | `_id`, `student_id`, `course_id`, grade (float64) |

Relationships: **User** (teacher) → **Course**; **User** (student) + **Course** → **Enrollment**; **User** + **Course** → **Grade**.

---

## 3. Role Permissions Matrix

| Endpoint | Method | Admin | Teacher | Student |
|----------|--------|:-----:|:-------:|:-------:|
| `/auth/register` | POST | ✓ | ✓ | ✓ |
| `/auth/login` | POST | ✓ | ✓ | ✓ |
| `/api/users` | GET | ✓ | — | — |
| `/api/users` | POST | ✓ | — | — |
| `/api/courses` | GET | ✓ | ✓ | ✓ |
| `/api/courses` | POST | ✓ | ✓ | — |
| `/api/courses/:id/enroll` | POST | ✓ | ✓ | — |
| `/api/grades` | POST | — | ✓ | — |
| `/api/grades` | PUT | — | ✓ | — |
| `/api/grades/:studentId` | GET | ✓ | ✓ | ✓ |
| `/api/stats/gpa/:studentId` | GET | ✓ | ✓ | ✓ |
| `/api/stats/course-avg/:courseId` | GET | ✓ | ✓ | ✓ |

---

## 4. Seed Data Script

- **Location:** `seed/seed.go`
- **Usage:** Uncomment `seed.RunSeed()` in `cmd/main.go`, run the API once, then comment it back.
- **Sample data:** Users (admin, teacher, student), courses (e.g. MATH101, PHYS101), enrollments, and grades.

**Seed credentials:** admin@test.com / admin123 | teacher@test.com / teacher123 | student@test.com / student123

---

## 5. API Documentation

- **OpenAPI spec:** `openapi.yaml`
- **Postman:** Import `postman_collection.json`. Use Login to set `token` for protected endpoints.

All `/api/*` endpoints require: `Authorization: Bearer <token>`.

---

## 6. Setup Guide

**Prerequisites:** Go 1.21+, MongoDB, (optional) Node.js 18+ for frontend.

1. `git clone <repo>` then `cd student-grade-api`
2. `go mod tidy`
3. Uncomment `seed.RunSeed()` in `cmd/main.go`, run `go run cmd/main.go`, then comment it back
4. Run API: `go run cmd/main.go` → http://localhost:8080
5. Optional frontend: `cd student-grade-frontend && npm install && npm run dev` → http://localhost:5173

---

## 7. Deliverables Checklist

| Deliverable | Status |
|-------------|--------|
| Complete REST API with proper structure | ✓ |
| Database schema with relationships | ✓ |
| Seed data script | ✓ |
| API documentation (OpenAPI + Postman) | ✓ |
| README with role permissions matrix and setup guide | ✓ |

---

## 8. AI Tools and Prompts

If any AI tools were used, the prompts are documented in **[PROMPTS.md](./PROMPTS.md)** for transparency.

---

**Capstone Project 3 – Go (Golang)**  
**Submission:** GitHub repo + shared drive link as per instructions.
