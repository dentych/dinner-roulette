pipeline {
  agent {
    docker 'golang'
  }
  stages {
    stage('build') {
      steps {
        sh 'go build'
      }
    }
    stage('test') {
      steps {
        sh 'ls'
      }
    }
  }
}
