#!groovy

def app = null

pipeline {
    agent any
    
    environment {
        DISABLE_AUTH = 'true'
        DB_ENGINE    = 'sqlite'
    }

    stages {
            stage('Checkout SCM') {
                steps{
                   checkout scm
                }
            }
            stage('Build Image') {
               steps{
                   sh 'echo "test"'
                   app = docker.build("test-development")
               }
            }
    }
}
