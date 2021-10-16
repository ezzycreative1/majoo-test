package response_mapping

const (
	//General
	GeneralError = "G-400"
	Unauthorized = "G-401"
	InvalidParam = "G-001"

	//RegisterInvalidDomain
	RegisterEmailDuplicate = "REG-001"
	RegisterInvalidDomain  = "REG-002"

	//Login
	LoginUserNotFound = "AUTH-001"

	//Forgot Password
	EmailNotRegister = "FP-001"

	//Change Password
	ChangePasswordInvalidParam = "CP-001"

	// Project
	ErrorGetPrj    = "PRJ-001"
	CreateProjects = "PRJ-002"
	DeletedProject = "PRJ-003"
	EditProject    = "PRJ-004"
	ValdiatePrj    = "PRJ-005"

	// User
	AccountSuspend = "USR-001"

	// Organization
	ErrorGetOrg                 = "ORG-001"
	LimitOrg                    = "ORG-002"
	DeactiveStatus              = "ORG-003"
	ErrorEditOrg                = "ORG-004"
	ErrorCreateOrgRole          = "ORG-005"
	EditMemberOrg               = "ORG-006"
	ErrorAddMemberOrg           = "ORG-007"
	ErrorDeleteMemberOrg        = "ORG-008"
	ErrorCreateOrg              = "ORG-009"
	ErrorUpdateOrganizationRole = "ORG-010"
	ErrorDeleteOrganizationRole = "ORG-011"

	//Midelware
	Middleware = "MID-001"

	// Instance
	GetInstance             = "INS-001"
	DeleteInstance          = "INS-002"
	StartInstance           = "INS-003"
	ShutdownInstance        = "INS-004"
	RestartInstance         = "INS-005"
	StatusInstance          = "INS-006"
	CreateInstance   string = "INS-007"

	// Storage
	GetStorage    = "STR-001"
	AddStorage    = "STR-002"
	AttachStorage = "STR-003"
	DetachStorage = "STR-004"
	ResizeStorage = "STR-005"
)

var ResponseMappingID = map[string]string{
	GeneralError: "General error",
	Unauthorized: "Tidak punya akses",
	InvalidParam: "Parameter tidak lengkap / tidak valid.",

	RegisterEmailDuplicate: "Email sudah terdaftar",
	RegisterInvalidDomain:  "Tidak bisa mendaftar dengan domain email ini",

	LoginUserNotFound: "User tidak ditemukan",

	EmailNotRegister: "Email belum terdaftar",

	ChangePasswordInvalidParam: "Password lama atau password baru atau konfirm password diperlukan",
}

var ResponseMappingEN = map[string]string{
	GeneralError: "General error",
	Unauthorized: "Unauthorized",
	InvalidParam: "Incomplete / invalid Parameter(s).",

	RegisterEmailDuplicate: "Email already exist",
	RegisterInvalidDomain:  "Email detected as free email provider domain",

	LoginUserNotFound: "User not found",
	CreateInstance:    "Failed Create Instance",
	EmailNotRegister:  "Email not register",

	ChangePasswordInvalidParam: "Old password or new password or confirm password required",

	// str
	GetStorage:    "Failed get Storage",
	AddStorage:    "Failed add storage",
	AttachStorage: "Failed Attach Storage",
	DetachStorage: "Failed Detach Storage",
	ResizeStorage: "Failed Resize Storage",
	EditMemberOrg: "Failed Edit member organization",

	// Middleware
	Middleware: "Failed in Middleware",

	// organization
	DeactiveStatus:              "Deactive organization",
	ErrorEditOrg:                "Error when edit organization",
	ErrorCreateOrgRole:          "Error when create organization role",
	ErrorAddMemberOrg:           "Error when add member organization",
	ErrorDeleteMemberOrg:        "Error when deleted member organization",
	ErrorCreateOrg:              "Error when create organization",
	ErrorUpdateOrganizationRole: "Error when update organization role",
	ErrorDeleteOrganizationRole: "Error when delete organization role",

	// User
	AccountSuspend: "Your account suspended",

	// org
	ErrorGetOrg: "Error get organization",
}
