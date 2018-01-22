package model

// Admins is a list of admins for this Team
var Admins []Admin

// Admin is the model representing a Team admin
type Admin struct {
	ID       string
	RealName string
}
