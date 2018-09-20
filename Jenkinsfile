pipeline {
  agent {
    docker {
      image 'golang'
      label 'docker'
    }
  }
  stages {
    stage('build') {
      steps {
        sh 'pwd'
        sh 'ls'
      }
    }
  }
}
