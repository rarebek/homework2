package model

type Company struct {
	Id             string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CompanyName    string `protobuf:"bytes,2,opt,name=company_name,json=companyName,proto3" json:"company_name,omitempty"`
	Description    string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Email          string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Password       string `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	Address        string `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
	ProfilePicture string `protobuf:"bytes,7,opt,name=profile_picture,json=profilePicture,proto3" json:"profile_picture,omitempty"`
	Website        string `protobuf:"bytes,8,opt,name=website,proto3" json:"website,omitempty"`
	Industry       string `protobuf:"bytes,9,opt,name=industry,proto3" json:"industry,omitempty"`
	EmployeeCount  int64  `protobuf:"varint,10,opt,name=employee_count,json=employeeCount,proto3" json:"employee_count,omitempty"`
	PhoneNumber    string `protobuf:"bytes,11,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	RefreshToken   string `protobuf:"bytes,12,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
}

type Freelancer struct {
	Id             string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName      string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName       string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Username       string `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	Biography      string `protobuf:"bytes,5,opt,name=biography,proto3" json:"biography,omitempty"`
	ProfilePicture string `protobuf:"bytes,6,opt,name=profile_picture,json=profilePicture,proto3" json:"profile_picture,omitempty"`
	Email          string `protobuf:"bytes,7,opt,name=email,proto3" json:"email,omitempty"`
	Password       string `protobuf:"bytes,8,opt,name=password,proto3" json:"password,omitempty"`
	PhoneNumber    string `protobuf:"bytes,9,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Address        string `protobuf:"bytes,10,opt,name=address,proto3" json:"address,omitempty"`
	Resume         string `protobuf:"bytes,11,opt,name=resume,proto3" json:"resume,omitempty"`
	RefreshToken   string `protobuf:"bytes,12,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
}

type CompanyRegisterRequest struct {
	CompanyName    string `protobuf:"bytes,1,opt,name=company_name,json=companyName,proto3" json:"company_name,omitempty"`
	Description    string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Email          string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password       string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Address        string `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	ProfilePicture string `protobuf:"bytes,6,opt,name=profile_picture,json=profilePicture,proto3" json:"profile_picture,omitempty"`
	Website        string `protobuf:"bytes,7,opt,name=website,proto3" json:"website,omitempty"`
	Industry       string `protobuf:"bytes,8,opt,name=industry,proto3" json:"industry,omitempty"`
	EmployeeCount  int64  `protobuf:"varint,9,opt,name=employee_count,json=employeeCount,proto3" json:"employee_count,omitempty"`
	PhoneNumber    string `protobuf:"bytes,10,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	RefreshToken   string `protobuf:"bytes,11,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	Code           int64
}

type CompanyRegisterResponse struct {
	Id             string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CompanyName    string `protobuf:"bytes,2,opt,name=company_name,json=companyName,proto3" json:"company_name,omitempty"`
	Description    string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Email          string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Password       string `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	Address        string `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
	ProfilePicture string `protobuf:"bytes,7,opt,name=profile_picture,json=profilePicture,proto3" json:"profile_picture,omitempty"`
	Website        string `protobuf:"bytes,8,opt,name=website,proto3" json:"website,omitempty"`
	Industry       string `protobuf:"bytes,9,opt,name=industry,proto3" json:"industry,omitempty"`
	EmployeeCount  int64  `protobuf:"varint,10,opt,name=employee_count,json=employeeCount,proto3" json:"employee_count,omitempty"`
	PhoneNumber    string `protobuf:"bytes,11,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	RefreshToken   string `protobuf:"bytes,12,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	Message        string `protobuf:"bytes,13,opt,name=message,proto3" json:"message,omitempty"`
}

type CompanyLogInRequest struct {
	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

type CompanyLogInResponse struct {
	Id             string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CompanyName    string `protobuf:"bytes,2,opt,name=company_name,json=companyName,proto3" json:"company_name,omitempty"`
	Description    string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Email          string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Password       string `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	Address        string `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
	ProfilePicture string `protobuf:"bytes,7,opt,name=profile_picture,json=profilePicture,proto3" json:"profile_picture,omitempty"`
	Website        string `protobuf:"bytes,8,opt,name=website,proto3" json:"website,omitempty"`
	Industry       string `protobuf:"bytes,9,opt,name=industry,proto3" json:"industry,omitempty"`
	EmployeeCount  int64  `protobuf:"varint,10,opt,name=employee_count,json=employeeCount,proto3" json:"employee_count,omitempty"`
	PhoneNumber    string `protobuf:"bytes,11,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	RefreshToken   string `protobuf:"bytes,12,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	Message        string `protobuf:"bytes,13,opt,name=message,proto3" json:"message,omitempty"`
}

type RegisterCompanyResponse struct {
	Id             string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CompanyName    string `protobuf:"bytes,2,opt,name=company_name,json=companyName,proto3" json:"company_name,omitempty"`
	Description    string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Email          string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Password       string `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	Address        string `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
	ProfilePicture string `protobuf:"bytes,7,opt,name=profile_picture,json=profilePicture,proto3" json:"profile_picture,omitempty"`
	Website        string `protobuf:"bytes,8,opt,name=website,proto3" json:"website,omitempty"`
	Industry       string `protobuf:"bytes,9,opt,name=industry,proto3" json:"industry,omitempty"`
	EmployeeCount  int64  `protobuf:"varint,10,opt,name=employee_count,json=employeeCount,proto3" json:"employee_count,omitempty"`
	PhoneNumber    string `protobuf:"bytes,11,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	AccessToken    string `protobuf:"bytes,12,opt,name=access,json=accessToken,proto3" json:"access_token,omitempty"`
}

type CheckUniquenessRequest struct {
	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

type CheckUniquenessResponse struct {
	IsUnique bool   `protobuf:"varint,1,opt,name=is_unique,json=isUnique,proto3" json:"is_unique,omitempty"`
	Message  string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

type JwtRequestModel struct {
	Token string `json:"token"`
}

type ResponseError struct {
	Error interface{} `json:"error"`
}

// ServerError ...
type ServerError struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
