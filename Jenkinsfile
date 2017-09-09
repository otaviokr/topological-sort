pipeline {
  agent any
  stages {
    stage('Preparation') {
      steps {
        echo 'Starting the Preparation step...'
        isUnix()
        catchError() {
          echo 'Erro???'
        }
        
        echo 'Continuing...'
      }
    }
  }
}