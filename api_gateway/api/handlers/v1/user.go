package v1

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/smtp"
	"strconv"
	"strings"
	"time"

	"FreelancerMarketplace/api-gateway/api/model"
	pbu "FreelancerMarketplace/api-gateway/genproto/user_service"
	codegen "FreelancerMarketplace/api-gateway/pkg/codegen"
	l "FreelancerMarketplace/api-gateway/pkg/logger"
	"FreelancerMarketplace/api-gateway/queue/kafka/producer"

	// "FreelancerMarketplace/api_gateway/queue/kafka/producer"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	uuid "github.com/google/uuid"
	"google.golang.org/protobuf/encoding/protojson"
)

// RegisterCompany
// CreateCompany Creates company account creates user
// @Summary Create company summary
// @Description This api is using for creating new company
// @Tags company
// @Accept json
// @Produce json
// @Param user body model.CompanyRegisterRequest true "user"
// @Success 200 {object} model.CompanyRegisterResponse
// @Router /v1/company/register [post]
func (h *handlerV1) RegisterCompany(c *gin.Context) {
	var (
		body        model.CompanyRegisterRequest
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	body.Email = strings.TrimSpace(body.Email)
	body.Email = strings.ToLower(body.Email)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	exists, err := h.serviceManager.CompanyService().CheckUniqueness(ctx, &pbu.CheckUniquenessRequest{
		Field: "email",
		Value: body.Email,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to check email uniqueness", l.Error(err))
		return
	}

	if !exists.IsUnique {
		c.JSON(http.StatusConflict, gin.H{
			"error": "This email already in use, please use another email address",
		})
		h.log.Error("failed to check email uniqueness", l.Error(err))
		return
	}

	// generate 6 digits number
	code := codegen.GenerateCode()
	body.Code = code
	// send to email of user this number
	type PageData struct {
		OTP string
	}
	tpl := template.Must(template.ParseFiles("index.html"))
	data := PageData{
		OTP: strconv.Itoa(int(code)),
	}
	var buf bytes.Buffer
	tpl.Execute(&buf, data)
	htmlContent := buf.Bytes()

	auth := smtp.PlainAuth("", "nodirbekgolang@gmail.com", "ecncwhvfdyvjghux", "smtp.gmail.com")
	err = smtp.SendMail("smtp.gmail.com:587", auth, "nodirbekgolang@gmail.com", []string{body.Email}, []byte("To: "+body.Email+"\r\nSubject: Email verification\r\nMIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"+string(htmlContent)))
	if err != nil {
		log.Fatalf("Error sending otp to email: %v", err)
	}
	log.Println("Email sent successfully")
	// save generated 6 digits number to user data in redis
	// body.Code = generated_6_digits_code
	byteUser, err := json.Marshal(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while marshalling user data", l.Error(err))
		return
	}

	err = h.redisStorage.SetWithTTL(body.Email, string(byteUser), int64(time.Second*300))
	fmt.Println("redis set")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while setting user data to redis", l.Error(err))
		return
	}

	err = producer.ProduceMessage("test-topic", string(byteUser))
	if err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "code sent to your email. Please verify your email",
	})

}

// VerifyCompany
// @Summary Verify
// @Description Api for Verifying
// @Tags company
// @Accept json
// @Produce json
// @Param email path string true "Verification email"
// @Param code path string true "Verification code"
// @Success 200 {object} model.RegisterCompanyResponse
// @Failure 400 {object} model.StandardErrorModel
// @Failure 500 {object} model.StandardErrorModel
// @Router /v1/company/verify/{email}/{code} [post]
func (h *handlerV1) VerifyCompany(c *gin.Context) {
	var companyData model.CompanyRegisterRequest
	email := c.Param("email")
	code := c.Param("code")
	fmt.Println("email: ", email)
	fmt.Println("code: ", code)

	intCode, _ := strconv.Atoi(code)

	data, err := redis.String(h.redisStorage.Get(email))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting data from redis", l.Error(err))
		return
	}

	err = json.Unmarshal([]byte(data), &companyData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while unmarshalling user data", l.Error(err))
		return
	}
	//pp.Print(companyData)

	if intCode != int(companyData.Code) {
		if err != nil {
			c.JSON(http.StatusConflict, gin.H{
				"error": "your code is not match",
			})
			h.log.Error("code is invalid", l.Error(err))
			return
		}
	}

	id, err := uuid.NewUUID()
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "error while generating uuid",
		})
		h.log.Error("error generate new uuid", l.Error(err))
		return
	}
	//pp.Print(pbu.CompanyRegisterRequest{
	//	CompanyName:    id.String(),
	//	Description:    companyData.Description,
	//	Email:          companyData.Email,
	//	Password:       companyData.Password,
	//	Address:        companyData.Address,
	//	ProfilePicture: companyData.ProfilePicture,
	//	Website:        companyData.Website,
	//	Industry:       companyData.Industry,
	//	EmployeeCount:  companyData.EmployeeCount,
	//	PhoneNumber:    companyData.PhoneNumber,
	//	RefreshToken:   refresh,
	//})
	response, err := h.serviceManager.CompanyService().CompanyRegister(context.Background(), &pbu.CompanyRegisterRequest{
		Id:             id.String(),
		CompanyName:    companyData.CompanyName,
		Description:    companyData.Description,
		Email:          companyData.Email,
		Password:       companyData.Password,
		Address:        companyData.Address,
		ProfilePicture: companyData.ProfilePicture,
		Website:        companyData.Website,
		Industry:       companyData.Industry,
		EmployeeCount:  companyData.EmployeeCount,
		PhoneNumber:    companyData.PhoneNumber,
		//RefreshToken:   refresh,
	})
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	// Create user logic
	c.JSON(http.StatusOK, &model.RegisterCompanyResponse{
		Id:             response.Id,
		CompanyName:    response.CompanyName,
		Description:    response.Description,
		Email:          response.Email,
		Password:       response.Password,
		Address:        response.Address,
		ProfilePicture: response.ProfilePicture,
		Website:        response.Website,
		Industry:       response.Industry,
		EmployeeCount:  response.EmployeeCount,
		PhoneNumber:    response.PhoneNumber,
		//AccessToken:    access,
	})
}

// LogInCompany
// @Summary LogIn
// @Description Api for Logging in
// @Tags company
// @Accept json
// @Produce json
// @Param email path string true "Your email"
// @Param password path string true "Your password"
// @Success 200 {object} model.CompanyLogInResponse
// @Failure 400 {object} model.StandardErrorModel
// @Failure 500 {object} model.StandardErrorModel
// @Router /v1/company/login/{email}/{password} [post]
func (h *handlerV1) LogInCompany(c *gin.Context) {
	email := c.Param("email")
	password := c.Param("password")
	response, err := h.serviceManager.CompanyService().CompanyLogIn(context.Background(), &pbu.CompanyLogInRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusAccepted, response)
}
