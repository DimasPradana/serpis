package main

/**
 * @author snub on Mon 19 Apr 2021 09:54:09
 * /home/snub/go/src/github.com/DimasPradana/kantor/serpis/main.go
 */

/** {{{ TODO snub on Mon 19 Apr 2021 09:54:17 template todo
	᚛ mutation tambahTodo {
		  createTodo(input:{text:"yang ke-1", userId:"1"}) {
			user {
			  id
			}
			text
			done
		  }
	}

	᚛ query cariTodos {
		todos {
			id
			text
			done
			user {
				name
			}
		}
	}
}}} */

/** {{{ TODO snub on Mon 19 Apr 2021 09:59:00 what to do 😀
	᚛ pakai env ✓
	᚛ koneksi ke database
		• oracle ✓
		• mysql
	᚛ all queries ✓
	᚛ single query → find single by nm_login ✓
}}}*/

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/DimasPradana/kantor/serpis/config"
	"github.com/DimasPradana/kantor/serpis/graph"
	"github.com/DimasPradana/kantor/serpis/graph/generated"
)

const defaultPort = "8888"

func init() {
	config.GetEnvfile()
}

func main() {
	port := os.Getenv("gqlgenport")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
