noAuthChain := alice.New(contextMiddleware, loggerMiddleware)
authChain := alice.New(contextMiddleware, loggerMiddleware, apiKeyAuthMiddleware)
adminChain := alice.New(contextMiddleware, loggerMiddleware, adminAuthMiddleware)
