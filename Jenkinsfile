pipeline {
  agent {
    label 'docker'
  }
  stages {
    stage('build') {
      steps {
        sh 'pwd'
        sh 'ls'
        sh 'docker run --rm -v $(pwd):/work -w /work golang go build'
      }
    }
  }
}
