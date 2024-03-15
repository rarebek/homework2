package postgres

import (
	pb "FreelancerMarketplace/user-service/genproto/user_service"
	"FreelancerMarketplace/user-service/pkg/logger"
	"FreelancerMarketplace/user-service/pkg/password"
	"database/sql"
	"errors"
	"fmt"
	"github.com/k0kubun/pp"
	"log"
)

type userRepo struct {
	db  *sql.DB
	log logger.Logger
}

// NewUserRepo ...
func NewUserRepo(db *sql.DB, log logger.Logger) *userRepo {
	return &userRepo{
		db:  db,
		log: log,
	}
}

func (r *userRepo) CompanyRegister(req *pb.CompanyRegisterRequest) (*pb.CompanyRegisterResponse, error) {
	var registeredCompany pb.Company
	hashedPassword, err := password.HashPassword(req.Password)
	if err != nil {
		log.Println("error while hashing password", err)
		return &pb.CompanyRegisterResponse{
			Id:             "",
			CompanyName:    "",
			Description:    "",
			Email:          "",
			Password:       "",
			Address:        "",
			ProfilePicture: "",
			Website:        "",
			Industry:       "",
			EmployeeCount:  0,
			PhoneNumber:    "",
			RefreshToken:   "",
			Message:        "Error while registering user",
		}, err
	}
	req.Password = hashedPassword
	query := "INSERT INTO companies(id, company_name, description, email, password, address, profile_picture, website, industry, employee_count, phone_number, refresh_token) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING id, company_name, description, email, password, address, profile_picture, website, industry, employee_count, phone_number, refresh_token"
	row := r.db.QueryRow(query, req.Id, req.CompanyName, req.Description, req.Email, req.Password, req.Address, req.ProfilePicture, req.Website, req.Industry, req.EmployeeCount, req.PhoneNumber, req.RefreshToken)
	if err := row.Scan(&registeredCompany.Id, &registeredCompany.CompanyName, &registeredCompany.Description, &registeredCompany.Email, &registeredCompany.Password, &registeredCompany.Address, &registeredCompany.ProfilePicture, &registeredCompany.Website, &registeredCompany.Industry, &registeredCompany.EmployeeCount, &registeredCompany.PhoneNumber, &registeredCompany.RefreshToken); err != nil {
		log.Printf("Error while registering user: %s", err)
		return &pb.CompanyRegisterResponse{
			Id:             "",
			CompanyName:    "",
			Description:    "",
			Email:          "",
			Password:       "",
			Address:        "",
			ProfilePicture: "",
			Website:        "",
			Industry:       "",
			EmployeeCount:  0,
			PhoneNumber:    "",
			RefreshToken:   "",
			Message:        "Error while registering user",
		}, err
	}

	pp.Print(pb.CompanyRegisterResponse{
		Id:             registeredCompany.Id,
		CompanyName:    registeredCompany.CompanyName,
		Description:    registeredCompany.Description,
		Email:          registeredCompany.Email,
		Password:       registeredCompany.Password,
		Address:        registeredCompany.Address,
		ProfilePicture: registeredCompany.ProfilePicture,
		Website:        registeredCompany.Website,
		Industry:       registeredCompany.Industry,
		EmployeeCount:  registeredCompany.EmployeeCount,
		PhoneNumber:    registeredCompany.PhoneNumber,
		RefreshToken:   registeredCompany.RefreshToken,
		Message:        "User registered successfully",
	})

	return &pb.CompanyRegisterResponse{
		Id:             registeredCompany.Id,
		CompanyName:    registeredCompany.CompanyName,
		Description:    registeredCompany.Description,
		Email:          registeredCompany.Email,
		Password:       registeredCompany.Password,
		Address:        registeredCompany.Address,
		ProfilePicture: registeredCompany.ProfilePicture,
		Website:        registeredCompany.Website,
		Industry:       registeredCompany.Industry,
		EmployeeCount:  registeredCompany.EmployeeCount,
		PhoneNumber:    registeredCompany.PhoneNumber,
		RefreshToken:   registeredCompany.RefreshToken,
		Message:        "User registered successfully",
	}, nil
}

