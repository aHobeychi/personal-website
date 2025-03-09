npx nodemon \
  --watch "**" \
  --ext "go,html,js,json,css" \
  --signal SIGTERM \
  --exec "go run ${PWD}/main.go"
