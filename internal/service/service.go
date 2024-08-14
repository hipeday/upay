package service

import (
	"github.com/hipeday/upay/internal/repository"
	"github.com/jmoiron/sqlx"
)

func GetTokenServiceInstance(db *sqlx.DB) TokenService {
	var tokenService TokenService = new(TokenServiceImpl)
	var tokenRepository repository.TokenRepository = new(repository.TokenRepositoryImpl)
	tokenRepository.Setup(db)
	tokenService.Setup(tokenRepository)
	return tokenService
}

func GetAccountServiceInstance(db *sqlx.DB) AccountService {
	var accountService AccountService = new(AccountServiceImpl)
	var accountRepository repository.AccountRepository = new(repository.AccountRepositoryImpl)
	accountRepository.Setup(db)
	accountService.Setup(accountRepository)
	return accountService
}

func GetSettingsServiceInstance(db *sqlx.DB) SettingsService {
	var settingsService SettingsService = new(SettingsServiceImpl)
	var settingsRepository repository.SettingsRepository = new(repository.SettingsRepositoryImpl)
	settingsRepository.Setup(db)
	settingsService.Setup(settingsRepository)
	return settingsService
}
