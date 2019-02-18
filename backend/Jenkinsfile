pipeline {
  agent {
    docker 'golang'
  }
  stages {
    stage('build') {
      steps {
        sh 'go build -v'
      }
    }
    stage('test') {
      steps {
        sh 'go test'
      }
    }
  }
}
