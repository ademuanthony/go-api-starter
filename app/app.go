package app

import (
	"net/http"

	"deficonnect/sonicflare/web"
)

type module struct {
	server *web.Server
	db     store

	MgDomain string
	MgKey    string
}

const (
	TxTypeCredit = "credit"
	TxTypeDebit  = "debit"
)

func Start(
	server *web.Server,
	db store,
	mgDomain, mgKey string) error {
	log.Info("starting...")

	app := module{
		db:       db,
		MgDomain: mgDomain,
		MgKey:    mgKey,
		server:   server,
	}

	app.buildRoute()

	return nil
}

func (m module) buildRoute() {
	m.server.AddRoute("/", web.GET, welcome)

	//ACCOUNT
	m.server.AddRoute("/api/account/update", web.POST, m.UpdateAccountDetail, m.server.RequireLogin, m.server.NoReentry)
	m.server.AddRoute("/api/account/me", web.GET, m.GetAccountDetail, m.server.RequireLogin)
	m.server.AddRoute("/api/account/referral-link", web.GET, m.referralLink, m.server.RequireLogin)
	m.server.AddRoute("/api/account/referral-count", web.GET, m.GetReferralCount, m.server.RequireLogin)
	m.server.AddRoute("/api/account/downlines", web.GET, m.MyDownlines, m.server.RequireLogin)
	m.server.AddRoute("/api/account/team-info", web.GET, m.TeamInformation, m.server.RequireLogin)
	m.server.AddRoute("/api/account/get-by-email", web.GET, wrapHandler(m.getUserByEmail))

	// ACCOUNTS
	m.server.AddRoute("/api/accounts/count", web.GET, m.GetAllAccountsCount, m.server.RequireLogin)
	m.server.AddRoute("/api/accounts/list", web.GET, m.GetAllAccounts, m.server.RequireLogin)
}

func welcome(w http.ResponseWriter, r *http.Request) {
	web.SendJSON(w, "welcome to sonicflare api. download the app from app store to start earning")
}
