//
//  Practicing GraphQL
//
//  Copyright © 2020. All rights reserved.
//

package api_test

import (
	"github.com/moemoe89/go-graphql-gendhis/config"
	"github.com/moemoe89/go-graphql-gendhis/routers"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	lang, _ := config.InitLang()
	log := config.InitLog()

	router := routers.GetRouter(lang, log, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
