#!groovy

def var = null

pipeline {
    agent any
    
    environment {
        DISABLE_AUTH = 'true'
        DB_ENGINE    = 'sqlite'
    }

    stages {
            stage('Checkout SCM') {
                steps {
                   checkout scm
                }
            }
         stage('Build Image') {
                steps {
                   app = docker.build("test-development")
                }
            }
    }
}
