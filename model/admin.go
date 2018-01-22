package model

// Admin is a list of admins for this Team
var Admins []Admin

type Admin struct {
	ID       string
	RealName string
}
