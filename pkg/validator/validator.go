package validator

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

func Required(value string, fieldName string) error {
	if strings.TrimSpace(value) == "" {
		return fmt.Errorf("%s is required", fieldName)
	}
	return nil
}

func IsValidEmail(email string) error {
	if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
		return fmt.Errorf("invalid email format")
	}
	return nil
}

func IsValidID(value string) error {
	if _, err := strconv.ParseUint(value, 10, 64); err != nil {
		return fmt.Errorf("invalid id format")
	}
	return nil
}

func ParsePagination(params url.Values) (limit int, offset int, err error) {
	limit = 10
	offset = 0

	if l := params.Get("limit"); l != "" {
		limit, err = strconv.Atoi(l)
		if err != nil || limit < 1 || limit > 100 {
			return 0, 0, fmt.Errorf("invalid limit parameter")
		}
	}

	if p := params.Get("page"); p != "" {
		page, parseErr := strconv.Atoi(p)
		if parseErr != nil || page < 1 {
			return 0, 0, fmt.Errorf("invalid page parameter")
		}
		offset = (page - 1) * limit
	}

	return limit, offset, nil
}
