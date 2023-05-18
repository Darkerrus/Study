package controllers

import (
	"fmt"
	"strconv"
	"test_RestApi/database"
	"test_RestApi/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "secret"

func Register(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		return err
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: string(password),
	}

	insert, err := database.DB.Query(fmt.Sprintf("INSERT INTO `users` (`name`, `email`, `password`) VALUES ('%s', '%s', '%s')", user.Name, user.Email, user.Password))
	if err != nil {
		panic(err)
	}
	defer insert.Close()

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	email := c.FormValue("email")
	pass := c.FormValue("pass")
	var user models.User

	res, err := database.DB.Query(fmt.Sprintf("SELECT * FROM users WHERE email = '%s'", email))
	if err != nil {
		panic(err)
	}
	for res.Next() {
		err = res.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
		if err != nil {
			panic(err)
		}
	}
	defer res.Close()

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON((fiber.Map{
			"message": "user not found",
		}))
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass)); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.Redirect("/")

}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	res, _ := database.DB.Query(fmt.Sprintf("SELECT * FROM users WHERE id = '%s'", claims.Issuer))
	if err != nil {
		panic(err)
	}
	for res.Next() {
		err = res.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
		if err != nil {
			panic(err)
		}
	}
	defer res.Close()

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.Redirect("/")
}

func Index(c *fiber.Ctx) error {
	cook := c.Cookies("jwt")
	return c.Render("index", fiber.Map{"cook": cook}, "layouts/main")

}

func Create(c *fiber.Ctx) error {
	if c.Cookies("jwt") == "" {
		c.Redirect("/")
	}
	return c.Render("create", fiber.Map{}, "layouts/main")
}
func Save(c *fiber.Ctx) error {
	name := c.FormValue("name")
	age := c.FormValue("age")
	description := c.FormValue("description")
	kind := c.FormValue("kind")
	img, err := c.FormFile("photo")

	if err != nil {
		return err
	}
	if name == "" || age == "" || description == "" || kind == "" {
		return c.Render("create", fiber.Map{"Title": "Не все данные заполнены!"}, "layouts/main")
	}

	c.SaveFile(img, "public/photos/"+img.Filename)

	insert, err := database.DB.Query(fmt.Sprintf("INSERT INTO `animals` (`name`, `age`, `description`, `kind`, `img`) VALUES ('%s', '%s', '%s', '%s', '%s')",
		name, age, description, kind, img.Filename))
	if err != nil {
		panic(err)
	}
	defer insert.Close()

	return c.Redirect("/")
}

func Animals(c *fiber.Ctx) error {
	var animal models.Animal
	var animals = []models.Animal{}
	res, err := database.DB.Query(fmt.Sprintf("SELECT * FROM animals"))
	if err != nil {
		panic(err)
	}
	for res.Next() {
		err = res.Scan(&animal.Id, &animal.Name, &animal.Age, &animal.Description, &animal.Kind, &animal.Img)
		if err != nil {
			panic(err)
		}
		animals = append(animals, animal)
	}
	defer res.Close()
	return c.Render("animals", animals, "layouts/main")

}

func Animal(c *fiber.Ctx) error {
	id := c.Params("id")
	var animal models.Animal
	res, err := database.DB.Query(fmt.Sprintf("SELECT * FROM animals WHERE id = '%s'", id))
	if err != nil {
		panic(err)
	}
	for res.Next() {
		err = res.Scan(&animal.Id, &animal.Name, &animal.Age, &animal.Description, &animal.Kind, &animal.Img)
		if err != nil {
			panic(err)
		}
	}
	defer res.Close()

	return c.Render("animal", animal, "layouts/main")

}

func Auth(c *fiber.Ctx) error {
	return c.Render("auth", fiber.Map{}, "layouts/main")
}
