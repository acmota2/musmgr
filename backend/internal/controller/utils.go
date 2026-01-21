package controller

import "github.com/jackc/pgx/v5/pgtype"

func textOrNull(s *string) pgtype.Text {
	if s != nil {
		return pgtype.Text{
			Valid:  true,
			String: *s,
		}
	}
	return pgtype.Text{
		Valid: false,
	}
}
