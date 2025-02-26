npx nodemon \
  --watch "**" \
  --ext "go,html,js,json" \
  --signal SIGTERM \
  --exec "go run ${PWD}/main.go"
