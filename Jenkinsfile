pipeline {
  agent {
    docker 'golang'
  }
  stages {
    stage('build') {
      steps {
        sh 'go build'
        sh 'ls'
      }
    }
    stage('test') {
      steps {
        sh 'ls'
        sh 'go test'
      }
    }
  }
}
