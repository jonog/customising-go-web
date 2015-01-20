func ShowWidget(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamIdStr := vars["team_id"]
	widgetIdStr := vars["widget_id"]
    ...
}