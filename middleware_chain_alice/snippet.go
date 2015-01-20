chain := alice.New(middlewareOne, middlewareTwo)
http.Handle("/", chain.Then(finalHandler))