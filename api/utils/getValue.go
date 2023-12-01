package utils

import "database/sql"

type GetValue struct{}

func (gv *GetValue) NullString(value *string) *sql.NullString {
	if value != nil {
		return &sql.NullString{
			String: *value,
			Valid:  true,
		}
	}

	return &sql.NullString{
		String: "",
		Valid:  false,
	}
}

func (gv *GetValue) NullBool(value *bool) *sql.NullBool {
	if value != nil {
		return &sql.NullBool{
			Bool:  *value,
			Valid: true,
		}
	}

	return &sql.NullBool{
		Bool:  false,
		Valid: false,
	}
}
