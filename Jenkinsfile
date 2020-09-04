pipeline {
  agent any
  stages {
    stage('Preparation') {
      steps {
        echo 'Starting the Preparation step...'
      }
    }
    stage('Run tests') {
      steps {
        sh 'go test -v ./toposort'
        catchError() {
          cleanWs(cleanWhenSuccess: true)
        }
      }
    }
  }
}