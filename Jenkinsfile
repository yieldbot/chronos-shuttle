node {
  def utils = load 'utils.groovy'
  utils.setEnv('go', '1.6')

  stage 'Build'
  sh '''
    if [ ! -d src/github.com/yieldbot/chronos-shuttle/ ]; then
      git clone --depth 1 https://github.com/yieldbot/chronos-shuttle.git src/github.com/yieldbot/chronos-shuttle
    fi
    cd src/github.com/yieldbot/chronos-shuttle
    git pull
    go get -t -v ./...
    go get github.com/golang/lint/golint
    OUT=`gofmt -l .`; if [ "$OUT" ]; then echo $OUT; exit 1; fi
    OUT=`golint ./...`; if [ "$OUT" ]; then echo $OUT; exit 1; fi
    go vet ./...
  '''

  stage 'Publish'
  sh '''
    PACKAGE_NAME="chronos-shuttle" # read from git
    PACKAGE_VERSION="1.2.3" # read from git
    PACKAGE_FILE="$PACKAGE_NAME-$PACKAGE_VERSION-linux-amd64.tar.gz"
    tar -cvzf $PACKAGE_FILE bin/$PACKAGE_NAME
    # not implemented yet
    echo "jfrog rt u $PACKAGE_FILE yieldbot-golang/$PACKAGE_NAME/$PACKAGE_VERSION/ --url=https://artifactory.yb0t.cc/artifactory --dry-run"
  '''

  stage 'Deploy'
  sh '''
    # not implemented yet
  '''
}