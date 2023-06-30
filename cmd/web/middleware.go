package main

//func WriteToConsole(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
//		fmt.Println("Hit the page")
//		next.ServeHTTP(writer, request)
//	})
//}

//func NoSurf(next http.Handler) http.Handler {
//	csrfHandler := nosurf.New(next)
//	csrfHandler.SetBaseCookie(http.Cookie{
//		HttpOnly: true,
//		Path:     "/",
//		Secure:   false,
//		SameSite: http.SameSiteLaxMode,
//	})
//	return csrfHandler
//}

//func SessionLoad(next http.Handler) http.Handler {
//	return session.LoadAndSave(next)
//}
