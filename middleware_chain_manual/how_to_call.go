http.Handle("/", middlewareOne(middlewareTwo(http.HandlerFunc(final))))