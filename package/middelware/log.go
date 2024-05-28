package middelware

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
)

func RequestLog(c *fiber.Ctx) error {
	userAgent := c.Get("User-Agent")
	ip := c.IP()
	timeZone := c.Get("Time-Zone") // Note: This depends on the client sending the timezone in the headers.
	timestamp := time.Now().Format(time.RFC3339)

	log.Printf("Timestamp: %s, IP: %s, User-Agent: %s, Time-Zone: %s\n", timestamp, ip, userAgent, timeZone)

	return c.Next()
}
