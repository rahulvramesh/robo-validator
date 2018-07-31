#!groovy

pipeline {
    agent any
    
    environment {
        DISABLE_AUTH = 'true'
        DB_ENGINE    = 'sqlite'
        app = {}
    }

    stages {
           
            stage('Build') {
                steps {
                    sh 'printenv'
                }
            }
    }
}
