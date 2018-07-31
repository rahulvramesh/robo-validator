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
              
                   checkout scm
                
            }
            stage('Build Image') {
               
                   app = docker.build("test-development")
                
            }
    }
}
