package main

import (
	"HR-monitor/pkg/auth"
	"HR-monitor/pkg/config"
	"HR-monitor/pkg/enums"
	"HR-monitor/pkg/handlers"
	"HR-monitor/pkg/repository"
	"HR-monitor/pkg/repository/postgres"
	"HR-monitor/pkg/service"
	"context"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Dependencies struct {
	DB     *pgxpool.Pool
	Config config.Config
}

func NewDependencies(db *pgxpool.Pool, config config.Config) *Dependencies {
	return &Dependencies{
		DB:     db,
		Config: config,
	}
}

func (d *Dependencies) NewAuthHandler() *handlers.AuthHandler {
	userRepo := postgres.NewPostgresUserRepository(d.DB)
	jwtService := auth.NewJWTService(d.Config.JWTSecret)
	authService := service.NewAuthService(userRepo, jwtService)
	return handlers.NewAuthHandler(authService)
}

func (d *Dependencies) NewResumeHandler() *handlers.ResumeHandler {
	resumeRepo := postgres.NewPostgresResumeRepository(d.DB)
	statsRepo := postgres.NewPostgresStatsRepository(d.DB)
	resumeService := service.NewResumeService(resumeRepo, statsRepo)
	return handlers.NewResumeHandler(resumeService)
}

func (d *Dependencies) NewVacancyHandler() *handlers.VacancyHandler {
	vacancyRepo := postgres.NewPostgresVacancyRepository(d.DB)
	vacancyService := service.NewVacancyService(vacancyRepo)
	return handlers.NewVacancyHandler(vacancyService)
}

func main() {
	ctx := context.Background()
	config := config.LoadConfig()

	err := repository.InitDB(ctx, config)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	db := repository.GetDB()
	defer repository.CloseDB()

	deps := NewDependencies(db, config)

	authHandler := deps.NewAuthHandler()
	resumeHandler := deps.NewResumeHandler()
	vacancyHandler := deps.NewVacancyHandler()

	jwtService := auth.NewJWTService(config.JWTSecret)
	jwtMiddleware := auth.JWTAuthMiddleware(jwtService)

	http.HandleFunc("/api/auth/register", authHandler.Register)
	http.HandleFunc("/api/auth/login", authHandler.Login)

	hrMux := http.NewServeMux()
	hrMux.HandleFunc("/resumes", resumeHandler.GetResumes)
	hrMux.HandleFunc("/resumes/create", resumeHandler.CreateResume)
	hrMux.HandleFunc("/resumes/move", resumeHandler.MoveResumeToStage)
	hrMux.HandleFunc("/resumes/stats", resumeHandler.GetResumeStats)
	hrMux.HandleFunc("/vacancies", vacancyHandler.GetVacancyByID)

	teamLeadMux := http.NewServeMux()
	teamLeadMux.HandleFunc("/vacancies/create", vacancyHandler.CreateVacancy)
	teamLeadMux.HandleFunc("/vacancies/delete", vacancyHandler.DeleteVacancy)
	teamLeadMux.HandleFunc("/vacancies/status", vacancyHandler.ChangeVacancyStatus)

	http.Handle("/api/hr/", http.StripPrefix("/api/hr", jwtMiddleware(auth.RequireRoles(enums.HRRole)(hrMux))))
	http.Handle("/api/team-lead/", http.StripPrefix("/api/team-lead", jwtMiddleware(auth.RequireRoles(enums.TeamLeadRole)(teamLeadMux))))

	log.Printf("Server starting on port %s", config.ServerPort)
	log.Fatal(http.ListenAndServe(":"+config.ServerPort, nil))
}
