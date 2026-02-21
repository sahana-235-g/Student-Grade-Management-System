import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";
import Login from "./pages/Login";
import Register from "./pages/Register";
import ListUsers from "./pages/ListUsers";
import Admin from "./pages/Admin";
import Teacher from "./pages/Teacher";
import Student from "./pages/Student";
import AdminDashboard from "./pages/AdminDashboard";
import TeacherDashboard from "./pages/TeacherDashboard";
import StudentDashboard from "./pages/StudentDashboard";
import CreateCourse from "./pages/CreateCourse";
import EnrollStudent from "./pages/EnrollStudent";
import AssignGrade from "./pages/AssignGrade";
import ViewGrades from "./pages/ViewGrades";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Login />} />
        <Route path="/register" element={<Register />} />
        <Route path="/admin" element={<Admin />} />
        <Route path="/teacher" element={<Teacher />} />
        <Route path="/student" element={<Student />} />
        <Route path="/admin/dashboard" element={<AdminDashboard />} />
        <Route path="/teacher/dashboard" element={<TeacherDashboard />} />
        <Route path="/student/dashboard" element={<StudentDashboard />} />
        <Route path="/create-course" element={<CreateCourse />} />
        <Route path="/list-users" element={<ListUsers />} />
        <Route path="/enroll-student" element={<EnrollStudent />} />
        <Route path="/assign-grade" element={<AssignGrade />} />
        <Route path="/view-grades" element={<ViewGrades />} />
        <Route path="*" element={<Navigate to="/" replace />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
