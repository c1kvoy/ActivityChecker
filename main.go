package main

import (
	"log"
	"site-checker/pkg/checker"
	"site-checker/pkg/config"
	"time"
)

func main() {
	// Загружаем конфигурацию
	cfg, err := config.LoadConfig("config/urls.yaml")
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	// Создаем чекер
	siteChecker := checker.NewSiteChecker(10 * time.Second)
	results := siteChecker.CheckConcurrent(cfg.URLs)

	// Проверяем каждый URL
	for _, result := range results {
		if result.Error != nil {
			log.Printf("❌ %s (%s): ошибка: %v (время: %v)", result.Name, result.URL, result.Error, result.Elapsed)
		} else {
			log.Printf("✅ %s (%s): доступен (время: %v)", result.Name, result.URL, result.Elapsed)
		}
	}
}
