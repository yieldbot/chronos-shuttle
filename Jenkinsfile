node {
  stage 'Tooling'
  setToolEnv('go', '1.5')
}

node {
  stage 'Build'
  sh 'go get github.com/golang/lint/golint'
  sh 'go get golang.org/x/tools/cmd/vet'
  sh 'go get -t -v ./...'
  sh 'OUT=`gofmt -l .`; if [ "$OUT" ]; then echo $OUT; exit 1; fi'
  sh 'OUT=`golint ./...`; if [ "$OUT" ]; then echo $OUT; exit 1; fi'
  sh 'go vet ./...'
}

def setToolEnv(toolName, toolVer) {
  def toolPath = tool toolName + '-' + toolVer
  def envPath = env.PATH
  // https://issues.jenkins-ci.org/browse/JENKINS-33511
  sh 'pwd > pwd.current'
  env.WORKSPACE = readFile('pwd.current')

  if(toolName == 'go') {
    env.GOROOT = toolPath + '/bin'
    env.GOPATH = env.WORKSPACE
  }
  env.PATH = toolPath + '/bin:' + envPath
}
