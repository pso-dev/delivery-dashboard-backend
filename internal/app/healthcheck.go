package app

import "net/http"

func (app *application) handleHealthcheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		env := envelope{
			"status": "available",
			"system_info": map[string]string{
				"environment": app.cfg.Env,
				"version":     app.cfg.Version,
			},
		}

		err := app.writeJSON(w, http.StatusOK, env, nil)
		if err != nil {
			app.serverErrorResponse(w, r, err)
		}
	}
}
