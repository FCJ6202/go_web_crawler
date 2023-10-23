package retry

import (
	"fmt"
	"net/http"
)

func RetryPage(w http.ResponseWriter, URL string) {
	// Display a retry page with a form to enter the URL again
	fmt.Fprint(w, `
		<html>
		<head>
			<title>Retry Page</title>
		</head>
		<body>
			<h1>Retry</h1>
			<p>There was an issue processing your request. Please try again.</p>
			<form method="post">
				<input type="text" name="url" value="`+URL+`">
				<input type="submit" value="Retry">
			</form>
		</body>
		</html>
	`)
}
