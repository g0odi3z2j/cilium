package middleware

import (
	"bytes"
	"fmt"
	"net/http"
	"path"
	"text/template"
)

func SwaggerUIOAuth2Callback(opts SwaggerUIOpts, next http.Handler) http.Handler {
	opts.EnsureDefaults()

	pth := opts.OAuthCallbackURL
	tmpl := template.Must(template.New("swaggeroauth").Parse(swaggerOAuthTemplate))

	buf := bytes.NewBuffer(nil)
	_ = tmpl.Execute(buf, &opts)
	b := buf.Bytes()

	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if path.Join(r.URL.Path) == pth {
			rw.Header().Set("Content-Type", "text/html; charset=utf-8")
			rw.WriteHeader(http.StatusOK)

			_, _ = rw.Write(b)
			return
		}

		if next == nil {
			rw.Header().Set("Content-Type", "text/plain")
			rw.WriteHeader(http.StatusNotFound)
			_, _ = rw.Write([]byte(fmt.Sprintf("%q not found", pth)))
			return
		}
		next.ServeHTTP(rw, r)
	})
}

const (
	swaggerOAuthTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <title>{{ .Title }}</title>
</head>
<body>
<script>
    'use strict';
    function run () {
        var oauth2 = window.opener.swaggerUIRedirectOauth2;
        var sentState = oauth2.state;
        var redirectUrl = oauth2.redirectUrl;
        var isValid, qp, arr;

        if (/code|token|error/.test(window.location.hash)) {
            qp = window.location.hash.substring(1).replace('?', '&');
        } else {
            qp = location.search.substring(1);
        }

        arr = qp.split("&");
        arr.forEach(function (v,i,_arr) { _arr[i] = '"' + v.replace('=', '":"') + '"';});
        qp = qp ? JSON.parse('{' + arr.join() + '}',
                function (key, value) {
                    return key === "" ? value : decodeURIComponent(value);
                }
        ) : {};

        isValid = qp.state === sentState;

        if ((
          oauth2.auth.schema.get("flow") === "accessCode" ||
          oauth2.auth.schema.get("flow") === "authorizationCode" ||
          oauth2.auth.schema.get("flow") === "authorization_code"
        ) && !oauth2.auth.code) {
            if (!isValid) {
                oauth2.errCb({
                    authId: oauth2.auth.name,
                    source: "auth",
                    level: "warning",
                    message: "Authorization may be unsafe, passed state was changed in server. The passed state wasn't returned from auth server."
                });
            }

            if (qp.code) {
                delete oauth2.state;
                oauth2.auth.code = qp.code;
                oauth2.callback({auth: oauth2.auth, redirectUrl: redirectUrl});
            } else {
                let oauthErrorMsg;
                if (qp.error) {
                    oauthErrorMsg = "["+qp.error+"]: " +
                        (qp.error_description ? qp.error_description+ ". " : "no accessCode received from the server. ") +
                        (qp.error_uri ? "More info: "+qp.error_uri : "");
                }

                oauth2.errCb({
                    authId: oauth2.auth.name,
                    source: "auth",
                    level: "error",
                    message: oauthErrorMsg || "[Authorization failed]: no accessCode received from the server."
                });
            }
        } else {
            oauth2.callback({auth: oauth2.auth, token: qp, isValid: isValid, redirectUrl: redirectUrl});
        }
        window.close();
    }

    if (document.readyState !== 'loading') {
        run();
    } else {
        document.addEventListener('DOMContentLoaded', function () {
            run();
        });
    }
</script>
</body>
</html>
`
)
