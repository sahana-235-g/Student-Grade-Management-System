package middleware

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func RequireRole(roles ...string) gin.HandlerFunc {
    return func(c *gin.Context) {
        userRole, exists := c.Get("role")
        if !exists {
            c.JSON(http.StatusForbidden, gin.H{"error": "Role not found"})
            c.Abort()
            return
        }

        // Check if user role is allowed
        for _, role := range roles {
            if userRole == role {
                c.Next()
                return
            }
        }

        c.JSON(http.StatusForbidden, gin.H{
            "error": "Access denied. Required roles: " + formatRoles(roles),
        })
        c.Abort()
    }
}

// Helper to show roles in error message
func formatRoles(roles []string) string {
    result := ""
    for i, r := range roles {
        if i > 0 {
            result += ", "
        }
        result += r
    }
    return result
}