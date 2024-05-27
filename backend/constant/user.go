package constant

const (
	USER_ROLE_MAKER    = "MAKER"
	USER_ROLE_APPROVER = "APPROVER"
)

var USER_ROLES = map[string]bool{
	USER_ROLE_MAKER:    true,
	USER_ROLE_APPROVER: true,
}
