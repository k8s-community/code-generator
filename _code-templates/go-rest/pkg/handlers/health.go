// Copyright 2017 Kubernetes Community Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package handlers

import (
	"net/http"

	"{[( .ProjectPath )]}/pkg/router"
)

// Health returns "OK" if service is alive
func (h *Handler) Health(c router.Control) {
	c.Code(http.StatusOK)
	c.Write(http.StatusText(http.StatusOK))
}