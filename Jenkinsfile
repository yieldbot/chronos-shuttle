node {
  stage 'Test Go'
  echo 'Testing Go'
  setToolEnv('go', '1.5')
  sh 'go get github.com/golang/lint/golint'
  sh 'go get golang.org/x/tools/cmd/vet'
  sh 'go get -t -v ./...'
  sh 'OUT=`gofmt -l .`; if [ "$OUT" ]; then echo $OUT; exit 1; fi'
  sh 'OUT=`golint ./...`; if [ "$OUT" ]; then echo $OUT; exit 1; fi'
  sh 'go vet ./...'
}
