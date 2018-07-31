#!groovy

pipeline {
    agent any
    def app
    
    environment {
        DISABLE_AUTH = 'true'
        DB_ENGINE    = 'sqlite'
    }

    stages {
            stage('Build') {
                steps {
                    sh 'printenv'
                }
            }
    }
}
