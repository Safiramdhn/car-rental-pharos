package routes

import (
	"car-rental/infra"
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRoutes(ctx infra.ServiceContext) {
	routes := gin.Default()
	routes.Use(cors.Default())

	// customer routes
	customerRoutes := routes.Group("/customer")
	{
		customerRoutes.GET("/", ctx.Ctl.Customer.GetAllCustomers)
		customerRoutes.GET("/:id", ctx.Ctl.Customer.GetCustomer)
		customerRoutes.POST("/", ctx.Ctl.Customer.CreateCustomer)
		customerRoutes.PUT("/:id", ctx.Ctl.Customer.UpdateCustomer)
		customerRoutes.DELETE("/:id", ctx.Ctl.Customer.DeleteCustomer)
	}

	// car routes
	carRoutes := routes.Group("/car")
	{
		carRoutes.GET("/", ctx.Ctl.Car.GetAllCars)
		carRoutes.GET("/:id", ctx.Ctl.Car.GetCar)
		carRoutes.POST("/", ctx.Ctl.Car.CreateCar)
		carRoutes.PUT("/:id", ctx.Ctl.Car.UpdateCar)
		carRoutes.DELETE("/:id", ctx.Ctl.Car.DeleteCar)
	}

	// booking routes
	bookingRoutes := routes.Group("/booking")
	{
		bookingRoutes.GET("/", ctx.Ctl.Booking.GetAllBookings)
		bookingRoutes.GET("/:id", ctx.Ctl.Booking.GetBooking)
		bookingRoutes.POST("/", ctx.Ctl.Booking.CreateBooking)
		bookingRoutes.PUT("/:id", ctx.Ctl.Booking.UpdateBooking)
		bookingRoutes.DELETE("/:id", ctx.Ctl.Booking.DeleteBooking)
	}

	// membership routes
	membershipRoutes := routes.Group("/membership")
	{
		membershipRoutes.PUT("/:customer_id", ctx.Ctl.Membership.SetMembership)
	}

	// driver routes
	driverRoutes := routes.Group("/driver")
	{
		driverRoutes.GET("/", ctx.Ctl.Driver.GetAllDrivers)
		driverRoutes.GET("/:id", ctx.Ctl.Driver.GetDriver)
		driverRoutes.POST("/", ctx.Ctl.Driver.CreateDriver)
		driverRoutes.PUT("/:id", ctx.Ctl.Driver.UpdateDriver)
		driverRoutes.DELETE("/:id", ctx.Ctl.Driver.DeleteDriver)
	}

	// driver incentive routes
	driverIncentiveRoutes := routes.Group("/driver-incentive")
	{
		driverIncentiveRoutes.GET("/", ctx.Ctl.DriverIncentive.GetAllDriverIncentives)
	}

	// booking type routes
	bookingTypeRoutes := routes.Group("/booking-type")
	{
		bookingTypeRoutes.GET("/", ctx.Ctl.BookingType.GetAllBookingTypes)
	}

	gracefulShutdown(ctx, routes.Handler())
}

func gracefulShutdown(ctx infra.ServiceContext, handler http.Handler) {
	srv := &http.Server{
		Addr:    ctx.Cfg.Port,
		Handler: handler,
	}

	if ctx.Cfg.ShutdownTimeout == 0 {
		launchServer(srv, ctx.Cfg.Port)
		return
	}

	go func() {
		launchServer(srv, ctx.Cfg.Port)
	}()

	// Wait for interrupt signal to gracefully shut down the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")
	appContext, cancel := context.WithTimeout(context.Background(), time.Duration(ctx.Cfg.ShutdownTimeout)*time.Second)
	defer cancel()
	if err := srv.Shutdown(appContext); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching appContext.Done(). timeout of ShutdownTimeout seconds.
	select {
	case <-appContext.Done():
		log.Printf("timeout of %d seconds.", ctx.Cfg.ShutdownTimeout)
	}
	log.Println("Server exiting")
}

func launchServer(server *http.Server, port string) {
	// service connections
	log.Println("Listening and serving HTTP on", port)
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("listen: %s\n", err)
	}
}
