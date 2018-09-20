pipeline {
  agent {
    docker {
      image 'golang'
    }
  }
  stages {
    stage('build') {
      steps {
        sh 'go build'
      }
    }
    stage('test') {
      steps {
        sh 'go test'
      }
    }
    stage('deploy test') {
      steps {
        sh 'echo deploying'
      }
    }
  }
}