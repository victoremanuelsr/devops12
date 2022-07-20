pipeline{
  agent any
  stages{
    stage('Deploy microservice docker'){
      steps{
        sh 'docker run -p 8082:8082 victoremanuelsr/calculatorgo:1.0'
      }
    }
  }
}