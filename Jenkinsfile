pipeline {
  agent {
    docker 'golang:alpine'
  }
  stages {
    stage('build') {
      steps {
        sh 'go build'
      }
    }
  }
}
