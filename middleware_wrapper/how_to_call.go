var logger *AppLogger = NewLogger()
loggerMiddleware := simpleLoggerMiddlewareWrapper(logger) // HL
http.Handle("/", loggerMiddleware(http.HandlerFunc(final)))