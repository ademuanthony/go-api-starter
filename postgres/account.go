package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"deficonnect/sonicflare/app"
	"deficonnect/sonicflare/postgres/models"

	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var (
	welcomeBonus   = 250 * 500 * 1e4
	level1Referral = 150 * 500 * 1e4
	level2Referral = 100 * 500 * 1e4
	level3Referral = 50 * 500 * 1e4
)

func (pg PgDb) CreateAccount(ctx context.Context, input app.CreateAccountInput) error {
	account := models.Account{
		ID:          uuid.NewString(),
		ReferralID:  null.StringFrom(input.ReferralCode),
		Username:    input.Username,
		Password:    input.Password,
		Email:       input.Email,
		PhoneNumber: input.PhoneNumber,
		FirstName:   input.Name,
		CreatedAt:   time.Now().Unix(),
	}

	referralCode := strings.ReplaceAll(uuid.NewString(), "-", "")[0:6]
	for {
		if ex, _ := models.Accounts(models.AccountWhere.ReferralCode.EQ(referralCode)).Exists(ctx, pg.Db); !ex {
			break
		}

		referralCode = strings.ReplaceAll(uuid.NewString(), "-", "")[0:6]
	}

	account.ReferralCode = referralCode

	tx, err := pg.Db.Begin()
	if err != nil {
		return err
	}

	if input.ReferralCode != "" {
		father, err := models.Accounts(
			models.AccountWhere.ReferralCode.EQ(input.ReferralCode),
		).One(ctx, pg.Db)
		if err != nil {
			log.Info("no ref", input.ReferralCode)
			tx.Rollback()
			return fmt.Errorf("invalid referral code")
		}
		if err := pg.AddReferralEarning(ctx, tx, father.ID, int64(level1Referral)); err != nil {
			tx.Rollback()
			return err
		}

		if err := pg.createNotificationTx(ctx, tx, father.ID, "New referral earning",
			fmt.Sprintf("You have earned %f from a new direct referral", level1Referral/1e4), "", "",
			app.NOTIFICATION_TYPE_TOOLBAR); err != nil {
			tx.Rollback()
			return err
		}

		account.ReferralID2 = father.ReferralID
		account.ReferralID3 = father.ReferralID2

		if father.ReferralID2.String != "" {
			grandFather, err := models.Accounts(
				models.AccountWhere.ReferralCode.EQ(account.ReferralID2.String),
			).One(ctx, pg.Db)
			if err == nil {
				if err := pg.AddReferralEarning(ctx, tx, grandFather.ID, int64(level2Referral)); err != nil {
					tx.Rollback()
					return err
				}
			}
		}

		if father.ReferralID3.String != "" {
			greatGrandFather, err := models.Accounts(
				models.AccountWhere.ReferralCode.EQ(account.ReferralID3.String),
			).One(ctx, pg.Db)
			if err == nil {
				if err := pg.AddReferralEarning(ctx, tx, greatGrandFather.ID, int64(level3Referral)); err != nil {
					tx.Rollback()
					return err
				}
			}
		}

	}

	err = account.Insert(ctx, tx, boil.Infer())
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := pg.createNotificationTx(ctx, tx, account.ID, "Welcome Bonus",
		fmt.Sprintf("You have being credited with $%f welcome bonus. Congratulation!", welcomeBonus/1e4), "View Balance", "/app/wallet/balance",
		app.NOTIFICATION_TYPE_DASHBOARD); err != nil {
		tx.Rollback()
		return err
	}

	wallet := models.Wallet{
		ID:         uuid.NewString(),
		AccountID:  account.ID,
		Address:    input.DepositWalletAddress,
		PrivateKey: input.PrivateKey,
		CoinSymbol: "BEP20-USDT",
	}

	if err = wallet.Insert(ctx, tx, boil.Infer()); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (pg PgDb) AddLogin(ctx context.Context, accountID, ip, platform string, date int64) error {
	info := models.LoginInfo{
		AccountID: accountID,
		Platform:  platform,
		IP:        ip,
		Date:      date,
	}

	return info.Insert(ctx, pg.Db, boil.Infer())
}

func (pg PgDb) LastLogin(ctx context.Context) (*models.LoginInfo, error) {
	maxDate := time.Now().Add(-1 * time.Minute).Unix()
	return models.LoginInfos(
		models.LoginInfoWhere.Date.LTE(maxDate),
		qm.OrderBy(models.LoginInfoColumns.Date+" desc"),
	).One(ctx, pg.Db)
}

func (pg PgDb) CreateDepositWallet(ctx context.Context, accountID, address, privateKey string) (*models.Wallet, error) {
	wallet := models.Wallet{
		ID:         uuid.NewString(),
		AccountID:  accountID,
		Address:    address,
		PrivateKey: privateKey,
		CoinSymbol: "BEP20-USDT",
	}

	err := wallet.Insert(ctx, pg.Db, boil.Infer())
	return &wallet, err
}

func (pg PgDb) GetAccount(ctx context.Context, id string) (*models.Account, error) {
	acc, err := models.FindAccount(ctx, pg.Db, id)
	if err != nil {
		return nil, err
	}
	bal, err := pg.AccountBalance(ctx, id)
	if err != nil {
		return nil, err
	}
	acc.Balance = bal

	return acc, nil

}

func (pg PgDb) GetAccountByUsername(ctx context.Context, username string) (*models.Account, error) {
	acc, err := models.Accounts(
		models.AccountWhere.Username.EQ(username),
	).One(ctx, pg.Db)

	if err != nil {
		return nil, err
	}
	bal, err := pg.AccountBalance(ctx, acc.ID)
	if err != nil {
		return nil, err
	}
	acc.Balance = bal

	return acc, nil
}

func (pg PgDb) GetAccountByEmail(ctx context.Context, email string) (*models.Account, error) {
	acc, err := models.Accounts(
		models.AccountWhere.Email.EQ(email),
	).One(ctx, pg.Db)

	if err != nil {
		return nil, err
	}
	bal, err := pg.AccountBalance(ctx, acc.ID)
	if err != nil {
		return nil, err
	}
	acc.Balance = bal

	return acc, nil
}

func (pg PgDb) GetAccountByReferralCode(ctx context.Context, referralCode string) (*models.Account, error) {
	acc, err := models.Accounts(
		models.AccountWhere.ReferralCode.EQ(referralCode),
	).One(ctx, pg.Db)

	if err != nil {
		return nil, err
	}
	bal, err := pg.AccountBalance(ctx, acc.ID)
	if err != nil {
		return nil, err
	}
	acc.Balance = bal

	return acc, nil
}

func (pg PgDb) GetAccountByTelegramID(ctx context.Context, telegramID int64) (*models.Account, error) {
	acc, err := models.Accounts(
		models.AccountWhere.TelegramID.EQ(telegramID),
	).One(ctx, pg.Db)

	if err != nil {
		return nil, err
	}
	bal, err := pg.AccountBalance(ctx, acc.ID)
	if err != nil {
		return nil, err
	}
	acc.Balance = bal

	return acc, nil
}

func (pg PgDb) GetPasswordResetCode(ctx context.Context, accountID string) (string, error) {
	// delete expired code
	minDate := time.Now().Add(-15 * time.Minute)
	if _, err := models.SecurityCodes(models.SecurityCodeWhere.Date.LTE(minDate.Unix())).DeleteAll(ctx, pg.Db); err != nil {
		return "", err
	}

	code, err := models.SecurityCodes(models.SecurityCodeWhere.Date.GT(minDate.Unix())).One(ctx, pg.Db)
	if err == sql.ErrNoRows {
		code = &models.SecurityCode{
			Code:      randomCode(6),
			AccountID: accountID,
			Date:      time.Now().Unix(),
		}
		if err = code.Insert(ctx, pg.Db, boil.Infer()); err != nil {
			return "", err
		}
	}

	if err != nil {
		return "", err
	}

	return code.Code, err
}

func (pg PgDb) ValidatePasswordResetCode(ctx context.Context, accountID, code string) (bool, error) {
	minDate := time.Now().Add(-15 * time.Minute)
	if _, err := models.SecurityCodes(models.SecurityCodeWhere.Date.LTE(minDate.Unix())).DeleteAll(ctx, pg.Db); err != nil {
		return false, err
	}

	lastCode, err := models.SecurityCodes(
		models.SecurityCodeWhere.AccountID.EQ(accountID),
		models.SecurityCodeWhere.Date.GT(minDate.Unix()),
		qm.OrderBy(models.SecurityCodeColumns.Date+" desc"),
	).One(ctx, pg.Db)
	if err != nil {
		return false, err
	}

	return code == lastCode.Code, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("0123456789")

func randomCode(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func (pg PgDb) GetAllAccountsCount(ctx context.Context) (int64, error) {
	return models.Accounts().Count(ctx, pg.Db)
}

func (pg PgDb) GetAccounts(ctx context.Context, skip, limit int) ([]*models.Account, error) {
	return models.Accounts(
		qm.Offset(skip),
		qm.Limit(limit),
	).All(ctx, pg.Db)
}

func (pg PgDb) GetAccountIDs(ctx context.Context) ([]string, error) {
	accounts, err := models.Accounts(
		qm.Select(models.AccountColumns.ID),
	).All(ctx, pg.Db)
	if err != nil {
		return nil, err
	}

	var ids []string
	for _, acc := range accounts {
		ids = append(ids, acc.ID)
	}

	return ids, nil
}

func (pg PgDb) UpdateAccountDetail(ctx context.Context, accountID string, input app.UpdateDetailInput) error {
	var upCol = models.M{}
	if input.FirstName != "" {
		upCol[models.AccountColumns.FirstName] = input.FirstName
	}
	if input.PhoneNumber != "" {
		upCol[models.AccountColumns.PhoneNumber] = input.PhoneNumber
	}
	if input.LastName != "" {
		upCol[models.AccountColumns.LastName] = input.LastName
	}

	_, err := models.Accounts(models.AccountWhere.ID.EQ(accountID)).UpdateAll(ctx, pg.Db, upCol)
	return err
}

func (pg PgDb) ChangePassword(ctx context.Context, accountID, password string) error {
	colUp := models.M{
		models.AccountColumns.Password: password,
	}
	_, err := models.Accounts(models.AccountWhere.ID.EQ(accountID)).UpdateAll(ctx, pg.Db, colUp)
	return err
}

func (pg PgDb) GetRefferalCount(ctx context.Context, accountID string) (int64, error) {
	account, err := models.FindAccount(ctx, pg.Db, accountID)
	if err != nil {
		return 0, err
	}
	return models.Accounts(
		models.AccountWhere.ReferralID.EQ(null.StringFrom(account.ReferralCode)),
	).Count(ctx, pg.Db)
}

func (pg PgDb) GetTeamInformation(ctx context.Context, accountID string) (*app.TeamInfo, error) {
	account, err := models.FindAccount(ctx, pg.Db, accountID)
	if err != nil {
		return nil, err
	}

	g1, err := models.Accounts(
		models.AccountWhere.ReferralID.EQ(null.StringFrom(account.ReferralCode)),
	).Count(ctx, pg.Db)
	if err != nil {
		return nil, err
	}

	g2, err := models.Accounts(
		models.AccountWhere.ReferralID2.EQ(null.StringFrom(account.ReferralCode)),
	).Count(ctx, pg.Db)
	if err != nil {
		return nil, err
	}

	g3, err := models.Accounts(
		models.AccountWhere.ReferralID3.EQ(null.StringFrom(account.ReferralCode)),
	).Count(ctx, pg.Db)
	if err != nil {
		return nil, err
	}

	statement := "select coalesce(sum(principal), 0) as principal from account where referral_id = $1"
	acc, err := models.Accounts(qm.SQL(statement, account.ReferralCode)).One(ctx, pg.Db)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	var p1 int64
	if acc != nil {
		p1 = acc.Principal
	}

	statement = "select coalesce(sum(principal), 0) as principal from account where referral_id_2 = $1"
	acc, err = models.Accounts(qm.SQL(statement, account.ReferralCode)).One(ctx, pg.Db)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	var p2 int64
	if acc != nil {
		p2 = acc.Principal
	}

	statement = "select coalesce(sum(principal), 0) as principal from account where referral_id_3 = $1"
	acc, err = models.Accounts(qm.SQL(statement, account.ReferralCode)).One(ctx, pg.Db)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	var p3 int64
	if acc != nil {
		p3 = acc.Principal
	}

	return &app.TeamInfo{
		FirstGeneration:   g1,
		SecoundGeneration: g2,
		ThirdGeneration:   g3,
		Pool1:             p1,
		Pool2:             p2,
		Pool3:             p3,
	}, nil
}

func (pg PgDb) GetDepositAddress(ctx context.Context, accountID string) (*models.Wallet, error) {
	return models.Wallets(models.WalletWhere.AccountID.EQ(accountID)).One(ctx, pg.Db)
}

func (pg PgDb) GetDeposits(ctx context.Context, accountID string, offset, limit int) ([]*models.Deposit, int64, error) {
	deposits, err := models.Deposits(
		models.DepositWhere.AccountID.EQ(accountID),
		qm.OrderBy(models.DepositColumns.Date+" desc"),
		qm.Limit(limit), qm.Offset(offset),
	).All(ctx, pg.Db)

	if err != nil {
		return nil, 0, err
	}

	totalCount, err := models.Deposits(models.DepositWhere.AccountID.EQ(accountID)).Count(ctx, pg.Db)

	if err != nil {
		return nil, 0, err
	}

	return deposits, totalCount, nil
}

func (pg PgDb) CreditAccount(ctx context.Context, accountID string, amount, date int64, ref string) error {
	tx, err := pg.Db.Begin()
	if err != nil {
		return err
	}
	if err := pg.CreditAccountTx(ctx, tx, accountID, amount, date, ref); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (pg PgDb) CreditAccountTx(ctx context.Context, tx *sql.Tx, accountID string, amount, date int64, ref string) error {
	transaction := models.AccountTransaction{
		AccountID:   accountID,
		Amount:      amount,
		TXType:      app.TxTypeCredit,
		Date:        date,
		Description: ref,
	}

	if err := transaction.Insert(ctx, tx, boil.Infer()); err != nil {
		return err
	}

	return nil
}

func (pg PgDb) DebitAccountTx(ctx context.Context, tx *sql.Tx, accountID string, amount, date int64, ref string) error {
	acc, err := pg.GetAccount(ctx, accountID)
	if err != nil {
		return err
	}
	if acc.Balance < amount {
		return errors.New("insufficient balance")
	}
	transaction := models.AccountTransaction{
		AccountID:   accountID,
		Amount:      amount,
		TXType:      app.TxTypeDebit,
		Date:        date,
		Description: ref,
	}

	if err := transaction.Insert(ctx, tx, boil.Infer()); err != nil {
		return err
	}

	return nil
}

func (pg PgDb) DebitAccount(ctx context.Context, accountID string, amount, date int64, ref string) error {
	tx, err := pg.Db.Begin()
	if err != nil {
		return err
	}
	err = pg.DebitAccountTx(ctx, tx, accountID, amount, date, ref)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (pg PgDb) AccountBalance(ctx context.Context, accountId string) (int64, error) {
	var statement = `SELECT 
	SUM(amount) AS balance FROM (
		SELECT
			CASE WHEN tx.tx_type = 'credit' THEN tx.amount ELSE -1 * tx.amount END AS amount 
		FROM account_transaction tx
		WHERE tx.account_id = $1
	) res`

	var result null.Int64
	err := pg.Db.QueryRow(statement, accountId).Scan(&result)
	if err != nil && err.Error() == sql.ErrNoRows.Error() {
		return 0, nil
	}
	if err != nil {
		return 0, fmt.Errorf("pg.Db.QueryRow %v", err)
	}
	return result.Int64, err
}

func (pg PgDb) MyDownlines(ctx context.Context, accountID string, generation int64, offset, limit int) ([]app.DownlineInfo, int64, error) {
	account, err := models.FindAccount(ctx, pg.Db, accountID)
	if err != nil {
		return nil, 0, err
	}
	query := []qm.QueryMod{}
	switch generation {
	case 1:
		query = append(query, models.AccountWhere.ReferralID.EQ(null.StringFrom(account.ReferralCode)))
	case 2:
		query = append(query, models.AccountWhere.ReferralID2.EQ(null.StringFrom(account.ReferralCode)))
	case 3:
		query = append(query, models.AccountWhere.ReferralID3.EQ(null.StringFrom(account.ReferralCode)))
	}

	totalCount, err := models.Accounts(query...).Count(ctx, pg.Db)
	if err != nil {
		return nil, 0, err
	}

	query = append(query, qm.Load(models.AccountRels.Deposits),
		qm.OrderBy(models.AccountColumns.CreatedAt+" desc"),
		qm.Offset(offset),
		qm.Limit(limit))

	accounts, err := models.Accounts(
		query...,
	).All(ctx, pg.Db)

	if err != nil {
		return nil, 0, err
	}

	var downlines []app.DownlineInfo
	for _, acc := range accounts {
		downline := app.DownlineInfo{
			ID:          acc.ID,
			Username:    acc.Username,
			FirstName:   acc.FirstName,
			LastName:    acc.LastName,
			PhoneNumber: acc.PhoneNumber,
			Date:        acc.CreatedAt,
		}
		downlines = append(downlines, downline)
	}

	return downlines, totalCount, nil
}

func (pg PgDb) AddReferralEarning(ctx context.Context, tx *sql.Tx, accountID string, amount int64) error {
	statement := `update account set bonus = bonus + $1 where id = $2`
	_, err := models.Accounts(qm.SQL(statement, amount, accountID)).ExecContext(ctx, tx)
	return err
}

func (pg PgDb) ReduceReferralEarning(ctx context.Context, tx *sql.Tx, accountID string, amount int64) error {
	statement := `update account set bonus = bonus - $1 where id = $2`
	_, err := models.Accounts(qm.SQL(statement, amount, accountID)).ExecContext(ctx, tx)
	return err
}

func (pg PgDb) AddStableNaira(ctx context.Context, tx *sql.Tx, accountID string, amount int64) error {
	statement := `update account set stable_naira = stable_naira + $1 where id = $2`
	_, err := models.Accounts(qm.SQL(statement, amount, accountID)).ExecContext(ctx, tx)
	return err
}

func (pg PgDb) SpendStableNaira(ctx context.Context, tx *sql.Tx, accountID string, amount int64) error {
	sender, err := models.FindAccount(ctx, tx, accountID, models.AccountColumns.StableNaira)
	if err != nil {
		return err
	}
	if sender.StableNaira < amount {
		return fmt.Errorf("insufficient balance")
	}
	statement := `update account set stable_naira = stable_naira - $1 where id = $2`
	_, err = models.Accounts(qm.SQL(statement, amount, accountID)).ExecContext(ctx, tx)
	return err
}

func (pg PgDb) TransferStableNaira(ctx context.Context, tx *sql.Tx,
	senderAccountID, receiverAccountID string, amount int64) error {
	if err := pg.SpendStableNaira(ctx, tx, senderAccountID, amount); err != nil {
		return err
	}

	return pg.AddStableNaira(ctx, tx, receiverAccountID, amount)
}
