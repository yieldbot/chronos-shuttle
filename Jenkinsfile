node {
  stage 'Build'
  def utils = load 'utils.groovy'
  utils.setEnv('go', '1.6')
  sh '''
    rm -rf src/github.com/yieldbot/chronos-shuttle
    git clone --depth 1 https://github.com/yieldbot/chronos-shuttle.git src/github.com/yieldbot/chronos-shuttle
    cd src/github.com/yieldbot/chronos-shuttle
    go get -t -v ./...
    go get github.com/golang/lint/golint
    OUT=`gofmt -l .`; if [ "$OUT" ]; then echo $OUT; exit 1; fi
    OUT=`golint ./...`; if [ "$OUT" ]; then echo $OUT; exit 1; fi
    go vet ./...
  '''
}
