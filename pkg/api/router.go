/****************************************************************************
 * Copyright 2019, Optimizely, Inc. and contributors                        *
 *                                                                          *
 * Licensed under the Apache License, Version 2.0 (the "License");          *
 * you may not use this file except in compliance with the License.         *
 * You may obtain a copy of the License at                                  *
 *                                                                          *
 *    http://www.apache.org/licenses/LICENSE-2.0                            *
 *                                                                          *
 * Unless required by applicable law or agreed to in writing, software      *
 * distributed under the License is distributed on an "AS IS" BASIS,        *
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. *
 * See the License for the specific language governing permissions and      *
 * limitations under the License.                                           *
 ***************************************************************************/
package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/optimizely/sidedoor/pkg/api/handlers"
)

// NewRouter returns HTTP API router
func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	// r.Post("/features/activate", ActivateAllFeatures)
	r.Post("/features/{featureKey}/activate", handlers.ActivateFeature)
	// r.Post("/track", handlers.TrackEvent)
	return r
}
