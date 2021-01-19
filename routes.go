package main

func (app *App) SetupRoutes() {
	app.router.HandleFunc("/ping", HandlePingRequest()).Methods("GET")

}
