package app

import (
	"context"
	"database/sql"
	"deficonnect/go-api-starter/postgres/models"
)

type Trade struct {
	ID        string `boil:"id" json:"id" toml:"id" yaml:"id"`
	AccountID string `boil:"account_id" json:"account_id" toml:"account_id" yaml:"account_id"`
	TradeNo   int    `boil:"trade_no" json:"trade_no" toml:"trade_no" yaml:"trade_no"`
	Date      int64  `boil:"date" json:"date" toml:"date" yaml:"date"`
	StartDate int64  `boil:"start_date" json:"start_date" toml:"start_date" yaml:"start_date"`
	EndDate   int64  `boil:"end_date" json:"end_date" toml:"end_date" yaml:"end_date"`
	Amount    int64  `boil:"amount" json:"amount" toml:"amount" yaml:"amount"`
	Profit    int64  `boil:"profit" json:"profit" toml:"profit" yaml:"profit"`
}

type store interface {
	BeginTransaction() (*sql.Tx, error)

	CreateAccount(ctx context.Context, input CreateAccountInput) error
	GetAccount(ctx context.Context, id string) (*models.Account, error)
	AccountBalance(ctx context.Context, accountId string) (int64, error)
	GetAccounts(ctx context.Context, skip, limit int) ([]*models.Account, error)
	GetPasswordResetCode(ctx context.Context, accountID string) (string, error)
	ValidatePasswordResetCode(ctx context.Context, accountID, code string) (bool, error)
	ChangePassword(ctx context.Context, accountID, password string) error
	GetAccountIDs(ctx context.Context) ([]string, error)
	GetAllAccountsCount(ctx context.Context) (int64, error)
	GetAccountByUsername(ctx context.Context, username string) (*models.Account, error)
	GetAccountByReferralCode(ctx context.Context, referralCode string) (*models.Account, error)
	UpdateAccountDetail(ctx context.Context, accountID string, input UpdateDetailInput) error
	MyDownlines(ctx context.Context, accountID string, generation int64, offset, limit int) ([]DownlineInfo, int64, error)
	GetRefferalCount(ctx context.Context, accountID string) (int64, error)
	GetTeamInformation(ctx context.Context, accountID string) (*TeamInfo, error)
	GetAccountByEmail(ctx context.Context, email string) (*models.Account, error)
	DebitAccount(ctx context.Context, accountID string, amount, date int64, ref string) error
	DebitAccountTx(ctx context.Context, tx *sql.Tx, accountID string, amount, date int64, ref string) error
	CreditAccount(ctx context.Context, accountID string, amount, date int64, ref string) error
	SpendStableNaira(ctx context.Context, tx *sql.Tx, accountID string, amount int64) error

	CreateNotification(ctx context.Context, accountID, title, message, actionText, actionLink string, notificationType int) error
	NotifyAll(ctx context.Context, titile string, content, actionText, actionLink string, notificationType int) error
	UnReadNotificationCount(ctx context.Context, accountID string, notificationType int) (int64, error)
	GetNotifications(ctx context.Context, accountID string, notificationType int, offset, limit int) (models.NotificationSlice, int64, error)
	GetNewNotifications(ctx context.Context, accountID string, notificationType int, offset, limit int) (models.NotificationSlice, int64, error)
	GetNotification(ctx context.Context, id string) (*models.Notification, error)

	SetConfigValue(ctx context.Context, accountID, key string, value ConfigValue) error
	GetConfigValue(ctx context.Context, accountID, key string) (ConfigValue, error)
	GetMasterConfigValue(ctx context.Context, key string) (ConfigValue, error)
	GetConfigs(ctx context.Context, accountID string) (models.UserSettingSlice, error)
	AddLogin(ctx context.Context, accountID, ip, platform string, date int64) error
	LastLogin(ctx context.Context) (*models.LoginInfo, error)
}
