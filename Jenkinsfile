pipeline {
  agent {
    docker 'golang'
  }
  stages {
    stage('build') {
      steps {
        sh 'go build'
        stash name: 'binary', includes: 'main'
      }
    }
    stage('test') {
      steps {
        unstash 'binary'
        sh 'ls'
      }
    }
  }
}
