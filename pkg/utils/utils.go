package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math"
	mathRand "math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
)

func Find[T any](slice []T, f func(T) bool) (T, bool) {
	for _, item := range slice {
		if f(item) {
			return item, true
		}
	}
	var zeroT T
	return zeroT, false
}

func Filter[T any](slice []T, f func(T) bool) []T {
	var result []T
	for _, item := range slice {
		if f(item) {
			result = append(result, item)
		}
	}
	return result
}

func GenerateUUID() string {
	return uuid.New().String()
}

func Contains[T comparable](arr []T, value T) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

func ConvertToSlug(text string) string {
	text = strings.ToLower(text)
	text = strings.ReplaceAll(text, " ", "-")
	reg := regexp.MustCompile("[^a-z0-9-]+")
	text = reg.ReplaceAllString(text, "")
	text = strings.Trim(text, "-")
	reg = regexp.MustCompile("-{2,}")
	text = reg.ReplaceAllString(text, "-")
	return text
}

func GenerateUniqueCode() (string, error) {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func GenerateCompactTicketNumber(inputNumber int) string {
	timestamp := time.Now().UnixNano() % 10000
	randomBytes := make([]byte, 2)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic("Error generating ticket bytes")
	}
	randomPart := hex.EncodeToString(randomBytes)
	return fmt.Sprintf("%04d%d%s", inputNumber%10000, timestamp, randomPart)[:10]
}

func GenerateToken(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := mathRand.New(mathRand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func RoundToTwoDecimals(value float64) float64 {
	return math.Round(value*100) / 100
}

func ExtractDigits(text string) string {
	reg := regexp.MustCompile(`[^0-9]+`)
	return reg.ReplaceAllString(text, "")
}

// MaskURL oculta las credenciales en una cadena de conexión para logging
func MaskURL(url string) string {
	if url == "" {
		return ""
	}
	// Buscar la posición de @
	atIndex := strings.Index(url, "@")
	if atIndex == -1 {
		return url // No hay credenciales o formato desconocido
	}

	// Buscar el inicio de las credenciales después de ://
	prefixIndex := strings.Index(url, "://")
	if prefixIndex == -1 {
		return "***@" + url[atIndex+1:]
	}
	prefixIndex += 3 // Saltar ://

	return url[:prefixIndex] + "***" + url[atIndex:]
}