func (r *userRepo) CompanyLogIn(req *pb.CompanyLogInRequest) (*pb.CompanyLogInResponse, error) {
	var loggedInCompany pb.CompanyLogInResponse
	query := "SELECT id, company_name, description, email, password, address, profile_picture, website, industry, employee_count, phone_number, refresh_token from companies where email = $1"
	row := r.db.QueryRow(query, req.Email)
	if err := row.Scan(&loggedInCompany.Id, &loggedInCompany.CompanyName, &loggedInCompany.Description, &loggedInCompany.Email, &loggedInCompany.Password, &loggedInCompany.Address, &loggedInCompany.ProfilePicture, &loggedInCompany.Website, &loggedInCompany.Industry, &loggedInCompany.EmployeeCount, &loggedInCompany.PhoneNumber, &loggedInCompany.RefreshToken); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			r.log.Warn("Company with this email have not registered yet")
			return &pb.CompanyLogInResponse{
				Id:             "",
				CompanyName:    "",
				Description:    "",
				Email:          "",
				Password:       "",
				Address:        "",
				ProfilePicture: "",
				Website:        "",
				Industry:       "",
				EmployeeCount:  0,
				PhoneNumber:    "",
				RefreshToken:   "",
				Message:        "Company with this email have not registered yet",
			}, nil
		} else {
			r.log.Error("Error inside CompanyLogIn: %s" + err.Error())
			return &pb.CompanyLogInResponse{
				Id:             "",
				CompanyName:    "",
				Description:    "",
				Email:          "",
				Password:       "",
				Address:        "",
				ProfilePicture: "",
				Website:        "",
				Industry:       "",
				EmployeeCount:  0,
				PhoneNumber:    "",
				RefreshToken:   "",
				Message:        "Error while logging in",
			}, err
		}
	}

	correct := password.CheckPasswordHash(req.Password, loggedInCompany.Password)
	if correct {
		r.log.Info("User logged in\n", logger.String("email", req.Email))
		return &pb.CompanyLogInResponse{
			Id:             loggedInCompany.Id,
			CompanyName:    loggedInCompany.CompanyName,
			Description:    loggedInCompany.Description,
			Email:          loggedInCompany.Email,
			Password:       loggedInCompany.Password,
			Address:        loggedInCompany.Address,
			ProfilePicture: loggedInCompany.ProfilePicture,
			Website:        loggedInCompany.Website,
			Industry:       loggedInCompany.Industry,
			EmployeeCount:  loggedInCompany.EmployeeCount,
			PhoneNumber:    loggedInCompany.PhoneNumber,
			RefreshToken:   loggedInCompany.RefreshToken,
			Message:        "Logged in successfully!",
		}, nil
	} else {
		r.log.Warn("Wrong password for email: 	", logger.String("email", req.Email))
		return &pb.CompanyLogInResponse{
			Id:             "",
			CompanyName:    "",
			Description:    "",
			Email:          "",
			Password:       "",
			Address:        "",
			ProfilePicture: "",
			Website:        "",
			Industry:       "",
			EmployeeCount:  0,
			PhoneNumber:    "",
			RefreshToken:   "",
			Message:        "Wrong password. Try again!!!",
		}, nil
	}

}

func (r *userRepo) CheckUniqueness(req *pb.CheckUniquenessRequest) (*pb.CheckUniquenessResponse, error) {
	var count int
	query := fmt.Sprintf("SELECT count(1) from companies WHERE %s = $1 ", req.Field)
	row := r.db.QueryRow(query, req.Value)
	if err := row.Scan(&count); err != nil {
		log.Println("Error while scanning in checkUniqueness method", err)
		return &pb.CheckUniquenessResponse{
			IsUnique: false,
			Message:  "Error while checking uniqueness",
		}, err
	}
	if count == 1 {
		return &pb.CheckUniquenessResponse{
			IsUnique: false,
			Message:  fmt.Sprintf("This %s is already taken. Please try entering different %s", req.Field, req.Field),
		}, nil
	} else {
		return &pb.CheckUniquenessResponse{
			IsUnique: true,
			Message:  fmt.Sprintf("You can use this %s", req.Field),
		}, nil
	}
}
