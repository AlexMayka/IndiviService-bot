package core

type FMS interface {
	Get(userID int) string
	Set(userID int, state string)
}
