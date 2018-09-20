pipeline {
  agent {
    docker {
      image 'golang'
      args '-v /root:/work -w /work'
      label 'docker'
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
