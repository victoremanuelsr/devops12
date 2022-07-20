pipeline{
  agent any
  stages{
    stage('Bake a docker image with packer'){
      steps{
        sh 'packer build packer.json'
      }
    }
  }
}