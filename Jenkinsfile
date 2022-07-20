pipeline{
  agent any
  stages{
    stage('Deploy microservice docker'){
      steps{
        sh 'docker run -p 8080:8080 victoremanuelsr/calculatorgo:1.0'
      }
    }
  }
}